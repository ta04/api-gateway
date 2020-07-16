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
proto "github.com/ta04/payment-method-service/model/proto"
)

// HandlePaymentMethod handles all the requests to payment method APIs
func HandlePaymentMethod(s web.Service) {
	paymentMethodSC := client.NewPaymentMethodSC()

	s.HandleFunc("/payment-method/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.GetAllPaymentMethodsRequest
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentMethodSC.GetAllPaymentMethods(r.Context(), request)
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

	s.HandleFunc("/payment-method/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.GetOnePaymentMethodRequest
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentMethodSC.GetOnePaymentMethod(r.Context(), request)
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

	s.Handle("/payment-method/store", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var paymentMethod *proto.PaymentMethod
			err = json.Unmarshal(body, &paymentMethod)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentMethodSC.CreateOnePaymentMethod(r.Context(), paymentMethod)
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

	s.Handle("/payment-method/update", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var paymentMethod *proto.PaymentMethod
			err = json.Unmarshal(body, &paymentMethod)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentMethodSC.UpdateOnePaymentMethod(r.Context(), paymentMethod)
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
