package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (e *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@") {
		panic("Email should contain @")
	}
	e.email.from = from
	return e
}

func (e *EmailBuilder) To(to string) *EmailBuilder {
	if !strings.Contains(to, "@") {
		panic("Email should contain @")
	}
	e.email.to = to
	return e
}

func (e *EmailBuilder) Subject(subject string) *EmailBuilder {
	e.email.subject = subject
	return e
}

func (e *EmailBuilder) Body(body string) *EmailBuilder {
	e.email.body = body
	return e
}

func SendEmailImpl(email *email) {

	fmt.Println(email)

}

// BuilderParameter

type Build func(*EmailBuilder)

func SendEmail(action Build) {

	builder := EmailBuilder{}
	action(&builder)
	SendEmailImpl(&builder.email)
}

func main() {

	SendEmail(func(builder *EmailBuilder) {
		builder.From("foo@bar.com").To("bar@baz.com").Subject("Meeting").Body(
			"Hello do you want to meet")

	})

}
