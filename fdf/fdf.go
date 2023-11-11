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
package fdf ;import (_ad "bufio";_bb "bytes";_ac "encoding/hex";_a "errors";_ge "fmt";_f "github.com/unidoc/unipdf/v3/common";_da "github.com/unidoc/unipdf/v3/core";_ab "io";_c "os";_gg "regexp";_d "sort";_e "strconv";_b "strings";);

// LoadFromPath loads FDF form data from file path `fdfPath`.
func LoadFromPath (fdfPath string )(*Data ,error ){_ef ,_fe :=_c .Open (fdfPath );if _fe !=nil {return nil ,_fe ;};defer _ef .Close ();return Load (_ef );};func (_abfc *fdfParser )parseNumber ()(_da .PdfObject ,error ){return _da .ParseNumber (_abfc ._gbc )};
func (_fcce *fdfParser )parseFdfVersion ()(int ,int ,error ){_fcce ._fac .Seek (0,_ab .SeekStart );_fce :=20;_cbf :=make ([]byte ,_fce );_fcce ._fac .Read (_cbf );_cceb :=_db .FindStringSubmatch (string (_cbf ));if len (_cceb )< 3{_bdf ,_egg ,_fcef :=_fcce .seekFdfVersionTopDown ();
if _fcef !=nil {_f .Log .Debug ("F\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065\u0063\u006f\u0076\u0065\u0072\u0079\u0020\u002d\u0020\u0075n\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0066\u0069nd\u0020\u0076\u0065r\u0073i\u006f\u006e");return 0,0,_fcef ;
};return _bdf ,_egg ,nil ;};_ggf ,_dcb :=_e .Atoi (_cceb [1]);if _dcb !=nil {return 0,0,_dcb ;};_dad ,_dcb :=_e .Atoi (_cceb [2]);if _dcb !=nil {return 0,0,_dcb ;};_f .Log .Debug ("\u0046\u0064\u0066\u0020\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020%\u0064\u002e\u0025\u0064",_ggf ,_dad );
return _ggf ,_dad ,nil ;};func (_gfc *fdfParser )setFileOffset (_abf int64 ){_gfc ._fac .Seek (_abf ,_ab .SeekStart );_gfc ._gbc =_ad .NewReader (_gfc ._fac );};func (_efc *fdfParser )readTextLine ()(string ,error ){var _fdg _bb .Buffer ;for {_cdg ,_cag :=_efc ._gbc .Peek (1);
if _cag !=nil {_f .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_cag .Error ());return _fdg .String (),_cag ;};if (_cdg [0]!='\r')&&(_cdg [0]!='\n'){_ed ,_ :=_efc ._gbc .ReadByte ();_fdg .WriteByte (_ed );}else {break ;};};return _fdg .String (),nil ;
};func (_aee *fdfParser )parseNull ()(_da .PdfObjectNull ,error ){_ ,_gbe :=_aee ._gbc .Discard (4);return _da .PdfObjectNull {},_gbe ;};func (_efa *fdfParser )getFileOffset ()int64 {_acg ,_ :=_efa ._fac .Seek (0,_ab .SeekCurrent );_acg -=int64 (_efa ._gbc .Buffered ());
return _acg ;};var _cf =_gg .MustCompile ("^\u005c\u0073\u002a\u0028\\d\u002b)\u005c\u0073\u002b\u0028\u005cd\u002b\u0029\u005c\u0073\u002b\u0052");var _bfc =_gg .MustCompile ("\u005e\u005b\u005c+-\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e]\u002b)\u0065[\u005c+\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d\u0039\u002e\u005d\u002b\u0029");


// FieldDictionaries returns a map of field names to field dictionaries.
func (fdf *Data )FieldDictionaries ()(map[string ]*_da .PdfObjectDictionary ,error ){_daf :=map[string ]*_da .PdfObjectDictionary {};for _dg :=0;_dg < fdf ._be .Len ();_dg ++{_fb ,_efd :=_da .GetDict (fdf ._be .Get (_dg ));if _efd {_ce ,_ :=_da .GetString (_fb .Get ("\u0054"));
if _ce !=nil {_daf [_ce .Str ()]=_fb ;};};};return _daf ,nil ;};func (_agb *fdfParser )skipComments ()error {if _ ,_bfca :=_agb .skipSpaces ();_bfca !=nil {return _bfca ;};_dcc :=true ;for {_eg ,_aed :=_agb ._gbc .Peek (1);if _aed !=nil {_f .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_aed .Error ());
return _aed ;};if _dcc &&_eg [0]!='%'{return nil ;};_dcc =false ;if (_eg [0]!='\r')&&(_eg [0]!='\n'){_agb ._gbc .ReadByte ();}else {break ;};};return _agb .skipComments ();};func (_fgec *fdfParser )seekFdfVersionTopDown ()(int ,int ,error ){_fgec ._fac .Seek (0,_ab .SeekStart );
_fgec ._gbc =_ad .NewReader (_fgec ._fac );_afb :=20;_eca :=make ([]byte ,_afb );for {_fgg ,_abg :=_fgec ._gbc .ReadByte ();if _abg !=nil {if _abg ==_ab .EOF {break ;}else {return 0,0,_abg ;};};if _da .IsDecimalDigit (_fgg )&&_eca [_afb -1]=='.'&&_da .IsDecimalDigit (_eca [_afb -2])&&_eca [_afb -3]=='-'&&_eca [_afb -4]=='F'&&_eca [_afb -5]=='D'&&_eca [_afb -6]=='P'{_bga :=int (_eca [_afb -2]-'0');
_acc :=int (_fgg -'0');return _bga ,_acc ,nil ;};_eca =append (_eca [1:_afb ],_fgg );};return 0,0,_a .New ("\u0076\u0065\u0072\u0073\u0069\u006f\u006e\u0020\u006e\u006f\u0074\u0020f\u006f\u0075\u006e\u0064");};func (_ebc *fdfParser )readAtLeast (_adcf []byte ,_bgg int )(int ,error ){_ebb :=_bgg ;
_gf :=0;_eaf :=0;for _ebb > 0{_bed ,_abe :=_ebc ._gbc .Read (_adcf [_gf :]);if _abe !=nil {_f .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0046\u0061i\u006c\u0065\u0064\u0020\u0072\u0065\u0061d\u0069\u006e\u0067\u0020\u0028\u0025\u0064\u003b\u0025\u0064\u0029\u0020\u0025\u0073",_bed ,_eaf ,_abe .Error ());
return _gf ,_a .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0072\u0065a\u0064\u0069\u006e\u0067");};_eaf ++;_gf +=_bed ;_ebb -=_bed ;};return _gf ,nil ;};func (_gcea *fdfParser )parseDict ()(*_da .PdfObjectDictionary ,error ){_f .Log .Trace ("\u0052\u0065\u0061\u0064\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020D\u0069\u0063\u0074\u0021");
_cfe :=_da .MakeDict ();_gda ,_ :=_gcea ._gbc .ReadByte ();if _gda !='<'{return nil ,_a .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");};_gda ,_ =_gcea ._gbc .ReadByte ();if _gda !='<'{return nil ,_a .New ("\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0064\u0069\u0063\u0074");
};for {_gcea .skipSpaces ();_gcea .skipComments ();_fcb ,_afe :=_gcea ._gbc .Peek (2);if _afe !=nil {return nil ,_afe ;};_f .Log .Trace ("D\u0069c\u0074\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_fcb ),string (_fcb ));
if (_fcb [0]=='>')&&(_fcb [1]=='>'){_f .Log .Trace ("\u0045\u004f\u0046\u0020\u0064\u0069\u0063\u0074\u0069o\u006e\u0061\u0072\u0079");_gcea ._gbc .ReadByte ();_gcea ._gbc .ReadByte ();break ;};_f .Log .Trace ("\u0050a\u0072s\u0065\u0020\u0074\u0068\u0065\u0020\u006e\u0061\u006d\u0065\u0021");
_cdb ,_afe :=_gcea .parseName ();_f .Log .Trace ("\u004be\u0079\u003a\u0020\u0025\u0073",_cdb );if _afe !=nil {_f .Log .Debug ("E\u0052\u0052\u004f\u0052\u0020\u0052e\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u006ea\u006d\u0065\u0020e\u0072r\u0020\u0025\u0073",_afe );
return nil ,_afe ;};if len (_cdb )> 4&&_cdb [len (_cdb )-4:]=="\u006e\u0075\u006c\u006c"{_gbcd :=_cdb [0:len (_cdb )-4];_f .Log .Debug ("\u0054\u0061\u006b\u0069n\u0067\u0020\u0063\u0061\u0072\u0065\u0020\u006f\u0066\u0020n\u0075l\u006c\u0020\u0062\u0075\u0067\u0020\u0028%\u0073\u0029",_cdb );
_f .Log .Debug ("\u004e\u0065\u0077\u0020ke\u0079\u0020\u0022\u0025\u0073\u0022\u0020\u003d\u0020\u006e\u0075\u006c\u006c",_gbcd );_gcea .skipSpaces ();_acgf ,_ :=_gcea ._gbc .Peek (1);if _acgf [0]=='/'{_cfe .Set (_gbcd ,_da .MakeNull ());continue ;};};
_gcea .skipSpaces ();_ebf ,_afe :=_gcea .parseObject ();if _afe !=nil {return nil ,_afe ;};_cfe .Set (_cdb ,_ebf );_f .Log .Trace ("\u0064\u0069\u0063\u0074\u005b\u0025\u0073\u005d\u0020\u003d\u0020\u0025\u0073",_cdb ,_ebf .String ());};_f .Log .Trace ("\u0072\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0046\u0044\u0046\u0020\u0044\u0069\u0063\u0074\u0021");
return _cfe ,nil ;};

// Load loads FDF form data from `r`.
func Load (r _ab .ReadSeeker )(*Data ,error ){_ff ,_bg :=_ffcf (r );if _bg !=nil {return nil ,_bg ;};_cb ,_bg :=_ff .Root ();if _bg !=nil {return nil ,_bg ;};_ca ,_cg :=_da .GetArray (_cb .Get ("\u0046\u0069\u0065\u006c\u0064\u0073"));if !_cg {return nil ,_a .New ("\u0066\u0069\u0065\u006c\u0064\u0073\u0020\u006d\u0069s\u0073\u0069\u006e\u0067");
};return &Data {_be :_ca ,_dd :_cb },nil ;};func (_aef *fdfParser )readComment ()(string ,error ){var _ged _bb .Buffer ;_ ,_ba :=_aef .skipSpaces ();if _ba !=nil {return _ged .String (),_ba ;};_cd :=true ;for {_cfa ,_ace :=_aef ._gbc .Peek (1);if _ace !=nil {_f .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_ace .Error ());
return _ged .String (),_ace ;};if _cd &&_cfa [0]!='%'{return _ged .String (),_a .New ("c\u006f\u006d\u006d\u0065\u006e\u0074 \u0073\u0068\u006f\u0075\u006c\u0064\u0020\u0073\u0074a\u0072\u0074\u0020w\u0069t\u0068\u0020\u0025");};_cd =false ;if (_cfa [0]!='\r')&&(_cfa [0]!='\n'){_ga ,_ :=_aef ._gbc .ReadByte ();
_ged .WriteByte (_ga );}else {break ;};};return _ged .String (),nil ;};var _ffd =_gg .MustCompile ("\u0028\u005c\u0064\u002b)\\\u0073\u002b\u0028\u005c\u0064\u002b\u0029\u005c\u0073\u002b\u006f\u0062\u006a");func (_adbe *fdfParser )parseIndirectObject ()(_da .PdfObject ,error ){_egc :=_da .PdfIndirectObject {};
_f .Log .Trace ("\u002dR\u0065a\u0064\u0020\u0069\u006e\u0064i\u0072\u0065c\u0074\u0020\u006f\u0062\u006a");_fdba ,_cbdg :=_adbe ._gbc .Peek (20);if _cbdg !=nil {_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return &_egc ,_cbdg ;};_f .Log .Trace ("\u0028\u0069\u006edi\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0020\u0070\u0065\u0065\u006b\u0020\u0022\u0025\u0073\u0022",string (_fdba ));_bfcdc :=_ffd .FindStringSubmatchIndex (string (_fdba ));if len (_bfcdc )< 6{_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_fdba ));
return &_egc ,_a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_adbe ._gbc .Discard (_bfcdc [0]);_f .Log .Trace ("O\u0066\u0066\u0073\u0065\u0074\u0073\u0020\u0025\u0020\u0064",_bfcdc );_cbad :=_bfcdc [1]-_bfcdc [0];_bdb :=make ([]byte ,_cbad );_ ,_cbdg =_adbe .readAtLeast (_bdb ,_cbad );if _cbdg !=nil {_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0075\u006e\u0061\u0062l\u0065\u0020\u0074\u006f\u0020\u0072\u0065\u0061\u0064\u0020-\u0020\u0025\u0073",_cbdg );
return nil ,_cbdg ;};_f .Log .Trace ("\u0074\u0065\u0078t\u006c\u0069\u006e\u0065\u003a\u0020\u0025\u0073",_bdb );_ccaa :=_ffd .FindStringSubmatch (string (_bdb ));if len (_ccaa )< 3{_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_bdb ));
return &_egc ,_a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_bdd ,_ :=_e .Atoi (_ccaa [1]);_fbb ,_ :=_e .Atoi (_ccaa [2]);_egc .ObjectNumber =int64 (_bdd );_egc .GenerationNumber =int64 (_fbb );for {_gfd ,_ecgd :=_adbe ._gbc .Peek (2);if _ecgd !=nil {return &_egc ,_ecgd ;};_f .Log .Trace ("I\u006ed\u002e\u0020\u0070\u0065\u0065\u006b\u003a\u0020%\u0073\u0020\u0028\u0025 x\u0029\u0021",string (_gfd ),string (_gfd ));
if _da .IsWhiteSpace (_gfd [0]){_adbe .skipSpaces ();}else if _gfd [0]=='%'{_adbe .skipComments ();}else if (_gfd [0]=='<')&&(_gfd [1]=='<'){_f .Log .Trace ("\u0043\u0061\u006c\u006c\u0020\u0050\u0061\u0072\u0073e\u0044\u0069\u0063\u0074");_egc .PdfObject ,_ecgd =_adbe .parseDict ();
_f .Log .Trace ("\u0045\u004f\u0046\u0020Ca\u006c\u006c\u0020\u0050\u0061\u0072\u0073\u0065\u0044\u0069\u0063\u0074\u003a\u0020%\u0076",_ecgd );if _ecgd !=nil {return &_egc ,_ecgd ;};_f .Log .Trace ("\u0050\u0061\u0072\u0073\u0065\u0064\u0020\u0064\u0069\u0063t\u0069\u006f\u006e\u0061\u0072\u0079\u002e.\u002e\u0020\u0066\u0069\u006e\u0069\u0073\u0068\u0065\u0064\u002e");
}else if (_gfd [0]=='/')||(_gfd [0]=='(')||(_gfd [0]=='[')||(_gfd [0]=='<'){_egc .PdfObject ,_ecgd =_adbe .parseObject ();if _ecgd !=nil {return &_egc ,_ecgd ;};_f .Log .Trace ("P\u0061\u0072\u0073\u0065\u0064\u0020o\u0062\u006a\u0065\u0063\u0074\u0020\u002e\u002e\u002e \u0066\u0069\u006ei\u0073h\u0065\u0064\u002e");
}else {if _gfd [0]=='e'{_ee ,_ffc :=_adbe .readTextLine ();if _ffc !=nil {return nil ,_ffc ;};if len (_ee )>=6&&_ee [0:6]=="\u0065\u006e\u0064\u006f\u0062\u006a"{break ;};}else if _gfd [0]=='s'{_gfd ,_ =_adbe ._gbc .Peek (10);if string (_gfd [:6])=="\u0073\u0074\u0072\u0065\u0061\u006d"{_bbc :=6;
if len (_gfd )> 6{if _da .IsWhiteSpace (_gfd [_bbc ])&&_gfd [_bbc ]!='\r'&&_gfd [_bbc ]!='\n'{_f .Log .Debug ("\u004e\u006fn\u002d\u0063\u006f\u006e\u0066\u006f\u0072\u006d\u0061\u006e\u0074\u0020\u0046\u0044\u0046\u0020\u006e\u006f\u0074 \u0065\u006e\u0064\u0069\u006e\u0067 \u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006c\u0069\u006e\u0065\u0020\u0070\u0072o\u0070\u0065r\u006c\u0079\u0020\u0077i\u0074\u0068\u0020\u0045\u004fL\u0020\u006d\u0061\u0072\u006b\u0065\u0072");
_bbc ++;};if _gfd [_bbc ]=='\r'{_bbc ++;if _gfd [_bbc ]=='\n'{_bbc ++;};}else if _gfd [_bbc ]=='\n'{_bbc ++;};};_adbe ._gbc .Discard (_bbc );_gfcff ,_aefe :=_egc .PdfObject .(*_da .PdfObjectDictionary );if !_aefe {return nil ,_a .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006di\u0073s\u0069\u006e\u0067\u0020\u0064\u0069\u0063\u0074\u0069\u006f\u006e\u0061\u0072\u0079");
};_f .Log .Trace ("\u0053\u0074\u0072\u0065\u0061\u006d\u0020\u0064\u0069c\u0074\u0020\u0025\u0073",_gfcff );_dee ,_cbca :=_gfcff .Get ("\u004c\u0065\u006e\u0067\u0074\u0068").(*_da .PdfObjectInteger );if !_cbca {return nil ,_a .New ("\u0073\u0074re\u0061\u006d\u0020l\u0065\u006e\u0067\u0074h n\u0065ed\u0073\u0020\u0074\u006f\u0020\u0062\u0065 a\u006e\u0020\u0069\u006e\u0074\u0065\u0067e\u0072");
};_cgbc :=*_dee ;if _cgbc < 0{return nil ,_a .New ("\u0073\u0074\u0072\u0065\u0061\u006d\u0020\u006e\u0065\u0065\u0064\u0073\u0020\u0074\u006f \u0062e\u0020\u006c\u006f\u006e\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0030");};if int64 (_cgbc )> _adbe ._dc {_f .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0053t\u0072\u0065\u0061\u006d\u0020l\u0065\u006e\u0067\u0074\u0068\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062\u0065\u0020\u006c\u0061\u0072\u0067\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0066\u0069\u006c\u0065\u0020\u0073\u0069\u007a\u0065");
return nil ,_a .New ("\u0069n\u0076\u0061l\u0069\u0064\u0020\u0073t\u0072\u0065\u0061m\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u002c\u0020la\u0072\u0067\u0065r\u0020\u0074h\u0061\u006e\u0020\u0066\u0069\u006ce\u0020\u0073i\u007a\u0065");};_fcg :=make ([]byte ,_cgbc );
_ ,_ecgd =_adbe .readAtLeast (_fcg ,int (_cgbc ));if _ecgd !=nil {_f .Log .Debug ("E\u0052\u0052\u004f\u0052 s\u0074r\u0065\u0061\u006d\u0020\u0028%\u0064\u0029\u003a\u0020\u0025\u0058",len (_fcg ),_fcg );_f .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025v",_ecgd );
return nil ,_ecgd ;};_aegc :=_da .PdfObjectStream {};_aegc .Stream =_fcg ;_aegc .PdfObjectDictionary =_egc .PdfObject .(*_da .PdfObjectDictionary );_aegc .ObjectNumber =_egc .ObjectNumber ;_aegc .GenerationNumber =_egc .GenerationNumber ;_adbe .skipSpaces ();
_adbe ._gbc .Discard (9);_adbe .skipSpaces ();return &_aegc ,nil ;};};_egc .PdfObject ,_ecgd =_adbe .parseObject ();return &_egc ,_ecgd ;};};_f .Log .Trace ("\u0052\u0065\u0074\u0075rn\u0069\u006e\u0067\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0021");
return &_egc ,nil ;};func (_ae *fdfParser )skipSpaces ()(int ,error ){_efdg :=0;for {_gbcg ,_gc :=_ae ._gbc .ReadByte ();if _gc !=nil {return 0,_gc ;};if _da .IsWhiteSpace (_gbcg ){_efdg ++;}else {_ae ._gbc .UnreadByte ();break ;};};return _efdg ,nil ;
};func (_ec *fdfParser )parseName ()(_da .PdfObjectName ,error ){var _gbcb _bb .Buffer ;_cce :=false ;for {_aedc ,_gde :=_ec ._gbc .Peek (1);if _gde ==_ab .EOF {break ;};if _gde !=nil {return _da .PdfObjectName (_gbcb .String ()),_gde ;};if !_cce {if _aedc [0]=='/'{_cce =true ;
_ec ._gbc .ReadByte ();}else if _aedc [0]=='%'{_ec .readComment ();_ec .skipSpaces ();}else {_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020N\u0061\u006d\u0065\u0020\u0073\u0074\u0061\u0072\u0074\u0069\u006e\u0067\u0020w\u0069\u0074\u0068\u0020\u0025\u0073\u0020(\u0025\u0020\u0078\u0029",_aedc ,_aedc );
return _da .PdfObjectName (_gbcb .String ()),_ge .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u006ea\u006d\u0065:\u0020\u0028\u0025\u0063\u0029",_aedc [0]);};}else {if _da .IsWhiteSpace (_aedc [0]){break ;}else if (_aedc [0]=='/')||(_aedc [0]=='[')||(_aedc [0]=='(')||(_aedc [0]==']')||(_aedc [0]=='<')||(_aedc [0]=='>'){break ;
}else if _aedc [0]=='#'{_feg ,_dfg :=_ec ._gbc .Peek (3);if _dfg !=nil {return _da .PdfObjectName (_gbcb .String ()),_dfg ;};_ec ._gbc .Discard (3);_aa ,_dfg :=_ac .DecodeString (string (_feg [1:3]));if _dfg !=nil {return _da .PdfObjectName (_gbcb .String ()),_dfg ;
};_gbcb .Write (_aa );}else {_caf ,_ :=_ec ._gbc .ReadByte ();_gbcb .WriteByte (_caf );};};};return _da .PdfObjectName (_gbcb .String ()),nil ;};

// Data represents forms data format (FDF) file data.
type Data struct{_dd *_da .PdfObjectDictionary ;_be *_da .PdfObjectArray ;};func (_bfd *fdfParser )parseHexString ()(*_da .PdfObjectString ,error ){_bfd ._gbc .ReadByte ();var _gcf _bb .Buffer ;for {_gdg ,_bd :=_bfd ._gbc .Peek (1);if _bd !=nil {return _da .MakeHexString (""),_bd ;
};if _gdg [0]=='>'{_bfd ._gbc .ReadByte ();break ;};_gef ,_ :=_bfd ._gbc .ReadByte ();if !_da .IsWhiteSpace (_gef ){_gcf .WriteByte (_gef );};};if _gcf .Len ()%2==1{_gcf .WriteRune ('0');};_agc ,_cfd :=_ac .DecodeString (_gcf .String ());if _cfd !=nil {_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u0020\u0050\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0068\u0065\u0078\u0020\u0073\u0074r\u0069\u006e\u0067\u003a\u0020\u0027\u0025\u0073\u0027 \u002d\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067\u0020\u0061n\u0020\u0065\u006d\u0070\u0074\u0079 \u0073\u0074\u0072i\u006e\u0067",_gcf .String ());
return _da .MakeHexString (""),nil ;};return _da .MakeHexString (string (_agc )),nil ;};var _gee =_gg .MustCompile ("\u0025\u0025\u0045O\u0046");var _db =_gg .MustCompile ("\u0025F\u0044F\u002d\u0028\u005c\u0064\u0029\u005c\u002e\u0028\u005c\u0064\u0029");


// FieldValues implements interface model.FieldValueProvider.
// Returns a map of field names to values (PdfObjects).
func (fdf *Data )FieldValues ()(map[string ]_da .PdfObject ,error ){_bgd ,_gd :=fdf .FieldDictionaries ();if _gd !=nil {return nil ,_gd ;};var _ced []string ;for _cc :=range _bgd {_ced =append (_ced ,_cc );};_d .Strings (_ced );_fd :=map[string ]_da .PdfObject {};
for _ ,_gb :=range _ced {_ea :=_bgd [_gb ];_abc :=_da .TraceToDirectObject (_ea .Get ("\u0056"));_fd [_gb ]=_abc ;};return _fd ,nil ;};type fdfParser struct{_ebe int ;_cbd int ;_ag map[int64 ]_da .PdfObject ;_fac _ab .ReadSeeker ;_gbc *_ad .Reader ;_dc int64 ;
_bgf *_da .PdfObjectDictionary ;};

// Root returns the Root of the FDF document.
func (_fdf *fdfParser )Root ()(*_da .PdfObjectDictionary ,error ){if _fdf ._bgf !=nil {if _bbd ,_gecd :=_fdf .trace (_fdf ._bgf .Get ("\u0052\u006f\u006f\u0074")).(*_da .PdfObjectDictionary );_gecd {if _adbed ,_dda :=_fdf .trace (_bbd .Get ("\u0046\u0044\u0046")).(*_da .PdfObjectDictionary );
_dda {return _adbed ,nil ;};};};var _cfeb []int64 ;for _fab :=range _fdf ._ag {_cfeb =append (_cfeb ,_fab );};_d .Slice (_cfeb ,func (_agce ,_ccg int )bool {return _cfeb [_agce ]< _cfeb [_ccg ]});for _ ,_gae :=range _cfeb {_bdfb :=_fdf ._ag [_gae ];if _ead ,_ega :=_fdf .trace (_bdfb ).(*_da .PdfObjectDictionary );
_ega {if _deed ,_aea :=_fdf .trace (_ead .Get ("\u0046\u0044\u0046")).(*_da .PdfObjectDictionary );_aea {return _deed ,nil ;};};};return nil ,_a .New ("\u0046\u0044\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");};func (_gce *fdfParser )parseString ()(*_da .PdfObjectString ,error ){_gce ._gbc .ReadByte ();
var _ggd _bb .Buffer ;_cgb :=1;for {_fdb ,_gfcf :=_gce ._gbc .Peek (1);if _gfcf !=nil {return _da .MakeString (_ggd .String ()),_gfcf ;};if _fdb [0]=='\\'{_gce ._gbc .ReadByte ();_de ,_dfd :=_gce ._gbc .ReadByte ();if _dfd !=nil {return _da .MakeString (_ggd .String ()),_dfd ;
};if _da .IsOctalDigit (_de ){_ecg ,_fdd :=_gce ._gbc .Peek (2);if _fdd !=nil {return _da .MakeString (_ggd .String ()),_fdd ;};var _efde []byte ;_efde =append (_efde ,_de );for _ ,_cca :=range _ecg {if _da .IsOctalDigit (_cca ){_efde =append (_efde ,_cca );
}else {break ;};};_gce ._gbc .Discard (len (_efde )-1);_f .Log .Trace ("\u004e\u0075\u006d\u0065ri\u0063\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0020\u0022\u0025\u0073\u0022",_efde );_fed ,_fdd :=_e .ParseUint (string (_efde ),8,32);if _fdd !=nil {return _da .MakeString (_ggd .String ()),_fdd ;
};_ggd .WriteByte (byte (_fed ));continue ;};switch _de {case 'n':_ggd .WriteRune ('\n');case 'r':_ggd .WriteRune ('\r');case 't':_ggd .WriteRune ('\t');case 'b':_ggd .WriteRune ('\b');case 'f':_ggd .WriteRune ('\f');case '(':_ggd .WriteRune ('(');case ')':_ggd .WriteRune (')');
case '\\':_ggd .WriteRune ('\\');};continue ;}else if _fdb [0]=='('{_cgb ++;}else if _fdb [0]==')'{_cgb --;if _cgb ==0{_gce ._gbc .ReadByte ();break ;};};_gge ,_ :=_gce ._gbc .ReadByte ();_ggd .WriteByte (_gge );};return _da .MakeString (_ggd .String ()),nil ;
};func (_fdbe *fdfParser )parseArray ()(*_da .PdfObjectArray ,error ){_cdd :=_da .MakeArray ();_fdbe ._gbc .ReadByte ();for {_fdbe .skipSpaces ();_fea ,_ddf :=_fdbe ._gbc .Peek (1);if _ddf !=nil {return _cdd ,_ddf ;};if _fea [0]==']'{_fdbe ._gbc .ReadByte ();
break ;};_faa ,_ddf :=_fdbe .parseObject ();if _ddf !=nil {return _cdd ,_ddf ;};_cdd .Append (_faa );};return _cdd ,nil ;};func _abb (_eag string )(_da .PdfObjectReference ,error ){_aac :=_da .PdfObjectReference {};_bab :=_cf .FindStringSubmatch (_eag );
if len (_bab )< 3{_f .Log .Debug ("\u0045\u0072\u0072or\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0072\u0065\u0066\u0065\u0072\u0065\u006e\u0063\u0065");return _aac ,_a .New ("\u0075n\u0061\u0062\u006c\u0065 \u0074\u006f\u0020\u0070\u0061r\u0073e\u0020r\u0065\u0066\u0065\u0072\u0065\u006e\u0063e");
};_gcfd ,_gag :=_e .Atoi (_bab [1]);if _gag !=nil {_f .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0070a\u0072\u0073\u0069n\u0067\u0020\u006fb\u006a\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072 '\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u006e\u0075\u006d\u0020\u003d\u0020\u0030",_bab [1]);
return _aac ,nil ;};_aac .ObjectNumber =int64 (_gcfd );_ecd ,_gag :=_e .Atoi (_bab [2]);if _gag !=nil {_f .Log .Debug ("\u0045\u0072r\u006f\u0072\u0020\u0070\u0061r\u0073\u0069\u006e\u0067\u0020g\u0065\u006e\u0065\u0072\u0061\u0074\u0069\u006f\u006e\u0020\u006e\u0075\u006d\u0062\u0065\u0072\u0020\u0027\u0025\u0073\u0027\u0020\u002d\u0020\u0055\u0073\u0069\u006e\u0067\u0020\u0067\u0065\u006e\u0020\u003d\u0020\u0030",_bab [2]);
return _aac ,nil ;};_aac .GenerationNumber =int64 (_ecd );return _aac ,nil ;};func (_bgb *fdfParser )trace (_fdc _da .PdfObject )_da .PdfObject {switch _gabf :=_fdc .(type ){case *_da .PdfObjectReference :_dca ,_cde :=_bgb ._ag [_gabf .ObjectNumber ].(*_da .PdfIndirectObject );
if _cde {return _dca .PdfObject ;};_f .Log .Debug ("\u0054\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");return nil ;case *_da .PdfIndirectObject :return _gabf .PdfObject ;};return _fdc ;};func (_cbc *fdfParser )seekToEOFMarker (_gab int64 )error {_gbbd :=int64 (0);
_bc :=int64 (1000);for _gbbd < _gab {if _gab <=(_bc +_gbbd ){_bc =_gab -_gbbd ;};_ ,_ffgb :=_cbc ._fac .Seek (-_gbbd -_bc ,_ab .SeekEnd );if _ffgb !=nil {return _ffgb ;};_bge :=make ([]byte ,_bc );_cbc ._fac .Read (_bge );_f .Log .Trace ("\u004c\u006f\u006f\u006bi\u006e\u0067\u0020\u0066\u006f\u0072\u0020\u0045\u004f\u0046 \u006da\u0072\u006b\u0065\u0072\u003a\u0020\u0022%\u0073\u0022",string (_bge ));
_bbe :=_gee .FindAllStringIndex (string (_bge ),-1);if _bbe !=nil {_bgge :=_bbe [len (_bbe )-1];_f .Log .Trace ("\u0049\u006e\u0064\u003a\u0020\u0025\u0020\u0064",_bbe );_cbc ._fac .Seek (-_gbbd -_bc +int64 (_bgge [0]),_ab .SeekEnd );return nil ;};_f .Log .Debug ("\u0057\u0061\u0072\u006e\u0069\u006eg\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006b\u0065\u0072\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064\u0021\u0020\u002d\u0020\u0063\u006f\u006e\u0074\u0069\u006e\u0075\u0065\u0020s\u0065e\u006b\u0069\u006e\u0067");
_gbbd +=_bc ;};_f .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u003a\u0020\u0045\u004f\u0046\u0020\u006d\u0061\u0072\u006be\u0072 \u0077\u0061\u0073\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064\u002e");return _a .New ("\u0045\u004f\u0046\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075\u006e\u0064");
};func _ffcf (_gbbg _ab .ReadSeeker )(*fdfParser ,error ){_ffda :=&fdfParser {};_ffda ._fac =_gbbg ;_ffda ._ag =map[int64 ]_da .PdfObject {};_ccaae ,_ebed ,_cec :=_ffda .parseFdfVersion ();if _cec !=nil {_f .Log .Error ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0076e\u0072\u0073\u0069o\u006e:\u0020\u0025\u0076",_cec );
return nil ,_cec ;};_ffda ._ebe =_ccaae ;_ffda ._cbd =_ebed ;_cec =_ffda .parse ();return _ffda ,_cec ;};func _edf (_gedf string )(*fdfParser ,error ){_bacd :=fdfParser {};_dfb :=[]byte (_gedf );_egca :=_bb .NewReader (_dfb );_bacd ._fac =_egca ;_bacd ._ag =map[int64 ]_da .PdfObject {};
_cgbb :=_ad .NewReader (_egca );_bacd ._gbc =_cgbb ;_bacd ._dc =int64 (len (_gedf ));return &_bacd ,_bacd .parse ();};func (_gbb *fdfParser )parseBool ()(_da .PdfObjectBool ,error ){_ccf ,_fg :=_gbb ._gbc .Peek (4);if _fg !=nil {return _da .PdfObjectBool (false ),_fg ;
};if (len (_ccf )>=4)&&(string (_ccf [:4])=="\u0074\u0072\u0075\u0065"){_gbb ._gbc .Discard (4);return _da .PdfObjectBool (true ),nil ;};_ccf ,_fg =_gbb ._gbc .Peek (5);if _fg !=nil {return _da .PdfObjectBool (false ),_fg ;};if (len (_ccf )>=5)&&(string (_ccf [:5])=="\u0066\u0061\u006cs\u0065"){_gbb ._gbc .Discard (5);
return _da .PdfObjectBool (false ),nil ;};return _da .PdfObjectBool (false ),_a .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};var _cba =_gg .MustCompile ("\u005e\u005b\u005c\u002b\u002d\u002e\u005d\u002a\u0028\u005b\u0030\u002d9\u002e\u005d\u002b\u0029");
func (_fff *fdfParser )parseObject ()(_da .PdfObject ,error ){_f .Log .Trace ("\u0052e\u0061d\u0020\u0064\u0069\u0072\u0065c\u0074\u0020o\u0062\u006a\u0065\u0063\u0074");_fff .skipSpaces ();for {_bda ,_dgd :=_fff ._gbc .Peek (2);if _dgd !=nil {return nil ,_dgd ;
};_f .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_bda ));if _bda [0]=='/'{_ffb ,_ffg :=_fff .parseName ();_f .Log .Trace ("\u002d\u003e\u004ea\u006d\u0065\u003a\u0020\u0027\u0025\u0073\u0027",_ffb );
return &_ffb ,_ffg ;}else if _bda [0]=='('{_f .Log .Trace ("\u002d>\u0053\u0074\u0072\u0069\u006e\u0067!");return _fff .parseString ();}else if _bda [0]=='['{_f .Log .Trace ("\u002d\u003e\u0041\u0072\u0072\u0061\u0079\u0021");return _fff .parseArray ();
}else if (_bda [0]=='<')&&(_bda [1]=='<'){_f .Log .Trace ("\u002d>\u0044\u0069\u0063\u0074\u0021");return _fff .parseDict ();}else if _bda [0]=='<'{_f .Log .Trace ("\u002d\u003e\u0048\u0065\u0078\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u0021");return _fff .parseHexString ();
}else if _bda [0]=='%'{_fff .readComment ();_fff .skipSpaces ();}else {_f .Log .Trace ("\u002d\u003eN\u0075\u006d\u0062e\u0072\u0020\u006f\u0072\u0020\u0072\u0065\u0066\u003f");_bda ,_ =_fff ._gbc .Peek (15);_bdag :=string (_bda );_f .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_bdag );
if (len (_bdag )> 3)&&(_bdag [:4]=="\u006e\u0075\u006c\u006c"){_dbb ,_bggg :=_fff .parseNull ();return &_dbb ,_bggg ;}else if (len (_bdag )> 4)&&(_bdag [:5]=="\u0066\u0061\u006cs\u0065"){_fcc ,_gcec :=_fff .parseBool ();return &_fcc ,_gcec ;}else if (len (_bdag )> 3)&&(_bdag [:4]=="\u0074\u0072\u0075\u0065"){_dab ,_bae :=_fff .parseBool ();
return &_dab ,_bae ;};_af :=_cf .FindStringSubmatch (_bdag );if len (_af )> 1{_bda ,_ =_fff ._gbc .ReadBytes ('R');_f .Log .Trace ("\u002d\u003e\u0020\u0021\u0052\u0065\u0066\u003a\u0020\u0027\u0025\u0073\u0027",string (_bda [:]));_bfa ,_fgb :=_abb (string (_bda ));
return &_bfa ,_fgb ;};_edd :=_cba .FindStringSubmatch (_bdag );if len (_edd )> 1{_f .Log .Trace ("\u002d\u003e\u0020\u004e\u0075\u006d\u0062\u0065\u0072\u0021");return _fff .parseNumber ();};_edd =_bfc .FindStringSubmatch (_bdag );if len (_edd )> 1{_f .Log .Trace ("\u002d\u003e\u0020\u0045xp\u006f\u006e\u0065\u006e\u0074\u0069\u0061\u006c\u0020\u004e\u0075\u006d\u0062\u0065r\u0021");
_f .Log .Trace ("\u0025\u0020\u0073",_edd );return _fff .parseNumber ();};_f .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020U\u006e\u006b\u006e\u006f\u0077n\u0020(\u0070e\u0065\u006b\u0020\u0022\u0025\u0073\u0022)",_bdag );return nil ,_a .New ("\u006f\u0062\u006a\u0065\u0063t\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0065\u0072\u0072\u006fr\u0020\u002d\u0020\u0075\u006e\u0065\u0078\u0070\u0065\u0063\u0074\u0065\u0064\u0020\u0070\u0061\u0074\u0074\u0065\u0072\u006e");
};};};func (_gdgd *fdfParser )parse ()error {_gdgd ._fac .Seek (0,_ab .SeekStart );_gdgd ._gbc =_ad .NewReader (_gdgd ._fac );for {_gdgd .skipComments ();_aca ,_dafg :=_gdgd ._gbc .Peek (20);if _dafg !=nil {_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0046\u0061\u0069\u006c\u0020\u0074\u006f\u0020r\u0065a\u0064\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a");
return _dafg ;};if _b .HasPrefix (string (_aca ),"\u0074r\u0061\u0069\u006c\u0065\u0072"){_gdgd ._gbc .Discard (7);_gdgd .skipSpaces ();_gdgd .skipComments ();_gdc ,_ :=_gdgd .parseDict ();_gdgd ._bgf =_gdc ;break ;};_fddf :=_ffd .FindStringSubmatchIndex (string (_aca ));
if len (_fddf )< 6{_f .Log .Debug ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020U\u006e\u0061\u0062l\u0065\u0020\u0074\u006f \u0066\u0069\u006e\u0064\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020\u0073\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065\u0020\u0028\u0025\u0073\u0029",string (_aca ));
return _a .New ("\u0075\u006e\u0061b\u006c\u0065\u0020\u0074\u006f\u0020\u0064\u0065\u0074\u0065\u0063\u0074\u0020\u0069\u006e\u0064\u0069\u0072\u0065\u0063\u0074\u0020\u006f\u0062\u006a\u0065\u0063\u0074\u0020s\u0069\u0067\u006e\u0061\u0074\u0075\u0072\u0065");
};_gaea ,_dafg :=_gdgd .parseIndirectObject ();if _dafg !=nil {return _dafg ;};switch _eef :=_gaea .(type ){case *_da .PdfIndirectObject :_gdgd ._ag [_eef .ObjectNumber ]=_eef ;case *_da .PdfObjectStream :_gdgd ._ag [_eef .ObjectNumber ]=_eef ;default:return _a .New ("\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};};return nil ;};