// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"github.com/NF918/gooxml/common"
	"github.com/NF918/gooxml/measurement"
	pic "github.com/NF918/gooxml/schema/soo/dml/picture"
	"github.com/NF918/gooxml/schema/soo/wml"
)

// InlineDrawing is an inlined image within a run.
type InlineDrawing struct {
	d *Document
	x *wml.WdInline
}

// X returns the inner wrapped XML type.
func (i InlineDrawing) X() *wml.WdInline {
	return i.x
}

// GetImage returns the ImageRef associated with an InlineDrawing.
func (i InlineDrawing) GetImage() (common.ImageRef, bool) {
	any := i.x.Graphic.GraphicData.Any
	if len(any) > 0 {
		p, ok := any[0].(*pic.Pic)
		if ok {
			if p.BlipFill != nil && p.BlipFill.Blip != nil && p.BlipFill.Blip.EmbedAttr != nil {
				return i.d.GetImageByRelID(*p.BlipFill.Blip.EmbedAttr)
			}
		}
	}
	return common.ImageRef{}, false
}

// SetSize sets the size of the displayed image on the page.
func (i InlineDrawing) SetSize(w, h measurement.Distance) {
	i.x.Extent.CxAttr = int64(float64(w*measurement.Pixel72) / measurement.EMU)
	i.x.Extent.CyAttr = int64(float64(h*measurement.Pixel72) / measurement.EMU)
}
