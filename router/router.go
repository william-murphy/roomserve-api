package router

import (
	"net/http"

	"roomserve/handlers"
	"roomserve/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initialize(router *fiber.App) {
	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Hello, World!")
	})

	router.Use(middleware.Json)

	// auth
	authGroup := router.Group("/auth")
	authGroup.Post("/login", handlers.Login)

	// user
	userGroup := router.Group("/users")
	userGroup.Post("/", handlers.CreateUser)

	// building
	buildingGroup := router.Group("/buildings")
	buildingGroup.Post("/", handlers.CreateBuilding)
	buildingGroup.Get("/", handlers.GetBuildings)
	buildingGroup.Get("/:id", handlers.GetBuilding)

	// floor
	floorGroup := router.Group("/floors")
	floorGroup.Post("/", handlers.CreateFloor)
	floorGroup.Get("/", handlers.GetFloors)
	floorGroup.Get("/:id", handlers.GetFloor)

	// room
	roomGroup := router.Group("/rooms")
	roomGroup.Post("/", handlers.CreateRoom)
	roomGroup.Get("/", handlers.GetRooms)
	roomGroup.Get("/:id", handlers.GetRoom)

	// reservation
	reservationGroup := router.Group("/reservations")
	reservationGroup.Post("/", handlers.CreateReservation)
	reservationGroup.Get("/", handlers.GetReservations)
	reservationGroup.Get("/:id", handlers.GetReservation)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).SendString("Not Found")
	})
}
