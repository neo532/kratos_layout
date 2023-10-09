package v1

import (
	"context"
	"fmt"
	"runtime"

	"github.com/neo532/apitool/transport/http/xhttp/middleware"
)

func Demo() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(c context.Context, req, reply interface{}) (err error) {

			fmt.Println(runtime.Caller(0))

			return handler(c, req, reply)
		}
	}
}
