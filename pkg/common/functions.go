package common

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)


// GetAuthenticationFromRequest locates and validates the authentication token in the given query. If the
// authentication is missing, in bad shape or cannot be validated, an error is returned. In case of success, this 
// function will return a fully filled AuthenticationInformation structure. In case of an error, the function will return 
// one the following error codes: ErrNoAuthorization, ErrMalformedAuthorization, ErrSignatureInvalid, ErrTokenMalformed
// or ErrTokenTooOld. 
//
// The function takes three parameters: 
// - request: the HTTP request to parse
// - publicKey: the public key to validate the token
// - detailedLogs: a boolean indicating if the detailed logs should be generated

func GetAuthenticationFromRequest(
	request *http.Request,
	publicKey *rsa.PublicKey,
	detailedLogs bool) (*AuthenticationInformation, error) {
		
	// If no request given
	if request == nil {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: parameter request was given null")
		}
		return nil, ErrBadParameters
	}

	// If no public key given
	if request == nil {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: parameter publicKey was given null")
		}
		return nil, ErrBadParameters
	}
	
	// If no authorization (8 is the minimum for Bearer + 1 char token)
	authorization := request.Header.Get("Authorization")
	if len(authorization) == 0 {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: No valid Authorization header")
		}
		return nil, ErrNoAuthorization
	}

	// If no authorization (8 is the minimum for Bearer + 1 char token)
	if len(authorization) < 8 {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Malformed Authorization header - Too short")
		}
		return nil, ErrMalformedAuthorization
	}

	// Check the format
	bearer := authorization[0:7]
	authorizationValue := authorization[7:]

	if bearer != "Bearer " {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Malformed authorization header - No Bearer found")
		}
		return nil, ErrMalformedAuthorization
	}

	// Split by the dots
	parts := strings.Split(authorizationValue, ".")
	if len(parts) != 3 {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Malformed Authorization header - Bad Bearer value")
		}
		return nil, ErrMalformedAuthorization
	}

	// Check the signature
	err := jwt.SigningMethodRS512.Verify(strings.Join(parts[0:2], "."), parts[2], publicKey)
	if err != nil {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Error while verifying the token - Bad signature")
		}
		return nil, ErrSignatureInvalid
	}

	// Read the token
	tokenString := authorizationValue
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Error while verifying the token - Malformed token")
		}
		return nil, ErrTokenMalformed
	}

	// Read the claims
	claims, ok := token.Claims.(*CustomClaims) // claims.User and claims.Roles are what we are interested in.
	if !ok {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Error while verifying the token - Malformed claims")
		}
		return nil, ErrTokenMalformed
	}

	// Read the timeout
	if claims.ExpiresAt < time.Now().Unix() {
		if detailedLogs {
			log.Printf("EasySSOCommon::GetAuthenticationFromRequest: Error while verifying the token - Token too old")
		}
		return nil, ErrTokenTooOld
	}

	return &AuthenticationInformation{
		User:  claims.User,
		Roles: claims.Roles,
		Token: authorizationValue,
	}, nil
}
