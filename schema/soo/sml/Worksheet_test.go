// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml_test

import (
	"encoding/xml"
	"testing"

	"github.com/NF918/gooxml/schema/soo/sml"
)

func TestWorksheetConstructor(t *testing.T) {
	v := sml.NewWorksheet()
	if v == nil {
		t.Errorf("sml.NewWorksheet must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Worksheet should validate: %s", err)
	}
}

func TestWorksheetMarshalUnmarshal(t *testing.T) {
	v := sml.NewWorksheet()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewWorksheet()
	xml.Unmarshal(buf, v2)
}
