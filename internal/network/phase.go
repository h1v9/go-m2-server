package network

// Define a new type for the phase
type Phase byte

// Declare the constants using iota
const (
	PHASE_CLOSE Phase = iota
	PHASE_HANDSHAKE
	PHASE_LOGIN
	PHASE_SELECT
	PHASE_LOADING
	PHASE_GAME
	PHASE_DEAD

	PHASE_CLIENT_CONNECTING
	PHASE_DBCLIENT
	PHASE_P2P
	PHASE_AUTH
	PHASE_TEEN
	PHASE_PASSPOD
)
