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
package fjson ;import (_g "encoding/json";_b "github.com/unidoc/unipdf/v3/common";_e "github.com/unidoc/unipdf/v3/core";_bf "github.com/unidoc/unipdf/v3/model";_ae "io";_f "os";);

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_fee ,_fad :=_f .Open (filePath );if _fad !=nil {return nil ,_fad ;};defer _fee .Close ();return LoadFromPDF (_fee );};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _ae .Reader )(*FieldData ,error ){var _bd FieldData ;_ce :=_g .NewDecoder (r ).Decode (&_bd ._c );if _ce !=nil {return nil ,_ce ;};return &_bd ,nil ;};

// SetImage assign model.Image to a specific field identified by fieldName.
func (_bc *FieldData )SetImage (fieldName string ,img *_bf .Image ,opt []string )error {_ecg :=fieldValue {Name :fieldName ,ImageValue :img ,Options :opt };_bc ._c =append (_bc ._c ,_ecg );return nil ;};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_c []fieldValue };type fieldValue struct{Name string `json:"name"`;Value string `json:"value"`;ImageValue *_bf .Image `json:"-"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};

// JSON returns the field data as a string in JSON format.
func (_feb FieldData )JSON ()(string ,error ){_fcd ,_acb :=_g .MarshalIndent (_feb ._c ,"","\u0020\u0020\u0020\u0020");return string (_fcd ),_acb ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _ae .ReadSeeker )(*FieldData ,error ){_fa ,_ca :=_bf .NewPdfReader (rs );if _ca !=nil {return nil ,_ca ;};if _fa .AcroForm ==nil {return nil ,nil ;};var _ab []fieldValue ;_cec :=_fa .AcroForm .AllFields ();for _ ,_fg :=range _cec {var _be []string ;
_beb :=make (map[string ]struct{});_ge ,_d :=_fg .FullName ();if _d !=nil {return nil ,_d ;};if _ff ,_ee :=_fg .V .(*_e .PdfObjectString );_ee {_ab =append (_ab ,fieldValue {Name :_ge ,Value :_ff .Decoded ()});continue ;};var _eg string ;for _ ,_cf :=range _fg .Annotations {_cb ,_ba :=_e .GetName (_cf .AS );
if _ba {_eg =_cb .String ();};_af ,_fc :=_e .GetDict (_cf .AP );if !_fc {continue ;};_cbf ,_ :=_e .GetDict (_af .Get ("\u004e"));for _ ,_fd :=range _cbf .Keys (){_gg :=_fd .String ();if _ ,_fed :=_beb [_gg ];!_fed {_be =append (_be ,_gg );_beb [_gg ]=struct{}{};
};};_bdg ,_ :=_e .GetDict (_af .Get ("\u0044"));for _ ,_bg :=range _bdg .Keys (){_gd :=_bg .String ();if _ ,_ed :=_beb [_gd ];!_ed {_be =append (_be ,_gd );_beb [_gd ]=struct{}{};};};};_ege :=fieldValue {Name :_ge ,Value :_eg ,Options :_be };_ab =append (_ab ,_ege );
};_eed :=FieldData {_c :_ab };return &_eed ,nil ;};

// FieldImageValues implements model.FieldImageProvider interface.
func (_bfg *FieldData )FieldImageValues ()(map[string ]*_bf .Image ,error ){_efa :=make (map[string ]*_bf .Image );for _ ,_fbe :=range _bfg ._c {if _fbe .ImageValue !=nil {_efa [_fbe .Name ]=_fbe .ImageValue ;};};return _efa ,nil ;};

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_gef *FieldData )SetImageFromFile (fieldName string ,imagePath string ,opt []string )error {_fcb ,_ec :=_f .Open (imagePath );if _ec !=nil {return _ec ;};defer _fcb .Close ();_ag ,_ec :=_bf .ImageHandling .Read (_fcb );if _ec !=nil {_b .Log .Error ("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_ec );
return _ec ;};return _gef .SetImage (fieldName ,_ag ,opt );};

// FieldValues implements model.FieldValueProvider interface.
func (_caa *FieldData )FieldValues ()(map[string ]_e .PdfObject ,error ){_fdc :=make (map[string ]_e .PdfObject );for _ ,_fea :=range _caa ._c {if len (_fea .Value )> 0{_fdc [_fea .Name ]=_e .MakeString (_fea .Value );};};return _fdc ,nil ;};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_cd ,_cg :=_f .Open (filePath );if _cg !=nil {return nil ,_cg ;};defer _cd .Close ();return LoadFromJSON (_cd );};