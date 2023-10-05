package server

import (
	http "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/neo532/kratos_layout/internal/service/api"

	demo "github.com/neo532/kratos_layout/proto/api/demo/v1"
)

type Router struct {
}

// InitHTTPRouter register HTTP router.
func InitHTTPRouter(srv *http.Server,
	dm *api.DemoApi,
) (r *Router) {

	// router
	demo.RegisterDemoHTTPServer(srv, dm)

	return
}
