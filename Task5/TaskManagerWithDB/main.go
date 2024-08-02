package main

import (
	"taskmanagerdb/router"
)

func main() {
	r := router.SetupRouter()

	r.Run(":8080")
}
