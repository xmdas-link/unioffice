// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package spreadsheetDrawing

import (
	"encoding/xml"

	"github.com/xmdas-link/unioffice"
)

type CT_TwoCellAnchor struct {
	EditAsAttr ST_EditAs
	From       *CT_Marker
	To         *CT_Marker
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

func NewCT_TwoCellAnchor() *CT_TwoCellAnchor {
	ret := &CT_TwoCellAnchor{}
	ret.From = NewCT_Marker()
	ret.To = NewCT_Marker()
	ret.ClientData = NewCT_AnchorClientData()
	return ret
}

func (m *CT_TwoCellAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.EditAsAttr != ST_EditAsUnset {
		attr, err := m.EditAsAttr.MarshalXMLAttr(xml.Name{Local: "editAs"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	sefrom := xml.StartElement{Name: xml.Name{Local: "xdr:from"}}
	e.EncodeElement(m.From, sefrom)
	seto := xml.StartElement{Name: xml.Name{Local: "xdr:to"}}
	e.EncodeElement(m.To, seto)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	seclientData := xml.StartElement{Name: xml.Name{Local: "xdr:clientData"}}
	e.EncodeElement(m.ClientData, seclientData)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TwoCellAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = NewCT_Marker()
	m.To = NewCT_Marker()
	m.ClientData = NewCT_AnchorClientData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "editAs" {
			m.EditAsAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_TwoCellAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "from"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "from"}:
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "to"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "to"}:
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "sp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "sp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "grpSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "grpSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "graphicFrame"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "graphicFrame"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cxnSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cxnSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "pic"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "contentPart"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.ContentPart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "clientData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "clientData"}:
				if err := d.DecodeElement(m.ClientData, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_TwoCellAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TwoCellAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TwoCellAnchor and its children
func (m *CT_TwoCellAnchor) Validate() error {
	return m.ValidateWithPath("CT_TwoCellAnchor")
}

// ValidateWithPath validates the CT_TwoCellAnchor and its children, prefixing error messages with path
func (m *CT_TwoCellAnchor) ValidateWithPath(path string) error {
	if err := m.EditAsAttr.ValidateWithPath(path + "/EditAsAttr"); err != nil {
		return err
	}
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.To.ValidateWithPath(path + "/To"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if err := m.ClientData.ValidateWithPath(path + "/ClientData"); err != nil {
		return err
	}
	return nil
}
