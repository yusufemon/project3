package users_test

import (
	"testing"

	"github.com/yusufemon/project3/app/users"
)

func TestGet(t *testing.T) {
	user, err := users.New()
	if err != nil {
		t.Errorf("method user.Get() is error, error message %q", err)
	}
	_, err = user.Get()
	if err != nil {
		t.Errorf("method user.Get() is error, error message %q", err)
	}
}

func TestInsert(t *testing.T) {
	user, _ := users.New()
	err := user.Insert(3, "tester", 30000)
	if err != nil {
		t.Errorf("method user.Insert(3,'tester_insert',30000) is error, error message %q", err)
	}
}

func TestUpdate(t *testing.T) {
	user, _ := users.New()
	err := user.Update(3, "tester_update", 3000)
	if err != nil {
		t.Errorf("method user.Update(3,'tester_update',3000000) is error, error message %q", err)
	}
}

func TestDelete(t *testing.T) {
	user, _ := users.New()
	err := user.Delete(3)
	if err != nil {
		t.Errorf("method user.Delete(3) is error, error message %q", err)
	}
}
