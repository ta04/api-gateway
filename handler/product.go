/*
Dear Programmers,

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*                                                 *
*	This file belongs to Kevin Veros Hamonangan   *
*	and	Fandi Fladimir Dachi and is a part of     *
*	our	last project as the student of Del        *
*	Institute of Technology, Sitoluama.           *
*	Please contact us via Instagram:              *
*	sleepingnext and fandi_dachi                  *
*	before copying this file.                     *
*	Thank you, buddy. 😊                          *
*                                                 *
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
*/

package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/micro/go-micro/web"
	"github.com/ta04/api-gateway/client"
	"github.com/ta04/api-gateway/helper"
	"github.com/ta04/api-gateway/middleware"
	proto "github.com/ta04/product-service/model/proto"
)

// HandleProduct handles all the requests to product APIs
func HandleProduct(s web.Service) {
	productSC := client.NewProductSC()

	s.HandleFunc("/product/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.GetAllProductsRequest
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := productSC.GetAllProducts(r.Context(), request)
			if res == nil {
				http.Error(w, "no products returned", http.StatusInternalServerError)
				return
			}
			if err != nil {
				http.Error(w, res.Error.Message, int(res.Error.Code))
				return
			}

			marshaledRes, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaledRes)
			return
		} else if r.Method == "OPTIONS" {
			helper.SetAccessControlHeader(w)
			return
		}
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	})

	s.HandleFunc("/product/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.GetOneProductRequest
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := productSC.GetOneProduct(r.Context(), request)
			if res == nil {
				http.Error(w, "no product returned", http.StatusInternalServerError)
				return
			}
			if err != nil {
				http.Error(w, res.Error.Message, int(res.Error.Code))
				return
			}

			marshaledRes, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaledRes)
			return
		} else if r.Method == "OPTIONS" {
			helper.SetAccessControlHeader(w)
			return
		}
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	})

	s.Handle("/product/store", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var product *proto.Product
			err = json.Unmarshal(body, &product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := productSC.CreateOneProduct(r.Context(), product)
			if res == nil {
				http.Error(w, "no product returned", http.StatusInternalServerError)
				return
			}
			if err != nil {
				http.Error(w, res.Error.Message, int(res.Error.Code))
				return
			}

			marshaledRes, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaledRes)
			return
		} else if r.Method == "OPTIONS" {
			helper.SetAccessControlHeader(w)
			return
		}
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	})))

	s.Handle("/product/update", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var product *proto.Product
			err = json.Unmarshal(body, &product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := productSC.UpdateOneProduct(r.Context(), product)
			if res == nil {
				http.Error(w, "no product returned", http.StatusInternalServerError)
				return
			}
			if err != nil {
				http.Error(w, res.Error.Message, int(res.Error.Code))
				return
			}

			marshaledRes, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaledRes)
			return
		} else if r.Method == "OPTIONS" {
			helper.SetAccessControlHeader(w)
			return
		}
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	})))
}
