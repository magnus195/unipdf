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

package testutils ;import (_d "crypto/md5";_aa "encoding/hex";_gg "errors";_a "fmt";_fc "github.com/unidoc/unipdf/v3/common";_cc "github.com/unidoc/unipdf/v3/core";_db "image";_bg "image/png";_c "io";_b "os";_ff "os/exec";_de "path/filepath";_f "strings";
_e "testing";);func ComparePNGFiles (file1 ,file2 string )(bool ,error ){_ae ,_cg :=HashFile (file1 );if _cg !=nil {return false ,_cg ;};_ge ,_cg :=HashFile (file2 );if _cg !=nil {return false ,_cg ;};if _ae ==_ge {return true ,nil ;};_bbg ,_cg :=ReadPNG (file1 );
if _cg !=nil {return false ,_cg ;};_bgc ,_cg :=ReadPNG (file2 );if _cg !=nil {return false ,_cg ;};if _bbg .Bounds ()!=_bgc .Bounds (){return false ,nil ;};return CompareImages (_bbg ,_bgc );};func RenderPDFToPNGs (pdfPath string ,dpi int ,outpathTpl string )error {if dpi <=0{dpi =100;
};if _ ,_bc :=_ff .LookPath ("\u0067\u0073");_bc !=nil {return ErrRenderNotSupported ;};return _ff .Command ("\u0067\u0073","\u002d\u0073\u0044\u0045\u0056\u0049\u0043\u0045\u003d\u0070\u006e\u0067a\u006c\u0070\u0068\u0061","\u002d\u006f",outpathTpl ,_a .Sprintf ("\u002d\u0072\u0025\u0064",dpi ),pdfPath ).Run ();
};func CopyFile (src ,dst string )error {_ad ,_fg :=_b .Open (src );if _fg !=nil {return _fg ;};defer _ad .Close ();_bgg ,_fg :=_b .Create (dst );if _fg !=nil {return _fg ;};defer _bgg .Close ();_ ,_fg =_c .Copy (_bgg ,_ad );return _fg ;};var (ErrRenderNotSupported =_gg .New ("\u0072\u0065\u006e\u0064\u0065r\u0069\u006e\u0067\u0020\u0050\u0044\u0046\u0020\u0066\u0069\u006c\u0065\u0073 \u0069\u0073\u0020\u006e\u006f\u0074\u0020\u0073\u0075\u0070\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u006e\u0020\u0074\u0068\u0069\u0073\u0020\u0073\u0079\u0073\u0074\u0065m");
);func ParseIndirectObjects (rawpdf string )(map[int64 ]_cc .PdfObject ,error ){_aaf :=_cc .NewParserFromString (rawpdf );_bdc :=map[int64 ]_cc .PdfObject {};for {_bdd ,_ed :=_aaf .ParseIndirectObject ();if _ed !=nil {if _ed ==_c .EOF {break ;};return nil ,_ed ;
};switch _cb :=_bdd .(type ){case *_cc .PdfIndirectObject :_bdc [_cb .ObjectNumber ]=_bdd ;case *_cc .PdfObjectStream :_bdc [_cb .ObjectNumber ]=_bdd ;};};for _ ,_aef :=range _bdc {_ac (_aef ,_bdc );};return _bdc ,nil ;};func ReadPNG (file string )(_db .Image ,error ){_ag ,_gc :=_b .Open (file );
if _gc !=nil {return nil ,_gc ;};defer _ag .Close ();return _bg .Decode (_ag );};func CompareImages (img1 ,img2 _db .Image )(bool ,error ){_fba :=img1 .Bounds ();_fa :=0;for _ea :=0;_ea < _fba .Size ().X ;_ea ++{for _bggfc :=0;_bggfc < _fba .Size ().Y ;
_bggfc ++{_cd ,_gb ,_agb ,_ :=img1 .At (_ea ,_bggfc ).RGBA ();_bb ,_dea ,_ec ,_ :=img2 .At (_ea ,_bggfc ).RGBA ();if _cd !=_bb ||_gb !=_dea ||_agb !=_ec {_fa ++;};};};_gf :=float64 (_fa )/float64 (_fba .Dx ()*_fba .Dy ());if _gf > 0.0001{_a .Printf ("\u0064\u0069\u0066f \u0066\u0072\u0061\u0063\u0074\u0069\u006f\u006e\u003a\u0020\u0025\u0076\u0020\u0028\u0025\u0064\u0029\u000a",_gf ,_fa );
return false ,nil ;};return true ,nil ;};func HashFile (file string )(string ,error ){_gd ,_bggf :=_b .Open (file );if _bggf !=nil {return "",_bggf ;};defer _gd .Close ();_fb :=_d .New ();if _ ,_bggf =_c .Copy (_fb ,_gd );_bggf !=nil {return "",_bggf ;
};return _aa .EncodeToString (_fb .Sum (nil )),nil ;};func _ac (_agf _cc .PdfObject ,_ffe map[int64 ]_cc .PdfObject )error {switch _af :=_agf .(type ){case *_cc .PdfIndirectObject :_gba :=_af ;_ac (_gba .PdfObject ,_ffe );case *_cc .PdfObjectDictionary :_ca :=_af ;
for _ ,_cbf :=range _ca .Keys (){_dc :=_ca .Get (_cbf );if _eag ,_ggb :=_dc .(*_cc .PdfObjectReference );_ggb {_fbc ,_be :=_ffe [_eag .ObjectNumber ];if !_be {return _gg .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_ca .Set (_cbf ,_fbc );}else {_ac (_dc ,_ffe );};};case *_cc .PdfObjectArray :_cfd :=_af ;for _da ,_ce :=range _cfd .Elements (){if _cbg ,_dg :=_ce .(*_cc .PdfObjectReference );_dg {_cdf ,_dag :=_ffe [_cbg .ObjectNumber ];if !_dag {return _gg .New ("r\u0065\u0066\u0065\u0072\u0065\u006ec\u0065\u0020\u0074\u006f\u0020\u006f\u0075\u0074\u0073i\u0064\u0065\u0020o\u0062j\u0065\u0063\u0074");
};_cfd .Set (_da ,_cdf );}else {_ac (_ce ,_ffe );};};};return nil ;};func RunRenderTest (t *_e .T ,pdfPath ,outputDir ,baselineRenderPath string ,saveBaseline bool ){_ffd :=_f .TrimSuffix (_de .Base (pdfPath ),_de .Ext (pdfPath ));t .Run ("\u0072\u0065\u006e\u0064\u0065\u0072",func (_baf *_e .T ){_ggg :=_de .Join (outputDir ,_ffd );
_bd :=_ggg +"\u002d%\u0064\u002e\u0070\u006e\u0067";if _eaa :=RenderPDFToPNGs (pdfPath ,0,_bd );_eaa !=nil {_baf .Skip (_eaa );};for _eaac :=1;true ;_eaac ++{_eb :=_a .Sprintf ("\u0025s\u002d\u0025\u0064\u002e\u0070\u006eg",_ggg ,_eaac );_df :=_de .Join (baselineRenderPath ,_a .Sprintf ("\u0025\u0073\u002d\u0025\u0064\u005f\u0065\u0078\u0070\u002e\u0070\u006e\u0067",_ffd ,_eaac ));
if _ ,_cga :=_b .Stat (_eb );_cga !=nil {break ;};_baf .Logf ("\u0025\u0073",_df );if saveBaseline {_baf .Logf ("\u0043\u006fp\u0079\u0069\u006eg\u0020\u0025\u0073\u0020\u002d\u003e\u0020\u0025\u0073",_eb ,_df );_gcf :=CopyFile (_eb ,_df );if _gcf !=nil {_baf .Fatalf ("\u0045\u0052\u0052OR\u0020\u0063\u006f\u0070\u0079\u0069\u006e\u0067\u0020\u0074\u006f\u0020\u0025\u0073\u003a\u0020\u0025\u0076",_df ,_gcf );
};continue ;};_baf .Run (_a .Sprintf ("\u0070\u0061\u0067\u0065\u0025\u0064",_eaac ),func (_ebc *_e .T ){_ebc .Logf ("\u0043o\u006dp\u0061\u0072\u0069\u006e\u0067 \u0025\u0073 \u0076\u0073\u0020\u0025\u0073",_eb ,_df );_bcb ,_ee :=ComparePNGFiles (_eb ,_df );
if _b .IsNotExist (_ee ){_ebc .Fatal ("\u0069m\u0061g\u0065\u0020\u0066\u0069\u006ce\u0020\u006di\u0073\u0073\u0069\u006e\u0067");}else if !_bcb {_ebc .Fatal ("\u0077\u0072\u006f\u006eg \u0070\u0061\u0067\u0065\u0020\u0072\u0065\u006e\u0064\u0065\u0072\u0065\u0064");
};});};});};func CompareDictionariesDeep (d1 ,d2 *_cc .PdfObjectDictionary )bool {if len (d1 .Keys ())!=len (d2 .Keys ()){_fc .Log .Debug ("\u0044\u0069\u0063\u0074\u0020\u0065\u006e\u0074\u0072\u0069\u0065\u0073\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",len (d1 .Keys ()),len (d2 .Keys ()));
_fc .Log .Debug ("\u0057\u0061s\u0020\u0027\u0025s\u0027\u0020\u0076\u0073\u0020\u0027\u0025\u0073\u0027",d1 .WriteString (),d2 .WriteString ());return false ;};for _ ,_cgac :=range d1 .Keys (){if _cgac =="\u0050\u0061\u0072\u0065\u006e\u0074"{continue ;
};_cde :=_cc .TraceToDirectObject (d1 .Get (_cgac ));_eeb :=_cc .TraceToDirectObject (d2 .Get (_cgac ));if _cde ==nil {_fc .Log .Debug ("\u00761\u0020\u0069\u0073\u0020\u006e\u0069l");return false ;};if _eeb ==nil {_fc .Log .Debug ("\u00762\u0020\u0069\u0073\u0020\u006e\u0069l");
return false ;};switch _add :=_cde .(type ){case *_cc .PdfObjectDictionary :_dga ,_eec :=_eeb .(*_cc .PdfObjectDictionary );if !_eec {_fc .Log .Debug ("\u0054\u0079\u0070\u0065 m\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020\u0025\u0054\u0020\u0076\u0073\u0020%\u0054",_cde ,_eeb );
return false ;};if !CompareDictionariesDeep (_add ,_dga ){return false ;};continue ;case *_cc .PdfObjectArray :_dad ,_fe :=_eeb .(*_cc .PdfObjectArray );if !_fe {_fc .Log .Debug ("\u00762\u0020n\u006f\u0074\u0020\u0061\u006e\u0020\u0061\u0072\u0072\u0061\u0079");
return false ;};if _add .Len ()!=_dad .Len (){_fc .Log .Debug ("\u0061\u0072\u0072\u0061\u0079\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006d\u0069s\u006da\u0074\u0063\u0068\u0020\u0028\u0025\u0064\u0020\u0021\u003d\u0020\u0025\u0064\u0029",_add .Len (),_dad .Len ());
return false ;};for _aea :=0;_aea < _add .Len ();_aea ++{_cab :=_cc .TraceToDirectObject (_add .Get (_aea ));_ga :=_cc .TraceToDirectObject (_dad .Get (_aea ));if _adf ,_gac :=_cab .(*_cc .PdfObjectDictionary );_gac {_acd ,_fac :=_ga .(*_cc .PdfObjectDictionary );
if !_fac {return false ;};if !CompareDictionariesDeep (_adf ,_acd ){return false ;};}else {if _cab .WriteString ()!=_ga .WriteString (){_fc .Log .Debug ("M\u0069\u0073\u006d\u0061tc\u0068 \u0027\u0025\u0073\u0027\u0020!\u003d\u0020\u0027\u0025\u0073\u0027",_cab .WriteString (),_ga .WriteString ());
return false ;};};};continue ;};if _cde .String ()!=_eeb .String (){_fc .Log .Debug ("\u006b\u0065y\u003d\u0025\u0073\u0020\u004d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0021\u0020\u0027\u0025\u0073\u0027\u0020\u0021\u003d\u0020'%\u0073\u0027",_cgac ,_cde .String (),_eeb .String ());
_fc .Log .Debug ("\u0046o\u0072 \u0027\u0025\u0054\u0027\u0020\u002d\u0020\u0027\u0025\u0054\u0027",_cde ,_eeb );_fc .Log .Debug ("\u0046\u006f\u0072\u0020\u0027\u0025\u002b\u0076\u0027\u0020\u002d\u0020'\u0025\u002b\u0076\u0027",_cde ,_eeb );return false ;
};};return true ;};