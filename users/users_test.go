package users_test

import (
	"testing"

	"github.com/yusufemon/project3/users"
)

func TestGet(t *testing.T) {
	user, _ := users.New()
	_, err := user.Get()
	if err != nil {
		t.Errorf("method user.Get() is error, error message %q", err)
	}
}

func TestInsert(t *testing.T) {
	user, _ := users.New()
	got, _ := user.Insert(3, "tester", 30000)
	want := "Insert Success"
	if got != want {
		t.Errorf("Insert(3,'tester_insert',30000) == %q, want %q", got, want)
	}
}

func TestUpdate(t *testing.T) {
	user, _ := users.New()
	got, _ := user.Update(3, "tester_update", 3000)
	want := "Update Success"
	if got != want {
		t.Errorf("Update(3,'tester_update',3000000) == %q, want %q", got, want)
	}
}

func TestDelete(t *testing.T) {
	user, _ := users.New()
	got, _ := user.Delete(3)
	want := "Delete Success"
	if got != want {
		t.Errorf("Delete(3) == %q, want %q", got, want)
	}
}
