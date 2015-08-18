package main

import (
	"fmt"
	"github.com/gourd/goparser"
	"github.com/gourd/gourd/templates"
)

func init() {
	t, err := templates.Asset("service/upperio.tpl")
	if err != nil {
		panic(err)
	}
	tpls.Append("gen service:upperio", string(t))
	tpls.AddDeps("gen service:upperio", "gen:general")
	tpls.AddPrep("gen service:upperio", func(in interface{}) (interface{}, error) {
		var data map[string]interface{}
		var f interface{}
		var id *goparser.FieldSpec
		var ok bool

		if data, ok = in.(map[string]interface{}); !ok {
			return in, fmt.Errorf("Unable to prepare. Incorrect data provided: %#v", in)
		}

		if f, ok = data["Id"]; !ok {
			return in, fmt.Errorf("Unable to prepare. No Id found in given data: %#v",
				data)
		}

		if id, ok = f.(*goparser.FieldSpec); !ok {
			return in, fmt.Errorf("Unable to prepare. Wrong Id type: %#v", f)
		}

		// override the field spec
		data["Id"] = &UpperFieldSpec{
			id,
		}
		return data, nil
	})
}

// UpperFieldSpec represents information of an upper field.
// Includes its original field spec and database type suggestion
type UpperFieldSpec struct {
	*goparser.FieldSpec
}

func (s UpperFieldSpec) IsString() bool {
	return s.Type == "string"
}

func (s *UpperFieldSpec) IsInt() bool {
	return s.Type == "int" || s.Type == "uint" ||
		s.Type == "int32" || s.Type == "uint32" ||
		s.Type == "int64" || s.Type == "uint64"
}
