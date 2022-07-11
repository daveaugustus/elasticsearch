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

// NewName() return a random full name
func NewName() string {
	var letters = []rune(constants.AlphabetsLower)

	firstName := make([]rune, 5)
	lastName := make([]rune, 5)
	for i := range firstName {
		firstName[i] = letters[rand.Intn(len(letters))]
		lastName[i] = letters[rand.Intn(len(letters))]
	}
	return strings.Title(string(firstName)) + " " + strings.Title(string(lastName))
}

func NewAge() int {
	maxAge := 90
	minAge := 13
	age := rand.Intn(maxAge-minAge) + minAge
	return age
}

func NewEmail(name string) string {
	name = strings.ToLower(name)
	username := strings.Replace(name, " ", "_", 1)
	emailDomains := EmailProviders()

	domainIndex := rand.Intn(len(emailDomains))

	return username + emailDomains[domainIndex]
}

func NewTime() {

}

func EmailProviders() []string {
	return []string{"@gmail.com", "@outlook.com", "@yahoo.com", "@hotmail.com",
		"@rediffmail.com", "@protonmail.com", "@zoho.com", "@icloud.com", "@mail.com"}
}
