package entity

type User struct {
	ID       string `db:"id" json:"id,omitempty"`
	Username string `db:"username" json:"username,omitempty"`
	Fullname string `db:"full_name" json:"full_name,omitempty"`
	Email    string `db:"email" json:"email,omitempty"`
	Avatar   string `db:"avatar" json:"avatar,omitempty"`
	Role     string `db:"role" json:"role,omitempty"`
	Password string `db:"password" json:"password,omitempty"`
}
