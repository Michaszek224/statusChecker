package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Wrong command should be:")
		fmt.Println("statuschecker 'url'")
		log.Fatal()
	}
	resp, err := http.Get(args[1])
	if err != nil {
		log.Fatalf("error while handling %v", args[1])
	}

	fmt.Printf("Status of %v:\n", args[1])
	fmt.Println(resp.Status)
}
