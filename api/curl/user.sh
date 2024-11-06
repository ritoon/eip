curl --location 'http://localhost:8888/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46YWRtaW4=' \
--data-raw '{
    "name":"Jo",
    "email":"biden@test.go",
    "pass":"zepass123"
}'

curl --location 'http://localhost:8888/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"biden@test.go",
    "pass":"zepass123"
}'

curl --location 'http://localhost:8888/users/77416ded-75fd-476d-9b9b-000193fe8967'


curl --location --request DELETE 'http://localhost:8888/users/77416ded-75fd-476d-9b9b-000193fe8967' \
--header 'Authorization: Bearer null'