package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	f1(5)
	f1(3)
	f2(2)
	f3()
}

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("f1() Done: ", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1(): ", r)
	}
	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()
	select {
	case <-c2.Done():
		fmt.Println("f2() done :", c2.Err())
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
		return

	}
}

func f3() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "secretKey", "secretValue")
	fmt.Println(search(ctx, "secretKey"))
	fmt.Println(search(ctx, "mySecretKey"))
}

func search(ctx context.Context, key string) (string, error) {
	value := ctx.Value(key)
	if value != nil {
		return fmt.Sprintf("%v", value), nil
	} else {
		return "", errors.New("Key not found")
	}
}
