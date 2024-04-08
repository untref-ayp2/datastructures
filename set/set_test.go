package set

import "testing"

func TestNewSet(t *testing.T) {
	set := NewSet[int]()
	if set == nil {
		t.Error("NewSet() devolvió nil")
	}
	if set.Size() != 0 {
		t.Error("El tamaño del conjunto debería ser 0")
	}
}

func TestAdd(t *testing.T) {
	set := NewSet[int]()
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
}

func TestContains(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	if !set.Contains(1) {
		t.Error("El conjunto debería tener el número 1")
	}
}
func TestRemove(t *testing.T) {
	set := NewSet[int]()
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
	set := NewSet(1, 2)
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
func TestEqual(t *testing.T) {
	set1 := NewSet(1, 2)
	set2 := NewSet(1, 2)
	if !Equal(set1, set2) {
		t.Errorf("Los conjuntos deberian ser iguales")
	}
}

func TestUnion(t *testing.T) {
	set1 := NewSet(1, 2)
	set2 := NewSet(1, 3)
	set3 := NewSet(1, 2, 3)
	result := Union(set1, set2)
	if !Equal(result, set3) {
		t.Errorf("La union de set1 y set2 deberia dar [1, 2, 3]")
	}
}

func TestIntersection(t *testing.T) {
	set1 := NewSet(1, 2)
	set2 := NewSet(1, 3)
	set3 := NewSet(1)
	result := Intersection(set1, set2)
	if !Equal(result, set3) {
		t.Errorf("La intersection de set1 y set2 deberia dar [1]")
	}
}

func TestDifference(t *testing.T) {
	set1 := NewSet(1, 2)
	set2 := NewSet(1, 3)
	set3 := NewSet(2)
	result := Difference(set1, set2)
	if !Equal(result, set3) {
		t.Errorf("La difference de set1 y set2 deberia dar [2]")
	}
}

func TestSubset(t *testing.T) {
	set1 := NewSet(1, 2)
	set2 := NewSet(1)
	if !Subset(set1, set2) {
		t.Errorf("set2 deberia ser un subconjunto de set1")
	}
}
