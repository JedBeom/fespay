package models

import "time"

type UserType int

const (
	UserStudent UserType = iota + 1
	UserTeacher
)

type Status int

const (
	StatusWorking = iota
	StatusFrozen  // Frozen(동결): 로그인, 조회 등은 가능. 결제, 수정 등의 행동 불가
	StatusBlocked // Blocked(차단): 모든 것이 불가. 유저 입장에서는 삭제된 것과 마찬가지.
)

type User struct {
	ID       string
	WalletID string
	BoothID  string
	LoginID  string
	Password string

	Type UserType

	Grade    int
	Class    int
	Number   int
	Name     string
	CardCode string

	Status    Status
	UpdatedAt time.Time
}

type Booth struct {
	ID          string
	WalletID    string
	Name        string
	Description string
	Staffs      []*User
	Status      Status
	UpdatedAt   time.Time
}

type OwnerType int

const (
	OwnerUser = iota + 1
	OwnerBooth
)

type Wallet struct {
	ID        string
	OwnerType OwnerType
	OwnerID   string
	Coin      int
	UpdatedAt time.Time
}

type Order struct {
	ID            string
	StaffID       string
	FromID        string
	ToID          string
	Amount        int
	RefundOrderID string

	AccessLogID string
	CreatedAt   time.Time
}

type AccessLog struct {
	ID        string
	SessionID string
	IP        string
	Action    string
	Path      string
	CreatedAt time.Time
}

type Session struct {
	ID        string
	UserID    string
	UserAgent string
	CreatedAt time.Time
	DeletedAt time.Time
}
