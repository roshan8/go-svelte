#!/bin/bash

cd fe
npm install
npm run build
cd ..
go run main.go
