package main

import (
	"fmt"

	"github.com/dreygur/bkashgo/models"
)

func main() {
	a := models.BkashIPNPayload{
		Type: "ABC",
		// Message: "ABCD",
		Message: models.Message{
			DateTime: "ABCDR",
		},
	}
	fmt.Println(a.Message)
}
