package main

import (
	"fmt"
	_ "github.com/hramov/tg-bot-admin/docs"
	db2 "github.com/hramov/tg-bot-admin/internal/adapters/db"
	"github.com/hramov/tg-bot-admin/internal/composite"
	"github.com/hramov/tg-bot-admin/internal/config"
	"github.com/hramov/tg-bot-admin/pkg/db"
	"github.com/hramov/tg-bot-admin/pkg/logging"
	"github.com/hramov/tg-bot-admin/pkg/mail"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func initRouter(logger *logging.Logger) *httprouter.Router {
	logger.Info("create router")
	return httprouter.New()
}

func initPostgres(cfg *config.Config, logger *logging.Logger) db2.Connector {
	logger.Info("create postgres connection")
	pg, err := db.DatabaseFactory(db.Postgres, cfg)
	if err != nil {
		logger.Fatal("cannot start postgres: %v", err)
		os.Exit(1)
	}
	return pg
}

func initSwagger(logger *logging.Logger, router *httprouter.Router) {
	logger.Info("swagger docs initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)
}

func initModules(logger *logging.Logger, router *httprouter.Router, cfg *config.Config, pg db2.Connector) {
	logger.Info("initializing user module")
	var userComposite composite.Composite = &composite.UserComposite{}
	userComposite.Register(logger, cfg, pg, router)
}

func start(router *httprouter.Router, cfg *config.Config, logger *logging.Logger) {
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
	logger := logging.GetLogger()
	cfg := config.GetConfig()
	mail.New(cfg.Mail)
	pg := initPostgres(cfg, logger)
	router := initRouter(logger)
	initSwagger(logger, router)
	initModules(logger, router, cfg, pg)
	start(router, cfg, logger)
}
