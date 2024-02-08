CREATE TABLE
  IF NOT EXISTS `$mysql_database_name`.`$mysql_table_name` (
    -- Primary key
    `id` int (15) NOT NULL AUTO_INCREMENT,
    
    -- System information
    `Channel`           varchar(255)        NOT NULL COMMENT 'system.Channel',
    `Computer`          varchar(255)        NOT NULL COMMENT 'system.Computer',
    `EventID`           int (15)            NOT NULL COMMENT 'system.EventID',
    `EventRecordID`     int (15)            NOT NULL COMMENT 'system.EventRecordID',
    `ProcessID`         int (15)            NOT NULL COMMENT 'system.Execution.ProcessID',
    `ThreadID`          int (15)            NOT NULL COMMENT 'system.Execution.ThreadID',
    `Keywords`          int (16)            NOT NULL COMMENT 'system.Keywords',
    `Level`             int (15)            NOT NULL COMMENT 'system.Level',
    `Opcode`            int (15)            NOT NULL COMMENT 'system.Opcode',
    `ProviderGuid`      varchar(255)        NOT NULL COMMENT 'system.Provider.Guid',
    `ProviderName`      varchar(255)        NOT NULL COMMENT 'system.Provider.Name',
    `UserID`            varchar(255)        NOT NULL COMMENT 'system.Security.UserID',
    `Task`              int (15)            NOT NULL COMMENT 'system.Task',
    `SystemTime`        Timestamp           NOT NULL COMMENT 'system.TimeCreated.SystemTime',
    `Version`           int (5)             NOT NULL COMMENT 'system.Version',

    PRIMARY KEY (`id`)
  ) ENGINE = InnoDB CHARSET = utf8mb4;