# ps-tag-onboarding-go

## Running in Docker
1. Run the below command to spin up the app and mongo docker containers
    ```shell
    docker compose up --build -d
    ```
     
## Running locally
1. The app requires 2 environment variables to be set in order to run locally. NOTE: this step isn't required 
when running in docker since the environment variables are already set in docker-compose.override.yml file.
    ```
    APP_PORT=<port number the api runs on>
    MONGO_CONNECTION_STRING=<mongo db connection string>
    ```
2. Have a functional mongo db instance running, either locally or in the cloud. Make sure that the
`MONGO_CONNECTION_STRING` environment variable set in step 1 points to the mongo db instance.
3. Run the app using your favourite IDE or use the below command:
    ```shell
    cd cmd/user-api
    go run main.go
    ```

## CURL commands to hit the endpoints

Run the below curl commands to test the endpoints:
- Save a user
    ```shell
    curl localhost:8080/save -X POST \
      -H "Content-Type: application/json" \
      -d '{"id":"1","first_name":"John","last_name":"Doe","email":"john.doe@example.com","age":18}'
    ```
- Fetch a user
    ```shell
    curl localhost:8080/find/1
    ```