package main 

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}

	wording := string(sd)
	fmt.Println("result after decoding:", wording)

	/* puzzle running last char -> first char */

	for _, r := range []rune(strings.ReplaceAll(wording, ":", " ")) {
		whatIsIt = string(r) + whatIsIt	
	}

	fmt.Println("result:", whatIsIt)
}