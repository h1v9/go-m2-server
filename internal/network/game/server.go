package game

import (
	"fmt"
	"go-m2-server/internal/common"
	"go-m2-server/internal/network"
	"log"

	"github.com/panjf2000/gnet/v2"
)

type gameServer struct {
	gnet.BuiltinEventEngine

	eng       gnet.Engine
	addr      string
	multicore bool
}

func (es *gameServer) OnBoot(eng gnet.Engine) gnet.Action {
	es.eng = eng
	log.Printf("Game server with multi-core=%t is listening on %s\n", es.multicore, es.addr)
	return gnet.None
}

func (s *gameServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	common.OnConnectionopen(c)
	return
}

/* func (s *simpleServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {

} */

func (es *gameServer) OnTraffic(c gnet.Conn) gnet.Action {
	pid := common.ReadPacketHeader(c)

	act := gnet.None

	switch network.Header(pid) {
	case network.HEADER_CG_HANDSHAKE:
		act, _ = common.HandshakeReceived(c)
	case network.HEADER_CG_KEY_AGREEMENT:
		act, _ = common.KeyAgreementReceived(c, false)
	case network.HEADER_CG_STATE_CHECKER:
		c.Discard(1)
		return SendState(c)
	case network.HEADER_CG_LOGIN2:
		act, _ = login2Received(c)
	case network.HEADER_CG_CHARACTER_SELECT:
		act, _ = characterSelectReceived(c)
	case network.HEADER_CG_MARK_LOGIN:
		fmt.Println("Game: Not a mark server")
		common.SetPhase(c, network.PHASE_CLOSE)
		act = gnet.Close
	case network.HEADER_CG_CLIENT_VERSION2:
		act, _ = clientVersion2Received(c)
	case network.HEADER_CG_ENTERGAME:
		act, _ = enterGame(c)
	case network.HEADER_CG_MOVE:
		act, _ = handleMove(c)
	case network.HEADER_CG_CHAT:
		act, _ = handleChat(c)
	default:
		fmt.Println("Game: Received UNK header " + fmt.Sprintf("%02X / %d", pid, pid))
		act = gnet.Close
	}

	return act
}

func StartServer(port int, multicore bool) {
	echo := &gameServer{addr: fmt.Sprintf("tcp://:%d", port), multicore: multicore}
	log.Fatal(gnet.Run(echo, echo.addr, gnet.WithMulticore(multicore)))
}
