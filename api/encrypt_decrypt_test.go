package api

import (
	"encoding/hex"
	"fmt"
)

// sm4 standard benchmark
const (
	plainPreset  = "0123456789abcdeffedcba9876543210"
	keyPreset    = "0123456789abcdeffedcba9876543210"
	cipherPreset = "681edf34d206965e86b3e94f536e4246002a8a4efa863ccad024ac0300bb40d2"
)

func ExampleSm4ECBEncrypt() {

	plain, err := hex.DecodeString(plainPreset)
	if err != nil {
		fmt.Println(err)
	}

	key, err := hex.DecodeString(keyPreset)
	if err != nil {
		fmt.Println(err)
	}

	cipher, err := Sm4ECBEncrypt(key, plain)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x", cipher)

	// Output: 681edf34d206965e86b3e94f536e4246002a8a4efa863ccad024ac0300bb40d2
}

func ExampleSm4ECBDecrypt() {

	cipher, err := hex.DecodeString(cipherPreset)
	if err != nil {
		fmt.Println(err)
	}

	key, err := hex.DecodeString(keyPreset)
	if err != nil {
		fmt.Println(err)
	}

	plain, err := Sm4ECBDecrypt(key, cipher)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%x", plain)

	// Output: 0123456789abcdeffedcba9876543210
}
