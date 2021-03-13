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
package ps ;import (_ed "bufio";_b "bytes";_e "errors";_g "fmt";_dc "github.com/unidoc/unipdf/v3/common";_bb "github.com/unidoc/unipdf/v3/core";_c "io";_d "math";);func (_aeg *PSOperand )index (_dgg *PSStack )error {_dcac ,_dbg :=_dgg .Pop ();if _dbg !=nil {return _dbg ;
};_fecb ,_cgbg :=_dcac .(*PSInteger );if !_cgbg {return ErrTypeCheck ;};if _fecb .Val < 0{return ErrRangeCheck ;};if _fecb .Val > len (*_dgg )-1{return ErrStackUnderflow ;};_edda :=(*_dgg )[len (*_dgg )-1-_fecb .Val ];_dbg =_dgg .Push (_edda .Duplicate ());
return _dbg ;};func (_bd *PSInteger )DebugString ()string {return _g .Sprintf ("\u0069\u006e\u0074\u003a\u0025\u0064",_bd .Val );};

// Exec executes the program, typically leaving output values on the stack.
func (_eeb *PSProgram )Exec (stack *PSStack )error {for _ ,_bg :=range *_eeb {var _ggc error ;switch _da :=_bg .(type ){case *PSInteger :_daf :=_da ;_ggc =stack .Push (_daf );case *PSReal :_ef :=_da ;_ggc =stack .Push (_ef );case *PSBoolean :_abe :=_da ;
_ggc =stack .Push (_abe );case *PSProgram :_cca :=_da ;_ggc =stack .Push (_cca );case *PSOperand :_gd :=_da ;_ggc =_gd .Exec (stack );default:return ErrTypeCheck ;};if _ggc !=nil {return _ggc ;};};return nil ;};

// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal (val float64 )*PSReal {_dab :=PSReal {};_dab .Val =val ;return &_dab };func (_beb *PSInteger )String ()string {return _g .Sprintf ("\u0025\u0064",_beb .Val )};

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array (objects []PSObject )([]float64 ,error ){var _ad []float64 ;for _ ,_bbd :=range objects {if _edg ,_eb :=_bbd .(*PSInteger );_eb {_ad =append (_ad ,float64 (_edg .Val ));}else if _fa ,_gg :=_bbd .(*PSReal );_gg {_ad =append (_ad ,_fa .Val );
}else {return nil ,ErrTypeCheck ;};};return _ad ,nil ;};func (_dda *PSOperand )ne (_bba *PSStack )error {_ageb :=_dda .eq (_bba );if _ageb !=nil {return _ageb ;};_ageb =_dda .not (_bba );return _ageb ;};

// PSReal represents a real number.
type PSReal struct{Val float64 ;};func (_aa *PSBoolean )String ()string {return _g .Sprintf ("\u0025\u0076",_aa .Val )};func (_acd *PSParser )skipSpaces ()(int ,error ){_ebe :=0;for {_bfc ,_ffbe :=_acd ._cbd .Peek (1);if _ffbe !=nil {return 0,_ffbe ;};
if _bb .IsWhiteSpace (_bfc [0]){_acd ._cbd .ReadByte ();_ebe ++;}else {break ;};};return _ebe ,nil ;};func (_cdf *PSOperand )le (_gcg *PSStack )error {_cea ,_feed :=_gcg .PopNumberAsFloat64 ();if _feed !=nil {return _feed ;};_bef ,_feed :=_gcg .PopNumberAsFloat64 ();
if _feed !=nil {return _feed ;};if _d .Abs (_bef -_cea )< _f {_ade :=_gcg .Push (MakeBool (true ));return _ade ;}else if _bef < _cea {_abge :=_gcg .Push (MakeBool (true ));return _abge ;}else {_gdac :=_gcg .Push (MakeBool (false ));return _gdac ;};};func (_gedf *PSParser )parseNumber ()(PSObject ,error ){_gbbb ,_gcae :=_bb .ParseNumber (_gedf ._cbd );
if _gcae !=nil {return nil ,_gcae ;};switch _ebed :=_gbbb .(type ){case *_bb .PdfObjectFloat :return MakeReal (float64 (*_ebed )),nil ;case *_bb .PdfObjectInteger :return MakeInteger (int (*_ebed )),nil ;};return nil ,_g .Errorf ("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054",_gbbb );
};func (_ab *PSReal )Duplicate ()PSObject {_ac :=PSReal {};_ac .Val =_ab .Val ;return &_ac };

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject ;func (_cd *PSOperand )DebugString ()string {return _g .Sprintf ("\u006fp\u003a\u0027\u0025\u0073\u0027",*_cd );};func (_dcag *PSOperand )gt (_fcb *PSStack )error {_fgee ,_cfd :=_fcb .PopNumberAsFloat64 ();if _cfd !=nil {return _cfd ;
};_fcba ,_cfd :=_fcb .PopNumberAsFloat64 ();if _cfd !=nil {return _cfd ;};if _d .Abs (_fcba -_fgee )< _f {_eace :=_fcb .Push (MakeBool (false ));return _eace ;}else if _fcba > _fgee {_efd :=_fcb .Push (MakeBool (true ));return _efd ;}else {_gbbg :=_fcb .Push (MakeBool (false ));
return _gbbg ;};};var ErrTypeCheck =_e .New ("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072");func (_befg *PSOperand )not (_agbd *PSStack )error {_ebcd ,_fef :=_agbd .Pop ();if _fef !=nil {return _fef ;};if _ffd ,_bda :=_ebcd .(*PSBoolean );
_bda {_fef =_agbd .Push (MakeBool (!_ffd .Val ));return _fef ;}else if _cac ,_dbgd :=_ebcd .(*PSInteger );_dbgd {_fef =_agbd .Push (MakeInteger (^_cac .Val ));return _fef ;}else {return ErrTypeCheck ;};};

// PSInteger represents an integer.
type PSInteger struct{Val int ;};

// PSParser is a basic Postscript parser.
type PSParser struct{_cbd *_ed .Reader };func (_ffad *PSOperand )log (_agdf *PSStack )error {_cdd ,_aae :=_agdf .PopNumberAsFloat64 ();if _aae !=nil {return _aae ;};_ebb :=_d .Log10 (_cdd );_aae =_agdf .Push (MakeReal (_ebb ));return _aae ;};func (_aab *PSOperand )eq (_agef *PSStack )error {_bca ,_agg :=_agef .Pop ();
if _agg !=nil {return _agg ;};_def ,_agg :=_agef .Pop ();if _agg !=nil {return _agg ;};_aef ,_ece :=_bca .(*PSBoolean );_bbda ,_fag :=_def .(*PSBoolean );if _ece ||_fag {var _bff error ;if _ece &&_fag {_bff =_agef .Push (MakeBool (_aef .Val ==_bbda .Val ));
}else {_bff =_agef .Push (MakeBool (false ));};return _bff ;};var _ceb float64 ;var _cage float64 ;if _ded ,_bgdc :=_bca .(*PSInteger );_bgdc {_ceb =float64 (_ded .Val );}else if _ebd ,_bcg :=_bca .(*PSReal );_bcg {_ceb =_ebd .Val ;}else {return ErrTypeCheck ;
};if _ffbf ,_dgag :=_def .(*PSInteger );_dgag {_cage =float64 (_ffbf .Val );}else if _bee ,_gbb :=_def .(*PSReal );_gbb {_cage =_bee .Val ;}else {return ErrTypeCheck ;};if _d .Abs (_cage -_ceb )< _f {_agg =_agef .Push (MakeBool (true ));}else {_agg =_agef .Push (MakeBool (false ));
};return _agg ;};func (_efg *PSOperand )bitshift (_adc *PSStack )error {_afgb ,_ecc :=_adc .PopInteger ();if _ecc !=nil {return _ecc ;};_eaa ,_ecc :=_adc .PopInteger ();if _ecc !=nil {return _ecc ;};var _cad int ;if _afgb >=0{_cad =_eaa <<uint (_afgb );
}else {_cad =_eaa >>uint (-_afgb );};_ecc =_adc .Push (MakeInteger (_cad ));return _ecc ;};

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram ()*PSProgram {return &PSProgram {}};func (_ged *PSOperand )ifCondition (_dcaf *PSStack )error {_caeg ,_bfb :=_dcaf .Pop ();if _bfb !=nil {return _bfb ;};_bgaa ,_bfb :=_dcaf .Pop ();if _bfb !=nil {return _bfb ;};_bbdd ,_bgb :=_caeg .(*PSProgram );
if !_bgb {return ErrTypeCheck ;};_fd ,_bgb :=_bgaa .(*PSBoolean );if !_bgb {return ErrTypeCheck ;};if _fd .Val {_fed :=_bbdd .Exec (_dcaf );return _fed ;};return nil ;};func (_ace *PSOperand )String ()string {return string (*_ace )};func (_geef *PSOperand )div (_agb *PSStack )error {_cae ,_fec :=_agb .Pop ();
if _fec !=nil {return _fec ;};_gda ,_fec :=_agb .Pop ();if _fec !=nil {return _fec ;};_bbf ,_agd :=_cae .(*PSReal );_fafg ,_bgd :=_cae .(*PSInteger );if !_agd &&!_bgd {return ErrTypeCheck ;};if _agd &&_bbf .Val ==0{return ErrUndefinedResult ;};if _bgd &&_fafg .Val ==0{return ErrUndefinedResult ;
};_age ,_afd :=_gda .(*PSReal );_ced ,_accg :=_gda .(*PSInteger );if !_afd &&!_accg {return ErrTypeCheck ;};var _bfa float64 ;if _afd {_bfa =_age .Val ;}else {_bfa =float64 (_ced .Val );};if _agd {_bfa /=_bbf .Val ;}else {_bfa /=float64 (_fafg .Val );};
_fec =_agb .Push (MakeReal (_bfa ));return _fec ;};const _f =0.000001;func (_bea *PSParser )parseFunction ()(*PSProgram ,error ){_bdgf ,_ :=_bea ._cbd .ReadByte ();if _bdgf !='{'{return nil ,_e .New ("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_cedb :=NewPSProgram ();for {_bea .skipSpaces ();_bbc ,_bedf :=_bea ._cbd .Peek (2);if _bedf !=nil {if _bedf ==_c .EOF {break ;};return nil ,_bedf ;};_dc .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_bbc ));
if _bbc [0]=='}'{_dc .Log .Trace ("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");_bea ._cbd .ReadByte ();break ;}else if _bbc [0]=='{'{_dc .Log .Trace ("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!");_ead ,_ccgb :=_bea .parseFunction ();
if _ccgb !=nil {return nil ,_ccgb ;};_cedb .Append (_ead );}else if _bb .IsDecimalDigit (_bbc [0])||(_bbc [0]=='-'&&_bb .IsDecimalDigit (_bbc [1])){_dc .Log .Trace ("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!");_eacb ,_dcfe :=_bea .parseNumber ();if _dcfe !=nil {return nil ,_dcfe ;
};_cedb .Append (_eacb );}else {_dc .Log .Trace ("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f");_bbc ,_ =_bea ._cbd .Peek (5);_fgcg :=string (_bbc );_dc .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_fgcg );
if (len (_fgcg )> 4)&&(_fgcg [:5]=="\u0066\u0061\u006cs\u0065"){_dedc ,_febc :=_bea .parseBool ();if _febc !=nil {return nil ,_febc ;};_cedb .Append (_dedc );}else if (len (_fgcg )> 3)&&(_fgcg [:4]=="\u0074\u0072\u0075\u0065"){_feec ,_adbd :=_bea .parseBool ();
if _adbd !=nil {return nil ,_adbd ;};_cedb .Append (_feec );}else {_bacb ,_fbd :=_bea .parseOperand ();if _fbd !=nil {return nil ,_fbd ;};_cedb .Append (_bacb );};};};return _cedb ,nil ;};func (_geb *PSOperand )ceiling (_abg *PSStack )error {_bga ,_ccag :=_abg .Pop ();
if _ccag !=nil {return _ccag ;};if _fgf ,_ede :=_bga .(*PSReal );_ede {_ccag =_abg .Push (MakeReal (_d .Ceil (_fgf .Val )));}else if _aga ,_dbe :=_bga .(*PSInteger );_dbe {_ccag =_abg .Push (MakeInteger (_aga .Val ));}else {_ccag =ErrTypeCheck ;};return _ccag ;
};func (_cg *PSOperand )Duplicate ()PSObject {_db :=*_cg ;return &_db };func (_afdb *PSOperand )lt (_ffeg *PSStack )error {_dgb ,_aba :=_ffeg .PopNumberAsFloat64 ();if _aba !=nil {return _aba ;};_dea ,_aba :=_ffeg .PopNumberAsFloat64 ();if _aba !=nil {return _aba ;
};if _d .Abs (_dea -_dgb )< _f {_afb :=_ffeg .Push (MakeBool (false ));return _afb ;}else if _dea < _dgb {_fbaa :=_ffeg .Push (MakeBool (true ));return _fbaa ;}else {_dfc :=_ffeg .Push (MakeBool (false ));return _dfc ;};};var ErrStackOverflow =_e .New ("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077");
func (_fbe *PSProgram )Duplicate ()PSObject {_fe :=&PSProgram {};for _ ,_abb :=range *_fbe {_fe .Append (_abb .Duplicate ());};return _fe ;};

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger (val int )*PSInteger {_aafa :=PSInteger {};_aafa .Val =val ;return &_aafa };

// NewPSStack returns an initialized PSStack.
func NewPSStack ()*PSStack {return &PSStack {}};func (_ebc *PSOperand )ln (_faeg *PSStack )error {_fecce ,_aaf :=_faeg .PopNumberAsFloat64 ();if _aaf !=nil {return _aaf ;};_eaeg :=_d .Log (_fecce );_aaf =_faeg .Push (MakeReal (_eaeg ));return _aaf ;};func _bfg (_gff int )int {if _gff < 0{return -_gff ;
};return _gff ;};func (_ea *PSOperand )add (_bag *PSStack )error {_fgd ,_fee :=_bag .Pop ();if _fee !=nil {return _fee ;};_fab ,_fee :=_bag .Pop ();if _fee !=nil {return _fee ;};_eg ,_ff :=_fgd .(*PSReal );_gc ,_faf :=_fgd .(*PSInteger );if !_ff &&!_faf {return ErrTypeCheck ;
};_gf ,_dfe :=_fab .(*PSReal );_bf ,_aeb :=_fab .(*PSInteger );if !_dfe &&!_aeb {return ErrTypeCheck ;};if _faf &&_aeb {_fea :=_gc .Val +_bf .Val ;_fad :=_bag .Push (MakeInteger (_fea ));return _fad ;};var _dbc float64 ;if _ff {_dbc =_eg .Val ;}else {_dbc =float64 (_gc .Val );
};if _dfe {_dbc +=_gf .Val ;}else {_dbc +=float64 (_bf .Val );};_fee =_bag .Push (MakeReal (_dbc ));return _fee ;};func (_bac *PSOperand )ge (_dgf *PSStack )error {_bgdg ,_efac :=_dgf .PopNumberAsFloat64 ();if _efac !=nil {return _efac ;};_ffe ,_efac :=_dgf .PopNumberAsFloat64 ();
if _efac !=nil {return _efac ;};if _d .Abs (_ffe -_bgdg )< _f {_cbg :=_dgf .Push (MakeBool (true ));return _cbg ;}else if _ffe > _bgdg {_fac :=_dgf .Push (MakeBool (true ));return _fac ;}else {_ggebf :=_dgf .Push (MakeBool (false ));return _ggebf ;};};
var ErrStackUnderflow =_e .New ("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077");

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor (program *PSProgram )*PSExecutor {_cc :=&PSExecutor {};_cc .Stack =NewPSStack ();_cc ._ga =program ;return _cc ;};func (_bdf *PSOperand )roll (_baa *PSStack )error {_dcaeg ,_gebd :=_baa .Pop ();if _gebd !=nil {return _gebd ;};_eebc ,_gebd :=_baa .Pop ();
if _gebd !=nil {return _gebd ;};_gec ,_fgba :=_dcaeg .(*PSInteger );if !_fgba {return ErrTypeCheck ;};_ebf ,_fgba :=_eebc .(*PSInteger );if !_fgba {return ErrTypeCheck ;};if _ebf .Val < 0{return ErrRangeCheck ;};if _ebf .Val ==0||_ebf .Val ==1{return nil ;
};if _ebf .Val > len (*_baa ){return ErrStackUnderflow ;};for _gagd :=0;_gagd < _bfg (_gec .Val );_gagd ++{var _dac []PSObject ;_dac =(*_baa )[len (*_baa )-(_ebf .Val ):len (*_baa )];if _gec .Val > 0{_cab :=_dac [len (_dac )-1];_dac =append ([]PSObject {_cab },_dac [0:len (_dac )-1]...);
}else {_adff :=_dac [len (_dac )-_ebf .Val ];_dac =append (_dac [1:],_adff );};_bded :=append ((*_baa )[0:len (*_baa )-_ebf .Val ],_dac ...);_baa =&_bded ;};return nil ;};

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool (val bool )*PSBoolean {_decf :=PSBoolean {};_decf .Val =val ;return &_decf };var ErrUnsupportedOperand =_e .New ("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064");var ErrUndefinedResult =_e .New ("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072");
func (_fcd *PSOperand )idiv (_abga *PSStack )error {_ggga ,_beeg :=_abga .Pop ();if _beeg !=nil {return _beeg ;};_gdga ,_beeg :=_abga .Pop ();if _beeg !=nil {return _beeg ;};_afa ,_fafb :=_ggga .(*PSInteger );if !_fafb {return ErrTypeCheck ;};if _afa .Val ==0{return ErrUndefinedResult ;
};_ffa ,_fafb :=_gdga .(*PSInteger );if !_fafb {return ErrTypeCheck ;};_egfa :=_ffa .Val /_afa .Val ;_beeg =_abga .Push (MakeInteger (_egfa ));return _beeg ;};func (_feea *PSOperand )mul (_eddd *PSStack )error {_bed ,_afdbd :=_eddd .Pop ();if _afdbd !=nil {return _afdbd ;
};_bec ,_afdbd :=_eddd .Pop ();if _afdbd !=nil {return _afdbd ;};_cga ,_ccg :=_bed .(*PSReal );_cfe ,_bfe :=_bed .(*PSInteger );if !_ccg &&!_bfe {return ErrTypeCheck ;};_beef ,_ccae :=_bec .(*PSReal );_bbdc ,_edb :=_bec .(*PSInteger );if !_ccae &&!_edb {return ErrTypeCheck ;
};if _bfe &&_edb {_cfea :=_cfe .Val *_bbdc .Val ;_dcdc :=_eddd .Push (MakeInteger (_cfea ));return _dcdc ;};var _gcf float64 ;if _ccg {_gcf =_cga .Val ;}else {_gcf =float64 (_cfe .Val );};if _ccae {_gcf *=_beef .Val ;}else {_gcf *=float64 (_bbdc .Val );
};_afdbd =_eddd .Push (MakeReal (_gcf ));return _afdbd ;};

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser (content []byte )*PSParser {_eebb :=PSParser {};_bffb :=_b .NewBuffer (content );_eebb ._cbd =_ed .NewReader (_bffb );return &_eebb ;};func (_dafb *PSOperand )mod (_dgad *PSStack )error {_ceaa ,_bfd :=_dgad .Pop ();if _bfd !=nil {return _bfd ;
};_gfb ,_bfd :=_dgad .Pop ();if _bfd !=nil {return _bfd ;};_abd ,_cffc :=_ceaa .(*PSInteger );if !_cffc {return ErrTypeCheck ;};if _abd .Val ==0{return ErrUndefinedResult ;};_daaf ,_cffc :=_gfb .(*PSInteger );if !_cffc {return ErrTypeCheck ;};_afeb :=_daaf .Val %_abd .Val ;
_bfd =_dgad .Push (MakeInteger (_afeb ));return _bfd ;};func (_cge *PSParser )parseOperand ()(*PSOperand ,error ){var _caeba []byte ;for {_aada ,_fcc :=_cge ._cbd .Peek (1);if _fcc !=nil {if _fcc ==_c .EOF {break ;};return nil ,_fcc ;};if _bb .IsDelimiter (_aada [0]){break ;
};if _bb .IsWhiteSpace (_aada [0]){break ;};_adfg ,_ :=_cge ._cbd .ReadByte ();_caeba =append (_caeba ,_adfg );};if len (_caeba )==0{return nil ,_e .New ("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029");
};return MakeOperand (string (_caeba )),nil ;};var ErrRangeCheck =_e .New ("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072");func (_fb *PSProgram )String ()string {_af :="\u007b\u0020";for _ ,_afg :=range *_fb {_af +=_afg .String ();
_af +="\u0020";};_af +="\u007d";return _af ;};

// PSOperand represents a Postscript operand (text string).
type PSOperand string ;func (_dgc *PSOperand )round (_cba *PSStack )error {_fdd ,_eab :=_cba .Pop ();if _eab !=nil {return _eab ;};if _dbd ,_eag :=_fdd .(*PSReal );_eag {_eab =_cba .Push (MakeReal (_d .Floor (_dbd .Val +0.5)));}else if _cfdf ,_ecg :=_fdd .(*PSInteger );
_ecg {_eab =_cba .Push (MakeInteger (_cfdf .Val ));}else {return ErrTypeCheck ;};return _eab ;};func (_fgc *PSOperand )sub (_geee *PSStack )error {_fdc ,_dgfe :=_geee .Pop ();if _dgfe !=nil {return _dgfe ;};_acf ,_dgfe :=_geee .Pop ();if _dgfe !=nil {return _dgfe ;
};_efb ,_gba :=_fdc .(*PSReal );_acb ,_ggef :=_fdc .(*PSInteger );if !_gba &&!_ggef {return ErrTypeCheck ;};_aad ,_gdc :=_acf .(*PSReal );_ebgb ,_afae :=_acf .(*PSInteger );if !_gdc &&!_afae {return ErrTypeCheck ;};if _ggef &&_afae {_eaed :=_ebgb .Val -_acb .Val ;
_dae :=_geee .Push (MakeInteger (_eaed ));return _dae ;};var _bcag float64 =0;if _gdc {_bcag =_aad .Val ;}else {_bcag =float64 (_ebgb .Val );};if _gba {_bcag -=_efb .Val ;}else {_bcag -=float64 (_acb .Val );};_dgfe =_geee .Push (MakeReal (_bcag ));return _dgfe ;
};func (_aff *PSOperand )neg (_fcf *PSStack )error {_gdgad ,_ceed :=_fcf .Pop ();if _ceed !=nil {return _ceed ;};if _fdf ,_dbcf :=_gdgad .(*PSReal );_dbcf {_ceed =_fcf .Push (MakeReal (-_fdf .Val ));return _ceed ;}else if _caec ,_bffc :=_gdgad .(*PSInteger );
_bffc {_ceed =_fcf .Push (MakeInteger (-_caec .Val ));return _ceed ;}else {return ErrTypeCheck ;};};func (_bc *PSOperand )abs (_cag *PSStack )error {_fae ,_gge :=_cag .Pop ();if _gge !=nil {return _gge ;};if _cbe ,_adf :=_fae .(*PSReal );_adf {_bgg :=_cbe .Val ;
if _bgg < 0{_gge =_cag .Push (MakeReal (-_bgg ));}else {_gge =_cag .Push (MakeReal (_bgg ));};}else if _feb ,_fba :=_fae .(*PSInteger );_fba {_cgb :=_feb .Val ;if _cgb < 0{_gge =_cag .Push (MakeInteger (-_cgb ));}else {_gge =_cag .Push (MakeInteger (_cgb ));
};}else {return ErrTypeCheck ;};return _gge ;};func (_afgg *PSOperand )pop (_dcf *PSStack )error {_ ,_fcbd :=_dcf .Pop ();if _fcbd !=nil {return _fcbd ;};return nil ;};

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_bfdd *PSStack )DebugString ()string {_faec :="\u005b\u0020";for _ ,_edc :=range *_bfdd {_faec +=_edc .DebugString ();_faec +="\u0020";};_faec +="\u005d";return _faec ;};

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_fbde *PSStack )PopInteger ()(int ,error ){_acdg ,_eaac :=_fbde .Pop ();if _eaac !=nil {return 0,_eaac ;};if _fde ,_agfa :=_acdg .(*PSInteger );_agfa {return _fde .Val ,nil ;};return 0,ErrTypeCheck ;};func (_dfaa *PSOperand )cos (_dcd *PSStack )error {_bce ,_ggeb :=_dcd .PopNumberAsFloat64 ();
if _ggeb !=nil {return _ggeb ;};_egg :=_d .Cos (_bce *_d .Pi /180.0);_ggeb =_dcd .Push (MakeReal (_egg ));return _ggeb ;};

// Empty empties the stack.
func (_deaa *PSStack )Empty (){*_deaa =[]PSObject {}};func (_caf *PSOperand )exp (_adb *PSStack )error {_cf ,_dfd :=_adb .PopNumberAsFloat64 ();if _dfd !=nil {return _dfd ;};_fgaa ,_dfd :=_adb .PopNumberAsFloat64 ();if _dfd !=nil {return _dfd ;};if _d .Abs (_cf )< 1&&_fgaa < 0{return ErrUndefinedResult ;
};_edd :=_d .Pow (_fgaa ,_cf );_dfd =_adb .Push (MakeReal (_edd ));return _dfd ;};

// Pop pops an object from the top of the stack.
func (_edddc *PSStack )Pop ()(PSObject ,error ){if len (*_edddc )< 1{return nil ,ErrStackUnderflow ;};_bcf :=(*_edddc )[len (*_edddc )-1];*_edddc =(*_edddc )[0:len (*_edddc )-1];return _bcf ,nil ;};func (_deb *PSBoolean )Duplicate ()PSObject {_add :=PSBoolean {};
_add .Val =_deb .Val ;return &_add };func (_dcdd *PSOperand )xor (_ccbd *PSStack )error {_dge ,_eceg :=_ccbd .Pop ();if _eceg !=nil {return _eceg ;};_ecgd ,_eceg :=_ccbd .Pop ();if _eceg !=nil {return _eceg ;};if _gaef ,_cebf :=_dge .(*PSBoolean );_cebf {_cffca ,_gca :=_ecgd .(*PSBoolean );
if !_gca {return ErrTypeCheck ;};_eceg =_ccbd .Push (MakeBool (_gaef .Val !=_cffca .Val ));return _eceg ;};if _efgc ,_faff :=_dge .(*PSInteger );_faff {_aegg ,_feff :=_ecgd .(*PSInteger );if !_feff {return ErrTypeCheck ;};_eceg =_ccbd .Push (MakeInteger (_efgc .Val ^_aegg .Val ));
return _eceg ;};return ErrTypeCheck ;};func (_deg *PSParser )parseBool ()(*PSBoolean ,error ){_bedc ,_cagec :=_deg ._cbd .Peek (4);if _cagec !=nil {return MakeBool (false ),_cagec ;};if (len (_bedc )>=4)&&(string (_bedc [:4])=="\u0074\u0072\u0075\u0065"){_deg ._cbd .Discard (4);
return MakeBool (true ),nil ;};_bedc ,_cagec =_deg ._cbd .Peek (5);if _cagec !=nil {return MakeBool (false ),_cagec ;};if (len (_bedc )>=5)&&(string (_bedc [:5])=="\u0066\u0061\u006cs\u0065"){_deg ._cbd .Discard (5);return MakeBool (false ),nil ;};return MakeBool (false ),_e .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");
};func (_ccf *PSOperand )cvi (_geg *PSStack )error {_eaag ,_faea :=_geg .Pop ();if _faea !=nil {return _faea ;};if _geec ,_dd :=_eaag .(*PSReal );_dd {_egc :=int (_geec .Val );_faea =_geg .Push (MakeInteger (_egc ));}else if _eac ,_ffb :=_eaag .(*PSInteger );
_ffb {_dcbc :=_eac .Val ;_faea =_geg .Push (MakeInteger (_dcbc ));}else {return ErrTypeCheck ;};return _faea ;};

// String returns a string representation of the stack.
func (_gdd *PSStack )String ()string {_gbdf :="\u005b\u0020";for _ ,_dbdb :=range *_gdd {_gbdf +=_dbdb .String ();_gbdf +="\u0020";};_gbdf +="\u005d";return _gbdf ;};func (_ebg *PSOperand )atan (_gag *PSStack )error {_acc ,_gb :=_gag .PopNumberAsFloat64 ();
if _gb !=nil {return _gb ;};_dce ,_gb :=_gag .PopNumberAsFloat64 ();if _gb !=nil {return _gb ;};if _acc ==0{var _dbce error ;if _dce < 0{_dbce =_gag .Push (MakeReal (270));}else {_dbce =_gag .Push (MakeReal (90));};return _dbce ;};_gee :=_dce /_acc ;_agf :=_d .Atan (_gee )*180/_d .Pi ;
_gb =_gag .Push (MakeReal (_agf ));return _gb ;};

// Exec executes the operand `op` in the state specified by `stack`.
func (_efa *PSOperand )Exec (stack *PSStack )error {_abbc :=ErrUnsupportedOperand ;switch *_efa {case "\u0061\u0062\u0073":_abbc =_efa .abs (stack );case "\u0061\u0064\u0064":_abbc =_efa .add (stack );case "\u0061\u006e\u0064":_abbc =_efa .and (stack );
case "\u0061\u0074\u0061\u006e":_abbc =_efa .atan (stack );case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":_abbc =_efa .bitshift (stack );case "\u0063e\u0069\u006c\u0069\u006e\u0067":_abbc =_efa .ceiling (stack );case "\u0063\u006f\u0070\u0079":_abbc =_efa .copy (stack );
case "\u0063\u006f\u0073":_abbc =_efa .cos (stack );case "\u0063\u0076\u0069":_abbc =_efa .cvi (stack );case "\u0063\u0076\u0072":_abbc =_efa .cvr (stack );case "\u0064\u0069\u0076":_abbc =_efa .div (stack );case "\u0064\u0075\u0070":_abbc =_efa .dup (stack );
case "\u0065\u0071":_abbc =_efa .eq (stack );case "\u0065\u0078\u0063\u0068":_abbc =_efa .exch (stack );case "\u0065\u0078\u0070":_abbc =_efa .exp (stack );case "\u0066\u006c\u006fo\u0072":_abbc =_efa .floor (stack );case "\u0067\u0065":_abbc =_efa .ge (stack );
case "\u0067\u0074":_abbc =_efa .gt (stack );case "\u0069\u0064\u0069\u0076":_abbc =_efa .idiv (stack );case "\u0069\u0066":_abbc =_efa .ifCondition (stack );case "\u0069\u0066\u0065\u006c\u0073\u0065":_abbc =_efa .ifelse (stack );case "\u0069\u006e\u0064e\u0078":_abbc =_efa .index (stack );
case "\u006c\u0065":_abbc =_efa .le (stack );case "\u006c\u006f\u0067":_abbc =_efa .log (stack );case "\u006c\u006e":_abbc =_efa .ln (stack );case "\u006c\u0074":_abbc =_efa .lt (stack );case "\u006d\u006f\u0064":_abbc =_efa .mod (stack );case "\u006d\u0075\u006c":_abbc =_efa .mul (stack );
case "\u006e\u0065":_abbc =_efa .ne (stack );case "\u006e\u0065\u0067":_abbc =_efa .neg (stack );case "\u006e\u006f\u0074":_abbc =_efa .not (stack );case "\u006f\u0072":_abbc =_efa .or (stack );case "\u0070\u006f\u0070":_abbc =_efa .pop (stack );case "\u0072\u006f\u0075n\u0064":_abbc =_efa .round (stack );
case "\u0072\u006f\u006c\u006c":_abbc =_efa .roll (stack );case "\u0073\u0069\u006e":_abbc =_efa .sin (stack );case "\u0073\u0071\u0072\u0074":_abbc =_efa .sqrt (stack );case "\u0073\u0075\u0062":_abbc =_efa .sub (stack );case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":_abbc =_efa .truncate (stack );
case "\u0078\u006f\u0072":_abbc =_efa .xor (stack );};return _abbc ;};

// PSBoolean represents a boolean value.
type PSBoolean struct{Val bool ;};func (_fecc *PSOperand )dup (_feee *PSStack )error {_bdg ,_addc :=_feee .Pop ();if _addc !=nil {return _addc ;};_addc =_feee .Push (_bdg );if _addc !=nil {return _addc ;};_addc =_feee .Push (_bdg .Duplicate ());return _addc ;
};func (_dg *PSOperand )and (_ce *PSStack )error {_ega ,_fbad :=_ce .Pop ();if _fbad !=nil {return _fbad ;};_ge ,_fbad :=_ce .Pop ();if _fbad !=nil {return _fbad ;};if _ggec ,_cgd :=_ega .(*PSBoolean );_cgd {_gce ,_dca :=_ge .(*PSBoolean );if !_dca {return ErrTypeCheck ;
};_fbad =_ce .Push (MakeBool (_ggec .Val &&_gce .Val ));return _fbad ;};if _aec ,_fge :=_ega .(*PSInteger );_fge {_gdg ,_dcae :=_ge .(*PSInteger );if !_dcae {return ErrTypeCheck ;};_fbad =_ce .Push (MakeInteger (_aec .Val &_gdg .Val ));return _fbad ;};
return ErrTypeCheck ;};func (_ec *PSProgram )DebugString ()string {_dec :="\u007b\u0020";for _ ,_dcb :=range *_ec {_dec +=_dcb .DebugString ();_dec +="\u0020";};_dec +="\u007d";return _dec ;};func (_gac *PSBoolean )DebugString ()string {return _g .Sprintf ("\u0062o\u006f\u006c\u003a\u0025\u0076",_gac .Val );
};

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand (val string )*PSOperand {_eaeb :=PSOperand (val );return &_eaeb };func (_agec *PSOperand )or (_fdb *PSStack )error {_bebe ,_gdgd :=_fdb .Pop ();if _gdgd !=nil {return _gdgd ;};_acg ,_gdgd :=_fdb .Pop ();if _gdgd !=nil {return _gdgd ;};
if _bcc ,_gbd :=_bebe .(*PSBoolean );_gbd {_caeb ,_aaef :=_acg .(*PSBoolean );if !_aaef {return ErrTypeCheck ;};_gdgd =_fdb .Push (MakeBool (_bcc .Val ||_caeb .Val ));return _gdgd ;};if _efag ,_adfd :=_bebe .(*PSInteger );_adfd {_dbf ,_abc :=_acg .(*PSInteger );
if !_abc {return ErrTypeCheck ;};_gdgd =_fdb .Push (MakeInteger (_efag .Val |_dbf .Val ));return _gdgd ;};return ErrTypeCheck ;};

// Append appends an object to the PSProgram.
func (_ba *PSProgram )Append (obj PSObject ){*_ba =append (*_ba ,obj )};func (_aabg *PSOperand )ifelse (_fcg *PSStack )error {_eae ,_bab :=_fcg .Pop ();if _bab !=nil {return _bab ;};_efdb ,_bab :=_fcg .Pop ();if _bab !=nil {return _bab ;};_gae ,_bab :=_fcg .Pop ();
if _bab !=nil {return _bab ;};_gdae ,_bcaa :=_eae .(*PSProgram );if !_bcaa {return ErrTypeCheck ;};_fecdb ,_bcaa :=_efdb .(*PSProgram );if !_bcaa {return ErrTypeCheck ;};_gea ,_bcaa :=_gae .(*PSBoolean );if !_bcaa {return ErrTypeCheck ;};if _gea .Val {_acee :=_fecdb .Exec (_fcg );
return _acee ;};_bab =_gdae .Exec (_fcg );return _bab ;};

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject ;func (_ggg *PSOperand )exch (_cee *PSStack )error {_fecd ,_ebda :=_cee .Pop ();if _ebda !=nil {return _ebda ;};_feef ,_ebda :=_cee .Pop ();if _ebda !=nil {return _ebda ;};_ebda =_cee .Push (_fecd );if _ebda !=nil {return _ebda ;
};_ebda =_cee .Push (_feef );return _ebda ;};

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct{Stack *PSStack ;_ga *PSProgram ;};func (_fgbc *PSOperand )sin (_fdg *PSStack )error {_fda ,_fbeg :=_fdg .PopNumberAsFloat64 ();if _fbeg !=nil {return _fbeg ;};_dedb :=_d .Sin (_fda *_d .Pi /180.0);_fbeg =_fdg .Push (MakeReal (_dedb ));
return _fbeg ;};func (_gdb *PSOperand )sqrt (_afbe *PSStack )error {_ecec ,_dcg :=_afbe .PopNumberAsFloat64 ();if _dcg !=nil {return _dcg ;};if _ecec < 0{return ErrRangeCheck ;};_dfde :=_d .Sqrt (_ecec );_dcg =_afbe .Push (MakeReal (_dfde ));return _dcg ;
};func (_bbb *PSOperand )cvr (_faab *PSStack )error {_egf ,_eccd :=_faab .Pop ();if _eccd !=nil {return _eccd ;};if _caga ,_daa :=_egf .(*PSReal );_daa {_eccd =_faab .Push (MakeReal (_caga .Val ));}else if _afe ,_eef :=_egf .(*PSInteger );_eef {_eccd =_faab .Push (MakeReal (float64 (_afe .Val )));
}else {return ErrTypeCheck ;};return _eccd ;};

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_ag *PSExecutor )Execute (objects []PSObject )([]PSObject ,error ){for _ ,_ae :=range objects {_be :=_ag .Stack .Push (_ae );if _be !=nil {return nil ,_be ;};};_fc :=_ag ._ga .Exec (_ag .Stack );if _fc !=nil {_dc .Log .Debug ("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_fc );
return nil ,_fc ;};_df :=[]PSObject (*_ag .Stack );_ag .Stack .Empty ();return _df ,nil ;};func (_faa *PSReal )String ()string {return _g .Sprintf ("\u0025\u002e\u0035\u0066",_faa .Val )};

// Parse parses the postscript and store as a program that can be executed.
func (_bdee *PSParser )Parse ()(*PSProgram ,error ){_bdee .skipSpaces ();_gfbc ,_ffg :=_bdee ._cbd .Peek (2);if _ffg !=nil {return nil ,_ffg ;};if _gfbc [0]!='{'{return nil ,_e .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b");
};_bbg ,_ffg :=_bdee .parseFunction ();if _ffg !=nil &&_ffg !=_c .EOF {return nil ,_ffg ;};return _bbg ,_ffg ;};func (_eeg *PSOperand )truncate (_ecef *PSStack )error {_ebbf ,_gbc :=_ecef .Pop ();if _gbc !=nil {return _gbc ;};if _bfeg ,_aea :=_ebbf .(*PSReal );
_aea {_ebba :=int (_bfeg .Val );_gbc =_ecef .Push (MakeReal (float64 (_ebba )));}else if _efagb ,_agbe :=_ebbf .(*PSInteger );_agbe {_gbc =_ecef .Push (MakeInteger (_efagb .Val ));}else {return ErrTypeCheck ;};return _gbc ;};func (_dga *PSOperand )copy (_gbe *PSStack )error {_dfa ,_dbeb :=_gbe .PopInteger ();
if _dbeb !=nil {return _dbeb ;};if _dfa < 0{return ErrRangeCheck ;};if _dfa > len (*_gbe ){return ErrRangeCheck ;};*_gbe =append (*_gbe ,(*_gbe )[len (*_gbe )-_dfa :]...);return nil ;};

// PSObject represents a postscript object.
type PSObject interface{

// Duplicate makes a fresh copy of the PSObject.
Duplicate ()PSObject ;

// DebugString returns a descriptive representation of the PSObject with more information than String()
// for debugging purposes.
DebugString ()string ;

// String returns a string representation of the PSObject.
String ()string ;};

// Push pushes an object on top of the stack.
func (_ebbff *PSStack )Push (obj PSObject )error {if len (*_ebbff )> 100{return ErrStackOverflow ;};*_ebbff =append (*_ebbff ,obj );return nil ;};func (_bde *PSOperand )floor (_ebdd *PSStack )error {_gef ,_faad :=_ebdd .Pop ();if _faad !=nil {return _faad ;
};if _fgaf ,_eda :=_gef .(*PSReal );_eda {_faad =_ebdd .Push (MakeReal (_d .Floor (_fgaf .Val )));}else if _cff ,_gggb :=_gef .(*PSInteger );_gggb {_faad =_ebdd .Push (MakeInteger (_cff .Val ));}else {return ErrTypeCheck ;};return _faad ;};

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_cfg *PSStack )PopNumberAsFloat64 ()(float64 ,error ){_cfdb ,_fce :=_cfg .Pop ();if _fce !=nil {return 0,_fce ;};if _aaeg ,_gecf :=_cfdb .(*PSReal );_gecf {return _aaeg .Val ,nil ;}else if _aca ,_ffgc :=_cfdb .(*PSInteger );_ffgc {return float64 (_aca .Val ),nil ;
}else {return 0,ErrTypeCheck ;};};func (_dfg *PSInteger )Duplicate ()PSObject {_ccb :=PSInteger {};_ccb .Val =_dfg .Val ;return &_ccb };func (_fg *PSReal )DebugString ()string {return _g .Sprintf ("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f",_fg .Val );
};