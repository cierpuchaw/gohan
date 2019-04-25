package cli

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/cloudwan/gohan/schema"
	"github.com/cloudwan/gohan/util"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

//TemplateDef template def defines file
type TemplateDef struct {
	Type         string `yaml:"type"`          // All, ResourceGroup, Resource
	Package      string `yaml:"package"`       // Code generation directory. Default "gen".
	OutputPath   string `yaml:"output_path"`   // This could be template string. "dir/{{Resource.ID}}".go
	TemplatePath string `yaml:"template_path"` // Path to template
	Language     string `yaml:"language"`      // Type of programming language. Eg: go(Default), typescript
}

func getGenerateCommand() cli.Command {
	return cli.Command{
		Name:      "generate",
		ShortName: "gen",
		Usage:     "Generate ServerSide Code",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "templates", Value: "", Usage: "Template configuration"},
			cli.StringFlag{Name: "config-file, c", Value: "./gohan.yaml", Usage: "Gohan config file"},
			cli.StringFlag{Name: "language, l", Value: "go", Usage: "Programming language"},
			cli.StringFlag{Name: "output, o", Value: ".", Usage: "Dir of output"},
			cli.StringFlag{Name: "package, p", Value: "gen", Usage: "Package Name"},
			cli.StringFlag{Name: "dbname, d", Value: "gohan", Usage: "DB Name"},
			cli.BoolFlag{Name: "resetdb", Usage: "Reset Database on create"},
		},
		Action: gohanGenerate,
	}
}

func ensureDatabase(dbName string) error {
	config := util.GetConfig()
	databaseType := config.GetString("database/type", "mysql")
	databaseConnection := os.Getenv("DATABASE_CONNECTION")
	if databaseConnection == "" {
		databaseConnection = config.GetString("database/connection", "")
	}
	db, err := sql.Open(databaseType, databaseConnection)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("drop database if exists " + dbName)
	if err != nil {
		return err
	}
	_, err = db.Exec("create database " + dbName)
	if err != nil {
		return err
	}
	return nil
}

func dropDatabase(dbName string) error {
	config := util.GetConfig()
	databaseType := config.GetString("database/type", "mysql")
	databaseConnection := os.Getenv("DATABASE_CONNECTION")
	if databaseConnection == "" {
		databaseConnection = config.GetString("database/connection", "")
	}
	db, err := sql.Open(databaseType, databaseConnection)
	if err != nil {
		return nil
	}
	defer db.Close()
	_, err = db.Exec("drop database " + dbName)
	if err != nil {
		return nil
	}
	return nil
}

func gohanGenerate(c *cli.Context) {
	path := c.String("output")
	templates := c.String("templates")
	configFile := c.String("config-file")
	packageName := c.String("package")
	language := c.String("language")
	dbName := c.String("dbname")
	codeDir := filepath.Join(path, packageName)
	etcDir := filepath.Join(path, "etc")
	dbDir := filepath.Join(etcDir, "db")
	resetDB := c.Bool("resetdb")
	migrationDir := filepath.Join(dbDir, "migrations")
	schemaPath := filepath.Join(etcDir, "schema.json")
	manager := schema.GetManager()
	config := util.GetConfig()
	config.ReadConfig(configFile)
	schemaFiles := config.GetStringList("schemas", nil)
	if schemaFiles == nil {
		log.Fatal("No schema specified in configuration")
		return
	}
	if err := manager.LoadSchemasFromFiles(schemaFiles...); err != nil {
		log.Fatal(err)
		return
	}
	// Genrating schema json
	log.Info("Genrating: schema json")

	list := []interface{}{}

	for _, schema := range manager.OrderedSchemas() {
		if schema.IsAbstract() {
			continue
		}
		if schema.Metadata["type"] == "metaschema" {
			continue
		}
		s := schema.JSON()
		s["url"] = schema.URL
		list = append(list, s)
	}
	os.Mkdir(etcDir, 0777)
	os.Mkdir(dbDir, 0777)
	os.Mkdir(migrationDir, 0777)
	execCommand(fmt.Sprintf("rm %s/*_init_schema.sql", migrationDir))
	if resetDB {
		err := ensureDatabase(dbName)
		if err != nil {
			log.Error("Failed to reset database", err)
		}
	}
	execCommand(
		fmt.Sprintf(
			"gohan migrate init --config-file %s", configFile))
	execCommand(
		fmt.Sprintf(
			"gohan migrate up --config-file %s", configFile))
	// Running sqlboiler
	execCommand("sqlboiler mysql")

	util.SaveFile(schemaPath, map[string]interface{}{
		"schemas": list,
	})
	execCommand(
		fmt.Sprintf(
			"go-bindata -pkg %s -o %s/go-bindata.go %s", packageName, codeDir, schemaPath))

	//Generating application code
	log.Info("Generating: application code")
	if templates != "" {
		templateConfig := []*TemplateDef{}
		templateDir := filepath.Dir(templates)
		data, err := ioutil.ReadFile(templates)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		err = yaml.Unmarshal([]byte(data), &templateConfig)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for _, templateDef := range templateConfig {
			flag := ""
			if templateDef.Package != "" {
				packageName = templateDef.Package
				codeDir = filepath.Join(path, packageName)
				language = templateDef.Language
			}
			templatePath := filepath.Join(templateDir, templateDef.TemplatePath)
			switch templateDef.Type {
			case "resource":
				flag = "--split-by-resource true"
			case "group":
				flag = "--split-by-resource-group true"
			}
			execCommand(
				fmt.Sprintf(
					"gohan template --config-file %s --template %s %s --output-path %s",
					configFile, templatePath,
					flag, filepath.Join(codeDir, templateDef.OutputPath),
				))
			if language == "go" {
				execCommand(
					fmt.Sprintf(
						"goimports -w %s/*.go", codeDir))
			}
		}
	}
}

func execCommand(command string) {
	output, err := exec.Command("sh", "-c", command).Output()
	log.Info("Running: %s", command)
	outputStr := string(output[:])
	if outputStr != "" {
		log.Info("Output: %s", outputStr)
	}
	if err != nil {
		log.Error("Error: %s %s", err, output)
	}
}
