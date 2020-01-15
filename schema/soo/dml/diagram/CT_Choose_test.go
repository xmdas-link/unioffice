// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/xmdas-link/unioffice/schema/soo/dml/diagram"
)

func TestCT_ChooseConstructor(t *testing.T) {
	v := diagram.NewCT_Choose()
	if v == nil {
		t.Errorf("diagram.NewCT_Choose must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Choose should validate: %s", err)
	}
}

func TestCT_ChooseMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Choose()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Choose()
	xml.Unmarshal(buf, v2)
}
