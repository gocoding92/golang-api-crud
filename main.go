package main

import "go-echo-api/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":3030"))
}
