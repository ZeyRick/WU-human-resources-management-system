## Prerequisites

- Go 1.16 or later

## How To Run

1. Clone the repository.
2. Run `go install` to download all the packages needed.
3. Make sure you create the .env file by copying the .env.example and rename it to .env
4. Run `go run main.go` to start the application.

### Project Layout

```tree
├── adapters
│   ├── controllers
│   └── db
├── core
│   ├── models
├── └── services
└── pkg
```

- `adapter` is the code who adapt the business logic code with external. resource like a HTTP request or database.
- `core` this is where we put our business logic related code.
- `core/model` this is for the database model.
- `main.go` this is the main file that run our backend.
- `pkg` this is where we store all the packages that we use thru out the program

### Database

- Database is Using MySql. You can run it in docker easily with docker desktop. tutorial https://www.youtube.com/watch?v=vNH3aOTDjsw
- For databsae orm we are using gorm you can read more about it here:
  - gorm: https://gorm.io/docs/index.html
  - what is orm: https://www.freecodecamp.org/news/what-is-an-orm-the-meaning-of-object-relational-mapping-database-tools/ (Basiclly just a technique to connect object in your code with data in your database)

### Migration
- After you successfully run the MySql In Docker you can connect to it using any type of MySql Client (Beekeeper ...etc)
- Follow the steps bellow to migrate all the tables into your database
  1. Create a database name "hr_management"
  2. Install migration tool with command "go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest"
  3. Go to backend file directory and set the DATABASE_CONNECTION_STRING in enviroment variable (.env) file.
  4. Run command "make db-up" this will migrate all the tables into your database
  - NOTE: if make up doesn't work try to see if your computer has "make" installed (How to install make: https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows RECOMMANDED TO USE Chocolatey for windows). If you still can't get make to work you can run the command "migrate -source file:./migrations -database mysql://[your DATABASE_CONNECTION_STRING] up"

  You can read more about the migrations tool here: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### How to migrate new table

- Run command "make db-craete name=[your migration name]" this will craete 2 sql file in ./migrations where you can write your sql code. after you done writing sql you can run "make db-up".
- Please read more about migration at "https://github.com/golang-migrate/migrate/tree/master/cmd/migrate" if you have any trouble
