package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func ENVInit() {
	if err := godotenv.Load(); err != nil {
		panic("env Failed to Connected")
	}
	fmt.Println("env Connection Succesfully")
}
