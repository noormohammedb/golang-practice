package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fileName := "./fileCreatedWithProgram.txt"
	contentToFile := "lorem ipsum dollor"

	fmt.Println("Working With Files")

	file, err := os.Create(fileName)

	// if err != nil {
	// 	panic(err)
	// }
	checkNotNill(err)

	wrightLength, err := io.WriteString(file, contentToFile)
	checkNotNill(err)

	fmt.Println("length : ", wrightLength)

	defer file.Close()

	// defer os.Remove(fileName) // for delete created file

	readFileContent(fileName)

}

func readFileContent(fileNameWithPath string) {
	fileRawData, err := ioutil.ReadFile(fileNameWithPath)
	checkNotNill(err)
	fileStringData := string(fileRawData)

	fmt.Println("file Raw Data : ", fileRawData)

	fmt.Println("file string data : ", fileStringData)

}

func checkNotNill(err error) {
	if err != nil {
		panic(err)
	}
}
