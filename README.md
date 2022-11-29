
# Golang Exam RestApi App


## Run Locally

Clone the project

```bash
  git clone https://github.com/rezairfanwijaya/go-exam-api.git
```

Go to the project directory

```bash
  cd go-exam-api
```

#### Before running the project, first create a database using the MySQL DBMS, then setup env file in root project

Setup env

```bash
  DATABASE_HOST       = "your localhost port"
  DATABASE_NAME       = "your database name"
  DATABASE_PORT       = "your database port"
  DATABASE_USERNAME   = "your database username"
  DATABASE_PASSWORD   = "your database password
```
#### If the database does not have a password, leave it blank like this ```DATABASE_PASSWORD = ""```

Run migration table

```bash
  go run main.go
```

#### after that 
run migration up data user
```bash
  migrate -database "mysql://yourdatabaseusername:yourdatabasepassword@tcp(yourdatabasehost:yourdatabaseport)/yourdatabasename" -path db/migration up
```


## User Data

| email | password     | role               |
| :-------- | :------- | :------------------------- |
| siswapertama@gmail.com | 12345678 | siswa |
| siswakedua@gmail.com | 12345678 | siswa |
| siswaketiga@gmail.com | 12345678 | siswa |
| siswakeempat@gmail.com | 12345678 | siswa |
| guru@gmail.com | 12345678 | guru |



