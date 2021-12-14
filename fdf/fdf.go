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

// Package fdf provides support for loading form field data from Form Field Data (FDF) files.
package fdf ;import (_gf "bufio";_fe "bytes";_cdd "encoding/hex";_g "errors";_ce "fmt";_ab "github.com/unidoc/unipdf/v3/common";_cb "github.com/unidoc/unipdf/v3/core";_a "io";_e "os";_b "regexp";_ca "sort";_cd "strconv";_c "strings";);var _cf =_b .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");
func (_cdaa *fdfParser )readComment ()(string ,error ){var _dac _fe .Buffer ;_ ,_bfc :=_cdaa .skipSpaces ();if _bfc !=nil {return _dac .String (),_bfc ;};_fbf :=true ;for {_cga ,_bg :=_cdaa ._fgc .Peek (1);if _bg !=nil {_ab .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_bg .Error ());
return _dac .String (),_bg ;};if _fbf &&_cga [0]!='%'{return _dac .String (),_g .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");};_fbf =false ;if (_cga [0]!='\r')&&(_cga [0]!='\n'){_dba ,_ :=_cdaa ._fgc .ReadByte ();
_dac .WriteByte (_dba );}else {break ;};};return _dac .String (),nil ;};func (_gc *fdfParser )parseBool ()(_cb .PdfObjectBool ,error ){_bef ,_egd :=_gc ._fgc .Peek (4);if _egd !=nil {return _cb .PdfObjectBool (false ),_egd ;};if (len (_bef )>=4)&&(string (_bef [:4])=="\u0074\u0072\u0075\u0065"){_gc ._fgc .Discard (4);
return _cb .PdfObjectBool (true ),nil ;};_bef ,_egd =_gc ._fgc .Peek (5);if _egd !=nil {return _cb .PdfObjectBool (false ),_egd ;};if (len (_bef )>=5)&&(string (_bef [:5])=="\u0066\u0061\u006cs\u0065"){_gc ._fgc .Discard (5);return _cb .PdfObjectBool (false ),nil ;
};return _cb .PdfObjectBool (false ),_g .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};func (_bcd *fdfParser )parseArray ()(*_cb .PdfObjectArray ,error ){_fbfb :=_cb .MakeArray ();
_bcd ._fgc .ReadByte ();for {_bcd .skipSpaces ();_geg ,_daf :=_bcd ._fgc .Peek (1);if _daf !=nil {return _fbfb ,_daf ;};if _geg [0]==']'{_bcd ._fgc .ReadByte ();break ;};_ebc ,_daf :=_bcd .parseObject ();if _daf !=nil {return _fbfb ,_daf ;};_fbfb .Append (_ebc );
};return _fbfb ,nil ;};

// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_cb .PdfObjectDictionary ,error ){_ga :=map[string ]*_cb .PdfObjectDictionary {};for _cgb :=0;_cgb < fdf ._cc .Len ();_cgb ++{_cda ,_df :=_cb .GetDict (fdf ._cc .Get (_cgb ));if _df {_cdb ,_ :=_cb .GetString (_cda .Get ("\u0054"));
if _cdb !=nil {_ga [_cdb .Str ()]=_cda ;};};};return _ga ,nil ;};var _fdbe =_b .MustCompile ("\u0025\u0025\u0045O\u0046");func (_beg *fdfParser )parseHexString ()(*_cb .PdfObjectString ,error ){_beg ._fgc .ReadByte ();var _dfb _fe .Buffer ;for {_ccg ,_aaf :=_beg ._fgc .Peek (1);
if _aaf !=nil {return _cb .MakeHexString (""),_aaf ;};if _ccg [0]=='>'{_beg ._fgc .ReadByte ();break ;};_cfd ,_ :=_beg ._fgc .ReadByte ();if !_cb .IsWhiteSpace (_cfd ){_dfb .WriteByte (_cfd );};};if _dfb .Len ()%2==1{_dfb .WriteRune ('0');};_gd ,_cccc :=_cdd .DecodeString (_dfb .String ());
if _cccc !=nil {_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_dfb .String ());
return _cb .MakeHexString (""),nil ;};return _cb .MakeHexString (string (_gd )),nil ;};func (_dfa *fdfParser )parseDict ()(*_cb .PdfObjectDictionary ,error ){_ab .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");
_aee :=_cb .MakeDict ();_ddf ,_ :=_dfa ._fgc .ReadByte ();if _ddf !='<'{return nil ,_g .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};_ddf ,_ =_dfa ._fgc .ReadByte ();if _ddf !='<'{return nil ,_g .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");
};for {_dfa .skipSpaces ();_dfa .skipComments ();_fge ,_dce :=_dfa ._fgc .Peek (2);if _dce !=nil {return nil ,_dce ;};_ab .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_fge ),string (_fge ));
if (_fge [0]=='>')&&(_fge [1]=='>'){_ab .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");_dfa ._fgc .ReadByte ();_dfa ._fgc .ReadByte ();break ;};_ab .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");
_cba ,_dce :=_dfa .parseName ();_ab .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_cba );if _dce !=nil {_ab .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_dce );
return nil ,_dce ;};if len (_cba )> 4&&_cba [len (_cba )-4:]=="\u006e\u0075\u006c\u006c"{_dfbf :=_cba [0:len (_cba )-4];_ab .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_cba );
_ab .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_dfbf );_dfa .skipSpaces ();_fee ,_ :=_dfa ._fgc .Peek (1);if _fee [0]=='/'{_aee .Set (_dfbf ,_cb .MakeNull ());continue ;};};_dfa .skipSpaces ();
_fad ,_dce :=_dfa .parseObject ();if _dce !=nil {return nil ,_dce ;};_aee .Set (_cba ,_fad );_ab .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_cba ,_fad .String ());};_ab .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");
return _aee ,nil ;};func _edfc (_acc string )(_cb .PdfObjectReference ,error ){_cbg :=_cb .PdfObjectReference {};_gg :=_cf .FindStringSubmatch (_acc );if len (_gg )< 3{_ab .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");
return _cbg ,_g .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");};_afa ,_dbg :=_cd .Atoi (_gg [1]);if _dbg !=nil {_ab .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_gg [1]);
return _cbg ,nil ;};_cbg .ObjectNumber =int64 (_afa );_aad ,_dbg :=_cd .Atoi (_gg [2]);if _dbg !=nil {_ab .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_gg [2]);
return _cbg ,nil ;};_cbg .GenerationNumber =int64 (_aad );return _cbg ,nil ;};func (_dgg *fdfParser )getFileOffset ()int64 {_bfb ,_ :=_dgg ._cge .Seek (0,_a .SeekCurrent );_bfb -=int64 (_dgg ._fgc .Buffered ());return _bfb ;};type fdfParser struct{_gea int ;
_efa int ;_af map[int64 ]_cb .PdfObject ;_cge _a .ReadSeeker ;_fgc *_gf .Reader ;_ae int64 ;_eae *_cb .PdfObjectDictionary ;};

// Root returns the Root of the FDF document.
func (_cec *fdfParser )Root ()(*_cb .PdfObjectDictionary ,error ){if _cec ._eae !=nil {if _eag ,_beac :=_cec .trace (_cec ._eae .Get ("\u0052\u006f\u006f\u0074")).(*_cb .PdfObjectDictionary );_beac {if _adb ,_bacc :=_cec .trace (_eag .Get ("\u0046\u0044\u0046")).(*_cb .PdfObjectDictionary );
_bacc {return _adb ,nil ;};};};var _fga []int64 ;for _cgc :=range _cec ._af {_fga =append (_fga ,_cgc );};_ca .Slice (_fga ,func (_baa ,_cddc int )bool {return _fga [_baa ]< _fga [_cddc ]});for _ ,_cgee :=range _fga {_dbe :=_cec ._af [_cgee ];if _ebf ,_bfdf :=_cec .trace (_dbe ).(*_cb .PdfObjectDictionary );
_bfdf {if _cead ,_fcdb :=_cec .trace (_ebf .Get ("\u0046\u0044\u0046")).(*_cb .PdfObjectDictionary );_fcdb {return _cead ,nil ;};};};return nil ,_g .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};func (_bega *fdfParser )parseFdfVersion ()(int ,int ,error ){_bega ._cge .Seek (0,_a .SeekStart );
_dbf :=20;_abe :=make ([]byte ,_dbf );_bega ._cge .Read (_abe );_eed :=_bbd .FindStringSubmatch (string (_abe ));if len (_eed )< 3{_ff ,_gebd ,_ace :=_bega .seekFdfVersionTopDown ();if _ace !=nil {_ab .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");
return 0,0,_ace ;};return _ff ,_gebd ,nil ;};_gdf ,_fdeg :=_cd .Atoi (_eed [1]);if _fdeg !=nil {return 0,0,_fdeg ;};_bafd ,_fdeg :=_cd .Atoi (_eed [2]);if _fdeg !=nil {return 0,0,_fdeg ;};_ab .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_gdf ,_bafd );
return _gdf ,_bafd ,nil ;};func (_bcb *fdfParser )seekToEOFMarker (_debb int64 )error {_ede :=int64 (0);_cgbf :=int64 (1000);for _ede < _debb {if _debb <=(_cgbf +_ede ){_cgbf =_debb -_ede ;};_ ,_cea :=_bcb ._cge .Seek (-_ede -_cgbf ,_a .SeekEnd );if _cea !=nil {return _cea ;
};_gcg :=make ([]byte ,_cgbf );_bcb ._cge .Read (_gcg );_ab .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_gcg ));_egf :=_fdbe .FindAllStringIndex (string (_gcg ),-1);
if _egf !=nil {_fac :=_egf [len (_egf )-1];_ab .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_egf );_bcb ._cge .Seek (-_ede -_cgbf +int64 (_fac [0]),_a .SeekEnd );return nil ;};_ab .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");
_ede +=_cgbf ;};_ab .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _g .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
};func (_bc *fdfParser )skipSpaces ()(int ,error ){_de :=0;for {_dcb ,_efad :=_bc ._fgc .ReadByte ();if _efad !=nil {return 0,_efad ;};if _cb .IsWhiteSpace (_dcb ){_de ++;}else {_bc ._fgc .UnreadByte ();break ;};};return _de ,nil ;};func (_deb *fdfParser )parseName ()(_cb .PdfObjectName ,error ){var _gfc _fe .Buffer ;
_ccc :=false ;for {_efc ,_edd :=_deb ._fgc .Peek (1);if _edd ==_a .EOF {break ;};if _edd !=nil {return _cb .PdfObjectName (_gfc .String ()),_edd ;};if !_ccc {if _efc [0]=='/'{_ccc =true ;_deb ._fgc .ReadByte ();}else if _efc [0]=='%'{_deb .readComment ();
_deb .skipSpaces ();}else {_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_efc ,_efc );return _cb .PdfObjectName (_gfc .String ()),_ce .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_efc [0]);
};}else {if _cb .IsWhiteSpace (_efc [0]){break ;}else if (_efc [0]=='/')||(_efc [0]=='[')||(_efc [0]=='(')||(_efc [0]==']')||(_efc [0]=='<')||(_efc [0]=='>'){break ;}else if _efc [0]=='#'{_deg ,_gb :=_deb ._fgc .Peek (3);if _gb !=nil {return _cb .PdfObjectName (_gfc .String ()),_gb ;
};_deb ._fgc .Discard (3);_feb ,_gb :=_cdd .DecodeString (string (_deg [1:3]));if _gb !=nil {return _cb .PdfObjectName (_gfc .String ()),_gb ;};_gfc .Write (_feb );}else {_aae ,_ :=_deb ._fgc .ReadByte ();_gfc .WriteByte (_aae );};};};return _cb .PdfObjectName (_gfc .String ()),nil ;
};func _agc (_dbb _a .ReadSeeker )(*fdfParser ,error ){_gac :=&fdfParser {};_gac ._cge =_dbb ;_gac ._af =map[int64 ]_cb .PdfObject {};_fdc ,_gbbg ,_fcf :=_gac .parseFdfVersion ();if _fcf !=nil {_ab .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_fcf );
return nil ,_fcf ;};_gac ._gea =_fdc ;_gac ._efa =_gbbg ;_fcf =_gac .parse ();return _gac ,_fcf ;};func (_bad *fdfParser )parseObject ()(_cb .PdfObject ,error ){_ab .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");
_bad .skipSpaces ();for {_fgd ,_gfd :=_bad ._fgc .Peek (2);if _gfd !=nil {return nil ,_gfd ;};_ab .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_fgd ));if _fgd [0]=='/'{_fdf ,_acde :=_bad .parseName ();
_ab .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_fdf );return &_fdf ,_acde ;}else if _fgd [0]=='('{_ab .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _bad .parseString ();}else if _fgd [0]=='['{_ab .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");
return _bad .parseArray ();}else if (_fgd [0]=='<')&&(_fgd [1]=='<'){_ab .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _bad .parseDict ();}else if _fgd [0]=='<'{_ab .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");
return _bad .parseHexString ();}else if _fgd [0]=='%'{_bad .readComment ();_bad .skipSpaces ();}else {_ab .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_fgd ,_ =_bad ._fgc .Peek (15);_feg :=string (_fgd );
_ab .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_feg );if (len (_feg )> 3)&&(_feg [:4]=="\u006e\u0075\u006c\u006c"){_dca ,_bfa :=_bad .parseNull ();return &_dca ,_bfa ;}else if (len (_feg )> 4)&&(_feg [:5]=="\u0066\u0061\u006cs\u0065"){_efd ,_acb :=_bad .parseBool ();
return &_efd ,_acb ;}else if (len (_feg )> 3)&&(_feg [:4]=="\u0074\u0072\u0075\u0065"){_cfe ,_abd :=_bad .parseBool ();return &_cfe ,_abd ;};_dde :=_cf .FindStringSubmatch (_feg );if len (_dde )> 1{_fgd ,_ =_bad ._fgc .ReadBytes ('R');_ab .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_fgd [:]));
_baf ,_acf :=_edfc (string (_fgd ));return &_baf ,_acf ;};_ddb :=_ed .FindStringSubmatch (_feg );if len (_ddb )> 1{_ab .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _bad .parseNumber ();};_ddb =_ge .FindStringSubmatch (_feg );
if len (_ddb )> 1{_ab .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");_ab .Log .Trace ("\u0025\u0020\u0073",_ddb );return _bad .parseNumber ();};_ab .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_feg );
return nil ,_g .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");
};};};func (_be *fdfParser )readAtLeast (_da []byte ,_dg int )(int ,error ){_bb :=_dg ;_dd :=0;_cgd :=0;for _bb > 0{_ec ,_ea :=_be ._fgc .Read (_da [_dd :]);if _ea !=nil {_ab .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_ec ,_cgd ,_ea .Error ());
return _dd ,_g .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");};_cgd ++;_dd +=_ec ;_bb -=_ec ;};return _dd ,nil ;};var _bbd =_b .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");
func (_gga *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_gga ._cge .Seek (0,_a .SeekStart );_gga ._fgc =_gf .NewReader (_gga ._cge );_efe :=20;_gbba :=make ([]byte ,_efe );for {_cccf ,_ead :=_gga ._fgc .ReadByte ();if _ead !=nil {if _ead ==_a .EOF {break ;
}else {return 0,0,_ead ;};};if _cb .IsDecimalDigit (_cccf )&&_gbba [_efe -1]=='.'&&_cb .IsDecimalDigit (_gbba [_efe -2])&&_gbba [_efe -3]=='-'&&_gbba [_efe -4]=='F'&&_gbba [_efe -5]=='D'&&_gbba [_efe -6]=='P'{_acac :=int (_gbba [_efe -2]-'0');_aff :=int (_cccf -'0');
return _acac ,_aff ,nil ;};_gbba =append (_gbba [1:_efe ],_cccf );};return 0,0,_g .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};var _bea =_b .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");
func (_aab *fdfParser )parseNumber ()(_cb .PdfObject ,error ){return _cb .ParseNumber (_aab ._fgc )};func (_fag *fdfParser )parseNull ()(_cb .PdfObjectNull ,error ){_ ,_abg :=_fag ._fgc .Discard (4);return _cb .PdfObjectNull {},_abg ;};func (_ccd *fdfParser )skipComments ()error {if _ ,_dfg :=_ccd .skipSpaces ();
_dfg !=nil {return _dfg ;};_fc :=true ;for {_aa ,_eef :=_ccd ._fgc .Peek (1);if _eef !=nil {_ab .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_eef .Error ());return _eef ;};if _fc &&_aa [0]!='%'{return nil ;};_fc =false ;if (_aa [0]!='\r')&&(_aa [0]!='\n'){_ccd ._fgc .ReadByte ();
}else {break ;};};return _ccd .skipComments ();};

// Load loads FDF form data from `r`.
func Load (r _a .ReadSeeker )(*Data ,error ){_fg ,_fb :=_agc (r );if _fb !=nil {return nil ,_fb ;};_cdde ,_fb :=_fg .Root ();if _fb !=nil {return nil ,_fb ;};_eb ,_d :=_cb .GetArray (_cdde .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_d {return nil ,_g .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");
};return &Data {_cc :_eb ,_fd :_cdde },nil ;};func (_ccb *fdfParser )parse ()error {_ccb ._cge .Seek (0,_a .SeekStart );_ccb ._fgc =_gf .NewReader (_ccb ._cge );for {_ccb .skipComments ();_dbc ,_aef :=_ccb ._fgc .Peek (20);if _aef !=nil {_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return _aef ;};if _c .HasPrefix (string (_dbc ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_ccb ._fgc .Discard (7);_ccb .skipSpaces ();_ccb .skipComments ();_aca ,_ :=_ccb .parseDict ();_ccb ._eae =_aca ;break ;};_gfbe :=_bea .FindStringSubmatchIndex (string (_dbc ));
if len (_gfbe )< 6{_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_dbc ));
return _g .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_gafa ,_aef :=_ccb .parseIndirectObject ();if _aef !=nil {return _aef ;};switch _fea :=_gafa .(type ){case *_cb .PdfIndirectObject :_ccb ._af [_fea .ObjectNumber ]=_fea ;case *_cb .PdfObjectStream :_ccb ._af [_fea .ObjectNumber ]=_fea ;default:return _g .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};};return nil ;};var _ge =_b .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");func (_fda *fdfParser )setFileOffset (_ba int64 ){_fda ._cge .Seek (_ba ,_a .SeekStart );
_fda ._fgc =_gf .NewReader (_fda ._cge );};func (_fcd *fdfParser )parseString ()(*_cb .PdfObjectString ,error ){_fcd ._fgc .ReadByte ();var _gbb _fe .Buffer ;_bac :=1;for {_ad ,_fdg :=_fcd ._fgc .Peek (1);if _fdg !=nil {return _cb .MakeString (_gbb .String ()),_fdg ;
};if _ad [0]=='\\'{_fcd ._fgc .ReadByte ();_abb ,_efb :=_fcd ._fgc .ReadByte ();if _efb !=nil {return _cb .MakeString (_gbb .String ()),_efb ;};if _cb .IsOctalDigit (_abb ){_eca ,_abba :=_fcd ._fgc .Peek (2);if _abba !=nil {return _cb .MakeString (_gbb .String ()),_abba ;
};var _eefe []byte ;_eefe =append (_eefe ,_abb );for _ ,_bbb :=range _eca {if _cb .IsOctalDigit (_bbb ){_eefe =append (_eefe ,_bbb );}else {break ;};};_fcd ._fgc .Discard (len (_eefe )-1);_ab .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_eefe );
_edf ,_abba :=_cd .ParseUint (string (_eefe ),8,32);if _abba !=nil {return _cb .MakeString (_gbb .String ()),_abba ;};_gbb .WriteByte (byte (_edf ));continue ;};switch _abb {case 'n':_gbb .WriteRune ('\n');case 'r':_gbb .WriteRune ('\r');case 't':_gbb .WriteRune ('\t');
case 'b':_gbb .WriteRune ('\b');case 'f':_gbb .WriteRune ('\f');case '(':_gbb .WriteRune ('(');case ')':_gbb .WriteRune (')');case '\\':_gbb .WriteRune ('\\');};continue ;}else if _ad [0]=='('{_bac ++;}else if _ad [0]==')'{_bac --;if _bac ==0{_fcd ._fgc .ReadByte ();
break ;};};_gbbc ,_ :=_fcd ._fgc .ReadByte ();_gbb .WriteByte (_gbbc );};return _cb .MakeString (_gbb .String ()),nil ;};func (_fgcc *fdfParser )readTextLine ()(string ,error ){var _geb _fe .Buffer ;for {_dad ,_ece :=_fgcc ._fgc .Peek (1);if _ece !=nil {_ab .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_ece .Error ());
return _geb .String (),_ece ;};if (_dad [0]!='\r')&&(_dad [0]!='\n'){_gaf ,_ :=_fgcc ._fgc .ReadByte ();_geb .WriteByte (_gaf );}else {break ;};};return _geb .String (),nil ;};func (_fae *fdfParser )parseIndirectObject ()(_cb .PdfObject ,error ){_bba :=_cb .PdfIndirectObject {};
_ab .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_cfc ,_def :=_fae ._fgc .Peek (20);if _def !=nil {_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return &_bba ,_def ;};_ab .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_cfc ));_cbae :=_bea .FindStringSubmatchIndex (string (_cfc ));if len (_cbae )< 6{_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_cfc ));
return &_bba ,_g .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_fae ._fgc .Discard (_cbae [0]);_ab .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_cbae );_abdg :=_cbae [1]-_cbae [0];_dbd :=make ([]byte ,_abdg );_ ,_def =_fae .readAtLeast (_dbd ,_abdg );if _def !=nil {_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_def );
return nil ,_def ;};_ab .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_dbd );_dcf :=_bea .FindStringSubmatch (string (_dbd ));if len (_dcf )< 3{_ab .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_dbd ));
return &_bba ,_g .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_aed ,_ :=_cd .Atoi (_dcf [1]);_ggb ,_ :=_cd .Atoi (_dcf [2]);_bba .ObjectNumber =int64 (_aed );_bba .GenerationNumber =int64 (_ggb );for {_beb ,_dedd :=_fae ._fgc .Peek (2);if _dedd !=nil {return &_bba ,_dedd ;};_ab .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_beb ),string (_beb ));
if _cb .IsWhiteSpace (_beb [0]){_fae .skipSpaces ();}else if _beb [0]=='%'{_fae .skipComments ();}else if (_beb [0]=='<')&&(_beb [1]=='<'){_ab .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_bba .PdfObject ,_dedd =_fae .parseDict ();
_ab .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_dedd );if _dedd !=nil {return &_bba ,_dedd ;};_ab .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");
}else if (_beb [0]=='/')||(_beb [0]=='(')||(_beb [0]=='[')||(_beb [0]=='<'){_bba .PdfObject ,_dedd =_fae .parseObject ();if _dedd !=nil {return &_bba ,_dedd ;};_ab .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");
}else {if _beb [0]=='e'{_bfbg ,_egc :=_fae .readTextLine ();if _egc !=nil {return nil ,_egc ;};if len (_bfbg )>=6&&_bfbg [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _beb [0]=='s'{_beb ,_ =_fae ._fgc .Peek (10);if string (_beb [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_gba :=6;
if len (_beb )> 6{if _cb .IsWhiteSpace (_beb [_gba ])&&_beb [_gba ]!='\r'&&_beb [_gba ]!='\n'{_ab .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");
_gba ++;};if _beb [_gba ]=='\r'{_gba ++;if _beb [_gba ]=='\n'{_gba ++;};}else if _beb [_gba ]=='\n'{_gba ++;};};_fae ._fgc .Discard (_gba );_cfdf ,_bcg :=_bba .PdfObject .(*_cb .PdfObjectDictionary );if !_bcg {return nil ,_g .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_ab .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_cfdf );_eaa ,_gbf :=_cfdf .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_cb .PdfObjectInteger );if !_gbf {return nil ,_g .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");
};_bbdb :=*_eaa ;if _bbdb < 0{return nil ,_g .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_bbdb )> _fae ._ae {_ab .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");
return nil ,_g .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_bd :=make ([]byte ,_bbdb );
_ ,_dedd =_fae .readAtLeast (_bd ,int (_bbdb ));if _dedd !=nil {_ab .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_bd ),_bd );_ab .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_dedd );
return nil ,_dedd ;};_gae :=_cb .PdfObjectStream {};_gae .Stream =_bd ;_gae .PdfObjectDictionary =_bba .PdfObject .(*_cb .PdfObjectDictionary );_gae .ObjectNumber =_bba .ObjectNumber ;_gae .GenerationNumber =_bba .GenerationNumber ;_fae .skipSpaces ();
_fae ._fgc .Discard (9);_fae .skipSpaces ();return &_gae ,nil ;};};_bba .PdfObject ,_dedd =_fae .parseObject ();return &_bba ,_dedd ;};};_ab .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");
return &_bba ,nil ;};func _ggbc (_agf string )(*fdfParser ,error ){_ecee :=fdfParser {};_cfde :=[]byte (_agf );_ebb :=_fe .NewReader (_cfde );_ecee ._cge =_ebb ;_ecee ._af =map[int64 ]_cb .PdfObject {};_bfd :=_gf .NewReader (_ebb );_ecee ._fgc =_bfd ;_ecee ._ae =int64 (len (_agf ));
return &_ecee ,_ecee .parse ();};

// Data represents forms data format (FDF) file data.
type Data struct{_fd *_cb .PdfObjectDictionary ;_cc *_cb .PdfObjectArray ;};var _ed =_b .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");func (_eaad *fdfParser )trace (_cfab _cb .PdfObject )_cb .PdfObject {switch _dbae :=_cfab .(type ){case *_cb .PdfObjectReference :_ecc ,_gfb :=_eaad ._af [_dbae .ObjectNumber ].(*_cb .PdfIndirectObject );
if _gfb {return _ecc .PdfObject ;};_ab .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");return nil ;case *_cb .PdfIndirectObject :return _dbae .PdfObject ;};return _cfab ;};

// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_cb .PdfObject ,error ){_ebd ,_ac :=fdf .FieldDictionaries ();if _ac !=nil {return nil ,_ac ;};var _cab []string ;for _ee :=range _ebd {_cab =append (_cab ,_ee );};_ca .Strings (_cab );_db :=map[string ]_cb .PdfObject {};
for _ ,_ag :=range _cab {_fdb :=_ebd [_ag ];_cee :=_cb .TraceToDirectObject (_fdb .Get ("\u0056"));_db [_ag ]=_cee ;};return _db ,nil ;};

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_ef ,_cg :=_e .Open (fdfPath );if _cg !=nil {return nil ,_cg ;};defer _ef .Close ();return Load (_ef );};