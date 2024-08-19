package AES

import (
	"crypto/aes"
	"crypto/cipher"
)

const (
	BLOCK_SIZE = 16
	KEY_LENGTH = 16
)

type AES struct {
	EncodeKey    []byte
	EncodeIv     []byte
	DecodeKey    []byte
	DecodeIv     []byte
	encodeStream cipher.Stream
	decodeStream cipher.Stream
	deblock      cipher.Block
	enblock      cipher.Block
}

// Encrypt encrypts the plaintext using AES-128 in CTR mode.
func (aesct *AES) Encrypt(plaintext []byte) (ciphertext []byte, err error) {

	stream := aesct.encodeStream

	// Allocate space for the ciphertext.
	ciphertext = make([]byte, len(plaintext))

	// Encrypt the plaintext.
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

// Decrypt decrypts the ciphertext using AES-128 in CTR mode.
func (aesct *AES) Decrypt(ciphertext []byte) (plaintext []byte, err error) {

	// Create the CTR stream.
	stream := aesct.decodeStream

	// Allocate space for the plaintext.
	plaintext = make([]byte, len(ciphertext))

	// Decrypt the ciphertext.
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}

func Init(finalKey [256]byte) *AES {
	var offset int
	aesS := AES{}
	aesS.EncodeKey = finalKey[:KEY_LENGTH]

	offset = min(KEY_LENGTH, len(finalKey)-KEY_LENGTH)
	aesS.DecodeKey = finalKey[offset : offset+KEY_LENGTH]

	offset = len(finalKey) - BLOCK_SIZE
	aesS.EncodeIv = finalKey[offset : offset+BLOCK_SIZE]

	if offset < BLOCK_SIZE {
		offset = 0
	} else {
		offset -= BLOCK_SIZE
	}
	aesS.DecodeIv = finalKey[offset : offset+BLOCK_SIZE]

	// Init ciphers
	aesS.deblock, _ = aes.NewCipher(aesS.DecodeKey)
	aesS.enblock, _ = aes.NewCipher(aesS.EncodeKey)
	aesS.decodeStream = cipher.NewCTR(aesS.deblock, aesS.DecodeIv)
	aesS.encodeStream = cipher.NewCTR(aesS.enblock, aesS.EncodeIv)

	return &aesS
}
