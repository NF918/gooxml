// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/NF918/gooxml"
	"github.com/NF918/gooxml/schema/soo/dml"
)

type CT_ColorTransformHeader struct {
	UniqueIdAttr string
	MinVerAttr   *string
	ResIdAttr    *int32
	Title        []*CT_CTName
	Desc         []*CT_CTDescription
	CatLst       *CT_CTCategories
	ExtLst       *dml.CT_OfficeArtExtensionList
}

func NewCT_ColorTransformHeader() *CT_ColorTransformHeader {
	ret := &CT_ColorTransformHeader{}
	return ret
}

func (m *CT_ColorTransformHeader) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueId"},
		Value: fmt.Sprintf("%v", m.UniqueIdAttr)})
	if m.MinVerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minVer"},
			Value: fmt.Sprintf("%v", *m.MinVerAttr)})
	}
	if m.ResIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "resId"},
			Value: fmt.Sprintf("%v", *m.ResIdAttr)})
	}
	e.EncodeToken(start)
	setitle := xml.StartElement{Name: xml.Name{Local: "title"}}
	for _, c := range m.Title {
		e.EncodeElement(c, setitle)
	}
	sedesc := xml.StartElement{Name: xml.Name{Local: "desc"}}
	for _, c := range m.Desc {
		e.EncodeElement(c, sedesc)
	}
	if m.CatLst != nil {
		secatLst := xml.StartElement{Name: xml.Name{Local: "catLst"}}
		e.EncodeElement(m.CatLst, secatLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorTransformHeader) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueIdAttr = parsed
			continue
		}
		if attr.Name.Local == "minVer" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MinVerAttr = &parsed
			continue
		}
		if attr.Name.Local == "resId" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.ResIdAttr = &pt
			continue
		}
	}
lCT_ColorTransformHeader:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "title"}:
				tmp := NewCT_CTName()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Title = append(m.Title, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "desc"}:
				tmp := NewCT_CTDescription()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Desc = append(m.Desc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "catLst"}:
				m.CatLst = NewCT_CTCategories()
				if err := d.DecodeElement(m.CatLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ColorTransformHeader %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorTransformHeader
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorTransformHeader and its children
func (m *CT_ColorTransformHeader) Validate() error {
	return m.ValidateWithPath("CT_ColorTransformHeader")
}

// ValidateWithPath validates the CT_ColorTransformHeader and its children, prefixing error messages with path
func (m *CT_ColorTransformHeader) ValidateWithPath(path string) error {
	for i, v := range m.Title {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Title[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Desc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Desc[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CatLst != nil {
		if err := m.CatLst.ValidateWithPath(path + "/CatLst"); err != nil {
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
