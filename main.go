package main

import (
	"fmt"
	"github.com/sritama/task-app-db/api"
)

func main() {
	fmt.Println("Starting a new server")
	api.Start()
}
