package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/KnightChaser/sentinela"
	"github.com/tidwall/gjson"
)

func MigrateEVTX2MySQL(filepath string, databaseConnection *sql.DB) {
	fmt.Println("Migrating EVTX to MySQL...")

	sysmonEvtxFile := filepath

	stats, err := sentinela.ParseEVTX(sysmonEvtxFile)
	if err != nil {
		log.Fatal(err)
	}

	// Display the statistics
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
		queryStructure, err := os.ReadFile("./database/insert.sql")
		if err != nil {
			log.Panic(err)
		}

		prepatedStatment, err := databaseConnection.Prepare(string(queryStructure))
		if err != nil {
			log.Fatal(err)
			continue
		}

		// Inserting the data
		err = nil
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
	}
}
