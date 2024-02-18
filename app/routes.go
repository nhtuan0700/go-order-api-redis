package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nhtuan0700/orders-api/handler"
	"github.com/nhtuan0700/orders-api/util"
)

func (a *App) loadRoutes(server *http.Server) {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		util.Response(w, http.StatusOK, "health is ok")
	})

	handler.NewOrderHandler(router, a.rdb)

	server.Handler = router
}
