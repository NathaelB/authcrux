package application

import (
	"encoding/json"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nathaelb/authcrux/application/handler/realm"
	"github.com/nathaelb/authcrux/domain/port"
)

type Services struct {
	realmService port.RealmService
}

// HTTPServer encapsule le serveur HTTP et sa configuration
type HTTPServer struct {
	app      *fiber.App
	isProd   bool
	services Services
	handlers struct {
		realm *realm.RealmHandler
	}
}

// NewHTTPServer crée une nouvelle instance du serveur HTTP
func NewHTTPServer(realmService port.RealmService) *HTTPServer {
	// Déterminer l'environnement (production ou développement)
	isProd := os.Getenv("ENV") == "production"

	// Création de l'application Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: customErrorHandler,
	})

	services := Services{
		realmService: realmService,
	}

	s := &HTTPServer{
		app:      app,
		isProd:   isProd,
		services: services,
	}

	s.initHandlers()
	s.registerRoutes()

	// Middlewares globaux
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: func() string {
			if isProd {
				return `{"time":"${time}","status":${status},"latency":"${latency}","method":"${method}","path":"${path}"}` + "\n"
			}
			return "${time} | ${status} | ${latency} | ${method} ${path}\n"
		}(),
		TimeFormat: "2006-01-02T15:04:05Z07:00",
		TimeZone:   "UTC",
	}))

	// Route de base
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API AuthCrux")
	})

	return s
}

func (s *HTTPServer) initHandlers() {
	s.handlers.realm = realm.NewRealmHandler(s.services.realmService)
}

func (s *HTTPServer) registerRoutes() {
	s.handlers.realm.RegisterRoutes(s.app)

	log.Info("Routes enregistrées avec succès")
}

// App retourne l'instance de l'application Fiber
func (s *HTTPServer) App() *fiber.App {
	return s.app
}

// Start démarre le serveur HTTP
func (s *HTTPServer) Start() error {
	port := getEnv("PORT", "3000")
	log.Info("Démarrage du serveur sur le port ", port)
	return s.app.Listen(":" + port)
}

// customErrorHandler gère les erreurs de manière globale
func customErrorHandler(c *fiber.Ctx, err error) error {
	// Par défaut, on retourne une erreur 500
	code := fiber.StatusInternalServerError

	// Si c'est une erreur Fiber, utiliser son code de statut
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Format de réponse JSON
	log.Info(err.Error())
	// Try to unmarshal the error message if it's in JSON format
	var errorData interface{}
	if err := json.Unmarshal([]byte(err.Error()), &errorData); err == nil {
		// If successful, return the parsed JSON
		return c.Status(code).JSON(fiber.Map{
			"success": false,
			"error":   errorData,
		})
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}

// getEnv récupère une variable d'environnement ou sa valeur par défaut
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
