package auth

import (
	"go-m2-server/internal/auth"
	"go-m2-server/internal/common"
	"go-m2-server/internal/network"

	"github.com/panjf2000/gnet/v2"
)

func login3Received(c gnet.Conn) (gnet.Action, error) {
	readP, err := common.ReceivePacket(c, &network.TPacketCGLogin3{})
	if err != nil {
		return gnet.Close, err
	}
	p := readP.(*network.TPacketCGLogin3)

	username := network.FromCString(p.Login[:])
	password := network.FromCString(p.Passwd[:])

	resC := make(chan uint32)
	go auth.Login(username, password, resC) // Goroutine
	lkey := <-resC

	if lkey > 0 {
		resP := network.TPacketGCAuthSuccess{
			Header:   network.HEADER_GC_AUTH_SUCCESS,
			LoginKey: lkey,
			Result:   1,
		}
		common.SendPacket(c, resP)
	} else {
		var statVar [network.ACCOUNT_STATUS_MAX_LEN + 1]byte
		copy(statVar[:], network.ToCString(network.LOGIN_WRONG_PASSWORD))
		resP := network.TPacketGCLoginFailure{
			Header: network.HEADER_GC_LOGIN_FAILURE,
			Status: statVar,
		}
		common.SendPacket(c, resP)
	}
	return gnet.None, nil
}
