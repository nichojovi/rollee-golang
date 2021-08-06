##### Rollee Project
```
Brought to you by Nicholas J.
Programming language: Go
Database: MySQL
How to run: make run
```

#### Brief about rollee project
```
On this repository
1. API was cover with authentication by username and password (standard encrypted SHA1)
2. Data is stored on DB (mysql), table provided
3. Sample unit test on repository level
4. Clean code architecture
```

#### 1. Get user by ID
```
curl --location --request GET 'localhost:9000/api/user' \
--header 'username: nichojovi' \
--header 'password: 448ed7416fce2cb66c285d182b1ba3df1e90016d' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "id": 1
}'
```

#### 2. Insert new user
```
curl --location --request POST 'localhost:9000/api/insert-user' \
--header 'username: nichojovi' \
--header 'password: 448ed7416fce2cb66c285d182b1ba3df1e90016d' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "username":"mtnike",
    "password":"c215eafb70a121a82858d5c2930fec30301503c3",
    "full_name":"mike lewis",
    "phone":"08727272727",
    "email":"mynameismike@gmail.com"
}'
```

#### 3. Update user's phone
```
curl --location --request PUT 'localhost:9000/api/update-user-phone' \
--header 'username: nichojovi' \
--header 'password: 448ed7416fce2cb66c285d182b1ba3df1e90016d' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "id":2,
    "phone":"0811111111111"
}'
```

#### 4. Delete user by id
```
curl --location --request DELETE 'localhost:9000/api/delete-user' \
--header 'username: nichojovi' \
--header 'password: 448ed7416fce2cb66c285d182b1ba3df1e90016d' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "id":2
}'
```

#### 5. Fibonacci
```
curl --location --request GET 'localhost:9000/api/fibonacci' \
--header 'username: nichojovi' \
--header 'password: 448ed7416fce2cb66c285d182b1ba3df1e90016d' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "n":1
}'
```