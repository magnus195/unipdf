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

package classer ;import (_e "github.com/unidoc/unipdf/v3/common";_b "github.com/unidoc/unipdf/v3/internal/jbig2/basic";_gf "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_gg "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_gd "image";_c "math";
);func (_ba *Classer )AddPage (inputPage *_gf .Bitmap ,pageNumber int ,method Method )(_gb error ){const _a ="\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065";_ba .Widths [pageNumber ]=inputPage .Width ;_ba .Heights [pageNumber ]=inputPage .Height ;
if _gb =_ba .verifyMethod (method );_gb !=nil {return _gg .Wrap (_gb ,_a ,"");};_ad ,_aa ,_gb :=inputPage .GetComponents (_ba .Settings .Components ,_ba .Settings .MaxCompWidth ,_ba .Settings .MaxCompHeight );if _gb !=nil {return _gg .Wrap (_gb ,_a ,"");
};_e .Log .Debug ("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076",_ad );if _gb =_ba .addPageComponents (inputPage ,_aa ,_ad ,pageNumber ,method );_gb !=nil {return _gg .Wrap (_gb ,_a ,"");};return nil ;};func (_dag *Classer )getULCorners (_ea *_gf .Bitmap ,_fd *_gf .Boxes )error {const _bb ="\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";
if _ea ==nil {return _gg .Error (_bb ,"\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070");};if _fd ==nil {return _gg .Error (_bb ,"\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};if _dag .PtaUL ==nil {_dag .PtaUL =&_gf .Points {};
};_ef :=len (*_fd );var (_bba ,_fdb ,_gbd ,_gbg int ;_dd ,_cc ,_fc ,_ca float32 ;_bbe error ;_abf *_gd .Rectangle ;_cf *_gf .Bitmap ;_df _gd .Point ;);for _ff :=0;_ff < _ef ;_ff ++{_bba =_dag .BaseIndex +_ff ;if _dd ,_cc ,_bbe =_dag .CentroidPoints .GetGeometry (_bba );
_bbe !=nil {return _gg .Wrap (_bbe ,_bb ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073");};if _fdb ,_bbe =_dag .ClassIDs .Get (_bba );_bbe !=nil {return _gg .Wrap (_bbe ,_bb ,"\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074");
};if _fc ,_ca ,_bbe =_dag .CentroidPointsTemplates .GetGeometry (_fdb );_bbe !=nil {return _gg .Wrap (_bbe ,_bb ,"\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_ge :=_fc -_dd ;
_egd :=_ca -_cc ;if _ge >=0{_gbd =int (_ge +0.5);}else {_gbd =int (_ge -0.5);};if _egd >=0{_gbg =int (_egd +0.5);}else {_gbg =int (_egd -0.5);};if _abf ,_bbe =_fd .Get (_ff );_bbe !=nil {return _gg .Wrap (_bbe ,_bb ,"");};_daf ,_fe :=_abf .Min .X ,_abf .Min .Y ;
_cf ,_bbe =_dag .UndilatedTemplates .GetBitmap (_fdb );if _bbe !=nil {return _gg .Wrap (_bbe ,_bb ,"\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029");
};_df ,_bbe =_ggb (_ea ,_daf ,_fe ,_gbd ,_gbg ,_cf );if _bbe !=nil {return _gg .Wrap (_bbe ,_bb ,"");};_dag .PtaUL .AddPoint (float32 (_daf -_gbd +_df .X ),float32 (_fe -_gbg +_df .Y ));};return nil ;};type Method int ;const (MaxConnCompWidth =350;MaxCharCompWidth =350;
MaxWordCompWidth =1000;MaxCompHeight =120;);func Init (settings Settings )(*Classer ,error ){const _be ="\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074";_d :=&Classer {Settings :settings ,Widths :map[int ]int {},Heights :map[int ]int {},TemplatesSize :_b .IntsMap {},TemplateAreas :&_b .IntSlice {},ComponentPageNumbers :&_b .IntSlice {},ClassIDs :&_b .IntSlice {},ComponentsNumber :&_b .IntSlice {},CentroidPoints :&_gf .Points {},CentroidPointsTemplates :&_gf .Points {},UndilatedTemplates :&_gf .Bitmaps {},DilatedTemplates :&_gf .Bitmaps {},ClassInstances :&_gf .BitmapsArray {},FgTemplates :&_b .NumSlice {}};
if _dc :=_d .Settings .Validate ();_dc !=nil {return nil ,_gg .Wrap (_dc ,_be ,"");};return _d ,nil ;};func (_agf *Classer )classifyRankHaus (_gbdg *_gf .Boxes ,_ed *_gf .Bitmaps ,_fdg int )error {const _abg ="\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073";
if _gbdg ==nil {return _gg .Error (_abg ,"\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};if _ed ==nil {return _gg .Error (_abg ,"\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");
};_bdff :=len (_ed .Values );if _bdff ==0{return _gg .Error (_abg ,"e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073");};_ddd :=_ed .CountPixels ();_bgd :=_agf .Settings .SizeHaus ;_agc :=_gf .SelCreateBrick (_bgd ,_bgd ,_bgd /2,_bgd /2,_gf .SelHit );
_cab :=&_gf .Bitmaps {Values :make ([]*_gf .Bitmap ,_bdff )};_cbf :=&_gf .Bitmaps {Values :make ([]*_gf .Bitmap ,_bdff )};var (_gbdf ,_agb ,_bfg *_gf .Bitmap ;_fefb error ;);for _ebfg :=0;_ebfg < _bdff ;_ebfg ++{_gbdf ,_fefb =_ed .GetBitmap (_ebfg );if _fefb !=nil {return _gg .Wrap (_fefb ,_abg ,"");
};_agb ,_fefb =_gbdf .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);if _fefb !=nil {return _gg .Wrap (_fefb ,_abg ,"");};_bfg ,_fefb =_gf .Dilate (nil ,_agb ,_agc );if _fefb !=nil {return _gg .Wrap (_fefb ,_abg ,"");};
_cab .Values [_bdff ]=_agb ;_cbf .Values [_bdff ]=_bfg ;};_ebba ,_fefb :=_gf .Centroids (_cab .Values );if _fefb !=nil {return _gg .Wrap (_fefb ,_abg ,"");};if _fefb =_ebba .Add (_agf .CentroidPoints );_fefb !=nil {_e .Log .Trace ("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064");
};if _agf .Settings .RankHaus ==1.0{_fefb =_agf .classifyRankHouseOne (_gbdg ,_ed ,_cab ,_cbf ,_ebba ,_fdg );}else {_fefb =_agf .classifyRankHouseNonOne (_gbdg ,_ed ,_cab ,_cbf ,_ebba ,_ddd ,_fdg );};if _fefb !=nil {return _gg .Wrap (_fefb ,_abg ,"");};
return nil ;};func (_fb *Classer )verifyMethod (_gfd Method )error {if _gfd !=RankHaus &&_gfd !=Correlation {return _gg .Error ("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064","\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064");
};return nil ;};const (MaxDiffWidth =2;MaxDiffHeight =2;);func (_bac *Classer )ComputeLLCorners ()(_da error ){const _bef ="\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";
if _bac .PtaUL ==nil {return _gg .Error (_bef ,"\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064");};_bg :=len (*_bac .PtaUL );_bac .PtaLL =&_gf .Points {};var (_ec ,_ag float32 ;_ab ,_gff int ;
_eb *_gf .Bitmap ;);for _f :=0;_f < _bg ;_f ++{_ec ,_ag ,_da =_bac .PtaUL .GetGeometry (_f );if _da !=nil {_e .Log .Debug ("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_da );
return _gg .Wrap (_da ,_bef ,"\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079");};_ab ,_da =_bac .ClassIDs .Get (_f );if _da !=nil {_e .Log .Debug ("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076",_da );
return _gg .Wrap (_da ,_bef ,"\u0043l\u0061\u0073\u0073\u0049\u0044");};_eb ,_da =_bac .UndilatedTemplates .GetBitmap (_ab );if _da !=nil {_e .Log .Debug ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_da );
return _gg .Wrap (_da ,_bef ,"\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_gff =_eb .Height ;_bac .PtaLL .AddPoint (_ec ,_ag +float32 (_gff ));};return nil ;};type Settings struct{MaxCompWidth int ;
MaxCompHeight int ;SizeHaus int ;RankHaus float64 ;Thresh float64 ;WeightFactor float64 ;KeepClassInstances bool ;Components _gf .Component ;Method Method ;};func (_gcg *Classer )classifyCorrelation (_egfg *_gf .Boxes ,_eff *_gf .Bitmaps ,_ggf int )error {const _ddc ="\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e";
if _egfg ==nil {return _gg .Error (_ddc ,"\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064");
};if _eff ==nil {return _gg .Error (_ddc ,"\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};_bbb :=len (_eff .Values );if _bbb ==0{_e .Log .Debug ("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079");
return nil ;};var (_ffe ,_fdf *_gf .Bitmap ;_ebb error ;);_beb :=&_gf .Bitmaps {Values :make ([]*_gf .Bitmap ,_bbb )};for _ada ,_bde :=range _eff .Values {_fdf ,_ebb =_bde .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"");};_beb .Values [_ada ]=_fdf ;};_gcc :=_gcg .FgTemplates ;_gdce :=_gf .MakePixelSumTab8 ();_cbb :=_gf .MakePixelCentroidTab8 ();_bad :=make ([]int ,_bbb );_fgg :=make ([][]int ,_bbb );_fee :=_gf .Points (make ([]_gf .Point ,_bbb ));
_ebg :=&_fee ;var (_eabg ,_bda int ;_cac ,_feb ,_bdf int ;_fca ,_dg int ;_bgb byte ;);for _de ,_bbaa :=range _beb .Values {_fgg [_de ]=make ([]int ,_bbaa .Height );_eabg =0;_bda =0;_feb =(_bbaa .Height -1)*_bbaa .RowStride ;_cac =0;for _dg =_bbaa .Height -1;
_dg >=0;_dg ,_feb =_dg -1,_feb -_bbaa .RowStride {_fgg [_de ][_dg ]=_cac ;_bdf =0;for _fca =0;_fca < _bbaa .RowStride ;_fca ++{_bgb =_bbaa .Data [_feb +_fca ];_bdf +=_gdce [_bgb ];_eabg +=_cbb [_bgb ]+_fca *8*_gdce [_bgb ];};_cac +=_bdf ;_bda +=_bdf *_dg ;
};_bad [_de ]=_cac ;if _cac > 0{(*_ebg )[_de ]=_gf .Point {X :float32 (_eabg )/float32 (_cac ),Y :float32 (_bda )/float32 (_cac )};}else {(*_ebg )[_de ]=_gf .Point {X :float32 (_bbaa .Width )/float32 (2),Y :float32 (_bbaa .Height )/float32 (2)};};};if _ebb =_gcg .CentroidPoints .Add (_ebg );
_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064");};var (_aab ,_dbcc ,_cce int ;_fge float64 ;_af ,_gdg ,_fa ,_fgb float32 ;_fed ,_ade _gf .Point ;_bf bool ;_bbg *similarTemplatesFinder ;_ecf int ;
_fedd *_gf .Bitmap ;_afe *_gd .Rectangle ;_cdd *_gf .Bitmaps ;);for _ecf ,_fdf =range _beb .Values {_dbcc =_bad [_ecf ];if _af ,_gdg ,_ebb =_ebg .GetGeometry (_ecf );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0070t\u0061\u0020\u002d\u0020\u0069");};_bf =false ;
_afd :=len (_gcg .UndilatedTemplates .Values );_bbg =_ecba (_gcg ,_fdf );for _daga :=_bbg .Next ();_daga > -1;{if _fedd ,_ebb =_gcg .UndilatedTemplates .GetBitmap (_daga );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032");
};if _cce ,_ebb =_gcc .GetInt (_daga );_ebb !=nil {_e .Log .Trace ("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_ebb );};if _fa ,_fgb ,_ebb =_gcg .CentroidPointsTemplates .GetGeometry (_daga );
_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020");
};if _gcg .Settings .WeightFactor > 0.0{if _aab ,_ebb =_gcg .TemplateAreas .Get (_daga );_ebb !=nil {_e .Log .Trace ("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076",_ebb );
};_fge =_gcg .Settings .Thresh +(1.0-_gcg .Settings .Thresh )*_gcg .Settings .WeightFactor *float64 (_cce )/float64 (_aab );}else {_fge =_gcg .Settings .Thresh ;};_gbb ,_cacf :=_gf .CorrelationScoreThresholded (_fdf ,_fedd ,_dbcc ,_cce ,_fed .X -_ade .X ,_fed .Y -_ade .Y ,MaxDiffWidth ,MaxDiffHeight ,_gdce ,_fgg [_ecf ],float32 (_fge ));
if _cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"");};if _ebf {var (_cca ,_dbg float64 ;_bgbf ,_abb int ;);_cca ,_cacf =_gf .CorrelationScore (_fdf ,_fedd ,_dbcc ,_cce ,_af -_fa ,_gdg -_fgb ,MaxDiffWidth ,MaxDiffHeight ,_gdce );if _cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_dbg ,_cacf =_gf .CorrelationScoreSimple (_fdf ,_fedd ,_dbcc ,_cce ,_af -_fa ,_gdg -_fgb ,MaxDiffWidth ,MaxDiffHeight ,_gdce );if _cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_bgbf =int (_c .Sqrt (_cca *float64 (_dbcc )*float64 (_cce )));_abb =int (_c .Sqrt (_dbg *float64 (_dbcc )*float64 (_cce )));if (_cca >=_fge )!=(_dbg >=_fge ){return _gg .Errorf (_ddc ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066",_bgbf ,_cca ,_cca >=_fge ,_abb ,_dbg ,_dbg >=_fge ,_cca -_dbg );
};if _cca >=_fge !=_gbb {return _gg .Errorf (_ddc ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076",_cca ,_cca *float64 (_dbcc )*float64 (_cce ),_bgbf ,_fge ,float32 (_fge )*float32 (_dbcc )*float32 (_cce ),_gbb );
};};if _gbb {_bf =true ;if _cacf =_gcg .ClassIDs .Add (_daga );_cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _cacf =_gcg .ComponentPageNumbers .Add (_ggf );_cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");
};if _gcg .Settings .KeepClassInstances {if _ffe ,_cacf =_eff .GetBitmap (_ecf );_cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069");};if _cdd ,_cacf =_gcg .ClassInstances .GetBitmaps (_daga );
_cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073");};_cdd .AddBitmap (_ffe );if _afe ,_cacf =_egfg .Get (_ecf );_cacf !=nil {return _gg .Wrap (_cacf ,_ddc ,"\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073");
};_cdd .AddBox (_afe );};break ;};};if !_bf {if _ebb =_gcg .ClassIDs .Add (_afd );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _ebb =_gcg .ComponentPageNumbers .Add (_ggf );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_cdd =&_gf .Bitmaps {};if _ffe ,_ebb =_eff .GetBitmap (_ecf );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_cdd .AddBitmap (_ffe );_ebfe ,_ecb :=_ffe .Width ,_ffe .Height ;_cfa :=uint64 (_ecb )*uint64 (_ebfe );_gcg .TemplatesSize .Add (_cfa ,_afd );
if _afe ,_ebb =_egfg .Get (_ecf );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_cdd .AddBox (_afe );_gcg .ClassInstances .AddBitmaps (_cdd );_gcg .CentroidPointsTemplates .AddPoint (_af ,_gdg );_gcg .FgTemplates .AddInt (_dbcc );
_gcg .UndilatedTemplates .AddBitmap (_ffe );_aab =(_fdf .Width -2*JbAddedPixels )*(_fdf .Height -2*JbAddedPixels );if _ebb =_gcg .TemplateAreas .Add (_aab );_ebb !=nil {return _gg .Wrap (_ebb ,_ddc ,"\u0021\u0066\u006f\u0075\u006e\u0064");};};};_gcg .NumberOfClasses =len (_gcg .UndilatedTemplates .Values );
return nil ;};func (_eed *Classer )classifyRankHouseOne (_ccc *_gf .Boxes ,_fad ,_efa ,_bgbb *_gf .Bitmaps ,_cbbd *_gf .Points ,_dab int )(_ggd error ){const _aba ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_efd ,_fcg ,_eeda ,_fdc float32 ;_bag int ;_ac ,_ddg ,_dcd ,_ga ,_ddcc *_gf .Bitmap ;_eaf ,_gdcd bool ;);for _bce :=0;_bce < len (_fad .Values );_bce ++{_ddg =_efa .Values [_bce ];_dcd =_bgbb .Values [_bce ];_efd ,_fcg ,_ggd =_cbbd .GetGeometry (_bce );
if _ggd !=nil {return _gg .Wrapf (_ggd ,_aba ,"\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079");};_aff :=len (_eed .UndilatedTemplates .Values );_eaf =false ;_ggda :=_ecba (_eed ,_ddg );for _bag =_ggda .Next ();_bag > -1;
{_ga ,_ggd =_eed .UndilatedTemplates .GetBitmap (_bag );if _ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"\u0062\u006d\u0033");};_ddcc ,_ggd =_eed .DilatedTemplates .GetBitmap (_bag );if _ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"\u0062\u006d\u0034");};_eeda ,_fdc ,_ggd =_eed .CentroidPointsTemplates .GetGeometry (_bag );
if _ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073");};_gdcd ,_ggd =_gf .HausTest (_ddg ,_dcd ,_ga ,_ddcc ,_efd -_eeda ,_fcg -_fdc ,MaxDiffWidth ,MaxDiffHeight );
if _ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"");};if _gdcd {_eaf =true ;if _ggd =_eed .ClassIDs .Add (_bag );_ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"");};if _ggd =_eed .ComponentPageNumbers .Add (_dab );_ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"");
};if _eed .Settings .KeepClassInstances {_gae ,_ccf :=_eed .ClassInstances .GetBitmaps (_bag );if _ccf !=nil {return _gg .Wrap (_ccf ,_aba ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_ac ,_ccf =_fad .GetBitmap (_bce );if _ccf !=nil {return _gg .Wrap (_ccf ,_aba ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");
};_gae .AddBitmap (_ac );_aeg ,_ccf :=_ccc .Get (_bce );if _ccf !=nil {return _gg .Wrap (_ccf ,_aba ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_gae .AddBox (_aeg );};break ;};};if !_eaf {if _ggd =_eed .ClassIDs .Add (_aff );_ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"");
};if _ggd =_eed .ComponentPageNumbers .Add (_dab );_ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"");};_cgb :=&_gf .Bitmaps {};_ac ,_ggd =_fad .GetBitmap (_bce );if _ggd !=nil {return _gg .Wrap (_ggd ,_aba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_cgb .Values =append (_cgb .Values ,_ac );
_cgg ,_fgbb :=_ac .Width ,_ac .Height ;_eed .TemplatesSize .Add (uint64 (_fgbb )*uint64 (_cgg ),_aff );_fba ,_bcc :=_ccc .Get (_bce );if _bcc !=nil {return _gg .Wrap (_bcc ,_aba ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_cgb .AddBox (_fba );_eed .ClassInstances .AddBitmaps (_cgb );
_eed .CentroidPointsTemplates .AddPoint (_efd ,_fcg );_eed .UndilatedTemplates .AddBitmap (_ddg );_eed .DilatedTemplates .AddBitmap (_dcd );};};return nil ;};type Classer struct{BaseIndex int ;Settings Settings ;ComponentsNumber *_b .IntSlice ;TemplateAreas *_b .IntSlice ;
Widths map[int ]int ;Heights map[int ]int ;NumberOfClasses int ;ClassInstances *_gf .BitmapsArray ;UndilatedTemplates *_gf .Bitmaps ;DilatedTemplates *_gf .Bitmaps ;TemplatesSize _b .IntsMap ;FgTemplates *_b .NumSlice ;CentroidPoints *_gf .Points ;CentroidPointsTemplates *_gf .Points ;
ClassIDs *_b .IntSlice ;ComponentPageNumbers *_b .IntSlice ;PtaUL *_gf .Points ;PtaLL *_gf .Points ;};func (_db *Classer )addPageComponents (_cb *_gf .Bitmap ,_ebd *_gf .Boxes ,_cd *_gf .Bitmaps ,_gc int ,_eg Method )error {const _dbd ="\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s";
if _cb ==nil {return _gg .Error (_dbd ,"\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065");};if _ebd ==nil ||_cd ==nil ||len (*_ebd )==0{_e .Log .Trace ("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064",_cb );
return nil ;};var _aae error ;switch _eg {case RankHaus :_aae =_db .classifyRankHaus (_ebd ,_cd ,_gc );case Correlation :_aae =_db .classifyCorrelation (_ebd ,_cd ,_gc );default:_e .Log .Debug ("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027",_eg );
return _gg .Error (_dbd ,"\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064");};if _aae !=nil {return _gg .Wrap (_aae ,_dbd ,"");};if _aae =_db .getULCorners (_cb ,_ebd );_aae !=nil {return _gg .Wrap (_aae ,_dbd ,"");
};_cg :=len (*_ebd );_db .BaseIndex +=_cg ;if _aae =_db .ComponentsNumber .Add (_cg );_aae !=nil {return _gg .Wrap (_aae ,_dbd ,"");};return nil ;};var _ebf bool ;func _ecba (_agbd *Classer ,_dabe *_gf .Bitmap )*similarTemplatesFinder {return &similarTemplatesFinder {Width :_dabe .Width ,Height :_dabe .Height ,Classer :_agbd };
};func DefaultSettings ()Settings {_bade :=&Settings {};_bade .SetDefault ();return *_bade };type similarTemplatesFinder struct{Classer *Classer ;Width int ;Height int ;Index int ;CurrentNumbers []int ;N int ;};func (_agaa *Settings )SetDefault (){if _agaa .MaxCompWidth ==0{switch _agaa .Components {case _gf .ComponentConn :_agaa .MaxCompWidth =MaxConnCompWidth ;
case _gf .ComponentCharacters :_agaa .MaxCompWidth =MaxCharCompWidth ;case _gf .ComponentWords :_agaa .MaxCompWidth =MaxWordCompWidth ;};};if _agaa .MaxCompHeight ==0{_agaa .MaxCompHeight =MaxCompHeight ;};if _agaa .Thresh ==0.0{_agaa .Thresh =0.9;};if _agaa .WeightFactor ==0.0{_agaa .WeightFactor =0.75;
};if _agaa .RankHaus ==0.0{_agaa .RankHaus =0.97;};if _agaa .SizeHaus ==0{_agaa .SizeHaus =2;};};func (_fgeg Settings )Validate ()error {const _fadc ="\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065";if _fgeg .Thresh < 0.4||_fgeg .Thresh > 0.98{return _gg .Error (_fadc ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d");
};if _fgeg .WeightFactor < 0.0||_fgeg .WeightFactor > 1.0{return _gg .Error (_fadc ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _fgeg .RankHaus < 0.5||_fgeg .RankHaus > 1.0{return _gg .Error (_fadc ,"\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _fgeg .SizeHaus < 1||_fgeg .SizeHaus > 10{return _gg .Error (_fadc ,"\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]");
};switch _fgeg .Components {case _gf .ComponentConn ,_gf .ComponentCharacters ,_gf .ComponentWords :default:return _gg .Error (_fadc ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et");
};return nil ;};const JbAddedPixels =6;var TwoByTwoWalk =[]int {0,0,0,1,-1,0,0,-1,1,0,-1,1,1,1,-1,-1,1,-1,0,-2,2,0,0,2,-2,0,-1,-2,1,-2,2,-1,2,1,1,2,-1,2,-2,1,-2,-1,-2,-2,2,-2,2,2,-2,2};func _ggb (_efc *_gf .Bitmap ,_gfdf ,_egf ,_eab ,_ee int ,_gdd *_gf .Bitmap )(_ggc _gd .Point ,_cfd error ){const _gce ="\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg";
if _efc ==nil {return _ggc ,_gg .Error (_gce ,"\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};if _gdd ==nil {return _ggc ,_gg .Error (_gce ,"t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064");
};_dce ,_eac :=_gdd .Width ,_gdd .Height ;_ffg ,_bbd :=_gfdf -_eab -JbAddedPixels ,_egf -_ee -JbAddedPixels ;_e .Log .Trace ("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027",_gfdf ,_egf ,_dce ,_eac ,_ffg ,_bbd );
_cgd ,_cfd :=_gf .Rect (_ffg ,_bbd ,_dce ,_eac );if _cfd !=nil {return _ggc ,_gg .Wrap (_cfd ,_gce ,"");};_ae ,_ ,_cfd :=_efc .ClipRectangle (_cgd );if _cfd !=nil {_e .Log .Error ("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076",_cgd );
return _ggc ,_gg .Wrap (_cfd ,_gce ,"");};_dbc :=_gf .New (_ae .Width ,_ae .Height );_bc :=_c .MaxInt32 ;var _aga ,_fg ,_cda ,_bd ,_dfc int ;for _aga =-1;_aga <=1;_aga ++{for _fg =-1;_fg <=1;_fg ++{if _ ,_cfd =_gf .Copy (_dbc ,_ae );_cfd !=nil {return _ggc ,_gg .Wrap (_cfd ,_gce ,"");
};if _cfd =_dbc .RasterOperation (_fg ,_aga ,_dce ,_eac ,_gf .PixSrcXorDst ,_gdd ,0,0);_cfd !=nil {return _ggc ,_gg .Wrap (_cfd ,_gce ,"");};_cda =_dbc .CountPixels ();if _cda < _bc {_bd =_fg ;_dfc =_aga ;_bc =_cda ;};};};_ggc .X =_bd ;_ggc .Y =_dfc ;return _ggc ,nil ;
};func (_eea *Classer )classifyRankHouseNonOne (_bcb *_gf .Boxes ,_bebb ,_dea ,_gdb *_gf .Bitmaps ,_bab *_gf .Points ,_ce *_b .NumSlice ,_aaa int )(_fac error ){const _dca ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_gdf ,_fda ,_dgg ,_cec float32 ;_bed ,_dafd ,_gdfd int ;_bcec ,_cdf ,_baa ,_bfgc ,_bgf *_gf .Bitmap ;_cgbe ,_ccd bool ;);_caa :=_gf .MakePixelSumTab8 ();for _bfgg :=0;_bfgg < len (_bebb .Values );_bfgg ++{if _cdf ,_fac =_dea .GetBitmap (_bfgg );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};if _bed ,_fac =_ce .GetInt (_bfgg );_fac !=nil {_e .Log .Trace ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_bfgg ,_fac );
};if _baa ,_fac =_gdb .GetBitmap (_bfgg );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _gdf ,_fda ,_fac =_bab .GetGeometry (_bfgg );_fac !=nil {return _gg .Wrapf (_fac ,_dca ,"\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079");
};_dae :=len (_eea .UndilatedTemplates .Values );_cgbe =false ;_edg :=_ecba (_eea ,_cdf );for _gdfd =_edg .Next ();_gdfd > -1;{if _bfgc ,_fac =_eea .UndilatedTemplates .GetBitmap (_gdfd );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");
};if _dafd ,_fac =_eea .FgTemplates .GetInt (_gdfd );_fac !=nil {_e .Log .Trace ("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_gdfd ,_fac );
};if _bgf ,_fac =_eea .DilatedTemplates .GetBitmap (_gdfd );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _dgg ,_cec ,_fac =_eea .CentroidPointsTemplates .GetGeometry (_gdfd );
_fac !=nil {return _gg .Wrap (_fac ,_dca ,"\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d");};_ccd ,_fac =_gf .RankHausTest (_cdf ,_baa ,_bfgc ,_bgf ,_gdf -_dgg ,_fda -_cec ,MaxDiffWidth ,MaxDiffHeight ,_bed ,_dafd ,float32 (_eea .Settings .RankHaus ),_caa );
if _fac !=nil {return _gg .Wrap (_fac ,_dca ,"");};if _ccd {_cgbe =true ;if _fac =_eea .ClassIDs .Add (_gdfd );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"");};if _fac =_eea .ComponentPageNumbers .Add (_aaa );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"");
};if _eea .Settings .KeepClassInstances {_bea ,_daa :=_eea .ClassInstances .GetBitmaps (_gdfd );if _daa !=nil {return _gg .Wrap (_daa ,_dca ,"\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029");
};if _bcec ,_daa =_bebb .GetBitmap (_bfgg );_daa !=nil {return _gg .Wrap (_daa ,_dca ,"\u0070i\u0078\u0061\u005b\u0069\u005d");};_bea .Values =append (_bea .Values ,_bcec );_fag ,_daa :=_bcb .Get (_bfgg );if _daa !=nil {return _gg .Wrap (_daa ,_dca ,"b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};_bea .Boxes =append (_bea .Boxes ,_fag );};break ;};};if !_cgbe {if _fac =_eea .ClassIDs .Add (_dae );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _fac =_eea .ComponentPageNumbers .Add (_aaa );_fac !=nil {return _gg .Wrap (_fac ,_dca ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_efae :=&_gf .Bitmaps {};_bcec =_bebb .Values [_bfgg ];_efae .AddBitmap (_bcec );_aef ,_ebc :=_bcec .Width ,_bcec .Height ;_eea .TemplatesSize .Add (uint64 (_aef )*uint64 (_ebc ),_dae );_bebd ,_cfg :=_bcb .Get (_bfgg );if _cfg !=nil {return _gg .Wrap (_cfg ,_dca ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_efae .AddBox (_bebd );_eea .ClassInstances .AddBitmaps (_efae );_eea .CentroidPointsTemplates .AddPoint (_gdf ,_fda );_eea .UndilatedTemplates .AddBitmap (_cdf );_eea .DilatedTemplates .AddBitmap (_baa );_eea .FgTemplates .AddInt (_bed );};};_eea .NumberOfClasses =len (_eea .UndilatedTemplates .Values );
return nil ;};const (RankHaus Method =iota ;Correlation ;);func (_gdgg *similarTemplatesFinder )Next ()int {var (_feddf ,_fdd ,_aad ,_eag int ;_dcf bool ;_feef *_gf .Bitmap ;_cea error ;);for {if _gdgg .Index >=25{return -1;};_fdd =_gdgg .Width +TwoByTwoWalk [2*_gdgg .Index ];
_feddf =_gdgg .Height +TwoByTwoWalk [2*_gdgg .Index +1];if _feddf < 1||_fdd < 1{_gdgg .Index ++;continue ;};if len (_gdgg .CurrentNumbers )==0{_gdgg .CurrentNumbers ,_dcf =_gdgg .Classer .TemplatesSize .GetSlice (uint64 (_fdd )*uint64 (_feddf ));if !_dcf {_gdgg .Index ++;
continue ;};_gdgg .N =0;};_aad =len (_gdgg .CurrentNumbers );for ;_gdgg .N < _aad ;_gdgg .N ++{_eag =_gdgg .CurrentNumbers [_gdgg .N ];_feef ,_cea =_gdgg .Classer .DilatedTemplates .GetBitmap (_eag );if _cea !=nil {_e .Log .Debug ("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020");
return 0;};if _feef .Width -2*JbAddedPixels ==_fdd &&_feef .Height -2*JbAddedPixels ==_feddf {return _eag ;};};_gdgg .Index ++;_gdgg .CurrentNumbers =nil ;};};