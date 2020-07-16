module github.com/ta04/api-gateway

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/hashicorp/go-msgpack v0.5.4 // indirect
	github.com/hashicorp/memberlist v0.1.4 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/ta04/auth-service v0.0.0-20200716033003-32dd5b738256
	github.com/ta04/order-service v0.0.0-20200709160327-f48007f58ee1
	github.com/ta04/payment-method-service v0.0.0-20200709173812-1603ab66c9ec
	github.com/ta04/payment-service v0.0.0-20200709173946-7ec309552ea8
	github.com/ta04/product-service v0.0.0-20200709174109-d689f10a1740
	github.com/ta04/user-service v0.0.0-20200709174013-ce781e76d996
)
