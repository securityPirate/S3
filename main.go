package main

import (
	"fmt"
	"github.com/securityPirate/S4/pkg/cryptor"
)

func main() {
	a := cryptor.Symmetric{}
	a.Generate()
	e := a.Encrypt([]byte("jjhelloman,we are in the see man ."))
	fmt.Println(e)

}
