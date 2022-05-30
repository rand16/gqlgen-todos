package graph

import "github.com/rand16/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos    []*model.Todo
	classes  []*model.ClassTeacher
	subjects []*model.SubjectTeacher
}
