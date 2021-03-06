// Package pkg ...
package pkg

func fn(m map[int]int) {
	if _, ok := m[0]; ok { // MATCH "unnecessary guard"
		delete(m, 0)
	}
	if _, ok := m[0]; !ok {
		delete(m, 0)
	}
	if _, ok := m[0]; ok {
		println("deleting")
		delete(m, 0)
	}
	if v, ok := m[0]; ok && v > 0 {
		delete(m, 0)
	}

	var key int
	if _, ok := m[key]; ok { // MATCH "unnecessary guard"
		delete(m, key)
	}
	if _, ok := m[key]; ok {
		delete(m, 0)
	}

	var ok bool
	if _, ok = m[key]; ok {
		delete(m, 0)
	}
	if ok {
		println("deleted")
	}
}
