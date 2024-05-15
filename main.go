/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"tactical/cmd"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	cmd.Execute()
}
