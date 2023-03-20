package main

import (
	"fmt"
	"net"
	netMail "net/mail"
	"strings"
)

/*
1. Get the input of mail address(s)
2. Loop through the email address
3. Check if its valid email address format
4. Check TXT records of the domain to identify authorised server
4. If valid check for the DMX record lookup
5. If it has then its valid

*/

func getInputEmails() string {
	fmt.Printf("Enter email address\n")
	var inputEmail string

	fmt.Scanln(&inputEmail)
	fmt.Printf("INFO : Here is the email address entered : %v\n", inputEmail)

	return inputEmail
}

func checkEmailValidity(ipMail *string) bool {
	isEmailValid, err := netMail.ParseAddress(*ipMail)

	if err != nil {
		fmt.Printf("ERROR : Invalid email :  %v\n", *ipMail)
		return false
	} else {
		fmt.Printf("SUCCESS : Address is in valid format %v\n", isEmailValid)
		return true
	}

}

func getDnsTxtRecords(ipMail *string) {

	domainName := strings.Split(*ipMail, "@")[1]

	if dnsTxtRecords, err := net.LookupTXT(domainName); err != nil {
		fmt.Printf("ERROR: Seems there are no DNS txt records for email : %v with error %v\n", *ipMail, err)
	} else {
		fmt.Printf("SUCCESS:  here are the dns txt records %v\n", dnsTxtRecords)
	}

}

func main() {

	inputEmail := getInputEmails()

	checkEmailValidity(&inputEmail)
	getDnsTxtRecords(&inputEmail)

}
