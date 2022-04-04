package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var conn pgx.Conn

func dbConnection() {
	connection, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	conn = *connection
	//defer conn.Close(context.Background())

	createSql := ` 
	create table if not exists seaports(
		id int primary key,
		name VARCHAR,
		code VARCHAR,
		city VARCHAR,
		state VARCHAR,
		country VARCHAR
	);
	`

	_, error := conn.Exec(context.Background(), createSql)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Table creation: %v\n", error)
		os.Exit(1)
	}
	//fmt.Println(success)
}

func Createnewport(id int64, name, code, city, state, country string) {

	_, newerr := conn.Exec(context.Background(), "insert into seaports(id,name,code,city,state,country) values ($1,$2,$3,$4,$5,$6)", id, name, code, city, state, country)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr)
		os.Exit(1)
	}
	fmt.Println("Create")
}

func Getportdetails(id int64) (Id int64, name, code, city, state, country string) {
	fmt.Println("Fetch")
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	AllRows, errorRet := conn.Query(context.Background(), "select * from seaports where id=$1", id)
	if errorRet != nil {
		fmt.Fprintf(os.Stderr, "Retrival Failed : %v\n", errorRet)
		os.Exit(1)
	}
	defer AllRows.Close()

	var ID int64
	var Name string
	var Code string
	var City string
	var State string
	var Country string

	for AllRows.Next() {
		error := AllRows.Scan(&ID, &Name, &Code, &City, &State, &Country)
		if error != nil {
			fmt.Fprintf(os.Stderr, "scanning failed:%v", error)
		}
	}
	log.Println(name, code)
	return ID, Name, Code, City, State, Country
}

func UpdatePortDetails(id int64, name, code, city, state, country string) {
	fmt.Println("Update")

	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	_, newerr := conn.Exec(context.Background(), "update seaports set name = $2,code=$3,city=$4,state=$5,country=$6 where id=$1", id, name, code, city, state, country)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "Updation failed: %v\n", newerr)
		os.Exit(1)
	}
	//fmt.Println("Updated")
	//fmt.Println(id, name)

}

func checkPortId(id int64) bool {
	if id == 0 {
		return false
	}
	return true
}

func DeletePortDetails(id int64) {
	fmt.Println("Delete")
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	_, newerr := conn.Exec(context.Background(), "Delete from seaports where id=$1", id)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "Deletion failed: %v\n", newerr)
		os.Exit(1)
	}

}
