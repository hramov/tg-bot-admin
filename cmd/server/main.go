package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/hramov/tg-bot-admin/docs"
	"github.com/hramov/tg-bot-admin/internal/adapters/api/middlewares"
	"github.com/hramov/tg-bot-admin/internal/adapters/db"
	db2 "github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/adapters/db/migrations"
	"github.com/hramov/tg-bot-admin/internal/composite"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/hramov/tg-bot-admin/pkg/mail"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func initRouter(logger *logging.Logger) *chi.Mux {
	logger.Info("create router")
	router := chi.NewRouter()

	router.Use(middlewares.ReqId)
	router.Use(middlewares.Log)

	return router
}

func initPostgres(cfg *config.Config, logger *logging.Logger) db2.Connector {
	logger.Info("create postgres connection")
	pg, err := db.DatabaseFactory(db.Postgres, cfg.Storage)
	if err != nil {
		logger.Fatal("cannot start postgres: %v", err)
		os.Exit(1)
	}

	logger.Info("init postgres migrations")
	migrations.Init(pg, logger)
	err = migrations.Start()
	if err != nil {
		logger.Errorf("postgres migrations error: %s", err.Error())
	}

	return pg
}

func initSwagger(logger *logging.Logger, router *chi.Mux) {
	logger.Info("swagger docs initializing")
	router.Handle("/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Get("/swagger/*", httpSwagger.WrapHandler)
}

func initModules(logger *logging.Logger, router *chi.Mux, cfg *config.Config, pg db2.Connector) {
	logger.Info("initializing user module")
	var userComposite composite.Composite = &composite.UserComposite{}
	userComposite.Register(logger, cfg, pg, router)

	logger.Info("initializing order module")
	var orderComposite composite.Composite = &composite.OrderComposite{}
	orderComposite.Register(logger, cfg, pg, router)

	logger.Info("initializing product module")
	var productComposite composite.Composite = &composite.ProductComposite{}
	productComposite.Register(logger, cfg, pg, router)
}

func start(router *chi.Mux, cfg *config.Config, logger *logging.Logger) {
	logger.Info("start application")
	var listener net.Listener
	var listenErr error
	if cfg.Listen.Type == config.UnixSockType {
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, cfg.Listen.SockPath)
		logger.Info("listen unix socket")
		listener, listenErr = net.Listen(config.UnixListener, socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen(config.TcpListener, fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	c := cors.New(cors.Options{
		AllowedMethods:     cfg.Cors.AllowedMethods,
		AllowedOrigins:     cfg.Cors.AllowedOrigins,
		AllowCredentials:   cfg.Cors.AllowCredentials,
		AllowedHeaders:     cfg.Cors.AllowedHeaders,
		OptionsPassthrough: cfg.Cors.OptionsPassthrough,
		ExposedHeaders:     cfg.Cors.ExposedHeaders,
		Debug:              cfg.Cors.Debug,
	})

	handler := c.Handler(router)

	server := &http.Server{
		Handler:      handler,
		WriteTimeout: cfg.Listen.WriteTimeout,
		ReadTimeout:  cfg.Listen.ReadTimeout,
	}

	logger.Fatal(server.Serve(listener))
}

func main() {
	cfg := config.GetConfig()
	logging.Init(cfg.Logger)
	logger := logging.GetLogger()
	mail.New(cfg.Mail)
	pg := initPostgres(cfg, logger)
	router := initRouter(logger)
	initSwagger(logger, router)
	initModules(logger, router, cfg, pg)
	start(router, cfg, logger)
}
