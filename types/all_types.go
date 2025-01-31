package types

type Response struct {
	Email           string `json:"email"`
	CurrentDateTime string `json:"current_date_time"`
	GithubUrl       string `json:"github_url"`
}
