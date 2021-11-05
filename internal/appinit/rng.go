package appinit

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	mathRand "math/rand"
	"paraguero_reloaded/internal/logger"
)

// SeedRNG seeds the math/rand RNG according to: https://stackoverflow.com/a/54491783.
// It should be initialized as soon as possible in the application.
func SeedRNG() {
	log := logger.GetLog()
	var b [8]byte
	_, err := cryptoRand.Read(b[:])
	if err != nil {
		log.Panicln("cannot seed math/rand")
	}
	mathRand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
