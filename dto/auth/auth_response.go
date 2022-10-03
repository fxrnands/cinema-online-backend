package authdto

type LoginResponse struct {
	ID        int    `json:"id"`
	FullName  string `gorm:"type: varchar(255)" json:"fullName"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Role      string `gorm:"type: varchar(100)" json:"role"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	FullName  string `gorm:"type: varchar(255)" json:"fullName"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Role      string `gorm:"type: varchar(100)" json:"role"`
}

type RegisterResponse struct {
	ID       int    `json:"id"`
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Role     string `gorm:"type: varchar(100)" json:"role"`
}