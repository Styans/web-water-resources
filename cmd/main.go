package main

import (
	"flag"
	"os"
	"test/configs"
	"test/internal/app"
	"test/internal/handlers"
	"test/internal/render"
	"test/internal/repository"
	"test/internal/service"
	"test/pkg/client/sqlite"

	"log"
)

func main() {
	log.Println("wait a minute...")

	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	cfg, err := configs.GetConfig(*configPath)
	if err != nil {
		log.Println(err)
		return
	}

	db, err := sqlite.OpenDB(cfg.DB.DSN)
	if err != nil {
		log.Println(err)
		return
	}

	repo := repository.NewRepository(db)
	file, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Error opening log file: ", err)
	}
	logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	// info := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)

	service := service.NewService(repo, logger)

	template, err := render.NewTemplateHTML(cfg.TemplateDir)
	if err != nil {
		log.Println(err)
		return
	}

	handler := handlers.NewHandler(service, template)

	err = app.Server(cfg, handler.Routes())

	if err != nil {

		log.Println("Ooopss...\n", err)
		return
	}
}
