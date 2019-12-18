package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

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
	grade := sToI(r.URL.Query().Get("grade"))
	class := sToI(r.URL.Query().Get("class"))

	if grade < 1 || grade > 3 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid Grade"))
		return
	}

	if class < 1 || class > 9 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid Grade"))
		return
	}

	numbersStr := r.URL.Query()["numbers"]
	var numbers []int
	for _, s := range numbersStr {
		i := sToI(s)
		if i == 0 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(s + " is not integer"))
			return
		}
		numbers = append(numbers, i)
	}

	if len(numbers) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("No numbers"))
		return
	}

	var us []models.User
	err := db.Model(&us).Where("grade = ?", grade).Where("class = ?", class).
		Where("number in (?)", pg.In(numbers)).
		Order("grade").Order("class").Order("number").Select()

	if len(us) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("No users"))
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	t := template.Must(template.ParseFiles("generate-card/card.html"))
	err = t.Execute(w, us)
	if err != nil {
		log.Println(err)
	}

	return
}

func getUser(db *pg.DB, id string) (models.User, error) {
	u := models.User{}
	if len(id) != 4 {
		return u, errors.New(id + "is not 4 length")
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
