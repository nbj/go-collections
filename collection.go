package Nbj

import (
	"slices"
)

// Collection
// The general structure of a collection
type Collection[T comparable] struct {
	items []T
}

// NewCollection
// Named constructor to create a new collection
func NewCollection[T comparable]() *Collection[T] {
	var collection Collection[T]

	return &collection
}

// Collect
// Named constructor to create a new collection from a slice
func Collect[T comparable](items []T) *Collection[T] {
	var collection Collection[T]
	collection.Fill(items)

	return &collection
}

// Fill
// Fills the collection with a slice of items
func (collection *Collection[T]) Fill(items []T) *Collection[T] {
	collection.items = items

	return collection
}

// Prepend
// Adds an item to the start of the collection
func (collection *Collection[T]) Prepend(item T) *Collection[T] {
	collection.items = append([]T{item}, collection.items...)

	return collection
}

// Add
// Adds an item to the end of the collection
func (collection *Collection[T]) Add(item T) *Collection[T] {
	collection.items = append(collection.items, item)

	return collection
}

// Push
// Syntactic sugar for adding (pushing) an item on to the collection
// This simply defers the call to Add() which does the exact same thing
func (collection *Collection[T]) Push(item T) *Collection[T] {
	return collection.Add(item)
}

// First
// Gets the first item in the collection
func (collection *Collection[T]) First() T {
	return collection.items[0]
}

// Last
// Gets the last item in the collection
func (collection *Collection[T]) Last() T {
	return collection.items[len(collection.items)-1]
}

// Shift
// Shifts the first item off the start of the collection and returns it
func (collection *Collection[T]) Shift() T {
	first, rest := collection.First(), collection.items[1:]

	collection.items = rest

	return first
}

// Pop
// Pops the last item of the collection and returns it
func (collection *Collection[T]) Pop() T {
	last, rest := collection.Last(), collection.items[:len(collection.items)-1]

	collection.items = rest

	return last
}

// Count
// Gets the number of items in the collection
func (collection *Collection[T]) Count() int {
	return len(collection.items)
}

// IsEmpty
// Tells if the collection is empty
func (collection *Collection[T]) IsEmpty() bool {
	return collection.Count() == 0
}

// IsNotEmpty
// Tells if the collection is NOT empty
func (collection *Collection[T]) IsNotEmpty() bool {
	return !collection.IsEmpty()
}

// Merge
// Merges another collection into this
func (collection *Collection[T]) Merge(collectionToMerge *Collection[T]) *Collection[T] {
	collection.items = append(collection.items, collectionToMerge.items...)

	return collection
}

// Contains
// Checks if the collection contains a specific item
func (collection *Collection[T]) Contains(search T) bool {
	return slices.Contains(collection.items, search)
}

// ForEach
// Iterates over all items in the collection and executes the closure for each
func (collection *Collection[T]) ForEach(closure func(item T)) {
	for _, item := range collection.items {
		closure(item)
	}
}

// Reduce
// Reduces the collection based on the closure passed to it
func (collection *Collection[T]) Reduce(closure func(carry any, item T) any, initial any) any {
	carry := initial

	for _, item := range collection.items {
		carry = closure(carry, item)
	}

	return carry
}

// Filter
// Filters the items in the collection based on a closure
func (collection *Collection[T]) Filter(closure func(item T) bool) *Collection[T] {
	var filteredCollection Collection[T]

	for _, item := range collection.items {
		if closure(item) {
			filteredCollection.Add(item)
		}
	}

	return &filteredCollection
}

// Reject
// Rejects items in the collection based on a closure
func (collection *Collection[T]) Reject(closure func(item T) bool) *Collection[T] {
	var filteredCollection Collection[T]

	for _, item := range collection.items {
		if !closure(item) {
			filteredCollection.Add(item)
		}
	}

	return &filteredCollection
}
