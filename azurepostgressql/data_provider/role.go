package data_provider


type Role struct {
	Name 		 string
	Privileges   []string
}

func (role Role) GetRoleName() string {
	return role.Name
}

func (role Role) GetRolePrivileges() []string {
	return role.Privileges
}

