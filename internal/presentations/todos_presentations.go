package presentations

type CreateTodoParam struct {
	Title string `json:"title"`
}

type CreateTodoGroupParam struct {
	Title string `json:"title"`
}

type AllTodoGroupResponse struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	UniqueName      string `json:"unique_name"`
	TodoDone        int    `json:"todo_done"`
	TodoDonePercent string `json:"todo_done_percent"`
	TodoTotal       int    `json:"todo_total"`
}
