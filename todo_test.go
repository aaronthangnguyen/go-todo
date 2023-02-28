package todo_test

import (
	"github.com/aaronthangnguyen/go-todo"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	task := "New task"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expected %q, got %q instead.", task, l[0].Task)
	}
}

func TestComplete(t *testing.T) {
	l := todo.List{}

	task := "New task"
	l.Add(task)

	if l[0].Task != task {
		t.Errorf("Expected %q, got %q instead.", task, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New task should not be completed.")
	}

	l.Complete(1)

	if !l[0].Done {
		t.Errorf("New task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	l := todo.List{}

	tasks := []string{
		"New task 1",
		"New task 2",
		"New task 3",
	}

	for _, task := range tasks {
		l.Add(task)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}

	l.Delete(2)

	if len(l) != 2 {
		t.Errorf("Expected list length %d, got %d instead", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	l1 := todo.List{}
	l2 := todo.List{}

	task := "New task"
	l1.Add(task)

	if l1[0].Task != task {
		t.Errorf("Expected %q, got %q instead", task, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := l1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.Get(tf.Name()); err != nil {
		t.Fatalf("Error getting list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should mawtch %q task", l1[0].Task, l2[0].Task)
	}
}