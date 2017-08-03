package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)
var (
	key := []byte("c43ac86d84469030f28c0a9656b1c533")
	iv := []byte("2981eeca66b5c3cd")

)

func main() {

	ciphertext, err := ioutil.ReadFile("readme.enc") // hello.txt의 내용을 읽어서 바이트 슬라이스 리턴
	if err != nil {
		fmt.Println(err)
		return
	}


	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at this point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	fmt.Printf("%s\n", ciphertext)
	// Output: exampleplaintext

	//    //ciphertext, _ := hex.DecodeString("22277966616d9bc47177bd02603d08c9a67d5380d0fe8cf3b44438dff7b9")
	//    //        hex.d
	//    //hex.de
	//    block, err := aes.NewCipher(key)
	//    if err != nil {
	//        panic(err)
	//    }

	//    // The IV needs to be unique, but not secure. Therefore it's common to
	//    // include it at the beginning of the ciphertext.
	//    if len(ciphertext) < aes.BlockSize {
	//        panic("ciphertext too short")
	//    }
	//    //iv := ciphertext[:aes.BlockSize]
	//    ciphertext = ciphertext[aes.BlockSize:]

	//    stream := cipher.NewCFBDecrypter(block, iv)

	//    // XORKeyStream can work in-place if the two arguments are the same.
	//    stream.XORKeyStream(ciphertext, ciphertext)
	//    fmt.Printf("%s", ciphertext)
	//    // Output: some plaintext

	//    //    calc.

	//    //    fmt.Println(string(data)) // 문자열로 변환하여 data의 내용 출력

}
