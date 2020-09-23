package models

import (
	postgresql "Tractor_Details/db"
	"fmt"

	"github.com/revel/revel"
)

// var (
// 	incrementalID int64 = 0
// 	tripDB              = make(map[int64]*Tractor)
// )

// const (
// 	queryInsertTicket = `INSERT INTO tractors (from_location,to_location,user_id) VALUES ($1,$2,$3) RETURNING id`
// 	querySelectTicket = `SELECT id, from_location, to_location, user_id FROM tickets WHERE id=$1`
// )

func (tractor *Tractor) Save() revel.Result {
	tx := postgresql.DB.Create(tractor)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil
	}
	return nil
}

func (tractor *Tractor) Get() revel.Result {

	tx := postgresql.DB.First(&tractor, tractor.ID)

	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil
	}

	return nil
}
