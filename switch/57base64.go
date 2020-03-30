package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // YWJjMTIzIT8kKiYoKSctPUB+


	// 解码可能会返回错误，如果不确定输入信息格式是否正确，那么久需要进行错误检查了
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec)) // abc123!?$*&()'-=@~

	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data)) // YWJjMTIzIT8kKiYoKSctPUB-

	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec)) // abc123!?$*&()'-=@~

}
