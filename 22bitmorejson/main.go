package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags, omitempty"`
}

func main() {
	fmt.Println("Welcome to json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []courses{
		{"Python Programming", 499, "Udemy", "python123", []string{"web-dev", "python", "ml"}},
		{"C Programming", 49, "Udemy", "c123", []string{"web-dev", "c", "core-lang"}},
		{"Java Programming", 99, "Udemy", "java123", []string{"web-dev", "python", "ml"}},
		{"Data Science", 999, "Udemy", "data123", []string{"web-dev", "python", "ml"}},
	}

	// package this data as Json data

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	 {
        "name": "Python Programming",
        "price": 499,
        "platform": "Udemy",
        "tags": [
                "web-dev",
                "python",
                "ml"
        ]
    }
	`)

	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Println((lcoCourse))
	} else {
		fmt.Println("JSON was invalid")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println(myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %s and value is %v and Type is : %T\n", k, v, v)
	}
}
