package main

type Set map[string]struct{}

// Add adds an element to the set.
func (s Set) Add(item string) {
	s[item] = struct{}{}
}

// Remove removes an element from the set.
func (s Set) Remove(item string) {
	delete(s, item)
}

// Contains checks if an element is present in the set.
func (s Set) Contains(item string) bool {
	_, exists := s[item]
	return exists
}

// Size returns the size of the set.
func (s Set) Size() int {
	return len(s)
}
