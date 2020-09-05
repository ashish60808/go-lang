---
---
The goal of this test is to assert (to some degree) your coding, testing, automation and documentation skills. You're given a simple problem, so you can focus on showcasing your techniques.

## Problem definition

The aim of test is to create a simple HTTP service that stores and returns configurations that satisfy certain conditions.
Since we love automating things, the service should be automatically deployed to kubernetes.


## Instructions

1. Clone this repository.
2. Create a new `dev` branch.
3. Solve the task and commit your code. Commit often, we like to see small commits that build up to the end result of your test, instead of one final commit with all the code.
4. Do a pull request from the `dev` branch to the `master` branch. More on that right below.
5. Reply to the thread you are having with our HR department so we can start reviewing your code.


### Endpoints

Your application **MUST** conform to the following endpoint structure and return the HTTP status codes appropriate to each operation.

Following are the endpoints that should be implemented:

| Name   | Method      | URL
| ---    | ---         | ---
| List   | `GET`       | `/configs`
| Create | `POST`      | `/configs`
| Get    | `GET`       | `/configs/{name}`
| Update | `PUT/PATCH` | `/configs/{name}`
| Delete | `DELETE`    | `/configs/{name}`
| Query  | `GET`       | `/search?name={config_name}&data.{key}={value}`

#### Query

The query endpoint **MUST** return all configs that satisfy the query argument.

#### Schema

- **Config**
  - Name (string)
  - Data (key:value pairs)

### Configuration

Your application **MUST** serve the API on the port defined by the environment variable `SERVE_PORT`.
The application **MUST** fail if the environment variable is not defined.

### Deployment

The application **MUST** be deployable on a kubernetes cluster. Please provide manifest files and a script that deploys the application on a minikube cluster.
The application **MUST** be accessible from outside the minikube cluster.

## Rules

- You can use any language / framework / SDK of your choice.
- The API **MUST** return valid JSON and **MUST** follow the endpoints set out above.
- You **SHOULD** write testable code and demonstrate unit testing it.
- You can use any testing, mocking libraries provided that you state the reasoning and it's simple to install and run.
- You **SHOULD** document your code and scripts.

# Solution

I Have implement HTTP services using  the go lang.
 Code is resides in ApplicationCode folder.

#### golang code

<a href="ApplicationCode/src/app/assignment.go">ApplicationCode</a>

<a href="ApplicationCode/README.md">Go lang Code documentation</a>

#### Minikube deployment manifests files location

<a href="minikube-deployment/">manifests files</a>

<a href="minikube-deployment/README.md">minikube documentation</a>

##### Dockerfile location and documentation

<a href="./Dockerfile"> Dockerfile location</a>

###### Build the docker :

docker build -t go-docker:latest .

###### RUN the docker :
docker run --rm -e SERVE_PORT=8082 -it  -p 80:8082 go-docker:latest


#############################################################
# Technical specification
## API Methods: 

| Name        | Method      | URL                                                  |
| ---         | ---         | ---                                                  |
| List        | `GET`       | `/configs`                                           |
| Create      | `POST`      | `/configs`                                           |
| Get         | `GET`       | `/configs/{name}`                                    |
| Update      | `PUT/PATCH` | `/configs/{name}`                                    |
| Delete      | `DELETE`    | `/configs/{name}`                                    |
| Query       | `GET`       | `/search?name={config_name}&data.{key}={value}`      |
| HealthCheck | `GET`       | `/healthcheck`                                       |


### Request 1: | List   | `GET` | `/configs`
curl http://localhost:3000/configs

#### Expected Output:

[{"name":"config1","data":{"key":"val1","id":1}},{"name":"config2","data":{"key":"val2","id":2}}]

### Request 2: | Create | `POST` | `/configs`
curl -X POST -d '{"name":"config1", "data": {"key": "val1", "id": 1} }' -H "Content-Type: application/json" http://localhost:3000/configs

#### Expected Output:

{"name":"config1","data":{"key":"val1","id":1}}

### Request 3: | Get    | `GET` | `/configs/{name}`
curl http://localhost:3000/configs/config2

#### Expected Output:

{"name":"config2","data":{"key":"val2","id":2}}

### Request 4: | Delete | `DELETE` | `/configs/{name}`
curl  -X DELETE http://localhost:3000/configs/config2

### Request 5: | Query  | `GET`  | `/search?name={config_name}&data.{key}={value}`
curl http://localhost:3000/search\?name\=config1\&data.id\=1

#### Expected Output:

{"name":"config1","data":{"key":"val1","id":1}}

### Request 6: | HealthCheck  | `GET`  | `/healthcheck`
curl http://localhost:3000/healthcheck

#### Expected Output:

Application is up and running

### Unit Test cases

assignment_test.go file contains the unit test cases implementation  for above REST APIs
 
To run the test cases run the "go test -v" , results will be displayed on the console
#### Expected Output: 

=== RUN   TestHealthcheckEndpoint

--- PASS: TestHealthcheckEndpoint (0.00s)

=== RUN   TestListConfigsEndpoint

--- PASS: TestListConfigsEndpoint (0.00s)

=== RUN   TestGetConfigEndpoint

--- PASS: TestGetConfigEndpoint (0.00s)

=== RUN   TestCreateConfigEndpoint

--- PASS: TestCreateConfigEndpoint (0.00s)

=== RUN   TestDeleteConfigEndpoint

--- PASS: TestDeleteConfigEndpoint (0.00s)

=== RUN   TestSearchConfigEndpoint

--- PASS: TestSearchConfigEndpoint (0.00s)

PASS

