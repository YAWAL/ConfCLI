# ConfCLI
Client for GetMeConf server

To run ConfCli first build binary file client.go. Then use 1 of the following commands:

- command **create** 
TDD

- command **read** reads config or configs from database and write them into JSON-files, has following flags:
  - **--config-type** (use this flag if you want to read all particular configs,
  config-type correspond to table's name in database. It will create JSON-files with each  config data);
  - **--config-name** (use this flag with **--config-type** 
if you want to read 1 particular config from database,
 where config-name is a name of particular config in table.
  It will create 1 JSON-file with particular config data).
  
- command **update** 



- command **delete** 


