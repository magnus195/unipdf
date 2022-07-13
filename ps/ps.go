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

// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
//
// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
package ps ;import (_ec "bufio";_e "bytes";_d "errors";_f "fmt";_c "github.com/unidoc/unipdf/v3/common";_de "github.com/unidoc/unipdf/v3/core";_g "io";_ge "math";);

// PSBoolean represents a boolean value.
type PSBoolean struct{Val bool ;};func (_edb *PSProgram )String ()string {_cg :="\u007b\u0020";for _ ,_dd :=range *_edb {_cg +=_dd .String ();_cg +="\u0020";};_cg +="\u007d";return _cg ;};

// Empty empties the stack.
func (_adba *PSStack )Empty (){*_adba =[]PSObject {}};

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor (program *PSProgram )*PSExecutor {_eb :=&PSExecutor {};_eb .Stack =NewPSStack ();_eb ._cc =program ;return _eb ;};

// PSReal represents a real number.
type PSReal struct{Val float64 ;};func (_beac *PSParser )skipSpaces ()(int ,error ){_afaee :=0;for {_gcg ,_agaa :=_beac ._bbga .Peek (1);if _agaa !=nil {return 0,_agaa ;};if _de .IsWhiteSpace (_gcg [0]){_beac ._bbga .ReadByte ();_afaee ++;}else {break ;
};};return _afaee ,nil ;};func (_ceg *PSOperand )String ()string {return string (*_ceg )};

// PSOperand represents a Postscript operand (text string).
type PSOperand string ;

// PSInteger represents an integer.
type PSInteger struct{Val int ;};func (_cae *PSOperand )roll (_aeab *PSStack )error {_aeaa ,_dbfd :=_aeab .Pop ();if _dbfd !=nil {return _dbfd ;};_afcda ,_dbfd :=_aeab .Pop ();if _dbfd !=nil {return _dbfd ;};_bcdaf ,_fec :=_aeaa .(*PSInteger );if !_fec {return ErrTypeCheck ;
};_ceaa ,_fec :=_afcda .(*PSInteger );if !_fec {return ErrTypeCheck ;};if _ceaa .Val < 0{return ErrRangeCheck ;};if _ceaa .Val ==0||_ceaa .Val ==1{return nil ;};if _ceaa .Val > len (*_aeab ){return ErrStackUnderflow ;};for _gegf :=0;_gegf < _aefb (_bcdaf .Val );
_gegf ++{var _dbe []PSObject ;_dbe =(*_aeab )[len (*_aeab )-(_ceaa .Val ):len (*_aeab )];if _bcdaf .Val > 0{_cfb :=_dbe [len (_dbe )-1];_dbe =append ([]PSObject {_cfb },_dbe [0:len (_dbe )-1]...);}else {_bdc :=_dbe [len (_dbe )-_ceaa .Val ];_dbe =append (_dbe [1:],_bdc );
};_ebf :=append ((*_aeab )[0:len (*_aeab )-_ceaa .Val ],_dbe ...);_aeab =&_ebf ;};return nil ;};const _b =0.000001;

// Exec executes the operand `op` in the state specified by `stack`.
func (_cgf *PSOperand )Exec (stack *PSStack )error {_deb :=ErrUnsupportedOperand ;switch *_cgf {case "\u0061\u0062\u0073":_deb =_cgf .abs (stack );case "\u0061\u0064\u0064":_deb =_cgf .add (stack );case "\u0061\u006e\u0064":_deb =_cgf .and (stack );case "\u0061\u0074\u0061\u006e":_deb =_cgf .atan (stack );
case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":_deb =_cgf .bitshift (stack );case "\u0063e\u0069\u006c\u0069\u006e\u0067":_deb =_cgf .ceiling (stack );case "\u0063\u006f\u0070\u0079":_deb =_cgf .copy (stack );case "\u0063\u006f\u0073":_deb =_cgf .cos (stack );
case "\u0063\u0076\u0069":_deb =_cgf .cvi (stack );case "\u0063\u0076\u0072":_deb =_cgf .cvr (stack );case "\u0064\u0069\u0076":_deb =_cgf .div (stack );case "\u0064\u0075\u0070":_deb =_cgf .dup (stack );case "\u0065\u0071":_deb =_cgf .eq (stack );case "\u0065\u0078\u0063\u0068":_deb =_cgf .exch (stack );
case "\u0065\u0078\u0070":_deb =_cgf .exp (stack );case "\u0066\u006c\u006fo\u0072":_deb =_cgf .floor (stack );case "\u0067\u0065":_deb =_cgf .ge (stack );case "\u0067\u0074":_deb =_cgf .gt (stack );case "\u0069\u0064\u0069\u0076":_deb =_cgf .idiv (stack );
case "\u0069\u0066":_deb =_cgf .ifCondition (stack );case "\u0069\u0066\u0065\u006c\u0073\u0065":_deb =_cgf .ifelse (stack );case "\u0069\u006e\u0064e\u0078":_deb =_cgf .index (stack );case "\u006c\u0065":_deb =_cgf .le (stack );case "\u006c\u006f\u0067":_deb =_cgf .log (stack );
case "\u006c\u006e":_deb =_cgf .ln (stack );case "\u006c\u0074":_deb =_cgf .lt (stack );case "\u006d\u006f\u0064":_deb =_cgf .mod (stack );case "\u006d\u0075\u006c":_deb =_cgf .mul (stack );case "\u006e\u0065":_deb =_cgf .ne (stack );case "\u006e\u0065\u0067":_deb =_cgf .neg (stack );
case "\u006e\u006f\u0074":_deb =_cgf .not (stack );case "\u006f\u0072":_deb =_cgf .or (stack );case "\u0070\u006f\u0070":_deb =_cgf .pop (stack );case "\u0072\u006f\u0075n\u0064":_deb =_cgf .round (stack );case "\u0072\u006f\u006c\u006c":_deb =_cgf .roll (stack );
case "\u0073\u0069\u006e":_deb =_cgf .sin (stack );case "\u0073\u0071\u0072\u0074":_deb =_cgf .sqrt (stack );case "\u0073\u0075\u0062":_deb =_cgf .sub (stack );case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":_deb =_cgf .truncate (stack );case "\u0078\u006f\u0072":_deb =_cgf .xor (stack );
};return _deb ;};func (_beba *PSOperand )copy (_bcfe *PSStack )error {_feaa ,_gdc :=_bcfe .PopInteger ();if _gdc !=nil {return _gdc ;};if _feaa < 0{return ErrRangeCheck ;};if _feaa > len (*_bcfe ){return ErrRangeCheck ;};*_bcfe =append (*_bcfe ,(*_bcfe )[len (*_bcfe )-_feaa :]...);
return nil ;};func (_ed *PSReal )String ()string {return _f .Sprintf ("\u0025\u002e\u0035\u0066",_ed .Val )};

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser (content []byte )*PSParser {_gcfea :=PSParser {};_fdf :=_e .NewBuffer (content );_gcfea ._bbga =_ec .NewReader (_fdf );return &_gcfea ;};func (_dbg *PSOperand )ge (_dff *PSStack )error {_ffcd ,_fba :=_dff .PopNumberAsFloat64 ();if _fba !=nil {return _fba ;
};_dgg ,_fba :=_dff .PopNumberAsFloat64 ();if _fba !=nil {return _fba ;};if _ge .Abs (_dgg -_ffcd )< _b {_agg :=_dff .Push (MakeBool (true ));return _agg ;}else if _dgg > _ffcd {_egda :=_dff .Push (MakeBool (true ));return _egda ;}else {_dag :=_dff .Push (MakeBool (false ));
return _dag ;};};func (_ffc *PSBoolean )DebugString ()string {return _f .Sprintf ("\u0062o\u006f\u006c\u003a\u0025\u0076",_ffc .Val );};func (_edg *PSOperand )or (_dba *PSStack )error {_bad ,_bca :=_dba .Pop ();if _bca !=nil {return _bca ;};_badf ,_bca :=_dba .Pop ();
if _bca !=nil {return _bca ;};if _ead ,_agac :=_bad .(*PSBoolean );_agac {_afd ,_gec :=_badf .(*PSBoolean );if !_gec {return ErrTypeCheck ;};_bca =_dba .Push (MakeBool (_ead .Val ||_afd .Val ));return _bca ;};if _gcfe ,_gfgb :=_bad .(*PSInteger );_gfgb {_eaec ,_ceac :=_badf .(*PSInteger );
if !_ceac {return ErrTypeCheck ;};_bca =_dba .Push (MakeInteger (_gcfe .Val |_eaec .Val ));return _bca ;};return ErrTypeCheck ;};func (_gfg *PSOperand )lt (_acbe *PSStack )error {_afa ,_ggd :=_acbe .PopNumberAsFloat64 ();if _ggd !=nil {return _ggd ;};_agdc ,_ggd :=_acbe .PopNumberAsFloat64 ();
if _ggd !=nil {return _ggd ;};if _ge .Abs (_agdc -_afa )< _b {_egf :=_acbe .Push (MakeBool (false ));return _egf ;}else if _agdc < _afa {_ebaa :=_acbe .Push (MakeBool (true ));return _ebaa ;}else {_dfff :=_acbe .Push (MakeBool (false ));return _dfff ;};
};

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram ()*PSProgram {return &PSProgram {}};var ErrUndefinedResult =_d .New ("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072");func (_ce *PSInteger )Duplicate ()PSObject {_gd :=PSInteger {};
_gd .Val =_ce .Val ;return &_gd };

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool (val bool )*PSBoolean {_ddb :=PSBoolean {};_ddb .Val =val ;return &_ddb };func (_gdf *PSProgram )DebugString ()string {_ebe :="\u007b\u0020";for _ ,_ba :=range *_gdf {_ebe +=_ba .DebugString ();_ebe +="\u0020";};_ebe +="\u007d";return _ebe ;
};func (_cfdd *PSOperand )ne (_fcd *PSStack )error {_ffab :=_cfdd .eq (_fcd );if _ffab !=nil {return _ffab ;};_ffab =_cfdd .not (_fcd );return _ffab ;};

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_ebad *PSStack )PopNumberAsFloat64 ()(float64 ,error ){_cgbe ,_def :=_ebad .Pop ();if _def !=nil {return 0,_def ;};if _bddg ,_bcb :=_cgbe .(*PSReal );_bcb {return _bddg .Val ,nil ;}else if _eefa ,_abg :=_cgbe .(*PSInteger );_abg {return float64 (_eefa .Val ),nil ;
}else {return 0,ErrTypeCheck ;};};func (_abb *PSOperand )dup (_gdfc *PSStack )error {_cfc ,_ffa :=_gdfc .Pop ();if _ffa !=nil {return _ffa ;};_ffa =_gdfc .Push (_cfc );if _ffa !=nil {return _ffa ;};_ffa =_gdfc .Push (_cfc .Duplicate ());return _ffa ;};
func (_gdea *PSOperand )sin (_fegd *PSStack )error {_aff ,_acef :=_fegd .PopNumberAsFloat64 ();if _acef !=nil {return _acef ;};_acg :=_ge .Sin (_aff *_ge .Pi /180.0);_acef =_fegd .Push (MakeReal (_acg ));return _acef ;};func (_eae *PSOperand )mul (_gagb *PSStack )error {_bgg ,_gdcc :=_gagb .Pop ();
if _gdcc !=nil {return _gdcc ;};_dbb ,_gdcc :=_gagb .Pop ();if _gdcc !=nil {return _gdcc ;};_ccb ,_fgd :=_bgg .(*PSReal );_ggg ,_fbga :=_bgg .(*PSInteger );if !_fgd &&!_fbga {return ErrTypeCheck ;};_dfca ,_gafe :=_dbb .(*PSReal );_gcae ,_fab :=_dbb .(*PSInteger );
if !_gafe &&!_fab {return ErrTypeCheck ;};if _fbga &&_fab {_dccf :=_ggg .Val *_gcae .Val ;_fcc :=_gagb .Push (MakeInteger (_dccf ));return _fcc ;};var _aae float64 ;if _fgd {_aae =_ccb .Val ;}else {_aae =float64 (_ggg .Val );};if _gafe {_aae *=_dfca .Val ;
}else {_aae *=float64 (_gcae .Val );};_gdcc =_gagb .Push (MakeReal (_aae ));return _gdcc ;};

// Append appends an object to the PSProgram.
func (_aac *PSProgram )Append (obj PSObject ){*_aac =append (*_aac ,obj )};func (_fcfb *PSOperand )cvr (_gfd *PSStack )error {_dbd ,_cbgd :=_gfd .Pop ();if _cbgd !=nil {return _cbgd ;};if _ceb ,_ffe :=_dbd .(*PSReal );_ffe {_cbgd =_gfd .Push (MakeReal (_ceb .Val ));
}else if _cce ,_fbg :=_dbd .(*PSInteger );_fbg {_cbgd =_gfd .Push (MakeReal (float64 (_cce .Val )));}else {return ErrTypeCheck ;};return _cbgd ;};func (_bfba *PSOperand )ln (_bebf *PSStack )error {_gaf ,_cbb :=_bebf .PopNumberAsFloat64 ();if _cbb !=nil {return _cbb ;
};_agba :=_ge .Log (_gaf );_cbb =_bebf .Push (MakeReal (_agba ));return _cbb ;};func (_db *PSReal )Duplicate ()PSObject {_aa :=PSReal {};_aa .Val =_db .Val ;return &_aa };

// String returns a string representation of the stack.
func (_efe *PSStack )String ()string {_gegfd :="\u005b\u0020";for _ ,_dgebc :=range *_efe {_gegfd +=_dgebc .String ();_gegfd +="\u0020";};_gegfd +="\u005d";return _gegfd ;};

// Pop pops an object from the top of the stack.
func (_ddf *PSStack )Pop ()(PSObject ,error ){if len (*_ddf )< 1{return nil ,ErrStackUnderflow ;};_aba :=(*_ddf )[len (*_ddf )-1];*_ddf =(*_ddf )[0:len (*_ddf )-1];return _aba ,nil ;};func (_ac *PSOperand )Duplicate ()PSObject {_gg :=*_ac ;return &_gg };
func (_bfeb *PSOperand )ifCondition (_fee *PSStack )error {_cea ,_gcf :=_fee .Pop ();if _gcf !=nil {return _gcf ;};_dgc ,_gcf :=_fee .Pop ();if _gcf !=nil {return _gcf ;};_egb ,_cge :=_cea .(*PSProgram );if !_cge {return ErrTypeCheck ;};_gbd ,_cge :=_dgc .(*PSBoolean );
if !_cge {return ErrTypeCheck ;};if _gbd .Val {_bead :=_egb .Exec (_fee );return _bead ;};return nil ;};func (_gaga *PSParser )parseBool ()(*PSBoolean ,error ){_geaa ,_eeda :=_gaga ._bbga .Peek (4);if _eeda !=nil {return MakeBool (false ),_eeda ;};if (len (_geaa )>=4)&&(string (_geaa [:4])=="\u0074\u0072\u0075\u0065"){_gaga ._bbga .Discard (4);
return MakeBool (true ),nil ;};_geaa ,_eeda =_gaga ._bbga .Peek (5);if _eeda !=nil {return MakeBool (false ),_eeda ;};if (len (_geaa )>=5)&&(string (_geaa [:5])=="\u0066\u0061\u006cs\u0065"){_gaga ._bbga .Discard (5);return MakeBool (false ),nil ;};return MakeBool (false ),_d .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");
};

// Parse parses the postscript and store as a program that can be executed.
func (_aada *PSParser )Parse ()(*PSProgram ,error ){_aada .skipSpaces ();_fdb ,_agca :=_aada ._bbga .Peek (2);if _agca !=nil {return nil ,_agca ;};if _fdb [0]!='{'{return nil ,_d .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b");
};_adb ,_agca :=_aada .parseFunction ();if _agca !=nil &&_agca !=_g .EOF {return nil ,_agca ;};return _adb ,_agca ;};func (_eeea *PSOperand )idiv (_baee *PSStack )error {_ged ,_cec :=_baee .Pop ();if _cec !=nil {return _cec ;};_agad ,_cec :=_baee .Pop ();
if _cec !=nil {return _cec ;};_agga ,_cgd :=_ged .(*PSInteger );if !_cgd {return ErrTypeCheck ;};if _agga .Val ==0{return ErrUndefinedResult ;};_cda ,_cgd :=_agad .(*PSInteger );if !_cgd {return ErrTypeCheck ;};_bcda :=_cda .Val /_agga .Val ;_cec =_baee .Push (MakeInteger (_bcda ));
return _cec ;};func (_gdg *PSOperand )ceiling (_fbf *PSStack )error {_ecg ,_bbe :=_fbf .Pop ();if _bbe !=nil {return _bbe ;};if _aad ,_ag :=_ecg .(*PSReal );_ag {_bbe =_fbf .Push (MakeReal (_ge .Ceil (_aad .Val )));}else if _cdd ,_cfd :=_ecg .(*PSInteger );
_cfd {_bbe =_fbf .Push (MakeInteger (_cdd .Val ));}else {_bbe =ErrTypeCheck ;};return _bbe ;};func (_ccag *PSOperand )log (_bdf *PSStack )error {_fgf ,_bdd :=_bdf .PopNumberAsFloat64 ();if _bdd !=nil {return _bdd ;};_eef :=_ge .Log10 (_fgf );_bdd =_bdf .Push (MakeReal (_eef ));
return _bdd ;};var ErrStackUnderflow =_d .New ("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077");func (_acb *PSOperand )le (_cdf *PSStack )error {_gff ,_eacg :=_cdf .PopNumberAsFloat64 ();if _eacg !=nil {return _eacg ;
};_dffe ,_eacg :=_cdf .PopNumberAsFloat64 ();if _eacg !=nil {return _eacg ;};if _ge .Abs (_dffe -_gff )< _b {_acf :=_cdf .Push (MakeBool (true ));return _acf ;}else if _dffe < _gff {_eggg :=_cdf .Push (MakeBool (true ));return _eggg ;}else {_fegb :=_cdf .Push (MakeBool (false ));
return _fegb ;};};func (_efg *PSOperand )neg (_cbd *PSStack )error {_cba ,_acaf :=_cbd .Pop ();if _acaf !=nil {return _acaf ;};if _baff ,_gbff :=_cba .(*PSReal );_gbff {_acaf =_cbd .Push (MakeReal (-_baff .Val ));return _acaf ;}else if _bbbd ,_gcd :=_cba .(*PSInteger );
_gcd {_acaf =_cbd .Push (MakeInteger (-_bbbd .Val ));return _acaf ;}else {return ErrTypeCheck ;};};func (_egad *PSOperand )mod (_ccf *PSStack )error {_eda ,_dbfg :=_ccf .Pop ();if _dbfg !=nil {return _dbfg ;};_abe ,_dbfg :=_ccf .Pop ();if _dbfg !=nil {return _dbfg ;
};_gfe ,_gag :=_eda .(*PSInteger );if !_gag {return ErrTypeCheck ;};if _gfe .Val ==0{return ErrUndefinedResult ;};_afca ,_gag :=_abe .(*PSInteger );if !_gag {return ErrTypeCheck ;};_ffda :=_afca .Val %_gfe .Val ;_dbfg =_ccf .Push (MakeInteger (_ffda ));
return _dbfg ;};func (_dbfc *PSOperand )ifelse (_aaag *PSStack )error {_bfbf ,_eac :=_aaag .Pop ();if _eac !=nil {return _eac ;};_feg ,_eac :=_aaag .Pop ();if _eac !=nil {return _eac ;};_gccf ,_eac :=_aaag .Pop ();if _eac !=nil {return _eac ;};_gbb ,_fceb :=_bfbf .(*PSProgram );
if !_fceb {return ErrTypeCheck ;};_ace ,_fceb :=_feg .(*PSProgram );if !_fceb {return ErrTypeCheck ;};_bdaa ,_fceb :=_gccf .(*PSBoolean );if !_fceb {return ErrTypeCheck ;};if _bdaa .Val {_geg :=_ace .Exec (_aaag );return _geg ;};_eac =_gbb .Exec (_aaag );
return _eac ;};

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger (val int )*PSInteger {_dcge :=PSInteger {};_dcge .Val =val ;return &_dcge };func (_fcfd *PSOperand )cvi (_eab *PSStack )error {_egg ,_aef :=_eab .Pop ();if _aef !=nil {return _aef ;};if _eea ,_fbc :=_egg .(*PSReal );_fbc {_aefg :=int (_eea .Val );
_aef =_eab .Push (MakeInteger (_aefg ));}else if _baf ,_gbe :=_egg .(*PSInteger );_gbe {_fed :=_baf .Val ;_aef =_eab .Push (MakeInteger (_fed ));}else {return ErrTypeCheck ;};return _aef ;};func (_bcd *PSOperand )bitshift (_dg *PSStack )error {_gdec ,_cbe :=_dg .PopInteger ();
if _cbe !=nil {return _cbe ;};_afc ,_cbe :=_dg .PopInteger ();if _cbe !=nil {return _cbe ;};var _dca int ;if _gdec >=0{_dca =_afc <<uint (_gdec );}else {_dca =_afc >>uint (-_gdec );};_cbe =_dg .Push (MakeInteger (_dca ));return _cbe ;};func (_ea *PSReal )DebugString ()string {return _f .Sprintf ("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f",_ea .Val );
};

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_ade *PSStack )PopInteger ()(int ,error ){_fcda ,_fcde :=_ade .Pop ();if _fcde !=nil {return 0,_fcde ;};if _gbde ,_dgca :=_fcda .(*PSInteger );_dgca {return _gbde .Val ,nil ;};return 0,ErrTypeCheck ;};func (_fgg *PSOperand )atan (_gdff *PSStack )error {_fad ,_fd :=_gdff .PopNumberAsFloat64 ();
if _fd !=nil {return _fd ;};_aeg ,_fd :=_gdff .PopNumberAsFloat64 ();if _fd !=nil {return _fd ;};if _fad ==0{var _dfa error ;if _aeg < 0{_dfa =_gdff .Push (MakeReal (270));}else {_dfa =_gdff .Push (MakeReal (90));};return _dfa ;};_ef :=_aeg /_fad ;_ga :=_ge .Atan (_ef )*180/_ge .Pi ;
_fd =_gdff .Push (MakeReal (_ga ));return _fd ;};func (_agag *PSOperand )exp (_dad *PSStack )error {_bec ,_gaa :=_dad .PopNumberAsFloat64 ();if _gaa !=nil {return _gaa ;};_fcfe ,_gaa :=_dad .PopNumberAsFloat64 ();if _gaa !=nil {return _gaa ;};if _ge .Abs (_bec )< 1&&_fcfe < 0{return ErrUndefinedResult ;
};_afcd :=_ge .Pow (_fcfe ,_bec );_gaa =_dad .Push (MakeReal (_afcd ));return _gaa ;};

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand (val string )*PSOperand {_gad :=PSOperand (val );return &_gad };func (_cgbc *PSOperand )sub (_abcc *PSStack )error {_fca ,_edf :=_abcc .Pop ();if _edf !=nil {return _edf ;};_dcf ,_edf :=_abcc .Pop ();if _edf !=nil {return _edf ;};_fdd ,_ccae :=_fca .(*PSReal );
_bge ,_cde :=_fca .(*PSInteger );if !_ccae &&!_cde {return ErrTypeCheck ;};_fcb ,_aefd :=_dcf .(*PSReal );_afef ,_fddd :=_dcf .(*PSInteger );if !_aefd &&!_fddd {return ErrTypeCheck ;};if _cde &&_fddd {_fcfg :=_afef .Val -_bge .Val ;_cebf :=_abcc .Push (MakeInteger (_fcfg ));
return _cebf ;};var _ebea float64 =0;if _aefd {_ebea =_fcb .Val ;}else {_ebea =float64 (_afef .Val );};if _ccae {_ebea -=_fdd .Val ;}else {_ebea -=float64 (_bge .Val );};_edf =_abcc .Push (MakeReal (_ebea ));return _edf ;};

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject ;

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array (objects []PSObject )([]float64 ,error ){var _bf []float64 ;for _ ,_fc :=range objects {if _gea ,_bfa :=_fc .(*PSInteger );_bfa {_bf =append (_bf ,float64 (_gea .Val ));}else if _ca ,_ff :=_fc .(*PSReal );_ff {_bf =append (_bf ,_ca .Val );
}else {return nil ,ErrTypeCheck ;};};return _bf ,nil ;};func (_gef *PSInteger )String ()string {return _f .Sprintf ("\u0025\u0064",_gef .Val )};func (_egga *PSOperand )gt (_dgga *PSStack )error {_agd ,_gdcf :=_dgga .PopNumberAsFloat64 ();if _gdcf !=nil {return _gdcf ;
};_eaf ,_gdcf :=_dgga .PopNumberAsFloat64 ();if _gdcf !=nil {return _gdcf ;};if _ge .Abs (_eaf -_agd )< _b {_bcc :=_dgga .Push (MakeBool (false ));return _bcc ;}else if _eaf > _agd {_fde :=_dgga .Push (MakeBool (true ));return _fde ;}else {_bac :=_dgga .Push (MakeBool (false ));
return _bac ;};};func (_gdba *PSOperand )round (_acag *PSStack )error {_bbbf ,_gab :=_acag .Pop ();if _gab !=nil {return _gab ;};if _egde ,_bgf :=_bbbf .(*PSReal );_bgf {_gab =_acag .Push (MakeReal (_ge .Floor (_egde .Val +0.5)));}else if _ceed ,_cebc :=_bbbf .(*PSInteger );
_cebc {_gab =_acag .Push (MakeInteger (_ceed .Val ));}else {return ErrTypeCheck ;};return _gab ;};func (_gbfg *PSOperand )not (_bba *PSStack )error {_dcg ,_gbed :=_bba .Pop ();if _gbed !=nil {return _gbed ;};if _ebb ,_faa :=_dcg .(*PSBoolean );_faa {_gbed =_bba .Push (MakeBool (!_ebb .Val ));
return _gbed ;}else if _dagd ,_bef :=_dcg .(*PSInteger );_bef {_gbed =_bba .Push (MakeInteger (^_dagd .Val ));return _gbed ;}else {return ErrTypeCheck ;};};var ErrTypeCheck =_d .New ("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072");


// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct{Stack *PSStack ;_cc *PSProgram ;};func (_ecc *PSProgram )Duplicate ()PSObject {_cag :=&PSProgram {};for _ ,_dbf :=range *_ecc {_cag .Append (_dbf .Duplicate ());};return _cag ;};func (_cga *PSOperand )add (_fe *PSStack )error {_ecce ,_beg :=_fe .Pop ();
if _beg !=nil {return _beg ;};_bcea ,_beg :=_fe .Pop ();if _beg !=nil {return _beg ;};_gde ,_cd :=_ecce .(*PSReal );_bfb ,_aab :=_ecce .(*PSInteger );if !_cd &&!_aab {return ErrTypeCheck ;};_cee ,_aeb :=_bcea .(*PSReal );_ab ,_bfe :=_bcea .(*PSInteger );
if !_aeb &&!_bfe {return ErrTypeCheck ;};if _aab &&_bfe {_ada :=_bfb .Val +_ab .Val ;_bda :=_fe .Push (MakeInteger (_ada ));return _bda ;};var _gdb float64 ;if _cd {_gdb =_gde .Val ;}else {_gdb =float64 (_bfb .Val );};if _aeb {_gdb +=_cee .Val ;}else {_gdb +=float64 (_ab .Val );
};_beg =_fe .Push (MakeReal (_gdb ));return _beg ;};var ErrRangeCheck =_d .New ("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072");

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal (val float64 )*PSReal {_gedc :=PSReal {};_gedc .Val =val ;return &_gedc };func (_fb *PSOperand )and (_bcf *PSStack )error {_ceeb ,_cbc :=_bcf .Pop ();if _cbc !=nil {return _cbc ;};_geac ,_cbc :=_bcf .Pop ();if _cbc !=nil {return _cbc ;};if _gfc ,_af :=_ceeb .(*PSBoolean );
_af {_beb ,_eag :=_geac .(*PSBoolean );if !_eag {return ErrTypeCheck ;};_cbc =_bcf .Push (MakeBool (_gfc .Val &&_beb .Val ));return _cbc ;};if _cfg ,_fea :=_ceeb .(*PSInteger );_fea {_fg ,_eed :=_geac .(*PSInteger );if !_eed {return ErrTypeCheck ;};_cbc =_bcf .Push (MakeInteger (_cfg .Val &_fg .Val ));
return _cbc ;};return ErrTypeCheck ;};func (_ceae *PSOperand )sqrt (_fgea *PSStack )error {_afcaa ,_gffe :=_fgea .PopNumberAsFloat64 ();if _gffe !=nil {return _gffe ;};if _afcaa < 0{return ErrRangeCheck ;};_cgdd :=_ge .Sqrt (_afcaa );_gffe =_fgea .Push (MakeReal (_cgdd ));
return _gffe ;};

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_ae *PSExecutor )Execute (objects []PSObject )([]PSObject ,error ){for _ ,_gc :=range objects {_bg :=_ae .Stack .Push (_gc );if _bg !=nil {return nil ,_bg ;};};_cbg :=_ae ._cc .Exec (_ae .Stack );if _cbg !=nil {_c .Log .Debug ("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_cbg );
return nil ,_cbg ;};_eg :=[]PSObject (*_ae .Stack );_ae .Stack .Empty ();return _eg ,nil ;};func (_dagb *PSOperand )index (_cagf *PSStack )error {_cfa ,_bbba :=_cagf .Pop ();if _bbba !=nil {return _bbba ;};_dce ,_fae :=_cfa .(*PSInteger );if !_fae {return ErrTypeCheck ;
};if _dce .Val < 0{return ErrRangeCheck ;};if _dce .Val > len (*_cagf )-1{return ErrStackUnderflow ;};_cgdc :=(*_cagf )[len (*_cagf )-1-_dce .Val ];_bbba =_cagf .Push (_cgdc .Duplicate ());return _bbba ;};func (_bd *PSInteger )DebugString ()string {return _f .Sprintf ("\u0069\u006e\u0074\u003a\u0025\u0064",_bd .Val );
};func (_cca *PSOperand )floor (_fge *PSStack )error {_ddd ,_agb :=_fge .Pop ();if _agb !=nil {return _agb ;};if _dcc ,_ebc :=_ddd .(*PSReal );_ebc {_agb =_fge .Push (MakeReal (_ge .Floor (_dcc .Val )));}else if _eaa ,_agf :=_ddd .(*PSInteger );_agf {_agb =_fge .Push (MakeInteger (_eaa .Val ));
}else {return ErrTypeCheck ;};return _agb ;};func (_ffd *PSOperand )abs (_aaa *PSStack )error {_fa ,_add :=_aaa .Pop ();if _add !=nil {return _add ;};if _fcf ,_bb :=_fa .(*PSReal );_bb {_geb :=_fcf .Val ;if _geb < 0{_add =_aaa .Push (MakeReal (-_geb ));
}else {_add =_aaa .Push (MakeReal (_geb ));};}else if _be ,_bdb :=_fa .(*PSInteger );_bdb {_aca :=_be .Val ;if _aca < 0{_add =_aaa .Push (MakeInteger (-_aca ));}else {_add =_aaa .Push (MakeInteger (_aca ));};}else {return ErrTypeCheck ;};return _add ;};
func (_aaf *PSParser )parseOperand ()(*PSOperand ,error ){var _cgbg []byte ;for {_gedg ,_aeeb :=_aaf ._bbga .Peek (1);if _aeeb !=nil {if _aeeb ==_g .EOF {break ;};return nil ,_aeeb ;};if _de .IsDelimiter (_gedg [0]){break ;};if _de .IsWhiteSpace (_gedg [0]){break ;
};_ceedg ,_ :=_aaf ._bbga .ReadByte ();_cgbg =append (_cgbg ,_ceedg );};if len (_cgbg )==0{return nil ,_d .New ("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029");};return MakeOperand (string (_cgbg )),nil ;
};var ErrUnsupportedOperand =_d .New ("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064");func (_dda *PSOperand )xor (_bbgf *PSStack )error {_dea ,_afae :=_bbgf .Pop ();if _afae !=nil {return _afae ;
};_dcd ,_afae :=_bbgf .Pop ();if _afae !=nil {return _afae ;};if _cac ,_edgf :=_dea .(*PSBoolean );_edgf {_dccd ,_adf :=_dcd .(*PSBoolean );if !_adf {return ErrTypeCheck ;};_afae =_bbgf .Push (MakeBool (_cac .Val !=_dccd .Val ));return _afae ;};if _edfb ,_eaeg :=_dea .(*PSInteger );
_eaeg {_agc ,_ege :=_dcd .(*PSInteger );if !_ege {return ErrTypeCheck ;};_afae =_bbgf .Push (MakeInteger (_edfb .Val ^_agc .Val ));return _afae ;};return ErrTypeCheck ;};func (_gaae *PSParser )parseFunction ()(*PSProgram ,error ){_ffg ,_ :=_gaae ._bbga .ReadByte ();
if _ffg !='{'{return nil ,_d .New ("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");};_dcdd :=NewPSProgram ();for {_gaae .skipSpaces ();_cfab ,_dfe :=_gaae ._bbga .Peek (2);if _dfe !=nil {if _dfe ==_g .EOF {break ;
};return nil ,_dfe ;};_c .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_cfab ));if _cfab [0]=='}'{_c .Log .Trace ("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");_gaae ._bbga .ReadByte ();
break ;}else if _cfab [0]=='{'{_c .Log .Trace ("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!");_ebca ,_fga :=_gaae .parseFunction ();if _fga !=nil {return nil ,_fga ;};_dcdd .Append (_ebca );}else if _de .IsDecimalDigit (_cfab [0])||(_cfab [0]=='-'&&_de .IsDecimalDigit (_cfab [1])){_c .Log .Trace ("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!");
_dcab ,_bbf :=_gaae .parseNumber ();if _bbf !=nil {return nil ,_bbf ;};_dcdd .Append (_dcab );}else {_c .Log .Trace ("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f");_cfab ,_ =_gaae ._bbga .Peek (5);_ccdf :=string (_cfab );
_c .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_ccdf );if (len (_ccdf )> 4)&&(_ccdf [:5]=="\u0066\u0061\u006cs\u0065"){_bde ,_aee :=_gaae .parseBool ();if _aee !=nil {return nil ,_aee ;};_dcdd .Append (_bde );}else if (len (_ccdf )> 3)&&(_ccdf [:4]=="\u0074\u0072\u0075\u0065"){_dcca ,_acec :=_gaae .parseBool ();
if _acec !=nil {return nil ,_acec ;};_dcdd .Append (_dcca );}else {_cef ,_cdec :=_gaae .parseOperand ();if _cdec !=nil {return nil ,_cdec ;};_dcdd .Append (_cef );};};};return _dcdd ,nil ;};

// PSParser is a basic Postscript parser.
type PSParser struct{_bbga *_ec .Reader };func (_cad *PSParser )parseNumber ()(PSObject ,error ){_cbce ,_ffdf :=_de .ParseNumber (_cad ._bbga );if _ffdf !=nil {return nil ,_ffdf ;};switch _acad :=_cbce .(type ){case *_de .PdfObjectFloat :return MakeReal (float64 (*_acad )),nil ;
case *_de .PdfObjectInteger :return MakeInteger (int (*_acad )),nil ;};return nil ,_f .Errorf ("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054",_cbce );};func (_df *PSBoolean )String ()string {return _f .Sprintf ("\u0025\u0076",_df .Val )};
func (_ded *PSOperand )exch (_gcc *PSStack )error {_cgb ,_fff :=_gcc .Pop ();if _fff !=nil {return _fff ;};_adcf ,_fff :=_gcc .Pop ();if _fff !=nil {return _fff ;};_fff =_gcc .Push (_cgb );if _fff !=nil {return _fff ;};_fff =_gcc .Push (_adcf );return _fff ;
};

// PSObject represents a postscript object.
type PSObject interface{

// Duplicate makes a fresh copy of the PSObject.
Duplicate ()PSObject ;

// DebugString returns a descriptive representation of the PSObject with more information than String()
// for debugging purposes.
DebugString ()string ;

// String returns a string representation of the PSObject.
String ()string ;};var ErrStackOverflow =_d .New ("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077");func (_gb *PSOperand )DebugString ()string {return _f .Sprintf ("\u006fp\u003a\u0027\u0025\u0073\u0027",*_gb );};

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject ;func _aefb (_gfff int )int {if _gfff < 0{return -_gfff ;};return _gfff ;};

// Exec executes the program, typically leaving output values on the stack.
func (_bce *PSProgram )Exec (stack *PSStack )error {for _ ,_ced :=range *_bce {var _ffcg error ;switch _gf :=_ced .(type ){case *PSInteger :_aea :=_gf ;_ffcg =stack .Push (_aea );case *PSReal :_dfc :=_gf ;_ffcg =stack .Push (_dfc );case *PSBoolean :_cf :=_gf ;
_ffcg =stack .Push (_cf );case *PSProgram :_bae :=_gf ;_ffcg =stack .Push (_bae );case *PSOperand :_dc :=_gf ;_ffcg =_dc .Exec (stack );default:return ErrTypeCheck ;};if _ffcg !=nil {return _ffcg ;};};return nil ;};func (_da *PSOperand )cos (_eff *PSStack )error {_eee ,_bdbe :=_eff .PopNumberAsFloat64 ();
if _bdbe !=nil {return _bdbe ;};_fef :=_ge .Cos (_eee *_ge .Pi /180.0);_bdbe =_eff .Push (MakeReal (_fef ));return _bdbe ;};func (_afe *PSOperand )pop (_abfg *PSStack )error {_ ,_dge :=_abfg .Pop ();if _dge !=nil {return _dge ;};return nil ;};func (_bbb *PSOperand )eq (_aga *PSStack )error {_ccc ,_eede :=_aga .Pop ();
if _eede !=nil {return _eede ;};_baba ,_eede :=_aga .Pop ();if _eede !=nil {return _eede ;};_bbg ,_eca :=_ccc .(*PSBoolean );_gdgb ,_ccg :=_baba .(*PSBoolean );if _eca ||_ccg {var _gbf error ;if _eca &&_ccg {_gbf =_aga .Push (MakeBool (_bbg .Val ==_gdgb .Val ));
}else {_gbf =_aga .Push (MakeBool (false ));};return _gbf ;};var _adaf float64 ;var _ega float64 ;if _abc ,_cgaf :=_ccc .(*PSInteger );_cgaf {_adaf =float64 (_abc .Val );}else if _afg ,_gbg :=_ccc .(*PSReal );_gbg {_adaf =_afg .Val ;}else {return ErrTypeCheck ;
};if _adc ,_ece :=_baba .(*PSInteger );_ece {_ega =float64 (_adc .Val );}else if _cdb ,_afgb :=_baba .(*PSReal );_afgb {_ega =_cdb .Val ;}else {return ErrTypeCheck ;};if _ge .Abs (_ega -_adaf )< _b {_eede =_aga .Push (MakeBool (true ));}else {_eede =_aga .Push (MakeBool (false ));
};return _eede ;};func (_ee *PSBoolean )Duplicate ()PSObject {_ad :=PSBoolean {};_ad .Val =_ee .Val ;return &_ad };

// Push pushes an object on top of the stack.
func (_ffag *PSStack )Push (obj PSObject )error {if len (*_ffag )> 100{return ErrStackOverflow ;};*_ffag =append (*_ffag ,obj );return nil ;};func (_abea *PSOperand )truncate (_dfg *PSStack )error {_cedg ,_feb :=_dfg .Pop ();if _feb !=nil {return _feb ;
};if _dgeb ,_fefd :=_cedg .(*PSReal );_fefd {_dgf :=int (_dgeb .Val );_feb =_dfg .Push (MakeReal (float64 (_dgf )));}else if _ccd ,_cbbe :=_cedg .(*PSInteger );_cbbe {_feb =_dfg .Push (MakeInteger (_ccd .Val ));}else {return ErrTypeCheck ;};return _feb ;
};func (_dae *PSOperand )div (_adda *PSStack )error {_fce ,_bab :=_adda .Pop ();if _bab !=nil {return _bab ;};_bea ,_bab :=_adda .Pop ();if _bab !=nil {return _bab ;};_bebb ,_bag :=_fce .(*PSReal );_geaf ,_bcdd :=_fce .(*PSInteger );if !_bag &&!_bcdd {return ErrTypeCheck ;
};if _bag &&_bebb .Val ==0{return ErrUndefinedResult ;};if _bcdd &&_geaf .Val ==0{return ErrUndefinedResult ;};_bfee ,_fdc :=_bea .(*PSReal );_abf ,_gded :=_bea .(*PSInteger );if !_fdc &&!_gded {return ErrTypeCheck ;};var _gca float64 ;if _fdc {_gca =_bfee .Val ;
}else {_gca =float64 (_abf .Val );};if _bag {_gca /=_bebb .Val ;}else {_gca /=float64 (_geaf .Val );};_bab =_adda .Push (MakeReal (_gca ));return _bab ;};

// NewPSStack returns an initialized PSStack.
func NewPSStack ()*PSStack {return &PSStack {}};

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_gabb *PSStack )DebugString ()string {_defe :="\u005b\u0020";for _ ,_ddaf :=range *_gabb {_defe +=_ddaf .DebugString ();_defe +="\u0020";};_defe +="\u005d";return _defe ;};