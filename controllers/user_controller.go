package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "mygolangapp/services"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := services.GetUserByID(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    response, err := json.Marshal(user)
    if err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    users, err := services.GetAllUsers()
    if err != nil {
        http.Error(w, "Error retrieving users", http.StatusInternalServerError)
        return
    }

    response, err := json.Marshal(users)
    if err != nil {
        http.Error(w, "Error encoding response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(response)
}
