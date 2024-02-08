INSERT INTO `$mysql_database_name`.`$mysql_table_name`
    (`Channel`,
     `Computer`,
     `EventID`,
     `EventRecordID`,
     `ProcessID`,
     `ThreadID`,
     `Keywords`,
     `Level`,
     `Opcode`,
     `ProviderGuid`,
     `ProviderName`,
     `UserID`,
     `Task`,
     `SystemTime`,
     `Version`)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)