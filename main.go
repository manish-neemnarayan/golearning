package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type ContextKey string

// const MyKey ContextKey = "username" <-- ContextKey("username") both are same

func main() {
	ctx := context.WithValue(context.Background(), ContextKey("username"), "neemnarayan")

	start := time.Now()
	result, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response: %+v\ntook: %v\n", result, time.Since(start))
}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()

	val := ctx.Value(ContextKey("username"))
	fmt.Println("the value  = ", val)

	type result struct {
		userId string
		err    error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// Done()
	// -> context tiemout is exceeded or canceled by cancel()
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err
	}
}

func thirdHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 100)
	return "userID: 12", nil
}
