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

	"github.com/xmdas-link/unioffice"
)

type CT_TblBorders struct {
	// Table Top Border
	Top *CT_Border
	// Table Leading Edge Border
	Start *CT_Border
	// Table Leading Edge Border
	Left *CT_Border
	// Table Bottom Border
	Bottom *CT_Border
	// Table Trailing Edge Border
	End *CT_Border
	// Table Trailing Edge Border
	Right *CT_Border
	// Table Inside Horizontal Edges Border
	InsideH *CT_Border
	// Table Inside Vertical Edges Border
	InsideV *CT_Border
}

func NewCT_TblBorders() *CT_TblBorders {
	ret := &CT_TblBorders{}
	return ret
}

func (m *CT_TblBorders) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Top != nil {
		setop := xml.StartElement{Name: xml.Name{Local: "w:top"}}
		e.EncodeElement(m.Top, setop)
	}
	if m.Start != nil {
		sestart := xml.StartElement{Name: xml.Name{Local: "w:start"}}
		e.EncodeElement(m.Start, sestart)
	}
	if m.Left != nil {
		seleft := xml.StartElement{Name: xml.Name{Local: "w:left"}}
		e.EncodeElement(m.Left, seleft)
	}
	if m.Bottom != nil {
		sebottom := xml.StartElement{Name: xml.Name{Local: "w:bottom"}}
		e.EncodeElement(m.Bottom, sebottom)
	}
	if m.End != nil {
		seend := xml.StartElement{Name: xml.Name{Local: "w:end"}}
		e.EncodeElement(m.End, seend)
	}
	if m.Right != nil {
		seright := xml.StartElement{Name: xml.Name{Local: "w:right"}}
		e.EncodeElement(m.Right, seright)
	}
	if m.InsideH != nil {
		seinsideH := xml.StartElement{Name: xml.Name{Local: "w:insideH"}}
		e.EncodeElement(m.InsideH, seinsideH)
	}
	if m.InsideV != nil {
		seinsideV := xml.StartElement{Name: xml.Name{Local: "w:insideV"}}
		e.EncodeElement(m.InsideV, seinsideV)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblBorders) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TblBorders:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "top"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "top"}:
				m.Top = NewCT_Border()
				if err := d.DecodeElement(m.Top, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "start"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "start"}:
				m.Start = NewCT_Border()
				if err := d.DecodeElement(m.Start, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "left"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "left"}:
				m.Left = NewCT_Border()
				if err := d.DecodeElement(m.Left, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bottom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bottom"}:
				m.Bottom = NewCT_Border()
				if err := d.DecodeElement(m.Bottom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "end"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "end"}:
				m.End = NewCT_Border()
				if err := d.DecodeElement(m.End, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "right"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "right"}:
				m.Right = NewCT_Border()
				if err := d.DecodeElement(m.Right, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "insideH"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "insideH"}:
				m.InsideH = NewCT_Border()
				if err := d.DecodeElement(m.InsideH, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "insideV"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "insideV"}:
				m.InsideV = NewCT_Border()
				if err := d.DecodeElement(m.InsideV, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TblBorders %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblBorders
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblBorders and its children
func (m *CT_TblBorders) Validate() error {
	return m.ValidateWithPath("CT_TblBorders")
}

// ValidateWithPath validates the CT_TblBorders and its children, prefixing error messages with path
func (m *CT_TblBorders) ValidateWithPath(path string) error {
	if m.Top != nil {
		if err := m.Top.ValidateWithPath(path + "/Top"); err != nil {
			return err
		}
	}
	if m.Start != nil {
		if err := m.Start.ValidateWithPath(path + "/Start"); err != nil {
			return err
		}
	}
	if m.Left != nil {
		if err := m.Left.ValidateWithPath(path + "/Left"); err != nil {
			return err
		}
	}
	if m.Bottom != nil {
		if err := m.Bottom.ValidateWithPath(path + "/Bottom"); err != nil {
			return err
		}
	}
	if m.End != nil {
		if err := m.End.ValidateWithPath(path + "/End"); err != nil {
			return err
		}
	}
	if m.Right != nil {
		if err := m.Right.ValidateWithPath(path + "/Right"); err != nil {
			return err
		}
	}
	if m.InsideH != nil {
		if err := m.InsideH.ValidateWithPath(path + "/InsideH"); err != nil {
			return err
		}
	}
	if m.InsideV != nil {
		if err := m.InsideV.ValidateWithPath(path + "/InsideV"); err != nil {
			return err
		}
	}
	return nil
}
