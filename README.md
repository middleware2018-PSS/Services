# Back2School REST API
[![Go Report Card](https://goreportcard.com/badge/github.com/middleware2018-PSS/Services)](https://goreportcard.com/report/github.com/middleware2018-PSS/Services)

## How to run
The Go application and the Postgres database run in two separate containers.
A Docker Compose service is provided for convenience.  
Additionally, all the commands you will need to run the service are defined in a Makefile,
if in doubt run `make` to list all the available command and their help

### Generate RSA keys used by JWT
**N.B.:** The following command should be run just once in order to generate the RSA key pair used by JWT to generate tokens.
Once the keys have been created you won't need to run this command anymore, unless you accidentally delete them.
The keys will be saved in `config/back2school.rsa{,.pub}`.
```
make gen-keys
```

### Build images and start the service
1. If you have just cloned the project, or you have made some changes to the code and you want to
build the images before starting the service, run `make build-up` otherwise `make up` (💄) should be good enough.
2. Once the service is started and the database has been properly initialized (might take some seconds), populate the
database with the initial entries running `make db-init`.  
The API will be reachable at http://localhost:5000
3. When you are done with your tests, run `make down` to stop and destroy the service


### TODO
Explain how to use the swagger doc to fiddle with the API

