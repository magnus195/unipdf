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

package classer ;import (_d "github.com/unidoc/unipdf/v3/common";_c "github.com/unidoc/unipdf/v3/internal/jbig2/basic";_gf "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_af "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_b "image";_a "math";
);var TwoByTwoWalk =[]int {0,0,0,1,-1,0,0,-1,1,0,-1,1,1,1,-1,-1,1,-1,0,-2,2,0,0,2,-2,0,-1,-2,1,-2,2,-1,2,1,1,2,-1,2,-2,1,-2,-1,-2,-2,2,-2,2,2,-2,2};func _bdd (_gff *_gf .Bitmap ,_cg ,_fa ,_dfg ,_fdd int ,_gef *_gf .Bitmap )(_bgc _b .Point ,_ga error ){const _cc ="\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg";
if _gff ==nil {return _bgc ,_af .Error (_cc ,"\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};if _gef ==nil {return _bgc ,_af .Error (_cc ,"t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064");
};_ff ,_fca :=_gef .Width ,_gef .Height ;_eff ,_fda :=_cg -_dfg -JbAddedPixels ,_fa -_fdd -JbAddedPixels ;_d .Log .Trace ("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027",_cg ,_fa ,_ff ,_fca ,_eff ,_fda );
_fbe ,_ga :=_gf .Rect (_eff ,_fda ,_ff ,_fca );if _ga !=nil {return _bgc ,_af .Wrap (_ga ,_cc ,"");};_aca ,_ ,_ga :=_gff .ClipRectangle (_fbe );if _ga !=nil {_d .Log .Error ("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076",_fbe );
return _bgc ,_af .Wrap (_ga ,_cc ,"");};_ab :=_gf .New (_aca .Width ,_aca .Height );_gec :=_a .MaxInt32 ;var _fbf ,_bgf ,_bgg ,_dfgd ,_bb int ;for _fbf =-1;_fbf <=1;_fbf ++{for _bgf =-1;_bgf <=1;_bgf ++{if _ ,_ga =_gf .Copy (_ab ,_aca );_ga !=nil {return _bgc ,_af .Wrap (_ga ,_cc ,"");
};if _ga =_ab .RasterOperation (_bgf ,_fbf ,_ff ,_fca ,_gf .PixSrcXorDst ,_gef ,0,0);_ga !=nil {return _bgc ,_af .Wrap (_ga ,_cc ,"");};_bgg =_ab .CountPixels ();if _bgg < _gec {_dfgd =_bgf ;_bb =_fbf ;_gec =_bgg ;};};};_bgc .X =_dfgd ;_bgc .Y =_bb ;return _bgc ,nil ;
};func Init (settings Settings )(*Classer ,error ){const _e ="\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074";_eb :=&Classer {Settings :settings ,Widths :map[int ]int {},Heights :map[int ]int {},TemplatesSize :_c .IntsMap {},TemplateAreas :&_c .IntSlice {},ComponentPageNumbers :&_c .IntSlice {},ClassIDs :&_c .IntSlice {},ComponentsNumber :&_c .IntSlice {},CentroidPoints :&_gf .Points {},CentroidPointsTemplates :&_gf .Points {},UndilatedTemplates :&_gf .Bitmaps {},DilatedTemplates :&_gf .Bitmaps {},ClassInstances :&_gf .BitmapsArray {},FgTemplates :&_c .NumSlice {}};
if _ge :=_eb .Settings .Validate ();_ge !=nil {return nil ,_af .Wrap (_ge ,_e ,"");};return _eb ,nil ;};const (MaxDiffWidth =2;MaxDiffHeight =2;);func (_dff *Classer )classifyRankHouseNonOne (_cdgg *_gf .Boxes ,_bcf ,_fabg ,_dgd *_gf .Bitmaps ,_ddf *_gf .Points ,_egb *_c .NumSlice ,_afef int )(_aab error ){const _ccbd ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_cgb ,_ggbc ,_ddbe ,_dgfa float32 ;_dad ,_acaf ,_abf int ;_acdb ,_cbc ,_cfbb ,_ecb ,_fdaa *_gf .Bitmap ;_fdg ,_afdc bool ;);_cdd :=_gf .MakePixelSumTab8 ();for _fdac :=0;_fdac < len (_bcf .Values );_fdac ++{if _cbc ,_aab =_fabg .GetBitmap (_fdac );
_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _dad ,_aab =_egb .GetInt (_fdac );_aab !=nil {_d .Log .Trace ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_fdac ,_aab );
};if _cfbb ,_aab =_dgd .GetBitmap (_fdac );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _cgb ,_ggbc ,_aab =_ddf .GetGeometry (_fdac );_aab !=nil {return _af .Wrapf (_aab ,_ccbd ,"\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079");
};_aac :=len (_dff .UndilatedTemplates .Values );_fdg =false ;_dba :=_fbb (_dff ,_cbc );for _abf =_dba .Next ();_abf > -1;{if _ecb ,_aab =_dff .UndilatedTemplates .GetBitmap (_abf );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");
};if _acaf ,_aab =_dff .FgTemplates .GetInt (_abf );_aab !=nil {_d .Log .Trace ("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_abf ,_aab );
};if _fdaa ,_aab =_dff .DilatedTemplates .GetBitmap (_abf );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _ddbe ,_dgfa ,_aab =_dff .CentroidPointsTemplates .GetGeometry (_abf );
_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d");};_afdc ,_aab =_gf .RankHausTest (_cbc ,_cfbb ,_ecb ,_fdaa ,_cgb -_ddbe ,_ggbc -_dgfa ,MaxDiffWidth ,MaxDiffHeight ,_dad ,_acaf ,float32 (_dff .Settings .RankHaus ),_cdd );
if _aab !=nil {return _af .Wrap (_aab ,_ccbd ,"");};if _afdc {_fdg =true ;if _aab =_dff .ClassIDs .Add (_abf );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"");};if _aab =_dff .ComponentPageNumbers .Add (_afef );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"");
};if _dff .Settings .KeepClassInstances {_dag ,_afad :=_dff .ClassInstances .GetBitmaps (_abf );if _afad !=nil {return _af .Wrap (_afad ,_ccbd ,"\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029");
};if _acdb ,_afad =_bcf .GetBitmap (_fdac );_afad !=nil {return _af .Wrap (_afad ,_ccbd ,"\u0070i\u0078\u0061\u005b\u0069\u005d");};_dag .Values =append (_dag .Values ,_acdb );_aaf ,_afad :=_cdgg .Get (_fdac );if _afad !=nil {return _af .Wrap (_afad ,_ccbd ,"b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};_dag .Boxes =append (_dag .Boxes ,_aaf );};break ;};};if !_fdg {if _aab =_dff .ClassIDs .Add (_aac );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _aab =_dff .ComponentPageNumbers .Add (_afef );_aab !=nil {return _af .Wrap (_aab ,_ccbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_gcgd :=&_gf .Bitmaps {};_acdb =_bcf .Values [_fdac ];_gcgd .AddBitmap (_acdb );_aff ,_cef :=_acdb .Width ,_acdb .Height ;_dff .TemplatesSize .Add (uint64 (_aff )*uint64 (_cef ),_aac );_fcgb ,_gaa :=_cdgg .Get (_fdac );if _gaa !=nil {return _af .Wrap (_gaa ,_ccbd ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_gcgd .AddBox (_fcgb );_dff .ClassInstances .AddBitmaps (_gcgd );_dff .CentroidPointsTemplates .AddPoint (_cgb ,_ggbc );_dff .UndilatedTemplates .AddBitmap (_cbc );_dff .DilatedTemplates .AddBitmap (_cfbb );_dff .FgTemplates .AddInt (_dad );};};_dff .NumberOfClasses =len (_dff .UndilatedTemplates .Values );
return nil ;};func (_fcc *Classer )classifyRankHaus (_gbc *_gf .Boxes ,_bed *_gf .Bitmaps ,_ed int )error {const _ead ="\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073";if _gbc ==nil {return _af .Error (_ead ,"\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");
};if _bed ==nil {return _af .Error (_ead ,"\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};_aec :=len (_bed .Values );if _aec ==0{return _af .Error (_ead ,"e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073");
};_dg :=_bed .CountPixels ();_cdbc :=_fcc .Settings .SizeHaus ;_bca :=_gf .SelCreateBrick (_cdbc ,_cdbc ,_cdbc /2,_cdbc /2,_gf .SelHit );_ebdf :=&_gf .Bitmaps {Values :make ([]*_gf .Bitmap ,_aec )};_bea :=&_gf .Bitmaps {Values :make ([]*_gf .Bitmap ,_aec )};
var (_da ,_dgf ,_fge *_gf .Bitmap ;_eba error ;);for _ggg :=0;_ggg < _aec ;_ggg ++{_da ,_eba =_bed .GetBitmap (_ggg );if _eba !=nil {return _af .Wrap (_eba ,_ead ,"");};_dgf ,_eba =_da .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _eba !=nil {return _af .Wrap (_eba ,_ead ,"");};_fge ,_eba =_gf .Dilate (nil ,_dgf ,_bca );if _eba !=nil {return _af .Wrap (_eba ,_ead ,"");};_ebdf .Values [_aec ]=_dgf ;_bea .Values [_aec ]=_fge ;};_gab ,_eba :=_gf .Centroids (_ebdf .Values );if _eba !=nil {return _af .Wrap (_eba ,_ead ,"");
};if _eba =_gab .Add (_fcc .CentroidPoints );_eba !=nil {_d .Log .Trace ("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064");};if _fcc .Settings .RankHaus ==1.0{_eba =_fcc .classifyRankHouseOne (_gbc ,_bed ,_ebdf ,_bea ,_gab ,_ed );
}else {_eba =_fcc .classifyRankHouseNonOne (_gbc ,_bed ,_ebdf ,_bea ,_gab ,_dg ,_ed );};if _eba !=nil {return _af .Wrap (_eba ,_ead ,"");};return nil ;};func (_fcgd *Settings )SetDefault (){if _fcgd .MaxCompWidth ==0{switch _fcgd .Components {case _gf .ComponentConn :_fcgd .MaxCompWidth =MaxConnCompWidth ;
case _gf .ComponentCharacters :_fcgd .MaxCompWidth =MaxCharCompWidth ;case _gf .ComponentWords :_fcgd .MaxCompWidth =MaxWordCompWidth ;};};if _fcgd .MaxCompHeight ==0{_fcgd .MaxCompHeight =MaxCompHeight ;};if _fcgd .Thresh ==0.0{_fcgd .Thresh =0.9;};if _fcgd .WeightFactor ==0.0{_fcgd .WeightFactor =0.75;
};if _fcgd .RankHaus ==0.0{_fcgd .RankHaus =0.97;};if _fcgd .SizeHaus ==0{_fcgd .SizeHaus =2;};};func _fbb (_cec *Classer ,_eaf *_gf .Bitmap )*similarTemplatesFinder {return &similarTemplatesFinder {Width :_eaf .Width ,Height :_eaf .Height ,Classer :_cec };
};func (_ac *Classer )ComputeLLCorners ()(_fc error ){const _fd ="\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _ac .PtaUL ==nil {return _af .Error (_fd ,"\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064");
};_ae :=len (*_ac .PtaUL );_ac .PtaLL =&_gf .Points {};var (_afb ,_fg float32 ;_fcg ,_aa int ;_ee *_gf .Bitmap ;);for _bf :=0;_bf < _ae ;_bf ++{_afb ,_fg ,_fc =_ac .PtaUL .GetGeometry (_bf );if _fc !=nil {_d .Log .Debug ("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_fc );
return _af .Wrap (_fc ,_fd ,"\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079");};_fcg ,_fc =_ac .ClassIDs .Get (_bf );if _fc !=nil {_d .Log .Debug ("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076",_fc );
return _af .Wrap (_fc ,_fd ,"\u0043l\u0061\u0073\u0073\u0049\u0044");};_ee ,_fc =_ac .UndilatedTemplates .GetBitmap (_fcg );if _fc !=nil {_d .Log .Debug ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_fc );
return _af .Wrap (_fc ,_fd ,"\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_aa =_ee .Height ;_ac .PtaLL .AddPoint (_afb ,_fg +float32 (_aa ));};return nil ;};type Settings struct{MaxCompWidth int ;
MaxCompHeight int ;SizeHaus int ;RankHaus float64 ;Thresh float64 ;WeightFactor float64 ;KeepClassInstances bool ;Components _gf .Component ;Method Method ;};func (_de *Classer )addPageComponents (_gd *_gf .Bitmap ,_bd *_gf .Boxes ,_afa *_gf .Bitmaps ,_db int ,_gbe Method )error {const _ba ="\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s";
if _gd ==nil {return _af .Error (_ba ,"\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065");};if _bd ==nil ||_afa ==nil ||len (*_bd )==0{_d .Log .Trace ("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064",_gd );
return nil ;};var _gc error ;switch _gbe {case RankHaus :_gc =_de .classifyRankHaus (_bd ,_afa ,_db );case Correlation :_gc =_de .classifyCorrelation (_bd ,_afa ,_db );default:_d .Log .Debug ("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027",_gbe );
return _af .Error (_ba ,"\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064");};if _gc !=nil {return _af .Wrap (_gc ,_ba ,"");};if _gc =_de .getULCorners (_gd ,_bd );_gc !=nil {return _af .Wrap (_gc ,_ba ,"");
};_ddb :=len (*_bd );_de .BaseIndex +=_ddb ;if _gc =_de .ComponentsNumber .Add (_ddb );_gc !=nil {return _af .Wrap (_gc ,_ba ,"");};return nil ;};type similarTemplatesFinder struct{Classer *Classer ;Width int ;Height int ;Index int ;CurrentNumbers []int ;
N int ;};const (RankHaus Method =iota ;Correlation ;);const (MaxConnCompWidth =350;MaxCharCompWidth =350;MaxWordCompWidth =1000;MaxCompHeight =120;);func (_gea *Classer )AddPage (inputPage *_gf .Bitmap ,pageNumber int ,method Method )(_ebc error ){const _f ="\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065";
_gea .Widths [pageNumber ]=inputPage .Width ;_gea .Heights [pageNumber ]=inputPage .Height ;if _ebc =_gea .verifyMethod (method );_ebc !=nil {return _af .Wrap (_ebc ,_f ,"");};_dd ,_gb ,_ebc :=inputPage .GetComponents (_gea .Settings .Components ,_gea .Settings .MaxCompWidth ,_gea .Settings .MaxCompHeight );
if _ebc !=nil {return _af .Wrap (_ebc ,_f ,"");};_d .Log .Debug ("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076",_dd );if _ebc =_gea .addPageComponents (inputPage ,_gb ,_dd ,pageNumber ,method );_ebc !=nil {return _af .Wrap (_ebc ,_f ,"");
};return nil ;};func (_df *Classer )getULCorners (_dbb *_gf .Bitmap ,_bfd *_gf .Boxes )error {const _cd ="\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _dbb ==nil {return _af .Error (_cd ,"\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070");
};if _bfd ==nil {return _af .Error (_cd ,"\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};if _df .PtaUL ==nil {_df .PtaUL =&_gf .Points {};};_ef :=len (*_bfd );var (_gg ,_gdd ,_ca ,_bda int ;_fb ,_geag ,_cf ,_ebd float32 ;_gcc error ;_gdg *_b .Rectangle ;
_fe *_gf .Bitmap ;_be _b .Point ;);for _cab :=0;_cab < _ef ;_cab ++{_gg =_df .BaseIndex +_cab ;if _fb ,_geag ,_gcc =_df .CentroidPoints .GetGeometry (_gg );_gcc !=nil {return _af .Wrap (_gcc ,_cd ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073");
};if _gdd ,_gcc =_df .ClassIDs .Get (_gg );_gcc !=nil {return _af .Wrap (_gcc ,_cd ,"\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074");};if _cf ,_ebd ,_gcc =_df .CentroidPointsTemplates .GetGeometry (_gdd );_gcc !=nil {return _af .Wrap (_gcc ,_cd ,"\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");
};_bg :=_cf -_fb ;_acd :=_ebd -_geag ;if _bg >=0{_ca =int (_bg +0.5);}else {_ca =int (_bg -0.5);};if _acd >=0{_bda =int (_acd +0.5);}else {_bda =int (_acd -0.5);};if _gdg ,_gcc =_bfd .Get (_cab );_gcc !=nil {return _af .Wrap (_gcc ,_cd ,"");};_ec ,_ddd :=_gdg .Min .X ,_gdg .Min .Y ;
_fe ,_gcc =_df .UndilatedTemplates .GetBitmap (_gdd );if _gcc !=nil {return _af .Wrap (_gcc ,_cd ,"\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029");
};_be ,_gcc =_bdd (_dbb ,_ec ,_ddd ,_ca ,_bda ,_fe );if _gcc !=nil {return _af .Wrap (_gcc ,_cd ,"");};_df .PtaUL .AddPoint (float32 (_ec -_ca +_be .X ),float32 (_ddd -_bda +_be .Y ));};return nil ;};const JbAddedPixels =6;func (_aeee *similarTemplatesFinder )Next ()int {var (_bce ,_aae ,_cfge ,_eef int ;
_ecd bool ;_dbaf *_gf .Bitmap ;_fega error ;);for {if _aeee .Index >=25{return -1;};_aae =_aeee .Width +TwoByTwoWalk [2*_aeee .Index ];_bce =_aeee .Height +TwoByTwoWalk [2*_aeee .Index +1];if _bce < 1||_aae < 1{_aeee .Index ++;continue ;};if len (_aeee .CurrentNumbers )==0{_aeee .CurrentNumbers ,_ecd =_aeee .Classer .TemplatesSize .GetSlice (uint64 (_aae )*uint64 (_bce ));
if !_ecd {_aeee .Index ++;continue ;};_aeee .N =0;};_cfge =len (_aeee .CurrentNumbers );for ;_aeee .N < _cfge ;_aeee .N ++{_eef =_aeee .CurrentNumbers [_aeee .N ];_dbaf ,_fega =_aeee .Classer .DilatedTemplates .GetBitmap (_eef );if _fega !=nil {_d .Log .Debug ("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020");
return 0;};if _dbaf .Width -2*JbAddedPixels ==_aae &&_dbaf .Height -2*JbAddedPixels ==_bce {return _eef ;};};_aeee .Index ++;_aeee .CurrentNumbers =nil ;};};func (_fcb *Classer )verifyMethod (_bag Method )error {if _bag !=RankHaus &&_bag !=Correlation {return _af .Error ("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064","\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064");
};return nil ;};func (_ade Settings )Validate ()error {const _fbea ="\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065";if _ade .Thresh < 0.4||_ade .Thresh > 0.98{return _af .Error (_fbea ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d");
};if _ade .WeightFactor < 0.0||_ade .WeightFactor > 1.0{return _af .Error (_fbea ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _ade .RankHaus < 0.5||_ade .RankHaus > 1.0{return _af .Error (_fbea ,"\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _ade .SizeHaus < 1||_ade .SizeHaus > 10{return _af .Error (_fbea ,"\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]");
};switch _ade .Components {case _gf .ComponentConn ,_gf .ComponentCharacters ,_gf .ComponentWords :default:return _af .Error (_fbea ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et");
};return nil ;};func (_ag *Classer )classifyRankHouseOne (_bgb *_gf .Boxes ,_ada ,_cfad ,_fab *_gf .Bitmaps ,_ggbe *_gf .Points ,_cdbb int )(_gcf error ){const _cgf ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_feg ,_bee ,_def ,_abd float32 ;_dbbc int ;_gccb ,_fgd ,_bfe ,_acad ,_bgfa *_gf .Bitmap ;_bdaa ,_bba bool ;);for _faea :=0;_faea < len (_ada .Values );_faea ++{_fgd =_cfad .Values [_faea ];_bfe =_fab .Values [_faea ];_feg ,_bee ,_gcf =_ggbe .GetGeometry (_faea );
if _gcf !=nil {return _af .Wrapf (_gcf ,_cgf ,"\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079");};_fcd :=len (_ag .UndilatedTemplates .Values );_bdaa =false ;_ffc :=_fbb (_ag ,_fgd );for _dbbc =_ffc .Next ();_dbbc > -1;
{_acad ,_gcf =_ag .UndilatedTemplates .GetBitmap (_dbbc );if _gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"\u0062\u006d\u0033");};_bgfa ,_gcf =_ag .DilatedTemplates .GetBitmap (_dbbc );if _gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"\u0062\u006d\u0034");};_def ,_abd ,_gcf =_ag .CentroidPointsTemplates .GetGeometry (_dbbc );
if _gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073");};_bba ,_gcf =_gf .HausTest (_fgd ,_bfe ,_acad ,_bgfa ,_feg -_def ,_bee -_abd ,MaxDiffWidth ,MaxDiffHeight );
if _gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"");};if _bba {_bdaa =true ;if _gcf =_ag .ClassIDs .Add (_dbbc );_gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"");};if _gcf =_ag .ComponentPageNumbers .Add (_cdbb );_gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"");
};if _ag .Settings .KeepClassInstances {_cfg ,_afd :=_ag .ClassInstances .GetBitmaps (_dbbc );if _afd !=nil {return _af .Wrap (_afd ,_cgf ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_gccb ,_afd =_ada .GetBitmap (_faea );if _afd !=nil {return _af .Wrap (_afd ,_cgf ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");
};_cfg .AddBitmap (_gccb );_dgb ,_afd :=_bgb .Get (_faea );if _afd !=nil {return _af .Wrap (_afd ,_cgf ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_cfg .AddBox (_dgb );};break ;};};if !_bdaa {if _gcf =_ag .ClassIDs .Add (_fcd );_gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"");
};if _gcf =_ag .ComponentPageNumbers .Add (_cdbb );_gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"");};_eaa :=&_gf .Bitmaps {};_gccb ,_gcf =_ada .GetBitmap (_faea );if _gcf !=nil {return _af .Wrap (_gcf ,_cgf ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_eaa .Values =append (_eaa .Values ,_gccb );
_dbe ,_cdbg :=_gccb .Width ,_gccb .Height ;_ag .TemplatesSize .Add (uint64 (_cdbg )*uint64 (_dbe ),_fcd );_acb ,_fac :=_bgb .Get (_faea );if _fac !=nil {return _af .Wrap (_fac ,_cgf ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_eaa .AddBox (_acb );_ag .ClassInstances .AddBitmaps (_eaa );
_ag .CentroidPointsTemplates .AddPoint (_feg ,_bee );_ag .UndilatedTemplates .AddBitmap (_fgd );_ag .DilatedTemplates .AddBitmap (_bfe );};};return nil ;};func DefaultSettings ()Settings {_aga :=&Settings {};_aga .SetDefault ();return *_aga };func (_ce *Classer )classifyCorrelation (_fddb *_gf .Boxes ,_ccb *_gf .Bitmaps ,_dbc int )error {const _bde ="\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e";
if _fddb ==nil {return _af .Error (_bde ,"\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064");
};if _ccb ==nil {return _af .Error (_bde ,"\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};_ecc :=len (_ccb .Values );if _ecc ==0{_d .Log .Debug ("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079");
return nil ;};var (_gee ,_dcb *_gf .Bitmap ;_fbc error ;);_ad :=&_gf .Bitmaps {Values :make ([]*_gf .Bitmap ,_ecc )};for _eg ,_cca :=range _ccb .Values {_dcb ,_fbc =_cca .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _fbc !=nil {return _af .Wrap (_fbc ,_bde ,"");};_ad .Values [_eg ]=_dcb ;};_ceg :=_ce .FgTemplates ;_gbf :=_gf .MakePixelSumTab8 ();_bdec :=_gf .MakePixelCentroidTab8 ();_ffg :=make ([]int ,_ecc );_acc :=make ([][]int ,_ecc );_cgc :=_gf .Points (make ([]_gf .Point ,_ecc ));
_cfb :=&_cgc ;var (_ega ,_baf int ;_afc ,_acf ,_bbc int ;_gad ,_ace int ;_cda byte ;);for _cfa ,_cac :=range _ad .Values {_acc [_cfa ]=make ([]int ,_cac .Height );_ega =0;_baf =0;_acf =(_cac .Height -1)*_cac .RowStride ;_afc =0;for _ace =_cac .Height -1;
_ace >=0;_ace ,_acf =_ace -1,_acf -_cac .RowStride {_acc [_cfa ][_ace ]=_afc ;_bbc =0;for _gad =0;_gad < _cac .RowStride ;_gad ++{_cda =_cac .Data [_acf +_gad ];_bbc +=_gbf [_cda ];_ega +=_bdec [_cda ]+_gad *8*_gbf [_cda ];};_afc +=_bbc ;_baf +=_bbc *_ace ;
};_ffg [_cfa ]=_afc ;if _afc > 0{(*_cfb )[_cfa ]=_gf .Point {X :float32 (_ega )/float32 (_afc ),Y :float32 (_baf )/float32 (_afc )};}else {(*_cfb )[_cfa ]=_gf .Point {X :float32 (_cac .Width )/float32 (2),Y :float32 (_cac .Height )/float32 (2)};};};if _fbc =_ce .CentroidPoints .Add (_cfb );
_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064");};var (_ebe ,_bff ,_gbff int ;_cfaa float64 ;_egc ,_dbg ,_gaf ,_cge float32 ;_ffb ,_gae _gf .Point ;_fef bool ;_ggb *similarTemplatesFinder ;
_dce int ;_ebb *_gf .Bitmap ;_ggf *_b .Rectangle ;_dddg *_gf .Bitmaps ;);for _dce ,_dcb =range _ad .Values {_bff =_ffg [_dce ];if _egc ,_dbg ,_fbc =_cfb .GetGeometry (_dce );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0070t\u0061\u0020\u002d\u0020\u0069");
};_fef =false ;_afe :=len (_ce .UndilatedTemplates .Values );_ggb =_fbb (_ce ,_dcb );for _bgce :=_ggb .Next ();_bgce > -1;{if _ebb ,_fbc =_ce .UndilatedTemplates .GetBitmap (_bgce );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032");
};if _gbff ,_fbc =_ceg .GetInt (_bgce );_fbc !=nil {_d .Log .Trace ("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_fbc );};if _gaf ,_cge ,_fbc =_ce .CentroidPointsTemplates .GetGeometry (_bgce );
_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020");
};if _ce .Settings .WeightFactor > 0.0{if _ebe ,_fbc =_ce .TemplateAreas .Get (_bgce );_fbc !=nil {_d .Log .Trace ("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076",_fbc );
};_cfaa =_ce .Settings .Thresh +(1.0-_ce .Settings .Thresh )*_ce .Settings .WeightFactor *float64 (_gbff )/float64 (_ebe );}else {_cfaa =_ce .Settings .Thresh ;};_cdg ,_cb :=_gf .CorrelationScoreThresholded (_dcb ,_ebb ,_bff ,_gbff ,_ffb .X -_gae .X ,_ffb .Y -_gae .Y ,MaxDiffWidth ,MaxDiffHeight ,_gbf ,_acc [_dce ],float32 (_cfaa ));
if _cb !=nil {return _af .Wrap (_cb ,_bde ,"");};if _gcg {var (_bc ,_aee float64 ;_cdb ,_gbb int ;);_bc ,_cb =_gf .CorrelationScore (_dcb ,_ebb ,_bff ,_gbff ,_egc -_gaf ,_dbg -_cge ,MaxDiffWidth ,MaxDiffHeight ,_gbf );if _cb !=nil {return _af .Wrap (_cb ,_bde ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_aee ,_cb =_gf .CorrelationScoreSimple (_dcb ,_ebb ,_bff ,_gbff ,_egc -_gaf ,_dbg -_cge ,MaxDiffWidth ,MaxDiffHeight ,_gbf );if _cb !=nil {return _af .Wrap (_cb ,_bde ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_cdb =int (_a .Sqrt (_bc *float64 (_bff )*float64 (_gbff )));_gbb =int (_a .Sqrt (_aee *float64 (_bff )*float64 (_gbff )));if (_bc >=_cfaa )!=(_aee >=_cfaa ){return _af .Errorf (_bde ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066",_cdb ,_bc ,_bc >=_cfaa ,_gbb ,_aee ,_aee >=_cfaa ,_bc -_aee );
};if _bc >=_cfaa !=_cdg {return _af .Errorf (_bde ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076",_bc ,_bc *float64 (_bff )*float64 (_gbff ),_cdb ,_cfaa ,float32 (_cfaa )*float32 (_bff )*float32 (_gbff ),_cdg );
};};if _cdg {_fef =true ;if _cb =_ce .ClassIDs .Add (_bgce );_cb !=nil {return _af .Wrap (_cb ,_bde ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _cb =_ce .ComponentPageNumbers .Add (_dbc );_cb !=nil {return _af .Wrap (_cb ,_bde ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");
};if _ce .Settings .KeepClassInstances {if _gee ,_cb =_ccb .GetBitmap (_dce );_cb !=nil {return _af .Wrap (_cb ,_bde ,"\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069");};if _dddg ,_cb =_ce .ClassInstances .GetBitmaps (_bgce );
_cb !=nil {return _af .Wrap (_cb ,_bde ,"K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073");};_dddg .AddBitmap (_gee );if _ggf ,_cb =_fddb .Get (_dce );_cb !=nil {return _af .Wrap (_cb ,_bde ,"\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073");
};_dddg .AddBox (_ggf );};break ;};};if !_fef {if _fbc =_ce .ClassIDs .Add (_afe );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _fbc =_ce .ComponentPageNumbers .Add (_dbc );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_dddg =&_gf .Bitmaps {};if _gee ,_fbc =_ccb .GetBitmap (_dce );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_dddg .AddBitmap (_gee );_abb ,_ea :=_gee .Width ,_gee .Height ;_ggbf :=uint64 (_ea )*uint64 (_abb );_ce .TemplatesSize .Add (_ggbf ,_afe );
if _ggf ,_fbc =_fddb .Get (_dce );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_dddg .AddBox (_ggf );_ce .ClassInstances .AddBitmaps (_dddg );_ce .CentroidPointsTemplates .AddPoint (_egc ,_dbg );_ce .FgTemplates .AddInt (_bff );
_ce .UndilatedTemplates .AddBitmap (_gee );_ebe =(_dcb .Width -2*JbAddedPixels )*(_dcb .Height -2*JbAddedPixels );if _fbc =_ce .TemplateAreas .Add (_ebe );_fbc !=nil {return _af .Wrap (_fbc ,_bde ,"\u0021\u0066\u006f\u0075\u006e\u0064");};};};_ce .NumberOfClasses =len (_ce .UndilatedTemplates .Values );
return nil ;};type Classer struct{BaseIndex int ;Settings Settings ;ComponentsNumber *_c .IntSlice ;TemplateAreas *_c .IntSlice ;Widths map[int ]int ;Heights map[int ]int ;NumberOfClasses int ;ClassInstances *_gf .BitmapsArray ;UndilatedTemplates *_gf .Bitmaps ;
DilatedTemplates *_gf .Bitmaps ;TemplatesSize _c .IntsMap ;FgTemplates *_c .NumSlice ;CentroidPoints *_gf .Points ;CentroidPointsTemplates *_gf .Points ;ClassIDs *_c .IntSlice ;ComponentPageNumbers *_c .IntSlice ;PtaUL *_gf .Points ;PtaLL *_gf .Points ;
};var _gcg bool ;type Method int ;