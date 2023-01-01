package entity

type User struct {
	Username  string     `gorm:"primaryKey;column:username;type:varchar(100)"`
	Password  string     `gorm:"column:password;type:varchar(200)"`
	IsActive  bool       `gorm:"column:is_active;type:boolean"`
	UserRoles []UserRole `gorm:"ForeignKey:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "tb_user"
}
