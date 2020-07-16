package domain

// Note the use of the slanted slash before the json keyword which causes Go to consider this as json out-of-the-box
type User struct {
	Id 			uint64 `json:"id"`
	FirstName 	string `json:"first_name"`
	LastName 	string `json:"last_name"`
	Email 		string `json:"email"`
}
