package info

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestTOMLNull(t *testing.T) {
	info := (*Info)(nil)
	info.Add(nil)
	res, err := info.TOML()
	if err != nil {
		t.Errorf("TOML() err = %v, want nil.", err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	if buf.String() != "" {
		t.Errorf("TOML() = %v, want \"\".", buf.String())
	}
}

func TestTOMLEmpty(t *testing.T) {
	info := NewInfo()
	res, err := info.TOML()
	if err != nil {
		t.Errorf("TOML() err = %v, want nil.", err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	if buf.String() != "" {
		t.Errorf("TOML() = %v, want \"\".", buf.String())
	}
}

func TestTOML(t *testing.T) {
	norm := `
[[Packet]]
  name = "name1"
  value = "value1"
  dump = "00 01 02"
  note = "note1"

  [[Packet.Item]]
    name = "name2"
    dump = "03 04 05"
    note = "note2"
`
	output := strings.Trim(norm, " \t\n\r") + "\n"

	info := NewInfo()
	item1 := NewItem(
		Name("name1"),
		Value("value1"),
		Note("note1"),
		DumpStr("00 01 02"),
	)
	item2 := NewItem(
		Name("name2"),
		Note("note2"),
		DumpStr("03 04 05"),
	)
	item1.Add(item2)
	item1.Add(nil) //abnormal
	info.Add(item1)
	info.Add(nil) //abnormal
	toml, err := info.TOML()
	if err != nil {
		t.Errorf("MarshalTOML() = \"%v\"want nil.", err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, toml)
	str := buf.String()
	if str != output {
		t.Errorf("TOML output = \n%s\n want \n%s\n", str, output)
	}
}

func TestItemString(t *testing.T) {
	output := `name1: value1 (note1)
	00 01 02
	name2 (note2)
		03 04 05
`
	item1 := NewItem(
		Name("name1"),
		Value("value1"),
		Note("note1"),
		DumpStr("00 01 02"),
	)
	item2 := NewItem(
		Name("name2"),
		Note("note2"),
		DumpStr("03 04 05"),
	)
	item1.Add(item2)
	str := item1.String()
	if str != output {
		t.Errorf("TOML output = \n%s\n want \n%s\n", str, output)
	}
}

func ExampleNewInfo() {
	item := NewItem(
		Name("name"),
		Value("value"),
		Note("note"),
		DumpStr("00 01 02"),
	)
	fmt.Println(item.Dump)
	// Output:
	// 00 01 02
}

func TestJSONNull(t *testing.T) {
	info := (*Info)(nil)
	info.Add(nil)
	res, err := info.JSON()
	if err != nil {
		t.Errorf("JSON() err = %v, want nil.", err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	if buf.String() != "" {
		t.Errorf("JSON() = %v, want \"\".", buf.String())
	}
}

func TestJSONEmpty(t *testing.T) {
	info := NewInfo()
	res, err := info.JSON()
	if err != nil {
		t.Errorf("JSON() err = %v, want nil.", err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, res)
	str := buf.String()
	if str != "{}\n" {
		t.Errorf("TOML() = %v, want {}.", str)
	}

}

func TestJSON(t *testing.T) {
	norm := `
{
  "Packet": [
    {
      "name": "name1",
      "value": "value1",
      "dump": "00 01 02",
      "note": "note1",
      "Item": [
        {
          "name": "name2",
          "dump": "03 04 05",
          "note": "note2"
        }
      ]
    }
  ]
}
`
	output := strings.Trim(norm, " \t\n\r")

	info := NewInfo()
	item1 := NewItem(
		Name("name1"),
		Value("value1"),
		Note("note1"),
		DumpStr("00 01 02"),
	)
	item2 := NewItem(
		Name("name2"),
		Note("note2"),
		DumpStr("03 04 05"),
	)
	item1.Add(item2)
	info.Add(item1)
	json, err := info.JSON()
	if err != nil {
		t.Errorf("JSON() = \"%v\"want nil.", err)
		return
	}
	buf := new(bytes.Buffer)
	io.Copy(buf, json)
	str := buf.String()
	if str != output+"\n" {
		t.Errorf("JSON output = \"%s\" want \"%s\"", str, output)
	}
}

func TestStringer(t *testing.T) {
	norm := `name1: value1 (note1)
	00 01 02
	name2 (note2)
		03 04 05
`
	output := strings.Trim(norm, " \t\n\r")

	info := NewInfo()
	item1 := NewItem(
		Name("name1"),
		Value("value1"),
		Note("note1"),
		DumpStr("00 01 02"),
	)
	item2 := NewItem(
		Name("name2"),
		Note("note2"),
		DumpStr("03 04 05"),
	)
	item1.Add(item2)
	info.Add(item1)
	str := info.String()
	if str != output+"\n" {
		t.Errorf("String() = \"%s\" want \"%s\"", str, output)
	}
}

/* Copyright 2017 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */