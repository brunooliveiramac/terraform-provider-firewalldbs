package data_provider

type User struct {
	Username    string
	Password    string
	Database    string
	Role        Role
	NewUsername string
	NewPassword string
}

func (user *User) RoleName() string {
	return user.Role.GetRoleName()
}
