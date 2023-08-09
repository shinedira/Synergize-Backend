# Backend - Technical Test
Pattern :
Clean Architecture (with lil touch of NestJs)

## Run
1. Copy `.env.example` file to `.env`
```bash
cp .env.example .env
update .env with your machine credential
```
2. `go mod tidy -v`
3. `go run .`

###
by default it will run on port `3000` just hit `http://127.0.0.1:3000`, base URL wil be `http://127.0.0.1:3000/api/v1`


## Endpoint API
| METHOD             | URL                                                                | Param | Auth |
| ----------------- | ------------------------------------------------------------------ | --------- | --------- |
| POST | {base-url}/auth/register | - | FALSE |
| POST | {base-url}/auth/login | - | FALSE |
| POST | {base-url}/auth/logout | - | TRUE |
| POST | {base-url}/auth/refresh | - | FALSE |
| GET | {base-url}/player/user | filters[username], filters[name], filters[bank][name], filters[bank][account_name], filters[bank][account_number] page, limit  | TRUE |
| GET| {base-url}/player/user/{id} | - | TRUE |
| POST| {base-url}/player/bank-account | - | TRUE |
| POST| {base-url}/player/transaction/topup | - | TRUE |


#### Payload POST Example
##### Register
```
{
    "name": "User 1",
    "username": "player",
    "password": "password"
}

```
##### Login
```
{
    "username": "player",
    "password": "password"
}

```
##### refresh
```
{
	"refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImMyMDQ4MWFjLTkwOTEtNDQ0Yi1iNzY5LWY1OWQzMzdlZTAwOCIsInVzZXJuYW1lIjoidCIsImlzc3VlZF9hdCI6IjIwMjMtMDgtMDlUMjA6MzM6MDQuNzg0OTIzKzA3OjAwIiwiZXhwaXJlZF9hdCI6IjIwMjQtMDItMDhUMDg6MzM6MDQuNzg0OTIzKzA3OjAwIn0.z-NH05Gjqv4ibmSq00OpOsERiNIiiey10hTdGkXGdrc"
}

```

##### Bank Account
```
{
    "bank_name": "Mandiri",
    "account_name": "Player Name",
    "account_number": 12345566
}

```
##### Topup
```
{
    "amount": 50000
}

```


## OR If you use Insomnia 
- you can import the collection, I already include the collection with file name `Insomnia.json`