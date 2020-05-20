package main

import (
	s4 "./pkg"
)

func main() {
	a := s4.Symmetric{}
	a.Generate()
	e := a.Encrypt([]byte("jjhelloman,we are in the see man ."))
}
