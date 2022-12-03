
# Golang Exam RestApi App
Simple rest api to create questions and answer questions

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
  KEY                 = "your secret key for JWT"
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


<a name="userData"></a>
## User Data

| email | password     | role               |
| :-------- | :------- | :------------------------- |
| siswapertama@gmail.com | 12345678 | siswa |
| siswakedua@gmail.com | 12345678 | siswa |
| siswaketiga@gmail.com | 12345678 | siswa |
| siswakeempat@gmail.com | 12345678 | siswa |
| guru@gmail.com | 12345678 | guru |


## Endpoint

#### Login

```http
  POST localhost/login
```

| Body | Type     | Description                | Required | 
| :-------- | :------- | :------------------------- | :------------------------- |
| `email` | `string` | email user |  **Required** |
| `password` | `string` | password user with **len 8** | **Required** |

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses login"
    },
    "data": {
        "id": 2,
        "name": "siswa kedua",
        "email": "siswakedua@gmail.com",
        "role": "siswa",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ"
    }
}
```
Response Failed
```bash
{
    "meta": {
        "status": "gagal",
        "code": 400,
        "message": "gagal melakukan login"
    },
    "data": "email tidak terdaftar"
}
```

#### Get User By JWT Token

```http
  GET localhost/v1/user/info
```

| Header | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `Authorization`      | `string` | Set yuor header with value token when you get on login | **Required**

```bash
    Authorization = Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ
```

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses mengambil data user"
    },
    "data": {
        "id": 2,
        "email": "siswakedua@gmail.com",
        "role": "siswa"
    }
}
```
Response Failed
```bash
{
    "meta": {
        "status": "Unauthorized",
        "code": 401,
        "message": "gagal melakukan authorization"
    },
    "data": "signature is invalid"
}
```
#### Get All Question

```http
  GET localhost/v1/questions
```

| Header | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `Authorization`      | `string` | Set yuor header with value token when you get on login | **Required**

```bash
    Authorization = Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ
```

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses mengambil question"
    },
    "data": [
        {
            "id": 1,
            "question": "1 + 1 ?"
        },
        {
            "id": 2,
            "question": "1 + 1 ?"
        },
        {
            "id": 3,
            "question": "1 + 1 ?"
        },
        {
            "id": 4,
            "question": "1 + 1 ?"
        }
    ]
}
```

Response Unauthorized
```bash
{
    "meta": {
        "status": "Unauthorized",
        "code": 401,
        "message": "error"
    },
    "data": "Masukan string Bearer sebelum token"
}
```

#### Get Question By ID

```http
  GET localhost/v1/question/:id
```

| Param | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `id`      | `integer` | Question ID | **Required**

| Header | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `Authorization`      | `string` | Set yuor header with value token when you get on login | **Required**

```bash
    Authorization = Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ
```

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses mengambil question"
    },
    "data": {
        "id": 1,
        "number": 1,
        "question": "Diamanakah ibu kota negara Brazil"
    }
}
```
Response Failed
```bash
{
    "meta": {
        "status": "gagal",
        "code": 400,
        "message": "gagal mengambil question"
    },
    "data": "question dengan id 90 tidak ditemukan"
}
```

Response Unauthorized
```bash
{
    "meta": {
        "status": "Unauthorized",
        "code": 401,
        "message": "error"
    },
    "data": "Masukan string Bearer sebelum token"
}
```

#### Create New Question

```http
  POST localhost/v1/question
```

| Payload | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `question`      | `string` | Question | **Required**

| Header | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `Authorization`      | `string` | Set yuor header with value token when you get on login | **Required**

```bash
    Authorization = Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ
```

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses menyimpan question"
    },
    "data": {
        "id": 1,
        "question": "Diamanakah ibu kota negara Brazil"
    }
}
```
Response Failed
```bash
{
    "meta": {
        "status": "gagal validasi",
        "code": 400,
        "message": "gagal melakukan validasi"
    },
    "data": [
        "error on filed: Question, condition: required"
    ]
}
```

Response Unauthorized (only teacher can create question)
```bash
{
    "meta": {
        "status": "Unauthorized",
        "code": 401,
        "message": "error"
    },
    "data": "akses ditolak"
}
```
```only can create one question in one request```
#### Update Question

```http
  PUT localhost/v1/question/:id
```

| Payload | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `question`      | `string` | Question | **Required**

| Param | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `id`      | `int` | Question ID | **Required**

| Header | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `Authorization`      | `string` | Set yuor header with value token when you get on login | **Required**

```bash
    Authorization = Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ
```

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses mengupdate question"
    },
    "data": {
        "id": 2,
        "question": "2+2 berapa ?"
    }
}
```
Response Failed
```bash
{
    "meta": {
        "status": "gagal",
        "code": 400,
        "message": "gagal mengupdate question"
    },
    "data": "question dengan id 90 tidak ditemukan"
}
```

Response Unauthorized (only teacher can update question)
```bash
{
    "meta": {
        "status": "Unauthorized",
        "code": 401,
        "message": "error"
    },
    "data": "akses ditolak"
}
```
```only can update one question in one request```

#### Delete Question

```http
  DELETE localhost/v1/question/:id
```

| Param | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `id`      | `int` | Question ID | **Required**

| Header | Type     | Description                       |  Required | 
| :-------- | :------- | :-------------------------------- | :-------------------------------- |
| `Authorization`      | `string` | Set yuor header with value token when you get on login | **Required**

```bash
    Authorization = Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNpc3dha2VkdWFAZ21haWwuY29tIiwiZXhwaXJlZCI6MTY2OTk1MTQ1NiwidXNlcklEIjoyfQ.2DJa8GBpxWrYcgVUwoha6rNQmGZGHc7zv9njxPxdDbQ
```

Response Success
```bash
{
    "meta": {
        "status": "sukses",
        "code": 200,
        "message": "sukses menghapus soal"
    },
    "data": "soal berhasil dihapus"
}
```
Response Failed
```bash
{
    "meta": {
        "status": "gagal",
        "code": 500,
        "message": "gagal menghapus question"
    },
    "data": "question dengan id 99 tidak ditemukan"
}
```

Response Unauthorized (only teacher can delete question)
```bash
{
    "meta": {
        "status": "Unauthorized",
        "code": 401,
        "message": "error"
    },
    "data": "akses ditolak"
}
```
```only can delete one question in one request```
