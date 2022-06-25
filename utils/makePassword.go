package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MakePassword(password string) string {
	tempPassword := "<<<!@#$%^&152589" + password + "(*&^%$#@56372>>>"
	finalStr := ")*&^%$#@!13423423CodeThenDEcoDe" + tempPassword + "GetThisStringThenDecrypt!@#$%^&*()_03746253"
	m := md5.New()
	ErrorChecker(m.Write([]byte(finalStr)))
	newPassword := hex.EncodeToString(m.Sum(nil))
	return newPassword
}
