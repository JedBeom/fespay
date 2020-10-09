package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/JedBeom/fespay/models"

	"github.com/go-pg/pg"
	_ "github.com/joho/godotenv/autoload"
)

var (
	db *pg.DB
)

func connectDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	db = pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
}

func main() {
	connectDB()
	fs := http.FileServer(http.Dir("generate-card"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":5000", nil); err != nil {
		panic(err)
	}
}

func sToI(a string) int {
	b, _ := strconv.Atoi(a)
	return b
}

func index(w http.ResponseWriter, r *http.Request) {
	userType := sToI(r.URL.Query().Get("type"))
	ids := strings.Split(r.URL.Query().Get("ids"), "|")
	needsSort := r.URL.Query().Get("sort")

	if needsSort == "true" {
		sort.Strings(ids)
	}

	if len(ids) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("no ids"))
	}

	t := template.Must(template.ParseFiles("generate-card/card.html"))

	if userType == 2 {
		us, err := getTeachers(db)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("teacher: " + err.Error()))
			return
		}
		err = t.Execute(w, us)
		if err != nil {
			log.Println(err)
		}
		return
	}

	if userType == 4 {
		us := []models.User{
			{
				Name:     "김노관",
				CardCode: "T0013",
			},
			{
				Name:     "정혜인",
				CardCode: "T0015",
			},
			{
				Name:     "유재민",
				CardCode: "T0016",
			},
			{
				Name:     "남지혜",
				CardCode: "T0018",
			},
			{
				Name:     "이보라",
				CardCode: "T0022",
			},
			{
				Name:     "Jefferson",
				CardCode: "T0024",
			},
			{
				Name:     "박현숙",
				CardCode: "T0026",
			},
		}

		err := t.Execute(w, us)
		if err != nil {
			log.Println(err)
		}
		return
	}

	var us []models.User
	for _, id := range ids {
		u, err := getStudent(db, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(id + ": " + err.Error()))
			return
		}

		us = append(us, u)
	}

	err := t.Execute(w, us)
	if err != nil {
		log.Println(err)
	}

	return
}

func getStudent(db *pg.DB, id string) (models.User, error) {
	u := models.User{}
	if len(id) != 4 {
		return u, errors.New(id + " is not 4 length")
	}

	grade := int(id[0] - '0')
	class := int(id[1] - '0')
	number := sToI(id[2:])

	if grade < 1 || grade > 3 {
		return u, errors.New("grade invalid")
	}

	if class < 1 || class > 9 {
		return u, errors.New("class invalid")
	}

	if number < 1 || number > 32 {
		return u, errors.New("number invalid")
	}

	err := db.Model(&u).Where("grade = ?", grade).
		Where("class = ?", class).Where("number = ?", number).Select()

	return u, err
}

func getTeachers(db *pg.DB) ([]models.User, error) {
	var us []models.User

	err := db.Model(&us).Where("type = 2").Select()
	return us, err
}
