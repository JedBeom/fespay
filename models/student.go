package models

import (
	"io"

	"github.com/go-pg/pg"
)

func StudentByBarcodeID(db *pg.DB, barcodeID string) (student Student, err error) {
	err = db.Model(&student).Where("barcode_id = ?", barcodeID).Select()
	return
}

func CopyStudentsCSV(db *pg.DB, file io.Reader) (err error) {
	_, err = db.CopyFrom(file,
		`COPY students(grade, class, number, name) FROM STDIN DELIMITER ',' CSV HEADER`)
	return
}
