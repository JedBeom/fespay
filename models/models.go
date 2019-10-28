package models

import "time"

type UserType int

const (
	UserStudent UserType = iota + 1
	UserTeacher
)

type Status int

const (
	StatusWorking   = iota + 1
	StatusFrozen    // Frozen(동결): 로그인, 조회 등은 가능. 결제, 수정 등의 행동 불가
	StatusSuspended // Suspended(정지): 모든 것이 불가. 유저 입장에서는 삭제된 것과 마찬가지.
)

type User struct {
	ID string

	WalletID string
	Wallet   *Wallet
	BoothID  string
	Booth    *Booth
	LoginID  string
	Password string

	Type UserType `sql:",notnull"`

	Grade    int    `sql:",unique:gcn"`
	Class    int    `sql:",unique:gcn"`
	Number   int    `sql:",unique:gcn"`
	Name     string `sql:",unique" sql:",notnull"`
	CardCode string `sql:",unique" sql:",notnull" sql:"type:char(5)"`

	Status    Status `sql:"default:1"`
	UpdatedAt time.Time
}

type Booth struct {
	ID          string
	WalletID    string
	Wallet      *Wallet
	Name        string `sql:",unique" sql:",notnull" sql:"type:varchar(15)"`
	Description string `sql:"type:varchar(200)"`
	Staffs      []*User
	Status      Status `sql:"default:1"`
	UpdatedAt   time.Time
}

type OwnerType int

const (
	OwnerUser = iota + 1
	OwnerBooth
)

type Wallet struct {
	ID        string
	OwnerType OwnerType `sql:",notnull"`
	OwnerID   string    `sql:",notnull" sql:",unique"`
	Coin      int       `sql:"default:0"`
	UpdatedAt time.Time
}

type Order struct {
	ID            string `json:"id"`
	StaffID       string `sql:",notnull" json:"staffID"`
	Staff         *User
	FromID        string `sql:",notnull" json:"fromID"`
	From          *Wallet
	ToID          string
	To            *Wallet
	Amount        int `sql:",notnull"`
	RefundOrderID string
	RefundOrder   *Order

	AccessLogID string `sql:",notnull"`
	AccessLog   *AccessLog
	CreatedAt   time.Time `sql:"default:now()"`
	ClosedAt    time.Time
}

type AccessLog struct {
	ID        string
	SessionID string
	Session   *Session
	IP        string
	Method    string
	Path      string
	CreatedAt time.Time
}

type Session struct {
	ID        string
	UserID    string
	User      *User
	UserAgent string
	CreatedAt time.Time
	DeletedAt time.Time `pg:",soft_delete"`
}

type Action int

const (
	ActionRegister = iota + 1
)

type Token struct {
	ID          string
	Action      Action
	AccessLogID string
	AccessLog   *AccessLog
}
