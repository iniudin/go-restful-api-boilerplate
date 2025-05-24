package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/iniudin/boilerplate/internal/platform"
	"github.com/iniudin/boilerplate/pkg/web"

	_ "github.com/iniudin/boilerplate/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @Summary Hello World
// @Description Hello World
// @Tags Hello World
// @Success 200 {object} Response
// @Router / [get]
func Hello(w http.ResponseWriter, r *http.Request) {
	web.JSON(w, http.StatusOK, Response{
		Status:  "success",
		Message: "Hello, World!",
	})
}

func NewRouter(config *platform.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", Hello)

	r.Get("/docs/*", httpSwagger.Handler())

	return r
}
