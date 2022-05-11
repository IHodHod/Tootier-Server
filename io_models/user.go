package io_models



type FindUserName struct {
	Username string `json:"username" uri:"username"`
	UserID uint64 `json:"user_id"`
}
