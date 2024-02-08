CREATE TABLE IF NOT EXISTS `EVTX` (
  -- Primary key
  `id` int(15) NOT NULL AUTO_INCREMENT,

  -- System information
  `system.Channel`                  varchar(255) NOT NULL,
  `system.Computer`                 varchar(255) NOT NULL,
  `system.EventID`                  int(15) NOT NULL,
  `system.EventRecordID`            int(15) NOT NULL,
  `system.Execution.ProcessID`      int(15) NOT NULL,
  `system.Execution.ThreadID`       int(15) NOT NULL,
  `system.Keywords`                 int(16) NOT NULL,
  `system.Level`                    int(15) NOT NULL,
  `system.Opcode`                   int(15) NOT NULL,
  `system.Provider.Guid`            varchar(255) NOT NULL,
  `system.Provider.Name`            varchar(255) NOT NULL,
  `system.Security.UserID`          varchar(255) NOT NULL,
  `system.Task`                     int(15) NOT NULL,
  `system.TimeCreated.SystemTime`   Timestamp NOT NULL,
  `system.Version`                  int(5) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB CHARSET=utf8mb4;