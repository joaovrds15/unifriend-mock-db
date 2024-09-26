package builder

type UserResponseBuilder interface {
	createMajor()
	createUser()
	createQuiz()
	createQuestions()
	createOptions()
	generateUserResponses()
}

func GetBuilder() UserResponseBuilder {
	return newResponseBuilder()
}
