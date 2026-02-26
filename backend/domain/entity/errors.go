package entity

import "errors"

// ErrNotFound はエンティティが見つからない場合のセンチネルエラー
var ErrNotFound = errors.New("not found")
