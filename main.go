package main

import (
	//"crypto"
	"crypto/cipher"
	"crypto/aes"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"fmt"
)

// Point represents a generic elliptic curve Point with a
// X and a Y coordinate.
type Point struct {
	X, Y *big.Int
}

func main(){
	//generating AES-256-GCM
	//key should be random
	plaintext := []byte("exampleplaintext")
	aesKey := make([]byte, 32)
	_ ,err := rand.Read(aesKey)
	block , err := aes.NewCipher(aesKey)
	AES256GCM, err := cipher.NewGCM(block)
	nonce := make([]byte, AES256GCM.NonceSize())
	_, err = rand.Read(nonce)
	ciphertext := AES256GCM.Seal(nil,nonce,plaintext,nil)
	plain1 , err := AES256GCM.Open(nil,nonce,ciphertext,nil)
	fmt.Printf("%s\n\n\n",ciphertext)
	fmt.Printf("%s\n\n\n",plain1)

	//Generate ECC
	c := elliptic.P384()
	private,x,y,err := elliptic.GenerateKey(c,rand.Reader)
	public := Point{X:x , Y:y}
	z := elliptic.Marshal(c,public.X,public.Y)
	w,q := elliptic.Unmarshal(c,z)
	//return shared secret
	s,_ := c.ScalarMult(public.X,public.Y,private)

	if err != nil {
        panic(err)
	}

	fmt.Print(private,"\n\n\n")
	fmt.Print(public,"\n\n\n")
	fmt.Print(s.Bytes(),"\n\n\n")
	fmt.Printf("%v\n\n\n",z)
	fmt.Printf("%v\n\n\n",w)
	fmt.Printf("%v\n\n\n",q)
	
}