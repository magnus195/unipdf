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

package sampling ;import (_c "github.com/unidoc/unipdf/v3/internal/bitwise";_fa "github.com/unidoc/unipdf/v3/internal/imageutil";_g "io";);type Writer struct{_ccc _fa .ImageBase ;_cee *_c .Writer ;_cg ,_cfg int ;_afd bool ;};func (_faf *Reader )ReadSamples (samples []uint32 )(_gac error ){for _d :=0;
_d < len (samples );_d ++{samples [_d ],_gac =_faf .ReadSample ();if _gac !=nil {return _gac ;};};return nil ;};type SampleWriter interface{WriteSample (_fb uint32 )error ;WriteSamples (_gf []uint32 )error ;};type SampleReader interface{ReadSample ()(uint32 ,error );
ReadSamples (_a []uint32 )error ;};type Reader struct{_e _fa .ImageBase ;_gd *_c .Reader ;_b ,_ga ,_ac int ;_be bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _ceb []uint32 ;_bd :=bitsPerOutputSample ;
var _cc uint32 ;var _gga uint32 ;_gba :=0;_abg :=0;_bg :=0;for _bg < len (data ){if _gba > 0{_bdd :=_gba ;if _bd < _bdd {_bdd =_bd ;};_cc =(_cc <<uint (_bdd ))|(_gga >>uint (bitsPerInputSample -_bdd ));_gba -=_bdd ;if _gba > 0{_gga =_gga <<uint (_bdd );
}else {_gga =0;};_bd -=_bdd ;if _bd ==0{_ceb =append (_ceb ,_cc );_bd =bitsPerOutputSample ;_cc =0;_abg ++;};}else {_dc :=data [_bg ];_bg ++;_dd :=bitsPerInputSample ;if _bd < _dd {_dd =_bd ;};_gba =bitsPerInputSample -_dd ;_cc =(_cc <<uint (_dd ))|(_dc >>uint (_gba ));
if _dd < bitsPerInputSample {_gga =_dc <<uint (_dd );};_bd -=_dd ;if _bd ==0{_ceb =append (_ceb ,_cc );_bd =bitsPerOutputSample ;_cc =0;_abg ++;};};};for _gba >=bitsPerOutputSample {_af :=_gba ;if _bd < _af {_af =_bd ;};_cc =(_cc <<uint (_af ))|(_gga >>uint (bitsPerInputSample -_af ));
_gba -=_af ;if _gba > 0{_gga =_gga <<uint (_af );}else {_gga =0;};_bd -=_af ;if _bd ==0{_ceb =append (_ceb ,_cc );_bd =bitsPerOutputSample ;_cc =0;_abg ++;};};if _bd > 0&&_bd < bitsPerOutputSample {_cc <<=uint (_bd );_ceb =append (_ceb ,_cc );};return _ceb ;
};func (_gdb *Reader )ReadSample ()(uint32 ,error ){if _gdb ._ga ==_gdb ._e .Height {return 0,_g .EOF ;};_ab ,_bc :=_gdb ._gd .ReadBits (byte (_gdb ._e .BitsPerComponent ));if _bc !=nil {return 0,_bc ;};_gdb ._ac --;if _gdb ._ac ==0{_gdb ._ac =_gdb ._e .ColorComponents ;
_gdb ._b ++;};if _gdb ._b ==_gdb ._e .Width {if _gdb ._be {_gdb ._gd .ConsumeRemainingBits ();};_gdb ._b =0;_gdb ._ga ++;};return uint32 (_ab ),nil ;};func (_ff *Writer )WriteSample (sample uint32 )error {if _ ,_db :=_ff ._cee .WriteBits (uint64 (sample ),_ff ._ccc .BitsPerComponent );
_db !=nil {return _db ;};_ff ._cfg --;if _ff ._cfg ==0{_ff ._cfg =_ff ._ccc .ColorComponents ;_ff ._cg ++;};if _ff ._cg ==_ff ._ccc .Width {if _ff ._afd {_ff ._cee .FinishByte ();};_ff ._cg =0;};return nil ;};func (_dbe *Writer )WriteSamples (samples []uint32 )error {for _ba :=0;
_ba < len (samples );_ba ++{if _fe :=_dbe .WriteSample (samples [_ba ]);_fe !=nil {return _fe ;};};return nil ;};func NewWriter (img _fa .ImageBase )*Writer {return &Writer {_cee :_c .NewWriterMSB (img .Data ),_ccc :img ,_cfg :img .ColorComponents ,_afd :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _ce []uint32 ;_cf :=bitsPerSample ;var _ed uint32 ;var _gb byte ;_ae :=0;_gdc :=0;_de :=0;for _de < len (data ){if _ae > 0{_fag :=_ae ;if _cf < _fag {_fag =_cf ;};_ed =(_ed <<uint (_fag ))|uint32 (_gb >>uint (8-_fag ));
_ae -=_fag ;if _ae > 0{_gb =_gb <<uint (_fag );}else {_gb =0;};_cf -=_fag ;if _cf ==0{_ce =append (_ce ,_ed );_cf =bitsPerSample ;_ed =0;_gdc ++;};}else {_gg :=data [_de ];_de ++;_bf :=8;if _cf < _bf {_bf =_cf ;};_ae =8-_bf ;_ed =(_ed <<uint (_bf ))|uint32 (_gg >>uint (_ae ));
if _bf < 8{_gb =_gg <<uint (_bf );};_cf -=_bf ;if _cf ==0{_ce =append (_ce ,_ed );_cf =bitsPerSample ;_ed =0;_gdc ++;};};};for _ae >=bitsPerSample {_fd :=_ae ;if _cf < _fd {_fd =_cf ;};_ed =(_ed <<uint (_fd ))|uint32 (_gb >>uint (8-_fd ));_ae -=_fd ;if _ae > 0{_gb =_gb <<uint (_fd );
}else {_gb =0;};_cf -=_fd ;if _cf ==0{_ce =append (_ce ,_ed );_cf =bitsPerSample ;_ed =0;_gdc ++;};};return _ce ;};func NewReader (img _fa .ImageBase )*Reader {return &Reader {_gd :_c .NewReader (img .Data ),_e :img ,_ac :img .ColorComponents ,_be :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};