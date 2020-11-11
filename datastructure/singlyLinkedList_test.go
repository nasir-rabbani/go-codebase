package main

import "testing"

func Test_list_insert(t *testing.T) {
	type args struct {
		key int
	}
	tests := []struct {
		name string
		ll   *list
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test 1",
			ll: &list{
				head: &node{},
			},
			args: args{key: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ll.insert(tt.args.key)
		})
	}
}
