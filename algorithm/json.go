/*
Created by zhhigh
Date :2013-11-24
*/
package algorithm

import(
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

func ReadJson(jsonFile string)(interface {}){
	file, e := ioutil.ReadFile(jsonFile)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(file))

	//m := new(Dispatch)
	//var m interface{}
	var jsontype interface{}
	json.Unmarshal(file, &jsontype)
	fmt.Printf("Results: %v\n", jsontype)
	return jsontype
}
