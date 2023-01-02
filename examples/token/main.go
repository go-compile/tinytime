package main

import (
	"bytes"
	"crypto/hmac"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/go-compile/tinytime"
	"golang.org/x/crypto/sha3"
)

func main() {
	// example stateless token
	expire := tinytime.Unix(
		time.Now().Add(time.Hour * 12).Unix(),
	).Hours()
	username := "go-compile"

	expireHex := hex.EncodeToString(tinytime.ToBytes(expire))

	// built the token
	token := username + "," + expireHex

	// setup hmac
	secret := "secure-hmac-key"
	h := hmac.New(sha3.New256, []byte(secret))
	h.Write(append([]byte(username), expireHex...))

	// add signature to token
	token += "," + base64.StdEncoding.EncodeToString(h.Sum(nil))

	fmt.Printf("Token: %s\n", token)

	verifyToken(token, secret)
}

func verifyToken(token, secret string) {
	split := strings.SplitN(token, ",", 3)
	username := split[0]
	expireBuf, _ := hex.DecodeString(split[1])
	expire := tinytime.FromBytes(expireBuf)

	signature, _ := base64.StdEncoding.DecodeString(split[2])

	fmt.Printf("Decoded Username: %s\n", username)
	fmt.Printf("Decoded Expire: %d\n", expire)

	h := hmac.New(sha3.New256, []byte(secret))
	h.Write(append([]byte(username), []byte(split[1])...))

	fmt.Printf("Signature Valid: %t\n", bytes.Equal(h.Sum(nil), signature))

	fmt.Printf("Expired: %t\n", tinytime.Now().Hours() >= expire)
}
