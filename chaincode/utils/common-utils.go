package utils

import (
	"encoding/hex"
	"math/rand"
)

const length = 8

const (
	UserRoleADMIN      = "ADMIN"
	UserRoleINSPECTOR  = "INSPECTOR"
	UserRoleHARVESTOR  = "PRODUCER"
	UserRoleEXPORTOR   = "EXPORTOR"
	UserRoleIMPORTOR   = "IMPORTOR"
	UserRolePROCESSOR  = "PROCESSOR"
	UserStatusACTIVE   = "ACTIVE"
	UserStatusINACTIVE = "INACTIVE"
)

// generateSecureRandomString generates a secure random string of a specified length.
func GenerateSecureRandomString() (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
