package main

import (
	"context"
	"encoding/json"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
	"strconv"
)

func main() {
	s := web.NewService(
		web.Name("com.ta04.web.skit"),
	)

	// Initialize the service
	s.Init()

	s.HandleFunc("/product/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// Create a new service
			s := micro.NewService(
				micro.Name("com.ta04.cli.product"),
			)

			// Initialize the service
			s.Init()

			productServiceClient := productPB.NewProductServiceClient("com.ta04.srv.product", s.Client())

			// Call IndexProducts rpc from grpc client
			res, err := productServiceClient.IndexProducts(context.Background(), &productPB.IndexProductsRequest{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Marshal the response
			js, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set the header and write the marshaled response
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}
	})

	s.HandleFunc("/product/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// Take params from request
			var id int32
			vars := r.URL.Query()
			log.Println(vars)
			ids, exists := vars["id"]
			if !exists || len(ids) != 1 {
				id = 0
			} else {
				id64, _ := strconv.ParseInt(ids[0], 10, 64)
				id = int32(id64)
			}

			// Create a new service
			s := micro.NewService(
				micro.Name("com.ta04.cli.product"),
			)

			// Initialize the service
			s.Init()

			productServiceClient := productPB.NewProductServiceClient("com.ta04.srv.product", s.Client())

			// Call ShowProduct rpc from grpc client
			res, err := productServiceClient.ShowProduct(context.Background(), &productPB.Product{Id: int32(id)})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Marshal the response
			js, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Set the header and write the marshaled response
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
			return
		}
	})

	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}