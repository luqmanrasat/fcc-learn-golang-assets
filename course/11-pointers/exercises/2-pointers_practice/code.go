package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "dang", "****")
	*message = strings.ReplaceAll(*message, "heck", "****")
	*message = strings.ReplaceAll(*message, "shoot", "*****")
}

// don't touch below this line

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	messages1 := []string{
		"well shoot, this is awful",
		"dang robots",
		"dang them to heck",
	}

	messages2 := []string{
		"well shoot",
		"Allan is going straight to heck",
		"dang... that's a tough break",
	}

	test(messages1)
	test(messages2)
}
