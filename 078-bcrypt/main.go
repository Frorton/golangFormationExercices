/*
bcrypt itself is a hashing algorithm.

Encryption is a two way operation and you can always get back the original piece of data with the correct key.

Hashing is a one way operation and there is no way to get back the original piece of data
(rainbow tables and weak hashing methods notwithstanding,
for example md5 was commonly used in the early days of the internet).
*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	ps := `password123`
	sb, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error :", err)
	}
	fmt.Println(ps)
	fmt.Println(sb)

	fmt.Println("-------")

	ps2 := `password123`
	err = bcrypt.CompareHashAndPassword(sb, []byte(ps2))
	if err != nil {
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("Logged in")
}
