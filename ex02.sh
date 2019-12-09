#!/bin/sh

# First run works as expected
#START OMIT
printf "The tests always run\n" # OMIT
go test ./... -count=1
#END OMIT

printf "The tests always run\n" # OMIT
go test ./... -count=1

printf "The tests always run\n" # OMIT
go test ./... -count=1
