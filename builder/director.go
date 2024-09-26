package builder

type Director struct {
	builder UserResponseBuilder
}

func NewDirector(b UserResponseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) BuildUserResponse() {
	d.builder.createMajor()
	d.builder.createUser()
	d.builder.createQuiz()
	d.builder.createQuestions()
	d.builder.createOptions()
	d.builder.generateUserResponses()
}
