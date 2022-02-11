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

	arr:= []string{"SELECT", "DELETE"}

	role := Role{
		Name:       "brunom",
		Privileges:  arr,
	}

	//When
	roleInserted, _ := GrantPrivileges(&credentials, &role)

	//Then
	assert.Equal(t, "brunom", roleInserted.GetRoleName())
}
