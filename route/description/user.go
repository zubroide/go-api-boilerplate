package description

import (
	"github.com/zubroide/go-api-boilerplate/model/entity"
	"github.com/zubroide/go-api-boilerplate/model/repository"
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
	Body *entity.UserFields
}

// swagger:parameters GetUsers
type UsersParameters struct {
	// in: query
	*repository.UserListParameters
}
