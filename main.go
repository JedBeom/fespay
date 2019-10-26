package main

type KeyValue map[string]interface{}

const JSONIndent = "  "
const Namespace = "fespay.aligo.space"

func main() {
	connectDB()
	insertUsersIfNotExists(db)
	run()
}
