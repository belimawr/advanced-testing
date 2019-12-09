#!/bin/sh

#START OMIT
# First run works as expected
printf "The tests run normally\n" # OMIT
go test ./...

# What happens here?
printf "\nNow they're cached\n" # OMIT
go test ./...

# We can clean the test cache
printf "\nCleaning the test cache\n" # OMIT
go clean -testcache

# Now the tests run again
printf "\nThe tests run normally\n" # OMIT
go test ./...
#END OMIT
