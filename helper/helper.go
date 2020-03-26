package helper

import (
	userPB "github.com/G0tYou/user-service/proto"
	authPB "github.com/SleepingNext/auth-service/proto"
	orderPB "github.com/SleepingNext/order-service/proto"
	paymentPB "github.com/SleepingNext/payment-service/proto"
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/micro/go-micro"
)

func NewProductClient() productPB.ProductServiceClient {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.api.product"),
	)

	// Initialize the service
	s.Init()

	productServiceClient := productPB.NewProductServiceClient("com.ta04.srv.product", s.Client())
	return productServiceClient
}

func NewAuthClient() authPB.AuthServiceClient {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.api.auth"),
	)

	// Initialize the service
	s.Init()

	authServiceClient := authPB.NewAuthServiceClient("com.ta04.srv.auth", s.Client())
	return authServiceClient
}

func NewUserClient() userPB.UserServiceClient {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.api.user"),
	)

	// Initialize the service
	s.Init()

	userServiceClient := userPB.NewUserServiceClient("com.ta04.srv.user", s.Client())
	return userServiceClient
}

func NewOrderClient() orderPB.OrderServiceClient {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.api.order"),
	)

	// Initialize the service
	s.Init()

	orderServiceClient := orderPB.NewOrderServiceClient("com.ta04.srv.order", s.Client())
	return orderServiceClient
}

func NewPaymentClient() paymentPB.PaymentServiceClient {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.api.payment"),
	)

	// Initialize the service
	s.Init()

	paymentServiceClient := paymentPB.NewPaymentServiceClient("com.ta04.srv.payment", s.Client())
	return paymentServiceClient
}