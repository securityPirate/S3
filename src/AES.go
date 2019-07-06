package s4

//generating AES-256-GCM

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// error
func logE(nerror error){
	if nerror != nil {
		panic(nerror)
	}
}

// Symmetric ...
type Symmetric struct {
	key   []byte
	nonce []byte
}

// Generate AES Key , nonce , block
func (sym *Symmetric) Generate() []byte {
	sym.key = make([]byte, 32)
	_, err := rand.Read(sym.key)
	logE(err)

	block, err := aes.NewCipher(sym.key)
	logE(err)

	AES256GCM, err := cipher.NewGCM(block)
	logE(err)

	sym.nonce = make([]byte, AES256GCM.NonceSize())
	_, err = rand.Read(sym.nonce)
	logE(err)

	return sym.key
}

// Encrypt ...
func (sym Symmetric) Encrypt( plain []byte) []byte {
	block, err := aes.NewCipher(sym.key)
	logE(err)

	AES256GCM, err := cipher.NewGCM(block)
	logE(err)

	theCipher := AES256GCM.Seal(nil, sym.nonce, plain , nil)
	return theCipher
}

// Decrypt ...
func (sym Symmetric) Decrypt(ciphered []byte) []byte {
	block, err := aes.NewCipher(sym.key)
	logE(err)

	AES256GCM, err := cipher.NewGCM(block)
	logE(err)

	plain, err := AES256GCM.Open(nil,sym.nonce,ciphered,nil)
	logE(err)

	return plain
}
