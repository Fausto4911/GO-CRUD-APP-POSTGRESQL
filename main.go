package main

import (
	"context"
	"fmt"
	"go-crud-app-postgresql/driver"
	"go-crud-app-postgresql/repository/client"
)

var ctx context.Context

func main () {
	fmt.Println("*** Running main ***")
	ctx = context.Background()
	repo := client.IClientRepository{
		Ctx: ctx,
	}
	fmt.Println("======== GetById ==========")
	u, err := repo.GetById(1, driver.GetPostgreSQLConnection())
	if err != nil {
		fmt.Println(err)
	}
    fmt.Println(u)

}
