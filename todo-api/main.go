package main

func main() {
	type todo struct {
		ID        string
		Item      string
		Completed bool
	}

	todos := []todo{
		{ID: "1", Item: "Clean Room", Completed: false},
		{ID: "2", Item: "Read Book", Completed: false},
		{ID: "3", Item: "Record Video", Completed: false},
	}
}
