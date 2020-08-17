package main

import (
	"flag"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var (
		password = flag.String("password", "password", "The password to be hashed.")
		cost     = flag.Int("cost", bcrypt.DefaultCost, "The hashing cost to be used to create the hashed password.")
	)
	flag.Parse()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), *cost)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(hashedPassword))
}
