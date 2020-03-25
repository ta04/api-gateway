package main

import (
	"context"
	"encoding/json"
	"github.com/SleepingNext/SKit/api-gateway/helper"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/web"
	"log"
	"net/http"
	"strings"
)

func main() {
	s := web.NewService(
		web.Name("com.ta04.web.skit"),
	)

	// Initialize the service
	s.Init()

	s.HandleFunc("/product/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		client := helper.NewClient()

		// Call IndexProducts rpc from grpc client
		res, err := client.IndexProducts(context.Background(), &productPB.IndexProductsRequest{})
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
	})

	s.HandleFunc("/product/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
		// Take params from request
		vars := r.URL.Query()
		product := helper.ParseShowAndDeleteQuery(vars)

		client := helper.NewClient()

		// Call ShowProduct rpc from grpc client
		res, err := client.ShowProduct(context.Background(), product)
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
	})

	s.HandleFunc("/product/store", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
		// Take params from request
		vars := r.URL.Query()
		product := helper.ParseStoreQuery(vars)

		client := helper.NewClient()

		// Get the authorization header
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// Create context with token
		ctx := metadata.NewContext(context.Background(), map[string]string{
			"token": token,
		})

		// Call StoreProduct rpc from grpc client
		res, err := client.StoreProduct(ctx, product)
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
	})

	s.HandleFunc("/product/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
		// Take params from request
		vars := r.URL.Query()
		product := helper.ParseUpdateQuery(vars)

		client := helper.NewClient()


		// Get the authorization header
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// Create context with token
		ctx := metadata.NewContext(context.Background(), map[string]string{
			"token": token,
		})


		// Call UpdateProduct rpc from grpc client
		res, err := client.UpdateProduct(ctx, product)
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
	})

	s.HandleFunc("/product/destroy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
		// Take params from request
		vars := r.URL.Query()
		product := helper.ParseShowAndDeleteQuery(vars)

		client := helper.NewClient()

		// Get the authorization header
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			http.Error(w, "Invalid token", http.StatusBadRequest)
			return
		}
		token := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		// Create context with token
		ctx := metadata.NewContext(context.Background(), map[string]string{
			"token": token,
		})

		// Call DestroyProduct rpc from grpc client
		res, err := client.DestroyProduct(ctx, product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return			}

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
	})

	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}