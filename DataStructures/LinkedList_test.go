package DataStructures

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var ll ILinkedList

func init() {
	ll = &LinkedList{}
}
func TestInsert(t *testing.T) {
	testData := []struct {
		input    interface{}
		expected *LinkedList
	}{
		{
			input:    5,
			expected: &LinkedList{Head: &ListNode{Data: 5}},
		},
		{
			input:    2,
			expected: &LinkedList{Head: &ListNode{Data: 5, Next: &ListNode{Data: 2}}},
		},
	}
	for _, td := range testData {
		ll.Insert(td.input)
		if !cmp.Equal(ll, td.expected) {
			t.Errorf("Expected : %v, Actual : %v", td.expected, ll)
		}
	}
}

func TestReverse(t *testing.T) {
	testData := []struct {
		input  *LinkedList
		output *LinkedList
	}{
		{
			input: &LinkedList{
				Head: &ListNode{
					Data: 1,
					Next: &ListNode{
						Data: 2,
						Next: &ListNode{
							Data: 3,
							Next: &ListNode{
								Data: 4,
							},
						},
					}}},
			output: &LinkedList{
				Head: &ListNode{
					Data: 4,
					Next: &ListNode{
						Data: 3,
						Next: &ListNode{
							Data: 2,
							Next: &ListNode{
								Data: 1,
							},
						},
					},
				},
			},
		},
		{
			input: &LinkedList{
				Head: &ListNode{
					Data: 1,
				},
			},
			output: &LinkedList{
				Head: &ListNode{
					Data: 1,
				},
			},
		},
		{
			input: &LinkedList{
				Head: &ListNode{
					Data: 1,
					Next: &ListNode{
						Data: 2,
						Next: &ListNode{
							Data: 3,
						},
					}}},
			output: &LinkedList{
				Head: &ListNode{
					Data: 3,
					Next: &ListNode{
						Data: 2,
						Next: &ListNode{
							Data: 1,
						},
					},
				},
			},
		},
	}
	for _, td := range testData {
		ll = td.input
		ll.Reverse()
		if !cmp.Equal(td.output, ll) {
			t.Errorf("Expected : %v, Actual : %v", td.output, ll)
		}
	}
}
