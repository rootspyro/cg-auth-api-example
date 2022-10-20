package db

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type ClientBuilder struct {
	host string
	port int
	dbName string
	username string
	password string
}

func NewClientBuilder(_host string,_port string,_db string, _user string, _password string) ClientBuilder {
	
	intPort, err := strconv.Atoi(_port)
	_ = err

	return ClientBuilder{
		host: _host,
		port: intPort,
		dbName: _db,
		username: _user,
		password: _password,
	}
}

func (b *ClientBuilder) BuildSqlClient() *sql.DB {

	//connString := "Driver={PostgreSQL};Server=localhost;Port=5434;Database=auth-api;Uid=rootspyro;Pwd=Fyy1I3xe7ssaRzyPla8obAXyJ;"
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", b.host, b.port, b.username, b.password, b.dbName)
	db, err := sql.Open("postgres", connString)

	if ( err != nil ) {
		log.Fatalln(err)
	}

	return db;

}
