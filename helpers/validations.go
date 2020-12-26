package helpers

import (
	"regexp"
)

func ValidateEmail(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return re.MatchString(email)

}

func ValidatePassword(password string) map[string]interface{} {

	r, _ := regexp.Compile(`[A-Z]`)
	if !r.MatchString(password) {
		return map[string]interface{}{
			"status": false, "message": "Password should contain a uppercase letter."}
	}

	r, _ = regexp.Compile(`[a-z]`)
	if !r.MatchString(password) {
		return map[string]interface{}{
			"status": false, "message": "Password should contain a lowercase letter."}

	}

	r, _ = regexp.Compile(`[0-9]`)
	if !r.MatchString(password) {
		return map[string]interface{}{
			"status": false, "message": "Password should contain a digit."}
	}
	if len(password) < 8 {
		return map[string]interface{}{
			"status": false, "message": "Password should be atleast 8 characters."}
	}

	return map[string]interface{}{
		"status": true}
}
