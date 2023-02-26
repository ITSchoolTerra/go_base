package main

import "strings"

func main() {
	strings.Contains("A,b,c,d,e,f;", "e f")
	// результат: false

	strings.Index("012345", "234")
	// результат: 2

	strings.Replace("A,b,c,d,e,f;", ",e,f", "", -1)
	// результат: "A,b,c,d;"

	strings.Replace("A,b,c,d,e,f;", ",", "-", -1)
	// результат: "A-b-c-d-e-f;"

	strings.TrimSuffix("A,b,c,d;", ";")
	// результат: "A,b,c,d"

	strings.Split("A,b,c,d", ",")
	// результат: []string{"A","b","c","d"}

	strings.ToLower("A,b,c,d")
	// результат: "a,b,c,d"

	strings.ToUpper("a,b,c,d")
	// результат: "A,B,C,D"

}
