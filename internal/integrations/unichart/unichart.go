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

package unichart ;import (_d "bytes";_f "fmt";_b "github.com/unidoc/unichart/render";
	_ee "github.com/magnus195/unipdf/v3/common";
	_bg "github.com/magnus195/unipdf/v3/contentstream";
	_ff "github.com/magnus195/unipdf/v3/contentstream/draw";
	_bd "github.com/magnus195/unipdf/v3/core";
	_g "github.com/magnus195/unipdf/v3/model";_e "image/color";_dg "io";_de "math";);func (_fba *Renderer )Text (text string ,x ,y int ){_fba ._fb .Add_q ();_fba .SetFont (_fba ._dga );_gbg ,_dd ,_aa ,_ :=_dgae (_fba ._gc );_fba ._fb .Add_rg (_gbg ,_dd ,_aa );
_fba ._fb .Translate (float64 (x ),float64 (y )).Scale (1,-1);if _ebg :=_fba ._cg ;_ebg !=0{_fba ._fb .RotateDeg (_ebg );};_fba ._fb .Add_BT ().Add_TL (_fba ._ce );var (_db =_fba ._dga .Encoder ();_gf =_fba .wrapText (text );_bb =len (_gf ););for _fdba ,_gde :=range _gf {_fba ._fb .Add_TJ (_bd .MakeStringFromBytes (_db .Encode (_gde )));
if _fdba !=_bb -1{_fba ._fb .Add_Tstar ();};};_fba ._fb .Add_ET ();_fba ._fb .Add_Q ();};func (_gg *Renderer )MeasureText (text string )_b .Box {_fcea :=_gg ._ce ;_bag ,_feb :=_gg ._dga .GetFontDescriptor ();if _feb !=nil {_ee .Log .Debug ("W\u0041\u0052\u004e\u003a\u0020\u0055n\u0061\u0062\u006c\u0065\u0020\u0074o\u0020\u0067\u0065\u0074\u0020\u0066\u006fn\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069\u0070\u0074o\u0072");
}else {_dce ,_dcf :=_bag .GetCapHeight ();if _dcf !=nil {_ee .Log .Debug ("\u0057\u0041\u0052\u004e\u003a\u0020\u0055\u006e\u0061\u0062\u006c\u0065\u0020t\u006f\u0020\u0067\u0065\u0074\u0020f\u006f\u006e\u0074\u0020\u0063\u0061\u0070\u0020\u0068\u0065\u0069\u0067\u0068t\u003a\u0020\u0025\u0076",_dcf );
}else {_fcea =_dce /1000.0*_gg ._ce ;};};var (_gbd =0.0;_dbg =_gg .wrapText (text ););for _ ,_fdf :=range _dbg {if _eag :=_gg .getTextWidth (_fdf );_eag > _gbd {_gbd =_eag ;};};_efg :=_b .NewBox (0,0,int (_gbd ),int (_fcea ));if _dcd :=_gg ._cg ;_dcd !=0{_efg =_efg .Corners ().Rotate (_dcd ).Box ();
};return _efg ;};func _dbd (_ac float64 )float64 {return _ac *180/_de .Pi };func (_cgd *Renderer )SetDPI (dpi float64 ){_cgd ._gb =dpi };func (_fd *Renderer )LineTo (x ,y int ){_fd ._fb .Add_l (float64 (x ),float64 (y ))};func _bgg (_bdc _e .Color )(uint8 ,uint8 ,uint8 ,uint8 ){_ada ,_aff ,_ab ,_bgb :=_bdc .RGBA ();
return uint8 (_ada >>8),uint8 (_aff >>8),uint8 (_ab >>8),uint8 (_bgb >>8);};func (_cadg *Renderer )SetFontSize (size float64 ){_cadg ._ce =size };func (_dbb *Renderer )wrapText (_cff string )[]string {var (_daa []string ;_gdf []rune ;);for _ ,_ggc :=range _cff {if _ggc =='\n'{_daa =append (_daa ,string (_gdf ));
_gdf =[]rune {};continue ;};_gdf =append (_gdf ,_ggc );};if len (_gdf )> 0{_daa =append (_daa ,string (_gdf ));};return _daa ;};func NewRenderer (cc *_bg .ContentCreator ,res *_g .PdfPageResources )func (int ,int )(_b .Renderer ,error ){return func (_ec ,_eg int )(_b .Renderer ,error ){_dc :=&Renderer {_fg :_ec ,_fa :_eg ,_gb :72,_fb :cc ,_ge :res ,_ea :map[*_g .PdfFont ]_bd .PdfObjectName {}};
_dc .ResetStyle ();return _dc ,nil ;};};func (_bf *Renderer )SetStrokeWidth (width float64 ){_bf ._c =width ;_bf ._fb .Add_w (width )};func (_cge *Renderer )QuadCurveTo (cx ,cy ,x ,y int ){_cge ._fb .Add_v (float64 (x ),float64 (y ),float64 (cx ),float64 (cy ));
};func (_ebc *Renderer )Save (w _dg .Writer )error {if w ==nil {return nil ;};_ ,_eda :=_dg .Copy (w ,_d .NewBuffer (_ebc ._fb .Bytes ()));return _eda ;};func (_deg *Renderer )Close (){_deg ._fb .Add_h ()};func (_fdc *Renderer )SetFont (font _b .Font ){_ecd ,_ead :=font .(*_g .PdfFont );
if !_ead {_ee .Log .Debug ("\u0045R\u0052\u004f\u0052\u003a\u0020\u0069\u006e\u0076\u0061\u006c\u0069d\u0020\u0066\u006f\u006e\u0074\u0020\u0074\u0079\u0070\u0065");return ;};_caf ,_ead :=_fdc ._ea [_ecd ];if !_ead {_caf =_cbc ("\u0046\u006f\u006e\u0074",1,_fdc ._ge .HasFontByName );
if _bfg :=_fdc ._ge .SetFontByName (_caf ,_ecd .ToPdfObject ());_bfg !=nil {_ee .Log .Debug ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0061\u0064d\u0020\u0066\u006f\u006e\u0074\u0020\u0025\u0076\u0020\u0074\u006f\u0020\u0072\u0065\u0073\u006f\u0075\u0072\u0063\u0065\u0073",_ecd );
};_fdc ._ea [_ecd ]=_caf ;};_fdc ._fb .Add_Tf (_caf ,_fdc ._ce );_fdc ._dga =_ecd ;};func (_ga *Renderer )Circle (radius float64 ,x ,y int ){_afb :=radius ;if _dgd :=_ga ._c ;_dgd !=0{_afb -=_dgd /2;};_bab :=_afb *0.551784;_dcb :=_ff .CubicBezierPath {Curves :[]_ff .CubicBezierCurve {_ff .NewCubicBezierCurve (-_afb ,0,-_afb ,_bab ,-_bab ,_afb ,0,_afb ),_ff .NewCubicBezierCurve (0,_afb ,_bab ,_afb ,_afb ,_bab ,_afb ,0),_ff .NewCubicBezierCurve (_afb ,0,_afb ,-_bab ,_bab ,-_afb ,0,-_afb ),_ff .NewCubicBezierCurve (0,-_afb ,-_bab ,-_afb ,-_afb ,-_bab ,-_afb ,0)}};
if _gac :=_ga ._c ;_gac !=0{_dcb =_dcb .Offset (_gac /2,_gac /2);};_dcb =_dcb .Offset (float64 (x ),float64 (y ));_ff .DrawBezierPathWithCreator (_dcb ,_ga ._fb );};func (_ffcf *Renderer )SetStrokeDashArray (dashArray []float64 ){_egg :=make ([]int64 ,len (dashArray ));
for _ece ,_cdb :=range dashArray {_egg [_ece ]=int64 (_cdb );};_ffcf ._fb .Add_d (_egg ,0);};func _adf (_eee float64 )float64 {return _eee *_de .Pi /180.0};func (_fe *Renderer )SetFontColor (color _e .Color ){_fe ._gc =color };func (_ba *Renderer )Fill (){_ba ._fb .Add_f ()};
func (_bfc *Renderer )FillStroke (){_bfc ._fb .Add_B ()};func (_eb *Renderer )ArcTo (cx ,cy int ,rx ,ry ,startAngle ,deltaAngle float64 ){startAngle =_dbd (2.0*_de .Pi -startAngle );deltaAngle =_dbd (-deltaAngle );_fc ,_fgb :=deltaAngle ,1;if _de .Abs (deltaAngle )> 90.0{_fgb =int (_de .Ceil (_de .Abs (deltaAngle )/90.0));
_fc =deltaAngle /float64 (_fgb );};var (_ebf =_adf (_fc /2);_df =_de .Abs (4.0/3.0*(1.0-_de .Cos (_ebf ))/_de .Sin (_ebf ));_bdb =float64 (cx );_cf =float64 (cy ););for _eceg :=0;_eceg < _fgb ;_eceg ++{_fce :=_adf (startAngle +float64 (_eceg )*_fc );_fdb :=_adf (startAngle +float64 (_eceg +1)*_fc );
_fcb :=_de .Cos (_fce );_efcb :=_de .Cos (_fdb );_da :=_de .Sin (_fce );_dgc :=_de .Sin (_fdb );var _dfd []float64 ;if _fc > 0{_dfd =[]float64 {_bdb +rx *_fcb ,_cf -ry *_da ,_bdb +rx *(_fcb -_df *_da ),_cf -ry *(_da +_df *_fcb ),_bdb +rx *(_efcb +_df *_dgc ),_cf -ry *(_dgc -_df *_efcb ),_bdb +rx *_efcb ,_cf -ry *_dgc };
}else {_dfd =[]float64 {_bdb +rx *_fcb ,_cf -ry *_da ,_bdb +rx *(_fcb +_df *_da ),_cf -ry *(_da -_df *_fcb ),_bdb +rx *(_efcb -_df *_dgc ),_cf -ry *(_dgc +_df *_efcb ),_bdb +rx *_efcb ,_cf -ry *_dgc };};if _eceg ==0{_eb ._fb .Add_l (_dfd [0],_dfd [1]);
};_eb ._fb .Add_c (_dfd [2],_dfd [3],_dfd [4],_dfd [5],_dfd [6],_dfd [7]);};};type Renderer struct{_fg int ;_fa int ;_gb float64 ;_fb *_bg .ContentCreator ;_ge *_g .PdfPageResources ;_fgc _e .Color ;_ef _e .Color ;_c float64 ;_dga *_g .PdfFont ;_ce float64 ;
_gc _e .Color ;_cg float64 ;_ea map[*_g .PdfFont ]_bd .PdfObjectName ;};func (_ecc *Renderer )ResetStyle (){_ecc .SetFillColor (_e .Black );_ecc .SetStrokeColor (_e .Transparent );_ecc .SetStrokeWidth (0);_ecc .SetFont (_g .DefaultFont ());_ecc .SetFontColor (_e .Black );
_ecc .SetFontSize (12);_ecc .SetTextRotation (0);};func (_gd *Renderer )Stroke (){_gd ._fb .Add_S ()};func (_cb *Renderer )SetTextRotation (radians float64 ){_cb ._cg =_dbd (-radians )};func (_edf *Renderer )getTextWidth (_cag string )float64 {var _gae float64 ;
for _ ,_fgdf :=range _cag {_dac ,_bdbg :=_edf ._dga .GetRuneMetrics (_fgdf );if !_bdbg {_ee .Log .Debug ("\u0045\u0052\u0052OR\u003a\u0020\u0075\u006e\u0073\u0075\u0070\u0070\u006fr\u0074e\u0064 \u0072u\u006e\u0065\u0020\u0025\u0076\u0020\u0069\u006e\u0020\u0066\u006f\u006e\u0074",_fgdf );
};_gae +=_dac .Wx ;};return _edf ._ce *_gae /1000.0;};func (_cafc *Renderer )ClearTextRotation (){_cafc ._cg =0};func (_ded *Renderer )MoveTo (x ,y int ){_ded ._fb .Add_m (float64 (x ),float64 (y ))};func (_cgf *Renderer )SetStrokeColor (color _e .Color ){_cgf ._ef =color ;
_ca ,_fgd ,_ffc ,_ :=_dgae (color );_cgf ._fb .Add_RG (_ca ,_fgd ,_ffc );};func (_ecb *Renderer )SetFillColor (color _e .Color ){_ecb ._fgc =color ;_cad ,_dge ,_af ,_ :=_dgae (color );_ecb ._fb .Add_rg (_cad ,_dge ,_af );};func (_efc *Renderer )GetDPI ()float64 {return _efc ._gb };
func _cbc (_ggd string ,_ffg int ,_baa func (_bd .PdfObjectName )bool )_bd .PdfObjectName {_ag :=_bd .PdfObjectName (_f .Sprintf ("\u0025\u0073\u0025\u0064",_ggd ,_ffg ));for _cc :=_ffg ;_baa (_ag );{_cc ++;_ag =_bd .PdfObjectName (_f .Sprintf ("\u0025\u0073\u0025\u0064",_ggd ,_cc ));
};return _ag ;};func _dgae (_dgaf _e .Color )(float64 ,float64 ,float64 ,float64 ){_ffa ,_dedc ,_dec ,_efgc :=_bgg (_dgaf );return float64 (_ffa )/255,float64 (_dedc )/255,float64 (_dec )/255,float64 (_efgc )/255;};func (_bgf *Renderer )SetClassName (name string ){};
