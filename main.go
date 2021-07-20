package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonString := `
	  [{ 
		  "Name": "Call of Duty", 
	      "Rating": 8.4,
		  "Users": [
			  {
				  "Name": "John",
				  "Rating": 10
			  },
			  {
				"Name": "Peter",
				"Rating": 5 
			  }
		  ]
	   },
	   { 
		"Name": "Sonic", 
		"Rating": 9.4,
		"Users": [
			{
				"Name": "John",
				"Rating": 4
			},
			{
			  "Name": "Peter",
			  "Rating": 9
			}
		]
	   }
	  ]`
	var data interface{}
	json.Unmarshal([]byte(jsonString), &data)
	firstElement := data.([]interface{})[0]
	fmt.Println(firstElement)
	printJSON(data)
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object:")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type", vv)
	}
}
