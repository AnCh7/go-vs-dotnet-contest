package main

func main() {
	initDB("anch:anch@tcp(127.0.0.1:3306)/employees")
	initServer()
}
