package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON")
	//EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abs123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}
	//package this data as JSON data
	//finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
func DecodeJSON(){
       
	jsonDataFromWeb := []byte(`
	     {
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "LearnCodeOnline.in",
                "tags": ["web-dev","js"]
        }
	
	`)
      var lcoCourse course

	  checkValid := json.Valid(jsonDataFromWeb) 

	  if checkValid{
		fmt.Println("JSON Data is Valid")
		json.Unmarshal(jsonDataFromWeb,&lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	  } else{
		fmt.Println("INVALID")
	  }

	  // cases where add data to key value pair
	  var myOnlineData map[string]interface{}
	  json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	  fmt.Printf("%#v\n",myOnlineData)

	  for k,v := range myOnlineData{
		fmt.Printf("The Key is : %v The Value is : %v and the Type is : %T\n",k,v,v)
	  }
}