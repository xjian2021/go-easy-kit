package codec

import (
	"crypto/md5"
	"crypto/rc4"
	"encoding/hex"
	"encoding/json"
)

func ToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func RC4(key []byte, src []byte) ([]byte, error) {
	c, err := rc4.NewCipher(key)
	dst := make([]byte, len(src))
	if err != nil {
		return dst, err
	}
	c.XORKeyStream(dst, src)
	return dst, nil
}

func Rc4Encrypt(key []byte, src interface{}) (string, error) {
	bs, err := json.Marshal(src)
	if err != nil {
		return "", err
	}

	rcBs, err := RC4(key, bs)
	if err != nil {
		return "", err
	}
	secHex := hex.EncodeToString(rcBs)
	return secHex, nil
}

func Rc4EncryptString(key []byte, src string) (string, error) {
	rcBs, err := RC4(key, []byte(src))
	if err != nil {
		return "", err
	}
	secHex := hex.EncodeToString(rcBs)
	return secHex, nil
}

func Rc4DecryptString(key []byte, src string) (string, error) {
	secHex, err := hex.DecodeString(src)
	if err != nil {
		return "", err
	}
	secretBs, err := RC4(key, secHex)
	return string(secretBs), err
}
