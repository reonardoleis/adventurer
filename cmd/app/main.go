package main

import (
	"github.com/reonardoleis/adventurer/internal/database"
	"github.com/reonardoleis/adventurer/internal/frontend"
)

func main() {
	database.NewJsonDatabase("adventurer.json")
	f := frontend.Get(frontend.Terminal)
	if err := f.Run(); err != nil {
		panic(err)
	}
}
