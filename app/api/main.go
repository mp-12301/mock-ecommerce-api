package main

import (
	"flag"
	"fmt"
)

type config struct {
	port int
	env string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "environment", "development", "Environment {development|staging|production}")
	
	flag.Parse()
	fmt.Println(cfg)


}
