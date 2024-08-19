package auth

import (
	"fmt"
	"go-m2-server/internal/common"
	"go-m2-server/internal/network"
	"log"

	"github.com/panjf2000/gnet/v2"
)

type authServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *authServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("Auth server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (s *authServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	common.OnConnectionopen(c)
	return
}

/* func (s *simpleServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {

} */

func (es *authServer) OnTraffic(c gnet.Conn) gnet.Action {
	pid := common.ReadPacketHeader(c)

	act := gnet.None

	switch network.Header(pid) {
	case network.HEADER_CG_HANDSHAKE:
		act, _ = common.HandshakeReceived(c)
	case network.HEADER_CG_KEY_AGREEMENT:
		act, _ = common.KeyAgreementReceived(c, true)
	case network.HEADER_CG_LOGIN3:
		act, _ = login3Received(c)
	default:
		fmt.Println("Auth: Received UNK header " + fmt.Sprintf("%02X / %d", pid, pid))
		return gnet.Close
	}

	return act
}

func StartServer(port int, multicore bool) {
	echo := &authServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
