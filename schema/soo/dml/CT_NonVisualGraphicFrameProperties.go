// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"

	"github.com/NF918/gooxml"
)

type CT_NonVisualGraphicFrameProperties struct {
	GraphicFrameLocks *CT_GraphicalObjectFrameLocking
	ExtLst            *CT_OfficeArtExtensionList
}

func NewCT_NonVisualGraphicFrameProperties() *CT_NonVisualGraphicFrameProperties {
	ret := &CT_NonVisualGraphicFrameProperties{}
	return ret
}

func (m *CT_NonVisualGraphicFrameProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.GraphicFrameLocks != nil {
		segraphicFrameLocks := xml.StartElement{Name: xml.Name{Local: "a:graphicFrameLocks"}}
		e.EncodeElement(m.GraphicFrameLocks, segraphicFrameLocks)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_NonVisualGraphicFrameProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_NonVisualGraphicFrameProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "graphicFrameLocks"}:
				m.GraphicFrameLocks = NewCT_GraphicalObjectFrameLocking()
				if err := d.DecodeElement(m.GraphicFrameLocks, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_NonVisualGraphicFrameProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_NonVisualGraphicFrameProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_NonVisualGraphicFrameProperties and its children
func (m *CT_NonVisualGraphicFrameProperties) Validate() error {
	return m.ValidateWithPath("CT_NonVisualGraphicFrameProperties")
}

// ValidateWithPath validates the CT_NonVisualGraphicFrameProperties and its children, prefixing error messages with path
func (m *CT_NonVisualGraphicFrameProperties) ValidateWithPath(path string) error {
	if m.GraphicFrameLocks != nil {
		if err := m.GraphicFrameLocks.ValidateWithPath(path + "/GraphicFrameLocks"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
