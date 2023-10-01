package main

import (
	"fmt"
	"time"
)

type ShoppingClient struct {
	endpoint string
	apiKey   string
	userId   string

	shouldRetry bool
	timeout     time.Duration
}

type ShoppingClientOption func(*ShoppingClient)

func WithApiKey(key string) ShoppingClientOption {
	return func(c *ShoppingClient) {
		c.apiKey = key
	}
}

func WithUserId(id string) ShoppingClientOption {
	return func(c *ShoppingClient) {
		c.userId = id
	}
}

func WithTimeout(timeout time.Duration) ShoppingClientOption {
	return func(c *ShoppingClient) {
		c.timeout = timeout
	}
}

func WithRetry(shouldRetry bool) ShoppingClientOption {
	return func(c *ShoppingClient) {
		c.shouldRetry = shouldRetry
	}
}

func NewShoppingClient(endpoint string, options ...ShoppingClientOption) *ShoppingClient {
	client := &ShoppingClient{
		endpoint: endpoint,
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func main() {
	client := NewShoppingClient("https://api.shopping.com/v1",
		WithApiKey("my-api-key"),
		WithTimeout(10*time.Second),
	)
	fmt.Printf("%+v\n", client) // &{endpoint:https://api.shopping.com/v1 apiKey:my-api-key userId: shouldRetry:false timeout:10000000000}

	clientV2 := NewShoppingClient("https://api.shopping.com/v2",
		WithRetry(true),
		WithUserId("my-user-id"),
	)
	fmt.Printf("%+v\n", clientV2) // &{endpoint:https://api.shopping.com/v2 apiKey: userId:my-user-id shouldRetry:true timeout:0}
}
