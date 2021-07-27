package script

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type RestaurantEventData struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Archived   bool   `json:"archived"`
	Locality   string `json:"locality"`
	LatLong    string `json:"lat_long"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
}

func dbConn(user, password, database, table string) *sql.DB {
	db, err := sql.Open("msql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, database, table))
	if err != nil {
		panic(err.Error())
	}
	return db
}
func main() {
	dbUser := flag.String("user", "root", "")
	dbPassword := flag.String("password", "root", "")
	database := flag.String("database", "127.0.0.1:3006", "")
	table := flag.String("table", "restaurants", "")
	dbConn := dbConn(*dbUser, *dbPassword, *database, *table)
	defer dbConn.Close()

	results, err := dbConn.Query(fmt.Sprintf("SELECT * FROM %s", *table))
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var id int64
		var archived bool
		var name, locality, address, lat_long, postal_code string
		err := results.Scan(id, name, archived, locality, address, lat_long, postal_code)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(RestaurantEventData{
			Id:         id,
			Name:       name,
			Archived:   archived,
			Locality:   locality,
			LatLong:    lat_long,
			Address:    address,
			PostalCode: postal_code,
		})
	}
}
