package app

import (
	"log"

	"rapdev-graphql/pkg/config"
	graphqldelivery "rapdev-graphql/pkg/delivery/graphql"
	"rapdev-graphql/pkg/repository"
	"rapdev-graphql/pkg/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp() *fiber.App {
	db := config.NewDatabase()

	articleRepo := repository.NewArticleRepository(db)
	educationRepo := repository.NewEducationRepository(db)
	experienceRepo := repository.NewExperienceRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	galleryRepo := repository.NewGalleryRepository(db)
	socialLinkRepo := repository.NewSocialLinkRepository(db)

	articleUC := usecase.NewArticleUsecase(articleRepo)
	educationUC := usecase.NewEducationUsecase(educationRepo)
	experienceUC := usecase.NewExperienceUsecase(experienceRepo)
	projectUC := usecase.NewProjectUsecase(projectRepo)
	galleryUC := usecase.NewGalleryUsecase(galleryRepo)
	socialLinkUC := usecase.NewSocialLinkUsecase(socialLinkRepo)

	resolver := graphqldelivery.NewResolver(
		articleUC,
		educationUC,
		experienceUC,
		projectUC,
		galleryUC,
		socialLinkUC,
	)

	schema, err := graphqldelivery.NewSchema(resolver)
	if err != nil {
		log.Fatalf("failed to create graphql schema: %v", err)
	}

	fiberApp := fiber.New(fiber.Config{
		AppName:      "Rapdev Portfolio Server",
		ErrorHandler: customErrorHandler,
	})

	fiberApp.Use(recover.New())
	fiberApp.Use(logger.New())
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, OPTIONS",
	}))

	graphqlHandler := graphqldelivery.NewHandler(schema)
	playgroundHandler := graphqldelivery.NewPlaygroundHandler()

	fiberApp.Get("/", playgroundHandler)
	fiberApp.Post("/graphql", graphqlHandler)
	fiberApp.Get("/graphql", graphqlHandler)

	fiberApp.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	return fiberApp
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	return c.Status(code).JSON(fiber.Map{
		"errors": []fiber.Map{{"message": err.Error()}},
	})
}
