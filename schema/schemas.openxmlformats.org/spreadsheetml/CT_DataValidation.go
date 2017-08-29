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
	"log"
	"strconv"

	"baliance.com/gooxml"
)

type CT_DataValidation struct {
	// Data Validation Type
	TypeAttr ST_DataValidationType
	// Data Validation Error Style
	ErrorStyleAttr ST_DataValidationErrorStyle
	// IME Mode Enforced
	ImeModeAttr ST_DataValidationImeMode
	// Operator
	OperatorAttr ST_DataValidationOperator
	// Allow Blank
	AllowBlankAttr *bool
	// Show Drop Down
	ShowDropDownAttr *bool
	// Show Input Message
	ShowInputMessageAttr *bool
	// Show Error Message
	ShowErrorMessageAttr *bool
	// Error Alert Text
	ErrorTitleAttr *string
	// Error Message
	ErrorAttr *string
	// Prompt Title
	PromptTitleAttr *string
	// Input Prompt
	PromptAttr *string
	// Sequence of References
	SqrefAttr ST_Sqref
	// Formula 1
	Formula1 *string
	// Formula 2
	Formula2 *string
}

func NewCT_DataValidation() *CT_DataValidation {
	ret := &CT_DataValidation{}
	return ret
}
func (m *CT_DataValidation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m == nil {
		return nil
	}
	if m.TypeAttr != ST_DataValidationTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ErrorStyleAttr != ST_DataValidationErrorStyleUnset {
		attr, err := m.ErrorStyleAttr.MarshalXMLAttr(xml.Name{Local: "errorStyle"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ImeModeAttr != ST_DataValidationImeModeUnset {
		attr, err := m.ImeModeAttr.MarshalXMLAttr(xml.Name{Local: "imeMode"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.OperatorAttr != ST_DataValidationOperatorUnset {
		attr, err := m.OperatorAttr.MarshalXMLAttr(xml.Name{Local: "operator"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AllowBlankAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "allowBlank"},
			Value: fmt.Sprintf("%v", *m.AllowBlankAttr)})
	}
	if m.ShowDropDownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showDropDown"},
			Value: fmt.Sprintf("%v", *m.ShowDropDownAttr)})
	}
	if m.ShowInputMessageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showInputMessage"},
			Value: fmt.Sprintf("%v", *m.ShowInputMessageAttr)})
	}
	if m.ShowErrorMessageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showErrorMessage"},
			Value: fmt.Sprintf("%v", *m.ShowErrorMessageAttr)})
	}
	if m.ErrorTitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "errorTitle"},
			Value: fmt.Sprintf("%v", *m.ErrorTitleAttr)})
	}
	if m.ErrorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "error"},
			Value: fmt.Sprintf("%v", *m.ErrorAttr)})
	}
	if m.PromptTitleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "promptTitle"},
			Value: fmt.Sprintf("%v", *m.PromptTitleAttr)})
	}
	if m.PromptAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "prompt"},
			Value: fmt.Sprintf("%v", *m.PromptAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sqref"},
		Value: fmt.Sprintf("%v", m.SqrefAttr)})
	e.EncodeToken(start)
	start.Attr = nil
	if m.Formula1 != nil {
		seformula1 := xml.StartElement{Name: xml.Name{Local: "x:formula1"}}
		gooxml.AddPreserveSpaceAttr(&seformula1, *m.Formula1)
		e.EncodeElement(m.Formula1, seformula1)
	}
	if m.Formula2 != nil {
		seformula2 := xml.StartElement{Name: xml.Name{Local: "x:formula2"}}
		gooxml.AddPreserveSpaceAttr(&seformula2, *m.Formula2)
		e.EncodeElement(m.Formula2, seformula2)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}
func (m *CT_DataValidation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "errorStyle" {
			m.ErrorStyleAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "imeMode" {
			m.ImeModeAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "operator" {
			m.OperatorAttr.UnmarshalXMLAttr(attr)
		}
		if attr.Name.Local == "allowBlank" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AllowBlankAttr = &parsed
		}
		if attr.Name.Local == "showDropDown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowDropDownAttr = &parsed
		}
		if attr.Name.Local == "showInputMessage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowInputMessageAttr = &parsed
		}
		if attr.Name.Local == "showErrorMessage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowErrorMessageAttr = &parsed
		}
		if attr.Name.Local == "errorTitle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ErrorTitleAttr = &parsed
		}
		if attr.Name.Local == "error" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ErrorAttr = &parsed
		}
		if attr.Name.Local == "promptTitle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PromptTitleAttr = &parsed
		}
		if attr.Name.Local == "prompt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PromptAttr = &parsed
		}
		if attr.Name.Local == "sqref" {
			parsed, err := ParseSliceST_Sqref(attr.Value)
			if err != nil {
				return err
			}
			m.SqrefAttr = parsed
		}
	}
lCT_DataValidation:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name.Local {
			case "formula1":
				m.Formula1 = new(string)
				if err := d.DecodeElement(m.Formula1, &el); err != nil {
					return err
				}
			case "formula2":
				m.Formula2 = new(string)
				if err := d.DecodeElement(m.Formula2, &el); err != nil {
					return err
				}
			default:
				log.Printf("skipping unsupported element %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DataValidation
		case xml.CharData:
		}
	}
	return nil
}
func (m *CT_DataValidation) Validate() error {
	return m.ValidateWithPath("CT_DataValidation")
}
func (m *CT_DataValidation) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.ErrorStyleAttr.ValidateWithPath(path + "/ErrorStyleAttr"); err != nil {
		return err
	}
	if err := m.ImeModeAttr.ValidateWithPath(path + "/ImeModeAttr"); err != nil {
		return err
	}
	if err := m.OperatorAttr.ValidateWithPath(path + "/OperatorAttr"); err != nil {
		return err
	}
	return nil
}