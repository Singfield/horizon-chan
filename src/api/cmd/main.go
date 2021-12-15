package main

import "log"

func Run() error {
	app := fiber.New()
	err := app.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
