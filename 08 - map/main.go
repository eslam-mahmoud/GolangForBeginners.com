// A map is a data type maps keys to values
// The zero value of a map is nil
package main

import "fmt"

func main() {
	// init map if you have initial values
	m := map[string]bool{
		"X": true,
		"Y": false,
	}
	fmt.Println("m", m)

	// init a map, no need for length
	players := make(map[string]int)
	fmt.Println("players", players)

	// add value to map
	players["Eslam"] = 99
	players["Mahmoud"] = 100
	fmt.Println("players", players)

	// get value from map
	fmt.Println(players["Eslam"])

	// delete value from map
	delete(players, "Mahmoud")
	// delete woring key does not affect the map
	delete(players, "x")
	fmt.Println(players)

	// check if key exist and get value
	// value will be the default value if key does not exist
	v, ok := players["y"]
	fmt.Println("Missing key, V:", v, "Ok:", ok)
	v, ok = players["Eslam"]
	fmt.Println("Valid key, V:", v, "Ok:", ok)
	v = players["y"]
	fmt.Println("Valid key, V:", v)

	// map of slices
	mp := map[string][]string{
		"Admins": []string{"Eslam", "Mahmoud"},
		"Users":  []string{"X", "Y", "Z"},
	}
	fmt.Println("mp", mp)

}
