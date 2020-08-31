package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	database "github.com/AlekseyPromet/quiz/db"
	"github.com/AlekseyPromet/quiz/generated"
	"github.com/AlekseyPromet/quiz/model"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.QuestionInput) (*model.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to databse", err)
		return nil, err
	}
	defer db.Close()

	question := model.Questions{}
	question.QuestionText = input.QuestionText
	question.pubDate = input.PubDate

	db.Create(&question)
	return &question, nil
}

func (r *mutationResolver) CreateChoice(ctx context.Context, input *model.ChoiceInput) (*model.Choice, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Questions(ctx context.Context) ([]*model.Question, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Choices(ctx context.Context) ([]*model.Choice, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
