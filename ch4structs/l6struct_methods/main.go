package main

import (
	"fmt"
)

type authenticationInfo struct {
	username string
	password string
}

// create the method below


func (value authenticationInfo) getBasicAuth () string {
	formattedValue := fmt.Sprintf("Authorization: Basic %s:%s", value.username, value.password )
	return formattedValue
}