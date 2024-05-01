#!/bin/bash

# Build Go files
go build quad_files/quadA.go
if [ $? -ne 0 ]; then
    echo "Build failed for quadA.go"
    exit 1
fi

go build quad_files/quadB.go
if [ $? -ne 0 ]; then
    echo "Build failed for quadB.go"
    exit 1
fi

go build quad_files/quadC.go
if [ $? -ne 0 ]; then
    echo "Build failed for quadC.go"
    exit 1
fi

go build quad_files/quadD.go
if [ $? -ne 0 ]; then
    echo "Build failed for quadD.go"
    exit 1
fi

go build quad_files/quadE.go
if [ $? -ne 0 ]; then
    echo "Build failed for quadE.go"
    exit 1
fi

go build quad_files/quadchecker.go
if [ $? -ne 0 ]; then
    echo "Build failed for quadchecker.go"
    exit 1
fi

# Echo success message
echo "Build completed successfully. Built files moved to the main directory."
