package network

type Header byte
type PSize uint16

const (
	HEADER_CG_HANDSHAKE     Header = 0xff
	HEADER_CG_PONG          Header = 0xfe
	HEADER_CG_TIME_SYNC     Header = 0xfc
	HEADER_CG_KEY_AGREEMENT Header = 0xfb

	HEADER_CG_LOGIN            Header = 1
	HEADER_CG_ATTACK           Header = 2
	HEADER_CG_CHAT             Header = 3
	HEADER_CG_CHARACTER_CREATE Header = 4
	HEADER_CG_CHARACTER_DELETE Header = 5
	HEADER_CG_CHARACTER_SELECT Header = 6
	HEADER_CG_MOVE             Header = 7
	HEADER_CG_SYNC_POSITION    Header = 8
	HEADER_CG_ENTERGAME        Header = 10

	HEADER_CG_ITEM_USE    Header = 11
	HEADER_CG_ITEM_DROP   Header = 12
	HEADER_CG_ITEM_MOVE   Header = 13
	HEADER_CG_ITEM_PICKUP Header = 15

	HEADER_CG_QUICKSLOT_ADD  Header = 16
	HEADER_CG_QUICKSLOT_DEL  Header = 17
	HEADER_CG_QUICKSLOT_SWAP Header = 18
	HEADER_CG_WHISPER        Header = 19
	HEADER_CG_ITEM_DROP2     Header = 20

	HEADER_CG_ON_CLICK           Header = 26
	HEADER_CG_EXCHANGE           Header = 27
	HEADER_CG_CHARACTER_POSITION Header = 28
	HEADER_CG_SCRIPT_ANSWER      Header = 29
	HEADER_CG_QUEST_INPUT_STRING Header = 30
	HEADER_CG_QUEST_CONFIRM      Header = 31

	HEADER_CG_SHOP              Header = 50
	HEADER_CG_FLY_TARGETING     Header = 51
	HEADER_CG_USE_SKILL         Header = 52
	HEADER_CG_ADD_FLY_TARGETING Header = 53
	HEADER_CG_SHOOT             Header = 54
	HEADER_CG_MYSHOP            Header = 55

	HEADER_CG_ITEM_USE_TO_ITEM Header = 60
	HEADER_CG_TARGET           Header = 61

	HEADER_CG_TEXT          Header = 64
	HEADER_CG_WARP          Header = 65
	HEADER_CG_SCRIPT_BUTTON Header = 66
	HEADER_CG_MESSENGER     Header = 67

	HEADER_CG_MALL_CHECKOUT    Header = 69
	HEADER_CG_SAFEBOX_CHECKIN  Header = 70
	HEADER_CG_SAFEBOX_CHECKOUT Header = 71

	HEADER_CG_PARTY_INVITE        Header = 72
	HEADER_CG_PARTY_INVITE_ANSWER Header = 73
	HEADER_CG_PARTY_REMOVE        Header = 74
	HEADER_CG_PARTY_SET_STATE     Header = 75
	HEADER_CG_PARTY_USE_SKILL     Header = 76
	HEADER_CG_SAFEBOX_ITEM_MOVE   Header = 77
	HEADER_CG_PARTY_PARAMETER     Header = 78

	HEADER_CG_GUILD             Header = 80
	HEADER_CG_ANSWER_MAKE_GUILD Header = 81

	HEADER_CG_FISHING Header = 82

	HEADER_CG_ITEM_GIVE Header = 83

	HEADER_CG_EMPIRE Header = 90

	HEADER_CG_REFINE Header = 96

	HEADER_CG_MARK_LOGIN   Header = 100
	HEADER_CG_MARK_CRCLIST Header = 101
	HEADER_CG_MARK_UPLOAD  Header = 102
	HEADER_CG_MARK_IDXLIST Header = 104

	HEADER_CG_HACK        Header = 105
	HEADER_CG_CHANGE_NAME Header = 106
	HEADER_CG_LOGIN2      Header = 109
	HEADER_CG_DUNGEON     Header = 110
	HEADER_CG_LOGIN3      Header = 111

	HEADER_CG_GUILD_SYMBOL_UPLOAD Header = 112
	HEADER_CG_SYMBOL_CRC          Header = 113

	HEADER_CG_SCRIPT_SELECT_ITEM Header = 114

	HEADER_CG_LOGIN5_OPENID Header = 116

	HEADER_CG_PASSPOD_ANSWER Header = 202

	HEADER_CG_HS_ACK    Header = 203
	HEADER_CG_XTRAP_ACK Header = 204

	HEADER_CG_DRAGON_SOUL_REFINE Header = 205
	HEADER_CG_STATE_CHECKER      Header = 206

	HEADER_CG_CLIENT_VERSION  Header = 0xfd
	HEADER_CG_CLIENT_VERSION2 Header = 0xf1

	HEADER_GC_KEY_AGREEMENT_COMPLETED Header = 0xfa
	HEADER_GC_KEY_AGREEMENT           Header = 0xfb
	HEADER_GC_TIME_SYNC               Header = 0xfc
	HEADER_GC_PHASE                   Header = 0xfd
	HEADER_GC_BINDUDP                 Header = 0xfe
	HEADER_GC_HANDSHAKE               Header = 0xff

	HEADER_GC_CHARACTER_ADD Header = 1
	HEADER_GC_CHARACTER_DEL Header = 2
	HEADER_GC_MOVE          Header = 3
	HEADER_GC_CHAT          Header = 4
	HEADER_GC_SYNC_POSITION Header = 5

	HEADER_GC_LOGIN_SUCCESS         Header = 6
	HEADER_GC_LOGIN_SUCCESS_NEWSLOT Header = 32
	HEADER_GC_LOGIN_FAILURE         Header = 7

	HEADER_GC_CHARACTER_CREATE_SUCCESS         Header = 8
	HEADER_GC_CHARACTER_CREATE_FAILURE         Header = 9
	HEADER_GC_CHARACTER_DELETE_SUCCESS         Header = 10
	HEADER_GC_CHARACTER_DELETE_WRONG_SOCIAL_ID Header = 11

	HEADER_GC_ATTACK Header = 12
	HEADER_GC_STUN   Header = 13
	HEADER_GC_DEAD   Header = 14

	HEADER_GC_MAIN_CHARACTER_OLD     Header = 15
	HEADER_GC_CHARACTER_POINTS       Header = 16
	HEADER_GC_CHARACTER_POINT_CHANGE Header = 17
	HEADER_GC_CHANGE_SPEED           Header = 18
	HEADER_GC_CHARACTER_UPDATE       Header = 19
	HEADER_GC_CHARACTER_UPDATE_NEW   Header = 24

	HEADER_GC_ITEM_DEL    Header = 20
	HEADER_GC_ITEM_SET    Header = 21
	HEADER_GC_ITEM_USE    Header = 22
	HEADER_GC_ITEM_DROP   Header = 23
	HEADER_GC_ITEM_UPDATE Header = 25

	HEADER_GC_ITEM_GROUND_ADD Header = 26
	HEADER_GC_ITEM_GROUND_DEL Header = 27

	HEADER_GC_QUICKSLOT_ADD  Header = 28
	HEADER_GC_QUICKSLOT_DEL  Header = 29
	HEADER_GC_QUICKSLOT_SWAP Header = 30

	HEADER_GC_ITEM_OWNERSHIP Header = 31

	HEADER_GC_WHISPER Header = 34

	HEADER_GC_MOTION Header = 36
	HEADER_GC_PARTS  Header = 37

	HEADER_GC_SHOP      Header = 38
	HEADER_GC_SHOP_SIGN Header = 39

	HEADER_GC_DUEL_START         Header = 40
	HEADER_GC_PVP                Header = 41
	HEADER_GC_EXCHANGE           Header = 42
	HEADER_GC_CHARACTER_POSITION Header = 43

	HEADER_GC_PING          Header = 44
	HEADER_GC_SCRIPT        Header = 45
	HEADER_GC_QUEST_CONFIRM Header = 46

	HEADER_GC_MOUNT     Header = 61
	HEADER_GC_OWNERSHIP Header = 62
	HEADER_GC_TARGET    Header = 63

	HEADER_GC_WARP Header = 65

	HEADER_GC_ADD_FLY_TARGETING Header = 69
	HEADER_GC_CREATE_FLY        Header = 70
	HEADER_GC_FLY_TARGETING     Header = 71
	HEADER_GC_SKILL_LEVEL_OLD   Header = 72
	HEADER_GC_SKILL_LEVEL       Header = 76

	HEADER_GC_MESSENGER Header = 74
	HEADER_GC_GUILD     Header = 75

	HEADER_GC_PARTY_INVITE       Header = 77
	HEADER_GC_PARTY_ADD          Header = 78
	HEADER_GC_PARTY_UPDATE       Header = 79
	HEADER_GC_PARTY_REMOVE       Header = 80
	HEADER_GC_QUEST_INFO         Header = 81
	HEADER_GC_REQUEST_MAKE_GUILD Header = 82
	HEADER_GC_PARTY_PARAMETER    Header = 83

	HEADER_GC_SAFEBOX_SET            Header = 85
	HEADER_GC_SAFEBOX_DEL            Header = 86
	HEADER_GC_SAFEBOX_WRONG_PASSWORD Header = 87
	HEADER_GC_SAFEBOX_SIZE           Header = 88

	HEADER_GC_FISHING Header = 89

	HEADER_GC_EMPIRE Header = 90

	HEADER_GC_PARTY_LINK   Header = 91
	HEADER_GC_PARTY_UNLINK Header = 92

	HEADER_GC_REFINE_INFORMATION_OLD Header = 95

	HEADER_GC_VIEW_EQUIP Header = 99

	HEADER_GC_MARK_BLOCK   Header = 100
	HEADER_GC_MARK_IDXLIST Header = 102

	HEADER_GC_TIME        Header = 106
	HEADER_GC_CHANGE_NAME Header = 107

	HEADER_GC_DUNGEON Header = 110

	HEADER_GC_WALK_MODE      Header = 111
	HEADER_GC_SKILL_GROUP    Header = 112
	HEADER_GC_MAIN_CHARACTER Header = 113

	HEADER_GC_SEPCIAL_EFFECT Header = 114

	HEADER_GC_NPC_POSITION Header = 115

	HEADER_GC_MATRIX_CARD        Header = 116
	HEADER_GC_LOGIN_KEY          Header = 118
	HEADER_GC_REFINE_INFORMATION Header = 119
	HEADER_GC_CHANNEL            Header = 121

	HEADER_GC_TARGET_UPDATE Header = 123
	HEADER_GC_TARGET_DELETE Header = 124
	HEADER_GC_TARGET_CREATE Header = 125

	HEADER_GC_AFFECT_ADD    Header = 126
	HEADER_GC_AFFECT_REMOVE Header = 127

	HEADER_GC_MALL_OPEN Header = 122
	HEADER_GC_MALL_SET  Header = 128
	HEADER_GC_MALL_DEL  Header = 129

	HEADER_GC_LAND_LIST         Header = 130
	HEADER_GC_LOVER_INFO        Header = 131
	HEADER_GC_LOVE_POINT_UPDATE Header = 132

	HEADER_GC_SYMBOL_DATA Header = 133

	HEADER_GC_DIG_MOTION Header = 134

	HEADER_GC_DAMAGE_INFO          Header = 135
	HEADER_GC_CHAR_ADDITIONAL_INFO Header = 136

	HEADER_GC_MAIN_CHARACTER3_BGM     Header = 137
	HEADER_GC_MAIN_CHARACTER4_BGM_VOL Header = 138

	HEADER_GC_AUTH_SUCCESS Header = 150

	HEADER_GC_PANAMA_PACK Header = 151

	HEADER_GC_HYBRIDCRYPT_KEYS Header = 152
	HEADER_GC_HYBRIDCRYPT_SDB  Header = 153

	HEADER_GC_AUTH_SUCCESS_OPENID Header = 154

	HEADER_GC_ROULETTE Header = 200

	HEADER_GC_REQUEST_PASSPOD        Header = 202
	HEADER_GC_REQUEST_PASSPOD_FAILED Header = 203

	HEADER_GC_HS_REQUEST        Header = 204
	HEADER_GC_XTRAP_CS1_REQUEST Header = 205

	HEADER_GC_SPECIFIC_EFFECT Header = 208

	HEADER_GC_DRAGON_SOUL_REFINE    Header = 209
	HEADER_GC_RESPOND_CHANNELSTATUS Header = 210

	HEADER_GG_LOGIN                    Header = 1
	HEADER_GG_LOGOUT                   Header = 2
	HEADER_GG_RELAY                    Header = 3
	HEADER_GG_NOTICE                   Header = 4
	HEADER_GG_SHUTDOWN                 Header = 5
	HEADER_GG_GUILD                    Header = 6
	HEADER_GG_DISCONNECT               Header = 7
	HEADER_GG_SHOUT                    Header = 8
	HEADER_GG_SETUP                    Header = 9
	HEADER_GG_MESSENGER_ADD            Header = 10
	HEADER_GG_MESSENGER_REMOVE         Header = 11
	HEADER_GG_FIND_POSITION            Header = 12
	HEADER_GG_WARP_CHARACTER           Header = 13
	HEADER_GG_MESSENGER_MOBILE         Header = 14
	HEADER_GG_GUILD_WAR_ZONE_MAP_INDEX Header = 15
	HEADER_GG_TRANSFER                 Header = 16
	HEADER_GG_XMAS_WARP_SANTA          Header = 17
	HEADER_GG_XMAS_WARP_SANTA_REPLY    Header = 18
	HEADER_GG_RELOAD_CRC_LIST          Header = 19
	HEADER_GG_LOGIN_PING               Header = 20
	HEADER_GG_CHECK_CLIENT_VERSION     Header = 21
	HEADER_GG_BLOCK_CHAT               Header = 22

	HEADER_GG_BLOCK_EXCEPTION  Header = 24
	HEADER_GG_SIEGE            Header = 25
	HEADER_GG_MONARCH_NOTICE   Header = 26
	HEADER_GG_MONARCH_TRANSFER Header = 27
	HEADER_GG_PCBANG_UPDATE    Header = 28

	HEADER_GG_CHECK_AWAKENESS Header = 29
)
