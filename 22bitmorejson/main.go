package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"courseName"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"keywords,omitempty"`
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	var lcoCourse []course = []course{
		// lcoCourse := []course{
		{"React", 149, "learncodeonline.com", "pass123", []string{"js", "web", "spa"}},
		{"Flutter", 249, "learncodeonline.com", "123pass", []string{"dart", "android", "ios"}},
		{"Angular", 199, "learncodeonline.com", "xyz123", []string{"js", "web", "spa"}},
		{"Node", 299, "learncodeonline.com", "abc123", nil},
	}

	fmt.Println(lcoCourse)

	// finalJson, err := json.MarshalIndent(lcoCourse, "", "	")
	finalJson, err := json.Marshal(lcoCourse)
	checkNotNill(err)

	// fmt.Println(string(finalJson))
	fmt.Printf("%s", finalJson)

}

func checkNotNill(err error) {
	if err != nil {
		panic(err)
	}
}
