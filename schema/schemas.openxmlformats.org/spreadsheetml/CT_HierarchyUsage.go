// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_HierarchyUsage struct {
	// Hierarchy Usage
	HierarchyUsageAttr int32
}

func NewCT_HierarchyUsage() *CT_HierarchyUsage {
	ret := &CT_HierarchyUsage{}
	return ret
}
func (m *CT_HierarchyUsage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "hierarchyUsage"},
		Value: fmt.Sprintf("%v", m.HierarchyUsageAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_HierarchyUsage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "hierarchyUsage" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.HierarchyUsageAttr = int32(parsed)
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_HierarchyUsage: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}
func (m *CT_HierarchyUsage) Validate() error {
	return m.ValidateWithPath("CT_HierarchyUsage")
}
func (m *CT_HierarchyUsage) ValidateWithPath(path string) error {
	return nil
}