package models

import "time"

type Student struct {
	ID int `json:"-"`

	// payload
	Grade  int `sql:",unique:gcn" sql:",notnull" json:"grade"`
	Class  int `sql:",unique:gcn" sql:",notnull" json:"class"`
	Number int `sql:",unique:gcn" sql:",notnull" json:"number"`

	// payload
	Name      string `sql:",notnull" sql:"type:varchar(7)" json:"name"`
	BarcodeID string `sql:",unique" sql:",notnull" sql:"type:char(5)" json:"barcode_id"`

	// payload
	Coin int `sql:",notnull" sql:"default:0" json:"coin"`

	UpdatedAt time.Time `json:"-"`
}

type Booth struct {
	ID int `json:"-"`

	// payload
	Name string `sql:",notnull" json:"name"`
	Coin int    `sql:"default:0" json:"coin"`

	Sellers  []*Seller  `json:"sellers,omitempty"`
	Products []*Product `json:"products,omitempty"`

	UpdatedAt *time.Time `json:"-"`
}

type Product struct {
	// payload
	ID    string `json:"id"`
	Name  string `sql:",notnull" json:"name"`
	Price int    `sql:"default:0" json:"price"`

	BoothID int `json:"-"`
}

type Seller struct {
	ID        int      `json:"-"`
	StudentID int      `sql:",unique" sql:",notnull" json:"-"`
	Student   *Student `json:"student,omitempty"`

	// payload
	LoginID string `sql:",unique" sql:",notnull" json:"login_id"`
	BoothID int    `sql:",notnull" json:"-"`
	Booth   *Booth `json:"booth,omitempty"`

	Pin string `sql:",notnull" sql:"type:char(6)" json:"-"`

	// Normal Seller: 0
	// Admin: 1
	Permission int `sql:"default:0" json:"-"`

	CreatedAt *time.Time `sql:"default:now()" json:"-"`
}

type Session struct {
	// payload
	ID string

	SellerID int
	Seller   *Seller

	CreatedAt time.Time `sql:"default:now()"`
	DeletedAt time.Time `pg:",soft_delete"`
}

type AccessLog struct {
	ID        int
	Date      time.Time
	Path      string
	SessionID int
	IP        string
	UserAgent string
}

type Order struct {
	// payload
	ID   string
	Date time.Time `sql:"default:now()"`

	// payload
	StudentID, SellerID            int
	SubTotal, Discount, GrandTotal int

	// payload
	Products []*Product

	// payload
	IsCanceled bool

	AccessLogID int
}
