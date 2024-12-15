#!/bin/sh

curl -X POST -v \
  -H "Content-Type: application/json" \
  -d @d.json \
  "http://localhost:8080/character"
