package dtos

type CreateUserRequest struct {
	Fname    string `json:"fname" label:"First name"  validate:"required,min=2,max=100"`
	Lname    string `json:"lname" label:"Last name"   validate:"required,min=2,max=100"`
	Age      int32  `json:"age"   label:"Age"         validate:"required,gte=18,lte=120"`
	Email    string `json:"email" label:"Email"       validate:"required,email"`
	Password string `json:"password" label:"Password" validate:"required,min=8,max=64"`
}