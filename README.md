# IITK Coin

This repository contains my implementation of backend of IITK Coin pclub snt summer project.

## Index
 - Directory Structure 
 - Database 
 - Endpoint Details
 
 ## Directory Structure
 
    ├── Dockerfile
    ├── README.md
    ├── controllers
    │   ├── auth.go
    │   ├── coins.go
    │   ├── db.go
    │   ├── item.go
    │   ├── middleware.go
    │   ├── otp.go
    │   ├── redeem.go
    │   ├── router.go
    │   ├── transaction.go
    │   ├── user.go
    │   ├── utils.go
    │   └── validation.go
    ├── docker-compose.yml
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── models
    │   ├── hash.go
    │   ├── item.go
    │   ├── params.go
    │   ├── request.go
    │   ├── token.go
    │   └── user.go
    ├── sqldb
    │   └── db.go
    └── utils
        ├── mail.go
        └── otp.go
    
    4 directories, 27 files
## Database: Sqlite
My database implementation containe five tables

- ### users

	    CREATE TABLE users (
	                    rollNo INTEGER PRIMARY KEY,
	                    name TEXT NOT NULL,
	                    password TEXT NOT NULL,
	                    coins TEXT NOT NULL,
	                    isAdmin INTEGER NOT NULL,
	                    isFreezed INTEGER NOT NULL
	            );

- ### transactions

		CREATE TABLE transactions (
		                id INTEGER PRIMARY KEY AUTOINCREMENT,
		                sender INTEGER REFERENCES users(rollNo),
		                reciever INTEGET NOT NULL REFERENCES users(rollNo),
		                amount INTEGER NOT NULL,
		                type TEXT NOT NULL,
		                madeAt INTEGER NOT NULL
		        );

- ### redeem_requests

	    CREATE TABLE redeem_requests (
	                    id INTEGER PRIMARY KEY AUTOINCREMENT,
	                    user INTEGER NOT NULL REFERENCES users(rollNo),
	                    itemCode INTEGER NOT NULL REFERENCES items(code),
	                    status TEXT NOT NULL,
	                    madeAt TEXT NOT NULL
	            );

- ### items

	    CREATE TABLE items (
	                    code INTEGER PRIMARY KEY AUTOINCREMENT,
	                    amount INTEGER NOT NULL,
	                    name TEXT NOT NULL,
	                    isAvailable INTEGER NOT NULL
	            );

- ### otps

	    CREATE TABLE otps (
	                    otp INTEGER NOT NULL,
	                    user INTEGER NOT NULL REFERENCES users(rollNo) PRIMARY KEY,
	                    madeAt TEXT NOT NULL
	            );
## Details of Endpoints

### Signup

    url: /user/signup
    method: POST
    
    Request: Body {
    	"rollNo": "",
    	"name": "",
    	"password": "",
    	"isAdmin": "",
    	"isFreezed": ""
    }

### Login

    url: /user/login
    method: POST
    
    Request: Body {
    	"rollNo": "",
    	"password": ""
    }

### Transfer Coins

    url: /coins/send
    method: POST
    
    Request: Body {
    	"rollNo": "",
    	"coins": ""
    }
	
	-> rollNo here refers to reciever's roll no.

### Reward Coins

    url: /coins/reward
    method: POST
    
    Request: Body {
    	"rollNo": "",
    	"coins": ""
    }
	
	-> rollNo here refers to reciever's roll no.
	Note:- This is an admin only route.

### Get coins

    url: /coins/balance
    method: GET
	
### Item List

    url: /items
    method: GET
    
### Make Redeem Request

    url: /coins/redeem
    method: POST
    
    Request: Body {
    	"itemCode: ""
    }

### Get Redeem Requests

    url: /redeemRequests
    method: GET
    
    Note:- This is an admin only route.

### Update Redeem Request Status

    url: /redeemRequests
    method: POST
	
	Request: Body {
    	"id": "",
    	"status": ""
    }
    
    Note:- This is an admin only route.

### Get Otp

    url: /otp
    method: GET

### Check Otp

    url: /otp
    method: POST
    
    Request: Body {
    	"otp": "",
    }

### Add Item

    url: /items
    method: POST
    
    Request: Body {
    	"amount": "",
    	"name": "",
    	"isAvailable": ""
    }
	
	Note:- This is an admin only route.

### Update Item

    url: /items/:itemcode
    method: PUT
    
    Request: Body {
    	"amount": "", //optional
    	"name": "",  //optional
    	"isAvailable": "" //optional
    }
	
	Note:- This is an admin only route.

### Delete Item

    url: /items/:itemcode
    method: DELETE
	
	Note:- This is an admin only route.
