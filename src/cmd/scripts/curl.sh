# CREATE USER
curl http://localhost:8080/api/users -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -X POST -d '{ "name": "a", "email": "a@gmail.com", "password": "123456" }' -i

# GET USERS
curl http://localhost:8080/api/users -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -X GET -i
curl "http://localhost:8080/api/users?limit=10&last_id=1" -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -X GET

# GET USER BY ID
curl http://localhost:8080/api/users/1 -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -X GET -i

# UPDATE USER
curl http://localhost:8080/api/users -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -X PUT --data '{ "id": 1, "name": "a_1_update" }' -i

# DELETE USER
curl http://localhost:8080/api/users -H "X-API-Key: RAHASIA" -H "Content-Type: application/json" -X DELETE --data '{ "id": 1 }' -i
