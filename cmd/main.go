package main

import (
	"log/slog"
	"os"
	"url-shortener-golang/pkg/config"
	"url-shortener-golang/pkg/repository"

	"github.com/jmoiron/sqlx"
)
const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)
func main(){
    // init config cleanEnv
    cfg := config.MustLoad()

    // init logger slog
    log := setupLogger(cfg.Env)
    log.Info("Starting server...")

    // init storage postgres
    db, err := initPostgresConnection(cfg) 
    if err != nil {
        log.Error("initConfig error", slog.String("error", err.Error()))
      
    } 
    repos := repository.NewRepository(db)
	// services := servece.

    // init router chi

    // run server
}


func setupLogger(env string) *slog.Logger{
	// create logger with slog
	var log *slog.Logger

	switch env{
	case envLocal: 
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return log
}

func initPostgresConnection(cfg *config.Config) (*sqlx.DB, error) {
    db, err := repository.NewPostgresDB(repository.Config{
        Host:     cfg.DB.Host,
        Port:     cfg.DB.Port,
        Username: cfg.DB.Username,
        DBName:   cfg.DB.Dbname,
        SSLMode:  cfg.DB.Sslmode,
        Password: os.Getenv("DB_PASSWORD"),
    })

    return db, err
}
