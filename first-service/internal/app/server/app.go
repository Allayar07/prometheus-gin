package server

import (
	"gin_prometheus/internal/app/handler"
	"gin_prometheus/internal/repository"
	"gin_prometheus/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func Init(port string) {
	//db, err := repository.NewPostgresDB(context.Background())
	//if err != nil {
	//	log.Fatal(err)
	//}

	repos := repository.NewRepository(&pgxpool.Pool{})
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(Server)

	if err := srv.Run(":"+port, handlers.InitRoutes()); err != nil {
		log.Fatalln(err)
	}

}
