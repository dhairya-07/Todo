package database

import(
	"log"
	"fmt"
	// "encoding/binary"
	"time"

	// "github.com/google/uuid"
	"github.com/gocql/gocql"
)

var(
	Session *gocql.Session
	ErrNotFound = gocql.ErrNotFound
)

type Todo struct {
    ID          gocql.UUID `json:"id"`
    UserID    string		`json:"user_id"`
    Title       string		`json:"title"`
    Description string		`json:"description"`
    Status      string		`json:"status"`
    CreatedAt   time.Time	`json:"created_at"`
    UpdatedAt   time.Time	`json:"updated_at"`
}

type TodoService interface {
    CreateTodo(user_id,title, description, status string, createdAt, updatedAt time.Time) (*Todo, error)
    GetTodo(user_id, todoId string) (*Todo, error)
    GetAllTodos(user_id string) ([]*Todo, error)
    UpdateTodo(user_id,todoId, newStatus string) (string, error)
    DeleteTodo(user_id,todoId string) (string, error)
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

	query := `
        CREATE KEYSPACE IF NOT EXISTS todo
        WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}
    `
    Session.Query(query).Exec()

	err = Session.Query("CREATE TABLE IF NOT EXISTS todo (id uuid PRIMARY KEY, user_id text, title text, description text, status text, created_at timestamp, updated_at timestamp);").Exec()

	if err!=nil{
		log.Fatal("Error creating table:",err)
	}
}

func CreateTodo(user_id,title, description, status string, createdAt, updatedAt time.Time) (*Todo, error){
	// user_id := uuid.New().String()
	id := gocql.TimeUUID()
	todo := &Todo{
		ID: id,
		UserID: user_id,
		Title: title,
		Description: description,
		Status: status,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	
	query := `INSERT INTO todo(id, user_id, title, description, status, created_at, updated_at) VALUES(?,?,?,?,?,?,?)`
	err := Session.Query(query, id, user_id, title, description, status, createdAt, updatedAt ).Exec()
	if err!=nil{
			return nil,err
		}
	
	return todo, nil
}

func GetTodo(user_id, todoId string) (*Todo, error){
	todo := &Todo{}

	query := `SELECT * FROM todo WHERE user_id =? AND id = ? ALLOW FILTERING`
	err := Session.Query(query,user_id,todoId).Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt)
	if err!=nil{
		return nil, err
	}
	return todo, nil
}


// func GetAllTodos(user_id string) ([]*Todo, error) {
//     var todos []*Todo

//     query := `SELECT * FROM todo WHERE user_id = ? ALLOW FILTERING`
//     iter := Session.Query(query, user_id).Iter()

//     var todo Todo
// 	var createdAtTime, updatedAtTime time.Time
//     for iter.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &createdAtTime,&updatedAtTime) {
// 		todo.CreatedAt = createdAtTime
// 		todo.UpdatedAt = updatedAtTime
      
//         todos = append(todos, &todo)
//         todo = Todo{}
//     }

//     if err := iter.Close(); err != nil {
//         return nil, err
//     }
//     return todos, nil
// }

func GetAllTodos(user_id string) ([]*Todo, error) {
    var todos []*Todo

    query := `SELECT * FROM todo WHERE user_id = ? ALLOW FILTERING`
    iter := Session.Query(query, user_id).Iter()

    var todo Todo
    for iter.Scan(&todo.ID, &todo.UserID, &todo.Title, &todo.Description, &todo.Status, &todo.CreatedAt,&todo.UpdatedAt) {      
        todos = append(todos, &todo)
        todo = Todo{}
    }

    if err := iter.Close(); err != nil {
        return nil, err
    }
    return todos, nil
}


func UpdateTodo(user_id,todoId, newStatus string) (string, error){
	query := `UPDATE todo SET status = ?, updated_at = ? WHERE user_id = ? AND id = ? ALLOW FILTERING`
	err := Session.Query(query, newStatus, time.Now(), user_id, todoId).Exec()
	if err!=nil{
		return "Error updating the status", err
	}
	return "Todo status updated successfully", nil
}

func DeleteTodo(user_id, todoId string) (string, error){
	query := `DELETE FROM todo WHERE user_id = ? AND id = ? ALLOW FILTERING`
	err := Session.Query(query,user_id, todoId).Exec()
	if err!=nil{
		return "Error deleting the todo",err
	}
	return "Todo deleted successfully",nil
}