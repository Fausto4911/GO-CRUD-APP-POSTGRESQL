package main

import (
	"context"
	"fmt"
	"go-crud-app-postgresql/driver"
	"go-crud-app-postgresql/model"
	"go-crud-app-postgresql/repository/client"
	"log"
	"time"
)

var ctx context.Context

func main() {
	fmt.Println("*** Running main ***")
	ctx = context.Background()
	repo := client.IClientRepository{
		Ctx: ctx,
	}

	db := driver.GetPostgreSQLConnection()
	fmt.Println("======== Store ==========")
	c := model.Client{
		Id:       1,
		Email:    "example@example.com",
		Password: "examplepassword",
		NickName: "NickNameExample",
		CreateAt: time.Now(),
	}

	err := repo.Store(c, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	db = driver.GetPostgreSQLConnection()
	fmt.Println("======== GetById =========")
	u, err := repo.GetById(1, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

	db = driver.GetPostgreSQLConnection()
	fmt.Println("======== GetAll =========")
	arr, err := repo.GetAll(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range arr {
		fmt.Println(c)
	}

	db = driver.GetPostgreSQLConnection()
	fmt.Println("======== Update =========")
	up := model.Client{
		Id:       1,
		Email:    "update@example.com",
		Password: "updatepassword",
		NickName: "NickNameUpdate",
		CreateAt: time.Now(),
	}
	err = repo.Update(up, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)

	db = driver.GetPostgreSQLConnection()
	fmt.Println("======== Delete =========")
	err = repo.Delete(1, db)
	if err != nil {
		log.Fatal(err)
	}

}
