package main

import (
	"fmt"
	"github.com/lafelthitz/go-interface/api"
)

func main() {
	q := api.NewQuery(api.NewTD())
	r, err := q.Query("CHG000001")

	if err == nil {
		fmt.Println(r)
		return
	}

	fmt.Println(err.Error())

}
