package router

import (
	"app/src/controller"
	m "app/src/middleware"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func HomeRoutes(v1 fiber.Router, u service.UserService, p service.ProductTokenService, ml service.MealService) {
	mealController := controller.NewMealController(ml)

	home := v1.Group("/home")

	home.Get("/statistic", m.Auth(u, p), mealController.GetHomeStatistics)
}
