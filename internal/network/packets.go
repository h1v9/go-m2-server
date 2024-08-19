package network

import (
	"go-m2-server/internal/common/types"
)

const (
	LOGIN_MAX_LEN          = 30
	PASSWD_MAX_LEN         = 16
	ACCOUNT_STATUS_MAX_LEN = 8

	LOGIN_WRONG_PASSWORD    = "WRONGPWD"
	LOGIN_BANNED            = "BLOCK"
	LOGIN_NOT_AVAILABLE     = "NOTAVAIL"
	LOGIN_ALREADY_LOGGED_IN = "ALREADY"
	LOGIN_BADSCLID          = "BADSCLID"
	LOGIN_AGE_LIMIT         = "AGELIMIT"
	LOGIN_SUCCESS           = "OK"
)

type TPacketGCStateCheck struct {
	Header Header
	Size   int32
	Port   int16
	Status byte
}

// Define a Go type that wraps the C struct
type TPacketGCHandshake struct {
	BHeader     Header
	DwHandshake uint32
	DwTime      uint32
	LDelta      int32
}

type TPacketKeyAgreement struct {
	BHeader       Header
	WAgreedLength uint16
	WDataLength   uint16
	Data          [256]byte
}

type TPacketKeyAgreementCompleted struct {
	BHeader Header
	Data    [3]byte // dummy (not used)
}

type TPacketGCPhase struct {
	Header Header
	Phase  Phase
}

type TPacketCGLogin3 struct {
	Header    Header
	Login     [LOGIN_MAX_LEN + 1]byte
	Passwd    [PASSWD_MAX_LEN + 1]byte
	ClientKey [4]uint32
}

type TPacketGCAuthSuccess struct {
	Header   Header
	LoginKey uint32
	Result   byte
}

type TPacketGCLoginFailure struct {
	Header Header
	Status [ACCOUNT_STATUS_MAX_LEN + 1]byte
}

type TPacketCGLogin2 struct {
	Header       Header
	Login        [LOGIN_MAX_LEN + 1]byte
	DwLoginKey   uint32
	AdwClientKey [4]uint32
}

type TPacketGCEmpire struct {
	Header Header
	Empire byte
}

type TPacketGCLoginSuccess struct {
	BHeader Header
	Players [types.PLAYER_PER_ACCOUNT]types.TSimplePlayer
	GuildID [types.PLAYER_PER_ACCOUNT]uint32

	GuildName [types.PLAYER_PER_ACCOUNT][types.GUILD_NAME_MAX_LEN + 1]byte

	Handle    uint32
	RandomKey uint32
}

type TPacketCGPlayerSelect struct {
	Header Header
	Index  byte
}

type TPacketGCMainCharacter struct {
	Header     Header
	DwVID      uint32
	WRaceNum   uint16
	SzName     [types.CHARACTER_NAME_MAX_LEN + 1]byte
	Lx, Ly, Lz int32
	Empire     byte
	SkillGroup byte
}

type TPacketGCPoints struct {
	Header Header
	Points [types.POINT_MAX_NUM]int32
}

type TPacketGCSkillLevel struct {
	Header Header
	Skills [types.SKILL_MAX_NUM]types.TPlayerSkill
}

type TPacketCGClientVersion2 struct {
	Header    Header
	Filename  [32 + 1]byte
	Timestamp [32 + 1]byte
}

type TPacketGCCharacterAdd struct {
	Header Header
	DwVID  uint32

	Angle float32
	X     int32
	Y     int32
	Z     int32

	BType        types.CharType
	WRaceNum     uint16
	BMovingSpeed uint8
	BAttackSpeed uint8

	BStateFlag   uint8
	DwAffectFlag [3]uint32
}

type TPacketGCCharacterAdditionalInfo struct {
	Header      Header
	DwVID       uint32
	Name        [types.CHARACTER_NAME_MAX_LEN + 1]byte
	AwPart      [types.CHR_EQUIPPART_NUM]uint16
	BEmpire     uint8
	DwGuildID   uint32
	DwLevel     uint32
	SAlignment  int16
	BPKMode     uint8
	DwMountVnum uint32
}

type TPacketCGMove struct {
	Header Header
	Func   byte
	Arg    byte
	Rot    byte
	X      int32
	Y      int32
	Time   uint32
}

type TPacketCGChat struct {
	Header Header
	Size   PSize
	Type   byte
	Chat   string
}
