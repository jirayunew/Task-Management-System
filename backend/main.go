package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var jwtKey = []byte("your_secret_key")
var db *sqlx.DB

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func connectToDatabase() {
	var err error
	for i := 0; i < 10; i++ {
		db, err = sqlx.Connect("postgres", "host=db user=postgres password=secret dbname=task_management sslmode=disable")
		if err == nil {
			return
		}
		time.Sleep(2 * time.Second)
	}
	panic("Could not connect to database")
}

func main() {
	connectToDatabase()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", login)

	authorized := r.Group("/")
	authorized.Use(authenticate)
	{
		authorized.GET("/tasks", getTasks)
		authorized.POST("/tasks", addTask)
		authorized.PUT("/tasks/:id", updateTask)
		authorized.DELETE("/tasks/:id", deleteTask)
	}

	r.Run(":8080")
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if creds.Username != "admin" || creds.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}

// Task Struct ที่ปรับให้ใช้ *string 
type Task struct {
	ID          int        `json:"id" db:"id"`
	Name        string     `json:"name" db:"name"`
	Status      string     `json:"status" db:"status"`
	Description *string    `json:"description" db:"description"` // รองรับ nil
	DueDate     *string `json:"dueDate" db:"due_date"`        // รองรับ nil
	Priority    string     `json:"priority" db:"priority"`
}

// ดึงรายการ Task ทั้งหมด
func getTasks(c *gin.Context) {
	var tasks []Task
	err := db.Select(&tasks, "SELECT * FROM tasks")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, tasks)
}

// เพิ่ม Task ใหม่
func addTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(
		"INSERT INTO tasks (name, status, description, due_date, priority) VALUES ($1, $2, $3, $4, $5)",
		task.Name, task.Status, task.Description, task.DueDate, task.Priority,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Task added successfully"})
}

// อัปเดต Task
func updateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(
		"UPDATE tasks SET name = $1, status = $2, description = $3, due_date = $4, priority = $5 WHERE id = $6",
		task.Name, task.Status, task.Description, task.DueDate, task.Priority, id,
	)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task updated successfully"})
}

// ลบ Task
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
