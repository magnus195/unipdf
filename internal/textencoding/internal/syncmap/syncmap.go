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

package syncmap ;import _b "sync";type RuneUint16Map struct{_ab map[rune ]uint16 ;_abb _b .RWMutex ;};func (_ggg *RuneSet )Write (r rune ){_ggg ._dca .Lock ();defer _ggg ._dca .Unlock ();_ggg ._bg [r ]=struct{}{};};type RuneByteMap struct{_bd map[rune ]byte ;
_cc _b .RWMutex ;};type StringsMap struct{_egd map[string ]string ;_cf _b .RWMutex ;};func NewByteRuneMap (m map[byte ]rune )*ByteRuneMap {return &ByteRuneMap {_e :m }};func MakeRuneByteMap (length int )*RuneByteMap {_gg :=make (map[rune ]byte ,length );
return &RuneByteMap {_bd :_gg };};func (_bga *RuneSet )Range (f func (_bf rune )(_cdb bool )){_bga ._dca .RLock ();defer _bga ._dca .RUnlock ();for _baf :=range _bga ._bg {if f (_baf ){break ;};};};func (_efff *StringRuneMap )Write (g string ,r rune ){_efff ._cad .Lock ();
defer _efff ._cad .Unlock ();_efff ._gf [g ]=r ;};func NewRuneStringMap (m map[rune ]string )*RuneStringMap {return &RuneStringMap {_fb :m }};func (_eaf *StringsMap )Write (g1 ,g2 string ){_eaf ._cf .Lock ();defer _eaf ._cf .Unlock ();_eaf ._egd [g1 ]=g2 ;
};func (_edd *RuneStringMap )Read (r rune )(string ,bool ){_edd ._ed .RLock ();defer _edd ._ed .RUnlock ();_fg ,_ffg :=_edd ._fb [r ];return _fg ,_ffg ;};func (_ce *RuneByteMap )Read (r rune )(byte ,bool ){_ce ._cc .RLock ();defer _ce ._cc .RUnlock ();
_ea ,_gc :=_ce ._bd [r ];return _ea ,_gc ;};type StringRuneMap struct{_gf map[string ]rune ;_cad _b .RWMutex ;};func (_acf *StringRuneMap )Length ()int {_acf ._cad .RLock ();defer _acf ._cad .RUnlock ();return len (_acf ._gf );};func (_dee *ByteRuneMap )Length ()int {_dee ._d .RLock ();
defer _dee ._d .RUnlock ();return len (_dee ._e )};func (_ca *ByteRuneMap )Range (f func (_a byte ,_ff rune )(_ac bool )){_ca ._d .RLock ();defer _ca ._d .RUnlock ();for _de ,_g :=range _ca ._e {if f (_de ,_g ){break ;};};};func (_gea *StringRuneMap )Range (f func (_ad string ,_baa rune )(_fd bool )){_gea ._cad .RLock ();
defer _gea ._cad .RUnlock ();for _cga ,_geb :=range _gea ._gf {if f (_cga ,_geb ){break ;};};};func (_eab *RuneStringMap )Write (r rune ,s string ){_eab ._ed .Lock ();defer _eab ._ed .Unlock ();_eab ._fb [r ]=s ;};func MakeRuneSet (length int )*RuneSet {return &RuneSet {_bg :make (map[rune ]struct{},length )}};
type RuneSet struct{_bg map[rune ]struct{};_dca _b .RWMutex ;};func (_dc *RuneByteMap )Length ()int {_dc ._cc .RLock ();defer _dc ._cc .RUnlock ();return len (_dc ._bd )};func (_c *ByteRuneMap )Read (b byte )(rune ,bool ){_c ._d .RLock ();defer _c ._d .RUnlock ();
_be ,_cg :=_c ._e [b ];return _be ,_cg ;};func (_aec *RuneUint16Map )RangeDelete (f func (_beb rune ,_db uint16 )(_dfe bool ,_dea bool )){_aec ._abb .Lock ();defer _aec ._abb .Unlock ();for _fgbb ,_gb :=range _aec ._ab {_gbc ,_cgc :=f (_fgbb ,_gb );if _gbc {delete (_aec ._ab ,_fgbb );
};if _cgc {break ;};};};func (_bda *RuneUint16Map )Delete (r rune ){_bda ._abb .Lock ();defer _bda ._abb .Unlock ();delete (_bda ._ab ,r );};func (_bfa *StringsMap )Copy ()*StringsMap {_bfa ._cf .RLock ();defer _bfa ._cf .RUnlock ();_eddd :=map[string ]string {};
for _ddc ,_caae :=range _bfa ._egd {_eddd [_ddc ]=_caae ;};return &StringsMap {_egd :_eddd };};func (_ee *StringsMap )Read (g string )(string ,bool ){_ee ._cf .RLock ();defer _ee ._cf .RUnlock ();_dfg ,_gd :=_ee ._egd [g ];return _dfg ,_gd ;};func (_af *RuneUint16Map )Range (f func (_fgd rune ,_ec uint16 )(_cgb bool )){_af ._abb .RLock ();
defer _af ._abb .RUnlock ();for _bb ,_ge :=range _af ._ab {if f (_bb ,_ge ){break ;};};};func (_ba *RuneByteMap )Range (f func (_fa rune ,_ga byte )(_fc bool )){_ba ._cc .RLock ();defer _ba ._cc .RUnlock ();for _eaa ,_ef :=range _ba ._bd {if f (_eaa ,_ef ){break ;
};};};func (_cb *StringRuneMap )Read (g string )(rune ,bool ){_cb ._cad .RLock ();defer _cb ._cad .RUnlock ();_bcg ,_eff :=_cb ._gf [g ];return _bcg ,_eff ;};func (_ddf *RuneUint16Map )Length ()int {_ddf ._abb .RLock ();defer _ddf ._abb .RUnlock ();return len (_ddf ._ab );
};type RuneStringMap struct{_fb map[rune ]string ;_ed _b .RWMutex ;};type ByteRuneMap struct{_e map[byte ]rune ;_d _b .RWMutex ;};func MakeRuneUint16Map (length int )*RuneUint16Map {return &RuneUint16Map {_ab :make (map[rune ]uint16 ,length )};};func (_egc *RuneUint16Map )Write (r rune ,g uint16 ){_egc ._abb .Lock ();
defer _egc ._abb .Unlock ();_egc ._ab [r ]=g ;};func (_aa *RuneSet )Exists (r rune )bool {_aa ._dca .RLock ();defer _aa ._dca .RUnlock ();_ ,_df :=_aa ._bg [r ];return _df ;};func (_bad *StringsMap )Range (f func (_cbg ,_cde string )(_deac bool )){_bad ._cf .RLock ();
defer _bad ._cf .RUnlock ();for _ag ,_bcf :=range _bad ._egd {if f (_ag ,_bcf ){break ;};};};type StringsTuple struct{Key ,Value string ;};func (_dd *ByteRuneMap )Write (b byte ,r rune ){_dd ._d .Lock ();defer _dd ._d .Unlock ();_dd ._e [b ]=r };func NewStringsMap (tuples []StringsTuple )*StringsMap {_cdd :=map[string ]string {};
for _ ,_edb :=range tuples {_cdd [_edb .Key ]=_edb .Value ;};return &StringsMap {_egd :_cdd };};func (_cd *RuneByteMap )Write (r rune ,b byte ){_cd ._cc .Lock ();defer _cd ._cc .Unlock ();_cd ._bd [r ]=b };func NewStringRuneMap (m map[string ]rune )*StringRuneMap {return &StringRuneMap {_gf :m }};
func (_caa *RuneStringMap )Length ()int {_caa ._ed .RLock ();defer _caa ._ed .RUnlock ();return len (_caa ._fb );};func (_cae *RuneUint16Map )Read (r rune )(uint16 ,bool ){_cae ._abb .RLock ();defer _cae ._abb .RUnlock ();_ae ,_fgb :=_cae ._ab [r ];return _ae ,_fgb ;
};func (_cdg *RuneStringMap )Range (f func (_bc rune ,_dg string )(_cdge bool )){_cdg ._ed .RLock ();defer _cdg ._ed .RUnlock ();for _aab ,_eg :=range _cdg ._fb {if f (_aab ,_eg ){break ;};};};func MakeByteRuneMap (length int )*ByteRuneMap {return &ByteRuneMap {_e :make (map[byte ]rune ,length )}};
func (_cge *RuneSet )Length ()int {_cge ._dca .RLock ();defer _cge ._dca .RUnlock ();return len (_cge ._bg )};