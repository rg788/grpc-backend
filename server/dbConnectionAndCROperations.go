package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var conn pgx.Conn

func dbConnection() {
	connection, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/nikil"))
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

func Createnewport(id int64, name, code, city, state, country string) error {

	_, newerr := conn.Exec(context.Background(), "insert into seaports(id,name,code,city,state,country) values ($1,$2,$3,$4,$5,$6)", id, name, code, city, state, country)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr)
		//os.Exit(1)
		return newerr
	}
	fmt.Println("Create")
	return newerr
}

func Getportdetails(id int64) (Id int64, name, code, city, state, country string) {
	fmt.Println("Fetch")

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
}

func CheckPortId(id int64) bool {

	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())
	data, newerr := conn.Query(context.Background(), "select id from seaports where id=$1", id)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "Updation failed: %v\n", newerr)
		os.Exit(1)
	}
	var ID int64
	for data.Next() {
		error := data.Scan(&ID)
		if error != nil {
			fmt.Fprintf(os.Stderr, "scanning failed:%v", error)
		}
	}
	if ID == 0 {
		return false
	}
	return true
}

func DeletePortDetails(id int64) error {
	fmt.Println("Delete")
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	_, newerr := conn.Exec(context.Background(), "Delete from seaports where id=$1", id)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "Deletion failed: %v\n", newerr)
		os.Exit(1)
		return newerr
	}

	return nil

}

type portAttributes struct {
	ID      int64
	Name    string
	Code    string
	City    string
	State   string
	Country string
}

func GetAllPorts(page, limit int32) []portAttributes {
	pageNumber := (page - 1) * limit
	query :=
		`Select * from seaports
         offset $1
         fetch first $2 row only `
	Ports, err := conn.Query(context.Background(), query, pageNumber, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Retrieving all the ports details from the database failed : %v", err)
	}
	defer Ports.Close()
	var port portAttributes
	portsInRanges := []portAttributes{}
	for Ports.Next() {
		error := Ports.Scan(&port.ID, &port.Name, &port.Code, &port.City, &port.State, &port.Country)
		if error != nil {
			fmt.Fprintf(os.Stderr, "Scanning failed when trying to retrive all the ports from the database:%v", error)
		}
		portsInRanges = append(portsInRanges, port)
	}

	return portsInRanges
}
