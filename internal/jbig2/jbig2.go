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

package jbig2 ;import (_f "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/jbig2/decoder";_fd "github.com/unidoc/unipdf/v3/internal/jbig2/document";_e "github.com/unidoc/unipdf/v3/internal/jbig2/document/segments";
_b "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_fe "sort";);func DecodeGlobals (encoded []byte )(Globals ,error ){const _c ="\u0044\u0065\u0063\u006f\u0064\u0065\u0047\u006c\u006f\u0062\u0061\u006c\u0073";_ag :=_f .NewReader (encoded );_ce ,_af :=_fd .DecodeDocument (_ag ,nil );
if _af !=nil {return nil ,_b .Wrap (_af ,_c ,"");};if _ce .GlobalSegments ==nil ||(_ce .GlobalSegments .Segments ==nil ){return nil ,_b .Error (_c ,"\u006eo\u0020\u0067\u006c\u006f\u0062\u0061\u006c\u0020\u0073\u0065\u0067m\u0065\u006e\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064");
};_gc :=Globals {};for _ ,_ceg :=range _ce .GlobalSegments .Segments {_gc [int (_ceg .SegmentNumber )]=_ceg ;};return _gc ,nil ;};func DecodeBytes (encoded []byte ,parameters _a .Parameters ,globals ...Globals )([]byte ,error ){var _eb Globals ;if len (globals )> 0{_eb =globals [0];
};_d ,_be :=_a .Decode (encoded ,parameters ,_eb .ToDocumentGlobals ());if _be !=nil {return nil ,_be ;};return _d .DecodeNextPage ();};type Globals map[int ]*_e .Header ;func (_ga Globals )ToDocumentGlobals ()*_fd .Globals {if _ga ==nil {return nil ;};
_ec :=[]*_e .Header {};for _ ,_ff :=range _ga {_ec =append (_ec ,_ff );};_fe .Slice (_ec ,func (_ef ,_bb int )bool {return _ec [_ef ].SegmentNumber < _ec [_bb ].SegmentNumber });return &_fd .Globals {Segments :_ec };};