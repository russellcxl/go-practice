# Assignment 3

## Features:
- TCP server/client connection
- Protobuf messages
- Checksum authentication
- User can check leave balance / apply for leave / cancel leave application / logout
- Manager functions (approve/reject leaves) not yet implemented

## How to start:
- run `go get` to sync all dependencies
- run sql migration script `database/migration.sql`
- change mysql DB credentials in `database/db.go`  
- run server at specified port `go run server/main.go 9000`
- dial in with client `go run client/main.go 9000`
- login with sample staff userid = `1001` and password = `123`

## Libraries used:
- [gORM](https://github.com/jinzhu/gorm)
- [Redis](https://github.com/go-redis/redis)
- [go-pretty](https://github.com/jedib0t/go-pretty) -- for pretty table printing

## Models:
**User**
```
message User {
  uint64 user_id = 1;
  string password = 2;
  string name = 3;
  uint32 team_id = 4;
  int32  role = 5;
  uint32 leave_balance = 6;
}
```
**Leave**
```
message Leave {
  uint64 leave_id = 1;
  uint64 user_id = 2;
  uint32 team_id = 3;
  uint32 start_time = 4;
  uint32 end_time = 5;
  uint32 days_taken = 6;
  int32  status = 7;
  uint64 approver_id = 8;
}
```

## Resources:
- [How gorm knows which table to make changes to](https://gorm.io/docs/conventions.html#Pluralized-Table-Name)
- [When to flush data from buffer](https://stackoverflow.com/questions/49166370/why-do-you-need-flush-at-all-if-close-is-enough)
- [Never pass a pointer to an interface](https://golang.org/doc/faq#pointer_to_interface)
- [When to use pointers](https://medium.com/@meeusdylan/when-to-use-pointers-in-go-44c15fe04eac)

## TO-DO:
- DB txn; rollback/commit