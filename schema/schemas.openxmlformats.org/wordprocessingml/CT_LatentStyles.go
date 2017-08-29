// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wordprocessingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"baliance.com/gooxml/schema/schemas.openxmlformats.org/officeDocument/2006/sharedTypes"
)

type CT_LatentStyles struct {
	// Default Style Locking Setting
	DefLockedStateAttr *sharedTypes.ST_OnOff
	// Default User Interface Priority Setting
	DefUIPriorityAttr *int32
	// Default Semi-Hidden Setting
	DefSemiHiddenAttr *sharedTypes.ST_OnOff
	// Default Hidden Until Used Setting
	DefUnhideWhenUsedAttr *sharedTypes.ST_OnOff
	// Default Primary Style Setting
	DefQFormatAttr *sharedTypes.ST_OnOff
	// Latent Style Count
	CountAttr *int32
	// Latent Style Exception
	LsdException []*CT_LsdException
}

func NewCT_LatentStyles() *CT_LatentStyles {
	ret := &CT_LatentStyles{}
	return ret
}
func (m *CT_LatentStyles) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.DefLockedStateAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:defLockedState"},
			Value: fmt.Sprintf("%v", *m.DefLockedStateAttr)})
	}
	if m.DefUIPriorityAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:defUIPriority"},
			Value: fmt.Sprintf("%v", *m.DefUIPriorityAttr)})
	}
	if m.DefSemiHiddenAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:defSemiHidden"},
			Value: fmt.Sprintf("%v", *m.DefSemiHiddenAttr)})
	}
	if m.DefUnhideWhenUsedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:defUnhideWhenUsed"},
			Value: fmt.Sprintf("%v", *m.DefUnhideWhenUsedAttr)})
	}
	if m.DefQFormatAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:defQFormat"},
			Value: fmt.Sprintf("%v", *m.DefQFormatAttr)})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w:count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.LsdException != nil {
		selsdException := xml.StartElement{Name: xml.Name{Local: "w:lsdException"}}
		e.EncodeElement(m.LsdException, selsdException)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_LatentStyles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "defLockedState" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DefLockedStateAttr = &parsed
		}
		if attr.Name.Local == "defUIPriority" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.DefUIPriorityAttr = &pt
		}
		if attr.Name.Local == "defSemiHidden" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DefSemiHiddenAttr = &parsed
		}
		if attr.Name.Local == "defUnhideWhenUsed" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DefUnhideWhenUsedAttr = &parsed
		}
		if attr.Name.Local == "defQFormat" {
			parsed, err := ParseUnionST_OnOff(attr.Value)
			if err != nil {
				return err
			}
			m.DefQFormatAttr = &parsed
		}
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.CountAttr = &pt
		}
	}
lCT_LatentStyles:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "lsdException":
				tmp := NewCT_LsdException()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.LsdException = append(m.LsdException, tmp)
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LatentStyles
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_LatentStyles) Validate() error {
	return m.ValidateWithPath("CT_LatentStyles")
}
func (m *CT_LatentStyles) ValidateWithPath(path string) error {
	if m.DefLockedStateAttr != nil {
		if err := m.DefLockedStateAttr.ValidateWithPath(path + "/DefLockedStateAttr"); err != nil {
			return err
		}
	}
	if m.DefSemiHiddenAttr != nil {
		if err := m.DefSemiHiddenAttr.ValidateWithPath(path + "/DefSemiHiddenAttr"); err != nil {
			return err
		}
	}
	if m.DefUnhideWhenUsedAttr != nil {
		if err := m.DefUnhideWhenUsedAttr.ValidateWithPath(path + "/DefUnhideWhenUsedAttr"); err != nil {
			return err
		}
	}
	if m.DefQFormatAttr != nil {
		if err := m.DefQFormatAttr.ValidateWithPath(path + "/DefQFormatAttr"); err != nil {
			return err
		}
	}
	for i, v := range m.LsdException {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/LsdException[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}