# ConfCLI
#### Client for GetMeConf server

To run ConfCli first build binary file from client.go. Then use following commands and flags:

- command **create** 
TDD

- command **read** reads config or configs from database and writes them to JSON-files, has following flags:
  - **--config-type**: use this flag if you want to read all particular configs,
  config-type correspond to table's name in database. It will create JSON-files with each  config data 
  (could be mongodb, tempconfig, tsconfig);
  - **--config-name**: use this flag with **--config-type** 
if you want to read 1 particular config from database,
 where config-name is a name of particular config in table.
  It will create 1 JSON-file with particular config data.
  - **--outpath**: use this flag to specify folder in which you want to write config-file or files
   (folder should be exist on disc). If flag is not specified file or files will be created in current folder with
   binary file.
  
- command **update** 



- command **delete** 


