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

package helper

import (
	userPB "github.com/G0tYou/user-service/proto"
	authPB "github.com/SleepingNext/auth-service/proto"
	orderPB "github.com/SleepingNext/order-service/proto"
	paymentPB "github.com/SleepingNext/payment-service/proto"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/micro/go-micro"
)

// NewProductSC creates a new product service client
func NewProductSC() productPB.ProductServiceClient {
	s := micro.NewService(
		micro.Name("com.ta04.api.product"),
	)
	s.Init()

	productServiceClient := productPB.NewProductServiceClient("com.ta04.srv.product", s.Client())
	return productServiceClient
}

// NewOrderSC creates a new order service client
func NewOrderSC() orderPB.OrderServiceClient {
	s := micro.NewService(
		micro.Name("com.ta04.api.order"),
	)
	s.Init()

	orderServiceClient := orderPB.NewOrderServiceClient("com.ta04.srv.order", s.Client())
	return orderServiceClient
}

// NewUserSC creates a new user service client
func NewUserSC() userPB.UserServiceClient {
	s := micro.NewService(
		micro.Name("com.ta04.api.user"),
	)
	s.Init()

	userServiceClient := userPB.NewUserServiceClient("com.ta04.srv.user", s.Client())
	return userServiceClient
}

// NewPaymentSC creates a new payment service client
func NewPaymentSC() paymentPB.PaymentServiceClient {
	s := micro.NewService(
		micro.Name("com.ta04.api.payment"),
	)
	s.Init()

	paymentServiceClient := paymentPB.NewPaymentServiceClient("com.ta04.srv.payment", s.Client())
	return paymentServiceClient
}

// NewAuthSC creates a new auth service client
func NewAuthSC() authPB.AuthServiceClient {
	s := micro.NewService(
		micro.Name("com.ta04.api.auth"),
	)
	s.Init()

	authServiceClient := authPB.NewAuthServiceClient("com.ta04.srv.auth", s.Client())
	return authServiceClient
}
