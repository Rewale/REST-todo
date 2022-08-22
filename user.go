package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binging:"required"`
	Password string `json:"password" binging:"required"`
}
