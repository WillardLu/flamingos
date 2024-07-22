// Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package youling_string

import (
	"fmt"
	"testing"

	"github.com/pelletier/go-toml"
)

func TestReadBetween(t *testing.T) {
	type args struct {
		s     string
		start string
		end   string
	}
	type tests struct {
		name string
		args args
		want string
	}
	// Reading test data from a TOML file
	datas, err := toml.LoadFile("testdata/test_data.toml")
	if err != nil {
		t.Error("\nAn error occurred while loading the test data file test_data/"+
			"test_data.toml:\n", err)
		// The t.Error() function will continue down the line after execution,
		// so a RETURN is needed.
		return
	}
	// Read the source string for testing
	var type1 interface{}
	type1 = datas.Get("source_string")
	if fmt.Sprintf("%T", type1) != "*toml.Tree" {
		t.Error("测试数据中没有找到 source_string 表")
		return
	}
	source := type1.(*toml.Tree)
	// Read test program
	type1 = datas.Get("ReadBetween")
	if fmt.Sprintf("%T", type1) != "[]*toml.Tree" {
		t.Error("测试数据中没有找到 ReadBetween 表数组")
		return
	}
	testSchemes := type1.([]*toml.Tree)
	// Reading test data
	var tData tests
	for _, tt := range testSchemes {
		type1 := tt.Get("name")
		if fmt.Sprintf("%T", type1) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 name 字段")
			return
		}
		type2 := tt.Get("args.start")
		if fmt.Sprintf("%T", type2) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 args.start 字段")
			return
		}
		type3 := tt.Get("args.end")
		if fmt.Sprintf("%T", type3) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 args.end 字段")
			return
		}
		s1 := tt.Get("source")
		if fmt.Sprintf("%T", s1) != "string" {
			t.Error("测试数据中 ReadBetween 表里面没有找到 source 字段")
			return
		}
		tData = tests{
			name: type1.(string),
			args: args{
				s:     source.Get(s1.(string)).(string),
				start: type2.(string),
				end:   type3.(string),
			},
			want: tt.Get("want").(string),
		}
		t.Run(tData.name, func(t *testing.T) {
			got := ReadBetween(tData.args.s, tData.args.start, tData.args.end)
			if got != tData.want {
				t.Errorf("ReadBetween() = %v, want %v", got, tData.want)
			}
		})
	}
	// Error Tests
	t.Run("Error test: return content does not match", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "d/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
	t.Run("Error test: left string not found", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/1AQQ", "d/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
	t.Run("Error test: right string not found", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "d1/AQQk"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
	t.Run("Error test: right string before left string", func(t *testing.T) {
		if got := ReadBetween("inbox/id/AQQkAD", "/id/AQQ", "inbox"); got != "" {
			t.Errorf("ReadBetween() = %v, want %v", got, "")
		}
	})
}
