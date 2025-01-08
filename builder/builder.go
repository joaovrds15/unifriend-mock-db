package builder

type UserResponseBuilder interface {
	createMajor()
	createUser()
	createUserImages()
	createQuiz()
	createQuestions()
	createOptions()
	generateUserResponses()
}

func GetBuilder() UserResponseBuilder {
	return newResponseBuilder()
}
