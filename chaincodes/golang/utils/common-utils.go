package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

const length = 8
const SharedSecret = "your-shared-secret"

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
func GenerateDeterministicID(sharedSecret string) string {
	// Hash the shared secret to get a consistent ID
	hash := sha256.New()
	hash.Write([]byte(sharedSecret))
	return hex.EncodeToString(hash.Sum(nil))
}
