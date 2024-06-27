package Nbj

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_a_collection_can_be_created_with_a_type_items(t *testing.T) {
	// Act
	collection := NewCollection[int]()

	// Assert
	assert.Equal(t, "*Nbj.Collection[int]", reflect.TypeOf(collection).String())
}

func Test_a_collection_knows_if_it_is_empty(t *testing.T) {
	// Arrange
	collection := NewCollection[int]()

	// Assert
	assert.True(t, collection.IsEmpty())
}

func Test_a_collection_can_have_items_added_to_it(t *testing.T) {
	// Arrange
	collection := NewCollection[int]()
	assert.True(t, collection.IsEmpty())

	// Act
	collection.Add(1000)

	// Assert
	assert.True(t, collection.IsNotEmpty())
}

func Test_a_collection_knows_how_many_items_it_is_storing(t *testing.T) {
	// Arrange
	collection := NewCollection[int]()
	assert.Equal(t, 0, collection.Count())

	// Act
	collection.Add(1000)
	collection.Add(1000)

	// Assert
	assert.Equal(t, 2, collection.Count())
}

func Test_a_collection_can_get_its_first_item(t *testing.T) {
	// Arrange
	collection := NewCollection[string]()
	collection.Add("first")
	collection.Add("middle")
	collection.Add("last")

	// Act
	first := collection.First()

	// Assert
	assert.Equal(t, "first", first)
}

func Test_a_collection_can_get_its_last_item(t *testing.T) {
	// Arrange
	collection := NewCollection[string]()
	collection.Add("first")
	collection.Add("middle")
	collection.Add("last")

	// Act
	last := collection.Last()

	// Assert
	assert.Equal(t, "last", last)
}

func Test_a_collection_can_shift_its_first_item_of_it(t *testing.T) {
	// Arrange
	collection := NewCollection[string]()
	collection.Add("first")
	collection.Add("middle")
	collection.Add("last")
	assert.Equal(t, 3, collection.Count())

	// Act
	first := collection.Shift()

	// Assert
	assert.Equal(t, "first", first)
	assert.Equal(t, 2, collection.Count())
	assert.Equal(t, "middle", collection.First())
	assert.Equal(t, "last", collection.Last())
}

func Test_a_collection_can_pop_its_last_item_of_it(t *testing.T) {
	// Arrange
	collection := NewCollection[string]()
	collection.Add("first")
	collection.Add("middle")
	collection.Add("last")
	assert.Equal(t, 3, collection.Count())

	// Act
	last := collection.Pop()

	// Assert
	assert.Equal(t, "last", last)
	assert.Equal(t, 2, collection.Count())
	assert.Equal(t, "first", collection.First())
	assert.Equal(t, "middle", collection.Last())
}

func Test_a_collection_can_be_filled_with_a_slice_containing_elements_of_the_collection_type(t *testing.T) {
	// Arrange
	collection := NewCollection[string]()
	assert.Equal(t, 0, collection.Count())

	sliceOfString := []string{
		"first",
		"middle",
		"last",
	}

	// Act
	collection.Fill(sliceOfString)

	// Assert
	assert.Equal(t, 3, collection.Count())
	assert.Equal(t, "first", collection.First())
	assert.Equal(t, "last", collection.Last())
}

func Test_a_collection_can_be_merged_with_another_collection_containing_items_of_the_same_type(t *testing.T) {
	// Arrange
	collectionA := NewCollection[string]()
	collectionA.Fill([]string{
		"first",
		"second",
	})

	assert.Equal(t, 2, collectionA.Count())

	collectionB := NewCollection[string]()
	collectionB.Fill([]string{
		"third",
		"fourth",
		"fifth",
	})

	assert.Equal(t, 3, collectionB.Count())

	// Act
	collectionA.Merge(collectionB)

	// Assert
	assert.Equal(t, 5, collectionA.Count())
	assert.Equal(t, "first", collectionA.First())
	assert.Equal(t, "fifth", collectionA.Last())
}

func Test_a_collection_can_be_created_from_a_slice(t *testing.T) {
	// Arrange
	slice := []string{
		"first",
		"middle",
		"last",
	}

	// Act
	collection := Collect[string](slice)

	// Assert
	assert.Equal(t, "*Nbj.Collection[string]", reflect.TypeOf(collection).String())
	assert.Equal(t, 3, collection.Count())
	assert.Equal(t, "first", collection.First())
	assert.Equal(t, "last", collection.Last())
}

func Test_a_collection_knows_if_it_contains_a_specific_item(t *testing.T) {
	// Arrange
	collection := Collect[string]([]string{
		"first",
		"middle",
		"last",
	})

	// Assert
	assert.True(t, collection.Contains(func(item string) bool {
		return item == "middle"
	}))
}

func Test_a_collection_can_iterate_over_all_it_items(t *testing.T) {
	// Arrange
	collection := Collect[string]([]string{
		"first",
		"middle",
		"last",
	})

	// Assert
	collection.ForEach(func(item string) {
		assert.True(t, collection.Contains(func(innerItem string) bool {
			return item == innerItem
		}))
	})
}

func Test_a_collection_can_be_reduced(t *testing.T) {
	// Arrange
	collection := Collect[int]([]int{1, 2, 3, 4, 5})

	// Act - Implement Sum() using Reduce()
	sum := collection.Reduce(func(carry any, item int) any {
		carry = carry.(int) + item

		return carry
	}, 0)

	// Assert - Sum() implementation
	assert.Equal(t, 15, sum)

	// Act - Implement Reverse() using Reduce()
	reversed := collection.Reduce(func(carry any, item int) any {
		carry = carry.(*Collection[int]).Prepend(item)

		return carry
	}, NewCollection[int]())

	// Assert
	assert.Equal(t, 1, collection.First())
	assert.Equal(t, 5, collection.Last())
	assert.Equal(t, 5, reversed.(*Collection[int]).First())
	assert.Equal(t, 1, reversed.(*Collection[int]).Last())
}

func Test_a_collections_items_can_be_filtered_using_a_closure(t *testing.T) {
	// Arrange
	collection := Collect[string]([]string{
		"first",
		"middle",
		"last",
	})

	assert.True(t, collection.Contains(func(item string) bool {
		return item == "middle"
	}))

	// Act
	collection = collection.Filter(func(item string) bool {
		return item == "first" || item == "last"
	})

	// Assert
	assert.False(t, collection.Contains(func(item string) bool {
		return item == "middle"
	}))
}

func Test_a_collections_items_can_be_rejected_using_a_closure(t *testing.T) {
	// Arrange
	collection := Collect[string]([]string{
		"first",
		"middle",
		"last",
	})

	assert.True(t, collection.Contains(func(item string) bool {
		return item == "middle"
	}))

	// Act
	collection = collection.Reject(func(item string) bool {
		return item == "middle"
	})

	// Assert
	assert.False(t, collection.Contains(func(item string) bool {
		return item == "middle"
	}))
}

func Test_a_collections_items_can_be_mapped_into_something_else(t *testing.T) {
	// Arrange
	collection := Collect[string]([]string{
		"first",
		"middle",
		"last",
	})

	// Act
	lengths := collection.Map(func(item string) any {
		return len(item)
	})

	// Assert
	assert.True(t, lengths.Contains(func(item any) bool {
		return item == 6
	}))

	assert.Equal(t, 5, lengths.First())
	assert.Equal(t, 4, lengths.Last())
}

func Test_a_collection_can_pluck_specific_fields_of_its_containing_items(t *testing.T) {
	// Arrange
	var objects []TestObject

	objects = append(objects, TestObject{Id: 1, Name: "John"})
	objects = append(objects, TestObject{Id: 2, Name: "Jane"})
	objects = append(objects, TestObject{Id: 3, Name: "Charlie"})

	// Act
	collectionA := Collect(objects).Pluck("Id")
	collectionB := Collect(objects).Pluck("Name")

	// Assert
	assert.Equal(t, 1, collectionA.First())
	assert.Equal(t, 3, collectionA.Last())

	assert.Equal(t, "John", collectionB.First())
	assert.Equal(t, "Charlie", collectionB.Last())
}

func Test_a_collection_can_return_all_its_items_as_an_array(t *testing.T) {
	// Arrange
	var objects []TestObject

	objects = append(objects, TestObject{Id: 1, Name: "John"})
	objects = append(objects, TestObject{Id: 2, Name: "Jane"})
	objects = append(objects, TestObject{Id: 3, Name: "Charlie"})

	collection := Collect(objects)

	// Act
	items := collection.All()

	// Assert
	assert.Equal(t, "[]Nbj.TestObject", reflect.TypeOf(items).String())
	assert.Equal(t, 1, items[0].Id)
	assert.Equal(t, "John", items[0].Name)
}

func Test_a_collection_can_return_a_specific_item_based_on_its_index_in_the_collection(t *testing.T) {
	// Arrange
	var objects []TestObject

	objects = append(objects, TestObject{Id: 1, Name: "John"})
	objects = append(objects, TestObject{Id: 2, Name: "Jane"})
	objects = append(objects, TestObject{Id: 3, Name: "Charlie"})

	collection := Collect(objects)

	// Act
	object := collection.Get(1)

	// Assert
	assert.Equal(t, 2, object.Id)
	assert.Equal(t, "Jane", object.Name)
}

func Test_a_collection_can_return_an_index_of_a_specific_item_based_on_a_closure_passed(t *testing.T) {
	// Arrange
	var objects []TestObject

	objects = append(objects, TestObject{Id: 1, Name: "John"})
	objects = append(objects, TestObject{Id: 2, Name: "Jane"})
	objects = append(objects, TestObject{Id: 3, Name: "Charlie"})

	collection := Collect(objects)
	object := TestObject{Id: 3, Name: "Charlie"}

	// Act
	index := collection.IndexOf(object)

	// Assert
	assert.Equal(t, 2, index)
}

func Test_a_collection_can_return_minus_one_if_index_of_could_not_find_the_specific_item(t *testing.T) {
	// Arrange
	var objects []TestObject

	collection := Collect(objects)

	object := TestObject{Id: 1, Name: "John"}

	// Act
	index := collection.IndexOf(object)

	// Assert
	assert.Equal(t, -1, index)
}

type TestObject struct {
	Id   int
	Name string
}
