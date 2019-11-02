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
	Coin     int    `sql:"default:0"`

	Status    Status `sql:"default:1"`
	UpdatedAt time.Time
}

type Booth struct {
	ID          string
	Name        string `sql:",unique" sql:",notnull" sql:"type:varchar(15)"`
	Description string `sql:"type:varchar(200)"`
	Location    string `json:"location"`
	Coin        int    `sql:"default:0"`
	Staffs      []*User
	Status      Status    `sql:"default:1" json:"status"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}

type RecordType int

const (
	RecordCharge RecordType = iota + 1
	RecordOrder
)

type Record struct {
	ID      string `json:"id"`
	StaffID string `sql:",notnull" json:"staffID"`
	Staff   *User  `json:"-"`
	BoothID string `sql:",notnull" json:"boothID"`
	Booth   *Booth `json:"-"`
	UserID  string `json:"userID,omitempty"`
	User    *User  `json:"-"`
	Amount  int    `sql:",notnull" json:"amount"`

	Type RecordType `json:"type"`

	AccessLogID string     `sql:",notnull" json:"-"`
	AccessLog   *AccessLog `json:"-"`

	CreatedAt  time.Time `sql:"default:now()" json:"createdAt"`
	PaidAt     time.Time `json:"paidAt,omitempty"`
	CanceledAt time.Time `json:"canceledAt,omitempty"`
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
