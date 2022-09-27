package main

import (
	"flag"
	"os"

	"github.com/mp-12301/mock-ecommerce-api/shared/logger"
)

type config struct {
	port int
	env  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "environment", "development", "Environment {development|staging|production}")

	flag.Parse()

	logger := logger.New(os.Stdout)

	logger.Print("hello world")
}
