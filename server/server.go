package server

import (
	"fmt"
	"time"
	"todocli/entity"
	"todocli/storage"
)

type Server struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) (*Server, error) {
	return &Server{
		storage: storage,
	}, nil
}

func (s *Server) GetCategoryTasks(categoryID int64) []entity.Task {
	tasks := make([]entity.Task, 0)
	for _, task := range s.storage.Tasks {
		if task.CategoryID == categoryID {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func (s *Server) RegisterUser(username, password, firstName, lastName string) error {
	user := &entity.User{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}
	user.SetPassword(password)

	s.storage.AddUser(*user)

	return nil
}

func (s *Server) LoginUser(username, password string) *entity.User {
	for _, user := range s.storage.Users {
		if user.Username == username && user.CheckPassword(password) {
			return &user
		}
	}

	return nil
}

func (s *Server) CreateCategory(title, description, color string, user *entity.User) error {
	if user == nil {
		return fmt.Errorf("error while creating category: user is not authenticated")
	}

	category := entity.Category{
		Title:       title,
		Description: description,
		Color:       color,
		UserID:      user.ID,
	}
	s.storage.AddCategory(category)

	return nil
}

func (s *Server) CreateTask(
	title string,
	description string,
	Status entity.Status,
	deadline time.Time,
	categoryID int64,
) {
	task := entity.Task{
		Title:       title,
		Description: description,
		Status:      Status,
		Deadline:    deadline,
		CreatedAt:   time.Now(),
		CategoryID:  categoryID,
	}

	s.storage.AddTask(task)
}

func (s *Server) GetUserCategories(user *entity.User) ([]entity.Category, error) {
	if user == nil {
		return nil, fmt.Errorf("error while getting user categories: user is not authenticated")
	}
	categories := make([]entity.Category, 0)
	for _, category := range s.storage.Categories {
		if category.UserID == user.ID {
			categories = append(categories, category)
		}
	}

	return categories, nil
}

func (s *Server) GetUserTasks(user *entity.User) ([]entity.Task, map[int64]string) {
	categoryTitles := make(map[int64]string)
	for _, category := range s.storage.Categories {
		if category.UserID == user.ID {
			categoryTitles[category.ID] = category.Title
		}
	}

	tasks := make([]entity.Task, 0)
	for _, task := range s.storage.Tasks {
		if _, ok := categoryTitles[task.CategoryID]; ok {
			tasks = append(tasks, task)
		}
	}

	return tasks, categoryTitles
}
