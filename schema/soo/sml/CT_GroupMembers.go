// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/xmdas-link/unioffice"
)

type CT_GroupMembers struct {
	// Group Member Count
	CountAttr *uint32
	// OLAP Group Member
	GroupMember []*CT_GroupMember
}

func NewCT_GroupMembers() *CT_GroupMembers {
	ret := &CT_GroupMembers{}
	return ret
}

func (m *CT_GroupMembers) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	segroupMember := xml.StartElement{Name: xml.Name{Local: "ma:groupMember"}}
	for _, c := range m.GroupMember {
		e.EncodeElement(c, segroupMember)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupMembers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_GroupMembers:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "groupMember"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "groupMember"}:
				tmp := NewCT_GroupMember()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GroupMember = append(m.GroupMember, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_GroupMembers %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupMembers
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupMembers and its children
func (m *CT_GroupMembers) Validate() error {
	return m.ValidateWithPath("CT_GroupMembers")
}

// ValidateWithPath validates the CT_GroupMembers and its children, prefixing error messages with path
func (m *CT_GroupMembers) ValidateWithPath(path string) error {
	for i, v := range m.GroupMember {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GroupMember[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
