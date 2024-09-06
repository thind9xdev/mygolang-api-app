package models



type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func GetUserByID(id int) (*User, error) {

    return &User{
        ID:    id,
        Name:  "John Doe",
        Email: "john@example.com",
    }, nil
}
func GetAllUsers() ([]User, error) {
 
    users := []User{
        {ID: 1, Name: "John Doe", Email: "john@example.com"},
        {ID: 2, Name: "Jane Doe", Email: "jane@example.com"},
    }
    return users, nil
}