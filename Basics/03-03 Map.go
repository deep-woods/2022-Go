package main

import "fmt"

func main() {
	var ascii_codes map[string]int
	ascii_codes["A"] = 65	  // panic: assignment to entry in nil map
	fmt.Println(ascii_codes)


	// Initialise: Create a map with values
	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	_, found := ascii_codes["B"]  // 0, found=false
	if found {
			fmt.Println("key B was not found")
	} // no output code.


	codes := map[string]int{"ko":1, "fr":2, "en":3}
	val, found := codes["ko"]
	fmt.Println(val, found)     // 1 true
	val, found = codes["ab"]
	fmt.Println(val, found)		// 0 false


		/*
	- create a map using make() function 
	  with key data type as string, 
	  and value data type as int.
	- add the following key_value pairs to it
	*/
	var m = make(map[string]int)
	m["A"] = 65
	m["F"] = 70
	m["K"] = 75
	fmt.Println(m)	// map[A:65 F:70 K:75]
	delete(m, "F")
	fmt.Println(m)  // map[A:65 K:75]

	// Override an existing map
	ascii_codes := make(map[string]int)
	ascii_codes["C"] = 65
	ascii_codes["O"] = 70
	ascii_codes["D"] = 75
	fmt.Println(ascii_codes) 	// map[C:65 D:75 O:70]

	ascii_codes = make(map[string]int)
	ascii_codes["E"] = 85
	fmt.Println(ascii_codes)	// map[E:85]


	// Update your map.
	langs := map[string]string{"ko": "Korean", "en": "English", "po": "portoges"}
	fmt.Println(langs)
	langs["po"] = "portugues" // Update the value here.
	fmt.Println(langs)
	// map[en:English ko:Korean po:portoges]
	// map[en:English ko:Korean po:portugues]
}