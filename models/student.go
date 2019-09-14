package models

import (
	"github.com/go-pg/pg"
)

func StudentByBarcodeID(db *pg.DB, barcodeID string) (student Student, err error) {
	err = db.Model(&student).Where("barcode_id = ?", barcodeID).Select()
	return
}
