package repository

import "errors"

// ErrNotificationNotFound はプッシュ通知が見つからなかったことを表すエラー
var ErrBearNotFound = errors.New("repository: bear not found")
