package utils

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Do laundry", Completed: true},
	{ID: "3", Item: "Program", Completed: false},
}
