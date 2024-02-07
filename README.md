# Todo
I have created the Todo API in GoLang with ScyllaDB as the database.<br>
I am running scyllaDB on docker container with the latest image.<br>
However there is an error regarding unmarshalling of timestamps returned from the scyllaDB into *string. Tried different solutions for it but haven't found a solution for it and the problem is still persisted. Here is the post link [Go Forum Post](https://forum.golangbridge.org/u/dhairya_srv/activity) Also there are errors stating syntax errors with gocql but when inspected couldn't find the syntax error. 
Instructions for running the API:<br>
 - Clone the repo into your local machine.
 - Run the command in your terminal: ```go get github.com/dhairya-07/Todo```.
 - Have Docker installed on your machine. Pull the scyllaDB image from docker registry with the command ```docker pull scylladb/scylla```<br>
 - Run the docker container for the scylladb instance using the command: ```docker run scylladb/scylla -d -p 9042:9042```<br>
 - In database/database.go at line 39 comment it. Then run ```go build``` and run ```./todo.exe```.
 - Then uncomment and again build and run the ```todo.exe```. First time it was done to create a keyspace in the database. Next time it is done to use that keyspace for the database.<br>
 
For now only the create todo handler is working fine rest have db related and syntax issues which I am working on.

Endpoints:
 - Create todo - ```http://localhost:9000/api/{userID}/todos``` POST REQUEST
   - Request body: JSON {"Title":"", "Description":""}<br>
 - Get All todos - ```http://localhost:9000/api/{userID}/todos``` GET REQUEST <br>
 - Get todo - ```http://localhost:9000/api/{userID}/todo/{todoID}``` GET REQUEST <br>
 - Update todo status - ```http://localhost:9000/api/{userID}/todo/{todoID}``` PUT REQUEST <br>
   - Request body: JSON {"Status":""}
 - Delete todo - ```http://localhost:9000/api/{userID}/todo/{todoID}``` DELETE REQUEST 
