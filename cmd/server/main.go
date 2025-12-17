package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"backend_dev_task/config"
	"backend_dev_task/db/sqlc"
	"backend_dev_task/internal/handler"
	"backend_dev_task/internal/logger"
	"backend_dev_task/internal/middleware"
	"backend_dev_task/internal/repository"
	"backend_dev_task/internal/routes"
	"backend_dev_task/internal/service"
)

func main() {
	logger.Init()
	defer logger.Sync()

	cfg := config.Load()

	db, err := sql.Open("pgx", cfg.DBURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Run migrations
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("Failed to create migrate instance:", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Failed to run migrations:", err)
	}

	queries := sqlc.New(db)
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.Logger())

	routes.SetupRoutes(app, userHandler)

	log.Fatal(app.Listen(":" + cfg.Port))
}
