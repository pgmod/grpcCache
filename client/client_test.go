package client

import (
	"context"
	"testing"
)

func TestSave(t *testing.T) {
	client, err := CreateClient("localhost:50051")
	if err != nil {
		t.Error(err)
		return
	}
	msg, err := client.Save(context.Background(), "test_id", "test_club_id", "test_extra_info")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(msg)
}
func TestClear(t *testing.T) {
	client, err := CreateClient("localhost:50051")
	if err != nil {
		t.Error(err)
		return
	}
	msg, err := client.Clear(context.Background(), "test_club_id")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(msg)
}
