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

package bitwise ;import (_ce "encoding/binary";_ac "errors";_e "fmt";
	_f "github.com/magnus195/unipdf/v3/common";
	_ef "github.com/magnus195/unipdf/v3/internal/jbig2/errors";_a "io";);func (_de *BufferedWriter )Write (d []byte )(int ,error ){_de .expandIfNeeded (len (d ));
if _de ._ffg ==0{return _de .writeFullBytes (d ),nil ;};return _de .writeShiftedBytes (d ),nil ;};func (_gbd *Reader )Length ()uint64 {return uint64 (_gbd ._af ._deg )};var _ _a .ByteWriter =&BufferedWriter {};func (_fde *Writer )FinishByte (){if _fde ._egb ==0{return ;
};_fde ._egb =0;_fde ._cdb ++;};func (_bgb *Reader )ReadBits (n byte )(_ccc uint64 ,_afa error ){if n < _bgb ._ade {_gde :=_bgb ._ade -n ;_ccc =uint64 (_bgb ._dge >>_gde );_bgb ._dge &=1<<_gde -1;_bgb ._ade =_gde ;return _ccc ,nil ;};if n > _bgb ._ade {if _bgb ._ade > 0{_ccc =uint64 (_bgb ._dge );
n -=_bgb ._ade ;};for n >=8{_cgd ,_aef :=_bgb .readBufferByte ();if _aef !=nil {return 0,_aef ;};_ccc =_ccc <<8+uint64 (_cgd );n -=8;};if n > 0{if _bgb ._dge ,_afa =_bgb .readBufferByte ();_afa !=nil {return 0,_afa ;};_cbc :=8-n ;_ccc =_ccc <<n +uint64 (_bgb ._dge >>_cbc );
_bgb ._dge &=1<<_cbc -1;_bgb ._ade =_cbc ;}else {_bgb ._ade =0;};return _ccc ,nil ;};_bgb ._ade =0;return uint64 (_bgb ._dge ),nil ;};func (_dcg *Reader )Mark (){_dcg ._afg =_dcg ._df ;_dcg ._gdcd =_dcg ._ade ;_dcg ._ae =_dcg ._dge ;_dcg ._bbda =_dcg ._bdc ;
};func (_bcf *Reader )Align ()(_gbb byte ){_gbb =_bcf ._ade ;_bcf ._ade =0;return _gbb };func (_cfd *Reader )read (_geg []byte )(int ,error ){if _cfd ._df >=int64 (_cfd ._af ._deg ){return 0,_a .EOF ;};_cfd ._ag =-1;_ca :=copy (_geg ,_cfd ._af ._efb [(int64 (_cfd ._af ._cb )+_cfd ._df ):(_cfd ._af ._cb +_cfd ._af ._deg )]);
_cfd ._df +=int64 (_ca );return _ca ,nil ;};func (_fgf *Writer )writeByte (_eac byte )error {if _fgf ._cdb > len (_fgf ._gfd )-1{return _a .EOF ;};if _fgf ._cdb ==len (_fgf ._gfd )-1&&_fgf ._egb !=0{return _a .EOF ;};if _fgf ._egb ==0{_fgf ._gfd [_fgf ._cdb ]=_eac ;
_fgf ._cdb ++;return nil ;};if _fgf ._gad {_fgf ._gfd [_fgf ._cdb ]|=_eac >>_fgf ._egb ;_fgf ._cdb ++;_fgf ._gfd [_fgf ._cdb ]=byte (uint16 (_eac )<<(8-_fgf ._egb )&0xff);}else {_fgf ._gfd [_fgf ._cdb ]|=byte (uint16 (_eac )<<_fgf ._egb &0xff);_fgf ._cdb ++;
_fgf ._gfd [_fgf ._cdb ]=_eac >>(8-_fgf ._egb );};return nil ;};var _ _a .Writer =&BufferedWriter {};func (_dad *BufferedWriter )grow (_acce int ){if _dad ._d ==nil &&_acce < _cd {_dad ._d =make ([]byte ,_acce ,_cd );return ;};_ecd :=len (_dad ._d );if _dad ._ffg !=0{_ecd ++;
};_fa :=cap (_dad ._d );switch {case _acce <=_fa /2-_ecd :_f .Log .Trace ("\u005b\u0042\u0075\u0066\u0066\u0065r\u0065\u0064\u0057\u0072\u0069t\u0065\u0072\u005d\u0020\u0067\u0072o\u0077\u0020\u002d\u0020\u0072e\u0073\u006c\u0069\u0063\u0065\u0020\u006f\u006e\u006c\u0079\u002e\u0020L\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0043\u0061\u0070\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u006e\u003a\u0020'\u0025\u0064\u0027",len (_dad ._d ),cap (_dad ._d ),_acce );
_f .Log .Trace ("\u0020\u006e\u0020\u003c\u003d\u0020\u0063\u0020\u002f\u0020\u0032\u0020\u002d\u006d\u002e \u0043:\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u006d\u003a\u0020\u0027\u0025\u0064\u0027",_fa ,_ecd );copy (_dad ._d ,_dad ._d [_dad .fullOffset ():]);
case _fa > _cea -_fa -_acce :_f .Log .Error ("\u0042\u0055F\u0046\u0045\u0052 \u0074\u006f\u006f\u0020\u006c\u0061\u0072\u0067\u0065");return ;default:_dc :=make ([]byte ,2*_fa +_acce );copy (_dc ,_dad ._d );_dad ._d =_dc ;};_dad ._d =_dad ._d [:_ecd +_acce ];
};func (_ffc *BufferedWriter )WriteBits (bits uint64 ,number int )(_gcf int ,_bb error ){const _eg ="\u0042u\u0066\u0066\u0065\u0072e\u0064\u0057\u0072\u0069\u0074e\u0072.\u0057r\u0069\u0074\u0065\u0072\u0042\u0069\u0074s";if number < 0||number > 64{return 0,_ef .Errorf (_eg ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};_dd :=number /8;if _dd > 0{_bc :=number -_dd *8;for _efg :=_dd -1;_efg >=0;_efg --{_ec :=byte ((bits >>uint (_efg *8+_bc ))&0xff);if _bb =_ffc .WriteByte (_ec );_bb !=nil {return _gcf ,_ef .Wrapf (_bb ,_eg ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_dd -_efg +1);
};};number -=_dd *8;if number ==0{return _dd ,nil ;};};var _da int ;for _ga :=0;_ga < number ;_ga ++{if _ffc ._aa {_da =int ((bits >>uint (number -1-_ga ))&0x1);}else {_da =int (bits &0x1);bits >>=1;};if _bb =_ffc .WriteBit (_da );_bb !=nil {return _gcf ,_ef .Wrapf (_bb ,_eg ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_ga );
};};return _dd ,nil ;};func (_ad *BufferedWriter )byteCapacity ()int {_eb :=len (_ad ._d )-_ad ._b ;if _ad ._ffg !=0{_eb --;};return _eb ;};func (_fg *Reader )readUnalignedByte ()(_agd byte ,_dae error ){_fdc :=_fg ._ade ;_agd =_fg ._dge <<(8-_fdc );_fg ._dge ,_dae =_fg .readBufferByte ();
if _dae !=nil {return 0,_dae ;};_agd |=_fg ._dge >>_fdc ;_fg ._dge &=1<<_fdc -1;return _agd ,nil ;};type Writer struct{_gfd []byte ;_egb uint8 ;_cdb int ;_gad bool ;};var _ BinaryWriter =&BufferedWriter {};func (_fbf *Writer )Data ()[]byte {return _fbf ._gfd };
func (_ed *BufferedWriter )fullOffset ()int {_bbc :=_ed ._b ;if _ed ._ffg !=0{_bbc ++;};return _bbc ;};func (_cbe *Reader )AbsolutePosition ()int64 {return _cbe ._df +int64 (_cbe ._af ._cb )};func (_dgf *Reader )BitPosition ()int {return int (_dgf ._ade )};
func NewReader (data []byte )*Reader {return &Reader {_af :readerSource {_efb :data ,_deg :len (data ),_cb :0}};};func (_gcd *BufferedWriter )writeFullBytes (_abe []byte )int {_gdg :=copy (_gcd ._d [_gcd .fullOffset ():],_abe );_gcd ._b +=_gdg ;return _gdg ;
};func (_dec *Writer )writeBit (_dfc uint8 )error {if len (_dec ._gfd )-1< _dec ._cdb {return _a .EOF ;};_gaa :=_dec ._egb ;if _dec ._gad {_gaa =7-_dec ._egb ;};_dec ._gfd [_dec ._cdb ]|=byte (uint16 (_dfc <<_gaa )&0xff);_dec ._egb ++;if _dec ._egb ==8{_dec ._cdb ++;
_dec ._egb =0;};return nil ;};type StreamReader interface{_a .Reader ;_a .ByteReader ;_a .Seeker ;Align ()byte ;BitPosition ()int ;Mark ();Length ()uint64 ;ReadBit ()(int ,error );ReadBits (_acg byte )(uint64 ,error );ReadBool ()(bool ,error );ReadUint32 ()(uint32 ,error );
Reset ();AbsolutePosition ()int64 ;};func (_cdc *BufferedWriter )SkipBits (skip int )error {if skip ==0{return nil ;};_cc :=int (_cdc ._ffg )+skip ;if _cc >=0&&_cc < 8{_cdc ._ffg =uint8 (_cc );return nil ;};_cc =int (_cdc ._ffg )+_cdc ._b *8+skip ;if _cc < 0{return _ef .Errorf ("\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073","\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_acc :=_cc /8;_cg :=_cc %8;_cdc ._ffg =uint8 (_cg );if _gc :=_acc -_cdc ._b ;_gc > 0&&len (_cdc ._d )-1< _acc {if _cdc ._ffg !=0{_gc ++;};_cdc .expandIfNeeded (_gc );};_cdc ._b =_acc ;return nil ;};func (_dac *Reader )ConsumeRemainingBits ()(uint64 ,error ){if _dac ._ade !=0{return _dac .ReadBits (_dac ._ade );
};return 0,nil ;};func (_eff *BufferedWriter )Reset (){_eff ._d =_eff ._d [:0];_eff ._b =0;_eff ._ffg =0};func (_afc *Writer )WriteBit (bit int )error {switch bit {case 0,1:return _afc .writeBit (uint8 (bit ));};return _ef .Error ("\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0062\u0069\u0074\u0020v\u0061\u006c\u0075\u0065");
};func (_fbb *Reader )RelativePosition ()int64 {return _fbb ._df };func (_bg *BufferedWriter )writeShiftedBytes (_bd []byte )int {for _ ,_ge :=range _bd {_bg .writeByte (_ge );};return len (_bd );};func (_ab *BufferedWriter )FinishByte (){if _ab ._ffg ==0{return ;
};_ab ._ffg =0;_ab ._b ++;};func (_ea *BufferedWriter )tryGrowByReslice (_bbd int )bool {if _fc :=len (_ea ._d );_bbd <=cap (_ea ._d )-_fc {_ea ._d =_ea ._d [:_fc +_bbd ];return true ;};return false ;};var _ BinaryWriter =&Writer {};func (_aaf *Reader )Read (p []byte )(_ba int ,_gce error ){if _aaf ._ade ==0{return _aaf .read (p );
};for ;_ba < len (p );_ba ++{if p [_ba ],_gce =_aaf .readUnalignedByte ();_gce !=nil {return 0,_gce ;};};return _ba ,nil ;};func (_ace *BufferedWriter )WriteByte (bt byte )error {if _ace ._b > len (_ace ._d )-1||(_ace ._b ==len (_ace ._d )-1&&_ace ._ffg !=0){_ace .expandIfNeeded (1);
};_ace .writeByte (bt );return nil ;};func NewWriterMSB (data []byte )*Writer {return &Writer {_gfd :data ,_gad :true }};func (_bf *BufferedWriter )Data ()[]byte {return _bf ._d };func (_g *BufferedWriter )ResetBitIndex (){_g ._ffg =0};func (_aad *Writer )WriteByte (c byte )error {return _aad .writeByte (c )};
func (_bbf *Writer )ResetBit (){_bbf ._egb =0};type BitWriter interface{WriteBit (_eea int )error ;WriteBits (_be uint64 ,_eaf int )(_gf int ,_gdc error );FinishByte ();SkipBits (_cef int )error ;};func (_gdga *Writer )WriteBits (bits uint64 ,number int )(_fe int ,_fdf error ){const _eecd ="\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065r\u0042\u0069\u0074\u0073";
if number < 0||number > 64{return 0,_ef .Errorf (_eecd ,"\u0062i\u0074\u0073 \u006e\u0075\u006db\u0065\u0072\u0020\u006d\u0075\u0073\u0074 \u0062\u0065\u0020\u0069\u006e\u0020r\u0061\u006e\u0067\u0065\u0020\u003c\u0030\u002c\u0036\u0034\u003e,\u0020\u0069\u0073\u003a\u0020\u0027\u0025\u0064\u0027",number );
};if number ==0{return 0,nil ;};_afad :=number /8;if _afad > 0{_gdf :=number -_afad *8;for _gfc :=_afad -1;_gfc >=0;_gfc --{_gcde :=byte ((bits >>uint (_gfc *8+_gdf ))&0xff);if _fdf =_gdga .WriteByte (_gcde );_fdf !=nil {return _fe ,_ef .Wrapf (_fdf ,_eecd ,"\u0062\u0079\u0074\u0065\u003a\u0020\u0027\u0025\u0064\u0027",_afad -_gfc +1);
};};number -=_afad *8;if number ==0{return _afad ,nil ;};};var _bff int ;for _aba :=0;_aba < number ;_aba ++{if _gdga ._gad {_bff =int ((bits >>uint (number -1-_aba ))&0x1);}else {_bff =int (bits &0x1);bits >>=1;};if _fdf =_gdga .WriteBit (_bff );_fdf !=nil {return _fe ,_ef .Wrapf (_fdf ,_eecd ,"\u0062i\u0074\u003a\u0020\u0025\u0064",_aba );
};};return _afad ,nil ;};func (_cbd *Reader )ReadBit ()(_accg int ,_dde error ){_beg ,_dde :=_cbd .readBool ();if _dde !=nil {return 0,_dde ;};if _beg {_accg =1;};return _accg ,nil ;};func NewWriter (data []byte )*Writer {return &Writer {_gfd :data }};
func (_fcd *Reader )ReadBool ()(bool ,error ){return _fcd .readBool ()};type readerSource struct{_efb []byte ;_cb int ;_deg int ;};const (_cd =64;_cea =int (^uint (0)>>1););func (_dg *BufferedWriter )writeByte (_gd byte ){switch {case _dg ._ffg ==0:_dg ._d [_dg ._b ]=_gd ;
_dg ._b ++;case _dg ._aa :_dg ._d [_dg ._b ]|=_gd >>_dg ._ffg ;_dg ._b ++;_dg ._d [_dg ._b ]=byte (uint16 (_gd )<<(8-_dg ._ffg )&0xff);default:_dg ._d [_dg ._b ]|=byte (uint16 (_gd )<<_dg ._ffg &0xff);_dg ._b ++;_dg ._d [_dg ._b ]=_gd >>(8-_dg ._ffg );
};};func (_gcec *Reader )readBool ()(_bfe bool ,_agdc error ){if _gcec ._ade ==0{_gcec ._dge ,_agdc =_gcec .readBufferByte ();if _agdc !=nil {return false ,_agdc ;};_bfe =(_gcec ._dge &0x80)!=0;_gcec ._dge ,_gcec ._ade =_gcec ._dge &0x7f,7;return _bfe ,nil ;
};_gcec ._ade --;_bfe =(_gcec ._dge &(1<<_gcec ._ade ))!=0;_gcec ._dge &=1<<_gcec ._ade -1;return _bfe ,nil ;};func (_cgda *Reader )ReadByte ()(byte ,error ){if _cgda ._ade ==0{return _cgda .readBufferByte ();};return _cgda .readUnalignedByte ();};type Reader struct{_af readerSource ;
_dge byte ;_ade byte ;_df int64 ;_bdc int ;_ag int ;_afg int64 ;_gdcd byte ;_ae byte ;_bbda int ;};func (_efbg *Writer )byteCapacity ()int {_bdf :=len (_efbg ._gfd )-_efbg ._cdb ;if _efbg ._egb !=0{_bdf --;};return _bdf ;};func (_aafa *Reader )Reset (){_aafa ._df =_aafa ._afg ;
_aafa ._ade =_aafa ._gdcd ;_aafa ._dge =_aafa ._ae ;_aafa ._bdc =_aafa ._bbda ;};func (_ee *BufferedWriter )WriteBit (bit int )error {if bit !=1&&bit !=0{return _ef .Errorf ("\u0042\u0075\u0066fe\u0072\u0065\u0064\u0057\u0072\u0069\u0074\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0042\u0069\u0074","\u0062\u0069\u0074\u0020\u0076\u0061\u006cu\u0065\u0020\u006du\u0073\u0074\u0020\u0062e\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u007b\u0030\u002c\u0031\u007d\u0020\u0062\u0075\u0074\u0020\u0069\u0073\u003a\u0020\u0025\u0064",bit );
};if len (_ee ._d )-1< _ee ._b {_ee .expandIfNeeded (1);};_aca :=_ee ._ffg ;if _ee ._aa {_aca =7-_ee ._ffg ;};_ee ._d [_ee ._b ]|=byte (uint16 (bit <<_aca )&0xff);_ee ._ffg ++;if _ee ._ffg ==8{_ee ._b ++;_ee ._ffg =0;};return nil ;};func (_aaa *BufferedWriter )expandIfNeeded (_ccg int ){if !_aaa .tryGrowByReslice (_ccg ){_aaa .grow (_ccg );
};};func BufferedMSB ()*BufferedWriter {return &BufferedWriter {_aa :true }};var (_ _a .Reader =&Reader {};_ _a .ByteReader =&Reader {};_ _a .Seeker =&Reader {};_ StreamReader =&Reader {};);type BinaryWriter interface{BitWriter ;_a .Writer ;_a .ByteWriter ;
Data ()[]byte ;};func (_aag *Reader )ReadUint32 ()(uint32 ,error ){_cf :=make ([]byte ,4);_ ,_aed :=_aag .Read (_cf );if _aed !=nil {return 0,_aed ;};return _ce .BigEndian .Uint32 (_cf ),nil ;};type BufferedWriter struct{_d []byte ;_ffg uint8 ;_b int ;
_aa bool ;};func (_fcc *Reader )AbsoluteLength ()uint64 {return uint64 (len (_fcc ._af ._efb ))};func (_gg *Reader )Seek (offset int64 ,whence int )(int64 ,error ){_gg ._ag =-1;_gg ._ade =0;_gg ._dge =0;_gg ._bdc =0;var _edd int64 ;switch whence {case _a .SeekStart :_edd =offset ;
case _a .SeekCurrent :_edd =_gg ._df +offset ;case _a .SeekEnd :_edd =int64 (_gg ._af ._deg )+offset ;default:return 0,_ac .New ("\u0072\u0065\u0061de\u0072\u002e\u0052\u0065\u0061\u0064\u0065\u0072\u002eS\u0065e\u006b:\u0020i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0077\u0068\u0065\u006e\u0063\u0065");
};if _edd < 0{return 0,_ac .New ("\u0072\u0065a\u0064\u0065\u0072\u002eR\u0065\u0061d\u0065\u0072\u002e\u0053\u0065\u0065\u006b\u003a \u006e\u0065\u0067\u0061\u0074\u0069\u0076\u0065\u0020\u0070\u006f\u0073i\u0074\u0069\u006f\u006e");};_gg ._df =_edd ;
_gg ._ade =0;return _edd ,nil ;};func (_fd *BufferedWriter )Len ()int {return _fd .byteCapacity ()};func (_dgc *Writer )UseMSB ()bool {return _dgc ._gad };func (_gcc *Writer )SkipBits (skip int )error {const _bfc ="\u0057r\u0069t\u0065\u0072\u002e\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073";
if skip ==0{return nil ;};_cde :=int (_gcc ._egb )+skip ;if _cde >=0&&_cde < 8{_gcc ._egb =uint8 (_cde );return nil ;};_cde =int (_gcc ._egb )+_gcc ._cdb *8+skip ;if _cde < 0{return _ef .Errorf (_bfc ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");
};_fbc :=_cde /8;_fdb :=_cde %8;_f .Log .Trace ("\u0053\u006b\u0069\u0070\u0042\u0069\u0074\u0073");_f .Log .Trace ("\u0042\u0069\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0042\u0079\u0074\u0065\u0049n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0046\u0075\u006c\u006c\u0042\u0069\u0074\u0073\u003a\u0020'\u0025\u0064\u0027\u002c\u0020\u004c\u0065\u006e\u003a\u0020\u0027\u0025\u0064\u0027,\u0020\u0043\u0061p\u003a\u0020\u0027\u0025\u0064\u0027",_gcc ._egb ,_gcc ._cdb ,int (_gcc ._egb )+(_gcc ._cdb )*8,len (_gcc ._gfd ),cap (_gcc ._gfd ));
_f .Log .Trace ("S\u006b\u0069\u0070\u003a\u0020\u0027%\u0064\u0027\u002c\u0020\u0064\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062i\u0074\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025d\u0027",skip ,_cde ,_fdb );_gcc ._egb =uint8 (_fdb );if _bfd :=_fbc -_gcc ._cdb ;
_bfd > 0&&len (_gcc ._gfd )-1< _fbc {_f .Log .Trace ("\u0042\u0079\u0074e\u0044\u0069\u0066\u0066\u003a\u0020\u0025\u0064",_bfd );return _ef .Errorf (_bfc ,"\u0069n\u0064e\u0078\u0020\u006f\u0075\u0074 \u006f\u0066 \u0072\u0061\u006e\u0067\u0065");};_gcc ._cdb =_fbc ;
_f .Log .Trace ("\u0042\u0069\u0074I\u006e\u0064\u0065\u0078:\u0020\u0027\u0025\u0064\u0027\u002c\u0020B\u0079\u0074\u0065\u0049\u006e\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027",_gcc ._egb ,_gcc ._cdb );return nil ;};func (_daf *Writer )Write (p []byte )(int ,error ){if len (p )> _daf .byteCapacity (){return 0,_a .EOF ;
};for _ ,_eec :=range p {if _ccfc :=_daf .writeByte (_eec );_ccfc !=nil {return 0,_ccfc ;};};return len (p ),nil ;};func (_afb *Reader )NewPartialReader (offset ,length int ,relative bool )(*Reader ,error ){if offset < 0{return nil ,_ac .New ("p\u0061\u0072\u0074\u0069\u0061\u006c\u0020\u0072\u0065\u0061\u0064\u0065\u0072\u0020\u006f\u0066\u0066\u0073e\u0074\u0020\u0063\u0061\u006e\u006e\u006f\u0074\u0020\u0062e \u006e\u0065\u0067a\u0074i\u0076\u0065");
};if relative {offset =_afb ._af ._cb +offset ;};if length > 0{_cge :=len (_afb ._af ._efb );if relative {_cge =_afb ._af ._deg ;};if offset +length > _cge {return nil ,_e .Errorf ("\u0070\u0061r\u0074\u0069\u0061l\u0020\u0072\u0065\u0061\u0064e\u0072\u0020\u006f\u0066\u0066se\u0074\u0028\u0025\u0064\u0029\u002b\u006c\u0065\u006e\u0067\u0074\u0068\u0028\u0025\u0064\u0029\u003d\u0025d\u0020i\u0073\u0020\u0067\u0072\u0065\u0061ter\u0020\u0074\u0068\u0061\u006e\u0020\u0074\u0068\u0065\u0020\u006f\u0072ig\u0069n\u0061\u006c\u0020\u0072e\u0061d\u0065r\u0020\u006ce\u006e\u0067th\u003a\u0020\u0025\u0064",offset ,length ,offset +length ,_afb ._af ._deg );
};};if length < 0{_eeg :=len (_afb ._af ._efb );if relative {_eeg =_afb ._af ._deg ;};length =_eeg -offset ;};return &Reader {_af :readerSource {_efb :_afb ._af ._efb ,_deg :length ,_cb :offset }},nil ;};func (_aac *Reader )readBufferByte ()(byte ,error ){if _aac ._df >=int64 (_aac ._af ._deg ){return 0,_a .EOF ;
};_aac ._ag =-1;_bcd :=_aac ._af ._efb [int64 (_aac ._af ._cb )+_aac ._df ];_aac ._df ++;_aac ._bdc =int (_bcd );return _bcd ,nil ;};