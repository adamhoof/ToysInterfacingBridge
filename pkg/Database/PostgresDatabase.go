package Database

import (
	"database/sql"
	"fmt"
	"github.com/adamhoof/ToysInterfacingBridge/pkg/Toy"
	"github.com/lib/pq"
	"os"
	"strconv"
)

type PostgresDatabase struct {
	db *sql.DB
}

const updateToyModeSQLStatement = `UPDATE HomeAppliances SET current_mode = $2 WHERE name = $1;`
const toysDataQuery = `SELECT name, available_modes, id, publish_topic, subscribe_topic FROM HomeAppliances;`

func (postgres *PostgresDatabase) Connect() {
	port, err := strconv.Atoi(os.Getenv("dbPort"))
	if err != nil {
		panic(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s",
		os.Getenv("dbEndpoint"), port, os.Getenv("dbUser"), os.Getenv("dbPassword"), os.Getenv("dbName"))

	postgres.db, err = sql.Open(os.Getenv("dbDriver"), psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
}

func (postgres *PostgresDatabase) TestConnection() {
	result := postgres.db.Ping()
	if result != nil {
		panic(result)
	}
}

func (postgres *PostgresDatabase) Disconnect() {
	err := postgres.db.Close()
	if err != nil {
		fmt.Println()
	}
}

func (postgres *PostgresDatabase) UpdateToyMode(toyName string, toyMode string) {
	_, err := postgres.db.Exec(updateToyModeSQLStatement, toyName, toyMode)
	if err != nil {
		fmt.Println("Couldn't update toy mode: ", err)
	}
}

func (postgres *PostgresDatabase) PullToysData(toyBag map[string]Toy.Toy) {
	rows, err := postgres.db.Query(toysDataQuery)
	if err != nil {
		fmt.Println("unable to query data", err)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			fmt.Println("unable to close rows", err)
		}
	}(rows)

	for rows.Next() {
		toy := Toy.Toy{}
		err = rows.Scan(&toy.Name, pq.Array(&toy.AvailableCommands), &toy.ID, &toy.PublishTopic, &toy.SubscribeTopic)
		if err != nil {
			fmt.Println("unable to fetch toy data into toy", err)
		}
		toyBag[toy.Name] = toy
	}
}
