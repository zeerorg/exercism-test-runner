This is a server for running if a solution passes all the checks in exercism.io

Current plan:
1. Take a solution uuid and return true if it passes the tests else return nothing
2. How to do this
    1. Create dockerfile for language
    2. Provide the docker container with uuid and download solution in the container
        1. 
    3. Run the container and see its output

Another way to do things:
1. Download the solution on the server using exercism api
