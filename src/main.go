package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoPrivResponse struct {
	Number int `json:"number"`
}

func main() {
	fmt.Println("Hi, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := http.Get("http://pod-service-go-priv.default.svc.cluster.local:80")
		if err != nil {
			log.Fatalf("Error sending request: %+v", err)
		}

		responseData, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Error reading request: %+v", err)
		}

		var response GoPrivResponse
		err = json.Unmarshal(responseData, &response)
		if err != nil {
			log.Fatalf("Error unmarshalling request: %+v", err)
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Hello, World! My number is %d"}`, response.Number)
	})

	log.Fatalf("Error: %+v", http.ListenAndServe("0.0.0.0:8080", nil))
}
