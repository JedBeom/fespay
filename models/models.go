package models

import "time"

type Student struct {
	ID int `json:"-"`

	Grade  int `sql:",unique:gcn" sql:",notnull" json:"grade"`
	Class  int `sql:",unique:gcn" sql:",notnull" json:"class"`
	Number int `sql:",unique:gcn" sql:",notnull" json:"number"`

	Name      string `sql:",notnull" sql:"type:varchar(7)" json:"name"`
	BarcodeID string `sql:",unique" sql:",notnull" sql:"type:char(5)" json:"barcode_id"`

	Coin int `sql:",notnull" sql:"default:0" json:"coin"`

	UpdatedAt time.Time `json:"-"`
}

type Booth struct {
	ID int `json:"-"`

	Name string `sql:",notnull" json:"name"`
	Coin int    `sql:"default:0" json:"coin"`

	Sellers  []*Seller  `json:"sellers,omitempty"`
	Products []*Product `json:"products,omitempty"`

	UpdatedAt time.Time `json:"-"`
}

type Product struct {
	ID    string `json:"id"`
	Name  string `sql:",notnull" json:"name"`
	Price int    `sql:"default:0" json:"price"`

	BoothID int `json:"-"`
}

type Seller struct {
	ID        int      `json:"-"`
	StudentID int      `sql:",unique" sql:",notnull" json:"-"`
	Student   *Student `json:"student,omitempty"`

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
	ID string

	SellerID int
	Seller   *Seller

	CreatedAt time.Time `sql:"default:now()"`
	DeletedAt time.Time `pg:",soft_delete"`
}

type AccessLog struct {
	ID   string    `sql:",notnull"`
	Date time.Time `sql:"default:now(),notnull"`
	Path string    `sql:",notnull"`

	SessionID string
	Session   *Session

	IP        string `sql:",notnull"`
	UserAgent string `sql:",notnull"`
}

type Order struct {
	ID   string    `json:"id" sql:",notnull"`
	Date time.Time `sql:"default:now()" json:"date" sql:",notnull"`

	StudentID int      `json:"-" sql:",notnull"`
	Student   *Student `json:"student,omitempty"`

	BoothID int    `json:"-" sql:",notnull"`
	Booth   *Booth `json:"-"`

	SellerID int     `json:"-" sql:",notnull"`
	Seller   *Seller `json:"seller,omitempty"`

	SubTotal   int `json:"sub_total" sql:",notnull"`
	Discount   int `json:"discount" sql:",notnull" sql:"default:0"`
	GrandTotal int `json:"grand_total" sql:",notnull"`

	Products []*Product `json:"products,omitempty" pg:"many2many:orders_to_products"`

	IsCanceled bool `json:"is_canceled" sql:"default:false"`

	AccessLogID string     `json:"-" sql:",notnull"`
	AccessLog   *AccessLog `json:"-"`
}

type OrderToProduct struct {
	tableName struct{} `sql:"orders_to_products"`
	OrderID   string   `sql:",pk"`
	ProductID string   `sql:",pk"`
	Amount    int      `sql:"default:1"`
}
