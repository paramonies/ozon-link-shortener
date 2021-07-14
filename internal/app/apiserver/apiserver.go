package apiserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/paramonies/ozon-link-shortener/internal/app/controller"
	"github.com/paramonies/ozon-link-shortener/internal/app/repository"
	"github.com/paramonies/ozon-link-shortener/internal/app/service"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func Start(config *Config) error {
	//DB init
	db, err := newDB(config)
	if err != nil {
		return err
	}

	//Controller, service, repository create instances
	repo := repository.NewLinkRepository(db)
	service := service.NewLinkService(repo)
	controller := controller.NewController(service)

	//Server start
	srv := new(Server)
	go func() {
		if err := srv.Run(config.SrvPort, controller.InitRoutes()); err != nil {
			log.Fatalf(err.Error())
		}
	}()
	log.Printf("APIServer Started on %s:%s", config.SrvHost, config.SrvPort)

	//Server shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("APIServer Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("Shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("error occured on db connection close: %s", err.Error())
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func newDB(config *Config) (*sqlx.DB, error) {
	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
