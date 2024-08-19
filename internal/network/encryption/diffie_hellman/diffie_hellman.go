package diffie_hellman

import (
	"math/big"

	"go-m2-server/internal/network/encryption/dhkx"
)

type DiffieHellmanInterface interface {
	GetGroup() *dhkx.DHGroup
	SetGroup(group *dhkx.DHGroup)
	GenerateDHKeys() [256]byte
	GetFinalKey(s [256]byte) [256]byte
}

type DiffieHellman struct {
	group     *dhkx.DHGroup
	static    *dhkx.DHKey
	ephemeral *dhkx.DHKey
}

func NewDiffieHellman() DiffieHellmanInterface {
	p, _ := new(big.Int).SetString("B10B8F96A080E01DDE92DE5EAE5D54EC52C99FBCFB06A3C69A6A9DCA52D23B616073E28675A23D189838EF1E2EE652C013ECB4AEA906112324975C3CD49B83BFACCBDD7D90C4BD7098488E9C219A73724EFFD6FAE5644738FAA31A4FF55BCCC0A151AF5F0DC8B4BD45BF37DF365C1A65E68CFDA76D4DA708DF1FB2BC2E4A4371", 16)
	g, _ := new(big.Int).SetString("A4D1CBD5C3FD34126765A442EFB99905F8104DD258AC507FD6406CFF14266D31266FEA1E5C41564B777E690F5504F213160217B4B01B886A5E91547F9E2749F4D7FBD7D3B9A92EE1909D0D2263F80A76A6A24C087A091F531DBF0A0169B6A28AD662A4D18E73AFA32D779D5918D08BC8858F4DCEF97C2A24855E6EEB22B3B2E5", 16)
	q, _ := new(big.Int).SetString("F518AA8781A8DF278ABA4E7D64B7CB9D49462353", 16)
	ngroup := dhkx.CreateGroup(p, g, q)

	return &DiffieHellman{
		group:     ngroup,
		static:    nil,
		ephemeral: nil,
	}
}

func (dh *DiffieHellman) GetGroup() *dhkx.DHGroup {
	return dh.group
}

func (dh *DiffieHellman) SetGroup(group *dhkx.DHGroup) {
	dh.group = group
}

func (dh *DiffieHellman) GenerateDHKeys() [256]byte {
	static, _ := dh.group.GeneratePrivateKey(nil)
	ephemeral, _ := dh.group.GeneratePrivateKey(nil)
	dh.static = static
	dh.ephemeral = ephemeral

	combinedKeys := append(static.Bytes(), ephemeral.Bytes()...)

	var array [256]byte
	copy(array[:], combinedKeys)

	return array
}

func (dh *DiffieHellman) GetFinalKey(s [256]byte) [256]byte {
	clientStaticKey := dhkx.NewPublicKey(s[:128])
	clientEphemeralKey := dhkx.NewPublicKey(s[128:])

	computedStatic, _ := dh.group.ComputeKey(clientStaticKey, dh.static)
	computedEphemeral, _ := dh.group.ComputeKey(clientEphemeralKey, dh.ephemeral)
	combinedKeys := append(computedStatic.Bytes(), computedEphemeral.Bytes()...)

	var array [256]byte
	copy(array[:], combinedKeys)
	return array
}
