package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	user := os.Getenv("USER")

	if user == "" {
		user = os.Getenv("USERNAME")
	}

	fmt.Println("Имя пользователя:", user)

	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Аргументы не переданы")
	} else {
		fmt.Println("Аргументы CLI:")
		for i, arg := range args {
			fmt.Printf("  arg[%d] = %s\n", i, arg)
		}
	}

	fmt.Println("Текущая версия Go:", runtime.Version())
}
