package main

import (
	"log"

	export "github.com/siangyeh8818/golang.exporter.tencentcloud/internal/server"
)

func main() {
	log.Println("Exporter is start ro running")

	export.Run_Exporter_Server()
}
