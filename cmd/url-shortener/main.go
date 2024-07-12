package main

import (
	"fmt"
	config2 "urlshortner/internal/config"
)

func main() {
	config := config2.MustLoad()
	fmt.Println(config)

	//TODO: init config: clean env

	//TODO: init logger: slog

	//TODO: init storage: sqlite3

	//TODO: init router: chi, render

	//TODO: run server
}
