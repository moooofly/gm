package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/moooofly/gm/api"
	"github.com/moooofly/gm/sm3"
)

var (
	plaintext  = flag.String("plain", "", "plain text in normal string, e.g. 'hello world'")
	ciphertext = flag.String("cipher", "", "cipher text in hex format, e.g. '1234' means []byte{0x12, 0x34}")
	factor1    = flag.String("f1", "", "factor1 for sm3 to generate protect Key which used by sm4 ecb later")
	factor2    = flag.String("f2", "", "factor2 for sm3 to generate protect Key which used by sm4 ecb later")
)

func genProtectKey(factor1, factor2 string, pkLen int) []byte {
	var hwInfo []byte
	h := sm3.New()

	if factor2 == "" {
		salt := "set whatever you set"
		h.Write([]byte(salt))
		hwInfo = h.Sum(nil)
	} else {
		hwInfo = []byte(factor2)
	}

	byHash := make([]byte, 0)
	byHash = append(byHash, []byte(factor1)...)
	byHash = append(byHash, hwInfo...)

	h.Reset()
	h.Write(byHash)
	protectKey := h.Sum(nil)

	return protectKey[0:pkLen]
}

func main() {
	flag.Parse()

	// obtain 128bit key for sm4
	pk := genProtectKey(*factor1, *factor2, 16)

	if *plaintext != "" && *ciphertext != "" {
		ct, _ := hex.DecodeString(*ciphertext)
		pt, err := api.Sm4ECBDecrypt(pk, ct)
		if err != nil {
			log.Fatal(err)
		}
		if strings.ToLower(*plaintext) == strings.ToLower(string(pt)) {
			fmt.Printf("[Compare]: %s <--> %s  match\n", *plaintext, *ciphertext)
		} else {
			fmt.Printf("[Compare]: %s <--> %s  not match\n", *plaintext, *ciphertext)
		}
	} else if *plaintext != "" {
		pt := *plaintext
		ct, err := api.Sm4ECBEncrypt(pk, []byte(pt))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[SM4 ECB encrypt] %s --> %x\n", pt, ct)

	} else if *ciphertext != "" {
		ct, _ := hex.DecodeString(*ciphertext)
		pt, err := api.Sm4ECBDecrypt(pk, ct)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[SM4 ECB decrypt] %x --> %s\n", ct, pt)
	} else {
		fmt.Println("Neither '-plain' nor '-cipher' is set, Nothing to do.")
		return
	}

}
