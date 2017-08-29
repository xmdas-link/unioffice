// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml

import (
	"encoding/xml"
	"log"
)

type CT_TLBehaviorAttributeNameList struct {
	// Attribute Name
	AttrName []string
}

func NewCT_TLBehaviorAttributeNameList() *CT_TLBehaviorAttributeNameList {
	ret := &CT_TLBehaviorAttributeNameList{}
	return ret
}
func (m *CT_TLBehaviorAttributeNameList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	e.EncodeToken(start)
	start.Attr = nil
	seattrName := xml.StartElement{Name: xml.Name{Local: "p:attrName"}}
	e.EncodeElement(m.AttrName, seattrName)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_TLBehaviorAttributeNameList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TLBehaviorAttributeNameList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "attrName":
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.AttrName = append(m.AttrName, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLBehaviorAttributeNameList
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_TLBehaviorAttributeNameList) Validate() error {
	return m.ValidateWithPath("CT_TLBehaviorAttributeNameList")
}
func (m *CT_TLBehaviorAttributeNameList) ValidateWithPath(path string) error {
	return nil
}