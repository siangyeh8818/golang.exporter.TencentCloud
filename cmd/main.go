package main

import (
	"fmt"
	"log"
	"os"

	export "github.com/siangyeh8818/golang.exporter.tencentcloud/internal/server"
)

func main() {
	log.Println("Exporter is start ro running")
	fmt.Printf("TENCENTCLOUD_SECRET_ID : %s\n", os.Getenv("TENCENTCLOUD_SECRET_ID"))
	fmt.Printf("TENCENTCLOUD_SECRET_KEY : %s\n", os.Getenv("TENCENTCLOUD_SECRET_KEY"))
	fmt.Printf("GET_API_INTERNAL_TIME : %s\n", os.Getenv("GET_API_INTERNAL_TIME"))

	go func() {
		export.Handler_API()
	}()

	export.Run_Exporter_Server()
}
