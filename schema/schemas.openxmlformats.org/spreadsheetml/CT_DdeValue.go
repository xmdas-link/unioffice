// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"log"

	"baliance.com/gooxml"
)

type CT_DdeValue struct {
	// DDE Value Type
	TAttr ST_DdeValueType
	// DDE Link Value
	Val string
}

func NewCT_DdeValue() *CT_DdeValue {
	ret := &CT_DdeValue{}
	return ret
}
func (m *CT_DdeValue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TAttr != ST_DdeValueTypeUnset {
		attr, err := m.TAttr.MarshalXMLAttr(xml.Name{Local: "t"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	start.Attr = nil
	seval := xml.StartElement{Name: xml.Name{Local: "x:val"}}
	gooxml.AddPreserveSpaceAttr(&seval, m.Val)
	e.EncodeElement(m.Val, seval)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DdeValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "t" {
			m.TAttr.UnmarshalXMLAttr(attr)
		}
	}
lCT_DdeValue:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "val":
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DdeValue
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DdeValue) Validate() error {
	return m.ValidateWithPath("CT_DdeValue")
}
func (m *CT_DdeValue) ValidateWithPath(path string) error {
	if err := m.TAttr.ValidateWithPath(path + "/TAttr"); err != nil {
		return err
	}
	return nil
}