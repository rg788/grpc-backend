package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
)

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/ports"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	/*extensions := `
	CREATE EXTENSION "pgcrypto";`

	_, error2 := conn.Exec(context.Background(), extensions)
	if error2 != nil {
		fmt.Fprintf(os.Stderr, "Extension creation: %v\n", error2)
		os.Exit(1)
	}*/

	createSql := ` 
	create table if not exists seaports(
		id UUID DEFAULT gen_random_uuid(),
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

	_, newerr := conn.Exec(context.Background(), "insert into seaports(id,name,code,city,state,country) values ($1,$2,$3,$4,$5,$6)", uuid.NewV4().String(), "nikil", "vdy789", "erode", "tamilnadu", "india")
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr)
		os.Exit(1)
	}
	_, newerr1 := conn.Exec(context.Background(), "insert into seaports(id,name,code,city,state,country) values ($1,$2,$3,$4,$5,$6)", uuid.NewV4().String(), "Rethenya", "vdu789", "tup", "tamilnadu", "india")
	if newerr1 != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr1)
		os.Exit(1)
	}
	results, newerr2 := conn.Exec(context.Background(), "select * from seaports")
	if newerr2 != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr2)
		os.Exit(1)
	}
	fmt.Println(results)

	AllRows, errorRet := conn.Query(context.Background(), "select * from seaports")
	if errorRet != nil {
		fmt.Fprintf(os.Stderr, "Retrival Failed : %v\n", errorRet)
		os.Exit(1)
	}

	defer AllRows.Close()

	for AllRows.Next() {
		values, err := AllRows.Values()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Retrival Failed : %v\n", err)
			os.Exit(1)
		}
		fmt.Println(values[0], values[1], values[2], values[3], values[4], values[5])
	}

	/*type Place struct {
	    Country string
	    City    sql.NullString
	    TelCode int
	}
	 places := []Place{}
	 err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	 if err != nil {
	     fmt.Println(err)
	      return
	 }*/

}
