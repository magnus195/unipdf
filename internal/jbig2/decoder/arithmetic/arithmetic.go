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

package arithmetic ;import (_a "fmt";_gc "github.com/unidoc/unipdf/v3/common";_ga "github.com/unidoc/unipdf/v3/internal/bitwise";_ag "github.com/unidoc/unipdf/v3/internal/jbig2/internal";_f "io";_gd "strings";);func (_cfc *Decoder )renormalize ()error {for {if _cfc ._ab ==0{if _aa :=_cfc .readByte ();
_aa !=nil {return _aa ;};};_cfc ._ac <<=1;_cfc ._e <<=1;_cfc ._ab --;if (_cfc ._ac &0x8000)!=0{break ;};};_cfc ._e &=0xffffffff;return nil ;};func (_eg *Decoder )readByte ()error {if _eg ._b .AbsolutePosition ()> _eg ._ce {if _ ,_gg :=_eg ._b .Seek (-1,_f .SeekCurrent );
_gg !=nil {return _gg ;};};_ead ,_gab :=_eg ._b .ReadByte ();if _gab !=nil {return _gab ;};_eg ._d =_ead ;if _eg ._d ==0xFF{_gdg ,_db :=_eg ._b .ReadByte ();if _db !=nil {return _db ;};if _gdg > 0x8F{_eg ._e +=0xFF00;_eg ._ab =8;if _ ,_age :=_eg ._b .Seek (-2,_f .SeekCurrent );
_age !=nil {return _age ;};}else {_eg ._e +=uint64 (_gdg )<<9;_eg ._ab =7;};}else {_ead ,_gab =_eg ._b .ReadByte ();if _gab !=nil {return _gab ;};_eg ._d =_ead ;_eg ._e +=uint64 (_eg ._d )<<8;_eg ._ab =8;};_eg ._e &=0xFFFFFFFFFF;return nil ;};func (_bg *DecoderStats )Copy ()*DecoderStats {_ebb :=&DecoderStats {_bef :_bg ._bef ,_ec :make ([]byte ,_bg ._bef )};
copy (_ebb ._ec ,_bg ._ec );return _ebb ;};func (_ef *Decoder )decodeIntBit (_dbc *DecoderStats )(int ,error ){_dbc .SetIndex (int32 (_ef ._bb ));_ad ,_ffb :=_ef .DecodeBit (_dbc );if _ffb !=nil {_gc .Log .Debug ("\u0041\u0072\u0069\u0074\u0068\u006d\u0065t\u0069\u0063\u0044e\u0063\u006f\u0064e\u0072\u0020'\u0064\u0065\u0063\u006f\u0064\u0065I\u006etB\u0069\u0074\u0027\u002d\u003e\u0020\u0044\u0065\u0063\u006f\u0064\u0065\u0042\u0069\u0074\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u002e\u0020\u0025\u0076",_ffb );
return _ad ,_ffb ;};if _ef ._bb < 256{_ef ._bb =((_ef ._bb <<uint64 (1))|int64 (_ad ))&0x1ff;}else {_ef ._bb =(((_ef ._bb <<uint64 (1)|int64 (_ad ))&511)|256)&0x1ff;};return _ad ,nil ;};func (_ebc *DecoderStats )SetIndex (index int32 ){_ebc ._egf =index };
func New (r *_ga .Reader )(*Decoder ,error ){_ea :=&Decoder {_b :r ,ContextSize :[]uint32 {16,13,10,10},ReferedToContextSize :[]uint32 {13,10}};if _fb :=_ea .init ();_fb !=nil {return nil ,_fb ;};return _ea ,nil ;};func (_agd *Decoder )init ()error {_agd ._ce =_agd ._b .AbsolutePosition ();
_cgb ,_gac :=_agd ._b .ReadByte ();if _gac !=nil {_gc .Log .Debug ("B\u0075\u0066\u0066\u0065\u0072\u0030 \u0072\u0065\u0061\u0064\u0042\u0079\u0074\u0065\u0020f\u0061\u0069\u006ce\u0064.\u0020\u0025\u0076",_gac );return _gac ;};_agd ._d =_cgb ;_agd ._e =uint64 (_cgb )<<16;
if _gac =_agd .readByte ();_gac !=nil {return _gac ;};_agd ._e <<=7;_agd ._ab -=7;_agd ._ac =0x8000;_agd ._cc ++;return nil ;};func (_gga *DecoderStats )cx ()byte {return _gga ._ec [_gga ._egf ]};func (_gfe *Decoder )lpsExchange (_ed *DecoderStats ,_fg int32 ,_efd uint32 )int {_bfg :=_ed .getMps ();
if _gfe ._ac < _efd {_ed .setEntry (int (_c [_fg ][1]));_gfe ._ac =_efd ;return int (_bfg );};if _c [_fg ][3]==1{_ed .toggleMps ();};_ed .setEntry (int (_c [_fg ][2]));_gfe ._ac =_efd ;return int (1-_bfg );};func (_bd *DecoderStats )Reset (){for _cae :=0;
_cae < len (_bd ._ec );_cae ++{_bd ._ec [_cae ]=0;_bd ._dg [_cae ]=0;};};func (_cg *Decoder )DecodeIAID (codeLen uint64 ,stats *DecoderStats )(int64 ,error ){_cg ._bb =1;var _ee uint64 ;for _ee =0;_ee < codeLen ;_ee ++{stats .SetIndex (int32 (_cg ._bb ));
_aca ,_gdb :=_cg .DecodeBit (stats );if _gdb !=nil {return 0,_gdb ;};_cg ._bb =(_cg ._bb <<1)|int64 (_aca );};_cf :=_cg ._bb -(1<<codeLen );return _cf ,nil ;};func (_eb *Decoder )DecodeBit (stats *DecoderStats )(int ,error ){var (_da int ;_bf =_c [stats .cx ()][0];
_ge =int32 (stats .cx ()););defer func (){_eb ._cc ++}();_eb ._ac -=_bf ;if (_eb ._e >>16)< uint64 (_bf ){_da =_eb .lpsExchange (stats ,_ge ,_bf );if _gb :=_eb .renormalize ();_gb !=nil {return 0,_gb ;};}else {_eb ._e -=uint64 (_bf )<<16;if (_eb ._ac &0x8000)==0{_da =_eb .mpsExchange (stats ,_ge );
if _gde :=_eb .renormalize ();_gde !=nil {return 0,_gde ;};}else {_da =int (stats .getMps ());};};return _da ,nil ;};func (_abf *DecoderStats )Overwrite (dNew *DecoderStats ){for _gag :=0;_gag < len (_abf ._ec );_gag ++{_abf ._ec [_gag ]=dNew ._ec [_gag ];
_abf ._dg [_gag ]=dNew ._dg [_gag ];};};func (_ca *Decoder )mpsExchange (_fd *DecoderStats ,_gbc int32 )int {_aae :=_fd ._dg [_fd ._egf ];if _ca ._ac < _c [_gbc ][0]{if _c [_gbc ][3]==1{_fd .toggleMps ();};_fd .setEntry (int (_c [_gbc ][2]));return int (1-_aae );
};_fd .setEntry (int (_c [_gbc ][1]));return int (_aae );};func (_fgf *DecoderStats )String ()string {_aee :=&_gd .Builder {};_aee .WriteString (_a .Sprintf ("S\u0074\u0061\u0074\u0073\u003a\u0020\u0020\u0025\u0064\u000a",len (_fgf ._ec )));for _fgg ,_fa :=range _fgf ._ec {if _fa !=0{_aee .WriteString (_a .Sprintf ("N\u006f\u0074\u0020\u007aer\u006f \u0061\u0074\u003a\u0020\u0025d\u0020\u002d\u0020\u0025\u0064\u000a",_fgg ,_fa ));
};};return _aee .String ();};func (_efa *DecoderStats )setEntry (_dge int ){_abfe :=byte (_dge &0x7f);_efa ._ec [_efa ._egf ]=_abfe };func NewStats (contextSize int32 ,index int32 )*DecoderStats {return &DecoderStats {_egf :index ,_bef :contextSize ,_ec :make ([]byte ,contextSize ),_dg :make ([]byte ,contextSize )};
};type DecoderStats struct{_egf int32 ;_bef int32 ;_ec []byte ;_dg []byte ;};func (_eed *DecoderStats )getMps ()byte {return _eed ._dg [_eed ._egf ]};var (_c =[][4]uint32 {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};
);type Decoder struct{ContextSize []uint32 ;ReferedToContextSize []uint32 ;_b *_ga .Reader ;_d uint8 ;_e uint64 ;_ac uint32 ;_bb int64 ;_ab int32 ;_cc int32 ;_ce int64 ;};func (_ccd *DecoderStats )toggleMps (){_ccd ._dg [_ccd ._egf ]^=1};func (_gf *Decoder )DecodeInt (stats *DecoderStats )(int32 ,error ){var (_gbe ,_ff int32 ;
_dab ,_fe ,_be int ;_ae error ;);if stats ==nil {stats =NewStats (512,1);};_gf ._bb =1;_fe ,_ae =_gf .decodeIntBit (stats );if _ae !=nil {return 0,_ae ;};_dab ,_ae =_gf .decodeIntBit (stats );if _ae !=nil {return 0,_ae ;};if _dab ==1{_dab ,_ae =_gf .decodeIntBit (stats );
if _ae !=nil {return 0,_ae ;};if _dab ==1{_dab ,_ae =_gf .decodeIntBit (stats );if _ae !=nil {return 0,_ae ;};if _dab ==1{_dab ,_ae =_gf .decodeIntBit (stats );if _ae !=nil {return 0,_ae ;};if _dab ==1{_dab ,_ae =_gf .decodeIntBit (stats );if _ae !=nil {return 0,_ae ;
};if _dab ==1{_be =32;_ff =4436;}else {_be =12;_ff =340;};}else {_be =8;_ff =84;};}else {_be =6;_ff =20;};}else {_be =4;_ff =4;};}else {_be =2;_ff =0;};for _ged :=0;_ged < _be ;_ged ++{_dab ,_ae =_gf .decodeIntBit (stats );if _ae !=nil {return 0,_ae ;};
_gbe =(_gbe <<1)|int32 (_dab );};_gbe +=_ff ;if _fe ==0{return _gbe ,nil ;}else if _fe ==1&&_gbe > 0{return -_gbe ,nil ;};return 0,_ag .ErrOOB ;};