package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/KnightChaser/sentinela"
	"github.com/schollz/progressbar/v3"
	"github.com/tidwall/gjson"
)

func MigrateEVTX2MySQL(filepath string, databaseConnection *sql.DB, tableName string) {

	fmt.Println("::: Migrating EVTX to MySQL...")

	sysmonEvtxFile := filepath

	stats, err := sentinela.ParseEVTX(sysmonEvtxFile)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a query structure to insert the data to MySQL (prepared statement)
	queryStructureByte, err := os.ReadFile("./database/evtxTableInsert.sql")
	queryStructure := string(queryStructureByte)
	queryStructure = strings.Replace(queryStructure, "$mysql_database_name", os.Getenv("MYSQL_DATABASE_NAME"), 1)
	queryStructure = strings.Replace(queryStructure, "$mysql_table_name", tableName, 1)

	if err != nil {
		log.Panic(err)
	}

	// Display the statistics
	bar := progressbar.Default(-1)

	for _, stat := range stats.Event {
		// Extracting system data(common) in the given EVTX file
		systemChannel := gjson.Get(stat, "Event.System.Channel").String()
		systemComputer := gjson.Get(stat, "Event.System.Computer").String()
		systemEventID := gjson.Get(stat, "Event.System.EventID").Uint()
		systemEventRecordID := gjson.Get(stat, "Event.System.EventRecordID").Uint()
		systemExecutionProcessID := gjson.Get(stat, "Event.System.Execution.ProcessID").Uint()
		systemExecutionThreadID := gjson.Get(stat, "Event.System.Execution.ThreadID").Uint()
		systemKeywords := gjson.Get(stat, "Event.System.Keywords").Uint()
		systemLevel := gjson.Get(stat, "Event.System.Level").Uint()
		systemOpcode := gjson.Get(stat, "Event.System.Opcode").Uint()
		systemProviderGuid := gjson.Get(stat, "Event.System.Provider.Guid").String()
		systemProviderName := gjson.Get(stat, "Event.System.Provider.Name").String()
		systemSecurityUserID := gjson.Get(stat, "Event.System.Security.UserID").String()
		systemTask := gjson.Get(stat, "Event.System.Task").Uint()
		systemTimeCreatedSystemTime := gjson.Get(stat, "Event.System.TimeCreated.SystemTime").Time()
		systemVersion := gjson.Get(stat, "Event.System.Version").Uint()

		// Registering the event data to MySQL
		prepatedStatment, err := databaseConnection.Prepare(string(queryStructure))
		if err != nil {
			log.Fatal(err)
			continue
		}

		// Inserting the data to MySQL
		_, err = prepatedStatment.Exec(
			systemChannel,
			systemComputer,
			systemEventID,
			systemEventRecordID,
			systemExecutionProcessID,
			systemExecutionThreadID,
			systemKeywords,
			systemLevel,
			systemOpcode,
			systemProviderGuid,
			systemProviderName,
			systemSecurityUserID,
			systemTask,
			systemTimeCreatedSystemTime,
			systemVersion)
		if err != nil {
			log.Fatal(err)
		}

		// log.Printf("Event data has been inserted to MySQL: %v", response)
		bar.Add(1)
	}

	fmt.Printf("::: The given EVTX file(%s) has been migrated to MySQL!", sysmonEvtxFile)
}
