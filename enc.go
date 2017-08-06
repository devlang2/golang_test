package main

import (
	//	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"strconv"

	//	"github.com/davecgh/go-spew/spew"
	//	"encoding/binary"
	"fmt"
	//	"math"
)

func Encrypt(key, iv, data []byte) ([]byte, error) {
	return key, nil
}

func Decrypt(key, iv, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""), err
	}

	if len(data) < aes.BlockSize {
		return []byte(""), fmt.Errorf("ciphertext too short")
	}

	if len(data)%aes.BlockSize != 0 {
		return []byte(""), fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	data_temp := make([]byte, len(data))
	copy(data_temp, data)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data_temp, data)

	return data_temp, nil
	//    return bytes.Trim(data_temp, "\x0e"), nil

}

func ByteToFloat64(b []byte) float64 {
	f, _ := strconv.ParseFloat(string(b), 64)
	return f
}

func ByteToInt64(b []byte) int64 {
	str := string(b)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}
