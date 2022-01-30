package tokenengine

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joshinjohnson/authentication-engine/pkg/models"
	"log"
	"strings"
)

type TokenGeneratorEngine struct {
	Log log.Logger
}

func (tm TokenGeneratorEngine) TokenGeneratorFunc() func(cred models.UserCredential, payload map[string]string, privateKey string) string {
	return func(cred models.UserCredential, payload map[string]string, privateKey string) string {
		header := cred.Email
		h := hmac.New(sha256.New, []byte(privateKey))
		payloadStr, err := json.Marshal(payload)
		if err != nil {
			tm.Log.Panic("Error generating token")
			return string(payloadStr)
		}

		payload64 := base64.StdEncoding.EncodeToString(payloadStr)
		header64 := base64.StdEncoding.EncodeToString([]byte(header))
		message := header64 + "." + payload64

		unsignedStr := header + string(payloadStr)
		h.Write([]byte(unsignedStr))
		signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

		tokenStr := message + "." + signature
		return tokenStr
	}
}

func (tm TokenGeneratorEngine) VerifyToken(tokenStr string, privateKey string) (bool, error) {
	splitToken := strings.Split(tokenStr, ".")
	if len(splitToken) != 3 {
		return false, errors.New("invalid token found")
	}

	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, err
	}
	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, err
	}

	unsignedStr := string(header) + string(payload)
	h := hmac.New(sha256.New, []byte(privateKey))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	fmt.Println(signature)

	if signature != splitToken[2] {
		return false, nil
	}
	return true, nil
}
