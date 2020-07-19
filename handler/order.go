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
	"github.com/ta04/api-gateway/middleware"
	"io/ioutil"
	"net/http"

	"github.com/micro/go-micro/web"
	"github.com/ta04/api-gateway/client"
	"github.com/ta04/api-gateway/helper"
	proto "github.com/ta04/order-service/model/proto"
)

// HandleOrder handles all the requests to order APIs
func HandleOrder(s web.Service) {
	orderSC := client.NewOrderSC()

	s.Handle("/order/index", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.GetAllOrdersRequest
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := orderSC.GetAllOrders(r.Context(), request)
			if res == nil {
				http.Error(w, "no orders returned", http.StatusInternalServerError)
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

	s.Handle("/order/show", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.GetOneOrderRequest
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := orderSC.GetOneOrder(r.Context(), request)
			if res == nil {
				http.Error(w, "no order returned", http.StatusInternalServerError)
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

	s.Handle("/order/store", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var order *proto.Order
			err = json.Unmarshal(body, &order)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := orderSC.CreateOneOrder(r.Context(), order)
			if res == nil {
				http.Error(w, "no order returned", http.StatusInternalServerError)
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

	s.Handle("/order/update", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var order *proto.Order
			err = json.Unmarshal(body, &order)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := orderSC.UpdateOneOrder(r.Context(), order)
			if res == nil {
				http.Error(w, "no order returned", http.StatusInternalServerError)
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
