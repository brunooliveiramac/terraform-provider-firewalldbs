package data_provider

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"

)

func DBClient(credential *Credential) (db *sql.DB, err error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		credential.GetHost(),
		credential.GetPort(),
		credential.GetUsername(),
		credential.GetPassword(),
		credential.GetDatabase())

	connection, err := sql.Open("postgres", psqlInfo)

	log.Printf(psqlInfo)

	if err != nil {
		log.Printf("Error!", err)

		return nil, err
	}

	log.Printf("Connected!")

	return connection, nil
}

type Credential struct {
	Host 		 string
	Database     string
	Port         int
	Username     string
	Password     string
}

func (credential Credential) GetPassword() string {
	return credential.Password
}
func (credential Credential) GetHost() string {
	return credential.Host
}
func (credential Credential) GetDatabase() string {
	return credential.Database
}
func (credential Credential) GetPort() int {
	return credential.Port
}
func (credential Credential) GetUsername() string {
	return credential.Username
}

