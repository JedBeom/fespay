package models

import "time"

type Student struct {
	ID int

	Grade int `sql:",unique:gcn"`
	Class int `sql:",unique:gcn"`
	Number int `sql:",unique:gcn"`

	Name string
	BarcodeID string `sql:",unique"`

	Coin int

	UpdatedAt time.Time
}

type Seller struct {
	ID int
	StudentID int `sql:",unique"`
	Student *Student

	LoginID string `sql:",unique"`
	PinCode int

	BoothID int
	Booth *Booth

	CreatedAt *time.Time `sql:"default:time.Now()"`
}

type Product struct {
	ID int
	Name string
	Price int
}

type Booth struct {
	ID int
	Name string
	Sellers []*Seller
	Products []*Product

	Coin int

	UpdateAt *time.Time
}

type Session struct {
	ID int
	SellerID int
	Seller *Seller

	UUID string
	CreatedAt time.Time
	DeletedAt time.Time `pg:",soft_delete"`
}

type AccessLog struct {
	ID int
	Date time.Time
	SellerID int
	Path string
	IP string
	UserAgent string
}

type PayLog struct {
	ID int
	Date time.Time

	StudentID, SellerID int
	Products []*Product
	Price, Discount, Total int

	IsCanceled bool
	AccessLogID int
}
