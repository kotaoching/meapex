package utils

import (
	"math/rand"
	"time"

	"github.com/nu7hatch/gouuid"
	"github.com/speps/go-hashids"
)

func GenerateUUID() (string, error) {
	var uid string

	u4, err := uuid.NewV4()
	if err == nil {
		uid = u4.String()
	} else {
		uid = ""
	}

	return uid, err
}

func GenerateHashid(salt string, input []int) string {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 16
	h := hashids.NewWithData(hd)
	e, _ := h.Encode(input)

	return e
}

func GenerateRandom(n int) string {
	alphanum := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(alphanum)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
