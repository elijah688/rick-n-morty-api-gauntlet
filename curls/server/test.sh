# curl "http://localhost:8080/character/list/episodes" \
#     -H "content-type: application/json" \
#     -d '{"ids":[1,2,3]}' | jq .

# curl "http://localhost:8080/location/list" \
# -H "content-type: application/json" \
# -d '{"ids":[1,2,3]}' | jq .

curl "http://localhost:8080/character/list/debut" \
    -H "content-type: application/json" \
    -d '{"ids":[1,2,3]}' | jq .
