#!/bin/bash

go build -o bookings cmd/web/*.go && 
./bookings -dbname=bookings -dbuser=root -dbpass=root -dbport=5442 -cache=false -production=false 