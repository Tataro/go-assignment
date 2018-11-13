# Notes

*Database setup, structure change, tests setup, etc.*
## Development
clone project to your local directory
```bash
$ git clone git@gitlab.com:Upaphong/go-assignment.git
```
go to folder
```bash
$ cd go-assignment
```
install dependencies
```bash
$ dep ensure
```
change db host to `localhost` in `providers/database/provider.go`  
after start postgres image, create table `knights`
```sql
CREATE TABLE IF NOT EXISTS knights
(
    id SERIAL,
    name TEXT NOT NULL,
    strength INTEGER NOT NULL,
    weapon_power  INTEGER NOT NULL,
    CONSTRAINT knights_pkey PRIMARY KEY (id)
);
```

## Docker Setup
change access permission to file `wait-for-it.sh`
```bash
$ chmod +x wait-for-it.sh
```
start docker
```bash
$ docker-compose up -d
```

## Testing
run test
```bash
$ go test -p 1 ./domain ./providers/database ./adapters/http
```

