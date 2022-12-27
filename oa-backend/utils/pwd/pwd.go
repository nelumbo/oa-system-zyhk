package pwd

import (
	"encoding/base64"

	"golang.org/x/crypto/scrypt"
)

func Scrypt(pwd string) (string, error) {
	const pwdLen = 10
	salt := []byte{1, 6, 9, 10, 99, 100, 199, 233}

	HashPwd, err := scrypt.Key([]byte(pwd), salt, 1<<15, 8, 1, pwdLen)
	if err != nil {
		return pwd, err
	}

	finalPwd := base64.StdEncoding.EncodeToString(HashPwd)
	return finalPwd, nil
}
