package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	A string
	B int
}

func main() {
	cfg := Config{"1234", 23}
	fmt.Println(cfg)
	data, err := json.Marshal(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	cfg2 := Config{}
	err = json.Unmarshal(data, &cfg2)
	fmt.Println("cfg2", cfg2)
}
