/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	lst := []string{"a", "b", "a", "c", "d"}
	set := NewSet(lst)
	assert.Equal(t, len(set.GetElements()), 4)

	set.Add("a")
	assert.Equal(t, len(set.GetElements()), 4)

	set.Add("b")
	assert.Equal(t, len(set.GetElements()), 4)

	set.Add("e")
	assert.Equal(t, len(set.GetElements()), 5)
	assert.True(t, set.Contains("a"), "Set should contain `eg`")

	set.Delete("a")
	assert.Equal(t, len(set.GetElements()), 4)
	assert.False(t, set.Contains("a"), "Set should not contain `a`")

	return
}
