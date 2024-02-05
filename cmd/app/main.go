package main

import (
	"github.com/joho/godotenv"
	"github.com/reonardoleis/adventurer/internal/database"
	"github.com/reonardoleis/adventurer/internal/frontend"
)

func main() {
	godotenv.Overload(".env")

	database.NewJsonDatabase("adventurer.json")

	f := frontend.Get(frontend.Terminal)
	if err := f.Run(); err != nil {
		panic(err)
	}
}
