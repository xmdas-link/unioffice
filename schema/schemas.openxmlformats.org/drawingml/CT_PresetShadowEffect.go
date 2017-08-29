// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package drawingml

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
)

type CT_PresetShadowEffect struct {
	PrstAttr  ST_PresetShadowVal
	DistAttr  *int64
	DirAttr   *int32
	ScrgbClr  *CT_ScRgbColor
	SrgbClr   *CT_SRgbColor
	HslClr    *CT_HslColor
	SysClr    *CT_SystemColor
	SchemeClr *CT_SchemeColor
	PrstClr   *CT_PresetColor
}

func NewCT_PresetShadowEffect() *CT_PresetShadowEffect {
	ret := &CT_PresetShadowEffect{}
	return ret
}
func (m *CT_PresetShadowEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	attr, err := m.PrstAttr.MarshalXMLAttr(xml.Name{Local: "prst"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	if m.DistAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dist"},
			Value: fmt.Sprintf("%v", *m.DistAttr)})
	}
	if m.DirAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dir"},
			Value: fmt.Sprintf("%v", *m.DirAttr)})
	}
	e.EncodeToken(start)
	start.Attr = nil
	if m.ScrgbClr != nil {
		sescrgbClr := xml.StartElement{Name: xml.Name{Local: "a:scrgbClr"}}
		e.EncodeElement(m.ScrgbClr, sescrgbClr)
	}
	if m.SrgbClr != nil {
		sesrgbClr := xml.StartElement{Name: xml.Name{Local: "a:srgbClr"}}
		e.EncodeElement(m.SrgbClr, sesrgbClr)
	}
	if m.HslClr != nil {
		sehslClr := xml.StartElement{Name: xml.Name{Local: "a:hslClr"}}
		e.EncodeElement(m.HslClr, sehslClr)
	}
	if m.SysClr != nil {
		sesysClr := xml.StartElement{Name: xml.Name{Local: "a:sysClr"}}
		e.EncodeElement(m.SysClr, sesysClr)
	}
	if m.SchemeClr != nil {
		seschemeClr := xml.StartElement{Name: xml.Name{Local: "a:schemeClr"}}
		e.EncodeElement(m.SchemeClr, seschemeClr)
	}
	if m.PrstClr != nil {
		seprstClr := xml.StartElement{Name: xml.Name{Local: "a:prstClr"}}
		e.EncodeElement(m.PrstClr, seprstClr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_PresetShadowEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "prst" {
			m.PrstAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "dist" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DistAttr = &parsed
		}
		if attr.Name.Local == "dir" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			// SPECIAL
			pt := int32(parsed)
			m.DirAttr = &pt
		}
	}
lCT_PresetShadowEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "scrgbClr":
				m.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(m.ScrgbClr, &el); err != nil {
					return err
				}
			case "srgbClr":
				m.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(m.SrgbClr, &el); err != nil {
					return err
				}
			case "hslClr":
				m.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(m.HslClr, &el); err != nil {
					return err
				}
			case "sysClr":
				m.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(m.SysClr, &el); err != nil {
					return err
				}
			case "schemeClr":
				m.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(m.SchemeClr, &el); err != nil {
					return err
				}
			case "prstClr":
				m.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(m.PrstClr, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PresetShadowEffect
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_PresetShadowEffect) Validate() error {
	return m.ValidateWithPath("CT_PresetShadowEffect")
}
func (m *CT_PresetShadowEffect) ValidateWithPath(path string) error {
	if m.PrstAttr == ST_PresetShadowValUnset {
		return fmt.Errorf("%s/PrstAttr is a mandatory field", path)
	}
	if err := m.PrstAttr.ValidateWithPath(path + "/PrstAttr"); err != nil {
		return err
	}
	if m.DistAttr != nil {
		if *m.DistAttr < 0 {
			return fmt.Errorf("%s/m.DistAttr must be >= 0 (have %v)", path, *m.DistAttr)
		}
		if *m.DistAttr > 27273042316900 {
			return fmt.Errorf("%s/m.DistAttr must be <= 27273042316900 (have %v)", path, *m.DistAttr)
		}
	}
	if m.DirAttr != nil {
		if *m.DirAttr < 0 {
			return fmt.Errorf("%s/m.DirAttr must be >= 0 (have %v)", path, *m.DirAttr)
		}
		if *m.DirAttr >= 21600000 {
			return fmt.Errorf("%s/m.DirAttr must be < 21600000 (have %v)", path, *m.DirAttr)
		}
	}
	if m.ScrgbClr != nil {
		if err := m.ScrgbClr.ValidateWithPath(path + "/ScrgbClr"); err != nil {
			return err
		}
	}
	if m.SrgbClr != nil {
		if err := m.SrgbClr.ValidateWithPath(path + "/SrgbClr"); err != nil {
			return err
		}
	}
	if m.HslClr != nil {
		if err := m.HslClr.ValidateWithPath(path + "/HslClr"); err != nil {
			return err
		}
	}
	if m.SysClr != nil {
		if err := m.SysClr.ValidateWithPath(path + "/SysClr"); err != nil {
			return err
		}
	}
	if m.SchemeClr != nil {
		if err := m.SchemeClr.ValidateWithPath(path + "/SchemeClr"); err != nil {
			return err
		}
	}
	if m.PrstClr != nil {
		if err := m.PrstClr.ValidateWithPath(path + "/PrstClr"); err != nil {
			return err
		}
	}
	return nil
}