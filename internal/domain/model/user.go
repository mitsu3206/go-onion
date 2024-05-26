package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID          uint         // 主キーの標準フィールド
	Name        string       // 通常の文字列フィールド
	Email       *string      // 文字列へのポインタ、nullを許可
	Age         uint8        // 符号なし8ビット整数
	Birthday    *time.Time   // time.Timeへのポインタ。nullを許可
	ActivatedAt sql.NullTime // sql.NullTimeを使用したnull可能な時間フィールド
	CreatedAt   time.Time    // GORMによって自動的に管理される作成時間
	UpdatedAt   time.Time    // GORMによって自動的に管理される更新時間
}

type UserRepository interface {
	CreateUser(name, email string, age uint8, birthday time.Time) (User, error)
	GetUserById(id uint) (User, error)
	UpdateUser(User) (User, error)
	DeleteUser(User) error
}
