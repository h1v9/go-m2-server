package network

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"hash/crc32"
	"sync"
	"time"
)

func PacketSize(p any) int {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, p)
	if err != nil {
		return 0
	}
	return len(buf.Bytes())
}

func Marshal(p any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, p)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(p any, data []byte) error {
	buf := bytes.NewReader(data)
	return binary.Read(buf, binary.LittleEndian, p)
}

func MakeHandshake() uint32 {
	val := make([]byte, 4)
	_, err := rand.Read(val)
	if err != nil {
		panic(err) // TODO handle error appropriately in production code
	}
	num := binary.LittleEndian.Uint32(val) % (1024 * 1024)

	buf := make([]byte, 8)
	binary.LittleEndian.PutUint32(buf, num)
	binary.LittleEndian.PutUint32(buf[4:], uint32(time.Now().Unix()))

	crcTable := crc32.MakeTable(crc32.Castagnoli)
	crc := crc32.Checksum(buf, crcTable)

	// TODO register this in memory

	return crc
}

var (
	tvBoot    time.Time
	bootSecMu sync.Mutex
)

// getBootSec returns the boot time in seconds.
func getBootSec() int64 {
	bootSecMu.Lock()
	defer bootSecMu.Unlock()

	// If tvBoot is uninitialized, set it to the current time
	if tvBoot.IsZero() {
		tvBoot = time.Now()
	}

	return tvBoot.Unix()
}

func GetDwordTime() uint32 {
	now := time.Now()
	bootSec := getBootSec()

	// Calculate the time difference in milliseconds
	seconds := now.Unix() - bootSec
	milliseconds := int64(now.Nanosecond()) / 1e6

	return uint32(seconds*1000 + milliseconds)
}

func ToCString(str string) []byte {
	return append([]byte(str), 0)
}

// FromCString converts a C-style string (null-terminated byte slice) to a Go string.
func FromCString(cstr []byte) string {
	// Find the null terminator and convert to Go string
	if idx := bytes.IndexByte(cstr, 0); idx != -1 {
		return string(cstr[:idx])
	}
	// If no null terminator is found, treat it as a regular byte slice
	return string(cstr)
}
