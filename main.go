package main

import (
	"fmt"

	s4 "./src"
)

func main() {
	a := s4.Symmetric{}
	a.Generate()
	e := a.Encrypt([]byte("jjhelloman,we are in the see man ."))
	

}
