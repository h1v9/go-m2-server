package common

import (
	"go-m2-server/internal/network"
	"go-m2-server/internal/network/encryption"
	dh "go-m2-server/internal/network/encryption/diffie_hellman"
	"unsafe"

	"github.com/panjf2000/gnet/v2"
)

type ClientContext struct {
	encryption.EncryptionContext
	ID         int
	lastHeader byte
}

func OnConnectionopen(c gnet.Conn) {
	c.SetContext(&ClientContext{
		EncryptionContext: encryption.EncryptionContext{
			DH: dh.NewDiffieHellman(),
		},
	})

	SetPhase(c, network.PHASE_HANDSHAKE)

	packet := network.TPacketGCHandshake{
		BHeader:     network.HEADER_GC_HANDSHAKE,
		DwHandshake: network.MakeHandshake(),
		DwTime:      network.GetDwordTime(),
		LDelta:      0,
	}

	SendPacket(c, packet)
}

func ReadPacketHeader(c gnet.Conn) byte {
	pid, _ := c.Peek(1)
	ctx, _ := c.Context().(*ClientContext)

	if ctx.Encrypting() {
		pid, _ = ctx.Decrypt(pid)
	}

	ctx.lastHeader = pid[0]

	return pid[0]
}

func SendPacket(c gnet.Conn, p any, skipEncryption ...bool) {
	// Default value for skipEncryption is false
	skip := false
	if len(skipEncryption) > 0 {
		skip = skipEncryption[0]
	}

	ctx := c.Context().(*ClientContext)
	packres, _ := network.Marshal(p)

	if ctx.Encrypting() && !skip {
		packres, _ = ctx.Encrypt(packres)
	}

	c.Write(packres)
}

func ReceivePacket(c gnet.Conn, p any) (any, error) {
	ctx := c.Context().(*ClientContext)
	size := network.PacketSize(p)

	data, err := c.Next(size)
	if err != nil {
		return p, err
	}

	if ctx.Encrypting() {
		data, err = ctx.Decrypt(data[1:])
	} else {
		data = data[1:]
	}
	if err != nil {
		return p, err
	}

	// Prepend header
	realData := make([]byte, len(data)+1)
	realData[0] = ctx.lastHeader
	copy(realData[1:], data)

	network.Unmarshal(p, realData)
	return p, nil
}

func ReceiveDynamicPacket(c gnet.Conn, p any) (any, error) {
	ctx := c.Context().(*ClientContext)

	// Load size
	var psize network.PSize
	toread := unsafe.Sizeof(psize)
	psizedata, err := c.Peek(int(toread) + 1)
	if err != nil {
		return p, err
	}
	if ctx.Encrypting() {
		psizedata, err = ctx.Decrypt(psizedata[1:])
	}
	if err != nil {
		return p, err
	}

	network.Unmarshal(&psize, psizedata[1:])

	data, err := c.Next(int(psize))
	if err != nil {
		return p, err
	}

	if ctx.Encrypting() {
		data, err = ctx.Decrypt(data[1+int(toread):])
	} else {
		data = data[1+int(toread):]
	}
	if err != nil {
		return p, err
	}

	// Prepend header
	realData := make([]byte, len(data)+1)
	realData[0] = ctx.lastHeader
	copy(realData[1:], psizedata)
	copy(realData[1+int(toread):], data)

	network.Unmarshal(p, data)
	return p, nil
}

func SetPhase(c gnet.Conn, phase network.Phase) {
	p := network.TPacketGCPhase{
		Header: network.HEADER_GC_PHASE,
		Phase:  phase,
	}

	SendPacket(c, p)
}
