package todo

var todos []Todo
var autoID = 1

func init() {
	loadedTodos, err := loadFromFile()
	if err == nil {
		todos = loadedTodos

		//Cari ID terbesar
		for _, t := range todos {
			if t.ID >= autoID {
				autoID = t.ID + 1
			}
		}
	} else {
		todos = []Todo{}
	}
}

func GetAll() []Todo {
	return todos
}

func Create(title, createdAt, priority string) Todo {
	todo := Todo{
		ID:    autoID,
		Title: title,
		CreatedAt: createdAt,
		Priority: priority,
		Completed:  false,
	}
	autoID++
	todos = append(todos, todo)

	//simpan ke file
	_ = saveToFile(todos)

	return todo
}

func DeleteByID(id int) bool {
	for i, todo := range todos {
		if todo.ID == id {
			// hapus item dari slice
			todos = append(todos[:i], todos[i+1:]...)
			_ = saveToFile(todos)
			return true
		}
	}
	return false
}

func ToggleDoneByID(id int) (Todo, bool) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = !todos[i].Completed
			_ = saveToFile(todos)
			return todos[i], true
		}
	}
	return Todo{}, false
}
