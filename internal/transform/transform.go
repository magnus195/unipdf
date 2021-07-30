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

package transform ;import (_d "fmt";_db "github.com/unidoc/unipdf/v3/common";_gc "math";);func (_ad Matrix )Translation ()(float64 ,float64 ){return _ad [6],_ad [7]};func (_be *Matrix )Clone ()Matrix {return NewMatrix (_be [0],_be [1],_be [3],_be [4],_be [6],_be [7])};
const _dg =1.0e-6;func NewMatrixFromTransforms (xScale ,yScale ,theta ,tx ,ty float64 )Matrix {return IdentityMatrix ().Scale (xScale ,yScale ).Rotate (theta ).Translate (tx ,ty );};func ScaleMatrix (x ,y float64 )Matrix {return NewMatrix (x ,0,0,y ,0,0)};
func ShearMatrix (x ,y float64 )Matrix {return NewMatrix (1,y ,x ,1,0,0)};func RotationMatrix (angle float64 )Matrix {_c :=_gc .Cos (angle );_gg :=_gc .Sin (angle );return NewMatrix (_c ,_gg ,-_gg ,_c ,0,0);};func TranslationMatrix (tx ,ty float64 )Matrix {return NewMatrix (1,0,0,1,tx ,ty )};
func (_ed Matrix )Rotate (theta float64 )Matrix {return _ed .Mult (RotationMatrix (theta ))};func (_ge Matrix )Transform (x ,y float64 )(float64 ,float64 ){_fc :=x *_ge [0]+y *_ge [3]+_ge [6];_ga :=x *_ge [1]+y *_ge [4]+_ge [7];return _fc ,_ga ;};func (_a Matrix )Round (precision float64 )Matrix {for _e :=range _a {_a [_e ]=_gc .Round (_a [_e ]/precision )*precision ;
};return _a ;};type Matrix [9]float64 ;func (_dbg *Point )Set (x ,y float64 ){_dbg .X ,_dbg .Y =x ,y };func (_cc Matrix )Translate (tx ,ty float64 )Matrix {return _cc .Mult (TranslationMatrix (tx ,ty ))};type Point struct{X float64 ;Y float64 ;};func (_cgb Point )Interpolate (b Point ,t float64 )Point {return Point {X :(1-t )*_cgb .X +t *b .X ,Y :(1-t )*_cgb .Y +t *b .Y };
};func (_cb Matrix )Mult (b Matrix )Matrix {_cb .Concat (b );return _cb };func NewPoint (x ,y float64 )Point {return Point {X :x ,Y :y }};const _bee =1e-6;func NewMatrix (a ,b ,c ,d ,tx ,ty float64 )Matrix {_aa :=Matrix {a ,b ,0,c ,d ,0,tx ,ty ,1};_aa .clampRange ();
return _aa ;};func (_bc *Matrix )Shear (x ,y float64 ){_bc .Concat (ShearMatrix (x ,y ))};func IdentityMatrix ()Matrix {return NewMatrix (1,0,0,1,0,0)};func (_ef Matrix )Identity ()bool {return _ef [0]==1&&_ef [1]==0&&_ef [2]==0&&_ef [3]==0&&_ef [4]==1&&_ef [5]==0&&_ef [6]==0&&_ef [7]==0&&_ef [8]==1;
};func (_aef *Matrix )Set (a ,b ,c ,d ,tx ,ty float64 ){_aef [0],_aef [1]=a ,b ;_aef [3],_aef [4]=c ,d ;_aef [6],_aef [7]=tx ,ty ;_aef .clampRange ();};func (_cd Matrix )String ()string {_af ,_f ,_b ,_ae ,_ggf ,_ce :=_cd [0],_cd [1],_cd [3],_cd [4],_cd [6],_cd [7];
return _d .Sprintf ("\u005b\u00257\u002e\u0034\u0066\u002c%\u0037\u002e4\u0066\u002c\u0025\u0037\u002e\u0034\u0066\u002c%\u0037\u002e\u0034\u0066\u003a\u0025\u0037\u002e\u0034\u0066\u002c\u00257\u002e\u0034\u0066\u005d",_af ,_f ,_b ,_ae ,_ggf ,_ce );};
func (_ac Matrix )Inverse ()(Matrix ,bool ){_efa ,_ea :=_ac [0],_ac [1];_geb ,_gb :=_ac [3],_ac [4];_cg ,_eff :=_ac [6],_ac [7];_cgd :=_efa *_gb -_ea *_geb ;if _gc .Abs (_cgd )< _dg {return Matrix {},false ;};_gaf ,_ec :=_gb /_cgd ,-_ea /_cgd ;_aad ,_df :=-_geb /_cgd ,_efa /_cgd ;
_ead :=-(_gaf *_cg +_aad *_eff );_ff :=-(_ec *_cg +_df *_eff );return NewMatrix (_gaf ,_ec ,_aad ,_df ,_ead ,_ff ),true ;};func (_ace *Matrix )clampRange (){for _fe ,_gd :=range _ace {if _gd > _cf {_db .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_gd ,_cf );
_ace [_fe ]=_cf ;}else if _gd < -_cf {_db .Log .Debug ("\u0043L\u0041M\u0050\u003a\u0020\u0025\u0067\u0020\u002d\u003e\u0020\u0025\u0067",_gd ,-_cf );_ace [_fe ]=-_cf ;};};};func (_fb Matrix )ScalingFactorX ()float64 {return _gc .Hypot (_fb [0],_fb [1])};
func (_aab Matrix )Scale (xScale ,yScale float64 )Matrix {return _aab .Mult (ScaleMatrix (xScale ,yScale ));};const _cf =1e9;func (_cfe *Point )Transform (a ,b ,c ,d ,tx ,ty float64 ){_bd :=NewMatrix (a ,b ,c ,d ,tx ,ty );_cfe .transformByMatrix (_bd );
};func (_edg *Point )transformByMatrix (_geff Matrix ){_edg .X ,_edg .Y =_geff .Transform (_edg .X ,_edg .Y )};func (_bb Point )Rotate (theta float64 )Point {_fg :=_gc .Hypot (_bb .X ,_bb .Y );_bed :=_gc .Atan2 (_bb .Y ,_bb .X );_ccf ,_cec :=_gc .Sincos (_bed +theta /180.0*_gc .Pi );
return Point {_fg *_cec ,_fg *_ccf };};func (_ccg Point )Displace (delta Point )Point {return Point {_ccg .X +delta .X ,_ccg .Y +delta .Y }};func (_gf Matrix )ScalingFactorY ()float64 {return _gc .Hypot (_gf [3],_gf [4])};func (_beb Matrix )Angle ()float64 {_dbd :=_gc .Atan2 (-_beb [1],_beb [0]);
if _dbd < 0.0{_dbd +=2*_gc .Pi ;};return _dbd /_gc .Pi *180.0;};const _eb =1e-10;func (_dgf Point )Distance (b Point )float64 {return _gc .Hypot (_dgf .X -b .X ,_dgf .Y -b .Y )};func (_gef Point )String ()string {return _d .Sprintf ("(\u0025\u002e\u0032\u0066\u002c\u0025\u002e\u0032\u0066\u0029",_gef .X ,_gef .Y );
};func (_gcg Matrix )Singular ()bool {return _gc .Abs (_gcg [0]*_gcg [4]-_gcg [1]*_gcg [3])< _eb };func (_bg *Matrix )Concat (b Matrix ){*_bg =Matrix {b [0]*_bg [0]+b [1]*_bg [3],b [0]*_bg [1]+b [1]*_bg [4],0,b [3]*_bg [0]+b [4]*_bg [3],b [3]*_bg [1]+b [4]*_bg [4],0,b [6]*_bg [0]+b [7]*_bg [3]+_bg [6],b [6]*_bg [1]+b [7]*_bg [4]+_bg [7],1};
_bg .clampRange ();};func (_ab Matrix )Unrealistic ()bool {_fd ,_ccb ,_da ,_bcc :=_gc .Abs (_ab [0]),_gc .Abs (_ab [1]),_gc .Abs (_ab [3]),_gc .Abs (_ab [4]);_efd :=_fd > _bee &&_bcc > _bee ;_cbd :=_ccb > _bee &&_da > _bee ;return !(_efd ||_cbd );};