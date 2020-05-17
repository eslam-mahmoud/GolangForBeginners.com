// JSON is a format that widely used to communicate between application and APIs
package main

import (
	"encoding/json"
	"fmt"
)

// type with unexported field
type human struct {
	name      string
	Expertise []string
}

// type with exported field & annotations
// annotations used mostly in JSON and DB
type human2 struct {
	Name      string   `json:"name"`
	Expertise []string `json:"expertise"`
}

func main() {
	{
		// slice to JSON
		expertise := []string{"GoLang", "PHP", "JS"}
		expertiseJSON, err := json.Marshal(expertise)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(expertiseJSON))
	}

	{
		// map to JSON
		expertise := map[string]string{"GoLang": "excellent", "PHP": "excellent", "JS": "Very good"}
		expertiseJSON, err := json.Marshal(expertise)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(expertiseJSON))
	}

	{
		// struct with unexported field to JSON
		eslam := human{
			name:      "Eslam",
			Expertise: []string{"GoLang", "PHP", "JS"},
		}
		eslamJSON, _ := json.Marshal(eslam)
		fmt.Println(string(eslamJSON))
	}

	{
		// with exported field & annotations
		eslam := human2{
			Name:      "Eslam",
			Expertise: []string{"GoLang", "PHP", "JS"},
		}
		eslamJSON, _ := json.Marshal(eslam)
		fmt.Println(string(eslamJSON))
	}

	{
		// pointer to JSON
		eslam := &human2{
			Name:      "Eslam",
			Expertise: []string{"GoLang", "PHP", "JS"},
		}
		eslamJSON, _ := json.Marshal(eslam)
		fmt.Println(string(eslamJSON))
	}

	{
		// JSON string to struct
		s := `{"name":"Eslam","expertise":["GoLang","PHP","JS"]}`
		b := []byte(s)    // we can only send bytes to json.Unmarshal
		eslam := human2{} // destination var
		fmt.Println(eslam)
		err := json.Unmarshal(b, &eslam)
		if err != nil {
			panic(err)
		}
		fmt.Println(eslam)
	}
	{
		// JSON string to struct
		s := `{"name":"Eslam","expertise":["GoLang","PHP","JS"]}`
		b := []byte(s)   // we can only send bytes to json.Unmarshal
		eslam := human{} // destination var
		err := json.Unmarshal(b, &eslam)
		if err != nil {
			panic(err)
		}
		fmt.Println(eslam)
	}
}
