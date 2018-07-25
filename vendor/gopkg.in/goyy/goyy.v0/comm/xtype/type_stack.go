// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xtype

// FILO : First In Last Out
type Stack []interface{}

func (me Stack) Len() int {
	return len(me)
}

func (me *Stack) Push(v interface{}) {
	*me = append(*me, v)
}

func (me *Stack) Pop() interface{} {
	old := *me
	n := len(old)
	if n == 0 {
		return nil
	}
	s := old[n-1]
	*me = old[0 : n-1]
	return s
}
