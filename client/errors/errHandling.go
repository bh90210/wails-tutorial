package errors

import (
	"log"
)

// Handle errors
func Handle(err error) {
	if err != nil {
		log.Print(err)
	}
}
