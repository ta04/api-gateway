module github.com/SleepingNext/api-gateway

go 1.13

require (
	github.com/G0tYou/user-service v0.0.0-20200406111304-e1323d4a0a54
	github.com/SleepingNext/auth-service v0.0.0-20200406165518-4b3ef33c067e
	github.com/SleepingNext/auth-service-cli v0.0.0-20200325115926-9aae17ac9ef1 // indirect
	github.com/SleepingNext/order-service v0.0.0-20200405104014-19c70a80fd75
	github.com/SleepingNext/payment-service v0.0.0-20200313155410-9f030ef23606
	github.com/SleepingNext/product-service v0.0.0-20200313155410-36993b85ea42
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
)

// +heroku goVersion go1.13.8
