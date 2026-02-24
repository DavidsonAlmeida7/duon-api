package helper

import (
	"encoding/json"
	"fmt"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
)

func PrintLikeJson(element interface{}) {
	bytes, _ := json.MarshalIndent(element, "", " ")
	fmt.Println(string(bytes))
}

func PrintFormatted(element interface{}) {
	fmt.Println("===========================")
	fmt.Printf("%+v", element)
	fmt.Println("\n===========================")
}

func Debug(label string, element interface{}) {
	fmt.Printf("\n%s[DEBUG - %s]%s\n", ColorRed, label, ColorReset)
	bytes, _ := json.MarshalIndent(element, "", "  ")
	fmt.Println(string(bytes))
	fmt.Printf("%s[END DEBUG - %s]%s\n\n", ColorRed, label, ColorReset)
}
