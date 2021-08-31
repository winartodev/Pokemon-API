package main

func main() {
	c := Controller{}
	c.db = c.connect()

	defer c.db.Close()

	c.endpointsHandler()

	run()
}