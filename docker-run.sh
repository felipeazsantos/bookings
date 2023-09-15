#!/bin/bash

docker run  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root --name postgres-bookings -p 5442:5432 -v /home/felipe/Documentos/Dev-pessoal/golang/go-build-modern-web-apps/bookings/data:/var/lib/postgresql/data  -d postgres