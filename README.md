It's a small service that can shorten urls, save them inside Postgresql database and retrieve them from it.
It's listening at 8080 port and has only one entrypoint: "/".
To test its functionality do the following:
  - Run `go run main.go` in project directory to test with your local postgres database,
  OR
  - Run `docker compose up --build` in project directory to test within docker containers.

Note that first way require postgresql installed on your machine.

It also requires database with "url_storage" name and user with "go_user" and "8246go" name and password accordingly.

To get this via psql CLI type the following (without postgres=#):

   **postgres=#** `CREATE ROLE go_user WITH LOGIN PASSWORD '8246go' CREATEDB;`   
   **postgres=#** `CREATE DATABASE url_storage WITH OWNER = go_user;`

  It assumed postgresql is listening on 5432 port.
  
  The second way requires docker installed on your machine (note that I use `docker compose` not `docker-compose`).
  
After service run successfully you can send GET and POST requests (using `curl`, for example) with urls in their body:

  ```curl -L -X GET 'localhost:8080/' -H 'Content-Type: application/json' --data-raw '{"url_short": "GkNk7Plccu"}'```
			
  OR
  
  ```curl -L -X POST 'localhost:8080/' -H 'Content-Type: application/json' --data-raw '{"url_long": "https://www.github.com/lesnoyleshyi"}'```


It'll be more handy to use `make` utility by editing *Makefile* and running `make get` or `make post`.
  
