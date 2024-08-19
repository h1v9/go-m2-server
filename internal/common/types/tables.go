package types

type TSimplePlayer struct {
	DwID          uint32
	SzName        [CHARACTER_NAME_MAX_LEN + 1]byte
	ByJob         byte
	ByLevel       byte
	DwPlayMinutes uint32
	ByST          byte
	ByHT          byte
	ByDX          byte
	ByIQ          byte
	WMainPart     uint16
	BChangeName   byte
	WHairPart     uint16
	BDummy        [4]byte
	X             int32
	Y             int32
	LAddr         [4]byte
	WPort         uint16
	SkillGroup    byte
}

type TPlayerSkill struct {
	MasterType byte
	Level      byte
	NextRead   int32 //time_t
}
