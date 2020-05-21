package cryptor

import (
	"crypto/sha256"
	"crypto/hmac"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"fmt"
)

/*
from http://www.secg.org/sec1-v2.pdf

The NIST-800-56-Catenation-KDF should be used, except for backwards compatability 
with implementations already using one of the three other key derivation functions.
https://csrc.nist.gov/publications/detail/sp/800-56a/rev-3/final
https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-56Cr1.pdf


*/


/* Now, when Bob wants to pass a note to Alice, he first picks a random value b,
// and computes the points bG and bA; he then gives the point bA to a key derivation
// function h, which produces a set of symmetric keys; he then uses the symmetric keys 
// to encrypt (and MAC) the message. He then sends the values bG and Encrypth(bA)(note) to Alice.
*/	

/*ECC ...
* struct hold keys values and the curve
*/
type ECC struct {
	c elliptic.Curve
	pub Public
	private []byte
	s1,s2 *big.Int
}

/*Public ...
* represents a generic elliptic curve Point with a
* X and a Y coordinate.
*/
type Public struct {
	X, Y *big.Int
}

/*eccGenerate
	generate ecc privatr , public keys 
*/
func eccGenerate(){
	//Generate curve
	var ecc1 ,ecc2 ECC
	ecc1.c = elliptic.P384()
	//generate the public and private keys
	ecc1.private,ecc1.pub.X,ecc1.pub.Y,_ = elliptic.GenerateKey(ecc1.c,rand.Reader)
	ecc2.private,ecc2.pub.X,ecc2.pub.Y,_ = elliptic.GenerateKey(ecc1.c,rand.Reader)
}

/*eccSharedGenerator
	generate shared secret key 
*/
func (ecc1 *ECC)eccSharedGenerator(ecc2 *ECC) *big.Int {
	ecc1.s1,_ = ecc1.c.ScalarMult(ecc2.pub.X,ecc2.pub.Y,ecc1.private)
	ecc2.s1 = ecc1.s1
	fmt.Printf("%v",ecc1.s1)
	return ecc1.s1
}

/*HMAC ...
*
*/
func HMAC(salt []byte , key []byte){
	x := hmac.New(sha256.New,key)
	fmt.Printf("%x",x)
		
}

//KDF ...
func KDF (ecc ECC,){
	print("hello")	
}


// EccTest ...
func EccTest(){	
	// sha-256
	// h := sha256.Sum256([]byte("helloworld"))
	// fmt.Printf("%x\n\n",h)

	var ecc1 ,ecc2 ECC
	ecc1.c = elliptic.P521()
	ecc2.c = ecc1.c
	//generate the public and private keys
	ecc1.private,ecc1.pub.X,ecc1.pub.Y,_ = elliptic.GenerateKey(ecc1.c,rand.Reader)
	ecc2.private,ecc2.pub.X,ecc2.pub.Y,_ = elliptic.GenerateKey(ecc1.c,rand.Reader)
	ecc1.s1,ecc1.s2 = ecc1.c.ScalarMult(ecc2.pub.X,ecc2.pub.Y,ecc1.private)
	ecc2.s1,ecc2.s2 = ecc1.c.ScalarMult(ecc1.pub.X,ecc1.pub.Y,ecc2.private)
	z := elliptic.Marshal(ecc1.c , ecc1.pub.X,ecc1.pub.Y  )
	fmt.Printf("%x\n\n len : %v\n",z , len(z)) 
	
}