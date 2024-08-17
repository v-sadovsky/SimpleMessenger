package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./swaggerui"))
	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui/", fs))

	fmt.Println("Swagger UI listening on :8084")
	if err := http.ListenAndServe(":8084", nil); err != nil {
		fmt.Println("Error starting Swagger UI server:", err)
	}
}
