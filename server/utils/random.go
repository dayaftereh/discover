package utils

import (
	"github.com/dayaftereh/discover/server/mathf"
	"encoding/base64"
	"math/rand"
	"time"

	"github.com/pkg/errors"
)

var random = createRand()

func createRand() *rand.Rand {
	now := time.Now()
	source := rand.NewSource(now.UnixNano())
	rand := rand.New(source)
	return rand
}

func RandBytes(size int64) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := random.Read(bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create random %d bytes", size)
	}
	return bytes, nil
}

func RandString(size int64) (string, error) {
	buffer, err := RandBytes(size)
	if err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(buffer)
	value := str[:size]
	return value, nil
}

func RandIntn(n int) int {
	return random.Intn(n)
}

func RandInt64(min int64, max int64) int64 {
	return min + random.Int63n(max-min)
}

func RandFloat64(min float64, max float64) float64 {
	return min + random.Float64()*(max-min)
}

func RandFromRange(r *mathf.Range) float64{
	return RandFloat64(r.Min, r.Max)
}
