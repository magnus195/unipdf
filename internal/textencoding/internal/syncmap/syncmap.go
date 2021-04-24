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

package syncmap ;import _c "sync";type StringsTuple struct{Key ,Value string ;};func (_db *ByteRuneMap )Read (b byte )(rune ,bool ){_db ._g .RLock ();defer _db ._g .RUnlock ();_e ,_fe :=_db ._d [b ];return _e ,_fe ;};func (_ea *ByteRuneMap )Range (f func (_a byte ,_b rune )(_fb bool )){_ea ._g .RLock ();
defer _ea ._g .RUnlock ();for _ff ,_bg :=range _ea ._d {if f (_ff ,_bg ){break ;};};};func (_cgg *RuneUint16Map )RangeDelete (f func (_gd rune ,_bac uint16 )(_ccc bool ,_cee bool )){_cgg ._dca .Lock ();defer _cgg ._dca .Unlock ();for _dd ,_aa :=range _cgg ._cag {_ag ,_gc :=f (_dd ,_aa );
if _ag {delete (_cgg ._cag ,_dd );};if _gc {break ;};};};func (_ae *RuneStringMap )Range (f func (_da rune ,_gafb string )(_ggd bool )){_ae ._gaf .RLock ();defer _ae ._gaf .RUnlock ();for _ede ,_ebc :=range _ae ._ccd {if f (_ede ,_ebc ){break ;};};};func (_af *RuneByteMap )Write (r rune ,b byte ){_af ._eac .Lock ();
defer _af ._eac .Unlock ();_af ._eb [r ]=b ;};func MakeRuneSet (length int )*RuneSet {return &RuneSet {_abd :make (map[rune ]struct{},length )}};func (_ge *RuneByteMap )Length ()int {_ge ._eac .RLock ();defer _ge ._eac .RUnlock ();return len (_ge ._eb )};
func (_gg *RuneSet )Exists (r rune )bool {_gg ._cc .RLock ();defer _gg ._cc .RUnlock ();_ ,_ad :=_gg ._abd [r ];return _ad ;};func (_abcc *RuneUint16Map )Length ()int {_abcc ._dca .RLock ();defer _abcc ._dca .RUnlock ();return len (_abcc ._cag );};func (_be *ByteRuneMap )Length ()int {_be ._g .RLock ();
defer _be ._g .RUnlock ();return len (_be ._d )};func (_beg *RuneSet )Length ()int {_beg ._cc .RLock ();defer _beg ._cc .RUnlock ();return len (_beg ._abd )};type RuneUint16Map struct{_cag map[rune ]uint16 ;_dca _c .RWMutex ;};func (_ebb *RuneSet )Range (f func (_fcf rune )(_gbdf bool )){_ebb ._cc .RLock ();
defer _ebb ._cc .RUnlock ();for _df :=range _ebb ._abd {if f (_df ){break ;};};};func (_fbc *RuneStringMap )Length ()int {_fbc ._gaf .RLock ();defer _fbc ._gaf .RUnlock ();return len (_fbc ._ccd );};func (_fc *RuneByteMap )Read (r rune )(byte ,bool ){_fc ._eac .RLock ();
defer _fc ._eac .RUnlock ();_dc ,_gbd :=_fc ._eb [r ];return _dc ,_gbd ;};func (_fg *RuneUint16Map )Delete (r rune ){_fg ._dca .Lock ();defer _fg ._dca .Unlock ();delete (_fg ._cag ,r );};type RuneByteMap struct{_eb map[rune ]byte ;_eac _c .RWMutex ;};
func MakeRuneUint16Map (length int )*RuneUint16Map {return &RuneUint16Map {_cag :make (map[rune ]uint16 ,length )};};func (_cef *StringRuneMap )Length ()int {_cef ._dgf .RLock ();defer _cef ._dgf .RUnlock ();return len (_cef ._gf );};func (_ffd *RuneUint16Map )Write (r rune ,g uint16 ){_ffd ._dca .Lock ();
defer _ffd ._dca .Unlock ();_ffd ._cag [r ]=g ;};func MakeByteRuneMap (length int )*ByteRuneMap {return &ByteRuneMap {_d :make (map[byte ]rune ,length )}};type StringRuneMap struct{_gf map[string ]rune ;_dgf _c .RWMutex ;};func (_cf *RuneByteMap )Range (f func (_ab rune ,_abc byte )(_feg bool )){_cf ._eac .RLock ();
defer _cf ._eac .RUnlock ();for _eag ,_ef :=range _cf ._eb {if f (_eag ,_ef ){break ;};};};func NewStringRuneMap (m map[string ]rune )*StringRuneMap {return &StringRuneMap {_gf :m }};func (_gdb *StringsMap )Copy ()*StringsMap {_gdb ._bdc .RLock ();defer _gdb ._bdc .RUnlock ();
_dfc :=map[string ]string {};for _beb ,_ccb :=range _gdb ._ffg {_dfc [_beb ]=_ccb ;};return &StringsMap {_ffg :_dfc };};func (_cgb *StringRuneMap )Range (f func (_caa string ,_bgb rune )(_ebd bool )){_cgb ._dgf .RLock ();defer _cgb ._dgf .RUnlock ();for _adc ,_eg :=range _cgb ._gf {if f (_adc ,_eg ){break ;
};};};func (_fbbf *StringsMap )Write (g1 ,g2 string ){_fbbf ._bdc .Lock ();defer _fbbf ._bdc .Unlock ();_fbbf ._ffg [g1 ]=g2 ;};func (_aed *RuneUint16Map )Read (r rune )(uint16 ,bool ){_aed ._dca .RLock ();defer _aed ._dca .RUnlock ();_bd ,_dg :=_aed ._cag [r ];
return _bd ,_dg ;};func (_gce *StringRuneMap )Read (g string )(rune ,bool ){_gce ._dgf .RLock ();defer _gce ._dgf .RUnlock ();_fbb ,_afb :=_gce ._gf [g ];return _fbb ,_afb ;};type RuneSet struct{_abd map[rune ]struct{};_cc _c .RWMutex ;};func (_ed *RuneStringMap )Read (r rune )(string ,bool ){_ed ._gaf .RLock ();
defer _ed ._gaf .RUnlock ();_cca ,_ca :=_ed ._ccd [r ];return _cca ,_ca ;};func NewStringsMap (tuples []StringsTuple )*StringsMap {_gca :=map[string ]string {};for _ ,_bdg :=range tuples {_gca [_bdg .Key ]=_bdg .Value ;};return &StringsMap {_ffg :_gca };
};func (_ade *StringsMap )Range (f func (_def ,_ec string )(_edd bool )){_ade ._bdc .RLock ();defer _ade ._bdc .RUnlock ();for _fgf ,_bdcf :=range _ade ._ffg {if f (_fgf ,_bdcf ){break ;};};};func (_de *RuneUint16Map )Range (f func (_cg rune ,_dcb uint16 )(_fec bool )){_de ._dca .RLock ();
defer _de ._dca .RUnlock ();for _ba ,_ceb :=range _de ._cag {if f (_ba ,_ceb ){break ;};};};type ByteRuneMap struct{_d map[byte ]rune ;_g _c .RWMutex ;};type StringsMap struct{_ffg map[string ]string ;_bdc _c .RWMutex ;};func NewRuneStringMap (m map[rune ]string )*RuneStringMap {return &RuneStringMap {_ccd :m }};
type RuneStringMap struct{_ccd map[rune ]string ;_gaf _c .RWMutex ;};func (_bc *RuneStringMap )Write (r rune ,s string ){_bc ._gaf .Lock ();defer _bc ._gaf .Unlock ();_bc ._ccd [r ]=s ;};func (_geb *StringsMap )Read (g string )(string ,bool ){_geb ._bdc .RLock ();
defer _geb ._bdc .RUnlock ();_ddf ,_ac :=_geb ._ffg [g ];return _ddf ,_ac ;};func (_ebcg *StringRuneMap )Write (g string ,r rune ){_ebcg ._dgf .Lock ();defer _ebcg ._dgf .Unlock ();_ebcg ._gf [g ]=r ;};func (_ga *RuneSet )Write (r rune ){_ga ._cc .Lock ();
defer _ga ._cc .Unlock ();_ga ._abd [r ]=struct{}{}};func MakeRuneByteMap (length int )*RuneByteMap {_bf :=make (map[rune ]byte ,length );return &RuneByteMap {_eb :_bf };};func (_gb *ByteRuneMap )Write (b byte ,r rune ){_gb ._g .Lock ();defer _gb ._g .Unlock ();
_gb ._d [b ]=r };func NewByteRuneMap (m map[byte ]rune )*ByteRuneMap {return &ByteRuneMap {_d :m }};