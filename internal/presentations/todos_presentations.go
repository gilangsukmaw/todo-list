package presentations

type CreateTodoParam struct {
	Title string `json:"title"`
}

type AllTodoGroupResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	UniqueName string `json:"unique_name"`
	TodoTotal  int    `json:"todo_total"`
}
