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

package optimize ;import (_ff "bytes";_fa "crypto/md5";_ba "errors";_ad "github.com/unidoc/unipdf/v3/common";_ffe "github.com/unidoc/unipdf/v3/contentstream";_a "github.com/unidoc/unipdf/v3/core";_faa "github.com/unidoc/unipdf/v3/extractor";_e "github.com/unidoc/unipdf/v3/internal/imageutil";
_c "github.com/unidoc/unipdf/v3/internal/textencoding";_d "github.com/unidoc/unipdf/v3/model";_eb "github.com/unidoc/unitype";_b "golang.org/x/image/draw";_fc "math";);

// Options describes PDF optimization parameters.
type Options struct{CombineDuplicateStreams bool ;CombineDuplicateDirectObjects bool ;ImageUpperPPI float64 ;ImageQuality int ;UseObjectStreams bool ;CombineIdenticalIndirectObjects bool ;CompressStreams bool ;CleanFonts bool ;SubsetFonts bool ;CleanContentstream bool ;
};type imageModifications struct{Scale float64 ;Encoding _a .StreamEncoder ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gbd *CombineDuplicateDirectObjects )Optimize (objects []_a .PdfObject )(_bba []_a .PdfObject ,_age error ){_fdf (objects );_febc :=make (map[string ][]*_a .PdfObjectDictionary );var _babd func (_fgb *_a .PdfObjectDictionary );_babd =func (_dedf *_a .PdfObjectDictionary ){for _ ,_face :=range _dedf .Keys (){_gff :=_dedf .Get (_face );
if _gfda ,_dfb :=_gff .(*_a .PdfObjectDictionary );_dfb {_cdc :=_fa .New ();_cdc .Write ([]byte (_gfda .WriteString ()));_bcd :=string (_cdc .Sum (nil ));_febc [_bcd ]=append (_febc [_bcd ],_gfda );_babd (_gfda );};};};for _ ,_affe :=range objects {_ccc ,_fbc :=_affe .(*_a .PdfIndirectObject );
if !_fbc {continue ;};if _caa ,_aabb :=_ccc .PdfObject .(*_a .PdfObjectDictionary );_aabb {_babd (_caa );};};_febe :=make ([]_a .PdfObject ,0,len (_febc ));_bdg :=make (map[_a .PdfObject ]_a .PdfObject );for _ ,_ddb :=range _febc {if len (_ddb )< 2{continue ;
};_fcafa :=_a .MakeDict ();_fcafa .Merge (_ddb [0]);_cgb :=_a .MakeIndirectObject (_fcafa );_febe =append (_febe ,_cgb );for _feba :=0;_feba < len (_ddb );_feba ++{_bfa :=_ddb [_feba ];_bdg [_bfa ]=_cgb ;};};_bba =make ([]_a .PdfObject ,len (objects ));
copy (_bba ,objects );_bba =append (_febe ,_bba ...);_efae (_bba ,_bdg );return _bba ,nil ;};

// CompressStreams compresses uncompressed streams.
// It implements interface model.Optimizer.
type CompressStreams struct{};

// Optimize optimizes PDF objects to decrease PDF size.
func (_da *Image )Optimize (objects []_a .PdfObject )(_afcd []_a .PdfObject ,_dfgf error ){if _da .ImageQuality <=0{return objects ,nil ;};_gcba :=_dea (objects );if len (_gcba )==0{return objects ,nil ;};_bcda :=make (map[_a .PdfObject ]_a .PdfObject );
_efb :=make (map[_a .PdfObject ]struct{});for _ ,_ggba :=range _gcba {_cfb :=_ggba .Stream .PdfObjectDictionary .Get (_a .PdfObjectName ("\u0053\u004d\u0061s\u006b"));_efb [_cfb ]=struct{}{};};for _fcg ,_dbga :=range _gcba {_geg :=_dbga .Stream ;if _ ,_egg :=_efb [_geg ];
_egg {continue ;};_ee ,_bdd :=_a .NewEncoderFromStream (_geg );if _bdd !=nil {_ad .Log .Warning ("\u0045\u0072\u0072\u006f\u0072 \u0067\u0065\u0074\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0066o\u0072\u0020\u0074\u0068\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0025\u0073");
continue ;};_dgfc ,_bdd :=_ee .DecodeStream (_geg );if _bdd !=nil {_ad .Log .Warning ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0063\u006f\u0064\u0065\u0020\u0074\u0068e\u0020i\u006d\u0061\u0067\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u0025\u0073");
continue ;};_efg :=_a .NewDCTEncoder ();_efg .ColorComponents =_dbga .ColorComponents ;_efg .Quality =_da .ImageQuality ;_efg .BitsPerComponent =_dbga .BitsPerComponent ;_efg .Width =_dbga .Width ;_efg .Height =_dbga .Height ;_ggg ,_bdd :=_efg .EncodeBytes (_dgfc );
if _bdd !=nil {_ad .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_bdd );return nil ,_bdd ;};var _dac _a .StreamEncoder ;_dac =_efg ;{_agecg :=_a .NewFlateEncoder ();_ecb :=_a .NewMultiEncoder ();_ecb .AddEncoder (_agecg );_ecb .AddEncoder (_efg );
_fgd ,_abcac :=_ecb .EncodeBytes (_dgfc );if _abcac !=nil {return nil ,_abcac ;};if len (_fgd )< len (_ggg ){_ad .Log .Debug ("\u004d\u0075\u006c\u0074\u0069\u0020\u0065\u006e\u0063\u0020\u0069\u006d\u0070\u0072\u006f\u0076\u0065\u0073\u003a\u0020\u0025\u0064\u0020\u0074o\u0020\u0025\u0064\u0020\u0028o\u0072\u0069g\u0020\u0025\u0064\u0029",len (_ggg ),len (_fgd ),len (_geg .Stream ));
_ggg =_fgd ;_dac =_ecb ;};};_ddf :=len (_geg .Stream );if _ddf < len (_ggg ){continue ;};_gcce :=&_a .PdfObjectStream {Stream :_ggg };_gcce .PdfObjectReference =_geg .PdfObjectReference ;_gcce .PdfObjectDictionary =_a .MakeDict ();_gcce .Merge (_geg .PdfObjectDictionary );
_gcce .Merge (_dac .MakeStreamDict ());_gcce .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_a .MakeInteger (int64 (len (_ggg ))));_bcda [_geg ]=_gcce ;_gcba [_fcg ].Stream =_gcce ;};_afcd =make ([]_a .PdfObject ,len (objects ));copy (_afcd ,objects );_efae (_afcd ,_bcda );
return _afcd ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_fbb *ObjectStreams )Optimize (objects []_a .PdfObject )(_ecccd []_a .PdfObject ,_gbbb error ){_adc :=&_a .PdfObjectStreams {};_dgc :=make ([]_a .PdfObject ,0,len (objects ));for _ ,_facee :=range objects {if _bacad ,_bfgd :=_facee .(*_a .PdfIndirectObject );
_bfgd &&_bacad .GenerationNumber ==0{_adc .Append (_facee );}else {_dgc =append (_dgc ,_facee );};};if _adc .Len ()==0{return _dgc ,nil ;};_ecccd =make ([]_a .PdfObject ,0,len (_dgc )+_adc .Len ()+1);if _adc .Len ()> 1{_ecccd =append (_ecccd ,_adc );};
_ecccd =append (_ecccd ,_adc .Elements ()...);_ecccd =append (_ecccd ,_dgc ...);return _ecccd ,nil ;};func _ec (_dc *_a .PdfObjectStream )error {_ga ,_ag :=_a .DecodeStream (_dc );if _ag !=nil {return _ag ;};_dd :=_ffe .NewContentStreamParser (string (_ga ));
_bf ,_ag :=_dd .Parse ();if _ag !=nil {return _ag ;};_bf =_g (_bf );_bc :=_bf .Bytes ();if len (_bc )>=len (_ga ){return nil ;};_cb ,_ag :=_a .MakeStream (_bf .Bytes (),_a .NewFlateEncoder ());if _ag !=nil {return _ag ;};_dc .Stream =_cb .Stream ;_dc .Merge (_cb .PdfObjectDictionary );
return nil ;};

// Chain allows to use sequence of optimizers.
// It implements interface model.Optimizer.
type Chain struct{_eg []_d .Optimizer };

// ObjectStreams groups PDF objects to object streams.
// It implements interface model.Optimizer.
type ObjectStreams struct{};

// CombineDuplicateDirectObjects combines duplicated direct objects by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateDirectObjects struct{};func _fdf (_bfeb []_a .PdfObject ){for _cecd ,_ddbbf :=range _bfeb {switch _fccg :=_ddbbf .(type ){case *_a .PdfIndirectObject :_fccg .ObjectNumber =int64 (_cecd +1);_fccg .GenerationNumber =0;case *_a .PdfObjectStream :_fccg .ObjectNumber =int64 (_cecd +1);
_fccg .GenerationNumber =0;case *_a .PdfObjectStreams :_fccg .ObjectNumber =int64 (_cecd +1);_fccg .GenerationNumber =0;};};};

// CombineDuplicateStreams combines duplicated streams by its data hash.
// It implements interface model.Optimizer.
type CombineDuplicateStreams struct{};type content struct{_ggfc string ;_ccg *_d .PdfPageResources ;};

// CombineIdenticalIndirectObjects combines identical indirect objects.
// It implements interface model.Optimizer.
type CombineIdenticalIndirectObjects struct{};

// Optimize optimizes PDF objects to decrease PDF size.
func (_fffga *CombineIdenticalIndirectObjects )Optimize (objects []_a .PdfObject )(_bge []_a .PdfObject ,_dcc error ){_fdf (objects );_aca :=make (map[_a .PdfObject ]_a .PdfObject );_bfdb :=make (map[_a .PdfObject ]struct{});_bgb :=make (map[string ][]*_a .PdfIndirectObject );
for _ ,_gfef :=range objects {_fdd ,_aee :=_gfef .(*_a .PdfIndirectObject );if !_aee {continue ;};if _ccb ,_bef :=_fdd .PdfObject .(*_a .PdfObjectDictionary );_bef {if _ffc ,_eaf :=_ccb .Get ("\u0054\u0079\u0070\u0065").(*_a .PdfObjectName );_eaf &&*_ffc =="\u0050\u0061\u0067\u0065"{continue ;
};_ggb :=_fa .New ();_ggb .Write ([]byte (_ccb .WriteString ()));_eafe :=string (_ggb .Sum (nil ));_bgb [_eafe ]=append (_bgb [_eafe ],_fdd );};};for _ ,_bfg :=range _bgb {if len (_bfg )< 2{continue ;};_acg :=_bfg [0];for _fea :=1;_fea < len (_bfg );_fea ++{_gbf :=_bfg [_fea ];
_aca [_gbf ]=_acg ;_bfdb [_gbf ]=struct{}{};};};_bge =make ([]_a .PdfObject ,0,len (objects )-len (_bfdb ));for _ ,_cfd :=range objects {if _ ,_cbc :=_bfdb [_cfd ];_cbc {continue ;};_bge =append (_bge ,_cfd );};_efae (_bge ,_aca );return _bge ,nil ;};

// CleanFonts cleans up embedded fonts, reducing font sizes.
type CleanFonts struct{

// Subset embedded fonts if encountered (if true).
// Otherwise attempts to reduce the font program.
Subset bool ;};func _bad (_fbg []_a .PdfObject )objectStructure {_cfcgc :=objectStructure {};_gbac :=false ;for _ ,_gabb :=range _fbg {switch _ddaa :=_gabb .(type ){case *_a .PdfIndirectObject :_faf ,_dgge :=_a .GetDict (_ddaa );if !_dgge {continue ;};
_cfgeg ,_dgge :=_a .GetName (_faf .Get ("\u0054\u0079\u0070\u0065"));if !_dgge {continue ;};switch _cfgeg .String (){case "\u0043a\u0074\u0061\u006c\u006f\u0067":_cfcgc ._cgbe =_faf ;_gbac =true ;};};if _gbac {break ;};};if !_gbac {return _cfcgc ;};_cdeb ,_abde :=_a .GetDict (_cfcgc ._cgbe .Get ("\u0050\u0061\u0067e\u0073"));
if !_abde {return _cfcgc ;};_cfcgc ._ffcf =_cdeb ;_bfga ,_abde :=_a .GetArray (_cdeb .Get ("\u004b\u0069\u0064\u0073"));if !_abde {return _cfcgc ;};for _ ,_ace :=range _bfga .Elements (){_abbc ,_dff :=_a .GetIndirect (_ace );if !_dff {break ;};_cfcgc ._fcafg =append (_cfcgc ._fcafg ,_abbc );
};return _cfcgc ;};type imageInfo struct{ColorSpace _a .PdfObjectName ;BitsPerComponent int ;ColorComponents int ;Width int ;Height int ;Stream *_a .PdfObjectStream ;PPI float64 ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_cefc *CompressStreams )Optimize (objects []_a .PdfObject )(_gcc []_a .PdfObject ,_abgg error ){_gcc =make ([]_a .PdfObject ,len (objects ));copy (_gcc ,objects );for _ ,_eaaa :=range objects {_cde ,_gdb :=_a .GetStream (_eaaa );if !_gdb {continue ;
};if _gddb :=_cde .Get ("\u0046\u0069\u006c\u0074\u0065\u0072");_gddb !=nil {if _ ,_facb :=_a .GetName (_gddb );_facb {continue ;};if _bbd ,_eab :=_a .GetArray (_gddb );_eab &&_bbd .Len ()> 0{continue ;};};_fgg :=_a .NewFlateEncoder ();var _gddbf []byte ;
_gddbf ,_abgg =_fgg .EncodeBytes (_cde .Stream );if _abgg !=nil {return _gcc ,_abgg ;};_ebcf :=_fgg .MakeStreamDict ();if len (_gddbf )+len (_ebcf .WriteString ())< len (_cde .Stream ){_cde .Stream =_gddbf ;_cde .PdfObjectDictionary .Merge (_ebcf );_cde .PdfObjectDictionary .Set ("\u004c\u0065\u006e\u0067\u0074\u0068",_a .MakeInteger (int64 (len (_cde .Stream ))));
};};return _gcc ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_gag *CombineDuplicateStreams )Optimize (objects []_a .PdfObject )(_cba []_a .PdfObject ,_aegb error ){_bcdg :=make (map[_a .PdfObject ]_a .PdfObject );_gfdb :=make (map[_a .PdfObject ]struct{});_dcg :=make (map[string ][]*_a .PdfObjectStream );for _ ,_abb :=range objects {if _gda ,_agfb :=_abb .(*_a .PdfObjectStream );
_agfb {_fbd :=_fa .New ();_fbd .Write (_gda .Stream );_gca :=string (_fbd .Sum (nil ));_dcg [_gca ]=append (_dcg [_gca ],_gda );};};for _ ,_aega :=range _dcg {if len (_aega )< 2{continue ;};_dfc :=_aega [0];for _dcga :=1;_dcga < len (_aega );_dcga ++{_dfd :=_aega [_dcga ];
_bcdg [_dfd ]=_dfc ;_gfdb [_dfd ]=struct{}{};};};_cba =make ([]_a .PdfObject ,0,len (objects )-len (_gfdb ));for _ ,_baf :=range objects {if _ ,_afc :=_gfdb [_baf ];_afc {continue ;};_cba =append (_cba ,_baf );};_efae (_cba ,_bcdg );return _cba ,nil ;};
func _ggdg (_gge *_a .PdfObjectStream ,_cce []rune ,_abe []_eb .GlyphIndex )error {_gge ,_gfa :=_a .GetStream (_gge );if !_gfa {_ad .Log .Debug ("\u0045\u006d\u0062\u0065\u0064\u0064\u0065\u0064\u0020\u0066\u006f\u006e\u0074\u0020\u006f\u0062\u006a\u0065c\u0074\u0020\u006e\u006f\u0074\u0020\u0066o\u0075\u006e\u0064\u0020\u002d\u002d\u0020\u0041\u0042\u004f\u0052T\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006e\u0067");
return _ba .New ("\u0066\u006f\u006e\u0074fi\u006c\u0065\u0032\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};_gcgb ,_ffd :=_a .DecodeStream (_gge );if _ffd !=nil {_ad .Log .Debug ("\u0044\u0065c\u006f\u0064\u0065 \u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_ffd );
return _ffd ;};_gef ,_ffd :=_eb .Parse (_ff .NewReader (_gcgb ));if _ffd !=nil {_ad .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070\u0061\u0072\u0073\u0069n\u0067\u0020\u0025\u0064\u0020\u0062\u0079\u0074\u0065\u0020f\u006f\u006e\u0074",len (_gge .Stream ));
return _ffd ;};_fca :=_abe ;if len (_cce )> 0{_geb :=_gef .LookupRunes (_cce );_fca =append (_fca ,_geb ...);};_gef ,_ffd =_gef .SubsetKeepIndices (_fca );if _ffd !=nil {_ad .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020s\u0075\u0062\u0073\u0065\u0074t\u0069n\u0067 \u0066\u006f\u006e\u0074\u003a\u0020\u0025v",_ffd );
return _ffd ;};var _bcc _ff .Buffer ;_ffd =_gef .Write (&_bcc );if _ffd !=nil {_ad .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_ffd );return _ffd ;};if _bcc .Len ()> len (_gcgb ){_ad .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
return nil ;};_egc ,_ffd :=_a .MakeStream (_bcc .Bytes (),_a .NewFlateEncoder ());if _ffd !=nil {_ad .Log .Debug ("\u0045\u0052\u0052\u004fR \u0057\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076",_ffd );return _ffd ;
};*_gge =*_egc ;_gge .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_a .MakeInteger (int64 (_bcc .Len ())));return nil ;};func _gdf (_dec *_d .XObjectImage ,_ead imageModifications )error {_fda ,_aage :=_dec .ToImage ();if _aage !=nil {return _aage ;};if _ead .Scale !=0{_fda ,_aage =_abd (_fda ,_ead .Scale );
if _aage !=nil {return _aage ;};};if _ead .Encoding !=nil {_dec .Filter =_ead .Encoding ;};_ggdf :=_a .MakeDict ();_ggdf .Set ("\u0051u\u0061\u006c\u0069\u0074\u0079",_a .MakeInteger (100));_ggdf .Set ("\u0050r\u0065\u0064\u0069\u0063\u0074\u006fr",_a .MakeInteger (1));
_dec .Decode =nil ;if _aage =_dec .SetImage (_fda ,nil );_aage !=nil {return _aage ;};_dec .ToPdfObject ();return nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_bcb *CleanFonts )Optimize (objects []_a .PdfObject )(_fac []_a .PdfObject ,_eca error ){var _bd map[*_a .PdfObjectStream ]struct{};if _bcb .Subset {var _ggf error ;_bd ,_ggf =_agb (objects );if _ggf !=nil {return nil ,_ggf ;};};for _ ,_egf :=range objects {_gb ,_fef :=_a .GetStream (_egf );
if !_fef {continue ;};if _ ,_cfg :=_bd [_gb ];_cfg {continue ;};_bcea ,_ac :=_a .NewEncoderFromStream (_gb );if _ac !=nil {_ad .Log .Debug ("\u0045\u0052RO\u0052\u0020\u0067e\u0074\u0074\u0069\u006eg e\u006eco\u0064\u0065\u0072\u003a\u0020\u0025\u0076 -\u0020\u0069\u0067\u006e\u006f\u0072\u0069n\u0067",_ac );
continue ;};_fb ,_ac :=_bcea .DecodeStream (_gb );if _ac !=nil {_ad .Log .Debug ("\u0044\u0065\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0065r\u0072\u006f\u0072\u0020\u003a\u0020\u0025v\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067",_ac );
continue ;};if len (_fb )< 4{continue ;};_aag :=string (_fb [:4]);if _aag =="\u004f\u0054\u0054\u004f"{continue ;};if _aag !="\u0000\u0001\u0000\u0000"&&_aag !="\u0074\u0072\u0075\u0065"{continue ;};_eag ,_ac :=_eb .Parse (_ff .NewReader (_fb ));if _ac !=nil {_ad .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020P\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_ac );
continue ;};_ac =_eag .Optimize ();if _ac !=nil {continue ;};var _fcd _ff .Buffer ;_ac =_eag .Write (&_fcd );if _ac !=nil {_ad .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020W\u0072\u0069\u0074\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_ac );
continue ;};if _fcd .Len ()> len (_fb ){_ad .Log .Debug ("\u0052\u0065-\u0077\u0072\u0069\u0074\u0074\u0065\u006e\u0020\u0066\u006f\u006e\u0074\u0020\u0069\u0073\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u002d\u0020\u0073\u006b\u0069\u0070");
continue ;};_dfed ,_ac :=_a .MakeStream (_fcd .Bytes (),_a .NewFlateEncoder ());if _ac !=nil {continue ;};*_gb =*_dfed ;_gb .Set ("\u004ce\u006e\u0067\u0074\u0068\u0031",_a .MakeInteger (int64 (_fcd .Len ())));};return objects ,nil ;};

// CleanContentstream cleans up redundant operands in content streams, including Page and XObject Form
// contents. This process includes:
// 1. Marked content operators are removed.
// 2. Some operands are simplified (shorter form).
// TODO: Add more reduction methods and improving the methods for identifying unnecessary operands.
type CleanContentstream struct{};

// Optimize optimizes PDF objects to decrease PDF size.
func (_ce *CleanContentstream )Optimize (objects []_a .PdfObject )(_gd []_a .PdfObject ,_be error ){_add :=map[*_a .PdfObjectStream ]struct{}{};var _ge []*_a .PdfObjectStream ;_fab :=func (_ebc *_a .PdfObjectStream ){if _ ,_dgg :=_add [_ebc ];!_dgg {_add [_ebc ]=struct{}{};
_ge =append (_ge ,_ebc );};};for _ ,_aeg :=range objects {switch _ab :=_aeg .(type ){case *_a .PdfIndirectObject :switch _bg :=_ab .PdfObject .(type ){case *_a .PdfObjectDictionary :if _gc ,_gf :=_a .GetName (_bg .Get ("\u0054\u0079\u0070\u0065"));!_gf ||_gc .String ()!="\u0050\u0061\u0067\u0065"{continue ;
};if _agf ,_ggad :=_a .GetStream (_bg .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_ggad {_fab (_agf );}else if _dga ,_fga :=_a .GetArray (_bg .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));_fga {for _ ,_gdd :=range _dga .Elements (){if _ef ,_dda :=_a .GetStream (_gdd );
_dda {_fab (_ef );};};};};case *_a .PdfObjectStream :if _bff ,_abc :=_a .GetName (_ab .Get ("\u0054\u0079\u0070\u0065"));!_abc ||_bff .String ()!="\u0058O\u0062\u006a\u0065\u0063\u0074"{continue ;};if _aed ,_bac :=_a .GetName (_ab .Get ("\u0053u\u0062\u0074\u0079\u0070\u0065"));
!_bac ||_aed .String ()!="\u0046\u006f\u0072\u006d"{continue ;};_fab (_ab );};};for _ ,_gfd :=range _ge {_be =_ec (_gfd );if _be !=nil {return nil ,_be ;};};return objects ,nil ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_aa *Chain )Optimize (objects []_a .PdfObject )(_ca []_a .PdfObject ,_ed error ){_ca =objects ;for _ ,_fg :=range _aa ._eg {_ca ,_ed =_fg .Optimize (_ca );if _ed !=nil {return _ca ,_ed ;};};return _ca ,nil ;};

// New creates a optimizers chain from options.
func New (options Options )*Chain {_fdad :=new (Chain );if options .CleanFonts ||options .SubsetFonts {_fdad .Append (&CleanFonts {Subset :options .SubsetFonts });};if options .CleanContentstream {_fdad .Append (new (CleanContentstream ));};if options .ImageUpperPPI > 0{_dfdd :=new (ImagePPI );
_dfdd .ImageUpperPPI =options .ImageUpperPPI ;_fdad .Append (_dfdd );};if options .ImageQuality > 0{_aged :=new (Image );_aged .ImageQuality =options .ImageQuality ;_fdad .Append (_aged );};if options .CombineDuplicateDirectObjects {_fdad .Append (new (CombineDuplicateDirectObjects ));
};if options .CombineDuplicateStreams {_fdad .Append (new (CombineDuplicateStreams ));};if options .CombineIdenticalIndirectObjects {_fdad .Append (new (CombineIdenticalIndirectObjects ));};if options .UseObjectStreams {_fdad .Append (new (ObjectStreams ));
};if options .CompressStreams {_fdad .Append (new (CompressStreams ));};return _fdad ;};

// Image optimizes images by rewrite images into JPEG format with quality equals to ImageQuality.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type Image struct{ImageQuality int ;};func _abd (_ageg *_d .Image ,_bfaf float64 )(*_d .Image ,error ){_ccf ,_abea :=_ageg .ToGoImage ();if _abea !=nil {return nil ,_abea ;};var _daf _e .Image ;_baca ,_eccc :=_ccf .(*_e .Monochrome );if _eccc {if _abea =_baca .ResolveDecode ();
_abea !=nil {return nil ,_abea ;};_daf ,_abea =_baca .Scale (_bfaf );if _abea !=nil {return nil ,_abea ;};}else {_ced :=int (_fc .RoundToEven (float64 (_ageg .Width )*_bfaf ));_fbce :=int (_fc .RoundToEven (float64 (_ageg .Height )*_bfaf ));_daf ,_abea =_e .NewImage (_ced ,_fbce ,int (_ageg .BitsPerComponent ),_ageg .ColorComponents ,nil ,nil ,nil );
if _abea !=nil {return nil ,_abea ;};_b .CatmullRom .Scale (_daf ,_daf .Bounds (),_ccf ,_ccf .Bounds (),_b .Over ,&_b .Options {});};_bcf :=_daf .Base ();_beb :=&_d .Image {Width :int64 (_bcf .Width ),Height :int64 (_bcf .Height ),BitsPerComponent :int64 (_bcf .BitsPerComponent ),ColorComponents :_bcf .ColorComponents ,Data :_bcf .Data };
_beb .SetDecode (_bcf .Decode );_beb .SetAlpha (_bcf .Alpha );return _beb ,nil ;};

// Append appends optimizers to the chain.
func (_de *Chain )Append (optimizers ..._d .Optimizer ){_de ._eg =append (_de ._eg ,optimizers ...)};func _beg (_gabf _a .PdfObject )[]content {if _gabf ==nil {return nil ;};_gafd ,_gcb :=_a .GetArray (_gabf );if !_gcb {_ad .Log .Debug ("\u0041\u006e\u006e\u006fts\u0020\u006e\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return nil ;};var _aab []content ;for _ ,_bgd :=range _gafd .Elements (){_ggdc ,_ede :=_a .GetDict (_bgd );if !_ede {_ad .Log .Debug ("I\u0067\u006e\u006f\u0072\u0069\u006eg\u0020\u006e\u006f\u006e\u002d\u0064i\u0063\u0074\u0020\u0065\u006c\u0065\u006de\u006e\u0074\u0020\u0069\u006e\u0020\u0041\u006e\u006e\u006ft\u0073");
continue ;};_fabf ,_ede :=_a .GetDict (_ggdc .Get ("\u0041\u0050"));if !_ede {_ad .Log .Debug ("\u004e\u006f\u0020\u0041P \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_aga :=_a .TraceToDirectObject (_fabf .Get ("\u004e"));
if _aga ==nil {_ad .Log .Debug ("N\u006f\u0020\u004e\u0020en\u0074r\u0079\u0020\u002d\u0020\u0073k\u0069\u0070\u0070\u0069\u006e\u0067");continue ;};var _dgd *_a .PdfObjectStream ;switch _cef :=_aga .(type ){case *_a .PdfObjectDictionary :_acb ,_ada :=_a .GetName (_ggdc .Get ("\u0041\u0053"));
if !_ada {_ad .Log .Debug ("\u004e\u006f\u0020\u0041S \u0065\u006e\u0074\u0072\u0079\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069n\u0067");continue ;};_dgd ,_ada =_a .GetStream (_cef .Get (*_acb ));if !_ada {_ad .Log .Debug ("\u0046o\u0072\u006d\u0020\u006eo\u0074\u0020\u0066\u006f\u0075n\u0064 \u002d \u0073\u006b\u0069\u0070\u0070\u0069\u006eg");
continue ;};case *_a .PdfObjectStream :_dgd =_cef ;};if _dgd ==nil {_ad .Log .Debug ("\u0046\u006f\u0072m\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0028n\u0069\u006c\u0029\u0020\u002d\u0020\u0073\u006b\u0069\u0070\u0070\u0069\u006e\u0067");
continue ;};_ded ,_ggc :=_d .NewXObjectFormFromStream (_dgd );if _ggc !=nil {_ad .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020l\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0066\u006f\u0072\u006d\u003a\u0020%\u0076\u0020\u002d\u0020\u0069\u0067\u006eo\u0072\u0069\u006e\u0067",_ggc );
continue ;};_gae ,_ggc :=_ded .GetContentStream ();if _ggc !=nil {_ad .Log .Debug ("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u006fn\u0074\u0065\u006et\u0073:\u0020\u0025\u0076",_ggc );continue ;};_aab =append (_aab ,content {_ggfc :string (_gae ),_ccg :_ded .Resources });
};return _aab ;};

// Optimize optimizes PDF objects to decrease PDF size.
func (_aea *ImagePPI )Optimize (objects []_a .PdfObject )(_efag []_a .PdfObject ,_fefc error ){if _aea .ImageUpperPPI <=0{return objects ,nil ;};_bda :=_dea (objects );if len (_bda )==0{return objects ,nil ;};_cdg :=make (map[_a .PdfObject ]struct{});for _ ,_fgf :=range _bda {_ebdd :=_fgf .Stream .PdfObjectDictionary .Get (_a .PdfObjectName ("\u0053\u004d\u0061s\u006b"));
_cdg [_ebdd ]=struct{}{};};_deb :=make (map[*_a .PdfObjectStream ]*imageInfo );for _ ,_dae :=range _bda {_deb [_dae .Stream ]=_dae ;};var _afb *_a .PdfObjectDictionary ;for _ ,_dag :=range objects {if _fdc ,_bga :=_a .GetDict (_dag );_afb ==nil &&_bga {if _gged ,_gbfg :=_a .GetName (_fdc .Get (_a .PdfObjectName ("\u0054\u0079\u0070\u0065")));
_gbfg &&*_gged =="\u0043a\u0074\u0061\u006c\u006f\u0067"{_afb =_fdc ;};};};if _afb ==nil {return objects ,nil ;};_egca ,_gcga :=_a .GetDict (_afb .Get (_a .PdfObjectName ("\u0050\u0061\u0067e\u0073")));if !_gcga {return objects ,nil ;};_acf ,_cgdd :=_a .GetArray (_egca .Get (_a .PdfObjectName ("\u004b\u0069\u0064\u0073")));
if !_cgdd {return objects ,nil ;};_efgf :=make (map[string ]*imageInfo );for _ ,_gbe :=range _acf .Elements (){_cfbg ,_bed :=_a .GetDict (_gbe );if !_bed {continue ;};_egda ,_ggdgd :=_a .GetArray (_cfbg .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
if !_ggdgd {continue ;};_cfge ,_gba :=_a .GetDict (_cfbg .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_gba {continue ;};_acfd ,_beca :=_a .GetDict (_cfge .Get ("\u0058O\u0062\u006a\u0065\u0063\u0074"));if !_beca {continue ;};_dfba :=_acfd .Keys ();
for _ ,_bccc :=range _dfba {if _cfa ,_edcd :=_a .GetStream (_acfd .Get (_bccc ));_edcd {if _bfe ,_adb :=_deb [_cfa ];_adb {_efgf [string (_bccc )]=_bfe ;};};};for _ ,_fdce :=range _egda .Elements (){if _gbg ,_eee :=_a .GetStream (_fdce );_eee {_eabg ,_fcgc :=_a .NewEncoderFromStream (_gbg );
if _fcgc !=nil {return nil ,_fcgc ;};_faccb ,_fcgc :=_eabg .DecodeStream (_gbg );if _fcgc !=nil {return nil ,_fcgc ;};_dfa :=_ffe .NewContentStreamParser (string (_faccb ));_gbb ,_fcgc :=_dfa .Parse ();if _fcgc !=nil {return nil ,_fcgc ;};_ggcf ,_fbcf :=1.0,1.0;
for _ ,_bfeg :=range *_gbb {if _bfeg .Operand =="\u0051"{_ggcf ,_fbcf =1.0,1.0;};if _bfeg .Operand =="\u0063\u006d"&&len (_bfeg .Params )==6{if _becb ,_ecg :=_a .GetFloatVal (_bfeg .Params [0]);_ecg {_ggcf =_ggcf *_becb ;};if _gbga ,_gcge :=_a .GetFloatVal (_bfeg .Params [3]);
_gcge {_fbcf =_fbcf *_gbga ;};if _bee ,_fbe :=_a .GetIntVal (_bfeg .Params [0]);_fbe {_ggcf =_ggcf *float64 (_bee );};if _fgdd ,_fbdg :=_a .GetIntVal (_bfeg .Params [3]);_fbdg {_fbcf =_fbcf *float64 (_fgdd );};};if _bfeg .Operand =="\u0044\u006f"&&len (_bfeg .Params )==1{_ade ,_aaf :=_a .GetName (_bfeg .Params [0]);
if !_aaf {continue ;};if _gcdf ,_daa :=_efgf [string (*_ade )];_daa {_dcd ,_gaac :=_ggcf /72.0,_fbcf /72.0;_ebf ,_gbbg :=float64 (_gcdf .Width )/_dcd ,float64 (_gcdf .Height )/_gaac ;if _dcd ==0||_gaac ==0{_ebf =72.0;_gbbg =72.0;};_gcdf .PPI =_fc .Max (_gcdf .PPI ,_ebf );
_gcdf .PPI =_fc .Max (_gcdf .PPI ,_gbbg );};};};};};};for _ ,_fge :=range _bda {if _ ,_eff :=_cdg [_fge .Stream ];_eff {continue ;};if _fge .PPI <=_aea .ImageUpperPPI {continue ;};_cfcg ,_gcbb :=_d .NewXObjectImageFromStream (_fge .Stream );if _gcbb !=nil {return nil ,_gcbb ;
};var _bbe imageModifications ;_bbe .Scale =_aea .ImageUpperPPI /_fge .PPI ;if _fge .BitsPerComponent ==1&&_fge .ColorComponents ==1{_begc :=_fc .Round (_fge .PPI /_aea .ImageUpperPPI );_fadc :=_e .NextPowerOf2 (uint (_begc ));if _e .InDelta (float64 (_fadc ),1/_bbe .Scale ,0.3){_bbe .Scale =float64 (1)/float64 (_fadc );
};if _ ,_bdf :=_cfcg .Filter .(*_a .JBIG2Encoder );!_bdf {_bbe .Encoding =_a .NewJBIG2Encoder ();};};if _gcbb =_gdf (_cfcg ,_bbe );_gcbb !=nil {_ad .Log .Debug ("\u0045\u0072\u0072\u006f\u0072 \u0073\u0063\u0061\u006c\u0065\u0020\u0069\u006d\u0061\u0067\u0065\u0020\u006be\u0065\u0070\u0020\u006f\u0072\u0069\u0067\u0069\u006e\u0061\u006c\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_gcbb );
continue ;};_bbe .Encoding =nil ;if _gec ,_dedfa :=_a .GetStream (_fge .Stream .PdfObjectDictionary .Get (_a .PdfObjectName ("\u0053\u004d\u0061s\u006b")));_dedfa {_gaaf ,_cbgd :=_d .NewXObjectImageFromStream (_gec );if _cbgd !=nil {return nil ,_cbgd ;
};if _cbgd =_gdf (_gaaf ,_bbe );_cbgd !=nil {return nil ,_cbgd ;};};};return objects ,nil ;};func _agb (_gac []_a .PdfObject )(_cc map[*_a .PdfObjectStream ]struct{},_abca error ){_cc =map[*_a .PdfObjectStream ]struct{}{};_gacd :=map[*_d .PdfFont ]struct{}{};
_bce :=_bad (_gac );for _ ,_af :=range _bce ._fcafg {_gcg ,_bab :=_a .GetDict (_af .PdfObject );if !_bab {continue ;};_ggd ,_bab :=_a .GetDict (_gcg .Get ("\u0052e\u0073\u006f\u0075\u0072\u0063\u0065s"));if !_bab {continue ;};_gab ,_ :=_fggd (_gcg .Get ("\u0043\u006f\u006e\u0074\u0065\u006e\u0074\u0073"));
_bca ,_cd :=_d .NewPdfPageResourcesFromDict (_ggd );if _cd !=nil {return nil ,_cd ;};_agbf :=[]content {{_ggfc :_gab ,_ccg :_bca }};_eae :=_beg (_gcg .Get ("\u0041\u006e\u006e\u006f\u0074\u0073"));if _eae !=nil {_agbf =append (_agbf ,_eae ...);};for _ ,_eaa :=range _agbf {_db ,_bcag :=_faa .NewFromContents (_eaa ._ggfc ,_eaa ._ccg );
if _bcag !=nil {return nil ,_bcag ;};_df ,_ ,_ ,_bcag :=_db .ExtractPageText ();if _bcag !=nil {return nil ,_bcag ;};for _ ,_fe :=range _df .Marks ().Elements (){if _fe .Font ==nil {continue ;};if _ ,_aff :=_gacd [_fe .Font ];!_aff {_gacd [_fe .Font ]=struct{}{};
};};};};_cf :=map[*_a .PdfObjectStream ][]*_d .PdfFont {};for _bfd :=range _gacd {_feb :=_bfd .FontDescriptor ();if _feb ==nil ||_feb .FontFile2 ==nil {continue ;};_cgf ,_cad :=_a .GetStream (_feb .FontFile2 );if !_cad {continue ;};_cf [_cgf ]=append (_cf [_cgf ],_bfd );
};for _agg :=range _cf {var _dde []rune ;var _fffg []_eb .GlyphIndex ;for _ ,_cbde :=range _cf [_agg ]{switch _bb :=_cbde .Encoder ().(type ){case *_c .IdentityEncoder :_cea :=_bb .RegisteredRunes ();_ecc :=make ([]_eb .GlyphIndex ,len (_cea ));for _gfe ,_fd :=range _cea {_ecc [_gfe ]=_eb .GlyphIndex (_fd );
};_fffg =append (_fffg ,_ecc ...);case *_c .TrueTypeFontEncoder :_dfe :=_bb .RegisteredRunes ();_dde =append (_dde ,_dfe ...);case _c .SimpleEncoder :_gfg :=_bb .Charcodes ();for _ ,_gfc :=range _gfg {_ebd ,_bec :=_bb .CharcodeToRune (_gfc );if !_bec {_ad .Log .Debug ("\u0043\u0068a\u0072\u0063\u006f\u0064\u0065\u003c\u002d\u003e\u0072\u0075\u006e\u0065\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064: \u0025\u0064",_gfc );
continue ;};_dde =append (_dde ,_ebd );};};};_abca =_ggdg (_agg ,_dde ,_fffg );if _abca !=nil {_ad .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0073\u0075\u0062\u0073\u0065\u0074\u0074\u0069\u006eg\u0020f\u006f\u006e\u0074\u0020\u0073\u0074\u0072\u0065\u0061\u006d\u003a\u0020\u0025\u0076",_abca );
return nil ,_abca ;};_cc [_agg ]=struct{}{};};return _cc ,nil ;};func _g (_ffee *_ffe .ContentStreamOperations )*_ffe .ContentStreamOperations {if _ffee ==nil {return nil ;};_dg :=_ffe .ContentStreamOperations {};for _ ,_gga :=range *_ffee {switch _gga .Operand {case "\u0042\u0044\u0043","\u0042\u004d\u0043","\u0045\u004d\u0043":continue ;
case "\u0054\u006d":if len (_gga .Params )==6{if _ae ,_fff :=_a .GetNumbersAsFloat (_gga .Params );_fff ==nil {if _ae [0]==1&&_ae [1]==0&&_ae [2]==0&&_ae [3]==1{_gga =&_ffe .ContentStreamOperation {Params :[]_a .PdfObject {_gga .Params [4],_gga .Params [5]},Operand :"\u0054\u0064"};
};};};};_dg =append (_dg ,_gga );};return &_dg ;};

// ImagePPI optimizes images by scaling images such that the PPI (pixels per inch) is never higher than ImageUpperPPI.
// TODO(a5i): Add support for inline images.
// It implements interface model.Optimizer.
type ImagePPI struct{ImageUpperPPI float64 ;};type objectStructure struct{_cgbe *_a .PdfObjectDictionary ;_ffcf *_a .PdfObjectDictionary ;_fcafg []*_a .PdfIndirectObject ;};func _efae (_dge []_a .PdfObject ,_adcf map[_a .PdfObject ]_a .PdfObject ){if len (_adcf )==0{return ;
};for _daad ,_bfc :=range _dge {if _gbec ,_egge :=_adcf [_bfc ];_egge {_dge [_daad ]=_gbec ;continue ;};_adcf [_bfc ]=_bfc ;switch _bbab :=_bfc .(type ){case *_a .PdfObjectArray :_efgb :=make ([]_a .PdfObject ,_bbab .Len ());copy (_efgb ,_bbab .Elements ());
_efae (_efgb ,_adcf );for _cfab ,_ggfa :=range _efgb {_bbab .Set (_cfab ,_ggfa );};case *_a .PdfObjectStreams :_efae (_bbab .Elements (),_adcf );case *_a .PdfObjectStream :_cec :=[]_a .PdfObject {_bbab .PdfObjectDictionary };_efae (_cec ,_adcf );_bbab .PdfObjectDictionary =_cec [0].(*_a .PdfObjectDictionary );
case *_a .PdfObjectDictionary :_aaa :=_bbab .Keys ();_baff :=make ([]_a .PdfObject ,len (_aaa ));for _cdf ,_acae :=range _aaa {_baff [_cdf ]=_bbab .Get (_acae );};_efae (_baff ,_adcf );for _geff ,_adda :=range _aaa {_bbab .Set (_adda ,_baff [_geff ]);};
case *_a .PdfIndirectObject :_ccae :=[]_a .PdfObject {_bbab .PdfObject };_efae (_ccae ,_adcf );_bbab .PdfObject =_ccae [0];};};};func _fggd (_edg _a .PdfObject )(_gaea string ,_eeb []_a .PdfObject ){var _cfde _ff .Buffer ;switch _cede :=_edg .(type ){case *_a .PdfIndirectObject :_eeb =append (_eeb ,_cede );
_edg =_cede .PdfObject ;};switch _eggec :=_edg .(type ){case *_a .PdfObjectStream :if _bgee ,_cfdc :=_a .DecodeStream (_eggec );_cfdc ==nil {_cfde .Write (_bgee );_eeb =append (_eeb ,_eggec );};case *_a .PdfObjectArray :for _ ,_affd :=range _eggec .Elements (){switch _abcb :=_affd .(type ){case *_a .PdfObjectStream :if _bfac ,_edgf :=_a .DecodeStream (_abcb );
_edgf ==nil {_cfde .Write (_bfac );_eeb =append (_eeb ,_abcb );};};};};return _cfde .String (),_eeb ;};func _dea (_bdgg []_a .PdfObject )[]*imageInfo {_gaa :=_a .PdfObjectName ("\u0053u\u0062\u0074\u0079\u0070\u0065");_agbb :=make (map[*_a .PdfObjectStream ]struct{});
var _cfc error ;var _bbc []*imageInfo ;for _ ,_fcc :=range _bdgg {_cab ,_cgbc :=_a .GetStream (_fcc );if !_cgbc {continue ;};if _ ,_gcgc :=_agbb [_cab ];_gcgc {continue ;};_agbb [_cab ]=struct{}{};_gfgd :=_cab .PdfObjectDictionary .Get (_gaa );_dfg ,_cgbc :=_a .GetName (_gfgd );
if !_cgbc ||string (*_dfg )!="\u0049\u006d\u0061g\u0065"{continue ;};_fcda :=&imageInfo {BitsPerComponent :8,Stream :_cab };if _fcda .ColorSpace ,_cfc =_d .DetermineColorspaceNameFromPdfObject (_cab .PdfObjectDictionary .Get ("\u0043\u006f\u006c\u006f\u0072\u0053\u0070\u0061\u0063\u0065"));
_cfc !=nil {_ad .Log .Error ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064\u0065\u0074\u0065r\u006d\u0069\u006e\u0065\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065\u0020\u0025\u0073",_cfc );continue ;};if _cgd ,_addg :=_a .GetIntVal (_cab .PdfObjectDictionary .Get ("\u0042\u0069t\u0073\u0050\u0065r\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074"));
_addg {_fcda .BitsPerComponent =_cgd ;};if _aeb ,_facc :=_a .GetIntVal (_cab .PdfObjectDictionary .Get ("\u0057\u0069\u0064t\u0068"));_facc {_fcda .Width =_aeb ;};if _cga ,_ddbb :=_a .GetIntVal (_cab .PdfObjectDictionary .Get ("\u0048\u0065\u0069\u0067\u0068\u0074"));
_ddbb {_fcda .Height =_cga ;};switch _fcda .ColorSpace {case "\u0044e\u0076\u0069\u0063\u0065\u0052\u0047B":_fcda .ColorComponents =3;case "\u0044\u0065\u0076\u0069\u0063\u0065\u0047\u0072\u0061\u0079":_fcda .ColorComponents =1;default:_ad .Log .Warning ("\u004f\u0070\u0074\u0069\u006d\u0069\u007a\u0061t\u0069\u006f\u006e i\u0073\u0020\u006e\u006f\u0074\u0020s\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0066\u006f\u0072\u0020\u0063\u006fl\u006f\u0072\u0020\u0073\u0070\u0061\u0063\u0065 \u0025\u0073",_fcda .ColorSpace );
continue ;};_bbc =append (_bbc ,_fcda );};return _bbc ;};