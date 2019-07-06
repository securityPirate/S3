package s4

//generating AES-256-CTR

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

// error
func logE(nerror error) {
	if nerror != nil {
		panic(nerror)
	}
}

// Symmetric ...
type Symmetric struct {
	Key, IV, iv8 []byte
}

// Generate AES Key , IV
func (sym *Symmetric) Generate() ([]byte, []byte) {

	sym.Key = make([]byte, 32)
	sym.IV = make([]byte, aes.BlockSize)
	sym.iv8 = make([]byte, aes.BlockSize)
	_, err := rand.Read(sym.Key)
	_, err = rand.Read(sym.IV)
	_, err = rand.Read(sym.iv8)
	logE(err)
	return sym.Key, sym.IV
}

// Encrypt ...
func (sym Symmetric) Encrypt(plain []byte) []byte {
	block, err := aes.NewCipher(sym.Key)
	logE(err)
	stream := cipher.NewCTR(block, sym.IV)
	ciphered := make([]byte, (2*aes.BlockSize)+len(plain))
	copy(ciphered[aes.BlockSize/2:], sym.IV)
	copy(ciphered[0:aes.BlockSize/2], sym.iv8[0:aes.BlockSize/2])
	copy(ciphered[aes.BlockSize/2+aes.BlockSize:], sym.iv8[aes.BlockSize/2:])
	stream.XORKeyStream(ciphered[2*aes.BlockSize:], plain)
	return ciphered
}

// Decrypt ...
func (sym Symmetric) Decrypt(ciphered []byte) []byte {
	block, err := aes.NewCipher(sym.Key)
	logE(err)
	stream := cipher.NewCTR(block, sym.IV)
	plain := make([]byte, len(ciphered)-2*aes.BlockSize)
	stream.XORKeyStream(plain, ciphered[2*aes.BlockSize:])
	return plain
}
