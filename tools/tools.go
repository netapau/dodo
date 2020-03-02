package tools

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// UserValidation confirm user condition with "yes" or "Yes" otherwise not valid.
func UserValidation() bool {
	ok := false
	r := bufio.NewReader(os.Stdin)
	//rp := regexp.MustCompile("^[Y|y]$|^[Y|y]es$") // US
	rp := regexp.MustCompile("^[O|o]$|^[O|o]ui$") // FR

	response, err := r.ReadString('\n')
	if err != nil {
		log.Panic(err)
	}
	response = strings.ReplaceAll(response, "\n", "")
	response = strings.ReplaceAll(response, "\r", "")

	match := rp.MatchString(response)
	if match == true {
		ok = true
	}

	return ok
}
