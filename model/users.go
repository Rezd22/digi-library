package model

type User struct {
	ID        int    `gorm:"type:int;primaryKey"`
	Username  string `gorm:"type:varchar(20);not null;unique" json:"username" form:"username"`
	Password  string `gorm:"type:varchar(20);not null" json:"password" form:"password"`
	User_type string `gorm:"type:varchar(20);not null" json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
}

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
	User_type  string `json:"user_type"`
}
