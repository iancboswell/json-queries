package main

import (
	"fmt"
	"log"

	// Make sure the MySQL driver has called `init()`
	_ "github.com/go-sql-driver/mysql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&multiStatements=true&readTimeout=1s&clientFoundRows=true",
		"root",
		"root",
		"(127.0.0.1:3306)",
		"json"))
	if err != nil {
		log.Fatalf("SQLX failed to open a connection to MySQL\n"+
			"sqlx.open() failed with: %v", err)
	}

	RawQueries()
	SquirrelQueries()
}

func PrintResults(results []string) {
	if len(results) == 0 {
		fmt.Println("There were no results.")
	} else {
		for _, name := range results {
			fmt.Println(name)
		}
	}

	fmt.Println("")
}

func RawQueries() {
	log.Println("***Performing raw JSON SQL queries***\n")

	var ret []string
	log.Println("Directed by Hideaki Anno:")
	err := db.Select(&ret, "SELECT name FROM mecha WHERE exts -> '$.director' = 'Hideaki Anno'")
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}
	PrintResults(ret)

	ret = make([]string, 0, 1)
	log.Println("Suffered budget woes:")
	err = db.Select(&ret, "SELECT name FROM mecha WHERE exts -> '$.budget_woes' = true")
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}
	PrintResults(ret)

	ret = make([]string, 0, 1)
	log.Println(">1000 space monsters:")
	err = db.Select(&ret, "SELECT name FROM mecha WHERE exts -> '$.space_monsters' > 1000")
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}
	PrintResults(ret)

	ret = make([]string, 0, 1)
	// Different syntax, same result
	log.Println("Directed by Noboru Ishiguro:")
	err = db.Select(&ret, "SELECT name FROM mecha WHERE JSON_EXTRACT(exts, '$.director') = ?", "Noboru Ishiguro")
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}
	PrintResults(ret)

	ret = make([]string, 0, 1)
	// Query an array within the JSON field
	log.Println("Gunbuster appears:")
	err = db.Select(&ret, "SELECT name FROM mecha WHERE JSON_CONTAINS(exts, JSON_QUOTE('Gunbuster'), \"$.mechs\")")
	if err != nil {
		log.Fatalf("SELECT failed: %v", err)
	}
	PrintResults(ret)
}

func SquirrelQueries() {
	log.Println("***Performing JSON SQL queries using Squirrel***\n")

	log.Println("Directed by Hideaki Anno:")
	var ret []string

	// Key and value for ext field.
	k := "director"
	v := "Hideaki Anno"
	query, args, err := sq.Select("name").From("mecha").
		Where(sq.Eq{fmt.Sprintf("exts -> '$.%s'", k): v}).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	err = db.Select(&ret, query, args...)
	if err != nil {
		log.Fatalf("SQL select failed: %v", err)
	}
	PrintResults(ret)

	log.Println("Suffered budget woes:")
	ret = make([]string, 0, 1)

	//

	// TODO not working...
	k = "budget_woes"
	v2 := true
	query, args, err = sq.Select("name").From("mecha").
		Where(sq.Eq{fmt.Sprintf("exts -> '$.%s'", k): v2}).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	log.Printf("Query:\n%s\n", query)
	log.Printf("Args:\n%v\n", args)

	err = db.Select(&ret, query, args...)
	if err != nil {
		log.Fatalf("SQL select failed: %v", err)
	}
	PrintResults(ret)

	//

	log.Println(">1000 space monsters:")
	ret = make([]string, 0, 1)

	k = "space_monsters"
	v3 := 1000
	query, args, err = sq.Select("name").From("mecha").
		Where(sq.Gt{fmt.Sprintf("exts -> '$.%s'", k): v3}).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	err = db.Select(&ret, query, args...)
	if err != nil {
		log.Fatalf("SQL select failed: %v", err)
	}
	PrintResults(ret)

	//

	log.Println("Directed by Noboru Ishiguro:")
	ret = make([]string, 0, 1)

	k = "director"
	v = "Noboru Ishiguro"
	query, args, err = sq.Select("name").From("mecha").
		Where(sq.Eq{fmt.Sprintf("JSON_EXTRACT(exts, '$.%s')", k): v}).
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	err = db.Select(&ret, query, args...)
	if err != nil {
		log.Fatalf("SQL select failed: %v", err)
	}
	PrintResults(ret)

	//

	log.Println("Gunbuster appears:")
	ret = make([]string, 0, 1)

	// TODO...
	query, args, err = sq.Select("name").From("mecha").
		Where("JSON_CONTAINS(exts, JSON_QUOTE('Gunbuster'), \"$.mechs\")").
		ToSql()
	if err != nil {
		log.Fatalf("SQL generation failed: %v", err)
	}

	err = db.Select(&ret, query, args...)
	if err != nil {
		log.Fatalf("SQL select failed: %v", err)
	}
	PrintResults(ret)

}
