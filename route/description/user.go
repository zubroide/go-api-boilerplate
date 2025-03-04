package description

import (
	"github.com/zubroide/go-api-boilerplate/controller"
	"github.com/zubroide/go-api-boilerplate/model/entity"
)

// User data
// swagger:response UserResponse
type UserResponse struct {
	// in: body
	Body struct {
		Entity *entity.User
		*BaseResponseBody
	}
}

// Users data
// swagger:response UsersResponse
type UsersResponse struct {
	// in: body
	Body struct {
		Entities []*entity.User
		*BaseResponseBody
	}
}

// swagger:parameters CreateUser UpdateUser
type UserParameters struct {
	// in: body
	Body *entity.User
}

// swagger:parameters GetUsers
type UsersParameters struct {
	// in: query
	*controller.UserListParameters
}
