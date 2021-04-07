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

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_a "encoding/json";_bf "github.com/unidoc/unipdf/v3/core";_cb "github.com/unidoc/unipdf/v3/model";_c "io";_b "os";);

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _c .Reader )(*FieldData ,error ){var _ba FieldData ;_fb :=_a .NewDecoder (r ).Decode (&_ba ._f );if _fb !=nil {return nil ,_fb ;};return &_ba ,nil ;};type fieldValue struct{Name string `json:"name"`;Value string `json:"value"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};

// JSON returns the field data as a string in JSON format.
func (_cge FieldData )JSON ()(string ,error ){_ce ,_fe :=_a .MarshalIndent (_cge ._f ,"","\u0020\u0020\u0020\u0020");return string (_ce ),_fe ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _c .ReadSeeker )(*FieldData ,error ){_cg ,_ag :=_cb .NewPdfReader (rs );if _ag !=nil {return nil ,_ag ;};if _cg .AcroForm ==nil {return nil ,nil ;};var _eg []fieldValue ;_bad :=_cg .AcroForm .AllFields ();for _ ,_bg :=range _bad {var _ad []string ;
_fbb :=make (map[string ]struct{});_ca ,_bfd :=_bg .FullName ();if _bfd !=nil {return nil ,_bfd ;};if _eba ,_fbc :=_bg .V .(*_bf .PdfObjectString );_fbc {_eg =append (_eg ,fieldValue {Name :_ca ,Value :_eba .Decoded ()});continue ;};var _bfe string ;for _ ,_agd :=range _bg .Annotations {_g ,_ef :=_bf .GetName (_agd .AS );
if _ef {_bfe =_g .String ();};_egg ,_ee :=_bf .GetDict (_agd .AP );if !_ee {continue ;};_edd ,_ :=_bf .GetDict (_egg .Get ("\u004e"));for _ ,_db :=range _edd .Keys (){_fd :=_db .String ();if _ ,_ae :=_fbb [_fd ];!_ae {_ad =append (_ad ,_fd );_fbb [_fd ]=struct{}{};
};};_da ,_ :=_bf .GetDict (_egg .Get ("\u0044"));for _ ,_cad :=range _da .Keys (){_gd :=_cad .String ();if _ ,_caf :=_fbb [_gd ];!_caf {_ad =append (_ad ,_gd );_fbb [_gd ]=struct{}{};};};};_fba :=fieldValue {Name :_ca ,Value :_bfe ,Options :_ad };_eg =append (_eg ,_fba );
};_agc :=FieldData {_f :_eg };return &_agc ,nil ;};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_ed ,_d :=_b .Open (filePath );if _d !=nil {return nil ,_d ;};defer _ed .Close ();return LoadFromJSON (_ed );};

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_agg ,_bcd :=_b .Open (filePath );if _bcd !=nil {return nil ,_bcd ;};defer _agg .Close ();return LoadFromPDF (_agg );};

// FieldValues implements model.FieldValueProvider interface.
func (_dbb *FieldData )FieldValues ()(map[string ]_bf .PdfObject ,error ){_edc :=make (map[string ]_bf .PdfObject );for _ ,_aa :=range _dbb ._f {if len (_aa .Value )> 0{_edc [_aa .Name ]=_bf .MakeString (_aa .Value );};};return _edc ,nil ;};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_f []fieldValue };