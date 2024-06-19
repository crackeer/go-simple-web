package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// HMACSha1
//
//	@param key
//	@param message
//	@return string
func HMACSha1(key, message string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

// Sha1Sum
//
//	@param raws
//	@return string
func Sha1Sum(raws string) string {
	hash := sha1.New()
	hash.Write([]byte(raws))
	return hex.EncodeToString(hash.Sum(nil))
}

// GetFileMD5
//
//	@param pathName
//	@return string
//	@return error
func GetFileMD5(pathName string) (string, error) {
	f, err := os.Open(pathName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5hash.Sum(nil)), nil
}

// JsonDecode
//
//	@param raws
//	@param dest
//	@return error
func JsonDecode(raws string, dest interface{}) error {
	return json.Unmarshal([]byte(raws), dest)
}

// MD5 ...
func MD5(input string) string {
	sum := md5.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}
