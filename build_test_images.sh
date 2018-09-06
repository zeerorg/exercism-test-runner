#!/usr/bin/env sh

docker build -t exercism_build dockerfiles/exercism_build
docker build -t python_test dockerfiles/python_test_environment
