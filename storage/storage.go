package storage

import (
	"fmt"
	"todocli/contract"
	"todocli/entity"
	"todocli/filehandler"
)

type Storage struct {
	handler    contract.StorageReaderWriter
	Users      []entity.User
	Categories []entity.Category
	Tasks      []entity.Task
}

func New(path string) (*Storage, error) {
	sHandler, hErr := filehandler.New(path)
	if hErr != nil {
		return nil, hErr
	}

	storage := &Storage{
		handler: sHandler,
	}
	lErr := storage.LoadData()
	if lErr != nil {
		return nil, lErr
	}

	return storage, nil
}

func (s *Storage) SaveData() error {
	uErr := s.SaveUsers()
	if uErr != nil {
		return fmt.Errorf("error while saving users: %w", uErr)
	}

	cErr := s.SaveCategories()
	if cErr != nil {
		return fmt.Errorf("error while saving categories: %w", cErr)
	}

	tErr := s.SaveTasks()
	if tErr != nil {
		return fmt.Errorf("error while saving tasks: %w", tErr)
	}

	return nil
}

func (s *Storage) LoadData() error {
	uErr := s.LoadUsers()
	if uErr != nil {
		return fmt.Errorf("error while loading data: %w", uErr)
	}

	cErr := s.LoadCategories()
	if cErr != nil {
		return fmt.Errorf("error while loading data: %w", cErr)
	}

	tErr := s.LoadTasks()
	if tErr != nil {
		return fmt.Errorf("error while loading data: %w", tErr)
	}

	return nil
}

func (s *Storage) LoadUsers() error {
	err := s.handler.Read("user", &s.Users)
	if err != nil {
		return fmt.Errorf("error while loading users: %w", err)
	}

	return nil
}

func (s *Storage) LoadCategories() error {
	err := s.handler.Read("category", &s.Categories)
	if err != nil {
		return fmt.Errorf("error while loading categories: %w", err)
	}

	return nil
}

func (s *Storage) LoadTasks() error {
	err := s.handler.Read("task", &s.Tasks)
	if err != nil {
		return fmt.Errorf("error while loading tasks: %w", err)
	}

	return nil
}

func (s *Storage) SaveUsers() error {
	err := s.handler.Write("user", s.Users)
	if err != nil {
		return fmt.Errorf("error while saving users: %w", err)
	}

	return nil
}

func (s *Storage) SaveCategories() error {
	err := s.handler.Write("category", s.Categories)
	if err != nil {
		return fmt.Errorf("error while saving categories: %w", err)
	}

	return nil
}

func (s *Storage) SaveTasks() error {
	err := s.handler.Write("task", s.Tasks)
	if err != nil {
		return fmt.Errorf("error while saving tasks: %w", err)
	}

	return nil
}

func (s *Storage) AddUser(u entity.User) {
	u.ID = s.GetLastUserID() + 1
	s.Users = append(s.Users, u)
}

func (s *Storage) AddCategory(c entity.Category) {
	c.ID = s.GetLastCategoryID() + 1
	s.Categories = append(s.Categories, c)
}

func (s *Storage) AddTask(t entity.Task) {
	t.ID = s.GetLastCategoryID() + 1
	s.Tasks = append(s.Tasks, t)
}

func (s *Storage) GetLastUserID() int64 {
	if len(s.Users) == 0 {
		return 0
	}

	return s.Users[len(s.Users)-1].ID
}

func (s *Storage) GetLastCategoryID() int64 {
	if len(s.Categories) == 0 {
		return 0
	}

	return s.Categories[len(s.Categories)-1].ID
}

func (s *Storage) GetLastTaskID() int64 {
	if len(s.Tasks) == 0 {
		return 0
	}

	return s.Tasks[len(s.Tasks)-1].ID
}
