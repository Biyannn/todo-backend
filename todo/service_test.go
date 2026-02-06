package todo

// import "testing"

// func resetTodos() {
// 	todos = []Todo{}
// 	autoID = 1
// }

// func testCreateTodo(t *testing.T) {
// 	resetTodos()

// 	todo := Create("Belajar Go", "2026-01-27", "high")

// 	if todo.ID != 1 {
// 		t.Errorf("Expected ID 1, got %d", todo.ID)
// 	}

// 	if todo.Completed != false {
// 		t.Errorf("Expected Completed false, got %v", todo.Completed)
// 	}

// 	if len(todos) != 1 {
// 		t.Errorf("Expected 1 todo, got %d", len(todos))
// 	}
// }

// func TestGetAllTodos(t *testing.T) {
// 	resetTodos()

// 	Create("Todo 1", "", "")
// 	Create("Todo 2", "", "")

// 	all := GetAll()

// 	if len(all) != 2 {
// 		t.Errorf("expected 2 todos, got %d", len(all))
// 	}
// }

// func TestToggleDone(t *testing.T) {
// 	resetTodos()

// 	todo := Create("Belajar Go", "", "")

// 	updated, success := ToggleDoneByID(todo.ID)
// 	if !success {
// 		t.Fatal("expected toggle success")
// 	}

// 	if updated.Completed != true {
// 		t.Errorf("expected Completed true, got %v", updated.Completed)
// 	}
// }

// func TestDeleteTodo(t *testing.T) {
// 	resetTodos()

// 	todo := Create("Delete me", "", "")

// 	success := DeleteByID(todo.ID)
// 	if !success {
// 		t.Fatal("expected delete success")
// 	}

// 	if len(todos) != 0 {
// 		t.Errorf("expected 0 todos, got %d", len(todos))
// 	}
// }

// func TestDeleteTodoNotFound(t *testing.T)  {
// 	resetTodos()

// 	success := DeleteByID(999)

// 	if success {
// 		t.Error("expected delete to fail, but succeeded")
// 	}
// }