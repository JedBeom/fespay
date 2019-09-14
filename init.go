package main

import (
	"io"

	"github.com/go-pg/pg"
)

func copyStudentsCSV(db *pg.DB, file io.Reader) (err error) {
	_, err = db.CopyFrom(file,
		`COPY students(grade, class, number, name) FROM STDIN DELIMITER ',' CSV HEADER`)
	return
}
