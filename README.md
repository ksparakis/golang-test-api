# Test API
I built this api using Gin a golang api framework and MYSQL as a database.

To Deploy:
- pull this github repo down locally
- In a terminal, go to the directory where this project is located locally and run:

docker-compose build
docker-compose up

To run GO locally 
go mod download


API Routes
POST
/post-data/{name of the data}

GET
/get-data/{name of the data}