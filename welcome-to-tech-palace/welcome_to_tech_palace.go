package techpalace

import (
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	panic("Please implement the WelcomeMessage() function")
}

//PrintStars prints n number of stars
func printStars(n int) string{
	stars:=""
	for i:=0;i<n;i++{
		stars +="*"
	}

	return stars
}
// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return printStars(numStarsPerLine)+"\n" + welcomeMsg +"\n"+ printStars(numStarsPerLine)

	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var newMsg string=""

	for index,character := range oldMsg{

		c:=string(character)
		if   (c == "\n") ||( c == "*"){
			continue
		}else if (c == " "){
			if index != 0{
				if string(oldMsg[index -1])!=" " && string(oldMsg[index +1])!=" " && string(oldMsg[index -1])!="\n" {
					newMsg+=" "
				}
			}
			continue

		}

		newMsg+=c
	}
	return newMsg
	panic("Please implement the CleanupMessage() function")
}
