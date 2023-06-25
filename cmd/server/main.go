package main

import "fmt"

func Run() error {
		fmt.Println("starting application")
		return nil
}

func main() {
		fmt.Println("Golang REST API")

		if err := Run(); err != nil {
				fmt.Println(err)
		}
}
