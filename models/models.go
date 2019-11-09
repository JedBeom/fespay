package models

import "time"

type UserType int

const (
	UserStudent UserType = iota + 1
	UserTeacher
	UserParent
	UserGuest
)

type Status int

const (
	StatusWorking   = iota + 1
	StatusFrozen    // Frozen(동결): 로그인, 조회 등은 가능. 결제, 수정 등의 행동 불가
	StatusSuspended // Suspended(정지): 모든 것이 불가. 유저 입장에서는 삭제된 것과 마찬가지.
)

type User struct {
	ID string `json:"id"`

	BoothID  string `json:"boothID"`
	Booth    *Booth `json:"booth,omitempty"`
	LoginID  string `json:"loginID,omitempty"`
	Password string `json:"-"`

	Type UserType `sql:",notnull" json:"userType"`

	Grade    int    `sql:",unique:gcn" json:"grade,omitempty"`
	Class    int    `sql:",unique:gcn" json:"class,omitempty"`
	Number   int    `sql:",unique:gcn" json:"number,omitempty"`
	Name     string `sql:",unique" sql:",notnull" sql:"type:varchar(7)" json:"name"`
	CardCode string `sql:",unique" sql:",notnull" sql:"type:char(5)" json:"cardCode"`
	// PayCode  string `sql:",unique" sql:",notnull" json:"payCode"`
	Coin int `sql:"default:0" json:"coin"`

	Status    Status     `sql:"default:1" json:"status"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type Booth struct {
	ID          string     `json:"id"`
	Name        string     `sql:",unique" sql:",notnull" sql:"type:varchar(10)" json:"name"`
	Description string     `sql:"type:varchar(200)" json:"description"`
	Location    string     `json:"location" json:"location"`
	Coin        int        `sql:"default:0" json:"coin"`
	Staffs      []*User    `json:"staffs,omitempty"`
	Status      Status     `sql:"default:1" json:"status"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type RecordType int

const (
	RecordCharge RecordType = iota + 1
	RecordOrder
)

type Record struct {
	ID      string `json:"id"`
	StaffID string `sql:",notnull" json:"staffID"`
	Staff   *User  `json:"staff,omitempty"`
	BoothID string `sql:",notnull" json:"boothID"`
	Booth   *Booth `json:"booth,omitempty"`
	UserID  string `json:"userID,omitempty"`
	User    *User  `json:"user,omitempty"`
	Amount  int    `sql:",notnull" json:"amount"`

	Type RecordType `json:"type"`

	AccessLogID string     `sql:",notnull" json:"-"`
	AccessLog   *AccessLog `json:"-"`

	CreatedAt  *time.Time `sql:"default:now()" json:"createdAt"`
	PaidAt     *time.Time `json:"paidAt,omitempty"`
	CanceledAt *time.Time `json:"canceledAt,omitempty"`
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

type AdminLog struct {
	ID          string
	AccessLog   *AccessLog
	AccessLogID string
}
