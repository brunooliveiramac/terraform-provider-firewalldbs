package data_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateRole(t *testing.T) {
	//Given
	credentials := Credential{
		Host:     "bees-eastus2-pim-sandbox.postgres.database.azure.com",
		Database: "postgres",
		Port:     5432,
		Username: "",
		Password: "",
	}

	arr:= []string{"READ", "DELETE"}

	role := Role{
		Name:       "brunom",
		Privileges:  arr,
	}

	//When
	roleInserted, _ := CreateRole(&credentials, &role)

	//Then
	assert.Equal(t, "brunom", roleInserted.GetRoleName())
}


func TestShouldGrantPrivilegesToRole(t *testing.T) {
	//Given
	credentials := Credential{
		Host:     "bees-eastus2-pim-sandbox.postgres.database.azure.com",
		Database: "postgres",
		Port:     5432,
		Username: "",
		Password: "",
	}

	privileges:= []string{"SELECT", "DELETE"}

	role := Role{
		Name:       "brunom",
		Privileges:  privileges,
		Schema: "public",
	}

	//When
	roleInserted, _ := GrantPrivilegesOnAllTables(&credentials, &role)

	//Then
	assert.Equal(t, "brunom", roleInserted.GetRoleName())
}

func TestShouldReadRole(t *testing.T) {
	//Given
	credentials := Credential{
		Host:     "bees-eastus2-pim-sandbox.postgres.database.azure.com",
		Database: "pim",
		Port:     5432,
		Username: "",
		Password: "",
	}

	role := Role{
		Name:       "read_write",
	}

	//When
	selected, _ := SelectRole(&credentials, &role)

	//Then
	assert.Equal(t, "public", selected.GetRoleSchema())
}

func TestShouldRevokeRole(t *testing.T) {
	//Given
	credentials := Credential{
		Host:     "bees-eastus2-pim-sandbox.postgres.database.azure.com",
		Database: "postgres",
		Port:     5432,
		Username: "",
		Password: "",
	}

	arr:= []string{"READ", "DELETE"}

	role := Role{
		Name:       "developer",
		Privileges:  arr,
		Schema: "public",
	}

	//When
	_, err := RevokeAll(&credentials, &role)

	//Then
	assert.Equal(t, err, nil)
}

func TestShouldSelectUser(t *testing.T) {
	//Given
	credentials := Credential{
		Host:     "bees-eastus2-pim-sandbox.postgres.database.azure.com",
		Database: "pim",
		Port:     5432,
		Username: "",
		Password: "",
	}

	role := User{
		Username:       "brunom",
	}

	//When
	selected, _ := SelectUser(&credentials, &role)

	//Then
	assert.Equal(t, "brunom", selected.Username)
}

func TestShouldAlterUsername(t *testing.T) {
	//Given
	credentials := Credential{
		Host:     "bees-eastus2-pim-sandbox.postgres.database.azure.com",
		Database: "pim",
		Port:     5432,
		Username: "",
		Password: "",
	}

	role := User{
		Username:    "fitness",
		NewUsername: "batatinha",
	}

	//When
	_, err  := AlterUsername(&credentials, &role)

	//Then
	assert.Equal(t, nil, err)
}



