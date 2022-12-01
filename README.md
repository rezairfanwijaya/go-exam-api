
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
  DATABASE_HOST       = "your localhost host"
  DATABASE_NAME       = "your database name"
  DATABASE_PORT       = "your database port"
  DATABASE_USERNAME   = "your database username"
  DATABASE_PASSWORD   = "your database password"
```
#### If the database does not have a password, leave it blank like this ```DATABASE_PASSWORD = ""```

Run migration table

```bash
  go run main.go
```
result


![image](https://user-images.githubusercontent.com/87264553/204498813-ad29cd5c-c293-4e3a-b8b0-176ea0b8aa4f.png)


#### after that 
run migration up data user with your password database
```bash
  migrate -database "mysql://yourdatabaseusername:yourdatabasepassword@tcp(yourdatabasehost:yourdatabaseport)/yourdatabasename" -path db/migration up
```
run migration up data user without your password database
```bash
  migrate -database "mysql://yourdatabaseusername@tcp(yourdatabasehost:yourdatabaseport)/yourdatabasename" -path db/migration up
```
result


![image](https://user-images.githubusercontent.com/87264553/204498884-7580de99-776b-4748-b2c8-01d151c17ebc.png)



## User Data

| email | password     | role               |
| :-------- | :------- | :------------------------- |
| siswapertama@gmail.com | 12345678 | siswa |
| siswakedua@gmail.com | 12345678 | siswa |
| siswaketiga@gmail.com | 12345678 | siswa |
| siswakeempat@gmail.com | 12345678 | siswa |
| guru@gmail.com | 12345678 | guru |



