# ch-xml2sql
ch-xml2sql is a simple tool to convert clickhouse xml dictionary definitions to SQL DDL queries.

> The status of this repository is **(WIP) Work In Progress** and lacks a lot of features.


# Usage
You can run the tool with the following command:
```shell
go run main.go -xmlDir <PATH_TO_DIRECTORY_OF_YOUR_XML_FILES> -out <PATH_TO_OUTPUT_FILE>
```


## Features
- [x] Convert XML dictionary definition to SQL DDL query
- [ ] Types of dictionary layouts
  - [x] hashed
  - [x] complex_key_hashed
  - [ ] flat
  - [ ] sparse_hashed
  - [ ] complex_key_sparse_hashed
  - [ ] hashed_array
  - [ ] complex_key_hashed_array
  - [ ] range_hashed
  - [ ] complex_key_range_hashed
  - [ ] cache
  - [ ] complex_key_cache
  - [ ] ssd_cache
  - [ ] complex_key_ssd_cache
  - [ ] direct
  - [ ] complex_key_direct
  - [ ] ip_trie
- [ ] Types of lifetimes
  - [x] with `min` and `max`
  - [ ] with `lifetime`
- [ ] Types of sources
  - [ ] Clickhouse
    - [x] `host`, `port`, `user`, `password`, `db`, `table`, `where`
    - [ ] `secure`
    - [ ] `query`
    - [ ] `invalidate_query`
  - [x] http
  - [x] file
  - [ ] executable file
  - [ ] executable pool
  - [ ] executable pool
  - [ ] DBMS
  - [ ] Null
- [x] Convert Structure
- [ ] Convert settings
- [ ] Convert comment