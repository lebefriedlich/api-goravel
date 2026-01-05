package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
)

func Api() {
	facades.Route().Prefix("/api").Group(func(r route.Router) {
		userController := controllers.NewUserController()
		r.Get("/users", userController.Index)
		r.Post("/users", userController.Register)
		r.Get("/users/{id}", userController.Show)
		r.Post("/users/{id}", userController.Update)
		r.Delete("/users/{id}", userController.Destroy)

		r.Get("/books", controllers.NewBookController().Index)
		r.Post("/books", controllers.NewBookController().Store)
		r.Get("/books/{id}", controllers.NewBookController().Show)
		r.Post("/books/{id}", controllers.NewBookController().Update)
		r.Delete("/books/{id}", controllers.NewBookController().Destroy)
	})
}
