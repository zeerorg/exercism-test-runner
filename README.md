This is a server for running if a solution passes all the checks in exercism.io
I want to keep this server as simple as it can be. So all it needs to get is the language and uuid and returns if the result is true or false.

# API:
A single GET request to `/test/{language}/{uuid}` will run the docker container for respective language and return the result.

## Current plan:
1. Take a solution uuid and return true if it passes the tests else return nothing
2. How to do this
    1. Create dockerfile for language
    2. Provide the docker container with uuid and download solution in the container
    3. Run the container and see its output

## Dependencies:
For both windows and linux `go` and `docker` are the only dependencies which are needed for project to run.

## How to build:
1. Run `build_test_images.sh` it essential builds all images for test environment.
2. `go build` and `go run` next should run the server.
3. Visit `localhost:8000` to see if you get a `Exercism test server is running` message.
4. Be sure to run the server from a user who has access to `docker` (root or in the docker group).

## Writing Tests
1. Take help from `handler_test.go` file. Where python and welcome page tests are written.

## Next steps:
1. Adding support for javascript. And then more languages.
2. Creating a frontend for this server. Web extension didn't work since I needed to send a request to another domain and firefox considers that a CORS violation. Input on this would be great. 

Email : r.g.gupta@outlook.com for any questions or for contributions.