package handle

import (
	"log"
)

// Handle errors
func Handle(err error) {
	if err != nil {
		log.Println(err)
	}
}
