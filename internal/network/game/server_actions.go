package game

import (
	"go-m2-server/internal/auth"
	"go-m2-server/internal/common"
	"go-m2-server/internal/common/types"
	"go-m2-server/internal/network"

	"github.com/panjf2000/gnet/v2"
)

func SendState(c gnet.Conn) gnet.Action {
	packet := network.TPacketGCStateCheck{
		Header: network.HEADER_GC_RESPOND_CHANNELSTATUS,
		Size:   1,
		Port:   23000,
		Status: 1,
	}

	common.SendPacket(c, packet)
	return gnet.Close
}

func login2Received(c gnet.Conn) (gnet.Action, error) {
	readP, err := common.ReceivePacket(c, &network.TPacketCGLogin2{})
	if err != nil {
		return gnet.Close, err
	}
	p := readP.(*network.TPacketCGLogin2)

	// Check login
	authC := make(chan bool)
	go auth.LoginByKey(p.DwLoginKey, authC)
	lres := <-authC

	if !lres {
		var statVar [network.ACCOUNT_STATUS_MAX_LEN + 1]byte
		copy(statVar[:], network.ToCString(network.LOGIN_WRONG_PASSWORD))
		resP := network.TPacketGCLoginFailure{
			Header: network.HEADER_GC_LOGIN_FAILURE,
			Status: statVar,
		}
		common.SendPacket(c, resP)
		return gnet.Close, nil
	}

	// Send empire
	empP := network.TPacketGCEmpire{
		Header: network.HEADER_GC_EMPIRE,
		Empire: 1,
	}
	common.SendPacket(c, empP)

	common.SetPhase(c, network.PHASE_SELECT)

	// Send characters
	var characters [types.PLAYER_PER_ACCOUNT]types.TSimplePlayer

	var charName [types.CHARACTER_NAME_MAX_LEN + 1]byte
	copy(charName[:], network.ToCString("Dummy"))

	characters[0].DwID = 1
	characters[0].SzName = charName
	characters[0].ByJob = 2
	characters[0].ByLevel = 120
	characters[0].DwPlayMinutes = 10
	characters[0].ByST = 90
	characters[0].ByHT = 90
	characters[0].ByDX = 90
	characters[0].ByIQ = 90
	characters[0].WMainPart = 12039
	characters[0].BChangeName = 0
	characters[0].WHairPart = 0
	characters[0].BDummy = [4]byte{0, 0, 0, 0}
	characters[0].X = 960300
	characters[0].Y = 272200
	characters[0].LAddr = [4]byte{127, 0, 0, 1}
	characters[0].WPort = 23000
	characters[0].SkillGroup = 1

	var guilds [types.PLAYER_PER_ACCOUNT]uint32
	var guildNames [types.PLAYER_PER_ACCOUNT][types.GUILD_NAME_MAX_LEN + 1]byte

	charPack := network.TPacketGCLoginSuccess{
		BHeader:   network.HEADER_GC_LOGIN_SUCCESS_NEWSLOT,
		Players:   characters,
		GuildID:   guilds,
		GuildName: guildNames,
		Handle:    10,
		RandomKey: 2834,
	}

	common.SendPacket(c, charPack)

	return gnet.None, nil
}

func characterSelectReceived(c gnet.Conn) (gnet.Action, error) {
	_, err := common.ReceivePacket(c, &network.TPacketCGPlayerSelect{})
	if err != nil {
		return gnet.Close, err
	}
	//p := readP.(*network.TPacketCGPlayerSelect)

	common.SetPhase(c, network.PHASE_LOADING)

	var charName [types.CHARACTER_NAME_MAX_LEN + 1]byte
	copy(charName[:], network.ToCString("Dummy"))

	mcp := network.TPacketGCMainCharacter{
		Header:     network.HEADER_GC_MAIN_CHARACTER,
		DwVID:      15,
		WRaceNum:   2,
		SzName:     charName,
		Lx:         959538,
		Ly:         273722,
		Lz:         0,
		Empire:     1,
		SkillGroup: 1,
	}

	common.SendPacket(c, mcp)

	var points [types.POINT_MAX_NUM]int32
	points[types.POINT_LEVEL] = 120
	points[types.POINT_EXP] = 0
	points[types.POINT_NEXT_EXP] = 0
	points[types.POINT_HP] = 100
	points[types.POINT_MAX_HP] = 200
	points[types.POINT_SP] = 100
	points[types.POINT_MAX_SP] = 200
	points[types.POINT_GOLD] = 5000
	points[types.POINT_STAMINA] = 100
	points[types.POINT_MOV_SPEED] = 170
	points[types.POINT_ATT_SPEED] = 180
	points[types.POINT_CASTING_SPEED] = 100

	// TODO other stats?

	pp := network.TPacketGCPoints{
		Header: network.HEADER_GC_CHARACTER_POINTS,
	}

	common.SendPacket(c, pp)
	var skills [types.SKILL_MAX_NUM]types.TPlayerSkill
	slp := network.TPacketGCSkillLevel{
		Header: network.HEADER_GC_SKILL_LEVEL,
		Skills: skills,
	}

	common.SendPacket(c, slp)

	return gnet.None, nil
}

func clientVersion2Received(c gnet.Conn) (gnet.Action, error) {
	_, err := common.ReceivePacket(c, &network.TPacketCGClientVersion2{})
	if err != nil {
		return gnet.Close, err
	}

	return gnet.None, nil
}

func enterGame(c gnet.Conn) (gnet.Action, error) {
	c.Discard(2)

	common.SetPhase(c, network.PHASE_GAME)

	pca := network.TPacketGCCharacterAdd{
		Header:       network.HEADER_GC_CHARACTER_ADD,
		DwVID:        15,
		Angle:        1,
		X:            959538,
		Y:            273722,
		Z:            0,
		BType:        types.CHAR_TYPE_PC,
		WRaceNum:     2,
		BMovingSpeed: 170,
		BAttackSpeed: 170,
		BStateFlag:   0,
		DwAffectFlag: [3]uint32{0, 0, 0},
	}
	common.SendPacket(c, pca)

	var charName [types.CHARACTER_NAME_MAX_LEN + 1]byte
	copy(charName[:], network.ToCString("Dummy"))
	info := network.TPacketGCCharacterAdditionalInfo{
		Header:      network.HEADER_GC_CHAR_ADDITIONAL_INFO,
		DwVID:       15,
		Name:        charName,
		AwPart:      [types.CHR_EQUIPPART_NUM]uint16{},
		BEmpire:     2,
		DwGuildID:   0,
		DwLevel:     120,
		SAlignment:  100,
		BPKMode:     3,
		DwMountVnum: 0,
	}

	common.SendPacket(c, info)

	return gnet.None, nil
}

func handleMove(c gnet.Conn) (gnet.Action, error) {
	_, err := common.ReceivePacket(c, &network.TPacketCGMove{})
	if err != nil {
		return gnet.None, err
	}
	//p := readP.(*network.TPacketCGMove)

	return gnet.None, nil
}

func handleChat(c gnet.Conn) (gnet.Action, error) {
	_, err := common.ReceiveDynamicPacket(c, &network.TPacketCGChat{})
	if err != nil {
		return gnet.None, err
	}
	//p := readP.(*network.TPacketCGChat)

	return gnet.None, nil
}
