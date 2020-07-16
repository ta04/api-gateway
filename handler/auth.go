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
	proto "github.com/ta04/auth-service/model/proto"
)

// HandleAuth handles all the requests to auth APIs
func HandleAuth(s web.Service) {
	authSC := client.NewAuthSC()
	s.HandleFunc("/auth/auth1", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.Auth1
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := authSC.AuthRPC1(r.Context(), request)
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

	s.HandleFunc("/auth/auth2", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var request *proto.Auth2
			err = json.Unmarshal(body, &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			res, err := authSC.AuthRPC2(r.Context(), request)
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
}
