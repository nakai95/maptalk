package main

import (
	"log"
	"maptalk/internal/infrastructure/router"
)

func main() {
    e := router.NewRouter()
    log.Fatal(e.Start(":8080"))
}

