package main

import (
	"git.sansera.com/mtkach/golang-test-assignment/fastlog"
	"git.sansera.com/mtkach/golang-test-assignment/generator"
)

func main() {
	err := generator.GenerateExampleCSV()
	if err != nil {
		fastlog.Fatal("Error generating example csv", "err", err)
	}
}
