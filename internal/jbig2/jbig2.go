//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package jbig2 ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_de "github.com/unidoc/unipdf/v3/internal/jbig2/decoder";_c "github.com/unidoc/unipdf/v3/internal/jbig2/document";_e "github.com/unidoc/unipdf/v3/internal/jbig2/document/segments";
_g "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_b "sort";);func DecodeBytes (encoded []byte ,parameters _de .Parameters ,globals ...Globals )([]byte ,error ){var _fc Globals ;if len (globals )> 0{_fc =globals [0];};_fcb ,_fa :=_de .Decode (encoded ,parameters ,_fc .ToDocumentGlobals ());
if _fa !=nil {return nil ,_fa ;};return _fcb .DecodeNextPage ();};func DecodeGlobals (encoded []byte )(Globals ,error ){const _bd ="\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073";_eg :=_f .NewReader (encoded );_a ,_ca :=_c .DecodeDocument (_eg ,nil );
if _ca !=nil {return nil ,_g .Wrap (_ca ,_bd ,"");};if _a .GlobalSegments ==nil ||(_a .GlobalSegments .Segments ==nil ){return nil ,_g .Error (_bd ,"\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064");
};_df :=Globals {};for _ ,_gb :=range _a .GlobalSegments .Segments {_df [int (_gb .SegmentNumber )]=_gb ;};return _df ,nil ;};type Globals map[int ]*_e .Header ;func (_cb Globals )ToDocumentGlobals ()*_c .Globals {if _cb ==nil {return nil ;};_egd :=[]*_e .Header {};
for _ ,_caf :=range _cb {_egd =append (_egd ,_caf );};_b .Slice (_egd ,func (_cab ,_db int )bool {return _egd [_cab ].SegmentNumber < _egd [_db ].SegmentNumber });return &_c .Globals {Segments :_egd };};