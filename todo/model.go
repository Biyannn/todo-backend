package todo

type Todo struct {
	ID int   `json:"id"`
	Title string `json:"title"`
	Priority string `json:"priority"` // low | medium | high
	Completed bool `json:"completed"`
	CreatedAt string `json:"createdAt"`
}