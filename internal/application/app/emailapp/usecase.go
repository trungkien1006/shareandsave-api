package emailapp

import "final_project/internal/domain/email"

type SendEmailUseCase struct {
	repo email.Repository
}

func NewSendEmailUseCase(repo email.Repository) *SendEmailUseCase {
	return &SendEmailUseCase{repo: repo}
}

func (uc *SendEmailUseCase) Execute(to, subject, body string) error {
	return uc.repo.Send(to, subject, body)
}
