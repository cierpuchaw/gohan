package autogen

// AUTO GENERATED CODE DO NOT MODIFY MANUALLY
import (
	"net/http/httptest"

	"github.com/cloudwan/gohan/extension/gohanscript"
	"github.com/cloudwan/gohan/extension/gohanscript/lib"
)

func init() {

	gohanscript.RegisterStmtParser("get_test_server_url",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			stmtErr := stmt.HasArgs(
				"server")
			if stmtErr != nil {
				return nil, stmtErr
			}
			return func(context *gohanscript.Context) (interface{}, error) {

				server := stmt.Arg("server", context).(*httptest.Server)

				result1 :=
					lib.GetTestServerURL(
						server)

				return result1, nil

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("GetTestServerURL",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			server := args[0].(*httptest.Server)

			result1 :=
				lib.GetTestServerURL(
					server)
			return []interface{}{
				result1}

		})

	gohanscript.RegisterStmtParser("stop_test_server",
		func(stmt *gohanscript.Stmt) (func(*gohanscript.Context) (interface{}, error), error) {
			stmtErr := stmt.HasArgs(
				"server")
			if stmtErr != nil {
				return nil, stmtErr
			}
			return func(context *gohanscript.Context) (interface{}, error) {

				server := stmt.Arg("server", context).(*httptest.Server)

				lib.StopTestServer(
					server)
				return nil, nil

			}, nil
		})
	gohanscript.RegisterMiniGoFunc("StopTestServer",
		func(vm *gohanscript.VM, args []interface{}) []interface{} {

			server := args[0].(*httptest.Server)

			lib.StopTestServer(
				server)
			return nil

		})

}