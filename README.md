# go-auth
An authentication service using jwt built in go

## .env
In order to run the project create a ```.env``` file and add the following configuration:
```
PORT=<port to run the app>
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=<redis password>

ACCESS_SECRET=<the access secret>
REFRESH_SECRET=<the refresh secret>
```

## Run the app
Simply run ```docker-compose up --build```

## Test the routes
Test the following routes using postman or another tool:
- ```/login``` - body: {"username": "username", "password": "password}
- ```/logout``` - include the access token
- ```/todo``` - include the access token
- ```/refresh``` - body: {"refresh_token": <the refresh token>}

## Improvements
- This app is intended to test jwt implementation in go. So the users do not persist in a database. There is only one static user created in the code.
In order to make the app dynamic save the users in a database, include a registration route
- Unit Testing
