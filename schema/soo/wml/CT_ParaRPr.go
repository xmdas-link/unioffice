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

type CT_ParaRPr struct {
	// Inserted Paragraph
	Ins *CT_TrackChange
	// Deleted Paragraph
	Del *CT_TrackChange
	// Move Source Paragraph
	MoveFrom *CT_TrackChange
	// Move Destination Paragraph
	MoveTo *CT_TrackChange
	// Referenced Character Style
	RStyle *CT_String
	// Run Fonts
	RFonts *CT_Fonts
	// Bold
	B *CT_OnOff
	// Complex Script Bold
	BCs *CT_OnOff
	// Italics
	I *CT_OnOff
	// Complex Script Italics
	ICs *CT_OnOff
	// Display All Characters As Capital Letters
	Caps *CT_OnOff
	// Small Caps
	SmallCaps *CT_OnOff
	// Single Strikethrough
	Strike *CT_OnOff
	// Double Strikethrough
	Dstrike *CT_OnOff
	// Display Character Outline
	Outline *CT_OnOff
	// Shadow
	Shadow *CT_OnOff
	// Embossing
	Emboss *CT_OnOff
	// Imprinting
	Imprint *CT_OnOff
	// Do Not Check Spelling or Grammar
	NoProof *CT_OnOff
	// Use Document Grid Settings For Inter-Character Spacing
	SnapToGrid *CT_OnOff
	// Hidden Text
	Vanish *CT_OnOff
	// Web Hidden Text
	WebHidden *CT_OnOff
	// Run Content Color
	Color *CT_Color
	// Character Spacing Adjustment
	Spacing *CT_SignedTwipsMeasure
	// Expanded/Compressed Text
	W *CT_TextScale
	// Font Kerning
	Kern *CT_HpsMeasure
	// Vertically Raised or Lowered Text
	Position *CT_SignedHpsMeasure
	// Non-Complex Script Font Size
	Sz *CT_HpsMeasure
	// Complex Script Font Size
	SzCs *CT_HpsMeasure
	// Text Highlighting
	Highlight *CT_Highlight
	// Underline
	U *CT_Underline
	// Animated Text Effect
	Effect *CT_TextEffect
	// Text Border
	Bdr *CT_Border
	// Run Shading
	Shd *CT_Shd
	// Manual Run Width
	FitText *CT_FitText
	// Subscript/Superscript Text
	VertAlign *CT_VerticalAlignRun
	// Right To Left Text
	Rtl *CT_OnOff
	// Use Complex Script Formatting on Run
	Cs *CT_OnOff
	// Emphasis Mark
	Em *CT_Em
	// Languages for Run Content
	Lang *CT_Language
	// East Asian Typography Settings
	EastAsianLayout *CT_EastAsianLayout
	// Paragraph Mark Is Always Hidden
	SpecVanish *CT_OnOff
	// Office Open XML Math
	OMath *CT_OnOff
	// Revision Information for Run Properties on the Paragraph Mark
	RPrChange *CT_ParaRPrChange
}

func NewCT_ParaRPr() *CT_ParaRPr {
	ret := &CT_ParaRPr{}
	return ret
}

func (m *CT_ParaRPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
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
	if m.RStyle != nil {
		serStyle := xml.StartElement{Name: xml.Name{Local: "w:rStyle"}}
		e.EncodeElement(m.RStyle, serStyle)
	}
	if m.RFonts != nil {
		serFonts := xml.StartElement{Name: xml.Name{Local: "w:rFonts"}}
		e.EncodeElement(m.RFonts, serFonts)
	}
	if m.B != nil {
		seb := xml.StartElement{Name: xml.Name{Local: "w:b"}}
		e.EncodeElement(m.B, seb)
	}
	if m.BCs != nil {
		sebCs := xml.StartElement{Name: xml.Name{Local: "w:bCs"}}
		e.EncodeElement(m.BCs, sebCs)
	}
	if m.I != nil {
		sei := xml.StartElement{Name: xml.Name{Local: "w:i"}}
		e.EncodeElement(m.I, sei)
	}
	if m.ICs != nil {
		seiCs := xml.StartElement{Name: xml.Name{Local: "w:iCs"}}
		e.EncodeElement(m.ICs, seiCs)
	}
	if m.Caps != nil {
		secaps := xml.StartElement{Name: xml.Name{Local: "w:caps"}}
		e.EncodeElement(m.Caps, secaps)
	}
	if m.SmallCaps != nil {
		sesmallCaps := xml.StartElement{Name: xml.Name{Local: "w:smallCaps"}}
		e.EncodeElement(m.SmallCaps, sesmallCaps)
	}
	if m.Strike != nil {
		sestrike := xml.StartElement{Name: xml.Name{Local: "w:strike"}}
		e.EncodeElement(m.Strike, sestrike)
	}
	if m.Dstrike != nil {
		sedstrike := xml.StartElement{Name: xml.Name{Local: "w:dstrike"}}
		e.EncodeElement(m.Dstrike, sedstrike)
	}
	if m.Outline != nil {
		seoutline := xml.StartElement{Name: xml.Name{Local: "w:outline"}}
		e.EncodeElement(m.Outline, seoutline)
	}
	if m.Shadow != nil {
		seshadow := xml.StartElement{Name: xml.Name{Local: "w:shadow"}}
		e.EncodeElement(m.Shadow, seshadow)
	}
	if m.Emboss != nil {
		seemboss := xml.StartElement{Name: xml.Name{Local: "w:emboss"}}
		e.EncodeElement(m.Emboss, seemboss)
	}
	if m.Imprint != nil {
		seimprint := xml.StartElement{Name: xml.Name{Local: "w:imprint"}}
		e.EncodeElement(m.Imprint, seimprint)
	}
	if m.NoProof != nil {
		senoProof := xml.StartElement{Name: xml.Name{Local: "w:noProof"}}
		e.EncodeElement(m.NoProof, senoProof)
	}
	if m.SnapToGrid != nil {
		sesnapToGrid := xml.StartElement{Name: xml.Name{Local: "w:snapToGrid"}}
		e.EncodeElement(m.SnapToGrid, sesnapToGrid)
	}
	if m.Vanish != nil {
		sevanish := xml.StartElement{Name: xml.Name{Local: "w:vanish"}}
		e.EncodeElement(m.Vanish, sevanish)
	}
	if m.WebHidden != nil {
		sewebHidden := xml.StartElement{Name: xml.Name{Local: "w:webHidden"}}
		e.EncodeElement(m.WebHidden, sewebHidden)
	}
	if m.Color != nil {
		secolor := xml.StartElement{Name: xml.Name{Local: "w:color"}}
		e.EncodeElement(m.Color, secolor)
	}
	if m.Spacing != nil {
		sespacing := xml.StartElement{Name: xml.Name{Local: "w:spacing"}}
		e.EncodeElement(m.Spacing, sespacing)
	}
	if m.W != nil {
		sew := xml.StartElement{Name: xml.Name{Local: "w:w"}}
		e.EncodeElement(m.W, sew)
	}
	if m.Kern != nil {
		sekern := xml.StartElement{Name: xml.Name{Local: "w:kern"}}
		e.EncodeElement(m.Kern, sekern)
	}
	if m.Position != nil {
		seposition := xml.StartElement{Name: xml.Name{Local: "w:position"}}
		e.EncodeElement(m.Position, seposition)
	}
	if m.Sz != nil {
		sesz := xml.StartElement{Name: xml.Name{Local: "w:sz"}}
		e.EncodeElement(m.Sz, sesz)
	}
	if m.SzCs != nil {
		seszCs := xml.StartElement{Name: xml.Name{Local: "w:szCs"}}
		e.EncodeElement(m.SzCs, seszCs)
	}
	if m.Highlight != nil {
		sehighlight := xml.StartElement{Name: xml.Name{Local: "w:highlight"}}
		e.EncodeElement(m.Highlight, sehighlight)
	}
	if m.U != nil {
		seu := xml.StartElement{Name: xml.Name{Local: "w:u"}}
		e.EncodeElement(m.U, seu)
	}
	if m.Effect != nil {
		seeffect := xml.StartElement{Name: xml.Name{Local: "w:effect"}}
		e.EncodeElement(m.Effect, seeffect)
	}
	if m.Bdr != nil {
		sebdr := xml.StartElement{Name: xml.Name{Local: "w:bdr"}}
		e.EncodeElement(m.Bdr, sebdr)
	}
	if m.Shd != nil {
		seshd := xml.StartElement{Name: xml.Name{Local: "w:shd"}}
		e.EncodeElement(m.Shd, seshd)
	}
	if m.FitText != nil {
		sefitText := xml.StartElement{Name: xml.Name{Local: "w:fitText"}}
		e.EncodeElement(m.FitText, sefitText)
	}
	if m.VertAlign != nil {
		severtAlign := xml.StartElement{Name: xml.Name{Local: "w:vertAlign"}}
		e.EncodeElement(m.VertAlign, severtAlign)
	}
	if m.Rtl != nil {
		sertl := xml.StartElement{Name: xml.Name{Local: "w:rtl"}}
		e.EncodeElement(m.Rtl, sertl)
	}
	if m.Cs != nil {
		secs := xml.StartElement{Name: xml.Name{Local: "w:cs"}}
		e.EncodeElement(m.Cs, secs)
	}
	if m.Em != nil {
		seem := xml.StartElement{Name: xml.Name{Local: "w:em"}}
		e.EncodeElement(m.Em, seem)
	}
	if m.Lang != nil {
		selang := xml.StartElement{Name: xml.Name{Local: "w:lang"}}
		e.EncodeElement(m.Lang, selang)
	}
	if m.EastAsianLayout != nil {
		seeastAsianLayout := xml.StartElement{Name: xml.Name{Local: "w:eastAsianLayout"}}
		e.EncodeElement(m.EastAsianLayout, seeastAsianLayout)
	}
	if m.SpecVanish != nil {
		sespecVanish := xml.StartElement{Name: xml.Name{Local: "w:specVanish"}}
		e.EncodeElement(m.SpecVanish, sespecVanish)
	}
	if m.OMath != nil {
		seoMath := xml.StartElement{Name: xml.Name{Local: "w:oMath"}}
		e.EncodeElement(m.OMath, seoMath)
	}
	if m.RPrChange != nil {
		serPrChange := xml.StartElement{Name: xml.Name{Local: "w:rPrChange"}}
		e.EncodeElement(m.RPrChange, serPrChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ParaRPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ParaRPr:
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
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rStyle"}:
				m.RStyle = NewCT_String()
				if err := d.DecodeElement(m.RStyle, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rFonts"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rFonts"}:
				m.RFonts = NewCT_Fonts()
				if err := d.DecodeElement(m.RFonts, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "b"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "b"}:
				m.B = NewCT_OnOff()
				if err := d.DecodeElement(m.B, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bCs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bCs"}:
				m.BCs = NewCT_OnOff()
				if err := d.DecodeElement(m.BCs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "i"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "i"}:
				m.I = NewCT_OnOff()
				if err := d.DecodeElement(m.I, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "iCs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "iCs"}:
				m.ICs = NewCT_OnOff()
				if err := d.DecodeElement(m.ICs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "caps"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "caps"}:
				m.Caps = NewCT_OnOff()
				if err := d.DecodeElement(m.Caps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "smallCaps"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "smallCaps"}:
				m.SmallCaps = NewCT_OnOff()
				if err := d.DecodeElement(m.SmallCaps, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "strike"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "strike"}:
				m.Strike = NewCT_OnOff()
				if err := d.DecodeElement(m.Strike, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "dstrike"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "dstrike"}:
				m.Dstrike = NewCT_OnOff()
				if err := d.DecodeElement(m.Dstrike, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "outline"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "outline"}:
				m.Outline = NewCT_OnOff()
				if err := d.DecodeElement(m.Outline, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "shadow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "shadow"}:
				m.Shadow = NewCT_OnOff()
				if err := d.DecodeElement(m.Shadow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "emboss"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "emboss"}:
				m.Emboss = NewCT_OnOff()
				if err := d.DecodeElement(m.Emboss, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "imprint"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "imprint"}:
				m.Imprint = NewCT_OnOff()
				if err := d.DecodeElement(m.Imprint, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "noProof"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "noProof"}:
				m.NoProof = NewCT_OnOff()
				if err := d.DecodeElement(m.NoProof, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "snapToGrid"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "snapToGrid"}:
				m.SnapToGrid = NewCT_OnOff()
				if err := d.DecodeElement(m.SnapToGrid, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "vanish"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "vanish"}:
				m.Vanish = NewCT_OnOff()
				if err := d.DecodeElement(m.Vanish, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "webHidden"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "webHidden"}:
				m.WebHidden = NewCT_OnOff()
				if err := d.DecodeElement(m.WebHidden, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "color"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "color"}:
				m.Color = NewCT_Color()
				if err := d.DecodeElement(m.Color, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "spacing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "spacing"}:
				m.Spacing = NewCT_SignedTwipsMeasure()
				if err := d.DecodeElement(m.Spacing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "w"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "w"}:
				m.W = NewCT_TextScale()
				if err := d.DecodeElement(m.W, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "kern"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "kern"}:
				m.Kern = NewCT_HpsMeasure()
				if err := d.DecodeElement(m.Kern, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "position"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "position"}:
				m.Position = NewCT_SignedHpsMeasure()
				if err := d.DecodeElement(m.Position, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "sz"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "sz"}:
				m.Sz = NewCT_HpsMeasure()
				if err := d.DecodeElement(m.Sz, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "szCs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "szCs"}:
				m.SzCs = NewCT_HpsMeasure()
				if err := d.DecodeElement(m.SzCs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "highlight"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "highlight"}:
				m.Highlight = NewCT_Highlight()
				if err := d.DecodeElement(m.Highlight, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "u"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "u"}:
				m.U = NewCT_Underline()
				if err := d.DecodeElement(m.U, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "effect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "effect"}:
				m.Effect = NewCT_TextEffect()
				if err := d.DecodeElement(m.Effect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "bdr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "bdr"}:
				m.Bdr = NewCT_Border()
				if err := d.DecodeElement(m.Bdr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "shd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "shd"}:
				m.Shd = NewCT_Shd()
				if err := d.DecodeElement(m.Shd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fitText"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fitText"}:
				m.FitText = NewCT_FitText()
				if err := d.DecodeElement(m.FitText, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "vertAlign"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "vertAlign"}:
				m.VertAlign = NewCT_VerticalAlignRun()
				if err := d.DecodeElement(m.VertAlign, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rtl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rtl"}:
				m.Rtl = NewCT_OnOff()
				if err := d.DecodeElement(m.Rtl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "cs"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "cs"}:
				m.Cs = NewCT_OnOff()
				if err := d.DecodeElement(m.Cs, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "em"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "em"}:
				m.Em = NewCT_Em()
				if err := d.DecodeElement(m.Em, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "lang"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "lang"}:
				m.Lang = NewCT_Language()
				if err := d.DecodeElement(m.Lang, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "eastAsianLayout"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "eastAsianLayout"}:
				m.EastAsianLayout = NewCT_EastAsianLayout()
				if err := d.DecodeElement(m.EastAsianLayout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "specVanish"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "specVanish"}:
				m.SpecVanish = NewCT_OnOff()
				if err := d.DecodeElement(m.SpecVanish, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "oMath"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "oMath"}:
				m.OMath = NewCT_OnOff()
				if err := d.DecodeElement(m.OMath, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "rPrChange"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "rPrChange"}:
				m.RPrChange = NewCT_ParaRPrChange()
				if err := d.DecodeElement(m.RPrChange, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ParaRPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ParaRPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ParaRPr and its children
func (m *CT_ParaRPr) Validate() error {
	return m.ValidateWithPath("CT_ParaRPr")
}

// ValidateWithPath validates the CT_ParaRPr and its children, prefixing error messages with path
func (m *CT_ParaRPr) ValidateWithPath(path string) error {
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
	if m.RStyle != nil {
		if err := m.RStyle.ValidateWithPath(path + "/RStyle"); err != nil {
			return err
		}
	}
	if m.RFonts != nil {
		if err := m.RFonts.ValidateWithPath(path + "/RFonts"); err != nil {
			return err
		}
	}
	if m.B != nil {
		if err := m.B.ValidateWithPath(path + "/B"); err != nil {
			return err
		}
	}
	if m.BCs != nil {
		if err := m.BCs.ValidateWithPath(path + "/BCs"); err != nil {
			return err
		}
	}
	if m.I != nil {
		if err := m.I.ValidateWithPath(path + "/I"); err != nil {
			return err
		}
	}
	if m.ICs != nil {
		if err := m.ICs.ValidateWithPath(path + "/ICs"); err != nil {
			return err
		}
	}
	if m.Caps != nil {
		if err := m.Caps.ValidateWithPath(path + "/Caps"); err != nil {
			return err
		}
	}
	if m.SmallCaps != nil {
		if err := m.SmallCaps.ValidateWithPath(path + "/SmallCaps"); err != nil {
			return err
		}
	}
	if m.Strike != nil {
		if err := m.Strike.ValidateWithPath(path + "/Strike"); err != nil {
			return err
		}
	}
	if m.Dstrike != nil {
		if err := m.Dstrike.ValidateWithPath(path + "/Dstrike"); err != nil {
			return err
		}
	}
	if m.Outline != nil {
		if err := m.Outline.ValidateWithPath(path + "/Outline"); err != nil {
			return err
		}
	}
	if m.Shadow != nil {
		if err := m.Shadow.ValidateWithPath(path + "/Shadow"); err != nil {
			return err
		}
	}
	if m.Emboss != nil {
		if err := m.Emboss.ValidateWithPath(path + "/Emboss"); err != nil {
			return err
		}
	}
	if m.Imprint != nil {
		if err := m.Imprint.ValidateWithPath(path + "/Imprint"); err != nil {
			return err
		}
	}
	if m.NoProof != nil {
		if err := m.NoProof.ValidateWithPath(path + "/NoProof"); err != nil {
			return err
		}
	}
	if m.SnapToGrid != nil {
		if err := m.SnapToGrid.ValidateWithPath(path + "/SnapToGrid"); err != nil {
			return err
		}
	}
	if m.Vanish != nil {
		if err := m.Vanish.ValidateWithPath(path + "/Vanish"); err != nil {
			return err
		}
	}
	if m.WebHidden != nil {
		if err := m.WebHidden.ValidateWithPath(path + "/WebHidden"); err != nil {
			return err
		}
	}
	if m.Color != nil {
		if err := m.Color.ValidateWithPath(path + "/Color"); err != nil {
			return err
		}
	}
	if m.Spacing != nil {
		if err := m.Spacing.ValidateWithPath(path + "/Spacing"); err != nil {
			return err
		}
	}
	if m.W != nil {
		if err := m.W.ValidateWithPath(path + "/W"); err != nil {
			return err
		}
	}
	if m.Kern != nil {
		if err := m.Kern.ValidateWithPath(path + "/Kern"); err != nil {
			return err
		}
	}
	if m.Position != nil {
		if err := m.Position.ValidateWithPath(path + "/Position"); err != nil {
			return err
		}
	}
	if m.Sz != nil {
		if err := m.Sz.ValidateWithPath(path + "/Sz"); err != nil {
			return err
		}
	}
	if m.SzCs != nil {
		if err := m.SzCs.ValidateWithPath(path + "/SzCs"); err != nil {
			return err
		}
	}
	if m.Highlight != nil {
		if err := m.Highlight.ValidateWithPath(path + "/Highlight"); err != nil {
			return err
		}
	}
	if m.U != nil {
		if err := m.U.ValidateWithPath(path + "/U"); err != nil {
			return err
		}
	}
	if m.Effect != nil {
		if err := m.Effect.ValidateWithPath(path + "/Effect"); err != nil {
			return err
		}
	}
	if m.Bdr != nil {
		if err := m.Bdr.ValidateWithPath(path + "/Bdr"); err != nil {
			return err
		}
	}
	if m.Shd != nil {
		if err := m.Shd.ValidateWithPath(path + "/Shd"); err != nil {
			return err
		}
	}
	if m.FitText != nil {
		if err := m.FitText.ValidateWithPath(path + "/FitText"); err != nil {
			return err
		}
	}
	if m.VertAlign != nil {
		if err := m.VertAlign.ValidateWithPath(path + "/VertAlign"); err != nil {
			return err
		}
	}
	if m.Rtl != nil {
		if err := m.Rtl.ValidateWithPath(path + "/Rtl"); err != nil {
			return err
		}
	}
	if m.Cs != nil {
		if err := m.Cs.ValidateWithPath(path + "/Cs"); err != nil {
			return err
		}
	}
	if m.Em != nil {
		if err := m.Em.ValidateWithPath(path + "/Em"); err != nil {
			return err
		}
	}
	if m.Lang != nil {
		if err := m.Lang.ValidateWithPath(path + "/Lang"); err != nil {
			return err
		}
	}
	if m.EastAsianLayout != nil {
		if err := m.EastAsianLayout.ValidateWithPath(path + "/EastAsianLayout"); err != nil {
			return err
		}
	}
	if m.SpecVanish != nil {
		if err := m.SpecVanish.ValidateWithPath(path + "/SpecVanish"); err != nil {
			return err
		}
	}
	if m.OMath != nil {
		if err := m.OMath.ValidateWithPath(path + "/OMath"); err != nil {
			return err
		}
	}
	if m.RPrChange != nil {
		if err := m.RPrChange.ValidateWithPath(path + "/RPrChange"); err != nil {
			return err
		}
	}
	return nil
}
