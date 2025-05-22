package db

import (
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() {
	fmt.Println("Connected to DB")
}
