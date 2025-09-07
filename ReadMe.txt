How to run

part1_Golang
- cd to part1_Golang/api_gateway
- docker compose up --build
- can use swagger to action: http://localhost:8080/swagger/index.html

User JWT: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoidXNlci0xMjMiLCJyb2xlIjoidXNlciJ9.1gojFYgFALzF54kIfXakaioIaDpAnmj4E17WPxXRNUY
Admin JWT: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiYWRtaW4tMTIzIiwicm9sZSI6ImFkbWluIn0.pwlzFI0womM5dmF4NMLKZBWc_8boXPVDiz2HMlUuC8M

- Use User JWT for List/Add/Get 
- Use Admin JWT for delete

For demo I set and validate expire time.

part2_Concurrency
- cd to part2_Concurrency
- Can edit 5 urls in a txt file.
- go run main.go

part3_BuggyCode
- cd to part3_BuggyCode
- go run main.go
- explain in main.go
