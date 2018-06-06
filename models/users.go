package models

import (
	"github.com/jinzhu/gorm"
	// postgress db driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// import sqlite3 driver
	"database/sql"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	gorm.Model       `json:"-"`
	Username         string  `gorm:"not null;unique" json:"username"`
	Password         string  `gorm:"not null" json:"-"`
	UUID             string  `gorm:"not null;unique" json:"uuid"`
	BillingAddress   Address // One-To-One relationship (belongs to - use BillingAddressID as foreign key)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One relationship (belongs to - use ShippingAddressID as foreign key)
	ShippingAddressID int
	Emails            []Email // One-To-Many relationship (has many - use Email's UserID as foreign key)

}

type Email struct {
	ID         int
	UserID     string `gorm:"index"`                          // Foreign key (belongs to), tag `index` will create index for this column
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` set sql type, `unique_index` will create unique index for this column
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // Set field as not nullable and unique
	Address2 string         `gorm:"type:varchar(100);unique"`
	ZipCode  sql.NullString `gorm:"not null"`
	State    string         `gorm:"type:varchar(100)"`
}

// UserManager struct
type UserManager struct {
	db *DB
}

// NewUserManager - Create a new *UserManager that can be used for managing users.
func NewUserManager(db *DB) (*UserManager, error) {

	db.AutoMigrate(&User{})

	usermgr := UserManager{}

	usermgr.db = db

	return &usermgr, nil
}

// HasUser - Check if the given username exists.
func (state *UserManager) HasUser(username string) bool {
	if err := state.db.Where("username=?", username).Find(&User{}).Error; err != nil {
		return false
	}
	return true
}

// FindUser -
func (state *UserManager) FindUser(username string) *User {
	user := User{}
	state.db.Where("username=?", username).Find(&user)
	return &user
}

// FindUserByUUID -
func (state *UserManager) FindUserByUUID(uuid string) *User {
	user := User{}
	state.db.Where("uuid=?", uuid).Find(&user)
	return &user
}

// AddUser - Creates a user and hashes the password
func (state *UserManager) AddUser(username, password string) *User {
	passwordHash := state.HashPassword(username, password)
	guid, _ := uuid.NewV4()
	user := &User{
		Username: username,
		Password: passwordHash,
		UUID:     guid.String(),
	}
	state.db.Create(&user)
	return user
}

// HashPassword - Hash the password (takes a username as well, it can be used for salting).
func (state *UserManager) HashPassword(username, password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic("Permissions: bcrypt password hashing unsuccessful")
	}
	return string(hash)
}

// CheckPassword - compare a hashed password with a possible plaintext equivalent
func (state *UserManager) CheckPassword(hashedPassword, password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) != nil {
		return false
	}
	return true
}

// ChangePassword - update the users password
func (state *UserManager) ChangePassword(username string, password string) {
	passwordHash := state.HashPassword(username, password)
	user := state.FindUser(username)
	state.db.Model(user).Update("password", passwordHash)
	return
}
