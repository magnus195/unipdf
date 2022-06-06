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

package docutil ;import (_d "errors";_e "fmt";_gc "github.com/unidoc/unipdf/v3/common";_dg "github.com/unidoc/unipdf/v3/core";);func (_df *Catalog )GetMetadata ()(*_dg .PdfObjectStream ,bool ){_eb ,_c :=_dg .GetStream (_df .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));
return _eb ,_c ;};func (_bc *Catalog )GetMarkInfo ()(*_dg .PdfObjectDictionary ,bool ){_cgf ,_ag :=_dg .GetDict (_bc .Object .Get ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f"));return _cgf ,_ag ;};type Catalog struct{Object *_dg .PdfObjectDictionary ;
_f *Document ;};func (_ecc Page )FindXObjectImages ()([]*Image ,error ){_bga ,_bbg :=_ecc .GetResourcesXObject ();if !_bbg {return nil ,nil ;};var _bgd []*Image ;var _ga error ;_fac :=map[*_dg .PdfObjectStream ]int {};_bfed :=map[*_dg .PdfObjectStream ]struct{}{};
var _cd int ;for _ ,_agd :=range _bga .Keys (){_bde ,_eda :=_dg .GetStream (_bga .Get (_agd ));if !_eda {continue ;};if _ ,_baf :=_fac [_bde ];_baf {continue ;};_dac ,_cdf :=_dg .GetName (_bde .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));if !_cdf ||_dac .String ()!="\u0049\u006d\u0061g\u0065"{continue ;
};_cae :=Image {BitsPerComponent :8,Stream :_bde ,Name :string (_agd )};if _cae .Colorspace ,_ga =_eg (_bde .PdfObjectDictionary .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));_ga !=nil {_gc .Log .Error ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073",_ga );
continue ;};if _geeb ,_bgac :=_dg .GetIntVal (_bde .PdfObjectDictionary .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));_bgac {_cae .BitsPerComponent =_geeb ;};if _bca ,_de :=_dg .GetIntVal (_bde .PdfObjectDictionary .Get ("\u0057\u0069\u0064t\u0068"));
_de {_cae .Width =_bca ;};if _aaac ,_gfc :=_dg .GetIntVal (_bde .PdfObjectDictionary .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));_gfc {_cae .Height =_aaac ;};if _caa ,_acb :=_dg .GetStream (_bde .Get ("\u0053\u004d\u0061s\u006b"));_acb {_cae .SMask =&ImageSMask {Image :&_cae ,Stream :_caa };
_bfed [_caa ]=struct{}{};};switch _cae .Colorspace {case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":_cae .ColorComponents =3;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":_cae .ColorComponents =1;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":_cae .ColorComponents =4;
default:_cae .ColorComponents =-1;};_fac [_bde ]=_cd ;_bgd =append (_bgd ,&_cae );_cd ++;};var _dge []int ;for _ ,_dfb :=range _bgd {if _dfb .SMask !=nil {_ddb ,_fee :=_fac [_dfb .SMask .Stream ];if _fee {_dge =append (_dge ,_ddb );};};};_gbe :=make ([]*Image ,len (_bgd )-len (_dge ));
_cd =0;_geg :for _fafa ,_edd :=range _bgd {for _ ,_gcc :=range _dge {if _fafa ==_gcc {continue _geg ;};};_gbe [_cd ]=_edd ;_cd ++;};return _bgd ,nil ;};func (_fe *Document )AddStream (stream *_dg .PdfObjectStream ){for _ ,_agf :=range _fe .Objects {if _agf ==stream {return ;
};};_fe .Objects =append (_fe .Objects ,stream );};func (_bge Page )FindXObjectForms ()[]*_dg .PdfObjectStream {_acab ,_gce :=_bge .GetResourcesXObject ();if !_gce {return nil ;};_gcca :=map[*_dg .PdfObjectStream ]struct{}{};var _bda func (_cdg *_dg .PdfObjectDictionary ,_feg map[*_dg .PdfObjectStream ]struct{});
_bda =func (_cge *_dg .PdfObjectDictionary ,_dbf map[*_dg .PdfObjectStream ]struct{}){for _ ,_gbc :=range _cge .Keys (){_bbd ,_bfc :=_dg .GetStream (_cge .Get (_gbc ));if !_bfc {continue ;};if _ ,_fec :=_dbf [_bbd ];_fec {continue ;};_dfgac ,_cb :=_dg .GetName (_bbd .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
if !_cb ||_dfgac .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_dbf [_bbd ]=struct{}{};_egc ,_cb :=_dg .GetDict (_bbd .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_cb {continue ;};_cfe ,_adad :=_dg .GetDict (_egc .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));
if _adad {_bda (_cfe ,_dbf );};};};_bda (_acab ,_gcca );var _gccg []*_dg .PdfObjectStream ;for _gec :=range _gcca {_gccg =append (_gccg ,_gec );};return _gccg ;};func (_ed *Catalog )SetVersion (){_ed .Object .Set ("\u0056e\u0072\u0073\u0069\u006f\u006e",_dg .MakeName (_e .Sprintf ("\u0025\u0064\u002e%\u0064",_ed ._f .Version .Major ,_ed ._f .Version .Minor )));
};type OutputIntent struct{Object *_dg .PdfObjectDictionary ;};func (_ddg Page )GetResourcesXObject ()(*_dg .PdfObjectDictionary ,bool ){_ece ,_eef :=_ddg .GetResources ();if !_eef {return nil ,false ;};return _dg .GetDict (_ece .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));
};func (_aaf *Content )SetData (data []byte )error {_dgf ,_feee :=_dg .MakeStream (data ,_dg .NewFlateEncoder ());if _feee !=nil {return _feee ;};_eac ,_ecef :=_dg .GetArray (_aaf ._dc .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
if !_ecef &&_aaf ._bcde ==0{_aaf ._dc .Object .Set ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073",_dgf );}else {if _feee =_eac .Set (_aaf ._bcde ,_dgf );_feee !=nil {return _feee ;};};_aaf ._dc ._cea .Objects =append (_aaf ._dc ._cea .Objects ,_dgf );
return nil ;};type Page struct{_bdg int ;Object *_dg .PdfObjectDictionary ;_cea *Document ;};func (_afd *Catalog )NewOutputIntents ()*OutputIntents {return &OutputIntents {_ge :_afd ._f }};type Image struct{Name string ;Width int ;Height int ;Colorspace _dg .PdfObjectName ;
ColorComponents int ;BitsPerComponent int ;SMask *ImageSMask ;Stream *_dg .PdfObjectStream ;};func (_ebb Page )GetContents ()([]Content ,bool ){_aac ,_bbb :=_dg .GetArray (_ebb .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if !_bbb {_gge ,_eaf :=_dg .GetStream (_ebb .Object .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
if !_eaf {return nil ,false ;};return []Content {{Stream :_gge ,_dc :_ebb ,_bcde :0}},true ;};_eag :=make ([]Content ,_aac .Len ());for _afc ,_ffg :=range _aac .Elements (){_ggg ,_bdge :=_dg .GetStream (_ffg );if !_bdge {continue ;};_eag [_afc ]=Content {Stream :_ggg ,_dc :_ebb ,_bcde :_afc };
};return _eag ,true ;};func (_aaa *Catalog )GetOutputIntents ()(*OutputIntents ,bool ){_bf :=_aaa .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");if _bf ==nil {return nil ,false ;};_faf ,_abab :=_dg .GetIndirect (_bf );
if !_abab {return nil ,false ;};_da ,_dgc :=_dg .GetArray (_faf .PdfObject );if !_dgc {return nil ,false ;};return &OutputIntents {_gea :_faf ,_ea :_da ,_ge :_aaa ._f },true ;};type Document struct{ID [2]string ;Version _dg .Version ;Objects []_dg .PdfObject ;
Info _dg .PdfObject ;Crypt *_dg .PdfCrypt ;UseHashBasedID bool ;};func (_gb *Catalog )SetMarkInfo (mi _dg .PdfObject ){_db :=_dg .MakeIndirectObject (mi );_gb .Object .Set ("\u004d\u0061\u0072\u006b\u0049\u006e\u0066\u006f",_db );_gb ._f .Objects =append (_gb ._f .Objects ,_db );
};func _eg (_gead _dg .PdfObject )(_dg .PdfObjectName ,error ){var _bd *_dg .PdfObjectName ;var _gee *_dg .PdfObjectArray ;if _cgc ,_eea :=_gead .(*_dg .PdfIndirectObject );_eea {if _eee ,_egf :=_cgc .PdfObject .(*_dg .PdfObjectArray );_egf {_gee =_eee ;
}else if _cga ,_bfb :=_cgc .PdfObject .(*_dg .PdfObjectName );_bfb {_bd =_cga ;};}else if _dfga ,_daf :=_gead .(*_dg .PdfObjectArray );_daf {_gee =_dfga ;}else if _dbbd ,_faa :=_gead .(*_dg .PdfObjectName );_faa {_bd =_dbbd ;};if _bd !=nil {switch *_bd {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":return *_bd ,nil ;
case "\u0050a\u0074\u0074\u0065\u0072\u006e":return *_bd ,nil ;};};if _gee !=nil &&_gee .Len ()> 0{if _ged ,_eeae :=_gee .Get (0).(*_dg .PdfObjectName );_eeae {switch *_ged {case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079","\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B","\u0044\u0065\u0076\u0069\u0063\u0065\u0043\u004d\u0059\u004b":if _gee .Len ()==1{return *_ged ,nil ;
};case "\u0043a\u006c\u0047\u0072\u0061\u0079","\u0043\u0061\u006c\u0052\u0047\u0042","\u004c\u0061\u0062":return *_ged ,nil ;case "\u0049\u0043\u0043\u0042\u0061\u0073\u0065\u0064","\u0050a\u0074\u0074\u0065\u0072\u006e","\u0049n\u0064\u0065\u0078\u0065\u0064":return *_ged ,nil ;
case "\u0053\u0065\u0070\u0061\u0072\u0061\u0074\u0069\u006f\u006e","\u0044e\u0076\u0069\u0063\u0065\u004e":return *_ged ,nil ;};};};return "",nil ;};func (_cef *Document )FindCatalog ()(*Catalog ,bool ){var _aca *_dg .PdfObjectDictionary ;for _ ,_dbg :=range _cef .Objects {_ace ,_dbb :=_dg .GetDict (_dbg );
if !_dbb {continue ;};if _ef ,_efc :=_dg .GetName (_ace .Get ("\u0054\u0079\u0070\u0065"));_efc &&*_ef =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_aca =_ace ;break ;};};if _aca ==nil {return nil ,false ;};return &Catalog {Object :_aca ,_f :_cef },true ;
};func (_ada Page )GetResources ()(*_dg .PdfObjectDictionary ,bool ){return _dg .GetDict (_ada .Object .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));};func (_ced *Document )GetPages ()([]Page ,bool ){_bbc ,_aae :=_ced .FindCatalog ();if !_aae {return nil ,false ;
};return _bbc .GetPages ();};func (_efa *Page )Number ()int {return _efa ._bdg };func (_ee *Catalog )HasMetadata ()bool {_afe :=_ee .Object .Get ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061");return _afe !=nil ;};func (_fg *Catalog )SetMetadata (data []byte )error {_cg ,_ab :=_dg .MakeStream (data ,nil );
if _ab !=nil {return _ab ;};_cg .Set ("\u0054\u0079\u0070\u0065",_dg .MakeName ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061"));_cg .Set ("\u0053u\u0062\u0074\u0079\u0070\u0065",_dg .MakeName ("\u0058\u004d\u004c"));_fg .Object .Set ("\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061",_cg );
_fg ._f .Objects =append (_fg ._f .Objects ,_cg );return nil ;};func (_abg Content )GetData ()([]byte ,error ){_ggf ,_dgcb :=_dg .NewEncoderFromStream (_abg .Stream );if _dgcb !=nil {return nil ,_dgcb ;};_adb ,_dgcb :=_ggf .DecodeStream (_abg .Stream );
if _dgcb !=nil {return nil ,_dgcb ;};return _adb ,nil ;};func (_dgd *Catalog )GetPages ()([]Page ,bool ){_a ,_b :=_dg .GetDict (_dgd .Object .Get ("\u0050\u0061\u0067e\u0073"));if !_b {return nil ,false ;};_ac ,_aa :=_dg .GetArray (_a .Get ("\u004b\u0069\u0064\u0073"));
if !_aa {return nil ,false ;};_fc :=make ([]Page ,_ac .Len ());for _ae ,_fa :=range _ac .Elements (){_af ,_bg :=_dg .GetDict (_fa );if !_bg {continue ;};_fc [_ae ]=Page {Object :_af ,_bdg :_ae +1,_cea :_dgd ._f };};return _fc ,true ;};type OutputIntents struct{_ea *_dg .PdfObjectArray ;
_ge *Document ;_gea *_dg .PdfIndirectObject ;};func (_ba *OutputIntents )Len ()int {return _ba ._ea .Len ()};func (_ggc *Document )AddIndirectObject (indirect *_dg .PdfIndirectObject ){for _ ,_gf :=range _ggc .Objects {if _gf ==indirect {return ;};};_ggc .Objects =append (_ggc .Objects ,indirect );
};type ImageSMask struct{Image *Image ;Stream *_dg .PdfObjectStream ;};type Content struct{Stream *_dg .PdfObjectStream ;_bcde int ;_dc Page ;};func (_ff *OutputIntents )Add (oi _dg .PdfObject )error {_bfe ,_ffb :=oi .(*_dg .PdfObjectDictionary );if !_ffb {return _d .New ("\u0069\u006e\u0070\u0075\u0074\u0020\u006f\u0075\u0074\u0070\u0075\u0074\u0020\u0069\u006e\u0074\u0065\u006et\u0020\u0073\u0068\u006f\u0075\u006c\u0064 \u0062\u0065\u0020\u0061\u006e\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};if _fbe ,_abd :=_dg .GetStream (_bfe .Get ("\u0044\u0065\u0073\u0074\u004f\u0075\u0074\u0070\u0075\u0074\u0050\u0072o\u0066\u0069\u006c\u0065"));_abd {_ff ._ge .Objects =append (_ff ._ge .Objects ,_fbe );};_gef ,_dd :=oi .(*_dg .PdfIndirectObject );if !_dd {_gef =_dg .MakeIndirectObject (oi );
};if _ff ._ea ==nil {_ff ._ea =_dg .MakeArray (_gef );}else {_ff ._ea .Append (_gef );};_ff ._ge .Objects =append (_ff ._ge .Objects ,_gef );return nil ;};func (_bb *OutputIntents )Get (i int )(OutputIntent ,bool ){if _bb ._ea ==nil {return OutputIntent {},false ;
};if i >=_bb ._ea .Len (){return OutputIntent {},false ;};_eed :=_bb ._ea .Get (i );_ec ,_ebf :=_dg .GetIndirect (_eed );if !_ebf {_ce ,_gg :=_dg .GetDict (_eed );return OutputIntent {Object :_ce },_gg ;};_be ,_ddf :=_dg .GetDict (_ec .PdfObject );return OutputIntent {Object :_be },_ddf ;
};func (_fb *Catalog )SetOutputIntents (outputIntents *OutputIntents ){if _cf :=_fb .Object .Get ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073");_cf !=nil {for _bcd ,_ca :=range _fb ._f .Objects {if _ca ==_cf {if outputIntents ._gea ==_cf {return ;
};_fb ._f .Objects =append (_fb ._f .Objects [:_bcd ],_fb ._f .Objects [_bcd +1:]...);break ;};};};_aba :=outputIntents ._gea ;if _aba ==nil {_aba =_dg .MakeIndirectObject (outputIntents ._ea );};_fb .Object .Set ("\u004f\u0075\u0074\u0070\u0075\u0074\u0049\u006e\u0074\u0065\u006e\u0074\u0073",_aba );
_fb ._f .Objects =append (_fb ._f .Objects ,_aba );};