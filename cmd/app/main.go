package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/joho/godotenv/autoload"

	"github.com/kewka/go-app-template/pkg/postgres"
	"github.com/kewka/go-app-template/pkg/service"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", "4000", "http server port")
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
	svc := service.New(dbpool)

	r := chi.NewRouter()
	r.Mount("/api", svc.HTTPHandler())
	r.Get("/*", http.FileServer(http.Dir("web/build")).ServeHTTP)
	log.Printf("http server running on %v", addr)
	return http.ListenAndServe(addr, r)
}
