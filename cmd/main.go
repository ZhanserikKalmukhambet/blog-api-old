package main

import (
	"fmt"
	"github.com/zhanserikKalmukhambet/blog-api/internal/config"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("%#v", cfg))
}
