package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// Set Global JWT
func SetGlobalJWTConfig(config JWTConfig) {

	if config.Skipper == nil {
		JTWConfigValidation.Skipper = DefaultJWTConfig.Skipper
	} else {
		JTWConfigValidation.Skipper = config.Skipper
	}

	// Signing Key Is Required
	if config.SigningKey == nil {
		panic("echo: jwt middleware requires signing key")
	}

	JTWConfigValidation.SigningKey = config.SigningKey

	if config.SigningMethod == "" {
		JTWConfigValidation.SigningMethod = DefaultJWTConfig.SigningMethod
	} else {
		JTWConfigValidation.SigningMethod = config.SigningMethod
	}

	if config.ContextKey == "" {
		JTWConfigValidation.ContextKey = DefaultJWTConfig.ContextKey
	} else {
		JTWConfigValidation.ContextKey = config.ContextKey
	}

	if config.TokenLookup == "" {
		JTWConfigValidation.TokenLookup = DefaultJWTConfig.TokenLookup
	} else {
		JTWConfigValidation.TokenLookup = config.TokenLookup
	}

	if config.AuthScheme == "" {
		JTWConfigValidation.AuthScheme = DefaultJWTConfig.AuthScheme
	} else {
		JTWConfigValidation.AuthScheme = config.AuthScheme
	}
	JTWConfigValidation.keyFunc = func(t *jwt.Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != JTWConfigValidation.SigningMethod {
			return nil, fmt.Errorf("Unexpected jwt signing method=%v", t.Header["alg"])
		}
		return config.SigningKey, nil
	}

}
