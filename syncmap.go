package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create the threadsafe map.
	var sm sync.Map

	// Fetch an item that doesn't exist yet.
	result, ok := sm.Load("hello")
	if ok {
		fmt.Println(result.(string))
	} else {
		fmt.Println("value not found for key: `hello`")
	}

	// Store an item in the map.
	sm.Store("hello", "world")
	fmt.Println("added value: `world` for key: `hello`")

	// Fetch the item we just stored.
	result, ok = sm.Load("hello")
	if ok {
		fmt.Printf("result: `%s` found for key: `hello`\n", result.(string))
	}

	fmt.Println("---------------------------")
}
