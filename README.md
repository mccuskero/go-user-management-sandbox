# Summary

This is a user management framework used to test out APIs, and JWT tokens.

## Overview

The simulator can:

- create users and their profiles
- simulate energy usages based on profiles

## Setting up Postgres

Postgres can be run within Docker, I am using Rancher to do so.

- ```helm install user-mgmt-postgresql bitnami/postgres```
- Setup port forwarding in rancher dashboard on port 5432 (matches user-mgmt-sandbox.env)
- Get the password from k8s secret
```export POSTGRES_PASSWORD=$(kubectl get secret --namespace default user-mgmt-postgresql -o jsonpath="{.data.postgres-password}" | base64 -d)```
- test connection to database: psql
```psql --host 127.0.0.1 -U postgres -p 5432```
- List databases
```postgres=# \l```
- Create the database
```CREATE DATABASE user_mgmt_sandbox;```
- Install UUID connection
```\c <database_name>```
- ADD in UUID extensions, to create them when adding in a user
```CREATE EXTENSION IF NOT EXISTS "uuid-ossp";```

## Create JWT public and private keys, and base64 encode them

You will need both private and public keys for access tokens, and also refresh token. The private key is used to create the token, and the public key is used to decode the token. See token.go. 

- Create a private pem key

```openssl genrsa -out jwtRSA256-private-access.pem 2048```

- Create the public key

```openssl rsa -in jwtRSA256-private-access.pem -pubout -outform PEM -out jwtRSA256-public-acceess.pem```

- base64 encode the private key, and cut/paste it to the user-mgmt-sandbox.env file

```cat jwtRSA256-private-access.pem | base64```

- base64 encode the public key, and cut/past it to the env file

```cat  jwtRSA256-public-acceess.pem | base64```

## Running the Simulator

- Run the simulator
```go run ./main.go```
or
```./go-user-management-sandbox```

output should be:

```output
➜  go-user-management-sandbox git:(main) ✗ ./go-user-management-sandbox
Connected Successfully to the Database
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/healthchecker        --> main.main.func1 (4 handlers)
[GIN-debug] POST   /api/auth/register        --> github.com/mccuskero/go-user-management-sandbox/pkg/services/controllers.(*AuthController).RegisterUser-fm (4 handlers)
[GIN-debug] POST   /api/auth/login           --> github.com/mccuskero/go-user-management-sandbox/pkg/services/controllers.(*AuthController).SignInUser-fm (4 handlers)
[GIN-debug] GET    /api/auth/profile         --> github.com/mccuskero/go-user-management-sandbox/pkg/services/controllers.(*AuthController).UserProfile-fm (4 handlers)
[GIN-debug] GET    /api/auth/refresh         --> github.com/mccuskero/go-user-management-sandbox/pkg/services/controllers.(*AuthController).RefreshAccessToken-fm (4 handlers)
[GIN-debug] GET    /api/auth/logout          --> github.com/mccuskero/go-user-management-sandbox/pkg/services/controllers.(*AuthController).LogoutUser-fm (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000


[GIN] 2024/02/25 - 14:09:25 | 200 |   69.496818ms |       127.0.0.1 | POST     "/api/auth/login"
[GIN] 2024/02/25 - 14:09:29 | 200 |   23.087965ms |       127.0.0.1 | GET      "/api/auth/profile"
[GIN] 2024/02/25 - 14:09:34 | 200 |     6.75038ms |       127.0.0.1 | GET      "/api/auth/refresh"
```

## Verifying the API with Insomnia 

- Check Healthchecker from browser and/or curl
Create a get command and enter... 
go to ```http://localhost:8000/api/healthchecker```
- or run curl
```curl http://localhost:8000/api/healthchecker```
response is 
```{"message":"Welcome to User Management Sandbox","status":"success"}```

- Start to register users using POST
go to ```localhost:8000/api/auth/register ```

enter in JSON...

```output
{
   "Name":"test2",
   "Email":"test2@test.com",
   "Password":"test123456",
   "PasswordConfirm":"test123456",
   "Photo":"testPhoto"
}
```

Reponse should be:

```output
{
    "data": {
        "user": {
            "id": "53143c8d-feed-4b42-baa4-2f14686470dd",
            "name": "test2",
            "email": "test2@test.com",
            "role": "user",
            "photo": "testPhoto",
            "provider": "local",
            "created_at": "2024-02-25T14:04:21.841391-05:00",
            "updated_at": "2024-02-25T14:04:21.841391-05:00"
        }
    },
    "status": "success"
}
```
