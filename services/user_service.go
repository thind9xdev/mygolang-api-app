package services

import (
    "mygolangapp/models"
)

func GetUserByID(id int) (*models.User, error) {
    return &models.User{
        ID:    id,
        Name:  "John Doe",
        Email: "john@example.com",
    }, nil
}

func GetAllUsers() ([]models.User, error) {
    users := []models.User{
        {ID: 1, Name: "John Doe", Email: "john@example.com"},
        {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
    }
    return users, nil
}
