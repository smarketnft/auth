package models

import "gorm.io/gorm"

type UserState string

var (
	UserStatePending = UserState("pending")
	UserStateActive  = UserState("active")
	UserStateBanned  = UserState("banned")
)

type UserRole string

var (
	UserRoleMember = UserRole("member")
	UserRoleAdmin  = UserRole("admin")
)

type User struct {
	gorm.Model
	UID      string    `gorm:"type:character varying(20);not null"`
	UserName string    `gorm:"type:character varying(50)"`
	Email    string    `gorm:"type:character varying(50);not null"`
	Password string    `gorm:"type:character varying(100);not null"`
	Address  string    `gorm:"type:character varying"`
	State    UserState `gorm:"type:character varying(10);not null"`
	Role     UserRole  `gorm:"type:character varying(10);not null"`
}
