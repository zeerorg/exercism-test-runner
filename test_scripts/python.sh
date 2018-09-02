#!/bin/bash
# script to run when testing

read download_path <<< $(exercism download --uuid=$1)
cd $download_path
pytest