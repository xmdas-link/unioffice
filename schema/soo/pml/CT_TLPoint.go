// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package pml

import (
	"encoding/xml"
	"fmt"

	"github.com/xmdas-link/unioffice/schema/soo/dml"
)

type CT_TLPoint struct {
	// X coordinate
	XAttr dml.ST_Percentage
	// Y coordinate
	YAttr dml.ST_Percentage
}

func NewCT_TLPoint() *CT_TLPoint {
	ret := &CT_TLPoint{}
	return ret
}

func (m *CT_TLPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "x"},
		Value: fmt.Sprintf("%v", m.XAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "y"},
		Value: fmt.Sprintf("%v", m.YAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "x" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.XAttr = parsed
			continue
		}
		if attr.Name.Local == "y" {
			parsed, err := ParseUnionST_Percentage(attr.Value)
			if err != nil {
				return err
			}
			m.YAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_TLPoint: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_TLPoint and its children
func (m *CT_TLPoint) Validate() error {
	return m.ValidateWithPath("CT_TLPoint")
}

// ValidateWithPath validates the CT_TLPoint and its children, prefixing error messages with path
func (m *CT_TLPoint) ValidateWithPath(path string) error {
	if err := m.XAttr.ValidateWithPath(path + "/XAttr"); err != nil {
		return err
	}
	if err := m.YAttr.ValidateWithPath(path + "/YAttr"); err != nil {
		return err
	}
	return nil
}
