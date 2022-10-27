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

package decoder ;import (_c "github.com/unidoc/unipdf/v3/internal/bitwise";_e "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_d "github.com/unidoc/unipdf/v3/internal/jbig2/document";_eg "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_a "image";
);func (_ea *Decoder )DecodePageImage (pageNumber int )(_a .Image ,error ){const _b ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";_bd ,_g :=_ea .decodePageImage (pageNumber );
if _g !=nil {return nil ,_eg .Wrap (_g ,_b ,"");};return _bd ,nil ;};type Decoder struct{_dd _c .StreamReader ;_cd *_d .Document ;_de int ;_db Parameters ;};func Decode (input []byte ,parameters Parameters ,globals *_d .Globals )(*Decoder ,error ){_bf :=_c .NewReader (input );
_ef ,_gf :=_d .DecodeDocument (_bf ,globals );if _gf !=nil {return nil ,_gf ;};return &Decoder {_dd :_bf ,_cd :_ef ,_db :parameters },nil ;};func (_fb *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _fb .decodePage (pageNumber )};func (_cf *Decoder )DecodeNextPage ()([]byte ,error ){_cf ._de ++;
_fg :=_cf ._de ;return _cf .decodePage (_fg );};type Parameters struct{UnpaddedData bool ;Color _e .Color ;};func (_ae *Decoder )decodePage (_df int )([]byte ,error ){const _gad ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _df < 0{return nil ,_eg .Errorf (_gad ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_df );
};if _df > int (_ae ._cd .NumberOfPages ){return nil ,_eg .Errorf (_gad ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_df );
};_ge ,_dc :=_ae ._cd .GetPage (_df );if _dc !=nil {return nil ,_eg .Wrap (_dc ,_gad ,"");};_ed ,_dc :=_ge .GetBitmap ();if _dc !=nil {return nil ,_eg .Wrap (_dc ,_gad ,"");};_ed .InverseData ();if !_ae ._db .UnpaddedData {return _ed .Data ,nil ;};return _ed .GetUnpaddedData ();
};func (_fac *Decoder )decodePageImage (_dcc int )(_a .Image ,error ){const _ac ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _dcc < 0{return nil ,_eg .Errorf (_ac ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_dcc );
};if _dcc > int (_fac ._cd .NumberOfPages ){return nil ,_eg .Errorf (_ac ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_dcc );
};_dbc ,_ege :=_fac ._cd .GetPage (_dcc );if _ege !=nil {return nil ,_eg .Wrap (_ege ,_ac ,"");};_ace ,_ege :=_dbc .GetBitmap ();if _ege !=nil {return nil ,_eg .Wrap (_ege ,_ac ,"");};_ace .InverseData ();return _ace .ToImage (),nil ;};func (_ga *Decoder )PageNumber ()(int ,error ){const _fa ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";
if _ga ._cd ==nil {return 0,_eg .Error (_fa ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");};return int (_ga ._cd .NumberOfPages ),nil ;};