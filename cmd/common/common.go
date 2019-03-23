package main

import (
	"fmt"
	"github.com/twuillemin/easy-sso-common/pkg/common"
)

// A dummy package so that Err* are not reported as unused
func main() {
	// Print the possible errors
	fmt.Println("Possible errors:")
	fmt.Println(common.ErrBadConfiguration)
	fmt.Println(common.ErrUnauthorized)
	fmt.Println(common.ErrUserNotFound)
	fmt.Println(common.ErrBadConfiguration)
	fmt.Println(common.ErrRefreshTokenNotFound)
	fmt.Println(common.ErrRefreshTooOld)
	fmt.Println(common.ErrNoAuthorization)
	fmt.Println(common.ErrMalformedAuthorization)
	fmt.Println(common.ErrSignatureInvalid)
	fmt.Println(common.ErrTokenMalformed)
	fmt.Println(common.ErrTokenTooOld)
	fmt.Println(common.ErrEmptyResponseFromServer)
}
