package main

import (
	"fmt"
	"github.com/securityPirate/S4/pkg/cryptor"
)

func main() {
	a := cryptor.Symmetric{}
	a.Generate()
	e := a.Encrypt([]byte("AES works man 123456789"))
	fmt.Println(string(e))
	fmt.Println(string(a.Decrypt(e)))
	fmt.Println("----________________----")
	cryptor.EccTest()
}
