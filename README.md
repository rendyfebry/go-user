# Simple HTTP Server

## Usage

```
go run cmd/main.go
```

## Request

### Create User

```
curl -L -X POST 'localhost:3000/users' \
-H 'Content-Type: application/json' \
--data-raw '{
    "name": "yo",
    "email": "yo@example.com"
}'
```

### Get All Users

```
curl -L -X GET 'localhost:3000/users'
```

### Get Single User

```
curl -L -X GET 'localhost:3000/users/d0465c53-94a4-47c1-8e0f-45eafaa96583'
```

### Update User
```
curl -L -X PUT 'localhost:3000/users/08e93f8a-316a-4e46-a416-7a12665111a0' \
-H 'Content-Type: application/json' \
--data-raw '{
    "name": "kiss",
    "email": "yo@example.com"
}'
```

### Delete User

```
curl -L -X DELETE 'localhost:3000/users/eb8c9481-e59f-40e6-b1dd-90b27f4de5ac'
```
