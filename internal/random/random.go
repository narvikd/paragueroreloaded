package random

import (
	"crypto/rand"
	"github.com/pkg/errors"
	"math/big"
	"paraguero_reloaded/internal/logger"
)

func Num(size int) int64 {
	bigNum, err := rand.Int(rand.Reader, big.NewInt(int64(size)))
	if err != nil {
		log := logger.GetLog()
		log.Errorln(errors.Wrap(err, "error converting int to bigNum on random.Num"))
	}
	return bigNum.Int64()
}
