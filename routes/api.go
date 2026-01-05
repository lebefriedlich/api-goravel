package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers"
	"goravel/app/http/middleware"
)

func Api() {
	userController := controllers.NewUserController()

	// Public routes
	facades.Route().Prefix("/api").Group(func(r route.Router) {
		r.Post("/login", userController.Login)
		r.Post("/register", userController.Register)
	})

	// Protected routes
	facades.Route().Prefix("/api").Middleware(middleware.Auth()).Group(func(r route.Router) {
		r.Post("/logout", userController.Logout)
		
		r.Get("/users", userController.Index)
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
