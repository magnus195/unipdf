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

package timeutils ;import (_c "errors";_ff "fmt";_b "regexp";_cf "strconv";_cb "time";);func ParsePdfTime (pdfTime string )(_cb .Time ,error ){_db :=_ab .FindAllStringSubmatch (pdfTime ,1);if len (_db )< 1{return _cb .Time {},_ff .Errorf ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020\u0064\u0061\u0074\u0065\u0020s\u0074\u0072\u0069\u006e\u0067\u0020\u0028\u0025\u0073\u0029",pdfTime );
};if len (_db [0])!=10{return _cb .Time {},_c .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0072\u0065\u0067\u0065\u0078p\u0020\u0067\u0072\u006f\u0075\u0070 \u006d\u0061\u0074\u0063\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020!\u003d\u0020\u0031\u0030");
};_bf ,_ :=_cf .ParseInt (_db [0][1],10,32);_da ,_ :=_cf .ParseInt (_db [0][2],10,32);_e ,_ :=_cf .ParseInt (_db [0][3],10,32);_fb ,_ :=_cf .ParseInt (_db [0][4],10,32);_ae ,_ :=_cf .ParseInt (_db [0][5],10,32);_fbg ,_ :=_cf .ParseInt (_db [0][6],10,32);
var (_dg byte ;_ed int64 ;_fdd int64 ;);if len (_db [0][7])> 0{_dg =_db [0][7][0];}else {_dg ='+';};if len (_db [0][8])> 0{_ed ,_ =_cf .ParseInt (_db [0][8],10,32);}else {_ed =0;};if len (_db [0][9])> 0{_fdd ,_ =_cf .ParseInt (_db [0][9],10,32);}else {_fdd =0;
};_aa :=int (_ed *60*60+_fdd *60);switch _dg {case '-':_aa =-_aa ;case 'Z':_aa =0;};_dd :=_ff .Sprintf ("\u0055\u0054\u0043\u0025\u0063\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064",_dg ,_ed ,_fdd );_dab :=_cb .FixedZone (_dd ,_aa );return _cb .Date (int (_bf ),_cb .Month (_da ),int (_e ),int (_fb ),int (_ae ),int (_fbg ),0,_dab ),nil ;
};func FormatPdfTime (in _cb .Time )string {_a :=in .Format ("\u002d\u0030\u0037\u003a\u0030\u0030");_d ,_ :=_cf .ParseInt (_a [1:3],10,32);_fg ,_ :=_cf .ParseInt (_a [4:6],10,32);_gg :=int64 (in .Year ());_cbf :=int64 (in .Month ());_cg :=int64 (in .Day ());
_fgb :=int64 (in .Hour ());_fd :=int64 (in .Minute ());_ge :=int64 (in .Second ());_de :=_a [0];return _ff .Sprintf ("\u0044\u003a\u0025\u002e\u0034\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e\u0032\u0064\u0025\u002e2\u0064\u0025\u0063\u0025\u002e2\u0064\u0027%\u002e\u0032\u0064\u0027",_gg ,_cbf ,_cg ,_fgb ,_fd ,_ge ,_de ,_d ,_fg );
};var _ab =_b .MustCompile ("\u005c\u0073\u002a\u0044\u005c\u0073\u002a:\u005c\u0073\u002a\u0028\u005c\u0064\u007b\u0034\u007d\u0029\u0028\u005c\u0064\u007b2\u007d)\u0028\u005c\u0064\u007b\u0032\u007d)\u0028\u005c\u0064\u007b\u0032\u007d\u0029(\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005c\u0064\u007b\u0032\u007d\u0029\u0028\u005b\u002b\u002d\u005a\u005d\u0029\u003f\u0028\u005cd\u007b\u0032\u007d\u0029\u003f\u0027\u003f\u0028\u005c\u0064\u007b\u0032\u007d)\u003f");
