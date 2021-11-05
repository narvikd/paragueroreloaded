package appinit

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	mathRand "math/rand"
	"paraguero_reloaded/internal/logger"
)

func SeedRNG() {
	log := logger.GetLog()
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		log.Panicln("cannot seed math/rand")
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
