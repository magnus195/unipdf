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

package docutil ;import (_g "errors";_d "fmt";
	_ad "github.com/magnus195/unipdf/v3/common";
	_gc "github.com/magnus195/unipdf/v3/core";);type Catalog struct{Object *_gc .PdfObjectDictionary ;_f *Document ;};func (_dg *Catalog )GetMarkInfo ()(*_gc .PdfObjectDictionary ,bool ){_ab ,_fbf :=_gc .GetDict (_dg .Object .Get ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"));
return _ab ,_fbf ;};func (_dag *Document )AddStream (stream *_gc .PdfObjectStream ){for _ ,_fcd :=range _dag .Objects {if _fcd ==stream {return ;};};_dag .Objects =append (_dag .Objects ,stream );};type Image struct{Name string ;Width int ;Height int ;
Colorspace _gc .PdfObjectName ;ColorComponents int ;BitsPerComponent int ;SMask *ImageSMask ;Stream *_gc .PdfObjectStream ;};type OutputIntents struct{_deg *_gc .PdfObjectArray ;_fe *Document ;_ba *_gc .PdfIndirectObject ;};func _gee (_aec _gc .PdfObject )(_gc .PdfObjectName ,error ){var _fcc *_gc .PdfObjectName ;
var _cfag *_gc .PdfObjectArray ;if _df ,_gf :=_aec .(*_gc .PdfIndirectObject );_gf {if _bd ,_dad :=_df .PdfObject .(*_gc .PdfObjectArray );_dad {_cfag =_bd ;}else if _cd ,_aee :=_df .PdfObject .(*_gc .PdfObjectName );_aee {_fcc =_cd ;};}else if _gb ,_gcd :=_aec .(*_gc .PdfObjectArray );
_gcd {_cfag =_gb ;}else if _cfab ,_gdac :=_aec .(*_gc .PdfObjectName );_gdac {_fcc =_cfab ;};if _fcc !=nil {switch *_fcc {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":return *_fcc ,nil ;
case "\u0050a\u0074\u0074\u0065\u0072\u006e":return *_fcc ,nil ;};};if _cfag !=nil &&_cfag .Len ()> 0{if _baf ,_fea :=_cfag .Get (0).(*_gc .PdfObjectName );_fea {switch *_baf {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":if _cfag .Len ()==1{return *_baf ,nil ;
};case "\u0043a\u006c\u0047\u0072\u0061\u0079","\u0043\u0061\u006c\u0052\u0047\u0042","\u004c\u0061\u0062":return *_baf ,nil ;case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064","\u0050a\u0074\u0074\u0065\u0072\u006e","\u0049n\u0064\u0065\u0078\u0065\u0064":return *_baf ,nil ;
case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e","\u0044e\u0076\u0069\u0063\u0065\u004e":return *_baf ,nil ;};};};return "",nil ;};type Page struct{_cff int ;Object *_gc .PdfObjectDictionary ;_gec *Document ;};func (_dcb *Catalog )GetOutputIntents ()(*OutputIntents ,bool ){_fa :=_dcb .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");
if _fa ==nil {return nil ,false ;};_adf ,_be :=_gc .GetIndirect (_fa );if !_be {return nil ,false ;};_cbe ,_da :=_gc .GetArray (_adf .PdfObject );if !_da {return nil ,false ;};return &OutputIntents {_ba :_adf ,_deg :_cbe ,_fe :_dcb ._f },true ;};func (_eb *OutputIntents )Get (i int )(OutputIntent ,bool ){if _eb ._deg ==nil {return OutputIntent {},false ;
};if i >=_eb ._deg .Len (){return OutputIntent {},false ;};_addd :=_eb ._deg .Get (i );_fce ,_ceb :=_gc .GetIndirect (_addd );if !_ceb {_bg ,_cg :=_gc .GetDict (_addd );return OutputIntent {Object :_bg },_cg ;};_cbef ,_dec :=_gc .GetDict (_fce .PdfObject );
return OutputIntent {Object :_cbef },_dec ;};func (_eg *Catalog )HasMetadata ()bool {_cab :=_eg .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061");return _cab !=nil ;};func (_aae Page )GetResources ()(*_gc .PdfObjectDictionary ,bool ){return _gc .GetDict (_aae .Object .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));
};type Content struct{Stream *_gc .PdfObjectStream ;_fdc int ;_cfad Page ;};func (_fgb Page )FindXObjectImages ()([]*Image ,error ){_beg ,_fcb :=_fgb .GetResourcesXObject ();if !_fcb {return nil ,nil ;};var _bf []*Image ;var _ace error ;_cbeb :=map[*_gc .PdfObjectStream ]int {};
_bafd :=map[*_gc .PdfObjectStream ]struct{}{};var _acd int ;for _ ,_eec :=range _beg .Keys (){_afg ,_bdd :=_gc .GetStream (_beg .Get (_eec ));if !_bdd {continue ;};if _ ,_gg :=_cbeb [_afg ];_gg {continue ;};_aaee ,_ecg :=_gc .GetName (_afg .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if !_ecg ||_aaee .String ()!="\u0049\u006d\u0061g\u0065"{continue ;};_fed :=Image {BitsPerComponent :8,Stream :_afg ,Name :string (_eec )};if _fed .Colorspace ,_ace =_gee (_afg .PdfObjectDictionary .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));
_ace !=nil {_ad .Log .Error ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073",_ace );continue ;};if _aba ,_fec :=_gc .GetIntVal (_afg .PdfObjectDictionary .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));
_fec {_fed .BitsPerComponent =_aba ;};if _fcdb ,_agd :=_gc .GetIntVal (_afg .PdfObjectDictionary .Get ("\u0057\u0069\u0064t\u0068"));_agd {_fed .Width =_fcdb ;};if _dcaa ,_fcea :=_gc .GetIntVal (_afg .PdfObjectDictionary .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));
_fcea {_fed .Height =_dcaa ;};if _ef ,_dfa :=_gc .GetStream (_afg .Get ("\u0053\u004d\u0061s\u006b"));_dfa {_fed .SMask =&ImageSMask {Image :&_fed ,Stream :_ef };_bafd [_ef ]=struct{}{};};switch _fed .Colorspace {case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":_fed .ColorComponents =3;
case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":_fed .ColorComponents =1;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":_fed .ColorComponents =4;default:_fed .ColorComponents =-1;};_cbeb [_afg ]=_acd ;_bf =append (_bf ,&_fed );
_acd ++;};var _fgbc []int ;for _ ,_cdd :=range _bf {if _cdd .SMask !=nil {_cgg ,_fda :=_cbeb [_cdd .SMask .Stream ];if _fda {_fgbc =append (_fgbc ,_cgg );};};};_fgc :=make ([]*Image ,len (_bf )-len (_fgbc ));_acd =0;_fac :for _beb ,_ebb :=range _bf {for _ ,_fbd :=range _fgbc {if _beb ==_fbd {continue _fac ;
};};_fgc [_acd ]=_ebb ;_acd ++;};return _bf ,nil ;};func (_cadb *Document )GetPages ()([]Page ,bool ){_ea ,_bag :=_cadb .FindCatalog ();if !_bag {return nil ,false ;};return _ea .GetPages ();};type ImageSMask struct{Image *Image ;Stream *_gc .PdfObjectStream ;
};func (_dee Page )GetContents ()([]Content ,bool ){_gdb ,_eac :=_gc .GetArray (_dee .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if !_eac {_faeg ,_aga :=_gc .GetStream (_dee .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
if !_aga {return nil ,false ;};return []Content {{Stream :_faeg ,_cfad :_dee ,_fdc :0}},true ;};_fg :=make ([]Content ,_gdb .Len ());for _dfg ,_cfd :=range _gdb .Elements (){_adb ,_cdc :=_gc .GetStream (_cfd );if !_cdc {continue ;};_fg [_dfg ]=Content {Stream :_adb ,_cfad :_dee ,_fdc :_dfg };
};return _fg ,true ;};func (_e *Catalog )GetPages ()([]Page ,bool ){_dc ,_aa :=_gc .GetDict (_e .Object .Get ("\u0050\u0061\u0067e\u0073"));if !_aa {return nil ,false ;};_ac ,_c :=_gc .GetArray (_dc .Get ("\u004b\u0069\u0064\u0073"));if !_c {return nil ,false ;
};_ee :=make ([]Page ,_ac .Len ());for _ca ,_acc :=range _ac .Elements (){_dcd ,_ae :=_gc .GetDict (_acc );if !_ae {continue ;};_ee [_ca ]=Page {Object :_dcd ,_cff :_ca +1,_gec :_e ._f };};return _ee ,true ;};func (_fb *Catalog )SetMetadata (data []byte )error {_ec ,_cad :=_gc .MakeStream (data ,nil );
if _cad !=nil {return _cad ;};_ec .Set ("\u0054\u0079\u0070\u0065",_gc .MakeName ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));_ec .Set ("\u0053u\u0062\u0074\u0079\u0070\u0065",_gc .MakeName ("\u0058\u004d\u004c"));_fb .Object .Set ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061",_ec );
_fb ._f .Objects =append (_fb ._f .Objects ,_ec );return nil ;};func (_gd *Catalog )SetMarkInfo (mi _gc .PdfObject ){_fd :=_gc .MakeIndirectObject (mi );_gd .Object .Set ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f",_fd );_gd ._f .Objects =append (_gd ._f .Objects ,_fd );
};func (_gdc Page )FindXObjectForms ()[]*_gc .PdfObjectStream {_ege ,_fga :=_gdc .GetResourcesXObject ();if !_fga {return nil ;};_fge :=map[*_gc .PdfObjectStream ]struct{}{};var _bdf func (_bfb *_gc .PdfObjectDictionary ,_ece map[*_gc .PdfObjectStream ]struct{});
_bdf =func (_cgb *_gc .PdfObjectDictionary ,_dgb map[*_gc .PdfObjectStream ]struct{}){for _ ,_eee :=range _cgb .Keys (){_gdad ,_eaa :=_gc .GetStream (_cgb .Get (_eee ));if !_eaa {continue ;};if _ ,_geb :=_dgb [_gdad ];_geb {continue ;};_dfc ,_ffc :=_gc .GetName (_gdad .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if !_ffc ||_dfc .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_dgb [_gdad ]=struct{}{};_bbb ,_ffc :=_gc .GetDict (_gdad .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_ffc {continue ;};_cbaf ,_ffa :=_gc .GetDict (_bbb .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));
if _ffa {_bdf (_cbaf ,_dgb );};};};_bdf (_ege ,_fge );var _cgd []*_gc .PdfObjectStream ;for _dfge :=range _fge {_cgd =append (_cgd ,_dfge );};return _cgd ;};func (_fdd *Content )SetData (data []byte )error {_gbc ,_gge :=_gc .MakeStream (data ,_gc .NewFlateEncoder ());
if _gge !=nil {return _gge ;};_abb ,_abae :=_gc .GetArray (_fdd ._cfad .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if !_abae &&_fdd ._fdc ==0{_fdd ._cfad .Object .Set ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073",_gbc );}else {if _gge =_abb .Set (_fdd ._fdc ,_gbc );
_gge !=nil {return _gge ;};};_fdd ._cfad ._gec .Objects =append (_fdd ._cfad ._gec .Objects ,_gbc );return nil ;};func (_eeb *Page )Number ()int {return _eeb ._cff };func (_cfa *Document )AddIndirectObject (indirect *_gc .PdfIndirectObject ){for _ ,_cc :=range _cfa .Objects {if _cc ==indirect {return ;
};};_cfa .Objects =append (_cfa .Objects ,indirect );};func (_fgf Content )GetData ()([]byte ,error ){_dgg ,_eacd :=_gc .NewEncoderFromStream (_fgf .Stream );if _eacd !=nil {return nil ,_eacd ;};_bff ,_eacd :=_dgg .DecodeStream (_fgf .Stream );if _eacd !=nil {return nil ,_eacd ;
};return _bff ,nil ;};func (_b *Catalog )SetVersion (){_b .Object .Set ("\u0056e\u0072\u0073\u0069\u006f\u006e",_gc .MakeName (_d .Sprintf ("\u0025\u0064\u002e%\u0064",_b ._f .Version .Major ,_b ._f .Version .Minor )));};func (_fad *OutputIntents )Add (oi _gc .PdfObject )error {_dca ,_dcc :=oi .(*_gc .PdfObjectDictionary );
if !_dcc {return _g .New ("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};if _bee ,_dgd :=_gc .GetStream (_dca .Get ("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065"));_dgd {_fad ._fe .Objects =append (_fad ._fe .Objects ,_bee );};_fbe ,_ce :=oi .(*_gc .PdfIndirectObject );
if !_ce {_fbe =_gc .MakeIndirectObject (oi );};if _fad ._deg ==nil {_fad ._deg =_gc .MakeArray (_fbe );}else {_fad ._deg .Append (_fbe );};_fad ._fe .Objects =append (_fad ._fe .Objects ,_fbe );return nil ;};func (_ge *Document )FindCatalog ()(*Catalog ,bool ){var _fbec *_gc .PdfObjectDictionary ;
for _ ,_gda :=range _ge .Objects {_bb ,_ag :=_gc .GetDict (_gda );if !_ag {continue ;};if _ceba ,_af :=_gc .GetName (_bb .Get ("\u0054\u0079\u0070\u0065"));_af &&*_ceba =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_fbec =_bb ;break ;};};if _fbec ==nil {return nil ,false ;
};return &Catalog {Object :_fbec ,_f :_ge },true ;};type Document struct{ID [2]string ;Version _gc .Version ;Objects []_gc .PdfObject ;Info _gc .PdfObject ;Crypt *_gc .PdfCrypt ;UseHashBasedID bool ;};func (_fc *Catalog )NewOutputIntents ()*OutputIntents {return &OutputIntents {_fe :_fc ._f }};
func (_cba Page )GetResourcesXObject ()(*_gc .PdfObjectDictionary ,bool ){_ed ,_abc :=_cba .GetResources ();if !_abc {return nil ,false ;};return _gc .GetDict (_ed .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));};func (_cb *Catalog )GetMetadata ()(*_gc .PdfObjectStream ,bool ){_acg ,_de :=_gc .GetStream (_cb .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));
return _acg ,_de ;};func (_db *Catalog )SetOutputIntents (outputIntents *OutputIntents ){if _cf :=_db .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");_cf !=nil {for _aea ,_add :=range _db ._f .Objects {if _add ==_cf {if outputIntents ._ba ==_cf {return ;
};_db ._f .Objects =append (_db ._f .Objects [:_aea ],_db ._f .Objects [_aea +1:]...);break ;};};};_ga :=outputIntents ._ba ;if _ga ==nil {_ga =_gc .MakeIndirectObject (outputIntents ._deg );};_db .Object .Set ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073",_ga );
_db ._f .Objects =append (_db ._f .Objects ,_ga );};type OutputIntent struct{Object *_gc .PdfObjectDictionary ;};func (_ecb *OutputIntents )Len ()int {return _ecb ._deg .Len ()};