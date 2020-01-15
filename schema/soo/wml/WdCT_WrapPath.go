// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/xmdas-link/unioffice"
	"github.com/xmdas-link/unioffice/schema/soo/dml"
)

type WdCT_WrapPath struct {
	EditedAttr *bool
	Start      *dml.CT_Point2D
	LineTo     []*dml.CT_Point2D
}

func NewWdCT_WrapPath() *WdCT_WrapPath {
	ret := &WdCT_WrapPath{}
	ret.Start = dml.NewCT_Point2D()
	return ret
}

func (m *WdCT_WrapPath) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EditedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "edited"},
			Value: fmt.Sprintf("%d", b2i(*m.EditedAttr))})
	}
	e.EncodeToken(start)
	sestart := xml.StartElement{Name: xml.Name{Local: "wp:start"}}
	e.EncodeElement(m.Start, sestart)
	selineTo := xml.StartElement{Name: xml.Name{Local: "wp:lineTo"}}
	for _, c := range m.LineTo {
		e.EncodeElement(c, selineTo)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_WrapPath) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Start = dml.NewCT_Point2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "edited" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EditedAttr = &parsed
			continue
		}
	}
lWdCT_WrapPath:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "start"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "start"}:
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "lineTo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "lineTo"}:
				tmp := dml.NewCT_Point2D()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LineTo = append(m.LineTo, tmp)
			default:
				unioffice.Log("skipping unsupported element on WdCT_WrapPath %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_WrapPath
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_WrapPath and its children
func (m *WdCT_WrapPath) Validate() error {
	return m.ValidateWithPath("WdCT_WrapPath")
}

// ValidateWithPath validates the WdCT_WrapPath and its children, prefixing error messages with path
func (m *WdCT_WrapPath) ValidateWithPath(path string) error {
	if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
		return err
	}
	for i, v := range m.LineTo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LineTo[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
