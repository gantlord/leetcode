package lc380

import (
	"testing"
)

func TestRandomizedSet_Insert(t *testing.T) {
	type fields struct {
		set map[int]bool
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedSet{
				set: tt.fields.set,
			}
			if got := this.Insert(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_Remove(t *testing.T) {
	type fields struct {
		set map[int]bool
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedSet{
				set: tt.fields.set,
			}
			if got := this.Remove(tt.args.val); got != tt.want {
				t.Errorf("RandomizedSet.Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandomizedSet_GetRandom(t *testing.T) {
	type fields struct {
		set map[int]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &RandomizedSet{
				set: tt.fields.set,
			}
			if got := this.GetRandom(); got != tt.want {
				t.Errorf("RandomizedSet.GetRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}
