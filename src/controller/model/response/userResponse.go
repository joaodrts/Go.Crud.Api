package response

type UserRequest struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
}
