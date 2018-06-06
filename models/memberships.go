package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
)

// Membership Struct
type Membership struct {
	gorm.Model  `json:"-"`
	Price       uint32 `gorm:"not null" json:"price"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"default:''" json:"description"`
	UUID        string `gorm:"not null;unique" json:"uuid"`
	id          int64  `gorm:"AUTO_INCREMENT:yes;primary_key:yes;column:membershipID" json:"id"`
	Users       []Users
}

// MembershipManager struct
type MembershipManager struct {
	db *DB
}

// User -
type Users struct {
	UserID string `gorm:"index"`
}

// NewMembershipManager - Create a new *UserManager that can be used for managing users.
func NewMembershipManager(db *DB) (*MembershipManager, error) {

	db.AutoMigrate(&User{})

	membershipmgr := MembershipManager{}

	membershipmgr.db = db

	return &membershipmgr, nil
}

// FindMembershipByUUID -
func (state *MembershipManager) FindMembershipByUUID(uuid string) *Membership {
	membership := Membership{}
	state.db.Where("uuid=?", uuid).Find(&membership)
	return &membership
}

// FindMembershipByID
func (state *MembershipManager) FindMembershipByID(id int64) *Membership {
	membership := Membership{}
	state.db.Where("id=?", id).Find(&membership)
	return &membership
}

// AddMembership - Creates a membership type
func (state *MembershipManager) AddMembership(name string, description string, price uint32) *Membership {
	guid, _ := uuid.NewV4()
	membership := &Membership{
		Name:        name,
		Description: description,
		UUID:        guid.String(),
		Price:       price,
	}
	state.db.Create(&membership)
	return membership
}

// LinkUserToMembership - Link a User to a Membership Type
func (state *MembershipManager) LinkUserToMembership(uuid string, UserUUID string) *Membership {

}
