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

package decoder ;import (_d "github.com/unidoc/unipdf/v3/internal/bitwise";_da "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_f "github.com/unidoc/unipdf/v3/internal/jbig2/document";_c "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_b "image";
);type Parameters struct{UnpaddedData bool ;Color _da .Color ;};func (_af *Decoder )DecodePage (pageNumber int )([]byte ,error ){return _af .decodePage (pageNumber )};func (_g *Decoder )DecodeNextPage ()([]byte ,error ){_g ._fc ++;_fg :=_g ._fc ;return _g .decodePage (_fg );
};func (_ef *Decoder )PageNumber ()(int ,error ){const _ec ="\u0044e\u0063o\u0064\u0065\u0072\u002e\u0050a\u0067\u0065N\u0075\u006d\u0062\u0065\u0072";if _ef ._de ==nil {return 0,_c .Error (_ec ,"d\u0065\u0063\u006f\u0064\u0065\u0072 \u006e\u006f\u0074\u0020\u0069\u006e\u0069\u0074\u0069a\u006c\u0069\u007ae\u0064 \u0079\u0065\u0074");
};return int (_ef ._de .NumberOfPages ),nil ;};func Decode (input []byte ,parameters Parameters ,globals *_f .Globals )(*Decoder ,error ){_eg :=_d .NewReader (input );_ce ,_bfg :=_f .DecodeDocument (_eg ,globals );if _bfg !=nil {return nil ,_bfg ;};return &Decoder {_a :_eg ,_de :_ce ,_bg :parameters },nil ;
};func (_fd *Decoder )DecodePageImage (pageNumber int )(_b .Image ,error ){const _ad ="\u0064\u0065\u0063od\u0065\u0072\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";_df ,_ee :=_fd .decodePageImage (pageNumber );
if _ee !=nil {return nil ,_c .Wrap (_ee ,_ad ,"");};return _df ,nil ;};func (_bf *Decoder )decodePageImage (_ea int )(_b .Image ,error ){const _dec ="\u0064e\u0063o\u0064\u0065\u0050\u0061\u0067\u0065\u0049\u006d\u0061\u0067\u0065";if _ea < 0{return nil ,_c .Errorf (_dec ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_ea );
};if _ea > int (_bf ._de .NumberOfPages ){return nil ,_c .Errorf (_dec ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_ea );
};_gg ,_ed :=_bf ._de .GetPage (_ea );if _ed !=nil {return nil ,_c .Wrap (_ed ,_dec ,"");};_ac ,_ed :=_gg .GetBitmap ();if _ed !=nil {return nil ,_c .Wrap (_ed ,_dec ,"");};_ac .InverseData ();return _ac .ToImage (),nil ;};type Decoder struct{_a _d .StreamReader ;
_de *_f .Document ;_fc int ;_bg Parameters ;};func (_db *Decoder )decodePage (_gd int )([]byte ,error ){const _gc ="\u0064\u0065\u0063\u006f\u0064\u0065\u0050\u0061\u0067\u0065";if _gd < 0{return nil ,_c .Errorf (_gc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0070\u0061\u0067\u0065 \u006eu\u006db\u0065\u0072\u003a\u0020\u0027\u0025\u0064'",_gd );
};if _gd > int (_db ._de .NumberOfPages ){return nil ,_c .Errorf (_gc ,"p\u0061\u0067\u0065\u003a\u0020\u0027%\u0064\u0027\u0020\u006e\u006f\u0074 \u0066\u006f\u0075\u006e\u0064\u0020\u0069n\u0020\u0074\u0068\u0065\u0020\u0064\u0065\u0063\u006f\u0064e\u0072",_gd );
};_gdd ,_bb :=_db ._de .GetPage (_gd );if _bb !=nil {return nil ,_c .Wrap (_bb ,_gc ,"");};_fb ,_bb :=_gdd .GetBitmap ();if _bb !=nil {return nil ,_c .Wrap (_bb ,_gc ,"");};_fb .InverseData ();if !_db ._bg .UnpaddedData {return _fb .Data ,nil ;};return _fb .GetUnpaddedData ();
};