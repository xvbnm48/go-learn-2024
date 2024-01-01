package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("TestIncrement", TestIncrement)
	t.Run("TestDecrement", TestDecrement)
	t.Run("TestConvertToString", TestConvertToString)
	t.Run("TestCheckIfString", TestCheckIfString)
}

func TestIncrement(t *testing.T) {
	i := 10
	i += 10
	expected := 20
	if i != expected {
		t.Errorf("Expected %d, but got %d", expected, i)
	}
}

func TestDecrement(t *testing.T) {
	i := 20
	i--
	i--
	expected := 18
	if i != expected {
		t.Errorf("Expected %d, but got %d", expected, i)
	}
}

func TestConvertToString(t *testing.T) {
	i := 18
	strI := strconv.Itoa(i)
	expected := "18"
	if strI != expected {
		t.Errorf("Expected %s, but got %s", expected, strI)
	}
}

func TestCheckIfString(t *testing.T) {
	i := 18
	strI := strconv.Itoa(i)
	strIType := reflect.TypeOf(strI).Kind()
	expected := reflect.String
	if strIType != expected {
		t.Errorf("Expected %s, but got %s", expected, strIType)
	}
}
