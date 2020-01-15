// Copyright 2017 FoxyUtils ehf. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased via https://unidoc.io website.

package docPropsVTypes_test

import (
	"encoding/xml"
	"testing"

	"github.com/xmdas-link/unioffice/schema/soo/ofc/docPropsVTypes"
)

func TestVectorConstructor(t *testing.T) {
	v := docPropsVTypes.NewVector()
	if v == nil {
		t.Errorf("docPropsVTypes.NewVector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed docPropsVTypes.Vector should validate: %s", err)
	}
}

func TestVectorMarshalUnmarshal(t *testing.T) {
	v := docPropsVTypes.NewVector()
	buf, _ := xml.Marshal(v)
	v2 := docPropsVTypes.NewVector()
	xml.Unmarshal(buf, v2)
}
