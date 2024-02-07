package database

import(
	"github.com/gocql/gocql"
	"log"
	"fmt"
	"time"
)

var(
	Session *gocql.Session
	ErrNotFound = gocql.ErrNotFound
)

type Todo struct {
    ID          gocql.UUID `json:"id"`
    Username    string		`json:"username"`
    Title       string		`json:"title"`
    Description string		`json:"description"`
    Status      string		`json:"status"`
    CreatedAt   time.Time	`json:"created_at"`
    UpdatedAt   time.Time	`json:"updated_at"`
}

type TodoService interface {
    CreateTodo(username, title, description, status string, createdAt, updatedAt time.Time) (*Todo, error)
    GetTodo(id string) (*Todo, error)
    GetAllTodos(username string) ([]*Todo, error)
    UpdateTodo(id, newStatus string) (string, error)
    DeleteTodo(id string) (string, error)
}

func init(){
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Consistency = gocql.One
	cluster.Keyspace = "todo"

	var err error
	Session, err = cluster.CreateSession()
	if err!=nil{
		log.Fatal("Error connecting to ScyllaDB:",err)
	}else{
		fmt.Println("Connected to DB successfully")
	}

	// err = Session.Query("CREATE TABLE IF NOT EXISTS todo (id uuid PRIMARY KEY, username text UNIQUE, title text, description text, status text, created_at timestamp, updated_at timestamp);").Exec()
	err = Session.Query("CREATE TABLE IF NOT EXISTS todo (id uuid PRIMARY KEY, username text UNIQUE, title text, description text, status text, created_at timestamp, updated_at timestamp);").Exec()

	if err!=nil{
		log.Fatal("Error creating table:",err)
	}
}

func CreateTodo(username, title, description, status string, createdAt, updatedAt time.Time) (*Todo, error){
	id := gocql.TimeUUID()
	todo := &Todo{
		ID: id,
		Username: username,
		Title: title,
		Description: description,
		Status: status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	
	query := `INSERT INTO todo(id, username, title, description, status, created_at, updated_at) VALUES(?,?,?,?,?,?,?)`
	err := Session.Query(query, id, username, title, description, status, createdAt, updatedAt ).Exec()
	if err!=nil{
			return nil,err
		}
	
	return todo, nil
}

func GetTodo(id string) (*Todo, error){
	todo := &Todo{}

	query := `SELECT * FROM todo WHERE id = ?`
	err := Session.Query(query,id).Scan(&todo.ID, &todo.Username, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
	if err!=nil{
		return nil, err
	}
	return todo, nil
}

func GetAllTodos(username string) ([]*Todo, error){
	var todos []*Todo

	query := `SELECT * FROM todo WHERE username = ?`
	iter := Session.Query(query, username).Iter()
	
	var todo Todo
	for iter.Scan(&todo.ID, &todo.Username, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt){
		todos = append(todos, &todo)
	}

	if err := iter.Close(); err!=nil{
		return nil, err
	}
	return todos, nil
}

func UpdateTodo(id, newStatus string) (string, error){
	query := `UPDATE todo SET status = ? updated_at = ? WHERE id = ?`
	err := Session.Query(query, newStatus, time.Now(), id).Exec()
	if err!=nil{
		return "Error updating the status", err
	}
	return "Todo status updated successfully", nil
}

func DeleteTodo(id string)(string, error){
	query := `DELETE todo WHERE id = ?`
	err := Session.Query(query, id).Exec()
	if err!=nil{
		return "Error deleting the todo",err
	}
	return "Todo deleted successfully",nil
}