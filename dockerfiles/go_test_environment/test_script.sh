#!/bin/bash
# script to run when testing

cd $(exercism download --uuid=$1)
# cd $download_path
go test
