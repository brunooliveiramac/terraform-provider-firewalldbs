package data_provider

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
