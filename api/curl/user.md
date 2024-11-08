## Register a new User
```sh
curl --location 'http://localhost:8888/register' \
--header 'Content-Type: application/json' \
--header 'Authorization: Basic YWRtaW46YWRtaW4=' \
--data '{
    "name":"Jo",
    "email":"toto",
		"address":{
			"city":"Strasbourg",
            "zip":"12345",
            "street":"la tour"
		},
		"games":[
			{"name":"Donjon Dragon"},
			{"name":"Warhammer"}
		],
    "pass":"toto"
}'
```

## Create a new User
```sh
curl --location 'http://localhost:8888/users' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data-raw '{
    "name":"Jo",
    "email":"biden@test.go",
		"address":{
			"city":"Strasbourg",
            "zip":"12345",
            "street":"la tour"
		},
		"games":[
			{"name":"Donjon Dragon"},
			{"name":"Warhammer"}
		],
    "pass":"zepass123"
}'
```

## Get a User

```sh
curl --location 'http://localhost:8888/users/{uuid}' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkX3VzZXIiOiJ1c3ItMm9aUXZPeWVkbXNLNzl3SHFZdWRyOVEyVmdJIiwiYWNjZXNzX2xldmVsIjoiIiwiZW1haWwiOiJ0b3RvIiwiZXhwIjoxNzMxMDc4NDcxfQ.oJzWfeidd5jE6sCUlEWJlVthCbwnxwHcdJTkLAVPGjI'
```

## Delete a User

```sh
curl --location --request DELETE 'http://localhost:8888/users/19b0a919-a182-48ba-bd33-5a5ceea68ed1' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkX3VzZXIiOiJ1c3ItMm9aUXZPeWVkbXNLNzl3SHFZdWRyOVEyVmdJIiwiYWNjZXNzX2xldmVsIjoiIiwiZW1haWwiOiJ0b3RvIiwiZXhwIjoxNzMxMDc4NDcxfQ.oJzWfeidd5jE6sCUlEWJlVthCbwnxwHcdJTkLAVPGjI'
```