INSERT INTO `test`.`evtx`
    (`system.Channel`,
     `system.Computer`,
     `system.EventID`,
     `system.EventRecordID`,
     `system.Execution.ProcessID`,
     `system.Execution.ThreadID`,
     `system.Keywords`,
     `system.Level`,
     `system.Opcode`,
     `system.Provider.Guid`,
     `system.Provider.Name`,
     `system.Security.UserID`,
     `system.Task`,
     `system.TimeCreated.SystemTime`,
     `system.Version`)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)