package main

import "fmt"

func main() {
	proLang := make(map[string]string)

	proLang["js"] = "JavaScript"
	proLang["py"] = "Python"
	proLang["cpp"] = "C Pluse Pluse"

	fmt.Printf("type of proLang: %T\n", proLang)
	fmt.Println(proLang)

	var tKey, dKey string = "js", "cpp"

	delete(proLang, dKey)
	fmt.Println(proLang)

	fmt.Println("targeting a key : ", proLang[tKey])

	// for _, value := range proLang {  // for only iterate values
	for key, value := range proLang {
		fmt.Println(key, value)
	}
}
