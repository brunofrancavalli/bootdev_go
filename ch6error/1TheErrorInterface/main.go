package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	iCust, error := sendSMS(msgToCustomer)

	if iCust == 0 {
		return 0, error
	}

	iSpouse, error := sendSMS(msgToSpouse)

	if iSpouse == 0 {
		return 0, error
	}

	return iCust + iSpouse, nil

}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
