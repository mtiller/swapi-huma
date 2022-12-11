package main

func main() {
	// Create CI container and populate it with services
	injector := dependencies()

	// Generate server application
	app := application(injector)

	// Run the CLI. When passed no arguments, it starts the server.
	app.Run()
}
