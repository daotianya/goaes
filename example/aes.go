package main

import (
	"fmt"
	"goaes"
)

func main() {
	aesEnc := goaes.NewEnc()
	aesEnc.Iv = `sdf234wef34efrfT`
	aesEnc.Key = `aaC5p6c5L2g6KeJ5`
	source := `i want go`
	des, err := aesEnc.Encrypt(source)
	if err != nil {
		fmt.Println("hahaha watele")
	}
	resource, err := aesEnc.Decrypt(des)
	fmt.Println(resource)
}
