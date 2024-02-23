package routes

//importing packages
import (
	models "GOTASK/models"
	"database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {

	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := models.Db.Exec("INSERT INTO tasks (title, description, status, duedate) VALUES (?, ?, ?, ?)", task.Title, task.Description, task.Status, task.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task Created Successfully"})

}

func GetAllTask(c *gin.Context) {

	rows, err := models.Db.Query("SELECT * FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)

}
func DeleteTask(c *gin.Context) {

	taskId := c.Param("id")

	result, err := models.Db.Exec("DELETE FROM tasks WHERE ID = ?", taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found with ID: " + taskId})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
func UpdateTask(c *gin.Context) {
	taskId := c.Param("id")

	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data in request body"})
		return
	}

	if updatedTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	result, err := models.Db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, duedate = ? WHERE id = ?",
		updatedTask.Title, updatedTask.Description, updatedTask.Status, updatedTask.DueDate, taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check affected rows"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error: no rows affected during update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})

}
func GetTask(c *gin.Context) {

	taskId := c.Param("id")

	var task models.Task
	err := models.Db.QueryRow("SELECT * FROM tasks WHERE ID = ?", taskId).Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found with ID: " + taskId})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, task)

}
