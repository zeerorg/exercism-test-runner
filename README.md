This is a server for running if a solution passes all the checks in exercism.io
I want to keep this server as simple as it can be. So all it needs to get is the language and uuid and returns if the result is true or false.

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
3. Visit `localhost:8000` to see if you get a `404 page not found` message.

## Next steps:
1. Setting up a test environment is the top priority.
2. Adding support for javascript. And then more languages.

Email : r.g.gupta@outlook.com for any questions or for contributions.