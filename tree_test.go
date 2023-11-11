package tree

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDir(t *testing.T) {
	tests := []struct {
		name    string
		dirname string
	}{
		{"simple", "/home/saspeh/VSProjects" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, err := NewDir(tt.dirname, nil)
			assert.Nil(t, err)
			fmt.Println(dir)
		})
	}
}

func TestNewFile(t *testing.T) {
	type args struct {
		filename string
		parent   *Dir
	}
	tests := []struct {
		name string
		args args
		want *File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFile(tt.args.filename, tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDir_Print(t *testing.T) {
	type fields struct {
		Name     string
		Parent   *Dir
		Children []any
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Dir{
				Name:     tt.fields.Name,
				Parent:   tt.fields.Parent,
				Children: tt.fields.Children,
			}
			p.Print()
		})
	}
}

func TestDir_GetPath(t *testing.T) {

	dir, err := NewDir("/home/saspeh/VSProjects", nil)
	assert.NotNil(t, dir)
	assert.Nil(t, err)
}
