module github.com/ta04/api-gateway

go 1.13

require (
	github.com/ta04/user-service v0.0.0-20200406111304-e1323d4a0a54
	github.com/ta04/auth-service v0.0.0-20200406234254-a5b46261da7a
	github.com/ta04/order-service v0.0.0-20200405104014-19c70a80fd75
	github.com/ta04/payment-service v0.0.0-20200405104014-900784da553d
	github.com/ta04/product-service v0.0.0-20200428090908-2c1ada5f5d57
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
)

// +heroku goVersion go1.13.8
