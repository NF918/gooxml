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

func TestCT_MailMergeDestConstructor(t *testing.T) {
	v := wml.NewCT_MailMergeDest()
	if v == nil {
		t.Errorf("wml.NewCT_MailMergeDest must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MailMergeDest should validate: %s", err)
	}
}

func TestCT_MailMergeDestMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MailMergeDest()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MailMergeDest()
	xml.Unmarshal(buf, v2)
}
