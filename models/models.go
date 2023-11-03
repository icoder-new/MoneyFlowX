package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	ID        string         `json:"uuid" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	UserID    string         `json:"user_id" gorm:"references:users(id);not null"`
	User      User           `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Number    string         `json:"number" gorm:"not null"`
	Balance   float64        `json:"balance" gorm:"not null"`
	UserType  string         `json:"user_type,omitempty" gorm:"references:users(type);not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type User struct {
	ID        string         `json:"uuid,omitempty" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	Firstname string         `json:"firstname" gorm:"not null"`
	Lastname  string         `json:"lastname" gorm:"not null"`
	Username  string         `json:"username,omitempty" gorm:"not null;unique"`
	Email     string         `json:"email,omitempty" gorm:"not null;unique"`
	Password  string         `json:"-" gorm:"not null"`
	Type      string         `json:"type" gorm:"not null"`
	IsActive  bool           `json:"is_active" gorm:"not null"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Transaction struct {
	ID        string         `json:"uuid" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	UserID    string         `json:"user_id" gorm:"references:users(id);not null"`
	User      User           `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	WalletID  string         `json:"wallet_id" gorm:"references:wallets(id);not null"`
	Wallet    Wallet         `json:"-" gorm:"foreignKey:WalletID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Amount    float64        `json:"amount,omitempty" gorm:"not null"`
	Comment   string         `json:"comment" gorm:"not null"`
	Type      string         `json:"type,omitempty" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type PasswordReset struct {
	ID        string    `json:"uuid" gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	UserID    string    `json:"user_id" gorm:"references:users(id);not null"`
	User      User      `json:"-" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email     string    `json:"email" gorm:"references:users(email);not null"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
