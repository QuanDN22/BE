package models

// =IF(SORTBY(SEQUENCE(10000),RANDARRAY(10000))<4500,"True","False")

import (
	"fmt"
	"log"
	"strconv"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func AddData_Init(db *gorm.DB) {
	// read excel document
	f, err := excelize.OpenFile("./data/data_server_1.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	// // Get value from cell by given worksheet name and cell reference.
	// cell, err := f.GetCellType("Sheet1", "C5")
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	// fmt.Printf("%T", cell)

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Println(err)
		return
	}

	// remove first element in the slice rows
	rows = append(rows[:0], rows[1:]...)

	// add server in database with three 3 fields: server_name, server_ip, server_status
	for _, row := range rows {
		status, err := strconv.ParseBool(row[2])
		if err != nil {
			log.Println(err)
			continue
		}
		result := db.Create(&Server{
			Server_Name:   row[0],
			Server_IPv4:   row[1],
			Server_Status: status,
		})
		fmt.Println(result.RowsAffected, result.Error)
		log.Printf("created server: %s", row)
	}
}
