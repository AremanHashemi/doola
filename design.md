# Doola

## Data Schema

THe main entity is a user

Users can be created, read, updated, and deleted.

User
ID, CreatedAt, First Name
0, now, Cam

CREATE INDEX todos_user_id ON user_id;

SELECT * FROM todos WHERE user_id = 0; O(log(all todos))

SELECT * FROM todos WHERE created_at > '01-01-2020'; O(all todos)

CREATE INDEX todos_created_at ON created_at;

SELECT * FROM todos WHERE created_at > '01-01-2020'; O(log(all todos))

## Architecture

We will use an MVC architecture for our server.

Model == Database
View == API -- frontend will talk to the API
Controller == business logic, domain layer