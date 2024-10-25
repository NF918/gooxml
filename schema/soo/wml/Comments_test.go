// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/NF918/gooxml/schema/soo/wml"
)

func TestCommentsConstructor(t *testing.T) {
	v := wml.NewComments()
	if v == nil {
		t.Errorf("wml.NewComments must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.Comments should validate: %s", err)
	}
}

func TestCommentsMarshalUnmarshal(t *testing.T) {
	v := wml.NewComments()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewComments()
	xml.Unmarshal(buf, v2)
}
