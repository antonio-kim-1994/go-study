package signatures

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

/*
Hmac256으로 인증키를 생성합니다.
Uri를 전달할 때 반드시 BaseURL 외의 모든 path를 함께 전달해야 합니다.
*/
func GenerateHmac(method string, uri string, timestamp string, accessKey string, secretKey string) (string, error) {
	data := method + " " + uri + "\n" + timestamp + "\n" + accessKey

	hmacSha256 := hmac.New(sha256.New, []byte(secretKey))
	hmacSha256.Write([]byte(data))
	mac := hmacSha256.Sum(nil)

	result := base64.StdEncoding.EncodeToString(mac)

	return result, nil
}
