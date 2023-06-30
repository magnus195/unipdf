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

package optimize ;import (_a "bytes";_dg "crypto/md5";_ec "errors";_c "github.com/unidoc/unipdf/v3/common";_bdb "github.com/unidoc/unipdf/v3/contentstream";_ge "github.com/unidoc/unipdf/v3/core";_cc "github.com/unidoc/unipdf/v3/extractor";_b "github.com/unidoc/unipdf/v3/internal/imageutil";
_ce "github.com/unidoc/unipdf/v3/internal/textencoding";_g "github.com/unidoc/unipdf/v3/model";_bd "github.com/unidoc/unitype";_d "golang.org/x/image/draw";_dc "math";);func _ffe (_af *_ge .PdfObjectStream )error {_aa ,_df :=_ge .DecodeStream (_af );if _df !=nil {return _df ;
};_gc :=_bdb .NewContentStreamParser (string (_aa ));_bb ,_df :=_gc .Parse ();if _df !=nil {return _df ;};_bb =_ff (_bb );_cda :=_bb .Bytes ();if len (_cda )>=len (_aa ){return nil ;};_gfg ,_df :=_ge .MakeStream (_bb .Bytes (),_ge .NewFlateEncoder ());
if _df !=nil {return _df ;};_af .Stream =_gfg .Stream ;_af .Merge (_gfg .PdfObjectDictionary );return nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_dbabd *ObjectStreams )Optimize (objects []_ge .PdfObject )(_gaac []_ge .PdfObject ,_gcff error ){_fafb :=&_ge .PdfObjectStreams {};_ceab :=make ([]_ge .PdfObject ,0,len (objects ));for _ ,_gaeb :=range objects {if _eae ,_eeg :=_gaeb .(*_ge .PdfIndirectObject );
_eeg &&_eae .GenerationNumber ==0{_fafb .Append (_gaeb );}else {_ceab =append (_ceab ,_gaeb );};};if _fafb .Len ()==0{return _ceab ,nil ;};_gaac =make ([]_ge .PdfObject ,0,len (_ceab )+_fafb .Len ()+1);if _fafb .Len ()> 1{_gaac =append (_gaac ,_fafb );
};_gaac =append (_gaac ,_fafb .Elements ()...);_gaac =append (_gaac ,_ceab ...);return _gaac ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_faef *CombineIdenticalIndirectObjects )Optimize (objects []_ge .PdfObject )(_fcee []_ge .PdfObject ,_abc error ){_bae (objects );_afc :=make (map[_ge .PdfObject ]_ge .PdfObject );_ebd :=make (map[_ge .PdfObject ]struct{});_dfa :=make (map[string ][]*_ge .PdfIndirectObject );
for _ ,_gdd :=range objects {_eff ,_afca :=_gdd .(*_ge .PdfIndirectObject );if !_afca {continue ;};if _dfd ,_cbe :=_eff .PdfObject .(*_ge .PdfObjectDictionary );_cbe {if _bgb ,_bfb :=_dfd .Get ("\u0054\u0079\u0070\u0065").(*_ge .PdfObjectName );_bfb &&*_bgb =="\u0050\u0061\u0067\u0065"{continue ;
};_gbd :=_dg .New ();_gbd .Write ([]byte (_dfd .WriteString ()));_dcb :=string (_gbd .Sum (nil ));_dfa [_dcb ]=append (_dfa [_dcb ],_eff );};};for _ ,_bgg :=range _dfa {if len (_bgg )< 2{continue ;};_gbc :=_bgg [0];for _aacfg :=1;_aacfg < len (_bgg );_aacfg ++{_edf :=_bgg [_aacfg ];
_afc [_edf ]=_gbc ;_ebd [_edf ]=struct{}{};};};_fcee =make ([]_ge .PdfObject ,0,len (objects )-len (_ebd ));for _ ,_bgc :=range objects {if _ ,_dbc :=_ebd [_bgc ];_dbc {continue ;};_fcee =append (_fcee ,_bgc );};_cab (_fcee ,_afc );return _fcee ,nil ;};


// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{_da []_g .Optimizer };

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ImageUpperPPI float64 ;};

// New creates a optimizers chain from options.
func New (options Options )*Chain {_cbf :=new (Chain );if options .CleanFonts ||options .SubsetFonts {_cbf .Append (&CleanFonts {Subset :options .SubsetFonts });};if options .CleanContentstream {_cbf .Append (new (CleanContentstream ));};if options .ImageUpperPPI > 0{_bdcb :=new (ImagePPI );
_bdcb .ImageUpperPPI =options .ImageUpperPPI ;_cbf .Append (_bdcb );};if options .ImageQuality > 0{_bfa :=new (Image );_bfa .ImageQuality =options .ImageQuality ;_cbf .Append (_bfa );};if options .CombineDuplicateDirectObjects {_cbf .Append (new (CombineDuplicateDirectObjects ));
};if options .CombineDuplicateStreams {_cbf .Append (new (CombineDuplicateStreams ));};if options .CombineIdenticalIndirectObjects {_cbf .Append (new (CombineIdenticalIndirectObjects ));};if options .UseObjectStreams {_cbf .Append (new (ObjectStreams ));
};if options .CompressStreams {_cbf .Append (new (CompressStreams ));};return _cbf ;};

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{};func _ff (_bf *_bdb .ContentStreamOperations )*_bdb .ContentStreamOperations {if _bf ==nil {return nil ;};_fa :=_bdb .ContentStreamOperations {};for _ ,_de :=range *_bf {switch _de .Operand {case "\u0042\u0044\u0043","\u0042\u004d\u0043","\u0045\u004d\u0043":continue ;
case "\u0054\u006d":if len (_de .Params )==6{if _bff ,_geb :=_ge .GetNumbersAsFloat (_de .Params );_geb ==nil {if _bff [0]==1&&_bff [1]==0&&_bff [2]==0&&_bff [3]==1{_de =&_bdb .ContentStreamOperation {Params :[]_ge .PdfObject {_de .Params [4],_de .Params [5]},Operand :"\u0054\u0064"};
};};};};_fa =append (_fa ,_de );};return &_fa ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_eb *CombineDuplicateDirectObjects )Optimize (objects []_ge .PdfObject )(_ceg []_ge .PdfObject ,_ad error ){_bae (objects );_dfb :=make (map[string ][]*_ge .PdfObjectDictionary );var _fad func (_fca *_ge .PdfObjectDictionary );_fad =func (_febb *_ge .PdfObjectDictionary ){for _ ,_cgc :=range _febb .Keys (){_dee :=_febb .Get (_cgc );
if _deaa ,_fda :=_dee .(*_ge .PdfObjectDictionary );_fda {_caaa :=_dg .New ();_caaa .Write ([]byte (_deaa .WriteString ()));_gdeg :=string (_caaa .Sum (nil ));_dfb [_gdeg ]=append (_dfb [_gdeg ],_deaa );_fad (_deaa );};};};for _ ,_ffd :=range objects {_fec ,_fgf :=_ffd .(*_ge .PdfIndirectObject );
if !_fgf {continue ;};if _gbb ,_bacf :=_fec .PdfObject .(*_ge .PdfObjectDictionary );_bacf {_fad (_gbb );};};_ccgg :=make ([]_ge .PdfObject ,0,len (_dfb ));_afg :=make (map[_ge .PdfObject ]_ge .PdfObject );for _ ,_ccc :=range _dfb {if len (_ccc )< 2{continue ;
};_aag :=_ge .MakeDict ();_aag .Merge (_ccc [0]);_dbg :=_ge .MakeIndirectObject (_aag );_ccgg =append (_ccgg ,_dbg );for _gea :=0;_gea < len (_ccc );_gea ++{_febbe :=_ccc [_gea ];_afg [_febbe ]=_dbg ;};};_ceg =make ([]_ge .PdfObject ,len (objects ));copy (_ceg ,objects );
_ceg =append (_ccgg ,_ceg ...);_cab (_ceg ,_afg );return _ceg ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gagb *Image )Optimize (objects []_ge .PdfObject )(_aeg []_ge .PdfObject ,_fdgb error ){if _gagb .ImageQuality <=0{return objects ,nil ;};_gbae :=_cfc (objects );if len (_gbae )==0{return objects ,nil ;};_bcd :=make (map[_ge .PdfObject ]_ge .PdfObject );
_eac :=make (map[_ge .PdfObject ]struct{});for _ ,_dbec :=range _gbae {_bbcd :=_dbec .Stream .Get ("\u0053\u004d\u0061s\u006b");_eac [_bbcd ]=struct{}{};};for _fcd ,_deeg :=range _gbae {_gdga :=_deeg .Stream ;if _ ,_abbf :=_eac [_gdga ];_abbf {continue ;
};_cfgf ,_defa :=_g .NewXObjectImageFromStream (_gdga );if _defa !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_defa );continue ;};switch _cfgf .Filter .(type ){case *_ge .JBIG2Encoder :continue ;case *_ge .CCITTFaxEncoder :continue ;
};_badd ,_defa :=_cfgf .ToImage ();if _defa !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_defa );continue ;};_adb :=_ge .NewDCTEncoder ();_adb .ColorComponents =_badd .ColorComponents ;_adb .Quality =_gagb .ImageQuality ;
_adb .BitsPerComponent =_deeg .BitsPerComponent ;_adb .Width =_deeg .Width ;_adb .Height =_deeg .Height ;_ebc ,_defa :=_adb .EncodeBytes (_badd .Data );if _defa !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_defa );
continue ;};var _edeg _ge .StreamEncoder ;_edeg =_adb ;{_eaa :=_ge .NewFlateEncoder ();_eacd :=_ge .NewMultiEncoder ();_eacd .AddEncoder (_eaa );_eacd .AddEncoder (_adb );_gga ,_dgbg :=_eacd .EncodeBytes (_badd .Data );if _dgbg !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_dgbg );
continue ;};if len (_gga )< len (_ebc ){_c .Log .Trace ("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029",len (_ebc ),len (_gga ),len (_gdga .Stream ));
_ebc =_gga ;_edeg =_eacd ;};};_eagf :=len (_gdga .Stream );if _eagf < len (_ebc ){continue ;};_bab :=&_ge .PdfObjectStream {Stream :_ebc };_bab .PdfObjectReference =_gdga .PdfObjectReference ;_bab .PdfObjectDictionary =_ge .MakeDict ();_bab .Merge (_gdga .PdfObjectDictionary );
_bab .Merge (_edeg .MakeStreamDict ());_bab .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_ge .MakeInteger (int64 (len (_ebc ))));_bcd [_gdga ]=_bab ;_gbae [_fcd ].Stream =_bab ;};_aeg =make ([]_ge .PdfObject ,len (objects ));copy (_aeg ,objects );_cab (_aeg ,_bcd );
return _aeg ,nil ;};func _bae (_ggeb []_ge .PdfObject ){for _dde ,_ffb :=range _ggeb {switch _dadb :=_ffb .(type ){case *_ge .PdfIndirectObject :_dadb .ObjectNumber =int64 (_dde +1);_dadb .GenerationNumber =0;case *_ge .PdfObjectStream :_dadb .ObjectNumber =int64 (_dde +1);
_dadb .GenerationNumber =0;case *_ge .PdfObjectStreams :_dadb .ObjectNumber =int64 (_dde +1);_dadb .GenerationNumber =0;};};};

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ImageQuality int ;};func _fga (_ea []_ge .PdfObject )(_bfd map[*_ge .PdfObjectStream ]struct{},_efa error ){_bfd =map[*_ge .PdfObjectStream ]struct{}{};_abf :=map[*_g .PdfFont ]struct{}{};_gdg :=_efb (_ea );for _ ,_gfd :=range _gdg ._dfgc {_aaa ,_feb :=_ge .GetDict (_gfd .PdfObject );
if !_feb {continue ;};_cgb ,_feb :=_ge .GetDict (_aaa .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_feb {continue ;};_ba ,_ :=_fgbb (_aaa .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_bfg ,_egb :=_g .NewPdfPageResourcesFromDict (_cgb );
if _egb !=nil {return nil ,_egb ;};_ed :=[]content {{_dba :_ba ,_ccg :_bfg }};_cbc :=_dbf (_aaa .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));if _cbc !=nil {_ed =append (_ed ,_cbc ...);};for _ ,_gee :=range _ed {_eeb ,_gdc :=_cc .NewFromContents (_gee ._dba ,_gee ._ccg );
if _gdc !=nil {return nil ,_gdc ;};_gg ,_ ,_ ,_gdc :=_eeb .ExtractPageText ();if _gdc !=nil {return nil ,_gdc ;};for _ ,_eed :=range _gg .Marks ().Elements (){if _eed .Font ==nil {continue ;};if _ ,_ceeb :=_abf [_eed .Font ];!_ceeb {_abf [_eed .Font ]=struct{}{};
};};};};_deb :=map[*_ge .PdfObjectStream ][]*_g .PdfFont {};for _fc :=range _abf {_egg :=_fc .FontDescriptor ();if _egg ==nil ||_egg .FontFile2 ==nil {continue ;};_caa ,_bag :=_ge .GetStream (_egg .FontFile2 );if !_bag {continue ;};_deb [_caa ]=append (_deb [_caa ],_fc );
};for _bdd :=range _deb {var _cdf []rune ;var _dfea []_bd .GlyphIndex ;for _ ,_egc :=range _deb [_bdd ]{switch _be :=_egc .Encoder ().(type ){case *_ce .IdentityEncoder :_gag :=_be .RegisteredRunes ();_bg :=make ([]_bd .GlyphIndex ,len (_gag ));for _gde ,_dfg :=range _gag {_bg [_gde ]=_bd .GlyphIndex (_dfg );
};_dfea =append (_dfea ,_bg ...);case *_ce .TrueTypeFontEncoder :_ag :=_be .RegisteredRunes ();_cdf =append (_cdf ,_ag ...);case _ce .SimpleEncoder :_bbb :=_be .Charcodes ();for _ ,_gcf :=range _bbb {_dbe ,_eaf :=_be .CharcodeToRune (_gcf );if !_eaf {_c .Log .Debug ("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064",_gcf );
continue ;};_cdf =append (_cdf ,_dbe );};};};_efa =_agb (_bdd ,_cdf ,_dfea );if _efa !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076",_efa );
return nil ,_efa ;};_bfd [_bdd ]=struct{}{};};return _bfd ,nil ;};

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{};

// Options describes PDF optimization parameters.
type Options struct{CombineDuplicateStreams bool ;CombineDuplicateDirectObjects bool ;ImageUpperPPI float64 ;ImageQuality int ;UseObjectStreams bool ;CombineIdenticalIndirectObjects bool ;CompressStreams bool ;CleanFonts bool ;SubsetFonts bool ;CleanContentstream bool ;
};func _cab (_dfbe []_ge .PdfObject ,_fdd map[_ge .PdfObject ]_ge .PdfObject ){if len (_fdd )==0{return ;};for _gbfd ,_bgd :=range _dfbe {if _gcfa ,_fgdd :=_fdd [_bgd ];_fgdd {_dfbe [_gbfd ]=_gcfa ;continue ;};_fdd [_bgd ]=_bgd ;switch _bfgf :=_bgd .(type ){case *_ge .PdfObjectArray :_cbde :=make ([]_ge .PdfObject ,_bfgf .Len ());
copy (_cbde ,_bfgf .Elements ());_cab (_cbde ,_fdd );for _abad ,_cae :=range _cbde {_bfgf .Set (_abad ,_cae );};case *_ge .PdfObjectStreams :_cab (_bfgf .Elements (),_fdd );case *_ge .PdfObjectStream :_agd :=[]_ge .PdfObject {_bfgf .PdfObjectDictionary };
_cab (_agd ,_fdd );_bfgf .PdfObjectDictionary =_agd [0].(*_ge .PdfObjectDictionary );case *_ge .PdfObjectDictionary :_bbg :=_bfgf .Keys ();_ebff :=make ([]_ge .PdfObject ,len (_bbg ));for _bde ,_edff :=range _bbg {_ebff [_bde ]=_bfgf .Get (_edff );};_cab (_ebff ,_fdd );
for _dfeb ,_ecb :=range _bbg {_bfgf .Set (_ecb ,_ebff [_dfeb ]);};case *_ge .PdfIndirectObject :_dabf :=[]_ge .PdfObject {_bfgf .PdfObject };_cab (_dabf ,_fdd );_bfgf .PdfObject =_dabf [0];};};};type content struct{_dba string ;_ccg *_g .PdfPageResources ;
};

// Optimize optimizes PDF objects to decrease PDF size.
func (_cggc *ImagePPI )Optimize (objects []_ge .PdfObject )(_gdgb []_ge .PdfObject ,_aab error ){if _cggc .ImageUpperPPI <=0{return objects ,nil ;};_gbee :=_cfc (objects );if len (_gbee )==0{return objects ,nil ;};_cegf :=make (map[_ge .PdfObject ]struct{});
for _ ,_dbfa :=range _gbee {_fbb :=_dbfa .Stream .PdfObjectDictionary .Get ("\u0053\u004d\u0061s\u006b");_cegf [_fbb ]=struct{}{};};_feg :=make (map[*_ge .PdfObjectStream ]*imageInfo );for _ ,_dae :=range _gbee {_feg [_dae .Stream ]=_dae ;};var _efga *_ge .PdfObjectDictionary ;
for _ ,_dcd :=range objects {if _fcdb ,_gge :=_ge .GetDict (_dcd );_efga ==nil &&_gge {if _dagf ,_acfd :=_ge .GetName (_fcdb .Get ("\u0054\u0079\u0070\u0065"));_acfd &&*_dagf =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_efga =_fcdb ;};};};if _efga ==nil {return objects ,nil ;
};_bec ,_acg :=_ge .GetDict (_efga .Get ("\u0050\u0061\u0067e\u0073"));if !_acg {return objects ,nil ;};_ebdc ,_bbd :=_ge .GetArray (_bec .Get ("\u004b\u0069\u0064\u0073"));if !_bbd {return objects ,nil ;};for _ ,_ffa :=range _ebdc .Elements (){_dbab :=make (map[string ]*imageInfo );
_edfa ,_bacc :=_ge .GetDict (_ffa );if !_bacc {continue ;};_fgd ,_ :=_fgbb (_edfa .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));if len (_fgd )==0{continue ;};_bfea ,_dcfa :=_ge .GetDict (_edfa .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));
if !_dcfa {continue ;};_bdg ,_edd :=_g .NewPdfPageResourcesFromDict (_bfea );if _edd !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073\u0020-\u0020\u0069\u0067\u006e\u006fr\u0069\u006eg\u003a\u0020\u0025\u0076",_edd );
continue ;};_cfab ,_bbcb :=_ge .GetDict (_bfea .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));if !_bbcb {continue ;};_fdc :=_cfab .Keys ();for _ ,_eebb :=range _fdc {if _bfgg ,_aff :=_ge .GetStream (_cfab .Get (_eebb ));_aff {if _dbb ,_age :=_feg [_bfgg ];
_age {_dbab [string (_eebb )]=_dbb ;};};};_caad :=_bdb .NewContentStreamParser (_fgd );_feda ,_edd :=_caad .Parse ();if _edd !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_edd );continue ;};_dgc :=_bdb .NewContentStreamProcessor (*_feda );
_dgc .AddHandler (_bdb .HandlerConditionEnumAllOperands ,"",func (_cfbg *_bdb .ContentStreamOperation ,_gccc _bdb .GraphicsState ,_dff *_g .PdfPageResources )error {switch _cfbg .Operand {case "\u0044\u006f":if len (_cfbg .Params )!=1{_c .Log .Debug ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020w\u0069\u0074\u0068\u0020\u006c\u0065\u006e\u0028\u0070\u0061ra\u006d\u0073\u0029 \u0021=\u0020\u0031");
return nil ;};_agab ,_fcaf :=_ge .GetName (_cfbg .Params [0]);if !_fcaf {_c .Log .Debug ("\u0045\u0052\u0052O\u0052\u003a\u0020\u0049\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0044\u006f\u0020\u0077\u0069\u0074\u0068\u0020\u006e\u006f\u006e\u0020\u004e\u0061\u006d\u0065\u0020p\u0061\u0072\u0061\u006d\u0065\u0074\u0065\u0072");
return nil ;};if _cba ,_cbd :=_dbab [string (*_agab )];_cbd {_fffg :=_gccc .CTM .ScalingFactorX ();_fab :=_gccc .CTM .ScalingFactorY ();_cde ,_fdggb :=_fffg /72.0,_fab /72.0;_cbg ,_gad :=float64 (_cba .Width )/_cde ,float64 (_cba .Height )/_fdggb ;if _cde ==0||_fdggb ==0{_cbg =72.0;
_gad =72.0;};_cba .PPI =_dc .Max (_cba .PPI ,_cbg );_cba .PPI =_dc .Max (_cba .PPI ,_gad );};};return nil ;});_edd =_dgc .Process (_bdg );if _edd !=nil {_c .Log .Debug ("E\u0052\u0052\u004f\u0052 p\u0072o\u0063\u0065\u0073\u0073\u0069n\u0067\u003a\u0020\u0025\u002b\u0076",_edd );
continue ;};};for _ ,_bddc :=range _gbee {if _ ,_efad :=_cegf [_bddc .Stream ];_efad {continue ;};if _bddc .PPI <=_cggc .ImageUpperPPI {continue ;};_ffg ,_dad :=_g .NewXObjectImageFromStream (_bddc .Stream );if _dad !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_dad );
continue ;};var _bfef imageModifications ;_bfef .Scale =_cggc .ImageUpperPPI /_bddc .PPI ;if _bddc .BitsPerComponent ==1&&_bddc .ColorComponents ==1{_add :=_dc .Round (_bddc .PPI /_cggc .ImageUpperPPI );_acb :=_b .NextPowerOf2 (uint (_add ));if _b .InDelta (float64 (_acb ),1/_bfef .Scale ,0.3){_bfef .Scale =float64 (1)/float64 (_acb );
};if _ ,_cdca :=_ffg .Filter .(*_ge .JBIG2Encoder );!_cdca {_bfef .Encoding =_ge .NewJBIG2Encoder ();};};if _dad =_ace (_ffg ,_bfef );_dad !=nil {_c .Log .Debug ("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_dad );
continue ;};_bfef .Encoding =nil ;if _cfe ,_eede :=_ge .GetStream (_bddc .Stream .PdfObjectDictionary .Get ("\u0053\u004d\u0061s\u006b"));_eede {_bea ,_dbda :=_g .NewXObjectImageFromStream (_cfe );if _dbda !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_dbda );
continue ;};if _dbda =_ace (_bea ,_bfef );_dbda !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0025\u002b\u0076",_dbda );continue ;};};};return objects ,nil ;};func _cfc (_fea []_ge .PdfObject )[]*imageInfo {_adc :=_ge .PdfObjectName ("\u0053u\u0062\u0074\u0079\u0070\u0065");
_cfgc :=make (map[*_ge .PdfObjectStream ]struct{});var _feaf []*imageInfo ;for _ ,_bagb :=range _fea {_baf ,_cafc :=_ge .GetStream (_bagb );if !_cafc {continue ;};if _ ,_faf :=_cfgc [_baf ];_faf {continue ;};_cfgc [_baf ]=struct{}{};_dcf :=_baf .PdfObjectDictionary .Get (_adc );
_aebg ,_cafc :=_ge .GetName (_dcf );if !_cafc ||string (*_aebg )!="\u0049\u006d\u0061g\u0065"{continue ;};_fdf :=&imageInfo {Stream :_baf ,BitsPerComponent :8};if _gebd ,_cdcd :=_ge .GetIntVal (_baf .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));
_cdcd {_fdf .BitsPerComponent =_gebd ;};if _cef ,_fgec :=_ge .GetIntVal (_baf .Get ("\u0057\u0069\u0064t\u0068"));_fgec {_fdf .Width =_cef ;};if _eeba ,_gaa :=_ge .GetIntVal (_baf .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));_gaa {_fdf .Height =_eeba ;
};_fffc ,_ebb :=_g .NewPdfColorspaceFromPdfObject (_baf .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));if _ebb !=nil {_c .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ebb );continue ;};if _fffc ==nil {_cfac ,_ecd :=_ge .GetName (_baf .Get ("\u0046\u0069\u006c\u0074\u0065\u0072"));
if _ecd {switch _cfac .String (){case "\u0043\u0043\u0049\u0054\u0054\u0046\u0061\u0078\u0044e\u0063\u006f\u0064\u0065","J\u0042\u0049\u0047\u0032\u0044\u0065\u0063\u006f\u0064\u0065":_fffc =_g .NewPdfColorspaceDeviceGray ();_fdf .BitsPerComponent =1;};
};};switch _bbf :=_fffc .(type ){case *_g .PdfColorspaceDeviceRGB :_fdf .ColorComponents =3;case *_g .PdfColorspaceDeviceGray :_fdf .ColorComponents =1;default:_c .Log .Debug ("\u004f\u0070\u0074\u0069\u006d\u0069\u007aa\u0074\u0069\u006fn\u0020\u0069\u0073 \u006e\u006ft\u0020\u0073\u0075\u0070\u0070\u006fr\u0074ed\u0020\u0066\u006f\u0072\u0020\u0063\u006f\u006c\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0054\u0020\u002d\u0020\u0073\u006b\u0069\u0070",_bbf );
continue ;};_feaf =append (_feaf ,_fdf );};return _feaf ;};type imageModifications struct{Scale float64 ;Encoding _ge .StreamEncoder ;};func _fgbb (_cgge _ge .PdfObject )(_gecc string ,_becc []_ge .PdfObject ){var _abcd _a .Buffer ;switch _efc :=_cgge .(type ){case *_ge .PdfIndirectObject :_becc =append (_becc ,_efc );
_cgge =_efc .PdfObject ;};switch _eea :=_cgge .(type ){case *_ge .PdfObjectStream :if _fbeg ,_ddgd :=_ge .DecodeStream (_eea );_ddgd ==nil {_abcd .Write (_fbeg );_becc =append (_becc ,_eea );};case *_ge .PdfObjectArray :for _ ,_facb :=range _eea .Elements (){switch _bef :=_facb .(type ){case *_ge .PdfObjectStream :if _ded ,_ebfff :=_ge .DecodeStream (_bef );
_ebfff ==nil {_abcd .Write (_ded );_becc =append (_becc ,_bef );};};};};return _abcd .String (),_becc ;};

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{};

// Optimize optimizes PDF objects to decrease PDF size.
func (_geg *Chain )Optimize (objects []_ge .PdfObject )(_gb []_ge .PdfObject ,_gefd error ){_f :=objects ;for _ ,_dd :=range _geg ._da {_gf ,_dce :=_dd .Optimize (_f );if _dce !=nil {_c .Log .Debug ("\u0045\u0052\u0052OR\u0020\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u002b\u0076",_dce );
continue ;};_f =_gf ;};return _f ,nil ;};

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{};

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct{

// Subset embedded fonts if encountered (if true).
// Otherwise attempts to reduce the font program.
Subset bool ;};

// Append appends optimizers to the chain.
func (_gef *Chain )Append (optimizers ..._g .Optimizer ){_gef ._da =append (_gef ._da ,optimizers ...)};

// Optimize optimizes PDF objects to decrease PDF size.
func (_cgg *CompressStreams )Optimize (objects []_ge .PdfObject )(_feff []_ge .PdfObject ,_eda error ){_feff =make ([]_ge .PdfObject ,len (objects ));copy (_feff ,objects );for _ ,_dgag :=range objects {_faec ,_ccb :=_ge .GetStream (_dgag );if !_ccb {continue ;
};if _eag :=_faec .Get ("\u0046\u0069\u006c\u0074\u0065\u0072");_eag !=nil {if _ ,_acaa :=_ge .GetName (_eag );_acaa {continue ;};if _dbaa ,_fdb :=_ge .GetArray (_eag );_fdb &&_dbaa .Len ()> 0{continue ;};};_dace :=_ge .NewFlateEncoder ();var _edc []byte ;
_edc ,_eda =_dace .EncodeBytes (_faec .Stream );if _eda !=nil {return _feff ,_eda ;};_fgg :=_dace .MakeStreamDict ();if len (_edc )+len (_fgg .WriteString ())< len (_faec .Stream ){_faec .Stream =_edc ;_faec .PdfObjectDictionary .Merge (_fgg );_faec .PdfObjectDictionary .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_ge .MakeInteger (int64 (len (_faec .Stream ))));
};};return _feff ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_aac *CleanFonts )Optimize (objects []_ge .PdfObject )(_fge []_ge .PdfObject ,_aec error ){var _dge map[*_ge .PdfObjectStream ]struct{};if _aac .Subset {var _acf error ;_dge ,_acf =_fga (objects );if _acf !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004fR\u003a\u0020\u0046\u0061\u0069\u006c\u0065\u0064\u0020\u0073u\u0062s\u0065\u0074\u0074\u0069\u006e\u0067\u003a \u0025\u0076",_acf );
return nil ,_acf ;};};for _ ,_aee :=range objects {_baa ,_ceb :=_ge .GetStream (_aee );if !_ceb {continue ;};if _ ,_gec :=_dge [_baa ];_gec {continue ;};_fgaa ,_bgf :=_ge .NewEncoderFromStream (_baa );if _bgf !=nil {_c .Log .Debug ("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067",_bgf );
continue ;};_acab ,_bgf :=_fgaa .DecodeStream (_baa );if _bgf !=nil {_c .Log .Debug ("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067",_bgf );
continue ;};if len (_acab )< 4{continue ;};_accf :=string (_acab [:4]);if _accf =="\u004f\u0054\u0054\u004f"{continue ;};if _accf !="\u0000\u0001\u0000\u0000"&&_accf !="\u0074\u0072\u0075\u0065"{continue ;};_dag ,_bgf :=_bd .Parse (_a .NewReader (_acab ));
if _bgf !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_bgf );continue ;};_bgf =_dag .Optimize ();
if _bgf !=nil {_c .Log .Debug ("\u0045\u0052RO\u0052\u0020\u004fp\u0074\u0069\u006d\u0069zin\u0067 f\u006f\u006e\u0074\u003a\u0020\u0025\u0076 -\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067",_bgf );continue ;};var _acfa _a .Buffer ;_bgf =_dag .Write (&_acfa );
if _bgf !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_bgf );continue ;};if _acfa .Len ()> len (_acab ){_c .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
continue ;};_cgf ,_bgf :=_ge .MakeStream (_acfa .Bytes (),_ge .NewFlateEncoder ());if _bgf !=nil {continue ;};*_baa =*_cgf ;_baa .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_ge .MakeInteger (int64 (_acfa .Len ())));};return objects ,nil ;};type objectStructure struct{_fcc *_ge .PdfObjectDictionary ;
_daa *_ge .PdfObjectDictionary ;_dfgc []*_ge .PdfIndirectObject ;};func _agb (_fef *_ge .PdfObjectStream ,_gcea []rune ,_egbd []_bd .GlyphIndex )error {_fef ,_gbeg :=_ge .GetStream (_fef );if !_gbeg {_c .Log .Debug ("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067");
return _ec .New ("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_fb ,_cfa :=_ge .DecodeStream (_fef );if _cfa !=nil {_c .Log .Debug ("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_cfa );
return _cfa ;};_cfg ,_cfa :=_bd .Parse (_a .NewReader (_fb ));if _cfa !=nil {_c .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074",len (_fef .Stream ));
return _cfa ;};_gcb :=_egbd ;if len (_gcea )> 0{_fff :=_cfg .LookupRunes (_gcea );_gcb =append (_gcb ,_fff ...);};_cfg ,_cfa =_cfg .SubsetKeepIndices (_gcb );if _cfa !=nil {_c .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v",_cfa );
return _cfa ;};var _bdf _a .Buffer ;_cfa =_cfg .Write (&_bdf );if _cfa !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_cfa );return _cfa ;};if _bdf .Len ()> len (_fb ){_c .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
return nil ;};_bfe ,_cfa :=_ge .MakeStream (_bdf .Bytes (),_ge .NewFlateEncoder ());if _cfa !=nil {_c .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_cfa );return _cfa ;
};*_fef =*_bfe ;_fef .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_ge .MakeInteger (int64 (_bdf .Len ())));return nil ;};func _acaf (_fag *_g .Image ,_aeef float64 )(*_g .Image ,error ){_cfb ,_faa :=_fag .ToGoImage ();if _faa !=nil {return nil ,_faa ;
};var _bfdgg _b .Image ;_afgd ,_deae :=_cfb .(*_b .Monochrome );if _deae {if _faa =_afgd .ResolveDecode ();_faa !=nil {return nil ,_faa ;};_bfdgg ,_faa =_afgd .Scale (_aeef );if _faa !=nil {return nil ,_faa ;};}else {_fbeb :=int (_dc .RoundToEven (float64 (_fag .Width )*_aeef ));
_cag :=int (_dc .RoundToEven (float64 (_fag .Height )*_aeef ));_bfdgg ,_faa =_b .NewImage (_fbeb ,_cag ,int (_fag .BitsPerComponent ),_fag .ColorComponents ,nil ,nil ,nil );if _faa !=nil {return nil ,_faa ;};_d .CatmullRom .Scale (_bfdgg ,_bfdgg .Bounds (),_cfb ,_cfb .Bounds (),_d .Over ,&_d .Options {});
};_egbf :=_bfdgg .Base ();_fdgg :=&_g .Image {Width :int64 (_egbf .Width ),Height :int64 (_egbf .Height ),BitsPerComponent :int64 (_egbf .BitsPerComponent ),ColorComponents :_egbf .ColorComponents ,Data :_egbf .Data };_fdgg .SetDecode (_egbf .Decode );
_fdgg .SetAlpha (_egbf .Alpha );return _fdgg ,nil ;};

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{};func _efb (_ceae []_ge .PdfObject )objectStructure {_bda :=objectStructure {};_gaf :=false ;for _ ,_gbfe :=range _ceae {switch _fcg :=_gbfe .(type ){case *_ge .PdfIndirectObject :_deaec ,_cgd :=_ge .GetDict (_fcg );
if !_cgd {continue ;};_cgff ,_cgd :=_ge .GetName (_deaec .Get ("\u0054\u0079\u0070\u0065"));if !_cgd {continue ;};switch _cgff .String (){case "\u0043a\u0074\u0061\u006c\u006f\u0067":_bda ._fcc =_deaec ;_gaf =true ;};};if _gaf {break ;};};if !_gaf {return _bda ;
};_fgb ,_fac :=_ge .GetDict (_bda ._fcc .Get ("\u0050\u0061\u0067e\u0073"));if !_fac {return _bda ;};_bda ._daa =_fgb ;_eafd ,_fac :=_ge .GetArray (_fgb .Get ("\u004b\u0069\u0064\u0073"));if !_fac {return _bda ;};for _ ,_daag :=range _eafd .Elements (){_aegg ,_ccgc :=_ge .GetIndirect (_daag );
if !_ccgc {break ;};_bda ._dfgc =append (_bda ._dfgc ,_aegg );};return _bda ;};

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{};func _ace (_adg *_g .XObjectImage ,_ebf imageModifications )error {_gae ,_dggf :=_adg .ToImage ();if _dggf !=nil {return _dggf ;};if _ebf .Scale !=0{_gae ,_dggf =_acaf (_gae ,_ebf .Scale );if _dggf !=nil {return _dggf ;
};};if _ebf .Encoding !=nil {_adg .Filter =_ebf .Encoding ;};_adg .Decode =nil ;switch _feae :=_adg .Filter .(type ){case *_ge .FlateEncoder :if _feae .Predictor !=1&&_feae .Predictor !=11{_feae .Predictor =1;};};if _dggf =_adg .SetImage (_gae ,nil );_dggf !=nil {_c .Log .Debug ("\u0045\u0072\u0072or\u0020\u0073\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0076",_dggf );
return _dggf ;};_adg .ToPdfObject ();return nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_fe *CleanContentstream )Optimize (objects []_ge .PdfObject )(_cf []_ge .PdfObject ,_fae error ){_ab :=map[*_ge .PdfObjectStream ]struct{}{};var _ga []*_ge .PdfObjectStream ;_gca :=func (_bc *_ge .PdfObjectStream ){if _ ,_cdc :=_ab [_bc ];!_cdc {_ab [_bc ]=struct{}{};
_ga =append (_ga ,_bc );};};_gfc :=map[_ge .PdfObject ]bool {};_gfgg :=map[_ge .PdfObject ]bool {};for _ ,_gcc :=range objects {switch _dab :=_gcc .(type ){case *_ge .PdfIndirectObject :switch _gbf :=_dab .PdfObject .(type ){case *_ge .PdfObjectDictionary :if _cee ,_cdb :=_ge .GetName (_gbf .Get ("\u0054\u0079\u0070\u0065"));
!_cdb ||_cee .String ()!="\u0050\u0061\u0067\u0065"{continue ;};if _fdg ,_ac :=_ge .GetStream (_gbf .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_ac {_gca (_fdg );}else if _ffee ,_cg :=_ge .GetArray (_gbf .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
_cg {var _cb []*_ge .PdfObjectStream ;for _ ,_gfb :=range _ffee .Elements (){if _acc ,_aca :=_ge .GetStream (_gfb );_aca {_cb =append (_cb ,_acc );};};if len (_cb )> 0{var _aaf _a .Buffer ;for _ ,_ae :=range _cb {if _db ,_ee :=_ge .DecodeStream (_ae );
_ee ==nil {_aaf .Write (_db );};_gfc [_ae ]=true ;};_ef ,_gce :=_ge .MakeStream (_aaf .Bytes (),_ge .NewFlateEncoder ());if _gce !=nil {return nil ,_gce ;};_gfgg [_ef ]=true ;_gbf .Set ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073",_ef );_gca (_ef );
};};};case *_ge .PdfObjectStream :if _bfff ,_defg :=_ge .GetName (_dab .Get ("\u0054\u0079\u0070\u0065"));!_defg ||_bfff .String ()!="\u0058O\u0062\u006a\u0065\u0063\u0074"{continue ;};if _bcf ,_fg :=_ge .GetName (_dab .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
!_fg ||_bcf .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_gca (_dab );};};for _ ,_afe :=range _ga {_fae =_ffe (_afe );if _fae !=nil {return nil ,_fae ;};};_cf =nil ;for _ ,_ca :=range objects {if _gfc [_ca ]{continue ;};_cf =append (_cf ,_ca );};
for _gbe :=range _gfgg {_cf =append (_cf ,_gbe );};return _cf ,nil ;};type imageInfo struct{BitsPerComponent int ;ColorComponents int ;Width int ;Height int ;Stream *_ge .PdfObjectStream ;PPI float64 ;};func _dbf (_gcd _ge .PdfObject )[]content {if _gcd ==nil {return nil ;
};_cgeg ,_dgb :=_ge .GetArray (_gcd );if !_dgb {_c .Log .Debug ("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");return nil ;};var _dgg []content ;for _ ,_bfdg :=range _cgeg .Elements (){_bdc ,_efg :=_ge .GetDict (_bfdg );
if !_efg {_c .Log .Debug ("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073");continue ;};_bfc ,_efg :=_ge .GetDict (_bdc .Get ("\u0041\u0050"));
if !_efg {_c .Log .Debug ("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_gfdc :=_ge .TraceToDirectObject (_bfc .Get ("\u004e"));if _gfdc ==nil {_c .Log .Debug ("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067");
continue ;};var _bbbc *_ge .PdfObjectStream ;switch _dacf :=_gfdc .(type ){case *_ge .PdfObjectDictionary :_dea ,_abe :=_ge .GetName (_bdc .Get ("\u0041\u0053"));if !_abe {_c .Log .Debug ("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");
continue ;};_bbbc ,_abe =_ge .GetStream (_dacf .Get (*_dea ));if !_abe {_c .Log .Debug ("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg");continue ;};case *_ge .PdfObjectStream :_bbbc =_dacf ;
};if _bbbc ==nil {_c .Log .Debug ("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067");continue ;};_dddg ,_agc :=_g .NewXObjectFormFromStream (_bbbc );
if _agc !=nil {_c .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_agc );continue ;};_febgc ,_agc :=_dddg .GetContentStream ();
if _agc !=nil {_c .Log .Debug ("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076",_agc );continue ;};_dgg =append (_dgg ,content {_dba :string (_febgc ),_ccg :_dddg .Resources });
};return _dgg ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_aga *CombineDuplicateStreams )Optimize (objects []_ge .PdfObject )(_gba []_ge .PdfObject ,_caf error ){_cfd :=make (map[_ge .PdfObject ]_ge .PdfObject );_ada :=make (map[_ge .PdfObject ]struct{});_aacf :=make (map[string ][]*_ge .PdfObjectStream );
for _ ,_fadb :=range objects {if _agba ,_eebc :=_fadb .(*_ge .PdfObjectStream );_eebc {_aeb :=_dg .New ();_aeb .Write (_agba .Stream );_aeb .Write ([]byte (_agba .PdfObjectDictionary .WriteString ()));_ecaa :=string (_aeb .Sum (nil ));_aacf [_ecaa ]=append (_aacf [_ecaa ],_agba );
};};for _ ,_aed :=range _aacf {if len (_aed )< 2{continue ;};_aba :=_aed [0];for _fefc :=1;_fefc < len (_aed );_fefc ++{_fgfg :=_aed [_fefc ];_cfd [_fgfg ]=_aba ;_ada [_fgfg ]=struct{}{};};};_gba =make ([]_ge .PdfObject ,0,len (objects )-len (_ada ));for _ ,_dgf :=range objects {if _ ,_fce :=_ada [_dgf ];
_fce {continue ;};_gba =append (_gba ,_dgf );};_cab (_gba ,_cfd );return _gba ,nil ;};