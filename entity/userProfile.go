package entity

type UserProfile struct {
	ID          int    `json:"id"`
	ProfileUser string `json:"profile_user"`
	UserID      int    `json:"user_id"`
}
