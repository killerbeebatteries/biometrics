# Biometrics

Tracking blood pressure and weight. 💪❤️

# TODO

- Implement HTMX for adding, updating and deleting records.
- Implement pretty graphs for weight and blood pressure
- Implement hosting both db and webapp in docker containers using docker-compose
- Document dev and hosting steps
- Annotations for comments

# Development

## Setup

Create a .env file and place it in the `golang_server` directory:
```
TARGET=dev
POSTGRES_PASSWORD=maybemakethisrandom

DB_USER="biometrics_user"
DB_NAME="biometrics"
DB_PASSWORD="lotsoffunwithgolang"
DB_HOST="localhost"
DB_PORT="5432"
WEB_PORT="8000"
```

The `TARGET` reflects the stage that will be built in the `Dockerfile` referenced in our `docker-compose.yaml`.

Example:
```
FROM base as dev
```
Docker Compose would use that stage of the Dockerfile if the target was set to `dev`.

## Run Environment

Using the magic of [docker-compose and air](https://firehydrant.com/blog/develop-a-go-app-with-docker-compose/), we can run a hot-reloadable environment that will run our app on whatever port we declared in `WEB_PORT`.

```
docker-compose up
```

# Production

## Setup

Create a .env file and place it in the `golang_server` directory:
```
TARGET=production
POSTGRES_PASSWORD=differenttodevhopefully

DB_USER="biometrics_user"
DB_NAME="biometrics"
DB_PASSWORD="nothingtoseehere"
DB_HOST="localhost"
DB_PORT="5432"
WEB_PORT="8000"
``

## Run Environment

```
docker-compose -d
```
