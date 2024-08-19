package encryption

import (
	"errors"
	aes "go-m2-server/internal/network/encryption/aes"
	dh "go-m2-server/internal/network/encryption/diffie_hellman"
)

type EncryptionContext struct {
	DH  dh.DiffieHellmanInterface
	AES aes.AES
}

func (ctx *EncryptionContext) Encrypting() bool {
	return len(ctx.AES.EncodeKey) > 0
}

func (ctx *EncryptionContext) Decrypt(data []byte) ([]byte, error) {
	if !ctx.Encrypting() {
		return nil, errors.New("encryption not ready")
	}

	return ctx.AES.Decrypt(data)
}

func (ctx *EncryptionContext) Encrypt(data []byte) ([]byte, error) {
	return ctx.AES.Encrypt(data)
}
