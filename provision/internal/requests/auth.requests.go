package requests

type SignupRequest struct {
	FirstName string `json:"first_name" validate:"required,alpha"`
	LastName  string `json:"last_name" validate:"alpha"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=100"`
	Phone     string `json:"phone" validate:"required,numeric"`
	Role      string `json:"role" validate:"required,oneof=admin user"`
}
