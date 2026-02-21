package persistence

import "time"

type UserModel struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime"`
	DeletedAt time.Time `gorm:"not null;default:'0001-01-01 00:00:00'"`
}

func (UserModel) TableName() string {
	return "users"
}

type TaskModel struct {
	ID          int64     `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"type:text;not null;default:''"`
	Status      string    `gorm:"not null;default:'todo'"`
	UserID      int64     `gorm:"not null;index"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime"`
	DeletedAt   time.Time `gorm:"not null;default:'0001-01-01 00:00:00'"`
	User        UserModel `gorm:"foreignKey:UserID"`
}

func (TaskModel) TableName() string {
	return "tasks"
}
