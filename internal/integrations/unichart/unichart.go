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

package unichart ;import (_g "bytes";_b "fmt";_ge "github.com/unidoc/unichart/render";_afd "github.com/unidoc/unipdf/v3/common";_e "github.com/unidoc/unipdf/v3/contentstream";_afa "github.com/unidoc/unipdf/v3/contentstream/draw";_gb "github.com/unidoc/unipdf/v3/core";
_d "github.com/unidoc/unipdf/v3/model";_bf "image/color";_af "io";_f "math";);func _dgb (_dce float64 )float64 {return _dce *_f .Pi /180.0};func (_gbg *Renderer )Close (){_gbg ._ee .Add_h ()};func (_fa *Renderer )ResetStyle (){_fa .SetFillColor (_bf .Black );
_fa .SetStrokeColor (_bf .Transparent );_fa .SetStrokeWidth (0);_fa .SetFont (_d .DefaultFont ());_fa .SetFontColor (_bf .Black );_fa .SetFontSize (12);_fa .SetTextRotation (0);};func (_dbc *Renderer )SetFontColor (color _bf .Color ){_dbc ._ff =color };
func _bdg (_cge _bf .Color )(uint8 ,uint8 ,uint8 ,uint8 ){_gdaf ,_aef ,_ecce ,_efb :=_cge .RGBA ();return uint8 (_gdaf >>8),uint8 (_aef >>8),uint8 (_ecce >>8),uint8 (_efb >>8);};func (_da *Renderer )SetDPI (dpi float64 ){_da ._c =dpi };func (_aeg *Renderer )getTextWidth (_fedd string )float64 {var _acc float64 ;
for _ ,_ddg :=range _fedd {_baa ,_aae :=_aeg ._bb .GetRuneMetrics (_ddg );if !_aae {_afd .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074",_ddg );
};_acc +=_baa .Wx ;};return _aeg ._bfd *_acc /1000.0;};func (_ebd *Renderer )MeasureText (text string )_ge .Box {_dad :=_ebd ._bfd ;_efc ,_dc :=_ebd ._bb .GetFontDescriptor ();if _dc !=nil {_afd .Log .Debug ("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072");
}else {_gaa ,_afgb :=_efc .GetCapHeight ();if _afgb !=nil {_afd .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076",_afgb );
}else {_dad =_gaa /1000.0*_ebd ._bfd ;};};var (_baf =0.0;_gc =_ebd .wrapText (text ););for _ ,_gac :=range _gc {if _dbd :=_ebd .getTextWidth (_gac );_dbd > _baf {_baf =_dbd ;};};_beg :=_ge .NewBox (0,0,int (_baf ),int (_dad ));if _gbe :=_ebd ._fg ;_gbe !=0{_beg =_beg .Corners ().Rotate (_gbe ).Box ();
};return _beg ;};func (_fce *Renderer )Text (text string ,x ,y int ){_fce ._ee .Add_q ();_fce .SetFont (_fce ._bb );_fd ,_ed ,_eb ,_ :=_fcb (_fce ._ff );_fce ._ee .Add_rg (_fd ,_ed ,_eb );_fce ._ee .Translate (float64 (x ),float64 (y )).Scale (1,-1);if _cgc :=_fce ._fg ;
_cgc !=0{_fce ._ee .RotateDeg (_cgc );};_fce ._ee .Add_BT ().Add_TL (_fce ._bfd );var (_cdc =_fce ._bb .Encoder ();_fed =_fce .wrapText (text );_dbg =len (_fed ););for _ccf ,_fga :=range _fed {_fce ._ee .Add_TJ (_gb .MakeStringFromBytes (_cdc .Encode (_fga )));
if _ccf !=_dbg -1{_fce ._ee .Add_Tstar ();};};_fce ._ee .Add_ET ();_fce ._ee .Add_Q ();};func (_fb *Renderer )FillStroke (){_fb ._ee .Add_B ()};func (_abb *Renderer )Fill (){_abb ._ee .Add_f ()};func (_gde *Renderer )ArcTo (cx ,cy int ,rx ,ry ,startAngle ,deltaAngle float64 ){startAngle =_fda (2.0*_f .Pi -startAngle );
deltaAngle =_fda (-deltaAngle );_eeb ,_agc :=deltaAngle ,1;if _f .Abs (deltaAngle )> 90.0{_agc =int (_f .Ceil (_f .Abs (deltaAngle )/90.0));_eeb =deltaAngle /float64 (_agc );};var (_aa =_dgb (_eeb /2);_fcg =_f .Abs (4.0/3.0*(1.0-_f .Cos (_aa ))/_f .Sin (_aa ));
_db =float64 (cx );_bfe =float64 (cy ););for _fab :=0;_fab < _agc ;_fab ++{_fcgb :=_dgb (startAngle +float64 (_fab )*_eeb );_gg :=_dgb (startAngle +float64 (_fab +1)*_eeb );_gec :=_f .Cos (_fcgb );_ea :=_f .Cos (_gg );_bd :=_f .Sin (_fcgb );_gab :=_f .Sin (_gg );
var _dg []float64 ;if _eeb > 0{_dg =[]float64 {_db +rx *_gec ,_bfe -ry *_bd ,_db +rx *(_gec -_fcg *_bd ),_bfe -ry *(_bd +_fcg *_gec ),_db +rx *(_ea +_fcg *_gab ),_bfe -ry *(_gab -_fcg *_ea ),_db +rx *_ea ,_bfe -ry *_gab };}else {_dg =[]float64 {_db +rx *_gec ,_bfe -ry *_bd ,_db +rx *(_gec +_fcg *_bd ),_bfe -ry *(_bd -_fcg *_gec ),_db +rx *(_ea -_fcg *_gab ),_bfe -ry *(_gab +_fcg *_ea ),_db +rx *_ea ,_bfe -ry *_gab };
};if _fab ==0{_gde ._ee .Add_l (_dg [0],_dg [1]);};_gde ._ee .Add_c (_dg [2],_dg [3],_dg [4],_dg [5],_dg [6],_dg [7]);};};func _fcb (_cbg _bf .Color )(float64 ,float64 ,float64 ,float64 ){_aag ,_dge ,_dfa ,_fdg :=_bdg (_cbg );return float64 (_aag )/255,float64 (_dge )/255,float64 (_dfa )/255,float64 (_fdg )/255;
};func (_ceg *Renderer )SetStrokeWidth (width float64 ){_ceg ._ef =width ;_ceg ._ee .Add_w (width )};func (_ba *Renderer )LineTo (x ,y int ){_ba ._ee .Add_l (float64 (x ),float64 (y ))};func (_def *Renderer )SetStrokeDashArray (dashArray []float64 ){_cb :=make ([]int64 ,len (dashArray ));
for _bee ,_gd :=range dashArray {_cb [_bee ]=int64 (_gd );};_def ._ee .Add_d (_cb ,0);};func (_fe *Renderer )Stroke (){_fe ._ee .Add_S ()};func (_aab *Renderer )Circle (radius float64 ,x ,y int ){_ged :=radius ;if _ac :=_aab ._ef ;_ac !=0{_ged -=_ac /2;
};_gga :=_ged *0.551784;_eef :=_afa .CubicBezierPath {Curves :[]_afa .CubicBezierCurve {_afa .NewCubicBezierCurve (-_ged ,0,-_ged ,_gga ,-_gga ,_ged ,0,_ged ),_afa .NewCubicBezierCurve (0,_ged ,_gga ,_ged ,_ged ,_gga ,_ged ,0),_afa .NewCubicBezierCurve (_ged ,0,_ged ,-_gga ,_gga ,-_ged ,0,-_ged ),_afa .NewCubicBezierCurve (0,-_ged ,-_gga ,-_ged ,-_ged ,-_gga ,-_ged ,0)}};
if _dab :=_aab ._ef ;_dab !=0{_eef =_eef .Offset (_dab /2,_dab /2);};_eef =_eef .Offset (float64 (x ),float64 (y ));_afa .DrawBezierPathWithCreator (_eef ,_aab ._ee );};func (_eeg *Renderer )SetStrokeColor (color _bf .Color ){_eeg ._de =color ;_dd ,_ae ,_dea ,_ :=_fcb (color );
_eeg ._ee .Add_RG (_dd ,_ae ,_dea );};func (_aad *Renderer )ClearTextRotation (){_aad ._fg =0};func (_ga *Renderer )MoveTo (x ,y int ){_ga ._ee .Add_m (float64 (x ),float64 (y ))};func (_cd *Renderer )SetFillColor (color _bf .Color ){_cd ._egb =color ;
_fff ,_ce ,_bg ,_ :=_fcb (color );_cd ._ee .Add_rg (_fff ,_ce ,_bg );};func (_abf *Renderer )SetClassName (name string ){};func _fda (_bfc float64 )float64 {return _bfc *180/_f .Pi };func (_gabd *Renderer )SetTextRotation (radians float64 ){_gabd ._fg =_fda (-radians )};
type Renderer struct{_df int ;_eg int ;_c float64 ;_ee *_e .ContentCreator ;_ag *_d .PdfPageResources ;_egb _bf .Color ;_de _bf .Color ;_ef float64 ;_bb *_d .PdfFont ;_bfd float64 ;_ff _bf .Color ;_fg float64 ;_ab map[*_d .PdfFont ]_gb .PdfObjectName ;
};func (_cc *Renderer )SetFont (font _ge .Font ){_cbe ,_afg :=font .(*_d .PdfFont );if !_afg {_afd .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065");return ;
};_gedb ,_afg :=_cc ._ab [_cbe ];if !_afg {_gedb =_cdb ("\u0046\u006f\u006e\u0074",1,_cc ._ag .HasFontByName );if _cg :=_cc ._ag .SetFontByName (_gedb ,_cbe .ToPdfObject ());_cg !=nil {_afd .Log .Debug ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073",_cbe );
};_cc ._ab [_cbe ]=_gedb ;};_cc ._ee .Add_Tf (_gedb ,_cc ._bfd );_cc ._bb =_cbe ;};func NewRenderer (cc *_e .ContentCreator ,res *_d .PdfPageResources )func (int ,int )(_ge .Renderer ,error ){return func (_efa ,_fc int )(_ge .Renderer ,error ){_aff :=&Renderer {_df :_efa ,_eg :_fc ,_c :72,_ee :cc ,_ag :res ,_ab :map[*_d .PdfFont ]_gb .PdfObjectName {}};
_aff .ResetStyle ();return _aff ,nil ;};};func _cdb (_geb string ,_fca int ,_defc func (_gb .PdfObjectName )bool )_gb .PdfObjectName {_ad :=_gb .PdfObjectName (_b .Sprintf ("\u0025\u0073\u0025\u0064",_geb ,_fca ));for _ca :=_fca ;_defc (_ad );{_ca ++;_ad =_gb .PdfObjectName (_b .Sprintf ("\u0025\u0073\u0025\u0064",_geb ,_ca ));
};return _ad ;};func (_cbf *Renderer )Save (w _af .Writer )error {if w ==nil {return nil ;};_ ,_agcb :=_af .Copy (w ,_g .NewBuffer (_cbf ._ee .Bytes ()));return _agcb ;};func (_bba *Renderer )QuadCurveTo (cx ,cy ,x ,y int ){_bba ._ee .Add_v (float64 (x ),float64 (y ),float64 (cx ),float64 (cy ));
};func (_bfb *Renderer )GetDPI ()float64 {return _bfb ._c };func (_ec *Renderer )SetFontSize (size float64 ){_ec ._bfd =size };func (_fcc *Renderer )wrapText (_ebb string )[]string {var (_fabf []string ;_cea []rune ;);for _ ,_bgc :=range _ebb {if _bgc =='\n'{_fabf =append (_fabf ,string (_cea ));
_cea =[]rune {};continue ;};_cea =append (_cea ,_bgc );};if len (_cea )> 0{_fabf =append (_fabf ,string (_cea ));};return _fabf ;};