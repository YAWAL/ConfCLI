# ConfCLI
#### Client for GetMeConf server

To run ConfCLI first build binary file from client.go. Then use following commands and flags:

- command **create** creates new config (reads csv file with config data and
writes new cortege to particular table in database), has following flags:
  - **--config-type**: use this flag to specify which config you want to
  write to database, config-type corresponds to table's name in database
  (could be mongodb, tempconfig, tsconfig).
  - **--file-name**: use this flag to specify csv file with config data
  to persist in database. CSV file should contains proper data according
   to config structure and be on the same level as binary.



- command **read** reads config or configs from database and writes them to JSON-files, has following flags:
  - **--config-type**: use this flag if you want to read all particular configs,
  config-type corresponds to table's name in database. It will create JSON-files with each  config data
  (could be mongodb, tempconfig, tsconfig);
  - **--config-name**: use this flag with **--config-type** 
if you want to read 1 particular config from database,
 where config-name is a name of particular config in table.
  It will create 1 JSON-file with particular config data.
  - **--outpath**: use this flag to specify folder in which you want to write config-file or files
   (folder should be exist on disc). If flag is not specified file or files will be created in current folder with
   binary file.
  
- command **update** updates existing config in database
 (reads csv file with config data to be updated and
 updates existing cortege in particular table), has following flags:
  - **--config-type**: use this flag to specify which config you want to
  update, config-type corresponds to table's name in database
  (could be mongodb, tempconfig, tsconfig).
  - **--file-name**: use this flag to specify csv file with config data
  to be updated. CSV file should contains proper data according
   to config structure and be on the same level as binary.



- command **delete** deletes particular config data from table in database, has
following flags:

  - **--config-type**: use this flag to specify config type
  (it corresponds to config table's name in database. Could be mongodb, tempconfig, tsconfig);
  - **--config-name**: use this flag to specify particular config name from table.


