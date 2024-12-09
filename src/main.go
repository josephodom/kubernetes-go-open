package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type GoPrivResponse struct {
	Number int `json:"number"`
}

func main() {
	fmt.Println("Hi, World!")

	client := redis.NewClient(&redis.Options{
		Addr:     "pod-service-redis.default.svc.cluster.local:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	ctx := context.Background()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numberStr, err := client.Get(ctx, "myNumber").Result()
		if err != nil {
			numberStr = ""
		}

		numberInt, err := strconv.Atoi(numberStr)
		if err != nil {
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

			numberInt = response.Number

			err = client.Set(ctx, "myNumber", numberInt, time.Duration(time.Second*10)).Err()
			if err != nil {
				log.Printf("Error setting to redis: %+v", err)
			}
		}

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Hey, World! My number is %d"}`, numberInt)
	})

	log.Fatalf("Error: %+v", http.ListenAndServe("0.0.0.0:8080", nil))
}
