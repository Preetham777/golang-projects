package main

import (
	"fmt"
	"net"
	netMail "net/mail"
	"net/smtp"
	"strings"

	// Additonal imports
	"github.com/jedib0t/go-pretty/v6/table"
)

/*
1. Get the input of mail address(s)
2. Loop through the email address
3. Check if its valid email address format
4. Check TXT records of the domain to identify authorised server
5. Check MX record lookup of the domain
6. If it has then its valid

*/

var (
	tableTitle                = "Email Verifier"
	colTitleIndex             = "#"
	colTitleValidationType    = "Validation Type"
	colTitleResult            = "Result"
	colTitleResultDescription = "Result Description"
	rowHeader                 = table.Row{colTitleIndex,
		colTitleValidationType,
		colTitleResult,
		colTitleResultDescription,
	}

	tableRows = []table.Row{}

	BULK_EMAIL_COUNT = 5
)

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
		tableRows = append(tableRows, table.Row{1, "Valid Email", "ERROR", err})
		return false
	} else {
		fmt.Printf("SUCCESS : Address is in valid format %v\n", isEmailValid)
		tableRows = append(tableRows, table.Row{1, "Valid Email", "SUCCESS", ""})
		return true
	}

}

func getDnsTxtRecords(ipMail *string) {

	domainName := strings.Split(*ipMail, "@")[1]

	if dnsTxtRecords, err := net.LookupTXT(domainName); err != nil {
		tableRows = append(tableRows, table.Row{2, "DNS TXT Records", "ERROR", err})
		fmt.Printf("ERROR: Seems there are no DNS txt records for email : %v with error %v\n", *ipMail, err)
	} else {
		fmt.Printf("SUCCESS:  here are the dns txt records %v\n", dnsTxtRecords)
		tableRows = append(tableRows, table.Row{2, "DNS TXT Records", "SUCCESS", dnsTxtRecords[1]})
	}

}

func getMxRecords(ipmail *string) {
	domainName := strings.Split(*ipmail, "@")[1]

	if mxRecords, err := net.LookupMX(domainName); err != nil {
		fmt.Printf("ERROR : MX records not found for email %v  with error %v\n", *ipmail, err)
		tableRows = append(tableRows, table.Row{3, "MX TXT Records", "ERROR", err})
	} else {
		fmt.Printf("INFO : MX recods found for email %v are %v\n", *ipmail, mxRecords)
		for i, mxHosts := range mxRecords {
			fmt.Printf("MX record%d Host: %v\n", i, mxHosts.Host)
		}
		tableRows = append(tableRows, table.Row{3, "MX TXT Records", "SUCCESS", mxRecords})
	}
}

func checkEmailDispoableOrNot(inputEmail *string) {
	inputEmailDomainName := strings.Split(*inputEmail, "@")[1]

	for _, mailDomain := range disposableEmailsList {
		if mailDomain == inputEmailDomainName {
			tableRows = append(tableRows, table.Row{4, "Non Disposable Email", "ERROR", "email is disposable"})
			return
		}
	}
	tableRows = append(tableRows, table.Row{4, "Non Disposable Email", "SUCCESS", "email is NOT disposable"})
}

func verifySMTPConnectivity(ipMail *string) {

	domainName := strings.Split(*ipMail, "@")[1]
	smtpPort := "25"

	emailHostName := strings.Split(*ipMail, "@")[0]
	newSmptConn, err := net.Dial("tcp", domainName+":"+smtpPort)

	if isError("dialing tcp to domain", err) {
		tableRows = append(tableRows, table.Row{5, "SMTP Verification", "ERROR", err})
		return
	}

	newSmtpClient, err := smtp.NewClient(newSmptConn, emailHostName)

	if isError("creating new SMTP client", err) {
		tableRows = append(tableRows, table.Row{5, "SMTP Verification", "ERROR", err})
		return
	}

	isSmtpValid := newSmtpClient.Verify(domainName)

	if isSmtpValid == nil {
		//valid
		tableRows = append(tableRows, table.Row{5, "SMTP Verification", "SUCCESS", "SMTP Verification done"})
		return
	} else {
		tableRows = append(tableRows, table.Row{5, "SMTP Verification", "ERROR", isSmtpValid})
		return
	}

}

func isError(errorAt string, e error) bool {
	if e != nil {
		fmt.Printf("ERROR: at %v with the error %v\n", errorAt, e)
		return true
	}
	return false
}

func renderTable(ipMail string) {
	tw := table.NewWriter()
	tw.SetTitle(tableTitle + " : " + ipMail)
	tw.AppendHeader(rowHeader)
	tw.AppendRows(tableRows)
	tw.AppendSeparator()
	tableRows = []table.Row{}
	fmt.Println(tw.Render())
}

func main() {

	inputEmail := "sample_mail@gmail.com"
	maxBulkEmailCountReached := 0

	getUpdatedDisposableEmails()

	for len(inputEmail) != 0 {
		inputEmail = getInputEmails()

		if !checkEmailValidity(&inputEmail) {
			continue
		}

		getDnsTxtRecords(&inputEmail)
		getMxRecords(&inputEmail)
		checkEmailDispoableOrNot(&inputEmail)
		verifySMTPConnectivity(&inputEmail)

		renderTable(inputEmail)

		maxBulkEmailCountReached += 1

		if maxBulkEmailCountReached == BULK_EMAIL_COUNT {
			break
		}

	}

}
