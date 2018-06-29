package users

import (
	// "fmt"
	"testing"
)

// func TestGet() {
	// got := Get()
	// if got != want {
	// 	t.Errorf("Delete(3,'tester',30000) == %q, want %q", got, want)
	// }
// }

func TestInsert(t *testing.T) {
	// fmt.Println("masuk1")
	got := Insert(3,"tester",30000)
	want := "Insert Success"
	// fmt.Println("masuk")
	if got != want {
		t.Errorf("Insert(3,'tester_insert',30000) == %q, want %q", got, want)
	}
}

func TestUpdate(t *testing.T) {
	got := Update(3,"tester_update",3000)
	want := "Update Success"
	if got != want {
		t.Errorf("Update(3,'tester_update',3000000) == %q, want %q", got, want)
	}
}

func TestDelete(t *testing.T) {
	got := Delete(3)
	want := "Delete Success"
	if got != want {
		t.Errorf("Delete(3) == %q, want %q", got, want)
	}
}