package duplicates

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckDuplicateIDs(t *testing.T) {
	// Duplicates at the same level
	testTreeOne := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	result, err := CheckDuplicateIDs(testTreeOne)
	expectedValue := 2
	assert.Equal(t, &Results{Value: &expectedValue, Level: 2}, result)
	assert.Equal(t, nil, err)
}

func TestCheckDuplicates2(t *testing.T) {
	// Duplicates at different levels (return level of dupe identified)
	testTreeTwo := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 10,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
	}
	result, err := CheckDuplicateIDs(testTreeTwo)
	expectedValue := 5
	// Expect it to return 5 and 2 (level index)
	assert.Equal(t, &Results{Value: &expectedValue, Level: 2}, result)
	assert.Equal(t, nil, err)
}

func TestCheckDuplicates3(t *testing.T) {
	// No duplicates
	testTreeThree := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 14,
			},
		},
	}
	result, err := CheckDuplicateIDs(testTreeThree)
	// Expect it to return nil and 0
	assert.Equal(t, &Results{Value: nil, Level: 0}, result)
	assert.Equal(t, nil, err)
}

func TestCheckDuplicates4(t *testing.T) {
	// Unequal branches
	testTreeFour := &TreeNode{
		Val: 12,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 7,
						Left: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 13,
				Right: &TreeNode{
					Val: 14,
				},
			},
			Right: &TreeNode{
				Val: 18,
				Right: &TreeNode{
					Val: 20,
				},
			},
		},
	}
	result, err := CheckDuplicateIDs(testTreeFour)
	expectedValue := 15
	// Expect it to return 15 and 3
	assert.Equal(t, &Results{Value: &expectedValue, Level: 3}, result)
	assert.Equal(t, nil, err)
}

func TestCheckDuplicates5(t *testing.T) {
	// Error case
	result, err := CheckDuplicateIDs(nil)
	// Expect it to return nil and 0
	assert.Equal(t, &Results{Value: nil, Level: 0}, result)
	assert.Equal(t, errors.New("no tree provided"), err)
}
