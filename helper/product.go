package helper

import (
	productPB "github.com/SleepingNext/product-service/proto"
	"github.com/micro/go-micro"
	"strconv"
)

func NewClient() productPB.ProductServiceClient {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.api.product"),
	)

	// Initialize the service
	s.Init()

	productServiceClient := productPB.NewProductServiceClient("com.ta04.srv.product", s.Client())
	return productServiceClient
}

func ParseShowAndDeleteQuery(vars map[string][]string) *productPB.Product {
	var id int32
	ids, exists := vars["id"]
	if !exists || len(ids) != 1 {
		id = 0
	} else {
		id64, _ := strconv.ParseInt(ids[0], 10, 64)
		id = int32(id64)
	}

	return &productPB.Product{
		Id: id,
	}
}

func ParseStoreQuery(vars map[string][]string) *productPB.Product {
	var name, description, picture, status string
	var price float64

	names, exists := vars["name"]
	if !exists || len(names) != 1 {
		name = ""
	} else {
		name = names[0]
	}

	descriptions, exists := vars["description"]
	if !exists || len(descriptions) != 1 {
		description = ""
	} else {
		description = descriptions[0]
	}

	prices, exists := vars["price"]
	if !exists || len(prices) != 1 {
		price = 0
	} else {
		price, _ = strconv.ParseFloat(prices[0], 64)
	}

	pictures, exists := vars["picture"]
	if !exists || len(pictures) != 1 {
		picture = ""
	} else {
		picture = pictures[0]
	}

	statuses, exists := vars["status"]
	if !exists || len(statuses) != 1 {
		status = ""
	} else {
		status = statuses[0]
	}

	return &productPB.Product{
		Name: name,
		Description: description,
		Price: price,
		Picture: picture,
		Status: status,
	}
}

func ParseUpdateQuery(vars map[string][]string) *productPB.Product {
	var id int32
	var name, description, picture, status string
	var price float64

	ids, exists := vars["id"]
	if !exists || len(ids) != 1 {
		id = 0
	} else {
		id64, _ := strconv.ParseInt(ids[0], 10, 64)
		id = int32(id64)
	}

	names, exists := vars["name"]
	if !exists || len(names) != 1 {
		name = ""
	} else {
		name = names[0]
	}

	descriptions, exists := vars["description"]
	if !exists || len(descriptions) != 1 {
		description = ""
	} else {
		description = descriptions[0]
	}

	prices, exists := vars["price"]
	if !exists || len(prices) != 1 {
		price = 0
	} else {
		price, _ = strconv.ParseFloat(prices[0], 64)
	}

	pictures, exists := vars["picture"]
	if !exists || len(pictures) != 1 {
		picture = ""
	} else {
		picture = pictures[0]
	}

	statuses, exists := vars["status"]
	if !exists || len(statuses) != 1 {
		status = ""
	} else {
		status = statuses[0]
	}

	return &productPB.Product{
		Id: id,
		Name: name,
		Description: description,
		Price: price,
		Picture: picture,
		Status: status,
	}
}
