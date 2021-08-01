package middleware

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"evermos-test/database/entity"
	"evermos-test/helper"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenetarateToken(user *entity.User) (string, error) {

	hexid := fmt.Sprintf("%x", string(user.Id))

	UserSessionValidation = &UserSession{
		UserId:   hexid,
		Fullname: user.FullName,
		NIM:      user.NIM,
		Email:    user.Email,
		Role:     user.Role,
	}

	jsonString, err := json.Marshal(UserSessionValidation) // Set Data
	if err != nil {
		fmt.Println("Error encoding JSON")
		return "", errors.New("Error encoding JSON")
	}

	key := []byte(InternalEncryptKey)

	encrypt, err := helper.Encrypt(jsonString, key)
	if err != nil {
		return "", err
	}

	pxl := hex.EncodeToString(encrypt)

	currentTimestamp := time.Now().UTC().Unix()
	var ttl int64 = (3600 * 10) // expired time in second
	// md5 of sub & iat
	h := md5.New()
	io.WriteString(h, hexid)
	io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.

	subRandom, _ := helper.GenerateRandomString(`[A-Z0-9]{7}-[A-Z0-9]{7}-[A-Z0-9]{7}-[A-Z0-9]{7}`)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subRandom,
		"iat": currentTimestamp,
		"exp": currentTimestamp + ttl,
		"nbf": currentTimestamp,
		"jti": h.Sum(nil),
		"pxl": pxl,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
