1. Clone the repo ```git clone https://github.com/paudelgaurav/atlas-gorm-example.git```
2. Install dependencies ```go mod download```
3. Populate the .env file using the .env.sample file
4. Run server: ```go run main.go```


Generate migration file
``` atlas migrate diff --env gorm ```

Check migration status
``` atlas migrate status --url "mysql://username:password@:port/dbname" ```

Apply migration

``` atlas migrate apply --url "mysql://username:password@:port/dbname"