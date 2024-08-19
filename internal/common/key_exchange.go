package common

import (
	"go-m2-server/internal/network"
	aes "go-m2-server/internal/network/encryption/aes"

	"github.com/panjf2000/gnet/v2"
)

func HandshakeReceived(c gnet.Conn) (gnet.Action, error) {
	// Read but ignore
	_, err := ReceivePacket(c, &network.TPacketGCHandshake{})
	if err != nil {
		return gnet.Close, err
	}

	context := c.Context().(*ClientContext)

	keys := context.DH.GenerateDHKeys()

	// Send ka
	kep := network.TPacketKeyAgreement{
		BHeader:       network.HEADER_GC_KEY_AGREEMENT,
		WAgreedLength: 256,
		WDataLength:   256,
		Data:          keys,
	}

	SendPacket(c, kep)

	return gnet.None, nil
}

func KeyAgreementReceived(c gnet.Conn, auth bool) (gnet.Action, error) {
	readP, err := ReceivePacket(c, &network.TPacketKeyAgreement{})
	if err != nil {
		return gnet.Close, err
	}
	p := readP.(*network.TPacketKeyAgreement)

	context := c.Context().(*ClientContext)

	finalKey := context.DH.GetFinalKey(p.Data)

	aesObj := aes.Init(finalKey)

	// send confirm
	kep := network.TPacketKeyAgreementCompleted{
		BHeader: network.HEADER_GC_KEY_AGREEMENT_COMPLETED,
		Data:    [3]byte{12, 22, 54},
	}
	SendPacket(c, kep)

	// Prepare for encryption
	c.Flush()

	if auth {
		SetPhase(c, network.PHASE_AUTH)
	} else {
		SetPhase(c, network.PHASE_LOGIN)
	}

	// Update context and enable encryption
	context.AES = *aesObj
	//c.SetContext(&context)

	return gnet.None, nil
}
