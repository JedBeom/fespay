package models

import (
	"net/http"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/lithammer/shortuuid"
)

func InsertProducts(db *pg.DB, boothID int, products *[]Product) error {
	for i := range *(products) {
		(*products)[i].ID = shortuuid.New()
		(*products)[i].BoothID = boothID
	}

	return db.Insert(products)
}

func SumAndValidateProducts(db *pg.DB, boothID int, productIDs []string) (total int, products []*Product, err error) {
	for _, productID := range productIDs {
		p := Product{ID: productID}
		if err := db.Model(&p).WherePK().Select(); err != nil {
			return 0, nil, err
		}

		if p.BoothID != boothID {
			return 0, nil, echo.NewHTTPError(http.StatusBadRequest, "Product and Booth mismatch")
		}

		products = append(products, &p)
		total += p.Price
	}

	return
}

func (order *Order) createOToPs(tx *pg.Tx) (err error) {
	m := make(map[string]int)
	for _, p := range order.Products {
		m[p.ID]++
	}

	for pID, amount := range m {
		oToP := OrderToProduct{
			OrderID:   order.ID,
			ProductID: pID,
			Amount:    amount,
		}

		if err := tx.Insert(&oToP); err != nil {
			return echo.ErrInternalServerError
		}
	}

	return nil
}
