package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"github.com/kewka/go-app-template/pkg/app"
	"github.com/kewka/go-app-template/pkg/httphandler"
	"github.com/kewka/go-app-template/pkg/postgres"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "4000", "server port")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	addr := fmt.Sprintf(":%v", port)
	postgresCfg, err := postgres.LoadConfig()
	if err != nil {
		return err
	}
	dbpool, err := postgres.NewPool(ctx, postgresCfg)
	if err != nil {
		return err
	}
	defer dbpool.Close()
	log.Printf("http server running on %v", addr)
	return http.ListenAndServe(addr, httphandler.New(
		&httphandler.Deps{
			ItemsService: app.NewItemsService(dbpool),
		},
	))
}
