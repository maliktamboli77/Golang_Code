package main

import (
	"fmt"
	"regexp"
)

func main() {
	//Sample inputs to validate
	name := "Malik Tamboli"
	email := "malikrtamboli@gmail.com"
	phoneNo := "+917709898503"

	//Regex patterns
	namePattern := `^[a-zA-Z]+(?: [a-zA-Z]+)?$`
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^(\+91[\-\s]?)?[0]?(91)?[789]\d{9}$`

	//Compile regex and match patterns
	matchName, _ := regexp.MatchString(namePattern, name)
	matchEmail, _ := regexp.MatchString(emailPattern, email)
	matchPhone, _ := regexp.MatchString(phonePattern, phoneNo)

	//Output Results
	fmt.Printf("Name valid: %t\n", matchName)
	fmt.Printf("Email valid: %t\n", matchEmail)
	fmt.Printf("Phone Number valid: %t\n", matchPhone)
}
