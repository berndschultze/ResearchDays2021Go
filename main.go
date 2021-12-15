package main

import (
	"fmt"
	"ttslight/subscription/subscription"
)

func main() {
	fmt.Println("Hello, World!")
	subscription1 := subscription.New("group_a", 5000, 1)
	fmt.Println(subscription1.ToString())
}
