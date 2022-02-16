package data_provider


type Role struct {
	Name 		 string
	Schema       string
	Tables       []string
	Privileges   []string
}

func (role *Role) SetName(name string) {
	role.Name = name
}

func (role *Role) SetSchema(schema string) {
	role.Schema = schema
}

func (role Role) GetRoleName() string {
	return role.Name
}

func (role Role) GetRolePrivileges() []string {
	return role.Privileges
}

func (role Role) GetRoleSchema() string {
	return role.Schema
}

func (role Role) GetRoleTables() []string {
	return role.Tables
}

func (role *Role) appendTable(table string) {
	role.Tables = append(role.Tables, table)
}

func (role *Role) appendPrivilege(privilege string) {
	role.Privileges = append(role.Privileges, privilege)
}

