package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/rand16/gqlgen-todos/graph/generated"
	"github.com/rand16/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//panic(fmt.Errorf("not implemented"))
	todo := &model.Todo{ID: fmt.Sprintf("T%d", rand.Intn(100)), Text: input.Text, UserID: &input.UserID}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateClass(ctx context.Context, name string, class string) (*model.ClassTeacher, error) {
	//panic(fmt.Errorf("not implemented"))
	newClass := &model.ClassTeacher{Name: name, Class: class}
	r.classes = append(r.classes, newClass)
	return newClass, nil
}

func (r *mutationResolver) CreateSubject(ctx context.Context, name string, subject string) (*model.SubjectTeacher, error) {
	//panic(fmt.Errorf("not implemented"))
	newSubject := &model.SubjectTeacher{Name: name, Subject: subject}
	r.subjects = append(r.subjects, newSubject)
	return newSubject, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	//panic(fmt.Errorf("not implemented"))
	return r.todos, nil
}

func (r *queryResolver) Classes(ctx context.Context) ([]*model.ClassTeacher, error) {
	//panic(fmt.Errorf("not implemented"))
	return r.classes, nil
}

func (r *queryResolver) Subjects(ctx context.Context) ([]*model.SubjectTeacher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Teacher(ctx context.Context) ([]model.Teacher, error) {
	//panic(fmt.Errorf("not implemented"))
	var teachers []model.Teacher
	index1 := len(r.classes)
	index2 := len(r.subjects)
	index := index1 + index2
	for i := 0; i < index; i++ {
		if i < index1 {
			teachers = append(teachers, r.classes[i])
		} else {
			teachers = append(teachers, r.subjects[i-index1])
		}
	}
	return teachers, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	//panic(fmt.Errorf("not implemented"))
	user := &model.User{ID: *obj.UserID, Name: "user" + *obj.UserID}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
