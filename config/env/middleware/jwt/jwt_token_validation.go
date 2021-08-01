package middleware

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"evermos-test/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ValidateToken(roles ...string) echo.MiddlewareFunc {

	parts := strings.Split(JTWConfigValidation.TokenLookup, ":")
	extractor := jwtFromHeader(parts[1], JTWConfigValidation.AuthScheme)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			/*if config.Skipper(c) {
				return next(c)
			}*/

			auth, err := extractor(c)

			if err != nil {
				//return echo.NewHTTPError(http.StatusBadRequest, err.Error())

				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":      http.StatusBadRequest,
					"code_type": "unAuthorized",
					"message":   err.Error(),
				})
			}
			token := new(jwt.Token)

			if token == nil {

				return c.JSON(http.StatusBadRequest, map[string]interface{}{
					"code":         http.StatusBadRequest,
					"code_message": "ErrorValidateToken",
					"message":      "Please Input Your Token",
				})

			}

			// Issue #647, #656
			claims, ok := JTWConfigValidation.Claims.(jwt.MapClaims)
			if ok && token.Valid && claims != nil {
				dataSession := claims["pxl"].(string)
				key := []byte(InternalEncryptKey)
				token, err = jwt.Parse(auth, JTWConfigValidation.keyFunc)

				decoded, err := hex.DecodeString(dataSession)
				if err != nil {

					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"code":         401,
						"code_message": "ErrorValidateToken",
						"message":      "Failed to Decrypt Data, Please Contact Your Administrator",
					})

				}

				plaintext, err := helper.Decrypt(decoded, key)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"code":         401,
						"code_message": "ErrorValidateToken",
						"message":      "Failed to Decrypt Data to Plaint Text, Please Contact Your Administrator",
					})

				}

				err = json.Unmarshal(plaintext, &UserSessionValidation)
				if err != nil {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{
						"code":         http.StatusBadRequest,
						"code_message": "ErrorValidateToken",
						"message":      "Failed to UnMarshal Data, Please Contact Your Administrator",
					})

				}

				if len(roles) > 0 {
					isAuthorized := helper.FindInArrayString(UserSessionValidation.Role, roles)
					if !isAuthorized {
						return c.JSON(http.StatusBadRequest, map[string]interface{}{
							"code":         http.StatusBadRequest,
							"code_message": "ErrorValidateToken",
							"message":      "You are not authorized to access this page, Please Contact Your Administrator",
						})
					}

					return next(c)
				}

				return next(c)
			}

			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":      http.StatusBadRequest,
				"code_type": "unAuthorized",
				"message":   "Please Contact your Administrator",
			})
		}
	}
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func jwtFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		l := len(authScheme)
		if len(auth) > l+1 && auth[:l] == authScheme {
			return auth[l+1:], nil
		}
		return "", errors.New("Missing or invalid jwt in the request header")
	}
}
