package entity

type UserTokenClaim struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}
