package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/zap"

	"service/config"
	projectCmds "service/internal/app/command/project"
	requirementCmds "service/internal/app/command/requirement"
	specificationCmds "service/internal/app/command/specification"
	userCmds "service/internal/app/command/user"
	"service/internal/infrastructure/postgres/project"
	"service/internal/infrastructure/postgres/requirement"
	"service/internal/infrastructure/postgres/specification"
	"service/internal/infrastructure/postgres/user"
	projectHandler "service/internal/infrastructure/server/http/fiber/project"
	requirementHandler "service/internal/infrastructure/server/http/fiber/requirement"
	specificationHandler "service/internal/infrastructure/server/http/fiber/specification"
	userHandler "service/internal/infrastructure/server/http/fiber/user"
	"service/pkg/storage"
	"service/pkg/utils"
)

func main() {
	l, err := utils.InitJSONLogger()
	if err != nil {
		panic(err)
	}

	defer l.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		l.Fatal("error reading config file", zap.Error(err))
	}

	l.Info("Config loaded", zap.Any("config", cfg))

	db, err := storage.InitPostgres(&cfg.Postgres)
	if err != nil {
		l.Fatal("failed to connect to PostgreSQL", zap.Error(err))
	}

	projectRepo := project.NewProjectRepository(db)
	requirementRepo := requirement.NewRequirementRepository(db)
	specificationRepo := specification.NewSpecificationRepository(db)
	userRepo := user.NewUserRepository(db)

	projectCmdCreate := projectCmds.NewCreateProjectCmdHandler(projectRepo, userRepo)
	projectCmdUpdate := projectCmds.NewUpdateProjectCmdHandler(projectRepo)
	projectCmdDelete := projectCmds.NewDeleteProjectCmdHandler(projectRepo)
	projectCmdAll := projectCmds.NewGetAllProjectsCmdHandler(projectRepo)

	requirementCmdCreate := requirementCmds.NewCreateRequirementCmdHandler(
		requirementRepo,
		projectRepo,
		userRepo,
	)
	requirementCmdUpdate := requirementCmds.NewUpdateRequirementCmdHandler(
		requirementRepo,
		specificationRepo,
		userRepo,
	)
	requirementCmdDelete := requirementCmds.NewDeleteRequirementCmdHandler(requirementRepo)
	requirementAddInSpecCmd := requirementCmds.NewAddInSpecificationCmdHandler(
		requirementRepo,
		specificationRepo,
	)
	requirementCmdGetSpec := requirementCmds.NewGetSpenRequirementsCmdHandler(
		requirementRepo,
		specificationRepo,
	)
	requirementCmdGetProjectReq := requirementCmds.NewGetProjectRepuirementsCmdHandler(
		requirementRepo,
		projectRepo,
	)

	specificationCmdCreate := specificationCmds.NewCreateSpecificationCmdHandler(
		specificationRepo,
		projectRepo,
	)
	specificationCmdUpdate := specificationCmds.NewUpdateSpecificationCmdHandler(specificationRepo)
	specificationCmdDelete := specificationCmds.NewDeleteSpecificationCmdHandler(specificationRepo)

	userCmdGet := userCmds.NewGetUserCmdHandler(userRepo)
	userCmdReg := userCmds.NewCreateUserCmdHandler(userRepo)
	userCmdGetAll := userCmds.NewGetAllUsersCmdHandler(userRepo)

	requirementHandler := requirementHandler.NewRequirementHandler(
		requirementCmdCreate,
		requirementCmdUpdate,
		requirementCmdDelete,
		requirementAddInSpecCmd,
		requirementCmdGetProjectReq,
		requirementCmdGetSpec,
	)

	projectHandler := projectHandler.NewProjectHandler(
		projectCmdCreate,
		projectCmdUpdate,
		projectCmdDelete,
		projectCmdAll,
	)

	specHandler := specificationHandler.NewSpecificationHandler(
		specificationCmdCreate,
		specificationCmdUpdate,
		specificationCmdDelete,
	)

	userHandler := userHandler.NewUserHandler(userCmdGet, userCmdReg, userCmdGetAll)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		// AllowOrigins: []{"*"},                          // Разрешает запросы с любого источника
		AllowMethods: []string{"GET,POST,OPTIONS"},           // Разрешённые методы
		AllowHeaders: []string{"Content-Type,Authorization"}, // Разрешённые заголовки
	}))

	app.Use(logger.New())
	app.Use(recover.New())

	projectHandler.Map(app)
	requirementHandler.Map(app)
	specHandler.Map(app)
	userHandler.Map(app)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		l.Info("Server started", zap.String("port", cfg.Server.Port))
		if err := app.Listen(":" + cfg.Server.Port); err != nil {
			l.Fatal("failed to start server", zap.Error(err))
		}
	}()

	<-quit
	l.Info("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		l.Fatal("Server forced to shutdown", zap.Error(err))
	}

	l.Info("Server exited gracefully")
}
