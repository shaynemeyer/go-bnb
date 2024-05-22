#!/bin/bash

go build -o go-bnb cmd/web/*.go && ./go-bnb -dbname=bookings -dbuser=shaynemeyer -cache=false -production=false
