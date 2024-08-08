package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Task Manager API Project")

	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTask)
	router.DELETE("/tasks/:id", removeTask)
	router.PUT("/tasks/:id", updateTask)
	router.POST("/tasks", addTask)
	router.Run()
}

func getTasks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func getTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, val := range tasks {
		if val.ID == id {
			ctx.JSON(http.StatusOK, val)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func removeTask(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, val := range tasks {
		if val.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func updateTask(ctx *gin.Context) {
	id := ctx.Param("id")

	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func addTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

var tasks = []Task{
	{ID: "1", Title: "Task Manager Project", Description: "Add/View/Delete Tasks", DueDate: time.Now(), Status: "In Progress"},
	{ID: "2", Title: "Books Management Project", Description: "Add/View/Delete Books", DueDate: time.Now().AddDate(0, 0, -1), Status: "Completed"},
}




// package main

// import (
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()
	
// 	router.GET("/tasks", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
// 	})

// 	router.GET("/tasks/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")
	
// 		for _, task := range tasks {
// 			if task.ID == id {
// 				ctx.JSON(http.StatusOK, task)
// 				return
// 			}
// 		}
	
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// 	})

// 	router.PUT("/tasks/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")
	
// 		var updatedTask Task
	
// 		if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
	
// 		for i, task := range tasks {
// 			if task.ID == id {
// 				// Update only the specified fields
// 				if updatedTask.Title != "" {
// 					tasks[i].Title = updatedTask.Title
// 				}
// 				if updatedTask.Description != "" {
// 					tasks[i].Description = updatedTask.Description
// 				}
// 				ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
// 				return
// 			}
// 		}
	
// 		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
// 	})

// 	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
// 		id := ctx.Param("id")
	
// 		for i, val := range tasks {
// 			if val.ID == id {
// 				tasks = append(tasks[:i], tasks[i+1:]...)
// 				ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
// 				return
// 			}
// 		}
	
// 		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
// 	})

// 	router.POST("/tasks", func(ctx *gin.Context) {
// 		var newTask Task
	
// 		if err := ctx.ShouldBindJSON(&newTask); err != nil {
// 			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
	
// 		tasks = append(tasks, newTask)
// 		ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
// 	})

// 	router.Run() // Listen and serve on 0.0.0.0:8080
// }

// // Task represents a task with its properties.
// type Task struct {
// 	ID          string    `json:"id"`
// 	Title       string    `json:"title"`
// 	Description string    `json:"description"`
// 	DueDate     time.Time `json:"due_date"`
// 	Status      string    `json:"status"`
//    }
   
//    // Mock data for tasks
//    var tasks = []Task{
// 	   {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
// 	   {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
// 	   {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
//    }