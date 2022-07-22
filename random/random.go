package random

import (
	"elasticsearch/constants"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewName() return a random full name first name will be 3 to 6 chars long
// last name will be 4 to 9 chars log
func NewName() string {
	var letters = []rune(constants.AlphabetsLower)
	lenFName := RandomNoNToN(3, 6)
	lenLName := RandomNoNToN(4, 9)

	firstName := make([]rune, lenFName)
	lastName := make([]rune, lenLName)

	for i := range firstName {
		firstName[i] = letters[rand.Intn(len(letters))]
	}

	for i := range lastName {
		lastName[i] = letters[rand.Intn(len(letters))]
	}
	return strings.Title(string(firstName)) + " " + strings.Title(string(lastName))
}

func NewEmail(name string) string {
	name = strings.ToLower(name)
	username := strings.Replace(name, " ", "_", 1)
	emailDomains := EmailDomains()

	domainIndex := rand.Intn(len(emailDomains))

	return username + emailDomains[domainIndex]
}

func NewTime() {

}

// EmailDomains returns a slice of email domains
func EmailDomains() []string {
	return []string{"@gmail.com", "@outlook.com", "@yahoo.com", "@hotmail.com",
		"@protonmail.com", "@zoho.com", "@icloud.com", "@mail.com"}
}

// RandomNoZeroToN returns a random no from 0 to n
func RandomNoZeroToN(n int) int {
	return rand.Intn(n)
}

// RandomNoZeroToN returns a random no from min to max
func RandomNoNToN(min, max int) int {
	return rand.Intn(max-min) + min
}
