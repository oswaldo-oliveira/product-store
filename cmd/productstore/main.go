package main

import (
	"fmt"

	"github.com/oswaldo-oliveira/productstore/configs"
)

func main() {
	conf, _ := configs.LoadConfig(".")
	fmt.Printf(conf.DBDriver)
}
