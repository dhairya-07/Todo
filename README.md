# Todo
I have created the Todo API in GoLang with ScyllaDB as the database.<br>
I am running scyllaDB on docker container with the latest image.<br>
However there is an error while creating the table of the Todos. I have searched on various places for the fix but nothing worked. At last I have posted on the scyllaDB forum for assistance.
When it will be resolved I will update the repo with the changes ASAP. Here is the post link [ScyllaDB Forum Post](https://forum.scylladb.com/u/dhairya_srv/activity)
Instructions for running the API:<br>
 - Clone the repo into your local machine.
 - Run the command in your terminal: ```go get github.com/dhairya-07/Todo```.
 - Have Docker installed on your machine. Pull the scyllaDB image from docker registry with the command ```docker pull scylladb/scylla```<br>
 - Run the docker container for the scylladb instance using the command: ```docker run scylladb/scylla -d -p 9042:9042```<br>
 - After that run ```go build```. This will build the code and generate a todo.exe. Run the exe: ```./todo.exe```
For now there will be an error for the db table creation I am working on that will update soon.

