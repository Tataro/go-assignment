# Notes

*Database setup, structure change, tests setup, etc.*
## Installing
clone project to your local directory
```bash
$ git clone git@gitlab.com:Upaphong/go-assignment.git
```
go to folder
```sh
$ cd go-assignment
```
change access permission to file `wait-for-it.sh`
```bash
$ chmod +x wait-for-it.sh
```
start docker
```bash
$ docker-compose up -d
```
after start postgres image, create table `knights`
```sql
CREATE TABLE IF NOT EXISTS knights
(
    id SERIAL,
    name TEXT NOT NULL,
    strength INTEGER NOT NULL,
    weapon_power  INTEGER NOT NULL,
    CONSTRAINT knights_pkey PRIMARY KEY (id)
)
```
## Testing
run test
```bash
$ go test -p 1 ./domain ./providers/database ./adapters/http
```

