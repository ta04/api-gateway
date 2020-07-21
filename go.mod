module github.com/ta04/api-gateway

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/hashicorp/go-msgpack v0.5.4 // indirect
	github.com/hashicorp/memberlist v0.1.4 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/ta04/auth-service v0.0.0-20200721015929-f1515dfbc3a3
	github.com/ta04/order-service v0.0.0-20200719185756-bbc7da825e94
	github.com/ta04/payment-method-service v0.0.0-20200719184732-81c713eec806
	github.com/ta04/payment-service v0.0.0-20200719182157-99f2bf3b6219
	github.com/ta04/product-service v0.0.0-20200719202405-ecbda318c0ab
	github.com/ta04/user-service v0.0.0-20200721034045-cef5fa22326d
)
