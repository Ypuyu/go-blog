package v1

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList()
	if list.head.next != list.tail {
		t.Errorf("Expected head.next to point to tail")
	}
	if list.tail.prev != list.head {
		t.Errorf("Expected tail.prev to point to head")
	}
	if list.len != 0 {
		t.Errorf("Expected length 0, got %d", list.len)
	}
}

func TestPushBack(t *testing.T) {
	list := NewList()
	list.PushBack(5)
	list.PushBack(6)

	expected := []any{5, 6}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPushFront(t *testing.T) {
	list := NewList()
	list.PushFront(5)
	list.PushFront(4)

	expected := []any{4, 5}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertAfter(t *testing.T) {
	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	node := list.Find(2)
	list.InsertAfter(node, 25)

	expected := []any{1, 2, 25, 3}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestInsertBefore(t *testing.T) {
	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	node := list.Find(2)
	list.InsertBefore(node, 15)

	expected := []any{1, 15, 2, 3}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestFind(t *testing.T) {
	list := NewList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	node := list.Find(20)
	if node == nil || node.val != 20 {
		t.Errorf("Expected to find node with value 20")
	}

	node = list.Find(40)
	if node != nil {
		t.Errorf("Expected nil when searching for non-existent value")
	}
}

func TestRemove(t *testing.T) {
	list := NewList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	list.Remove(20)
	expected := []any{10, 30}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveHead(t *testing.T) {
	list := NewList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	list.Remove(10)
	expected := []any{20, 30}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRemoveTail(t *testing.T) {
	list := NewList()
	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(30)

	list.Remove(30)
	expected := []any{10, 20}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetLen(t *testing.T) {
	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	if list.GetLen() != 3 {
		t.Errorf("Expected length 3, got %d", list.GetLen())
	}

	list.Remove(2)
	if list.GetLen() != 2 {
		t.Errorf("Expected length 2 after removal, got %d", list.GetLen())
	}
}

func TestForEach(t *testing.T) {
	list := NewList()
	list.PushBack(1)
	list.PushBack("hello")
	list.PushBack(3.14)

	expected := []any{1, "hello", 3.14}
	result := list.ForEach()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestForReverse(t *testing.T) {
	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	expected := []any{3, 2, 1}
	result := list.ForReverse()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPrint(t *testing.T) {
	// This test is more about ensuring Print doesn't panic.
	// You can capture stdout if needed for assertions.
	list := NewList()
	list.PushBack(1)
	list.PushBack(2)
	list.Print() // Just ensure it runs without error
}
