package security

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/dlclark/regexp2"
)

const (
	regex = `^(?=.*[0-9])[a-zA-Z0-9!@#$%^&*]{7,30}$`
)

var (
	ErrBadPassword = errors.New("bad password")
)

func Hash(str string) string {
	hash := sha256.Sum256([]byte(str))

	return hex.EncodeToString(hash[:])
}

func ValidatePassword(password string) error {
	regex, err := regexp2.Compile(regex, 0)
	if err != nil {
		return err
	}

	isMatch, err := regex.MatchString(password)
	if !isMatch {
		return err
	}

	if isMatch {
		return nil
	}

	return ErrBadPassword
}
