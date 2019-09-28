package data

import (
	"fmt"
	"github.com/aasumitro/ego-worker/helper"
)

const (
	DEFAULT = "DEFAULT"
	PENDING = "PENDING"
	PROCESS = "PROCESS"
	FAILED  = "FAILED"
	SENT    = "SENT"
)

type Message struct {
	ID       int
	From     string
	TO       string
	Subject  string
	Message  string
	Status   string
	Template string
}

func StoreOutbox(msg Message) error {
	//load config
	config := helper.GetConfig()
	// call database connection function
	db, _ := DBConnection(config.MySQL)
	//defer the close till after the main function has finished executing
	defer db.Close()
	// data query insert
	query := fmt.Sprintf(
		"INSERT INTO outbox (`from`, `to`, `subject`, `message`, `template`) VALUES ('%s', '%s',  '%s', '%s', '%s')",
		msg.From,
		msg.TO,
		msg.Subject,
		msg.Message,
		msg.Template)
	// perform a db.Query insert
	insert, err := db.Query(query)
	// if there is an error inserting, handle it
	helper.CheckError(err, "Failed insert new record")
	if err != nil {
		return err
	}
	// be careful deferring Queries if you are using transactions
	// defer the close till after the main function has finished executing
	defer insert.Close()

	return nil
}

func FetchOutbox() []Message {
	//load config
	config := helper.GetConfig()
	// call database connection function
	db, _ := DBConnection(config.MySQL)
	//defer the close till after the main function has finished executing
	defer db.Close()
	// data query insert
	// perform a db.Query insert
	query := fmt.Sprintf("SELECT * FROM outbox WHERE status<>'PROCESS' ORDER BY id ASC LIMIT 1")
	rows, err := db.Query(query)
	// if there is an error inserting, handle it
	helper.CheckError(err, "Failed fetch message")
	if err != nil {
		return nil
	}
	// be careful deferring Queries if you are using transactions
	// defer the close till after the main function has finished executing
	defer rows.Close()
	var result []Message
	// Next prepares the next result row for reading with the Scan method.
	for rows.Next() {
		var each = Message{}
		// Scan copies the columns in the current row into the values pointed
		var err = rows.Scan(
			&each.ID,
			&each.From,
			&each.TO,
			&each.Subject,
			&each.Message,
			&each.Status,
			&each.Template)
		if err != nil {
			helper.CheckError(err, "Failed scan row")
			return nil
		}
		// The append built-in function appends elements to the end of a slice.
		result = append(result, each)
	}
	// Err returns the error, if any, that was encountered during iteration.
	if err = rows.Err(); err != nil {
		helper.CheckError(err, "Failed get row")
		return nil
	}
	// return result
	return result
}

func MoveOutboxToSent(msg Message) {
	//load config
	config := helper.GetConfig()
	// call database connection function
	db, _ := DBConnection(config.MySQL)
	//defer the close till after the main function has finished executing
	defer db.Close()
	// delete query
	delQuery := fmt.Sprintf("DELETE FROM outbox where id=%d", msg.ID)
	// exec query
	del, errExec := db.Exec(delQuery)
	// handle error
	helper.CheckError(errExec, "Failed delete outbox")
	// get affected row
	delAffected, errDelete := del.RowsAffected()
	// handle error
	helper.CheckError(errDelete, "No row affected")
	// run if success delete outbox
	if delAffected == 1 {
		// insert query
		insQuery := fmt.Sprintf(
			"INSERT INTO sent (`from`, `to`, `subject`, `message`, `status`, `template`) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')",
			msg.From,
			msg.TO,
			msg.Subject,
			msg.Message,
			SENT,
			msg.Template)
		// perform a db.Query insert
		ins, errQuery := db.Query(insQuery)
		// if there is an error inserting, handle it
		helper.CheckError(errQuery, "Failed insert new record")
		// be careful deferring Queries if you are using transactions
		// defer the close till after the main function has finished executing
		defer ins.Close()
	}
}
