package tools

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"log"
)

func HashTool(s string) string {
	str := md5.New()
	str.Write([]byte(s))
	res := fmt.Sprintf("%x", str.Sum(nil))
	return res
}

//Generate a random form rand/Rand package.
//Change number n to set random lengths,default 16
func GenerateRandom() (string, error) {
	n := 16
	bslice := make([]byte, n)
	_, err := rand.Read(bslice)
	if err != nil {
		log.Fatal(err)
	}

	if bytes.Equal(bslice, make([]byte, n)) {
		err = fmt.Errorf("generate random error, random only zeroes")
		return "", err
	} else {
		err = nil
		str := fmt.Sprintf("%x", bslice)
		return str, err
	}
}
