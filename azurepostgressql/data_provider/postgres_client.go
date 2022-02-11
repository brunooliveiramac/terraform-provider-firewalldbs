package data_provider

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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
		log.Fatalf("Error when creating a new Role : %s", err)
		return r, err
	}

	defer connection.Close()

	return role, nil
}


func GrantPrivileges(credential *Credential, role *Role) (r *Role, err error) {

	connection := OpenTransaction(credential)

	sql := fmt.Sprintf("GRANT %s ON ALL TABLES IN SCHEMA PUBLIC TO  %s", strings.Join(role.GetRolePrivileges(), ","), role.Name)

	_, err = connection.Query(sql)

	if err != nil {
		log.Fatalf("Error when grating privileges to role : %s", err)
		return r, err
	}

	defer connection.Close()

	return role, nil
}



