package s4

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
	s *big.Int
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
	ecc1.s,_ = ecc1.c.ScalarMult(ecc2.pub.X,ecc2.pub.Y,ecc1.private)
	ecc2.s = ecc1.s
	return ecc1.s
}

/*HMAC ...
*
*/
func HMAC(salt []byte , key []byte){
	x := hmac.New(sha256.New,key)
	fmt.Printf("%t",x)
		
}

//KDF ...
func KDF (ecc ECC,){
	print("hello")	
}


// Eccbrb ...
func eccsfdsafds(){	
	// sha-256
	h := sha256.Sum256([]byte("helloworld"))
	fmt.Printf("%x\n\n",h)
}