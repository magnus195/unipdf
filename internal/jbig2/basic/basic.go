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

package basic ;import _f "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func (_da IntSlice )Get (index int )(int ,error ){if index > len (_da )-1{return 0,_f .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );
};return _da [index ],nil ;};func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func Min (x ,y int )int {if x < y {return x ;};return y ;};func (_fc IntsMap )Add (key uint64 ,value int ){_fc [key ]=append (_fc [key ],value )};
func (_ff *NumSlice )Add (v float32 ){*_ff =append (*_ff ,v )};func (_acf NumSlice )GetIntSlice ()[]int {_ffc :=make ([]int ,len (_acf ));for _cc ,_ea :=range _acf {_ffc [_cc ]=int (_ea );};return _ffc ;};func (_af *NumSlice )AddInt (v int ){*_af =append (*_af ,float32 (v ))};
type IntSlice []int ;func (_bd IntsMap )GetSlice (key uint64 )([]int ,bool ){_dc ,_dd :=_bd [key ];if !_dd {return nil ,false ;};return _dc ,true ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func (_dff *Stack )peek ()(interface{},bool ){_gc :=_dff .top ();
if _gc ==-1{return nil ,false ;};return _dff .Data [_gc ],true ;};func (_db NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_db )-1{return 0,_f .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};return _db [i ],nil ;};func (_e IntsMap )Get (key uint64 )(int ,bool ){_b ,_d :=_e [key ];if !_d {return 0,false ;};if len (_b )==0{return 0,false ;};return _b [0],true ;};type Stack struct{Data []interface{};Aux *Stack ;};func (_bc *IntSlice )Add (v int )error {if _bc ==nil {return _f .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");
};*_bc =append (*_bc ,v );return nil ;};func (_ba *Stack )Len ()int {return len (_ba .Data )};func (_a IntsMap )Delete (key uint64 ){delete (_a ,key )};func Max (x ,y int )int {if x > y {return x ;};return y ;};func (_df *Stack )Peek ()(_cb interface{},_ce bool ){return _df .peek ()};
func NewIntSlice (i int )*IntSlice {_fg :=IntSlice (make ([]int ,i ));return &_fg };func (_c NumSlice )GetInt (i int )(int ,error ){const _ab ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_c )-1{return 0,_f .Errorf (_ab ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};_ca :=_c [i ];return int (_ca +Sign (_ca )*0.5),nil ;};func (_ga *Stack )top ()int {return len (_ga .Data )-1};type NumSlice []float32 ;func (_dbd *Stack )Pop ()(_gd interface{},_fga bool ){_gd ,_fga =_dbd .peek ();if !_fga {return nil ,_fga ;};_dbd .Data =_dbd .Data [:_dbd .top ()];
return _gd ,true ;};func NewNumSlice (i int )*NumSlice {_dag :=NumSlice (make ([]float32 ,i ));return &_dag };func (_dae *Stack )Push (v interface{}){_dae .Data =append (_dae .Data ,v )};func (_bdg IntSlice )Size ()int {return len (_bdg )};func Abs (v int )int {if v > 0{return v ;
};return -v ;};type IntsMap map[uint64 ][]int ;func (_fcd *IntSlice )Copy ()*IntSlice {_ac :=IntSlice (make ([]int ,len (*_fcd )));copy (_ac ,*_fcd );return &_ac ;};