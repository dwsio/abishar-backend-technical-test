﻿# abishar-backend-technical-test

Steps to run:
1. Clone this repo 
2. Copy .env.example file to .env
3. Create database in postgreSQL with name `abishar-transactions`
4. Create table inside the database with name `transactions`
5. Go to inside the repo folder and run `make migrate-db`
6. Then run `make run`
7. Open browser and go to `http://127.0.0.1:3124/api/transaction/docs/#`
