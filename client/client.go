package client

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todocli/entity"
	"todocli/helper"
	"todocli/server"
)

type Client struct {
	scanner           bufio.Scanner
	server            *server.Server
	authenticatedUser *entity.User
}

func New(s *server.Server) (*Client, error) {
	return &Client{
		scanner:           *bufio.NewScanner(os.Stdin),
		server:            s,
		authenticatedUser: nil,
	}, nil
}

func (c *Client) HandleCommand() {
	fmt.Printf("Enter a command: ")
	c.scanner.Scan()
	switch strings.ToLower(c.scanner.Text()) {
	case "exit":
		c.HandleExit()
	case "register-user":
		c.HandleRegisterUser()
	case "login-user":
		c.HandleLoginUser()
	case "create-category":
		c.HandleCreateCategory()
	case "create-task":
		c.HandleCreateTask()
	case "list-tasks":
		c.HandleGetUserTasks()
	case "list-category-tasks":
		c.HandleGetCategoryTasks()
	default:
		fmt.Println("Invalid Command")
	}
}

func (c *Client) HandleGetUserTasks() {
	if c.authenticatedUser == nil {
		fmt.Println("You should login first to see list of your tasks")
		return
	}

	tasks, categoriesTitle := c.server.GetUserTasks(c.authenticatedUser)
	for i, task := range tasks {
		fmt.Printf("Task#%d- Title: %s, Description: %s, Status: %s, Deadline: %s, Category: %s\n",
			i, task.Title, task.Description, task.Status.GetTitle(), task.Deadline.Format("2006-01-02 15:04:05"), categoriesTitle[task.CategoryID])
	}
}

func (c *Client) HandleCreateTask() {
	if c.authenticatedUser == nil {
		fmt.Println("You should login first to create a task")
		return
	}

	categories, err := c.server.GetUserCategories(c.authenticatedUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	title := helper.ReadString(c.scanner, "Enter title of task: ")
	description := helper.ReadString(c.scanner, "Enter description of task: ")
	status := entity.Waiting
	deadline := helper.ReadTime(c.scanner, "Enter deadline of task (format: 2001-01-25 12:12:12): ")
	categoriesTitle := make([]string, len(categories))
	categoriesID := make([]int, len(categories))
	for i, category := range categories {
		categoriesTitle[i] = category.Title
		categoriesID[i] = int(category.ID)
	}
	categoryID := helper.ReadIntChoice(c.scanner, "Select category of task: ", categoriesTitle, categoriesID)

	c.server.CreateTask(title, description, status, deadline, int64(categoryID))
}

func (c *Client) HandleCreateCategory() {
	if c.authenticatedUser == nil {
		fmt.Println("You should login first to create a category")
		return
	}

	title := helper.ReadString(c.scanner, "Enter title of category: ")
	description := helper.ReadString(c.scanner, "Enter description of category: ")
	color := helper.ReadString(c.scanner, "Enter color of category: ")

	c.server.CreateCategory(title, description, color, c.authenticatedUser)

	fmt.Println("Category created successfully")
}

func (c *Client) HandleRegisterUser() {
	username := helper.ReadString(c.scanner, "Enter your username: ")
	password := helper.ReadString(c.scanner, "Enter your password: ")
	firstName := helper.ReadString(c.scanner, "Enter your first name: ")
	lastName := helper.ReadString(c.scanner, "Enter your last name: ")

	err := c.server.RegisterUser(username, password, firstName, lastName)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) HandleExit() {
	os.Exit(0)
}

func (c *Client) HandleLoginUser() {
	username := helper.ReadString(c.scanner, "Enter your username: ")
	password := helper.ReadString(c.scanner, "Enter your password: ")

	c.authenticatedUser = c.server.LoginUser(username, password)
	if c.authenticatedUser == nil {
		fmt.Println("Invalid username or password")
	} else {
		fmt.Println("Logged in successfully")
	}
}

func (c *Client) HandleGetCategoryTasks() {
	categories, err := c.server.GetUserCategories(c.authenticatedUser)
	if err != nil {
		fmt.Println(err)
		return
	}

	categoriesTitle := make([]string, len(categories))
	categoriesIDs := make([]int, len(categories))
	for i, category := range categories {
		categoriesTitle[i] = category.Title
		categoriesIDs[i] = int(category.ID)
	}
	categoryID := helper.ReadIntChoice(c.scanner, "Select your category: ", categoriesTitle, categoriesIDs)

	tasks := c.server.GetCategoryTasks(int64(categoryID))
	for i, task := range tasks {
		fmt.Printf("Task#%d- Title: %s, Description: %s, Status: %s, Deadline: %s\n",
			i, task.Title, task.Description, task.Status.GetTitle(), task.Deadline.Format("2006-01-02 15:04:05"))
	}
}
