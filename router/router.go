package router

import(
	"github.com/dhairya-07/todo/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	// router.HandleFunc("/api/todos", middleware.CreateTodoHandler).Methods("POST","OPTIONS")
	// router.HandleFunc("/api/{username}/todos", middleware.GetAllTodosHandler).Methods("GET","OPTIONS")
	// router.HandleFunc("api/todo/{id}", middleware.GetTodoHandler).Methods("GET","OPTIONS")
	// router.HandleFunc("api/todo/{id}", middleware.UpdateTodoHandler).Methods("PUT","OPTIONS")
	// router.HandleFunc("api/todo/{id}", middleware.DeleteTodoHandler).Methods("DELETE","OPTIONS")

	todoService := middleware.NewTodoService()

	router.HandleFunc("/api/{user_id}/todos", todoService.CreateTodoHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/{user_id}/todos", todoService.GetAllTodosHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/{user_id}/todo/{id}", todoService.GetTodoHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/{user_id}/todo/{id}", todoService.UpdateTodoHandler).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/{user_id}/todo/{id}", todoService.DeleteTodoHandler).Methods("DELETE", "OPTIONS")

	return router
}