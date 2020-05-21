package main

import (
	"context"
	"fmt"
)

var ctx context.Context

func main () {
	fmt.Println("*** Running main ***")
	ctx = context.Background()
	repo := IUsersRepository{
		Ctx: ctx,
	}
	fmt.Println("======== GetById ==========")
	u, err := repo.GetById(1)
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(u)

}
