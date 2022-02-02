package main

import (
	"encoding/json"
	"fmt"
)

//attributes starting with Capital letter will be included in JSON only
type student struct {
	First, Last string
	Email       string
	age         int
	IsMale      bool
}

//in this all all will be handeled by default
type studentMap map[string]interface{}

//adding json tags, these new names will be shown in json, - will be neglected,
type Teacher struct {
	First  string `json:"fname"`
	Last   string `json:"lname"`
	Email  string `json:"E-mail"`
	Age    int    `json:"-"`
	IsMale bool   `json:"-,"`
}

//general holder for json unmarshel
type General map[string]interface{}

func main() {

	//declaring struct, convert json to string while printing
	mujahid := student{"mujahid", "iqbal", "mjmalik66@gmail.com", 25, true}
	mujahidJSON, dd := json.Marshal(mujahid)
	fmt.Println(string(mujahidJSON), dd)

	//all fields will be available by default in this case, because now we have a map of interface{}
	sm := studentMap{"first": "mujahid", "last": "iqbal", "age": 25}
	smJSON, _ := json.Marshal(sm)
	fmt.Println(string(smJSON))

	//using json tags for field renaming
	nowTeacher := Teacher{"mujahid", "iqbal", "mjmalik66@gmail.com", 25, true}
	teachJSON, _ := json.MarshalIndent(nowTeacher, "", " ")
	fmt.Println(string(teachJSON))

	//check if json is valid
	isValid := json.Valid(teachJSON)
	fmt.Println(isValid)

	//unmarshel data
	var decodedTeacher Teacher
	unmarshelledData := json.Unmarshal(teachJSON, &decodedTeacher)
	fmt.Println(decodedTeacher)
	fmt.Println(unmarshelledData)

	var canStoreAnyJSON General
	unmarshelledDataTwo := json.Unmarshal(teachJSON, &canStoreAnyJSON)
	fmt.Println(canStoreAnyJSON)
	fmt.Println(unmarshelledDataTwo)
}
