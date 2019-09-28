package data

import (
	"database/sql"
	"github.com/aasumitro/ego-worker/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
)

// Messaging Queue Connection (RabbitMQ)
func MQConnection(mqConf *helper.RabbitMQConfig) (*amqp.Connection, error) {
	// Open up our messaging queue connection.
	mq, err := amqp.Dial(mqConf.URL)
	// if there is an error opening the connection, handle it
	helper.CheckError(err, "Failed to connect to MQ")
	// return database to use in any function
	return mq, nil
}

// Database Connection (MySQL)
func DBConnection(sqlConf *helper.MySQlConfig) (*sql.DB, error) {
	// Open up our database connection.
	db, err := sql.Open(sqlConf.Driver, sqlConf.URL)
	// if there is an error opening the connection, handle it
	helper.CheckError(err, "Failed connect to DB")
	// return database to use in any function
	return db, nil
}
