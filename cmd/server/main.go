package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/arulkarim/ngodingyuk-server/internal/config"
	"github.com/arulkarim/ngodingyuk-server/internal/domain"
	"github.com/arulkarim/ngodingyuk-server/internal/handler"
	"github.com/arulkarim/ngodingyuk-server/internal/middleware"
	"github.com/arulkarim/ngodingyuk-server/internal/repository"
	"github.com/arulkarim/ngodingyuk-server/internal/service"
	jwtpkg "github.com/arulkarim/ngodingyuk-server/pkg/jwt"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize JWT
	jwtpkg.SetSecret(cfg.JWTSecret)

	// Initialize DB (GORM)
	db := config.InitDB(cfg.DatabaseURL)

	// Get raw *sql.DB for SQL sandbox
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get raw SQL DB: ", err)
	}
	// We don't defer sqlDB.Close() here because GORM manages the connection pool.
	// But we need a SEPARATE connection for sandbox to avoid search_path conflicts.
	sandboxDB, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to create sandbox DB connection: ", err)
	}
	defer sandboxDB.Close()
	_ = sqlDB // suppress unused warning; GORM manages this

	// Auto-migrate new domain models only if DB_MIGRATE is true
	if os.Getenv("DB_MIGRATE") == "true" {
		log.Println("Running database migrations...")
		if err := db.AutoMigrate(
			&domain.User{},
			&domain.Challenge{},
			&domain.Course{},
			&domain.Lesson{},
			&domain.LessonQuiz{},
			&domain.UserChallengeProgress{},
			&domain.UserLessonProgress{},
			&domain.UserXPHistory{},
			&domain.UserCourseEnrollment{},
			&domain.Certificate{},
		); err != nil {
			log.Fatal("Failed to run migrations: ", err)
		}

		// Create unique constraints
		db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_user_challenge_unique ON user_challenge_progress(user_id, challenge_id)")
		db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_user_lesson_unique ON user_lesson_progress(user_id, lesson_id)")
		db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_user_course_enroll_unique ON user_course_enrollments(user_id, course_id)")
		db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS idx_certificate_unique ON certificates(user_id, course_id)")

		log.Println("Migrations completed successfully")
	} else {
		log.Println("Skipping migrations (DB_MIGRATE != true)")
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	challengeRepo := repository.NewChallengeRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	progressRepo := repository.NewProgressRepository(db)

	// Initialize services
	authSvc := service.NewAuthService(userRepo)
	executeSvc := service.NewExecuteService(sandboxDB)
	challengeSvc := service.NewChallengeService(challengeRepo, progressRepo, userRepo, executeSvc)
	courseSvc := service.NewCourseService(courseRepo, progressRepo, userRepo)
	leaderboardSvc := service.NewLeaderboardService(progressRepo)
	userSvc := service.NewUserService(userRepo, progressRepo, courseRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authSvc)
	challengeHandler := handler.NewChallengeHandler(challengeSvc)
	courseHandler := handler.NewCourseHandler(courseSvc)
	leaderboardHandler := handler.NewLeaderboardHandler(leaderboardSvc)
	userHandler := handler.NewUserHandler(userSvc)
	adminHandler := handler.NewAdminHandler(challengeRepo, courseRepo)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "NgodingYuk API",
	})

	// Global middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin,Content-Type,Accept,Authorization",
	}))

	// Health check
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "ngodingyuk-api",
		})
	})

	api := app.Group("/api")

	// === Auth Routes (Public) ===
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/refresh", authHandler.Refresh)
	auth.Post("/logout", authHandler.Logout)

	// Auth - Protected
	authProtected := auth.Group("", middleware.AuthMiddleware())
	authProtected.Get("/me", authHandler.GetProfile)

	// === Challenge Routes ===
	challenges := api.Group("/challenges")
	challenges.Get("/", middleware.OptionalAuthMiddleware(), challengeHandler.ListChallenges)
	challenges.Get("/:slug", challengeHandler.GetChallenge)
	challenges.Post("/:slug/run", challengeHandler.RunChallenge)

	// Challenge - Protected
	challengesProtected := challenges.Group("", middleware.AuthMiddleware())
	challengesProtected.Post("/:slug/submit", challengeHandler.SubmitChallenge)
	challengesProtected.Get("/:slug/my-progress", challengeHandler.GetMyProgress)

	// === Course Routes ===
	courses := api.Group("/courses")
	courses.Get("/", courseHandler.ListCourses)
	courses.Get("/:slug", middleware.OptionalAuthMiddleware(), courseHandler.GetCourse)
	courses.Get("/:slug/lessons/:lesson_id", middleware.OptionalAuthMiddleware(), courseHandler.GetLessonDetail)

	// Course - Protected
	coursesProtected := courses.Group("", middleware.AuthMiddleware())
	coursesProtected.Post("/:slug/enroll", courseHandler.EnrollCourse)
	coursesProtected.Post("/:slug/lessons/:lesson_id/complete", courseHandler.CompleteLesson)
	coursesProtected.Post("/:slug/lessons/:lesson_id/quiz", courseHandler.SubmitQuiz)
	coursesProtected.Get("/:slug/exam", courseHandler.GetExam)
	coursesProtected.Post("/:slug/exam/submit", courseHandler.SubmitExam)
	coursesProtected.Get("/:slug/my-progress", courseHandler.GetMyProgress)

	// === Leaderboard Routes (Public) ===
	leaderboard := api.Group("/leaderboard")
	leaderboard.Get("/weekly", leaderboardHandler.GetWeekly)
	leaderboard.Get("/all-time", leaderboardHandler.GetAllTime)

	// === User Routes (Protected) ===
	users := api.Group("/users", middleware.AuthMiddleware())
	users.Get("/me", userHandler.GetMe)
	users.Get("/me/xp-history", userHandler.GetXPHistory)
	users.Get("/me/challenge-stats", userHandler.GetChallengeStats)
	users.Get("/me/certificates", userHandler.GetCertificates)
	users.Patch("/me", userHandler.UpdateMe)

	// === Admin Routes (Protected) ===
	admin := api.Group("/admin", middleware.AuthMiddleware())
	admin.Post("/challenges", adminHandler.CreateChallenge)
	admin.Put("/challenges/:id", adminHandler.UpdateChallenge)
	admin.Delete("/challenges/:id", adminHandler.DeleteChallenge)
	admin.Post("/courses", adminHandler.CreateCourse)
	admin.Post("/courses/:id/lessons", adminHandler.CreateLesson)
	admin.Put("/lessons/:id", adminHandler.UpdateLesson)
	admin.Delete("/lessons/:id", adminHandler.DeleteLesson)
	admin.Post("/lessons/:id/quizzes", adminHandler.CreateQuiz)

	// Start server
	port := cfg.Port
	log.Printf("NgodingYuk API starting on port %s...", port)
	if err := app.Listen(fmt.Sprintf("0.0.0.0:%s", port)); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
