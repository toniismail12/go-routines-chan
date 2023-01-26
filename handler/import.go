package handler

import (
	"ch-export/database"
	"ch-export/models"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Import(filename string) error {

	fileName := filename

	xlsx, err := excelize.OpenFile("./files/" + fileName)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	c := make(chan string)
	var wg sync.WaitGroup

	start := time.Now()

	for _, sheetName := range xlsx.GetSheetMap() {

		rows := xlsx.GetRows(sheetName)

		jr := len(rows)

		for i := 2; i < jr+1; i++ {

			// if i == 100000 {
			// 	break
			// }

			periode := xlsx.GetCellValue(sheetName, fmt.Sprintf("K%d", i))
			date := periode[:4] + "-" + periode[4:] + "-01"
			// fmt.Println(date)

			wg.Add(1) // This tells the waitgroup, that there is now 1 pending operation here

			go insert_Data(
				xlsx.GetCellValue(sheetName, fmt.Sprintf("A%d", i)),
				xlsx.GetCellValue(sheetName, fmt.Sprintf("M%d", i)),
				xlsx.GetCellValue(sheetName, fmt.Sprintf("G%d", i)),
				date,
				xlsx.GetCellValue(sheetName, fmt.Sprintf("P%d", i)),
				xlsx.GetCellValue(sheetName, fmt.Sprintf("S%d", i)),
				c,
				&wg,
				i,
			)

		}

		go func() {
			wg.Wait() // this blocks the goroutine until WaitGroup counter is zero
			close(c)  // Channels need to be closed, otherwise the below loop will go on forever
		}() // This calls itself

		// this shorthand loop is syntactic sugar for an endless loop that just waits for results to come in through the 'c' channel

		for msg := range c {
			log.Println(msg)
		}

	}

	log.Println("import selesai")
	log.Println("finished in ", time.Since(start))

	return nil

}

func insert_Data(
	badge,
	cost_center,
	payroll_area,
	periode,
	wage_type,
	amount string,
	c chan string,
	wg *sync.WaitGroup,
	index int,
) {

	defer (*wg).Done()

	insert := models.Test_bareng_wtr{

		Badge:        badge,
		Cost_center:  cost_center,
		Payroll_area: payroll_area,
		Periode:      periode,
		Wage_type:    wage_type,
		Amount:       amount,
		Created_at:   time.Now().Local(),
		Updated_at:   time.Now().Local(),
		File_code:    "X1X1",
	}

	// save to table
	err := database.DB.Create(&insert)
	if err != nil {
		log.Println(err)
	}

	c <- strconv.Itoa(index) + " Success insert to table " + badge

}
