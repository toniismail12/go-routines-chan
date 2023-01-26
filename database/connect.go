package database

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	host := ""
	user := ""
	pass := ""
	dbport := ""
	db_name := ""

	// koneksi ke database
	dsn := "sqlserver://" + user + ":" + pass + "@" + host + ":" + dbport + "?" + "database=" + db_name

	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal connect database")
	} else {
		log.Println("connect success")
		log.Println("processing ...")
	}

	// migrate tabel
	DB = database

	// Migration()

}
