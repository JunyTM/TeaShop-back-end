package router

import (
	"net/http"
	"teashop/controller"
	"teashop/infrastructure"
	_"teashop/docs"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	//Declare controller
	userController := controller.NewUserController()

	// api swagger for develope mode
	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.GetHTTPSwagger()),
	))

	r.Route("/api/v1", func(router chi.Router) {

		router.Get("/user/{id}", userController.GetById)
		
		router.Group(func(protectedRoute chi.Router) {
			protectedRoute.Route("/user", func(subRoute chi.Router) {
				subRoute.Get("/all", userController.GetAll)
				subRoute.Post("/create", userController.Create)
				subRoute.Put("/update", userController.Update)
				subRoute.Delete("/delete/{uid}", userController.Delete)
			})
		})
	})
	return r
}
