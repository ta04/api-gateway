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

package main

import (
	"log"

	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"github.com/ta04/api-gateway/config"
	"github.com/ta04/api-gateway/handler"
)

func main() {
	name := config.MicroWebName()
	port := config.MicroWebPort()

	registry := consul.NewRegistry()

	s := web.NewService(
		web.Name(name),
		web.Address(port),
		web.Registry(registry),
	)
	s.Init()

	// Define handlers that'll handle all requests to all APIs
	handler.HandleProduct(s)
	handler.HandleOrder(s)
	handler.HandleUser(s)
	handler.HandlePayment(s)
	handler.HandleAuth(s)

	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
