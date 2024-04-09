package set

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSetList[int]()
	if set == nil {
		t.Error("NewSet() devolvió nil")
	}
	if set.Size() != 0 {
		t.Error("El tamaño del conjunto debería ser 0")
	}
}

func TestAdd(t *testing.T) {
	set := NewSetList[int]()
	set.Add(1)
	if set.Size() != 1 {
		t.Error("El tamaño del conjunto debería ser 1")
	}
	set.Add(2)
	if set.Size() != 2 {
		t.Error("El tamaño del conjunto debería ser 2")
	}
	set.Add(1)
	if set.Size() != 2 {
		t.Error("El tamaño del conjunto debería ser 2")
	}
	set.Add(3, 4)
	if set.Size() != 4 {
		t.Error("El tamaño del conjunto debería ser 4")
	}
}

func TestContains(t *testing.T) {
	set := NewSetList[int]()
	set.Add(1)
	if !set.Contains(1) {
		t.Error("El conjunto debería tener el número 1")
	}
}
func TestRemove(t *testing.T) {
	set := NewSetList[int]()
	set.Add(1)
	set.Add(2)
	set.Remove(1)
	if set.Contains(1) {
		t.Error("El conjunto debería no tener el número 1")
	}
	if set.Size() != 1 {
		t.Error("El tamaño del conjunto debería ser 1")
	}

}

func TestValues(t *testing.T) {
	set := NewSetList(1, 2)
	values := set.Values()
	if len(values) != 2 {
		t.Errorf("El tamaño de los valores del conjunto debería ser 2")
	}
	if values[0] != 1 {
		t.Error("El valor debería ser 1")
	}
	if values[1] != 2 {
		t.Error("El valor debería ser 2")
	}
}
