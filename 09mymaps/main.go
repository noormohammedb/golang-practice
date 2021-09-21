package main

import "fmt"

func main() {
	proLang := make(map[string]string)

	proLang["js"] = "JavaScript"
	proLang["py"] = "Python"
	proLang["cpp"] = "C Pluse Pluse"

	fmt.Println(proLang)
	fmt.Printf("type of proLang: %T\n", proLang)

	var tKey string = "js"

	fmt.Println("targeting a key : ", proLang[tKey])

	for key, value := range proLang {
		fmt.Println(key, value)
	}
}
