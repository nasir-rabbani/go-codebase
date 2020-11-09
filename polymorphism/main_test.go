package main

import (
	"testing"
)

func Test_person_nameLen(t *testing.T) {
	tests := []struct {
		name string
		p    person
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Happy scenario",
			p: person{
				personName: "Rabbani",
				personAge:  25,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.nameLen(); got != tt.want {
				t.Errorf("person.nameLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
