package builder

import (
	"unifriends-db-script/factories"
	"unifriends-db-script/models"
)

type Builder struct {
	Major        models.Major
	User         []models.User
	Quiz         models.QuizTable
	Questions    []models.QuestionTable
	Options      []models.OptionTable
	UserReponses []models.UserResponse
}

func newResponseBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) createMajor() {
	b.Major = factories.CreateMajor()
}

func (b *Builder) createUser() {
	users := factories.CreateUser(50, b.Major)
	b.User = append(b.User, users...)
}

func (b *Builder) createQuiz() {
	b.Quiz = factories.CreateQuiz()
}

func (b *Builder) createQuestions() {
	b.Questions = factories.CreateQuestions(5, b.Quiz.Id)
}

func (b *Builder) createOptions() {
	for _, question := range b.Questions {
		optionsFactory := factories.CreateOptions(5, question.Id)
		b.Options = append(b.Options, optionsFactory...)
	}
}

func (b *Builder) generateUserResponses() {
	for _, user := range b.User {
		for _, question := range b.Questions {
			randomOptionId := models.GenerateRandomOption(question.Id)
			userResponse := models.UserResponse{
				QuestionID: question.Id,
				OptionID:   randomOptionId,
				UserID:     user.Id,
			}

			userResponse.InsertUserResponse()
			b.UserReponses = append(b.UserReponses, userResponse)
		}
	}
}
