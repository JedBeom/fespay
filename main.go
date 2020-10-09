package main

import "github.com/teris-io/shortid"

type Map map[string]interface{}

const JSONIndent = "  "
const Namespace = "fespay.aligo.space"
const AdminBoothID = "fespay"

func init() {
	alphabet := `언제까지나,아이돌-꿈란루기위해있는것니@을처음바라고서오늘얼마온걸워눈부신저편ALSTR내일765NEViDOIor프로듀스터`
	sid, err := shortid.New(1, alphabet, 505)
	if err != nil {
		panic(err)
	}

	shortid.SetDefault(sid)
}

func main() {
	connectDB()
	insertUsersIfNotExists(db)
	run()
}
