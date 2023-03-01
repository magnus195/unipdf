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

package svg ;import (_e "encoding/xml";_c "fmt";_d "github.com/unidoc/unipdf/v3/common";_aa "github.com/unidoc/unipdf/v3/contentstream";_bc "github.com/unidoc/unipdf/v3/contentstream/draw";_dd "github.com/unidoc/unipdf/v3/internal/graphic2d";_fd "golang.org/x/net/html/charset";
_ad "io";_gf "math";_a "os";_gc "strconv";_g "strings";_f "unicode";);type GraphicSVGStyle struct{FillColor string ;StrokeColor string ;StrokeWidth float64 ;};func _caf (_eeg string )(_bbe ,_cef ,_ebf float64 ){if (len (_eeg )!=4&&len (_eeg )!=7)||_eeg [0]!='#'{_d .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_eeg );
return _bbe ,_cef ,_ebf ;};var _dgcb ,_ffaf ,_dff int ;if len (_eeg )==4{var _edf ,_agb ,_abdga int ;_bcga ,_fbef :=_c .Sscanf (_eeg ,"\u0023\u0025\u0031\u0078\u0025\u0031\u0078\u0025\u0031\u0078",&_edf ,&_agb ,&_abdga );if _fbef !=nil {_d .Log .Debug ("\u0049\u006e\u0076a\u006c\u0069\u0064\u0020h\u0065\u0078\u0020\u0063\u006f\u0064\u0065:\u0020\u0025\u0073\u002c\u0020\u0065\u0072\u0072\u006f\u0072\u003a\u0020\u0025\u0076",_eeg ,_fbef );
return _bbe ,_cef ,_ebf ;};if _bcga !=3{_d .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_eeg );return _bbe ,_cef ,_ebf ;};_dgcb =_edf *16+_edf ;_ffaf =_agb *16+_agb ;_dff =_abdga *16+_abdga ;
}else {_efd ,_dce :=_c .Sscanf (_eeg ,"\u0023\u0025\u0032\u0078\u0025\u0032\u0078\u0025\u0032\u0078",&_dgcb ,&_ffaf ,&_dff );if _dce !=nil {_d .Log .Debug ("I\u006ev\u0061\u006c\u0069\u0064\u0020\u0068\u0065\u0078 \u0063\u006f\u0064\u0065: \u0025\u0073",_eeg );
return _bbe ,_cef ,_ebf ;};if _efd !=3{_d .Log .Debug ("\u0049\u006e\u0076\u0061\u006c\u0069d\u0020\u0068\u0065\u0078\u0020\u0063\u006f\u0064\u0065\u003a\u0020\u0025\u0073,\u0020\u006e\u0020\u0021\u003d\u0020\u0033 \u0028\u0025\u0064\u0029",_eeg ,_efd );
return _bbe ,_cef ,_ebf ;};};_fea :=float64 (_dgcb )/255.0;_fegf :=float64 (_ffaf )/255.0;_adaf :=float64 (_dff )/255.0;return _fea ,_fegf ,_adaf ;};func (_dbc *GraphicSVGStyle )toContentStream (_afbb *_aa .ContentCreator ){if _dbc ==nil {return ;};if _dbc .FillColor !=""{var _fgcb ,_ggdg ,_ebd float64 ;
if _dgfge ,_dbgc :=_dd .ColorMap [_dbc .FillColor ];_dbgc {_fcg ,_bff ,_edg ,_ :=_dgfge .RGBA ();_fgcb ,_ggdg ,_ebd =float64 (_fcg ),float64 (_bff ),float64 (_edg );}else {_fgcb ,_ggdg ,_ebd =_caf (_dbc .FillColor );};_afbb .Add_rg (_fgcb ,_ggdg ,_ebd );
};if _dbc .StrokeColor !=""{var _gef ,_gde ,_fab float64 ;if _dbcd ,_gdca :=_dd .ColorMap [_dbc .StrokeColor ];_gdca {_bca ,_bcce ,_eef ,_ :=_dbcd .RGBA ();_gef ,_gde ,_fab =float64 (_bca )/255.0,float64 (_bcce )/255.0,float64 (_eef )/255.0;}else {_gef ,_gde ,_fab =_caf (_dbc .StrokeColor );
};_afbb .Add_RG (_gef ,_gde ,_fab );};if _dbc .StrokeWidth > 0{_afbb .Add_w (_dbc .StrokeWidth );};};type pathParserError struct{_fdgeg string };type token struct{_fe string ;_cabd bool ;};func (_baaf *GraphicSVG )setDefaultScaling (_ddd float64 ){_baaf ._ceb =_ddd ;
if _baaf .Style !=nil &&_baaf .Style .StrokeWidth > 0{_baaf .Style .StrokeWidth =_baaf .Style .StrokeWidth *_baaf ._ceb ;};for _ ,_gcda :=range _baaf .Children {_gcda .setDefaultScaling (_ddd );};};func _dbbe ()*GraphicSVGStyle {return &GraphicSVGStyle {FillColor :"\u00230\u0030\u0030\u0030\u0030\u0030",StrokeColor :"",StrokeWidth :0};
};func _fgcd (_afe []*Command )*Path {_ecbe :=&Path {};var _bee []*Command ;for _cadd ,_bfgf :=range _afe {switch _g .ToLower (_bfgf .Symbol ){case _gbf ._aae :if len (_bee )> 0{_ecbe .Subpaths =append (_ecbe .Subpaths ,&Subpath {_bee });};_bee =[]*Command {_bfgf };
case _gbf ._dbge :_bee =append (_bee ,_bfgf );_ecbe .Subpaths =append (_ecbe .Subpaths ,&Subpath {_bee });_bee =[]*Command {};default:_bee =append (_bee ,_bfgf );if len (_afe )==_cadd +1{_ecbe .Subpaths =append (_ecbe .Subpaths ,&Subpath {_bee });};};};
return _ecbe ;};func (_ace *GraphicSVG )ToContentCreator (cc *_aa .ContentCreator )*_aa .ContentCreator {if _ace .Name =="\u0073\u0076\u0067"{cc .Add_cm (1,0,0,1,0,0);_ace .setDefaultScaling (_ace ._ceb );for _ ,_cab :=range _ace .Children {_cab .ViewBox =_ace .ViewBox ;
_cab .toContentStream (cc );};return cc ;};return nil ;};func _ccc (_cdf string )(_acdd ,_acg string ){if _cdf ==""||(_cdf [len (_cdf )-1]>='0'&&_cdf [len (_cdf )-1]<='9'){return _cdf ,"";};_acdd =_cdf ;for _ ,_dcaa :=range _ca {if _g .Contains (_acdd ,_dcaa ){_acg =_dcaa ;
};_acdd =_g .TrimSuffix (_acdd ,_dcaa );};return ;};func _fdgb (_ecc []token )([]*Command ,error ){var (_afbd []*Command ;_feg []float64 ;);for _bef :=len (_ecc )-1;_bef >=0;_bef --{_gbb :=_ecc [_bef ];if _gbb ._cabd {_dbgf :=_gbf ._acd [_g .ToLower (_gbb ._fe )];
_gbe :=len (_feg );if _dbgf ==0&&_gbe ==0{_aec :=&Command {Symbol :_gbb ._fe };_afbd =append ([]*Command {_aec },_afbd ...);}else if _dbgf !=0&&_gbe %_dbgf ==0{_ffe :=_gbe /_dbgf ;for _fga :=0;_fga < _ffe ;_fga ++{_fggf :=_gbb ._fe ;if _fggf =="\u006d"&&_fga < _ffe -1{_fggf ="\u006c";
};if _fggf =="\u004d"&&_fga < _ffe -1{_fggf ="\u004c";};_ega :=&Command {_fggf ,_egdb (_feg [:_dbgf ])};_afbd =append ([]*Command {_ega },_afbd ...);_feg =_feg [_dbgf :];};}else {_bfa :=pathParserError {"I\u006e\u0063\u006f\u0072\u0072\u0065c\u0074\u0020\u006e\u0075\u006d\u0062e\u0072\u0020\u006f\u0066\u0020\u0070\u0061r\u0061\u006d\u0065\u0074\u0065\u0072\u0073\u0020\u0066\u006fr\u0020"+_gbb ._fe };
return nil ,_bfa ;};}else {_cbd ,_aga :=_abdf (_gbb ._fe ,64);if _aga !=nil {return nil ,_aga ;};_feg =append (_feg ,_cbd );};};return _afbd ,nil ;};func _ce (_fa *GraphicSVG ,_ebb *_aa .ContentCreator ){_ebb .Add_q ();_fa .Style .toContentStream (_ebb );
_dgfg ,_gdg :=_abdf (_fa .Attributes ["\u0063\u0078"],64);if _gdg !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gdg .Error ());
};_gfcc ,_gdg :=_abdf (_fa .Attributes ["\u0063\u0079"],64);if _gdg !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_gdg .Error ());
};_dab ,_gdg :=_abdf (_fa .Attributes ["\u0072"],64);if _gdg !=nil {_d .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0072\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_gdg .Error ());
};_bbd :=_dab *_fa ._ceb ;_gfa :=_dab *_fa ._ceb ;_bfg :=_bbd *_be ;_ecb :=_gfa *_be ;_abf :=_bc .NewCubicBezierPath ();_abf =_abf .AppendCurve (_bc .NewCubicBezierCurve (-_bbd ,0,-_bbd ,_ecb ,-_bfg ,_gfa ,0,_gfa ));_abf =_abf .AppendCurve (_bc .NewCubicBezierCurve (0,_gfa ,_bfg ,_gfa ,_bbd ,_ecb ,_bbd ,0));
_abf =_abf .AppendCurve (_bc .NewCubicBezierCurve (_bbd ,0,_bbd ,-_ecb ,_bfg ,-_gfa ,0,-_gfa ));_abf =_abf .AppendCurve (_bc .NewCubicBezierCurve (0,-_gfa ,-_bfg ,-_gfa ,-_bbd ,-_ecb ,-_bbd ,0));_abf =_abf .Offset (_dgfg *_fa ._ceb ,_gfcc *_fa ._ceb );
if _fa .Style .StrokeWidth > 0{_abf =_abf .Offset (_fa .Style .StrokeWidth /2,_fa .Style .StrokeWidth /2);};_bc .DrawBezierPathWithCreator (_abf ,_ebb );if _fa .Style .FillColor !=""&&_fa .Style .StrokeColor !=""{_ebb .Add_B ();}else if _fa .Style .FillColor !=""{_ebb .Add_f ();
}else if _fa .Style .StrokeColor !=""{_ebb .Add_S ();};_ebb .Add_h ();_ebb .Add_Q ();};func (_dac *Command )isAbsolute ()bool {return _dac .Symbol ==_g .ToUpper (_dac .Symbol )};func _dgf (_dbb *GraphicSVG ,_adc *_aa .ContentCreator ){_adc .Add_q ();_dbb .Style .toContentStream (_adc );
_dag ,_bg :=_abdf (_dbb .Attributes ["\u0078"],64);if _bg !=nil {_d .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0078\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_bg .Error ());
};_de ,_bg :=_abdf (_dbb .Attributes ["\u0079"],64);if _bg !=nil {_d .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020w\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020`\u0079\u0060\u0020\u0076\u0061\u006c\u0075e\u003a\u0020\u0025\u0076",_bg .Error ());
};_ec ,_bg :=_abdf (_dbb .Attributes ["\u0077\u0069\u0064t\u0068"],64);if _bg !=nil {_d .Log .Debug ("\u0045\u0072\u0072o\u0072\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061\u0072\u0073\u0069\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0077\u0069\u0064\u0074\u0068\u0020v\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_bg .Error ());
};_bcc ,_bg :=_abdf (_dbb .Attributes ["\u0068\u0065\u0069\u0067\u0068\u0074"],64);if _bg !=nil {_d .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0077h\u0069\u006c\u0065 \u0070\u0061\u0072\u0073i\u006e\u0067\u0020\u0073\u0074\u0072\u006f\u006b\u0065\u0020\u0068\u0065\u0069\u0067\u0068\u0074\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_bg .Error ());
};_adc .Add_re (_dag *_dbb ._ceb ,_de *_dbb ._ceb ,_ec *_dbb ._ceb ,_bcc *_dbb ._ceb );if _dbb .Style .FillColor !=""&&_dbb .Style .StrokeColor !=""{_adc .Add_B ();}else if _dbb .Style .FillColor !=""{_adc .Add_f ();}else if _dbb .Style .StrokeColor !=""{_adc .Add_S ();
};_adc .Add_Q ();};func _edb ()commands {var _fggg =map[string ]int {"\u006d":2,"\u007a":0,"\u006c":2,"\u0068":1,"\u0076":1,"\u0063":6,"\u0073":4,"\u0071":4,"\u0074":2,"\u0061":7};var _cadf []string ;for _gba :=range _fggg {_cadf =append (_cadf ,_gba );
};return commands {_cadf ,_fggg ,"\u006d","\u007a"};};func (_aee *GraphicSVG )SetScaling (xFactor ,yFactor float64 ){_baf :=_aee .Width /_aee .ViewBox .W ;_bbbfd :=_aee .Height /_aee .ViewBox .H ;_aee .setDefaultScaling (_gf .Max (_baf ,_bbbfd ));for _ ,_cad :=range _aee .Children {_cad .SetScaling (xFactor ,yFactor );
};};func (_ecf *GraphicSVG )Decode (decoder *_e .Decoder )error {for {_cgcg ,_ced :=decoder .Token ();if _cgcg ==nil &&_ced ==_ad .EOF {break ;};if _ced !=nil {return _ced ;};switch _fdcb :=_cgcg .(type ){case _e .StartElement :_gcaa :=_efa (_fdcb );_eba :=_gcaa .Decode (decoder );
if _eba !=nil {return _eba ;};_ecf .Children =append (_ecf .Children ,_gcaa );case _e .CharData :_gga :=_g .TrimSpace (string (_fdcb ));if _gga !=""{_ecf .Content =string (_fdcb );};case _e .EndElement :if _fdcb .Name .Local ==_ecf .Name {return nil ;};
};};return nil ;};func _abdf (_ebe string ,_cfaf int )(float64 ,error ){_dbfd ,_gefa :=_ccc (_ebe );_bbfd ,_cdb :=_gc .ParseFloat (_dbfd ,_cfaf );if _cdb !=nil {return 0,_cdb ;};if _dfe ,_aac :=_gfc [_gefa ];_aac {_bbfd =_bbfd *_dfe ;}else {_bbfd =_bbfd *_bce ;
};return _bbfd ,nil ;};func ParseFromFile (path string )(*GraphicSVG ,error ){_gebb ,_fbeg :=_a .Open (path );if _fbeg !=nil {return nil ,_fbeg ;};defer _gebb .Close ();return ParseFromStream (_gebb );};func (_acb *GraphicSVG )toContentStream (_cgf *_aa .ContentCreator ){_bcf ,_bbbf :=_ceg (_acb .Attributes ,_acb ._ceb );
if _bbbf !=nil {_d .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0073\u0074\u0079\u006c\u0065\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_bbbf );};_acb .Style =_bcf ;switch _acb .Name {case "\u0070\u0061\u0074\u0068":_gg (_acb ,_cgf );
for _ ,_cdd :=range _acb .Children {_cdd .toContentStream (_cgf );};case "\u0072\u0065\u0063\u0074":_dgf (_acb ,_cgf );for _ ,_cgb :=range _acb .Children {_cgb .toContentStream (_cgf );};case "\u0063\u0069\u0072\u0063\u006c\u0065":_ce (_acb ,_cgf );for _ ,_gcf :=range _acb .Children {_gcf .toContentStream (_cgf );
};case "\u0065l\u006c\u0069\u0070\u0073\u0065":_babg (_acb ,_cgf );for _ ,_gcd :=range _acb .Children {_gcd .toContentStream (_cgf );};case "\u0070\u006f\u006c\u0079\u006c\u0069\u006e\u0065":_cgc (_acb ,_cgf );for _ ,_acc :=range _acb .Children {_acc .toContentStream (_cgf );
};case "\u0070o\u006c\u0079\u0067\u006f\u006e":_cgd (_acb ,_cgf );for _ ,_efbg :=range _acb .Children {_efbg .toContentStream (_cgf );};case "\u006c\u0069\u006e\u0065":_gag (_acb ,_cgf );for _ ,_faf :=range _acb .Children {_faf .toContentStream (_cgf );
};case "\u0067":_bfb ,_add :=_acb .Attributes ["\u0066\u0069\u006c\u006c"];_cbc ,_bge :=_acb .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];_bgc ,_ddb :=_acb .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];for _ ,_cea :=range _acb .Children {if _ ,_dagf :=_cea .Attributes ["\u0066\u0069\u006c\u006c"];
!_dagf &&_add {_cea .Attributes ["\u0066\u0069\u006c\u006c"]=_bfb ;};if _ ,_gbdc :=_cea .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"];!_gbdc &&_bge {_cea .Attributes ["\u0073\u0074\u0072\u006f\u006b\u0065"]=_cbc ;};if _ ,_gdaa :=_cea .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];
!_gdaa &&_ddb {_cea .Attributes ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"]=_bgc ;};_cea .toContentStream (_cgf );};};};func _egdb (_fbfe []float64 )[]float64 {for _aea ,_dfgb :=0,len (_fbfe )-1;_aea < _dfgb ;_aea ,_dfgb =_aea +1,_dfgb -1{_fbfe [_aea ],_fbfe [_dfgb ]=_fbfe [_dfgb ],_fbfe [_aea ];
};return _fbfe ;};func _cgd (_aef *GraphicSVG ,_gfac *_aa .ContentCreator ){_gfac .Add_q ();_aef .Style .toContentStream (_gfac );_ggdf ,_eab :=_ecd (_aef .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);if _eab !=nil {_d .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_eab );
return ;};if len (_ggdf )%2> 0{_d .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _edc :=0;_edc < len (_ggdf );{if _edc ==0{_gfac .Add_m (_ggdf [_edc ]*_aef ._ceb ,_ggdf [_edc +1]*_aef ._ceb );}else {_gfac .Add_l (_ggdf [_edc ]*_aef ._ceb ,_ggdf [_edc +1]*_aef ._ceb );};_edc +=2;};_gfac .Add_l (_ggdf [0]*_aef ._ceb ,_ggdf [1]*_aef ._ceb );
if _aef .Style .FillColor !=""&&_aef .Style .StrokeColor !=""{_gfac .Add_B ();}else if _aef .Style .FillColor !=""{_gfac .Add_f ();}else if _aef .Style .StrokeColor !=""{_gfac .Add_S ();};_gfac .Add_h ();_gfac .Add_Q ();};func (_fcgd pathParserError )Error ()string {return _fcgd ._fdgeg };
func _ecd (_gbfg string )([]float64 ,error ){_eeag :=-1;var _bdc []float64 ;_badaa :=' ';for _eade ,_eda :=range _gbfg {if !_f .IsNumber (_eda )&&_eda !='.'&&!(_eda =='-'&&_badaa =='e')&&_eda !='e'{if _eeag !=-1{_aaf ,_fcge :=_bgcd (_gbfg [_eeag :_eade ]);
if _fcge !=nil {return _bdc ,_fcge ;};_bdc =append (_bdc ,_aaf ...);};if _eda =='-'{_eeag =_eade ;}else {_eeag =-1;};}else if _eeag ==-1{_eeag =_eade ;};_badaa =_eda ;};if _eeag !=-1&&_eeag !=len (_gbfg ){_gfag ,_eff :=_bgcd (_gbfg [_eeag :]);if _eff !=nil {return _bdc ,_eff ;
};_bdc =append (_bdc ,_gfag ...);};return _bdc ,nil ;};var _gbf commands ;type Subpath struct{Commands []*Command ;};func (_dca *Command )compare (_gcde *Command )bool {if _dca .Symbol !=_gcde .Symbol {return false ;};for _gcad ,_deg :=range _dca .Params {if _deg !=_gcde .Params [_gcad ]{return false ;
};};return true ;};func (_cdg *Subpath )compare (_bde *Subpath )bool {if len (_cdg .Commands )!=len (_bde .Commands ){return false ;};for _dgfe ,_bfcb :=range _cdg .Commands {if !_bfcb .compare (_bde .Commands [_dgfe ]){return false ;};};return true ;};
type GraphicSVG struct{ViewBox struct{X ,Y ,W ,H float64 ;};Name string ;Attributes map[string ]string ;Children []*GraphicSVG ;Content string ;Style *GraphicSVGStyle ;Width float64 ;Height float64 ;_ceb float64 ;};func _bgcd (_ebdg string )(_fge []float64 ,_beed error ){var _bgf float64 ;
_caa :=0;_aafe :=true ;for _daa ,_cfa :=range _ebdg {if _cfa =='.'{if _aafe {_aafe =false ;continue ;};_bgf ,_beed =_abdf (_ebdg [_caa :_daa ],64);if _beed !=nil {return ;};_fge =append (_fge ,_bgf );_caa =_daa ;};};_bgf ,_beed =_abdf (_ebdg [_caa :],64);
if _beed !=nil {return ;};_fge =append (_fge ,_bgf );return ;};func ParseFromStream (source _ad .Reader )(*GraphicSVG ,error ){_egd :=_e .NewDecoder (source );_egd .CharsetReader =_fd .NewReaderLabel ;_ebac ,_gdf :=_dfgd (_egd );if _gdf !=nil {return nil ,_gdf ;
};if _ddf :=_ebac .Decode (_egd );_ddf !=nil &&_ddf !=_ad .EOF {return nil ,_ddf ;};return _ebac ,nil ;};func _cgc (_cbe *GraphicSVG ,_bcd *_aa .ContentCreator ){_bcd .Add_q ();_cbe .Style .toContentStream (_bcd );_fgb ,_fda :=_ecd (_cbe .Attributes ["\u0070\u006f\u0069\u006e\u0074\u0073"]);
if _fda !=nil {_d .Log .Debug ("\u0045\u0052\u0052O\u0052\u0020\u0075\u006e\u0061\u0062\u006c\u0065\u0020\u0074\u006f\u0020\u0070\u0061\u0072\u0073\u0065\u0020\u0070\u006f\u0069\u006e\u0074\u0073\u0020\u0061\u0074\u0074\u0072i\u0062\u0075\u0074\u0065\u003a\u0020\u0025\u0076",_fda );
return ;};if len (_fgb )%2> 0{_d .Log .Debug ("\u0045\u0052R\u004f\u0052\u0020\u0069n\u0076\u0061l\u0069\u0064\u0020\u0070\u006f\u0069\u006e\u0074s\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u0020\u006ce\u006e\u0067\u0074\u0068");return ;
};for _dbf :=0;_dbf < len (_fgb );{if _dbf ==0{_bcd .Add_m (_fgb [_dbf ]*_cbe ._ceb ,_fgb [_dbf +1]*_cbe ._ceb );}else {_bcd .Add_l (_fgb [_dbf ]*_cbe ._ceb ,_fgb [_dbf +1]*_cbe ._ceb );};_dbf +=2;};if _cbe .Style .FillColor !=""&&_cbe .Style .StrokeColor !=""{_bcd .Add_B ();
}else if _cbe .Style .FillColor !=""{_bcd .Add_f ();}else if _cbe .Style .StrokeColor !=""{_bcd .Add_S ();};_bcd .Add_h ();_bcd .Add_Q ();};type commands struct{_edgc []string ;_acd map[string ]int ;_aae string ;_dbge string ;};func _ceg (_cbf map[string ]string ,_ged float64 )(*GraphicSVGStyle ,error ){_eea :=_dbbe ();
_ffbd ,_eed :=_cbf ["\u0066\u0069\u006c\u006c"];if _eed {_eea .FillColor =_ffbd ;if _ffbd =="\u006e\u006f\u006e\u0065"{_eea .FillColor ="";};};_cag ,_afa :=_cbf ["\u0073\u0074\u0072\u006f\u006b\u0065"];if _afa {_eea .StrokeColor =_cag ;if _cag =="\u006e\u006f\u006e\u0065"{_eea .StrokeColor ="";
};};_gedc ,_ead :=_cbf ["\u0073\u0074\u0072o\u006b\u0065\u002d\u0077\u0069\u0064\u0074\u0068"];if _ead {_age ,_ebdf :=_abdf (_gedc ,64);if _ebdf !=nil {return nil ,_ebdf ;};_eea .StrokeWidth =_age *_ged ;};return _eea ,nil ;};func _fdf (_cdde []token ,_acdb string )([]token ,string ){if _acdb !=""{_cdde =append (_cdde ,token {_acdb ,false });
_acdb ="";};return _cdde ,_acdb ;};func (_egc *Path )compare (_adb *Path )bool {if len (_egc .Subpaths )!=len (_adb .Subpaths ){return false ;};for _abb ,_fbec :=range _egc .Subpaths {if !_fbec .compare (_adb .Subpaths [_abb ]){return false ;};};return true ;
};func ParseFromString (svgStr string )(*GraphicSVG ,error ){return ParseFromStream (_g .NewReader (svgStr ));};func _efa (_fggd _e .StartElement )*GraphicSVG {_gae :=&GraphicSVG {};_bbf :=make (map[string ]string );for _ ,_geb :=range _fggd .Attr {_bbf [_geb .Name .Local ]=_geb .Value ;
};_gae .Name =_fggd .Name .Local ;_gae .Attributes =_bbf ;_gae ._ceb =1;if _gae .Name =="\u0073\u0076\u0067"{_daf ,_bfc :=_ecd (_bbf ["\u0076i\u0065\u0077\u0042\u006f\u0078"]);if _bfc !=nil {_d .Log .Debug ("\u0055\u006ea\u0062\u006c\u0065\u0020t\u006f\u0020p\u0061\u0072\u0073\u0065\u0020\u0076\u0069\u0065w\u0042\u006f\u0078\u0020\u0061\u0074\u0074\u0072\u0069\u0062\u0075\u0074e\u003a\u0020\u0025\u0076",_bfc );
return nil ;};_gae .ViewBox .X =_daf [0];_gae .ViewBox .Y =_daf [1];_gae .ViewBox .W =_daf [2];_gae .ViewBox .H =_daf [3];_gae .Width =_gae .ViewBox .W ;_gae .Height =_gae .ViewBox .H ;if _fbe ,_dgc :=_bbf ["\u0077\u0069\u0064t\u0068"];_dgc {_eca ,_dfg :=_abdf (_fbe ,64);
if _dfg !=nil {_d .Log .Debug ("U\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0070\u0061\u0072\u0073e\u0020\u0077\u0069\u0064\u0074\u0068\u0020a\u0074\u0074\u0072\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_dfg );return nil ;};_gae .Width =_eca ;
};if _fdge ,_bcdb :=_bbf ["\u0068\u0065\u0069\u0067\u0068\u0074"];_bcdb {_gcfc ,_abdg :=_abdf (_fdge ,64);if _abdg !=nil {_d .Log .Debug ("\u0055\u006eab\u006c\u0065\u0020t\u006f\u0020\u0070\u0061rse\u0020he\u0069\u0067\u0068\u0074\u0020\u0061\u0074tr\u0069\u0062\u0075\u0074\u0065\u003a\u0020%\u0076",_abdg );
return nil ;};_gae .Height =_gcfc ;};if _gae .Width > 0&&_gae .Height > 0{_gae ._ceb =_gae .Width /_gae .ViewBox .W ;};};return _gae ;};func (_gbab *commands )isCommand (_eec string )bool {for _ ,_baba :=range _gbab ._edgc {if _g .ToLower (_eec )==_baba {return true ;
};};return false ;};const (_bce =0.72;_eb =28.3464;_ab =_eb /10;_be =0.551784;);func _gag (_ee *GraphicSVG ,_gec *_aa .ContentCreator ){_gec .Add_q ();_ee .Style .toContentStream (_gec );_ecbf ,_baa :=_abdf (_ee .Attributes ["\u0078\u0031"],64);if _baa !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_baa .Error ());
};_bcda ,_baa :=_abdf (_ee .Attributes ["\u0079\u0031"],64);if _baa !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_baa .Error ());
};_cd ,_baa :=_abdf (_ee .Attributes ["\u0078\u0032"],64);if _baa !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_baa .Error ());
};_bbbb ,_baa :=_abdf (_ee .Attributes ["\u0079\u0032"],64);if _baa !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_baa .Error ());
};_gec .Add_m (_ecbf *_ee ._ceb ,_bcda *_ee ._ceb );_gec .Add_l (_cd *_ee ._ceb ,_bbbb *_ee ._ceb );if _ee .Style .FillColor !=""&&_ee .Style .StrokeColor !=""{_gec .Add_B ();}else if _ee .Style .FillColor !=""{_gec .Add_f ();}else if _ee .Style .StrokeColor !=""{_gec .Add_S ();
};_gec .Add_h ();_gec .Add_Q ();};func _dfb (_deb string )(*Path ,error ){_gbf =_edb ();_aab ,_agg :=_fdgb (_cfe (_deb ));if _agg !=nil {return nil ,_agg ;};return _fgcd (_aab ),nil ;};type Path struct{Subpaths []*Subpath ;};func _bafg (_eag float64 ,_bga int )float64 {_cdff :=_gf .Pow (10,float64 (_bga ));
return float64 (_ggab (_eag *_cdff ))/_cdff ;};func _dfgd (_bafc *_e .Decoder )(*GraphicSVG ,error ){for {_fce ,_gefb :=_bafc .Token ();if _fce ==nil &&_gefb ==_ad .EOF {break ;};if _gefb !=nil {return nil ,_gefb ;};switch _abfg :=_fce .(type ){case _e .StartElement :return _efa (_abfg ),nil ;
};};return &GraphicSVG {},nil ;};func _ggab (_ccf float64 )int {return int (_ccf +_gf .Copysign (0.5,_ccf ))};func _cfe (_gfd string )[]token {var (_fddd []token ;_bbbc string ;);for _ ,_dafb :=range _gfd {_fbb :=string (_dafb );switch {case _gbf .isCommand (_fbb ):_fddd ,_bbbc =_fdf (_fddd ,_bbbc );
_fddd =append (_fddd ,token {_fbb ,true });case _fbb =="\u002e":if _bbbc ==""{_bbbc ="\u0030";};if _g .Contains (_bbbc ,_fbb ){_fddd =append (_fddd ,token {_bbbc ,false });_bbbc ="\u0030";};fallthrough;case _fbb >="\u0030"&&_fbb <="\u0039"||_fbb =="\u0065":_bbbc +=_fbb ;
case _fbb =="\u002d":if _g .HasSuffix (_bbbc ,"\u0065"){_bbbc +=_fbb ;}else {_fddd ,_ =_fdf (_fddd ,_bbbc );_bbbc =_fbb ;};default:_fddd ,_bbbc =_fdf (_fddd ,_bbbc );};};_fddd ,_ =_fdf (_fddd ,_bbbc );return _fddd ;};type Command struct{Symbol string ;
Params []float64 ;};func _babg (_abc *GraphicSVG ,_dfc *_aa .ContentCreator ){_dfc .Add_q ();_abc .Style .toContentStream (_dfc );_efb ,_ccb :=_abdf (_abc .Attributes ["\u0063\u0078"],64);if _ccb !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ccb .Error ());
};_fbd ,_ccb :=_abdf (_abc .Attributes ["\u0063\u0079"],64);if _ccb !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0063\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ccb .Error ());
};_dbg ,_ccb :=_abdf (_abc .Attributes ["\u0072\u0078"],64);if _ccb !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0078\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ccb .Error ());
};_ffb ,_ccb :=_abdf (_abc .Attributes ["\u0072\u0079"],64);if _ccb !=nil {_d .Log .Debug ("\u0045\u0072\u0072or\u0020\u0077\u0068\u0069\u006c\u0065\u0020\u0070\u0061r\u0073i\u006eg\u0020`\u0072\u0079\u0060\u0020\u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0025\u0076",_ccb .Error ());
};_bccg :=_dbg *_abc ._ceb ;_gac :=_ffb *_abc ._ceb ;_fdd :=_efb *_abc ._ceb ;_fgc :=_fbd *_abc ._ceb ;_ddc :=_bccg *_be ;_cb :=_gac *_be ;_bdg :=_bc .NewCubicBezierPath ();_bdg =_bdg .AppendCurve (_bc .NewCubicBezierCurve (-_bccg ,0,-_bccg ,_cb ,-_ddc ,_gac ,0,_gac ));
_bdg =_bdg .AppendCurve (_bc .NewCubicBezierCurve (0,_gac ,_ddc ,_gac ,_bccg ,_cb ,_bccg ,0));_bdg =_bdg .AppendCurve (_bc .NewCubicBezierCurve (_bccg ,0,_bccg ,-_cb ,_ddc ,-_gac ,0,-_gac ));_bdg =_bdg .AppendCurve (_bc .NewCubicBezierCurve (0,-_gac ,-_ddc ,-_gac ,-_bccg ,-_cb ,-_bccg ,0));
_bdg =_bdg .Offset (_fdd ,_fgc );if _abc .Style .StrokeWidth > 0{_bdg =_bdg .Offset (_abc .Style .StrokeWidth /2,_abc .Style .StrokeWidth /2);};_bc .DrawBezierPathWithCreator (_bdg ,_dfc );if _abc .Style .FillColor !=""&&_abc .Style .StrokeColor !=""{_dfc .Add_B ();
}else if _abc .Style .FillColor !=""{_dfc .Add_f ();}else if _abc .Style .StrokeColor !=""{_dfc .Add_S ();};_dfc .Add_h ();_dfc .Add_Q ();};var (_ca =[]string {"\u0063\u006d","\u006d\u006d","\u0070\u0078","\u0070\u0074"};_gfc =map[string ]float64 {"\u0063\u006d":_eb ,"\u006d\u006d":_ab ,"\u0070\u0078":_bce ,"\u0070\u0074":1};
);func _gg (_ed *GraphicSVG ,_ff *_aa .ContentCreator ){_ff .Add_q ();_ed .Style .toContentStream (_ff );_fb ,_cf :=_dfb (_ed .Attributes ["\u0064"]);if _cf !=nil {_d .Log .Error ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0025s",_cf .Error ());};var (_cff ,_abd =0.0,0.0;
_cg ,_ba =0.0,0.0;_gb *Command ;);for _ ,_fdc :=range _fb .Subpaths {for _ ,_bad :=range _fdc .Commands {switch _g .ToLower (_bad .Symbol ){case "\u006d":_cg ,_ba =_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ;if !_bad .isAbsolute (){_cg ,_ba =_cff +_cg -_ed .ViewBox .X ,_abd +_ba -_ed .ViewBox .Y ;
};_ff .Add_m (_bafg (_cg ,3),_bafg (_ba ,3));_cff ,_abd =_cg ,_ba ;case "\u0063":_gd ,_ge ,_ggd ,_gbd ,_fg ,_eg :=_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ,_bad .Params [2]*_ed ._ceb ,_bad .Params [3]*_ed ._ceb ,_bad .Params [4]*_ed ._ceb ,_bad .Params [5]*_ed ._ceb ;
if !_bad .isAbsolute (){_gd ,_ge ,_ggd ,_gbd ,_fg ,_eg =_cff +_gd ,_abd +_ge ,_cff +_ggd ,_abd +_gbd ,_cff +_fg ,_abd +_eg ;};_ff .Add_c (_bafg (_gd ,3),_bafg (_ge ,3),_bafg (_ggd ,3),_bafg (_gbd ,3),_bafg (_fg ,3),_bafg (_eg ,3));_cff ,_abd =_fg ,_eg ;
case "\u0073":_ac ,_ea ,_dg ,_bf :=_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ,_bad .Params [2]*_ed ._ceb ,_bad .Params [3]*_ed ._ceb ;if !_bad .isAbsolute (){_ac ,_ea ,_dg ,_bf =_cff +_ac ,_abd +_ea ,_cff +_dg ,_abd +_bf ;};_ff .Add_c (_bafg (_cff ,3),_bafg (_abd ,3),_bafg (_ac ,3),_bafg (_ea ,3),_bafg (_dg ,3),_bafg (_bf ,3));
_cff ,_abd =_dg ,_bf ;case "\u006c":_bab ,_da :=_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ;if !_bad .isAbsolute (){_bab ,_da =_cff +_bab ,_abd +_da ;};_ff .Add_l (_bafg (_bab ,3),_bafg (_da ,3));_cff ,_abd =_bab ,_da ;case "\u0068":_cc :=_bad .Params [0]*_ed ._ceb ;
if !_bad .isAbsolute (){_cc =_cff +_cc ;};_ff .Add_l (_bafg (_cc ,3),_bafg (_abd ,3));_cff =_cc ;case "\u0076":_fc :=_bad .Params [0]*_ed ._ceb ;if !_bad .isAbsolute (){_fc =_abd +_fc ;};_ff .Add_l (_bafg (_cff ,3),_bafg (_fc ,3));_abd =_fc ;case "\u0071":_bd ,_ffa ,_gca ,_bb :=_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ,_bad .Params [2]*_ed ._ceb ,_bad .Params [3]*_ed ._ceb ;
if !_bad .isAbsolute (){_bd ,_ffa ,_gca ,_bb =_cff +_bd ,_abd +_ffa ,_cff +_gca ,_abd +_bb ;};_bbb ,_dde :=_dd .QuadraticToCubicBezier (_cff ,_abd ,_bd ,_ffa ,_gca ,_bb );_ff .Add_c (_bafg (_bbb .X ,3),_bafg (_bbb .Y ,3),_bafg (_dde .X ,3),_bafg (_dde .Y ,3),_bafg (_gca ,3),_bafg (_bb ,3));
_cff ,_abd =_gca ,_bb ;case "\u0074":var _ef ,_ga _dd .Point ;_ae ,_gfg :=_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ;if !_bad .isAbsolute (){_ae ,_gfg =_cff +_ae ,_abd +_gfg ;};if _gb !=nil &&_g .ToLower (_gb .Symbol )=="\u0071"{_df :=_dd .Point {X :_gb .Params [0]*_ed ._ceb ,Y :_gb .Params [1]*_ed ._ceb };
_gdc :=_dd .Point {X :_gb .Params [2]*_ed ._ceb ,Y :_gb .Params [3]*_ed ._ceb };_gcag :=_gdc .Mul (2.0).Sub (_df );_ef ,_ga =_dd .QuadraticToCubicBezier (_cff ,_abd ,_gcag .X ,_gcag .Y ,_ae ,_gfg );};_ff .Add_c (_bafg (_ef .X ,3),_bafg (_ef .Y ,3),_bafg (_ga .X ,3),_bafg (_ga .Y ,3),_bafg (_ae ,3),_bafg (_gfg ,3));
_cff ,_abd =_ae ,_gfg ;case "\u0061":_gcb ,_cae :=_bad .Params [0]*_ed ._ceb ,_bad .Params [1]*_ed ._ceb ;_cfc :=_bad .Params [2];_ggg :=_bad .Params [3]> 0;_abg :=_bad .Params [4]> 0;_db ,_fdg :=_bad .Params [5]*_ed ._ceb ,_bad .Params [6]*_ed ._ceb ;
if !_bad .isAbsolute (){_db ,_fdg =_cff +_db ,_abd +_fdg ;};_ddg :=_dd .EllipseToCubicBeziers (_cff ,_abd ,_gcb ,_cae ,_cfc ,_ggg ,_abg ,_db ,_fdg );for _ ,_aaa :=range _ddg {_ff .Add_c (_bafg (_aaa [1].X ,3),_bafg ((_aaa [1].Y ),3),_bafg ((_aaa [2].X ),3),_bafg ((_aaa [2].Y ),3),_bafg ((_aaa [3].X ),3),_bafg ((_aaa [3].Y ),3));
};_cff ,_abd =_db ,_fdg ;case "\u007a":_ff .Add_h ();};_gb =_bad ;};};if _ed .Style .FillColor !=""&&_ed .Style .StrokeColor !=""{_ff .Add_B ();}else if _ed .Style .FillColor !=""{_ff .Add_f ();}else if _ed .Style .StrokeColor !=""{_ff .Add_S ();};_ff .Add_h ();
_ff .Add_Q ();};