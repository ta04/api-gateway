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
	"github.com/ta04/api-gateway/helper"
	paymentPB "github.com/ta04/payment-service/proto"
)

// HandlePayment handles all the requests to payment APIs
func HandlePayment(s web.Service) {
	paymentSC := helper.NewPaymentSC()

	s.HandleFunc("/payment/index", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			res, err := paymentSC.IndexPayments(r.Context(), &paymentPB.IndexPaymentsRequest{})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
		}
		http.Error(w, "Unsupported http method", http.StatusBadRequest)
		return
	})

	s.HandleFunc("/payment/indexByUserID", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var user *paymentPB.User
			err = json.Unmarshal(body, &user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentSC.IndexPaymentsByUserID(r.Context(), user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
			return
		} else {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
	})

	s.HandleFunc("/payment/show", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var payment *paymentPB.Payment
			err = json.Unmarshal(body, &payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentSC.ShowPayment(r.Context(), payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
			return
		} else {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
	})

	s.HandleFunc("/payment/store", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var payment *paymentPB.Payment
			err = json.Unmarshal(body, &payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentSC.StorePayment(r.Context(), payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
			return
		} else {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
	})

	s.HandleFunc("/payment/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var payment *paymentPB.Payment
			err = json.Unmarshal(body, &payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentSC.UpdatePayment(r.Context(), payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
			return
		} else {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
	})

	s.HandleFunc("/payment/destroy", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var payment *paymentPB.Payment
			err = json.Unmarshal(body, &payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := paymentSC.DestroyPayment(r.Context(), payment)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
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
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
			return
		} else {
			http.Error(w, "Unsupported http method", http.StatusBadRequest)
			return
		}
	})
}
