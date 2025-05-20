package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.elastic.co/apm/module/apmhttp/v2"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, APM!\n")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if os.Getenv("ELASTIC_APM_SERVICE_NAME") == "" {
		os.Setenv("ELASTIC_APM_SERVICE_NAME", "my-sample-go-app") // Set a default if not provided
	}
	if os.Getenv("ELASTIC_APM_SERVER_URL") == "" {
		os.Setenv("ELASTIC_APM_SERVER_URL", "http://localhost:8200") // Set a default if not provided
	}

	http.Handle("/hello", apmhttp.Wrap(http.HandlerFunc(helloHandler)))

	fmt.Println("Server listening on port 8810...")
	fmt.Printf("Connect to APM by setting ELASTIC_APM_SERVICE_NAME (current: %s) and ELASTIC_APM_SERVER_URL (current: %s)\n", os.Getenv("ELASTIC_APM_SERVICE_NAME"), os.Getenv("ELASTIC_APM_SERVER_URL"))
	fmt.Println("Send requests to http://localhost:8810/hello to see traces in Kibana APM.")

	if err := http.ListenAndServe(":8810", nil); err != nil {
		log.Fatal(err)
	}
}
