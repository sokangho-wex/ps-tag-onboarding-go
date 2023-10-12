# ps-tag-onboarding-go

## Setup
1. Run the below command to spin up the app and mongo docker containers
    ```shell
    docker compose up --build -d
    ```
2. Run the below curl commands to test the endpoints:
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