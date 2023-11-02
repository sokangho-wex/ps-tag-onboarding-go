package users

type User struct {
	ID        string `json:"id" bson:"_id"`
	FirstName string `json:"first_name" bson:"firstname"`
	LastName  string `json:"last_name" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Age       int8   `json:"age" bson:"age"`
}

func NewUser(id string, fn string, ln string, email string, age int8) User {
	return User{id, fn, ln, email, age}
}
