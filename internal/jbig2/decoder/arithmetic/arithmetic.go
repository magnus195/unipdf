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

package arithmetic ;import (_e "fmt";_ee "github.com/unidoc/unipdf/v3/common";_ed "github.com/unidoc/unipdf/v3/internal/bitwise";_a "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_g "io";_b "strings";);func (_eea *Decoder )renormalize ()error {for {if _eea ._dg ==0{if _bee :=_eea .readByte ();
_bee !=nil {return _bee ;};};_eea ._f <<=1;_eea ._cad <<=1;_eea ._dg --;if (_eea ._f &0x8000)!=0{break ;};};_eea ._cad &=0xffffffff;return nil ;};func (_ef *Decoder )DecodeBit (stats *DecoderStats )(int ,error ){var (_cc int ;_eg =_gb [stats .cx ()][0];
_gbe =int32 (stats .cx ()););defer func (){_ef ._cag ++}();_ef ._f -=_eg ;if (_ef ._cad >>16)< uint64 (_eg ){_cc =_ef .lpsExchange (stats ,_gbe ,_eg );if _de :=_ef .renormalize ();_de !=nil {return 0,_de ;};}else {_ef ._cad -=uint64 (_eg )<<16;if (_ef ._f &0x8000)==0{_cc =_ef .mpsExchange (stats ,_gbe );
if _eae :=_ef .renormalize ();_eae !=nil {return 0,_eae ;};}else {_cc =int (stats .getMps ());};};return _cc ,nil ;};type DecoderStats struct{_gcb int32 ;_ac int32 ;_ggd []byte ;_ebe []byte ;};func (_cac *Decoder )DecodeInt (stats *DecoderStats )(int32 ,error ){var (_dd ,_gbc int32 ;
_edc ,_af ,_fc int ;_cf error ;);if stats ==nil {stats =NewStats (512,1);};_cac ._ba =1;_af ,_cf =_cac .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};_edc ,_cf =_cac .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _edc ==1{_edc ,_cf =_cac .decodeIntBit (stats );
if _cf !=nil {return 0,_cf ;};if _edc ==1{_edc ,_cf =_cac .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _edc ==1{_edc ,_cf =_cac .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;};if _edc ==1{_edc ,_cf =_cac .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;
};if _edc ==1{_fc =32;_gbc =4436;}else {_fc =12;_gbc =340;};}else {_fc =8;_gbc =84;};}else {_fc =6;_gbc =20;};}else {_fc =4;_gbc =4;};}else {_fc =2;_gbc =0;};for _afa :=0;_afa < _fc ;_afa ++{_edc ,_cf =_cac .decodeIntBit (stats );if _cf !=nil {return 0,_cf ;
};_dd =(_dd <<1)|int32 (_edc );};_dd +=_gbc ;if _af ==0{return _dd ,nil ;}else if _af ==1&&_dd > 0{return -_dd ,nil ;};return 0,_a .ErrOOB ;};func (_cfa *Decoder )mpsExchange (_gge *DecoderStats ,_da int32 )int {_fea :=_gge ._ebe [_gge ._gcb ];if _cfa ._f < _gb [_da ][0]{if _gb [_da ][3]==1{_gge .toggleMps ();
};_gge .setEntry (int (_gb [_da ][2]));return int (1-_fea );};_gge .setEntry (int (_gb [_da ][1]));return int (_fea );};func (_eff *DecoderStats )getMps ()byte {return _eff ._ebe [_eff ._gcb ]};func (_gc *Decoder )init ()error {_gc ._bab =_gc ._c .StreamPosition ();
_fdc ,_ff :=_gc ._c .ReadByte ();if _ff !=nil {_ee .Log .Debug ("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076",_ff );return _ff ;};_gc ._ca =_fdc ;_gc ._cad =uint64 (_fdc )<<16;
if _ff =_gc .readByte ();_ff !=nil {return _ff ;};_gc ._cad <<=7;_gc ._dg -=7;_gc ._f =0x8000;_gc ._cag ++;return nil ;};func (_gbcf *DecoderStats )Overwrite (dNew *DecoderStats ){for _daa :=0;_daa < len (_gbcf ._ggd );_daa ++{_gbcf ._ggd [_daa ]=dNew ._ggd [_daa ];
_gbcf ._ebe [_daa ]=dNew ._ebe [_daa ];};};var (_gb =[][4]uint32 {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
);func (_ebc *DecoderStats )setEntry (_ffc int ){_bg :=byte (_ffc &0x7f);_ebc ._ggd [_ebc ._gcb ]=_bg };func (_caag *DecoderStats )SetIndex (index int32 ){_caag ._gcb =index };func NewStats (contextSize int32 ,index int32 )*DecoderStats {return &DecoderStats {_gcb :index ,_ac :contextSize ,_ggd :make ([]byte ,contextSize ),_ebe :make ([]byte ,contextSize )};
};func (_cgb *Decoder )decodeIntBit (_aa *DecoderStats )(int ,error ){_aa .SetIndex (int32 (_cgb ._ba ));_fe ,_fg :=_cgb .DecodeBit (_aa );if _fg !=nil {_ee .Log .Debug ("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076",_fg );
return _fe ,_fg ;};if _cgb ._ba < 256{_cgb ._ba =((_cgb ._ba <<uint64 (1))|int64 (_fe ))&0x1ff;}else {_cgb ._ba =(((_cgb ._ba <<uint64 (1)|int64 (_fe ))&511)|256)&0x1ff;};return _fe ,nil ;};func (_eaa *Decoder )lpsExchange (_ab *DecoderStats ,_fcb int32 ,_afe uint32 )int {_fcd :=_ab .getMps ();
if _eaa ._f < _afe {_ab .setEntry (int (_gb [_fcb ][1]));_eaa ._f =_afe ;return int (_fcd );};if _gb [_fcb ][3]==1{_ab .toggleMps ();};_ab .setEntry (int (_gb [_fcb ][2]));_eaa ._f =_afe ;return int (1-_fcd );};func New (r _ed .StreamReader )(*Decoder ,error ){_ea :=&Decoder {_c :r ,ContextSize :[]uint32 {16,13,10,10},ReferedToContextSize :[]uint32 {13,10}};
if _eb :=_ea .init ();_eb !=nil {return nil ,_eb ;};return _ea ,nil ;};func (_bc *DecoderStats )toggleMps (){_bc ._ebe [_bc ._gcb ]^=1};func (_caa *DecoderStats )Copy ()*DecoderStats {_ede :=&DecoderStats {_ac :_caa ._ac ,_ggd :make ([]byte ,_caa ._ac )};
for _babg :=0;_babg < len (_caa ._ggd );_babg ++{_ede ._ggd [_babg ]=_caa ._ggd [_babg ];};return _ede ;};func (_cgce *DecoderStats )Reset (){for _egc :=0;_egc < len (_cgce ._ggd );_egc ++{_cgce ._ggd [_egc ]=0;_cgce ._ebe [_egc ]=0;};};func (_fa *Decoder )readByte ()error {if _fa ._c .StreamPosition ()> _fa ._bab {if _ ,_cce :=_fa ._c .Seek (-1,_g .SeekCurrent );
_cce !=nil {return _cce ;};};_bd ,_cgc :=_fa ._c .ReadByte ();if _cgc !=nil {return _cgc ;};_fa ._ca =_bd ;if _fa ._ca ==0xFF{_aeb ,_gg :=_fa ._c .ReadByte ();if _gg !=nil {return _gg ;};if _aeb > 0x8F{_fa ._cad +=0xFF00;_fa ._dg =8;if _ ,_ce :=_fa ._c .Seek (-2,_g .SeekCurrent );
_ce !=nil {return _ce ;};}else {_fa ._cad +=uint64 (_aeb )<<9;_fa ._dg =7;};}else {_bd ,_cgc =_fa ._c .ReadByte ();if _cgc !=nil {return _cgc ;};_fa ._ca =_bd ;_fa ._cad +=uint64 (_fa ._ca )<<8;_fa ._dg =8;};_fa ._cad &=0xFFFFFFFFFF;return nil ;};func (_fd *Decoder )DecodeIAID (codeLen uint64 ,stats *DecoderStats )(int64 ,error ){_fd ._ba =1;
var _ae uint64 ;for _ae =0;_ae < codeLen ;_ae ++{stats .SetIndex (int32 (_fd ._ba ));_eag ,_fb :=_fd .DecodeBit (stats );if _fb !=nil {return 0,_fb ;};_fd ._ba =(_fd ._ba <<1)|int64 (_eag );};_cg :=_fd ._ba -(1<<codeLen );return _cg ,nil ;};type Decoder struct{ContextSize []uint32 ;
ReferedToContextSize []uint32 ;_c _ed .StreamReader ;_ca uint8 ;_cad uint64 ;_f uint32 ;_ba int64 ;_dg int32 ;_cag int32 ;_bab int64 ;};func (_ad *DecoderStats )cx ()byte {return _ad ._ggd [_ad ._gcb ]};func (_eac *DecoderStats )String ()string {_cd :=&_b .Builder {};
_cd .WriteString (_e .Sprintf ("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a",len (_eac ._ggd )));for _ebed ,_cgd :=range _eac ._ggd {if _cgd !=0{_cd .WriteString (_e .Sprintf ("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a",_ebed ,_cgd ));
};};return _cd .String ();};