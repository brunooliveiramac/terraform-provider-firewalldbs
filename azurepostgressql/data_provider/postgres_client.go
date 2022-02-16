package data_provider

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sort"
	"strings"
)

func DBClientConnect(credential *Credential) (err error) {

	connection := OpenTransaction(credential)

	if err != nil {
		log.Printf("Error! %s", err)
		return  err
	}

	//TODO
	// Open Firewall port using the database api or azure cli

	err = connection.Ping()

	if err != nil {
		log.Printf("Error! %s", err)
		return  err
	}

	defer connection.Close()

	log.Printf("Connected!")

	return  nil
}


func OpenTransaction(credential *Credential) (db *sql.DB) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		credential.GetHost(),
		credential.GetPort(),
		credential.GetUsername(),
		credential.GetPassword(),
		credential.GetDatabase())

	connection, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error when openning connection ! %s", err)
	}

	log.Printf("Connected!")

	return connection
}


func CreateRole(credential *Credential, role *Role) (r *Role, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("CREATE ROLE %s", role.Name)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when creating a new Role : %s", err)
		return nil, err
	}

	defer connection.Close()

	return role, nil
}


func GrantPrivilegesOnAllTables(credential *Credential, role *Role) (r *Role, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("GRANT %s ON ALL TABLES IN SCHEMA %s TO  %s", strings.Join(role.GetRolePrivileges(), ","), role.GetRoleSchema(), role.Name)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when grating privileges to role : %s", err)
		return nil, err
	}

	defer connection.Close()

	return role, nil
}


func GrantPrivileges(credential *Credential, role *Role) (r *Role, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("GRANT %s ON %s IN SCHEMA %s TO  %s", strings.Join(role.GetRolePrivileges(), ","), strings.Join(role.GetRoleTables(), ","), role.GetRoleSchema(), role.Name)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when grating privileges to role : %s", err)
		return nil, err
	}

	defer connection.Close()

	return role, nil
}


func SelectRole(credential *Credential, role *Role) (r *Role, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("SELECT grantor, grantee, table_catalog, table_schema, table_name, privilege_type  FROM information_schema.role_table_grants where grantee='%s'", role.GetRoleName())

	data, err := connection.Query(sql)

	if err != nil {
		log.Printf("Error when selecting role: %s", err)
		return nil, err
	}

	var roleModels []RoleModelRow

	for data.Next() {
		var roleModelRow RoleModelRow

		err = data.Scan(&roleModelRow.Grantor, &roleModelRow.Name, &roleModelRow.Database, &roleModelRow.Schema,
						&roleModelRow.Table, &roleModelRow.Privilege)

		if err != nil {
			log.Printf("Error when selecting role: %s -> %s", role.Name, err)
			return nil, err
		}

		roleModels = append(roleModels, roleModelRow)

	}

	roleToReturn := Role{}

	for _, roleModel := range roleModels {
		if roleModel.Name != roleToReturn.Name {
			roleToReturn.SetName(roleModel.Name)
			roleToReturn.SetSchema(roleModel.Schema)
		}
		if !contains(roleToReturn.GetRoleTables(), roleModel.Table) {
			roleToReturn.appendTable(roleModel.Table)
		}
		if !contains(roleToReturn.GetRolePrivileges(), roleModel.Privilege) {
			roleToReturn.appendPrivilege(roleModel.Privilege)
		}
	}

	defer connection.Close()

	sort.Strings(roleToReturn.Tables)
	sort.Strings(roleToReturn.Privileges)

	log.Printf("Role! %s", roleToReturn)

	return &roleToReturn, nil
}


func SelectUser(credential *Credential, user *User) (r *User, err error) {

	connection := OpenTransaction(credential)

	userToReturn := User{}

	sql := fmt.Sprintf("SELECT rolname FROM pg_roles where rolname = $1")

	statement, err := connection.Prepare(sql)


	if err != nil {
		return nil, err
	}
	err = statement.QueryRow(user.Username).Scan(&userToReturn.Username)

	if err != nil {
		log.Printf("Error when selecting user: %s", err)
		return nil, err
	}


	return &userToReturn, nil
}

func RevokeAll(credential *Credential, role *Role) (r *Role, err error) {
	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("REVOKE ALL ON ALL TABLES IN SCHEMA %s FROM %s", role.GetRoleSchema(), role.Name)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when revoking privileges from role : %s", err)
		return nil, err
	}

	defer connection.Close()

	return role, nil
}

func DropRole(credential *Credential, role *Role) (r *Role, err error) {
	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("DROP ROLE %s", role.Name)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when droping role : %s", err)
		return nil, err
	}

	defer connection.Close()

	return role, nil
}

func CreateUser(credential *Credential, user *User) (u *User, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("CREATE ROLE %s WITH LOGIN ENCRYPTED PASSWORD '%s' IN ROLE %s;", user.Username, user.Password, user.RoleName())

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when creating a new User : %s", err)
		return nil, err
	}

	defer connection.Close()

	return user, nil
}

func AlterUsername(credential *Credential, user *User) (r *User, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("ALTER ROLE %s RENAME TO %s", user.Username, user.NewUsername)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when changing username : %s", err)
		return nil, err
	}

	defer connection.Close()

	return user, nil
}


func AlterPassword(credential *Credential, user *User) (r *User, err error) {
	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("ALTER ROLE %s WITH ENCRYPTED PASSWORD '%s';", user.Username, user.NewPassword)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when changing user password : %s", err)
		return nil, err
	}

	defer connection.Close()

	return user, nil
}

func DropUser(credential *Credential, user *User) (r *User, err error) {
	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("DROP ROLE %s", user.Username)

	_, err = connection.Query(sql)

	if err != nil {
		log.Printf("Error when droping user : %s", err)
		return nil, err
	}

	defer connection.Close()

	return user, nil
}


func contains(list []string, element string) bool {
	for _, item := range list {
		if item == element {
			return true
		}
	}
	return false
}

