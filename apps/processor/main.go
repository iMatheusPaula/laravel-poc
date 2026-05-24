package main

func main() {
	config := LoadConfig()
	startSubscriber(config)
}
