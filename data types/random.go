package main

import (
	c "crypto/rand"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	random_num := rand.Intn(100)
	fmt.Println(random_num)
	//secure random using crypto rand
	b := make([]byte, 6)
	_, err := c.Read(b)
	if err == nil {
		randomPassword := base64.URLEncoding.EncodeToString(b)
		fmt.Println("Random Password:", randomPassword)
	}

}
