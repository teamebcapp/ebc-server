package auth

import (
	"crypto/sha256"
	postgres "ebc-server/common/db"
	"encoding/hex"
)

func CreateToken(userId, passwod, url string) string {
	sha := sha256.New()

	token := Auth{}
	postgres.PostgresConn.Where("id = ?", userId).Find(&token)

	if token.Id == userId {
		postgres.PostgresConn.Delete(&token)
	}
	token.Id = userId

	sha.Write([]byte(userId + passwod + url))
	token.Access_token = hex.EncodeToString(sha.Sum(nil))
	sha.Reset()
	sha.Write([]byte(userId + passwod + url + "next"))
	token.Refresh_token = hex.EncodeToString(sha.Sum(nil))
	token.Grant_type = "NORMAL_USER"
	postgres.PostgresConn.Create(&token)
	return token.Access_token
}

func RefreshToken(userId, passwod, url string) error {
	return nil
}

func ValidTokenUser(userId, access_token string) (bool, string) {
	token := Auth{}
	postgres.PostgresConn.Find(&token, userId)

	if token.Id != "" {
		if token.Access_token == access_token {
			return true, "success"
		}
		return false, "invalid token"
	}
	return false, "not created token"
}

func ValidToken(access_token string) (bool, string) {
	token := Auth{}
	postgres.PostgresConn.Where("access_token = ?", access_token).Find(&token)

	if token.Id != "" {
		return true, "success"
	}
	return false, "invalid token"
}
