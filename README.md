# services

## How to run:

1. Install dep:
```curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh ```
and install project dependencies by ```dep ensure```
2. Install docker and run the db:
```sudo docker run --name back2school -p 5432:5432 postgres```. Then you will be able to run the db by ```sudo docker start back2school```
3. Restore the dump to the database ```psql -U postgres -p 5432 -h localhost -f dump.sql```
4. Generate private and public key to sign tokens inside the src/config folder```openssl genrsa -out back2school.rsa 1024 && openssl rsa -in back2school.rsa -pubout > back2school.rsa.pub```
5. run the application by ```go run main.go``` and it will be reachable at "localhost:5000"

### TODO:
Docker-compose with multistage build of the go app and populated db (or maybe the db could also be populated by the app at startup).