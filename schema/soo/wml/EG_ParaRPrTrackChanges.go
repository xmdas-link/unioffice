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

type EG_ParaRPrTrackChanges struct {
	// Inserted Paragraph
	Ins *CT_TrackChange
	// Deleted Paragraph
	Del *CT_TrackChange
	// Move Source Paragraph
	MoveFrom *CT_TrackChange
	// Move Destination Paragraph
	MoveTo *CT_TrackChange
}

func NewEG_ParaRPrTrackChanges() *EG_ParaRPrTrackChanges {
	ret := &EG_ParaRPrTrackChanges{}
	return ret
}

func (m *EG_ParaRPrTrackChanges) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.Ins != nil {
		seins := xml.StartElement{Name: xml.Name{Local: "w:ins"}}
		e.EncodeElement(m.Ins, seins)
	}
	if m.Del != nil {
		sedel := xml.StartElement{Name: xml.Name{Local: "w:del"}}
		e.EncodeElement(m.Del, sedel)
	}
	if m.MoveFrom != nil {
		semoveFrom := xml.StartElement{Name: xml.Name{Local: "w:moveFrom"}}
		e.EncodeElement(m.MoveFrom, semoveFrom)
	}
	if m.MoveTo != nil {
		semoveTo := xml.StartElement{Name: xml.Name{Local: "w:moveTo"}}
		e.EncodeElement(m.MoveTo, semoveTo)
	}
	return nil
}

func (m *EG_ParaRPrTrackChanges) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_ParaRPrTrackChanges:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "ins"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "ins"}:
				m.Ins = NewCT_TrackChange()
				if err := d.DecodeElement(m.Ins, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "del"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "del"}:
				m.Del = NewCT_TrackChange()
				if err := d.DecodeElement(m.Del, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveFrom"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveFrom"}:
				m.MoveFrom = NewCT_TrackChange()
				if err := d.DecodeElement(m.MoveFrom, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "moveTo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "moveTo"}:
				m.MoveTo = NewCT_TrackChange()
				if err := d.DecodeElement(m.MoveTo, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_ParaRPrTrackChanges %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_ParaRPrTrackChanges
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_ParaRPrTrackChanges and its children
func (m *EG_ParaRPrTrackChanges) Validate() error {
	return m.ValidateWithPath("EG_ParaRPrTrackChanges")
}

// ValidateWithPath validates the EG_ParaRPrTrackChanges and its children, prefixing error messages with path
func (m *EG_ParaRPrTrackChanges) ValidateWithPath(path string) error {
	if m.Ins != nil {
		if err := m.Ins.ValidateWithPath(path + "/Ins"); err != nil {
			return err
		}
	}
	if m.Del != nil {
		if err := m.Del.ValidateWithPath(path + "/Del"); err != nil {
			return err
		}
	}
	if m.MoveFrom != nil {
		if err := m.MoveFrom.ValidateWithPath(path + "/MoveFrom"); err != nil {
			return err
		}
	}
	if m.MoveTo != nil {
		if err := m.MoveTo.ValidateWithPath(path + "/MoveTo"); err != nil {
			return err
		}
	}
	return nil
}
