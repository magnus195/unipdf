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

package pdfaextension ;import (_f "fmt";_d "github.com/trimmer-io/go-xmp/models/dc";_fc "github.com/trimmer-io/go-xmp/models/pdf";_e "github.com/trimmer-io/go-xmp/models/xmp_base";_ab "github.com/trimmer-io/go-xmp/models/xmp_mm";_db "github.com/trimmer-io/go-xmp/xmp";
_bf "reflect";_b "sort";_ff "strings";_af "sync";);

// GetTag implements xmp.Model interface.
func (_ee *Model )GetTag (tag string )(string ,error ){_fcbe ,_cec :=_db .GetNativeField (_ee ,tag );if _cec !=nil {return "",_f .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_cec );};return _fcbe ,nil ;};func init (){_db .Register (Namespace ,_db .XmpMetadata );
_db .Register (SchemaNS );_db .Register (PropertyNS );_db .Register (ValueTypeNS );_db .Register (FieldNS );};

// IsZero checks if the resources list has no entries.
func (_ecab Schemas )IsZero ()bool {return len (_ecab )==0};const (ValueTypeNameBoolean ValueTypeName ="\u0042o\u006f\u006c\u0065\u0061\u006e";ValueTypeNameDate ValueTypeName ="\u0044\u0061\u0074\u0065";ValueTypeNameInteger ValueTypeName ="\u0049n\u0074\u0065\u0067\u0065\u0072";
ValueTypeNameReal ValueTypeName ="\u0052\u0065\u0061\u006c";ValueTypeNameText ValueTypeName ="\u0054\u0065\u0078\u0074";ValueTypeNameAgentName ValueTypeName ="\u0041g\u0065\u006e\u0074\u004e\u0061\u006de";ValueTypeNameProperName ValueTypeName ="\u0050\u0072\u006f\u0070\u0065\u0072\u004e\u0061\u006d\u0065";
ValueTypeNameXPath ValueTypeName ="\u0058\u0050\u0061t\u0068";ValueTypeNameGUID ValueTypeName ="\u0047\u0055\u0049\u0044";ValueTypeNameLocale ValueTypeName ="\u004c\u006f\u0063\u0061\u006c\u0065";ValueTypeNameMIMEType ValueTypeName ="\u004d\u0049\u004d\u0045\u0054\u0079\u0070\u0065";
ValueTypeNameRenditionClass ValueTypeName ="\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073";ValueTypeNameResourceRef ValueTypeName ="R\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0052\u0065\u0066";ValueTypeNameURL ValueTypeName ="\u0055\u0052\u004c";
ValueTypeNameURI ValueTypeName ="\u0055\u0052\u0049";ValueTypeNameVersion ValueTypeName ="\u0056e\u0072\u0073\u0069\u006f\u006e";);var PdfSchema =Schema {NamespaceURI :"\u0068\u0074\u0074\u0070:\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065.\u0063o\u006d\u002f\u0070\u0064\u0066\u002f\u0031.\u0033\u002f",Prefix :"\u0070\u0064\u0066",Schema :"\u0041\u0064o\u0062\u0065\u0020P\u0044\u0046\u0020\u0053\u0063\u0068\u0065\u006d\u0061",Property :Properties {{Category :PropertyCategoryInternal ,Description :"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",Name :"\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0050\u0044F\u0020\u0066\u0069l\u0065\u0020\u0076e\u0072\u0073\u0069\u006f\u006e\u0020\u0028\u0066\u006f\u0072 \u0065\u0078\u0061\u006d\u0070le\u003a\u0020\u0031\u002e\u0030\u002c\u0020\u0031\u002e\u0033\u002c\u0020\u0061\u006e\u0064\u0020\u0073\u006f\u0020\u006f\u006e\u0029\u002e",Name :"\u0050\u0044\u0046\u0056\u0065\u0072\u0073\u0069\u006f\u006e",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u006ea\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0074\u006f\u006fl\u0020\u0074\u0068\u0061\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0064\u0020\u0074\u0068\u0065\u0020\u0050\u0044\u0046\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074.",Name :"\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u0049\u006e\u0064\u0069\u0063\u0061\u0074\u0065\u0073\u0020\u0074\u0068\u0061\u0074\u0020\u0074\u0068\u0069s\u0020\u0069\u0073\u0020\u0061\u0020\u0072i\u0067\u0068\u0074\u0073\u002d\u006d\u0061\u006e\u0061\u0067\u0065d\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e\u002e",Name :"\u004d\u0061\u0072\u006b\u0065\u0064",ValueType :ValueTypeNameBoolean }}};


// Property is a schema that describes single property.
type Property struct{

// Category is the property category.
Category PropertyCategory `xmp:"pdfaProperty:category"`;

// Description of the property.
Description string `xmp:"pdfaProperty:description"`;

// Name is a property name.
Name string `xmp:"pdfaProperty:name"`;

// ValueType is the property value type.
ValueType ValueTypeName `xmp:"pdfaProperty:valueType"`;};

// FillModel fills in the document XMP model.
func FillModel (d *_db .Document ,extModel *Model )error {var _de []*_db .Namespace ;for _ ,_fa :=range d .Namespaces (){_de =append (_de ,_fa );};_b .Slice (_de ,func (_ffd ,_cc int )bool {return _de [_ffd ].Name < _de [_cc ].Name });for _ ,_dd :=range _de {switch _dd {case Namespace ,SchemaNS ,PropertyNS ,ValueTypeNS ,FieldNS ,_e .NsXmp ,_d .NsDc :continue ;
default:_ea ,_cd :=GetSchema (_dd .URI );if !_cd {continue ;};_bb :=d .FindModel (_dd );_fcb :=*_ea ;_fcb .Property =Properties {};for _ ,_faa :=range _ea .Property {_ce :=_cf (_bb ,_dd ,_faa );if _ce {_fcb .Property =append (_fcb .Property ,_faa );};};
if len (_fcb .Property )==0{continue ;};var _cdg bool ;for _fad ,_dcc :=range extModel .Schemas {if _dcc .Schema ==_fcb .Schema {_cdg =true ;extModel .Schemas [_fad ]=_fcb ;break ;};};if !_cdg {extModel .Schemas =append (extModel .Schemas ,_fcb );};};};
return nil ;};

// GetSchema for provided namespace.
func GetSchema (namespaceURI string )(*Schema ,bool ){_bbc .RLock ();defer _bbc .RUnlock ();_fce ,_eca :=_cda [namespaceURI ];return _fce ,_eca ;};

// SyncModel implements xmp.Model interface.
func (_fd *Model )SyncModel (d *_db .Document )error {return nil };

// PropertyCategory is the property category enumerator.
type PropertyCategory int ;

// SeqOfValueTypeName gets a value type name of a sequence of input value type names.
func SeqOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0073\u0065\u0071\u0020"+vt };

// Properties is a list of properties.
type Properties []Property ;

// Schemas is the array of xmp metadata extension resources.
type Schemas []Schema ;

// Model is the pdfa extension metadata model.
type Model struct{Schemas Schemas `xmp:"pdfaExtension:schemas"`;};var XmpIDQualSchema =Schema {NamespaceURI :"\u0068t\u0074\u0070:\u002f\u002f\u006e\u0073.\u0061\u0064\u006fb\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u006d\u0070/I\u0064\u0065\u006et\u0069\u0066i\u0065\u0072\u002f\u0071\u0075\u0061l\u002f\u0031.\u0030\u002f",Prefix :"\u0078\u006d\u0070\u0069\u0064\u0071",Schema :"X\u004dP\u0020\u0078\u006d\u0070\u0069\u0064\u0071\u0020q\u0075\u0061\u006c\u0069fi\u0065\u0072",Property :[]Property {{Category :PropertyCategoryInternal ,Name :"\u0053\u0063\u0068\u0065\u006d\u0065",Description :"\u0041\u0020\u0071\u0075\u0061\u006c\u0069\u0066\u0069\u0065\u0072\u0020\u0070\u0072o\u0076\u0069\u0064i\u006e\u0067\u0020\u0074h\u0065\u0020\u006e\u0020\u0061\u006d\u0065\u0020\u006f\u0066\u0020\u0074\u0068\u0065\u0020\u0066\u006f\u0072\u006d\u0061\u006c \u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0063\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u0068\u0065\u006d\u0065\u0020\u0075\u0073ed\u0020\u0066\u006fr\u0020\u0061\u006e\u0020\u0069\u0074\u0065\u006d \u0069\u006e\u0020\u0074\u0068\u0065\u0020\u0078\u006d\u0070\u003a\u0049\u0064\u0065\u006e\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0061\u0072\u0072\u0061\u0079\u002e",ValueType :ValueTypeNameText }},ValueType :nil };


// FieldValueType is a schema that describes a field in a structured type.
type FieldValueType struct{Name string `xmp:"pdfaField:description"`;Description string `xmp:"pdfaField:name"`;ValueType ValueTypeName `xmp:"pdfaField:valueType"`;};

// MarshalXMP implements xmp.Marshaler interface.
func (_gf Schemas )MarshalXMP (e *_db .Encoder ,node *_db .Node ,m _db .Model )error {return _db .MarshalArray (e ,node ,_gf .Typ (),_gf );};func init (){_cda =map[string ]*Schema {_ab .NsXmpMM .URI :&XmpMediaManagementSchema ,"\u0078\u006d\u0070\u0069\u0064\u0071":&XmpIDQualSchema ,_fc .NsPDF .URI :&PdfSchema ,"\u0073\u0074\u0045v\u0074":&FieldResourceEventSchema ,"\u0073\u0074\u0056e\u0072":&FieldVersionSchema };
};var (_cda =map[string ]*Schema {};_bbc _af .RWMutex ;);

// AltOfValueTypeName gets the ValueTypeName of the alt of given value type names.
func AltOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0061\u006c\u0074\u0020"+vt };

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_age *ValueTypes )UnmarshalXMP (d *_db .Decoder ,node *_db .Node ,m _db .Model )error {return _db .UnmarshalArray (d ,node ,_age .Typ (),_age );};var FieldVersionSchema =Schema {NamespaceURI :"\u0068\u0074\u0074p\u003a\u002f\u002f\u006e\u0073\u002e\u0061\u0064\u006f\u0062\u0065\u002e\u0063\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065/\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0023",Prefix :"\u0073\u0074\u0056e\u0072",Schema :"\u0042a\u0073\u0069\u0063\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u0074y\u0070\u0065\u0020\u0056\u0065\u0072\u0073\u0069\u006f\u006e",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0043\u006fmm\u0065\u006e\u0074s\u0020\u0063\u006f\u006ecer\u006ein\u0067\u0020\u0077\u0068\u0061\u0074\u0020wa\u0073\u0020\u0063\u0068\u0061\u006e\u0067e\u0064",Name :"\u0063\u006f\u006d\u006d\u0065\u006e\u0074\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0048\u0069\u0067\u0068\u0020\u006c\u0065\u0076\u0065\u006c\u002c\u0020\u0066\u006f\u0072\u006d\u0061\u006c\u0020\u0064e\u0073\u0063\u0072\u0069\u0070\u0074\u0069\u006f\u006e\u0020\u006f\u0066\u0020\u0077\u0068\u0061\u0074\u0020\u006f\u0070\u0065r\u0061\u0074\u0069\u006f\u006e\u0020\u0074\u0068\u0065\u0020\u0075\u0073\u0065\u0072 \u0070\u0065\u0072f\u006f\u0072\u006d\u0065\u0064\u002e",Name :"\u0065\u0076\u0065n\u0074",ValueType :ValueTypeResourceEvent },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068e\u0020\u0064\u0061\u0074\u0065\u0020\u006f\u006e\u0020\u0077\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u0077\u0061\u0073\u0020\u0063\u0068\u0065\u0063\u006b\u0065\u0064\u0020\u0069\u006e\u002e",Name :"\u006d\u006f\u0064\u0069\u0066\u0079\u0044\u0061\u0074\u0065",ValueType :ValueTypeNameDate },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068e\u0020\u0070\u0065\u0072s\u006f\u006e \u0077\u0068\u006f\u0020\u006d\u006f\u0064\u0069f\u0069\u0065\u0064\u0020\u0074\u0068\u0069\u0073\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u002e",Name :"\u006d\u006f\u0064\u0069\u0066\u0069\u0065\u0072",ValueType :ValueTypeNameProperName },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 n\u0065\u0077\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u002e",Name :"\u0076e\u0072\u0073\u0069\u006f\u006e",ValueType :ValueTypeNameText }},ValueType :nil };


// CanTag implements xmp.Model interface.
func (_ec *Model )CanTag (tag string )bool {_ ,_afb :=_db .GetNativeField (_ec ,tag );return _afb ==nil ;};

// ValueTypeName is the name of the value type.
type ValueTypeName string ;

// SyncToXMP implements xmp.Model interface.
func (_geg *Model )SyncToXMP (d *_db .Document )error {return nil };

// SetSchema sets the schema into given model.
func (_gd *Model )SetSchema (namespaceURI string ,s Schema ){for _ag :=0;_ag < len (_gd .Schemas );_ag ++{if _gd .Schemas [_ag ].NamespaceURI ==namespaceURI {_gd .Schemas [_ag ]=s ;return ;};};_gd .Schemas =append (_gd .Schemas ,s );};var (Namespace =_db .NewNamespace ("\u0070\u0064\u0066\u0061\u0045\u0078\u0074\u0065\u006e\u0073\u0069\u006f\u006e","\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002f\u0070\u0064\u0066\u0061/\u006e\u0073\u002f\u0065\u0078t\u0065\u006es\u0069\u006f\u006e\u002f",NewModel );
SchemaNS =_db .NewNamespace ("\u0070\u0064\u0066\u0061\u0053\u0063\u0068\u0065\u006d\u0061","h\u0074\u0074\u0070\u003a\u002f\u002fw\u0077\u0077\u002e\u0061\u0069\u0069m\u002e\u006f\u0072\u0067\u002f\u0070\u0064f\u0061\u002f\u006e\u0073\u002f\u0073\u0063\u0068\u0065\u006da\u0023",nil );
PropertyNS =_db .NewNamespace ("\u0070\u0064\u0066a\u0050\u0072\u006f\u0070\u0065\u0072\u0074\u0079","\u0068\u0074t\u0070\u003a\u002f\u002fw\u0077\u0077.\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067/\u0070\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0070\u0072\u006f\u0070e\u0072\u0074\u0079\u0023",nil );
ValueTypeNS =_db .NewNamespace ("\u0070\u0064\u0066\u0061\u0054\u0079\u0070\u0065","\u0068\u0074\u0074\u0070\u003a\u002f/\u0077\u0077\u0077\u002e\u0061\u0069\u0069\u006d\u002e\u006f\u0072\u0067\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002ft\u0079\u0070\u0065\u0023",nil );
FieldNS =_db .NewNamespace ("\u0070d\u0066\u0061\u0046\u0069\u0065\u006cd","\u0068\u0074\u0074p:\u002f\u002f\u0077\u0077\u0077\u002e\u0061\u0069\u0069m\u002eo\u0072g\u002fp\u0064\u0066\u0061\u002f\u006e\u0073\u002f\u0066\u0069\u0065\u006c\u0064\u0023",nil );
);const (PropertyCategoryUndefined PropertyCategory =iota ;PropertyCategoryInternal ;PropertyCategoryExternal ;);

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_ebg *Properties )UnmarshalXMP (d *_db .Decoder ,node *_db .Node ,m _db .Model )error {return _db .UnmarshalArray (d ,node ,_ebg .Typ (),_ebg );};

// Typ gets the array type of properties.
func (_eb Properties )Typ ()_db .ArrayType {return _db .ArrayTypeOrdered };

// BagOfValueTypeName gets the ValueTypeName of the bag of provided value type names.
func BagOfValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0062\u0061\u0067\u0020"+vt };

// Can implements xmp.Model interface.
func (_dba *Model )Can (nsName string )bool {return Namespace .GetName ()==nsName };func _cf (_fcbf _db .Model ,_cg *_db .Namespace ,_gg Property )bool {_ccf :=_bf .ValueOf (_fcbf ).Elem ();if _ccf .Kind ()==_bf .Ptr {_ccf =_ccf .Elem ();};_dcca :=_ccf .Type ();
for _ffg :=0;_ffg < _dcca .NumField ();_ffg ++{_ae :=_dcca .Field (_ffg );_ded :=_ae .Tag .Get ("\u0078\u006d\u0070");if _ded ==""{continue ;};if !_ff .HasPrefix (_ded ,_cg .Name ){continue ;};_afg :=_ff .IndexRune (_ded ,':');if _afg ==-1{continue ;};
_aa :=_ded [_afg +1:];if _aa ==_gg .Name {_ge :=_ccf .Field (_ffg );return !_ge .IsZero ();};};return false ;};

// MarshalXMP implements xmp.Marshaler interface.
func (_aea ValueTypes )MarshalXMP (e *_db .Encoder ,node *_db .Node ,m _db .Model )error {return _db .MarshalArray (e ,node ,_aea .Typ (),_aea );};

// Typ gets array type of the field value types.
func (_aeg ValueTypes )Typ ()_db .ArrayType {return _db .ArrayTypeOrdered };

// SetTag implements xmp.Model interface.
func (_eed *Model )SetTag (tag ,value string )error {if _ece :=_db .SetNativeField (_eed ,tag ,value );_ece !=nil {return _f .Errorf ("\u0025\u0073\u003a\u0020\u0025\u0076",Namespace .GetName (),_ece );};return nil ;};

// UnmarshalXMP implements xmp.Unmarshaler interface.
func (_dg *Schemas )UnmarshalXMP (d *_db .Decoder ,node *_db .Node ,m _db .Model )error {return _db .UnmarshalArray (d ,node ,_dg .Typ (),_dg );};const (ValueTypeResourceEvent ValueTypeName ="\u0052\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065\u006e\u0074";
);

// ValueTypes is the slice of field value types.
type ValueTypes []FieldValueType ;var FieldResourceEventSchema =Schema {NamespaceURI :"\u0068\u0074\u0074\u0070\u003a\u002f\u002f\u006es\u002e\u0061\u0064ob\u0065\u002e\u0063\u006f\u006d\u002fx\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u0073\u0054\u0079\u0070\u0065\u002f\u0052\u0065s\u006f\u0075\u0072\u0063\u0065\u0045\u0076\u0065n\u0074\u0023",Prefix :"\u0073\u0074\u0045v\u0074",Schema :"\u0044e\u0066\u0069n\u0069\u0074\u0069\u006fn\u0020\u006f\u0066 \u0062\u0061\u0073\u0069\u0063\u0020\u0076\u0061\u006cue\u0020\u0074\u0079p\u0065\u0020R\u0065\u0073\u006f\u0075\u0072\u0063e\u0045\u0076e\u006e\u0074",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0054he\u0020a\u0063t\u0069\u006f\u006e\u0020\u0074\u0068a\u0074\u0020\u006f\u0063c\u0075\u0072\u0072\u0065\u0064\u002e\u0020\u0044\u0065\u0066\u0069\u006e\u0065d \u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0061\u0072\u0065\u003a\u0020\u0063\u006f\u006e\u0076\u0065\u0072\u0074\u0065d\u002c \u0063\u006f\u0070\u0069\u0065\u0064\u002c\u0020\u0063\u0072\u0065\u0061\u0074e\u0064\u002c \u0063\u0072\u006fp\u0070\u0065\u0064\u002c\u0020\u0065\u0064\u0069\u0074ed\u002c\u0020\u0066i\u006c\u002d\u0074\u0065r\u0065\u0064\u002c\u0020\u0066\u006fr\u006d\u0061t\u0074\u0065\u0064\u002c\u0020\u0076\u0065\u0072s\u0069\u006f\u006e\u005f\u0075\u0070\u0064\u0061\u0074\u0065\u0064\u002c\u0020\u0070\u0072\u0069\u006e\u0074\u0065\u0064\u002c\u0020\u0070ubli\u0073\u0068\u0065\u0064\u002c\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0064\u002c\u0020\u0070\u0072\u006f\u0064\u0075\u0063\u0065\u0064\u002c\u0020\u0072\u0065\u0073i\u007ae\u0064.\u004e\u0065\u0077\u0020\u0076\u0061\u006c\u0075\u0065\u0073\u0020\u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0062\u0065 \u0076\u0065r\u0062\u0073 \u0069n\u0020\u0074\u0068\u0065\u0020\u0070a\u0073\u0074\u0020\u0074\u0065\u006e\u0073\u0065\u002e",Name :"\u0061\u0063\u0074\u0069\u006f\u006e",ValueType :"O\u0070\u0065\u006e\u0020\u0043\u0068\u006f\u0069\u0063\u0065"},{Category :PropertyCategoryInternal ,Description :"T\u0068\u0065\u0020\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0020\u0049\u0044\u0020\u006f\u0066\u0020t\u0068\u0065\u0020\u006d\u006f\u0064\u0069\u0066\u0069\u0065d \u0072\u0065\u0073o\u0075r\u0063\u0065",Name :"\u0069\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0041\u0064di\u0074\u0069\u006fn\u0061\u006c\u0020\u0064esc\u0072ip\u0074\u0069\u006f\u006e\u0020\u006f\u0066 t\u0068\u0065\u0020\u0061\u0063\u0074\u0069o\u006e",Name :"\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054h\u0065\u0020s\u006f\u0066\u0074\u0077a\u0072\u0065\u0020a\u0067\u0065\u006e\u0074\u0020\u0074\u0068\u0061\u0074 p\u0065\u0072\u0066o\u0072\u006de\u0064\u0020\u0074\u0068\u0065\u0020a\u0063\u0074i\u006f\u006e",Name :"\u0073\u006f\u0066\u0074\u0077\u0061\u0072\u0065\u0041\u0067\u0065\u006e\u0074",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u004f\u0070t\u0069\u006f\u006e\u0061\u006c\u0020\u0074\u0069\u006d\u0065\u0073\u0074\u0061\u006d\u0070\u0020\u006f\u0066\u0020\u0077\u0068\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0061\u0063\u0074\u0069\u006f\u006e\u0020\u006f\u0063\u0063\u0075\u0072\u0072\u0065\u0064",Name :"\u0077\u0068\u0065\u006e",ValueType :ValueTypeNameDate }},ValueType :nil };


// NewModel creates a new pdfAExtension model.
func NewModel (name string )_db .Model {return &Model {}};

// Schema is the pdfa extension schema.
type Schema struct{

// NamespaceURI is schema namespace URI.
NamespaceURI string `xmp:"pdfaSchema:namespaceURI"`;

// Prefix is preferred schema namespace prefix.
Prefix string `xmp:"pdfaSchema:prefix"`;

// Schema is optional description.
Schema string `xmp:"pdfaSchema:schema"`;

// Property is description of schema properties.
Property Properties `xmp:"pdfaSchema:property"`;

// ValueType is description of schema-specific value types.
ValueType ValueTypes `xmp:"pdfaSchema:valueType"`;};

// ValueType is the pdfa extension value type schema.
type ValueType struct{Description string `xmp:"pdfaType:description"`;Field []FieldValueType `xmp:"pdfaType:field"`;NamespaceURI string `xmp:"pdfaType:namespaceURI"`;Prefix string `xmp:"pdfaType:prefix"`;Type string `xmp:"pdfaType:type"`;};func (_bfe Schemas )Typ ()_db .ArrayType {return _db .ArrayTypeUnordered };


// SyncFromXMP implements xmp.Model interface.
func (_ccfg *Model )SyncFromXMP (d *_db .Document )error {return nil };

// Namespaces implements xmp.Model interface.
func (_dee *Model )Namespaces ()_db .NamespaceList {return _db .NamespaceList {Namespace }};

// MakeModel creates or gets a model from document.
func MakeModel (d *_db .Document )(*Model ,error ){_fb ,_ffa :=d .MakeModel (Namespace );if _ffa !=nil {return nil ,_ffa ;};return _fb .(*Model ),nil ;};

// ClosedChoiceValueTypeName gets the closed choice of provided value type name.
func ClosedChoiceValueTypeName (vt ValueTypeName )ValueTypeName {return "\u0043\u006c\u006f\u0073\u0065\u0064\u0020\u0043\u0068\u006f\u0069\u0063e\u0020\u006f\u0066\u0020"+vt ;};

// RegisterSchema registers schema extension definition.
func RegisterSchema (ns *_db .Namespace ,schema *Schema ){_bbc .Lock ();defer _bbc .Unlock ();_cda [ns .URI ]=schema ;};

// MarshalXMP implements xmp.Marshaler interface.
func (_da Properties )MarshalXMP (e *_db .Encoder ,node *_db .Node ,m _db .Model )error {return _db .MarshalArray (e ,node ,_da .Typ (),_da );};

// UnmarshalText implements encoding.TextUnmarshaler interface.
func (_ega *PropertyCategory )UnmarshalText (in []byte )error {switch string (in ){case "\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c":*_ega =PropertyCategoryInternal ;case "\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c":*_ega =PropertyCategoryExternal ;
default:*_ega =PropertyCategoryUndefined ;};return nil ;};

// MarshalText implements encoding.TextMarshaler interface.
func (_egb PropertyCategory )MarshalText ()([]byte ,error ){switch _egb {case PropertyCategoryInternal :return []byte ("\u0069\u006e\u0074\u0065\u0072\u006e\u0061\u006c"),nil ;case PropertyCategoryExternal :return []byte ("\u0065\u0078\u0074\u0065\u0072\u006e\u0061\u006c"),nil ;
case PropertyCategoryUndefined :return []byte (""),nil ;default:return nil ,_f .Errorf ("\u0075\u006ed\u0065\u0066\u0069\u006ee\u0064\u0020p\u0072\u006f\u0070\u0065\u0072\u0074\u0079\u0020c\u0061\u0074\u0065\u0067\u006f\u0072\u0079\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_egb );
};};var XmpMediaManagementSchema =Schema {NamespaceURI :"\u0068\u0074\u0074p\u003a\u002f\u002f\u006es\u002e\u0061\u0064\u006f\u0062\u0065\u002ec\u006f\u006d\u002f\u0078\u0061\u0070\u002f\u0031\u002e\u0030\u002f\u006d\u006d\u002f",Prefix :"\u0078\u006d\u0070M\u004d",Schema :"X\u004dP\u0020\u004d\u0065\u0064\u0069\u0061\u0020\u004da\u006e\u0061\u0067\u0065me\u006e\u0074",Property :[]Property {{Category :PropertyCategoryInternal ,Description :"\u0041\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0020\u0074\u006f\u0020\u0074\u0068\u0065\u0020\u006fr\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0066\u0072\u006f\u006d\u0020w\u0068\u0069\u0063\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u006f\u006e\u0065\u0020i\u0073\u0020\u0064e\u0072\u0069\u0076\u0065\u0064\u002e",Name :"D\u0065\u0072\u0069\u0076\u0065\u0064\u0046\u0072\u006f\u006d",ValueType :ValueTypeNameResourceRef },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0072\u0065\u0073\u006fu\u0072\u0063\u0065",Name :"\u0044\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"UU\u0049\u0044 \u0062\u0061\u0073\u0065\u0064\u0020\u0069\u0064\u0065n\u0074\u0069\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0073\u0070\u0065\u0063\u0069\u0066\u0069\u0063\u0020\u0069\u006e\u0063\u0061\u0072\u006e\u0061\u0074i\u006fn\u0020\u006f\u0066\u0020\u0061\u0020\u0064\u006fc\u0075m\u0065\u006et",Name :"\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0063\u006fm\u006d\u006f\u006e\u0020\u0069d\u0065\u006e\u0074\u0069\u0066\u0069e\u0072\u0020\u0066\u006f\u0072 \u0061\u006c\u006c\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0073 \u0061\u006e\u0064\u0020\u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0073\u0020o\u0066\u0020\u0061\u0020\u0064\u006f\u0063\u0075m\u0065\u006e\u0074",Name :"\u004fr\u0069g\u0069\u006e\u0061\u006c\u0044o\u0063\u0075m\u0065\u006e\u0074\u0049\u0044",ValueType :ValueTypeNameURI },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0020\u006e\u0061\u006d\u0065\u0020\u0066\u006f\u0072\u0020\u0074\u0068\u0069\u0073 \u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065",Name :"\u0052\u0065\u006e\u0064\u0069\u0074\u0069\u006f\u006eC\u006c\u0061\u0073\u0073",ValueType :ValueTypeNameRenditionClass },{Category :PropertyCategoryInternal ,Description :"\u0043\u0061n\u0020\u0062\u0065\u0020\u0075\u0073\u0065d\u0020t\u006f\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0020\u0061\u0064\u0064\u0069\u0074\u0069\u006f\u006e\u0061\u006c\u0020\u0072\u0065\u006e\u0064\u0069t\u0069\u006f\u006e\u0020\u0070\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0074\u0068a\u0074 \u0061r\u0065\u0020\u0074o\u006f\u0020\u0063\u006f\u006d\u0070\u006c\u0065\u0078\u0020\u006f\u0072\u0020\u0076\u0065\u0072\u0062o\u0073\u0065\u0020\u0074\u006f\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0020\u0069\u006e",Name :"\u0052e\u006ed\u0069\u0074\u0069\u006f\u006e\u0050\u0061\u0072\u0061\u006d\u0073",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0064\u006f\u0063\u0075\u006d\u0065\u006e\u0074\u0020\u0076\u0065\u0072\u0073\u0069o\u006e\u0020\u0069\u0064\u0065\u006e\u0074i\u0066\u0069\u0065\u0072\u0020\u0066\u006f\u0072\u0020\u0074\u0068i\u0073\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u002e",Name :"\u0056e\u0072\u0073\u0069\u006f\u006e\u0049D",ValueType :ValueTypeNameText },{Category :PropertyCategoryExternal ,Description :"\u0054\u0068\u0065\u0020\u0076\u0065r\u0073\u0069\u006f\u006e\u0020\u0068\u0069\u0073\u0074\u006f\u0072\u0079\u0020\u0061\u0073\u0073\u006f\u0063\u0069\u0061t\u0065\u0064\u0020\u0077\u0069\u0074\u0068\u0020\u0074\u0068\u0069\u0073\u0020\u0072e\u0073o\u0075\u0072\u0063\u0065",Name :"\u0056\u0065\u0072\u0073\u0069\u006f\u006e\u0073",ValueType :SeqOfValueTypeName (ValueTypeNameVersion )},{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065\u0020\u0072\u0065f\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075r\u0063\u0065\u0027\u0073\u0020\u006d\u0061n\u0061\u0067\u0065\u0072",Name :"\u004da\u006e\u0061\u0067\u0065\u0072",ValueType :ValueTypeNameAgentName },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e",Name :"\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074",ValueType :ValueTypeNameText },{Category :PropertyCategoryInternal ,Description :"\u0054\u0068\u0065 \u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065\u0064\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0027\u0073\u0020\u006d\u0061\u006e\u0061\u0067\u0065\u0072 \u0076\u0061\u0072\u0069\u0061\u006e\u0074\u002e",Name :"\u004d\u0061\u006e\u0061\u0067\u0065\u0072\u0056\u0061r\u0069\u0061\u006e\u0074",ValueType :ValueTypeNameText }}};
