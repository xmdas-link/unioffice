// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package document

import (
	"baliance.com/gooxml"
	"baliance.com/gooxml/measurement"
	pic "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/picture"
	wd "baliance.com/gooxml/schema/schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
)

// AnchoredDrawing is an absolutely positioned image within a document page.
type AnchoredDrawing struct {
	d *Document
	x *wd.Anchor
}

// X returns the inner wrapped XML type.
func (a AnchoredDrawing) X() *wd.Anchor {
	return a.x
}

// SetName sets the name of the image, visible in the properties of the image
// within Word.
func (a AnchoredDrawing) SetName(name string) {
	a.x.DocPr.NameAttr = name
	for _, a := range a.x.Graphic.GraphicData.Any {
		if p, ok := a.(*pic.Pic); ok {
			p.NvPicPr.CNvPr.DescrAttr = gooxml.String(name)
		}
	}
}

// SetOrigin sets the origin of the image.  It defaults to ST_RelFromHPage and
// ST_RelFromVPage
func (a AnchoredDrawing) SetOrigin(h wd.ST_RelFromH, v wd.ST_RelFromV) {
	a.x.PositionH.RelativeFromAttr = h
	a.x.PositionV.RelativeFromAttr = v
}

// SetOffset sets the offset of the image relative to the origin, which by
// default this is the top-left corner of the page. Offset is incompatible with
// SetAlignment, whichever is called last is applied.
func (a AnchoredDrawing) SetOffset(x, y measurement.Distance) {
	a.SetXOffset(x)
	a.SetYOffset(y)
}

// SetXOffset sets the X offset for an image relative to the origin.
func (a AnchoredDrawing) SetXOffset(x measurement.Distance) {
	a.x.PositionH.Choice = &wd.CT_PosHChoice{}
	a.x.PositionH.Choice.PosOffset = gooxml.Int32(int32(x / measurement.EMU))
}

// SetYOffset sets the Y offset for an image relative to the origin.
func (a AnchoredDrawing) SetYOffset(y measurement.Distance) {
	a.x.PositionV.Choice = &wd.CT_PosVChoice{}
	a.x.PositionV.Choice.PosOffset = gooxml.Int32(int32(y / measurement.EMU))
}

// SetAlignment positions an anchored image via alignment.  Offset is
// incompatible with SetOffset, whichever is called last is applied.
func (a AnchoredDrawing) SetAlignment(h wd.ST_AlignH, v wd.ST_AlignV) {
	a.SetHAlignment(h)
	a.SetVAlignment(v)
}

func (a AnchoredDrawing) SetHAlignment(h wd.ST_AlignH) {
	a.x.PositionH.Choice = &wd.CT_PosHChoice{}
	a.x.PositionH.Choice.Align = h
}

func (a AnchoredDrawing) SetVAlignment(v wd.ST_AlignV) {
	a.x.PositionV.Choice = &wd.CT_PosVChoice{}
	a.x.PositionV.Choice.Align = v
}

// SetSize sets the size of the displayed image on the page.
func (a AnchoredDrawing) SetSize(w, h measurement.Distance) {
	a.x.Extent.CxAttr = int64(float64(w*measurement.Pixel72) / measurement.EMU)
	a.x.Extent.CyAttr = int64(float64(h*measurement.Pixel72) / measurement.EMU)
}

// SetTextWrapNone unsets text wrapping so the image can float on top of the
// text. When used in conjunction with X/Y Offset relative to the page it can be
// used to place a logo at the top of a page at an absolute position that
// doesn't interfere with text.
func (a AnchoredDrawing) SetTextWrapNone() {
	a.x.Choice = &wd.EG_WrapTypeChoice{}
	a.x.Choice.WrapNone = wd.NewCT_WrapNone()
}

// SetTextWrapSquare sets the text wrap to square with a given wrap type.
func (a AnchoredDrawing) SetTextWrapSquare(t wd.ST_WrapText) {
	a.x.Choice = &wd.EG_WrapTypeChoice{}
	a.x.Choice.WrapSquare = wd.NewCT_WrapSquare()
	a.x.Choice.WrapSquare.WrapTextAttr = t
}