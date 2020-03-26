package main

import (
	"context"
	"encoding/json"
	userPB "github.com/G0tYou/user-service/proto"
	"github.com/SleepingNext/api-gateway/helper"
	orderPB "github.com/SleepingNext/order-service/proto"
	"log"

	//authPB "github.com/SleepingNext/auth-service/proto"
	//paymentPB "github.com/SleepingNext/payment-service/proto"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/web"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	s := web.NewService(
		web.Name("com.ta04.web.skit"),
	)

	// Initialize the service
	s.Init()

	// Product APIs
	s.HandleFunc("/product/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		client := helper.NewProductClient()

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
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var product *productPB.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewProductClient()

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
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var product *productPB.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewProductClient()

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
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var product *productPB.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewProductClient()


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
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var product *productPB.Product
		err = json.Unmarshal(body, &product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewProductClient()

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

	// User APIs
	s.HandleFunc("/user/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		client := helper.NewUserClient()

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

		// Call IndexUsers rpc from grpc client
		res, err := client.IndexUsers(ctx, &userPB.IndexUsersRequest{})
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

	s.HandleFunc("/user/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var user *userPB.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewUserClient()

		// Call ShowUser rpc from grpc client
		res, err := client.ShowUser(context.Background(), user)
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

	s.HandleFunc("/user/showByUsername", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var user *userPB.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewUserClient()

		// Call ShowUserByUsername( rpc from grpc client
		res, err := client.ShowUserByUsername(context.Background(), user)
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

	s.HandleFunc("/user/store", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var user *userPB.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewUserClient()

		// Call StoreUser rpc from grpc client
		res, err := client.StoreUser(context.Background(), user)
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

	s.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var user *userPB.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewUserClient()


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


		// Call UpdateUser rpc from grpc client
		res, err := client.UpdateUser(ctx, user)
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

	s.HandleFunc("/user/destroy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var user *userPB.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewUserClient()

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

		// Call DestroyUser rpc from grpc client
		res, err := client.DestroyUser(ctx, user)
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

	// Order APIs
	s.HandleFunc("/order/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		client := helper.NewOrderClient()

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

		// Call IndexOrders rpc from grpc client
		res, err := client.IndexOrders(ctx, &orderPB.IndexOrdersRequest{})
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

	s.HandleFunc("/order/indexByUserID", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var user *orderPB.User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewOrderClient()

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

		// Call IndexOrders rpc from grpc client
		res, err := client.IndexOrdersByUserID(ctx, user)
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

	s.HandleFunc("/order/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var order *orderPB.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewOrderClient()

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

		// Call ShowOrder rpc from grpc client
		res, err := client.ShowOrder(ctx, order)
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

	s.HandleFunc("/order/store", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var order *orderPB.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewOrderClient()

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

		// Call StoreOrder rpc from grpc client
		res, err := client.StoreOrder(ctx, order)
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

	s.HandleFunc("/order/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PUT" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var order *orderPB.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewOrderClient()

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

		// Call UpdateOrder rpc from grpc client
		res, err := client.UpdateOrder(ctx, order)
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

	s.HandleFunc("/order/destroy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}

		// Take params from request
		body, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Unmarshal the body
		var order *orderPB.Order
		err = json.Unmarshal(body, &order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		client := helper.NewOrderClient()

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

		// Call DestroyOrder rpc from grpc client
		res, err := client.DestroyOrder(ctx, order)
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
