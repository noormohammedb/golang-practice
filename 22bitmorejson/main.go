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
	// EncodeJson()
	DecodeJson()
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

func DecodeJson() {
	JsonDataFromWeb := []byte(`{
  "courseName": "React",
  "price": 149,
  "website": "learncodeonline.com",
  "keywords": ["js","web","spa"]
	}`)

	// var lcoCourse map[string]interface{}
	var lcoCourse course

	isValidJson := json.Valid(JsonDataFromWeb)
	if isValidJson {
		fmt.Println("json in valid")
		json.Unmarshal(JsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("json in invalid")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(JsonDataFromWeb, &myOnlineData)
	fmt.Printf("\n%#v\n", myOnlineData)

	for keyOfVal, value := range myOnlineData {
		fmt.Printf("%#v %#v \n", keyOfVal, value)
	}
}

func checkNotNill(err error) {
	if err != nil {
		panic(err)
	}
}
