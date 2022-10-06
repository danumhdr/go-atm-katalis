# go-atm-katalis

Demo flow if we using ATM/Transaction in Bank Account

Pre-requisites:
- Go
- MySQL

1. Download this Database & restore it
2. Check & configure Database in config/database.go
3. Build with run "go build ."
4. Use this comment :
    - login `{name:string}`
    - deposit `{amount:integer}`
    - withdraw `{amount:integer}`
    - transfer `{name:string}` `{amount:integer}`
    - logout
