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

package classer ;import (_ba "github.com/unidoc/unipdf/v3/common";_c "github.com/unidoc/unipdf/v3/internal/jbig2/basic";_ce "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_e "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_g "image";_a "math";
);type Method int ;func (_fg *Classer )addPageComponents (_fe *_ce .Bitmap ,_d *_ce .Boxes ,_dc *_ce .Bitmaps ,_fde int ,_bd Method )error {const _gge ="\u0043l\u0061\u0073\u0073\u0065r\u002e\u0041\u0064\u0064\u0050a\u0067e\u0043o\u006d\u0070\u006f\u006e\u0065\u006e\u0074s";
if _fe ==nil {return _e .Error (_gge ,"\u006e\u0069\u006c\u0020\u0069\u006e\u0070\u0075\u0074 \u0070\u0061\u0067\u0065");};if _d ==nil ||_dc ==nil ||len (*_d )==0{_ba .Log .Trace ("\u0041\u0064\u0064P\u0061\u0067\u0065\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u003a\u0020\u0025\u0073\u002e\u0020\u004e\u006f\u0020\u0063\u006f\u006d\u0070\u006f\u006e\u0065n\u0074\u0073\u0020\u0066\u006f\u0075\u006e\u0064",_fe );
return nil ;};var _bef error ;switch _bd {case RankHaus :_bef =_fg .classifyRankHaus (_d ,_dc ,_fde );case Correlation :_bef =_fg .classifyCorrelation (_d ,_dc ,_fde );default:_ba .Log .Debug ("\u0055\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064\u003a\u0020'%\u0076\u0027",_bd );
return _e .Error (_gge ,"\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006c\u0061\u0073\u0073\u0069\u0066\u0079\u0020\u006d\u0065\u0074\u0068\u006f\u0064");};if _bef !=nil {return _e .Wrap (_bef ,_gge ,"");};if _bef =_fg .getULCorners (_fe ,_d );_bef !=nil {return _e .Wrap (_bef ,_gge ,"");
};_bg :=len (*_d );_fg .BaseIndex +=_bg ;if _bef =_fg .ComponentsNumber .Add (_bg );_bef !=nil {return _e .Wrap (_bef ,_gge ,"");};return nil ;};const (RankHaus Method =iota ;Correlation ;);const (MaxDiffWidth =2;MaxDiffHeight =2;);func (_ed *Classer )getULCorners (_cg *_ce .Bitmap ,_bb *_ce .Boxes )error {const _fdb ="\u0067\u0065\u0074U\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";
if _cg ==nil {return _e .Error (_fdb ,"\u006e\u0069l\u0020\u0069\u006da\u0067\u0065\u0020\u0062\u0069\u0074\u006d\u0061\u0070");};if _bb ==nil {return _e .Error (_fdb ,"\u006e\u0069\u006c\u0020\u0062\u006f\u0075\u006e\u0064\u0073");};if _ed .PtaUL ==nil {_ed .PtaUL =&_ce .Points {};
};_ca :=len (*_bb );var (_fgf ,_cc ,_cbe ,_aac int ;_cgg ,_eg ,_fb ,_cfa float32 ;_ffe error ;_ab *_g .Rectangle ;_fge *_ce .Bitmap ;_dcc _g .Point ;);for _gf :=0;_gf < _ca ;_gf ++{_fgf =_ed .BaseIndex +_gf ;if _cgg ,_eg ,_ffe =_ed .CentroidPoints .GetGeometry (_fgf );
_ffe !=nil {return _e .Wrap (_ffe ,_fdb ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050o\u0069\u006e\u0074\u0073");};if _cc ,_ffe =_ed .ClassIDs .Get (_fgf );_ffe !=nil {return _e .Wrap (_ffe ,_fdb ,"\u0043\u006c\u0061s\u0073\u0049\u0044\u0073\u002e\u0047\u0065\u0074");
};if _fb ,_cfa ,_ffe =_ed .CentroidPointsTemplates .GetGeometry (_cc );_ffe !=nil {return _e .Wrap (_ffe ,_fdb ,"\u0043\u0065\u006etr\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074\u0073\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_ad :=_fb -_cgg ;
_dcg :=_cfa -_eg ;if _ad >=0{_cbe =int (_ad +0.5);}else {_cbe =int (_ad -0.5);};if _dcg >=0{_aac =int (_dcg +0.5);}else {_aac =int (_dcg -0.5);};if _ab ,_ffe =_bb .Get (_gf );_ffe !=nil {return _e .Wrap (_ffe ,_fdb ,"");};_fbc ,_caf :=_ab .Min .X ,_ab .Min .Y ;
_fge ,_ffe =_ed .UndilatedTemplates .GetBitmap (_cc );if _ffe !=nil {return _e .Wrap (_ffe ,_fdb ,"\u0055\u006e\u0064\u0069\u006c\u0061\u0074\u0065\u0064\u0054e\u006d\u0070\u006c\u0061\u0074\u0065\u0073.\u0047\u0065\u0074\u0028\u0069\u0043\u006c\u0061\u0073\u0073\u0029");
};_dcc ,_ffe =_ae (_cg ,_fbc ,_caf ,_cbe ,_aac ,_fge );if _ffe !=nil {return _e .Wrap (_ffe ,_fdb ,"");};_ed .PtaUL .AddPoint (float32 (_fbc -_cbe +_dcc .X ),float32 (_caf -_aac +_dcc .Y ));};return nil ;};type similarTemplatesFinder struct{Classer *Classer ;
Width int ;Height int ;Index int ;CurrentNumbers []int ;N int ;};var _ada bool ;var TwoByTwoWalk =[]int {0,0,0,1,-1,0,0,-1,1,0,-1,1,1,1,-1,-1,1,-1,0,-2,2,0,0,2,-2,0,-1,-2,1,-2,2,-1,2,1,1,2,-1,2,-2,1,-2,-1,-2,-2,2,-2,2,2,-2,2};func (_adf Settings )Validate ()error {const _bfg ="\u0053\u0065\u0074\u0074\u0069\u006e\u0067\u0073\u002e\u0056\u0061\u006ci\u0064\u0061\u0074\u0065";
if _adf .Thresh < 0.4||_adf .Thresh > 0.98{return _e .Error (_bfg ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0074\u0068\u0072\u0065\u0073\u0068\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0034\u0020\u002d\u0020\u0030\u002e\u0039\u0038\u005d");
};if _adf .WeightFactor < 0.0||_adf .WeightFactor > 1.0{return _e .Error (_bfg ,"\u006a\u0062i\u0067\u0032\u0020\u0065\u006ec\u006f\u0064\u0065\u0072\u0020w\u0065\u0069\u0067\u0068\u0074\u0020\u0066\u0061\u0063\u0074\u006f\u0072\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0030\u002e\u0030\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _adf .RankHaus < 0.5||_adf .RankHaus > 1.0{return _e .Error (_bfg ,"\u006a\u0062\u0069\u0067\u0032\u0020\u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0072a\u006e\u006b\u0020\u0068\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065 \u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065 [\u0030\u002e\u0035\u0020\u002d\u0020\u0031\u002e\u0030\u005d");
};if _adf .SizeHaus < 1||_adf .SizeHaus > 10{return _e .Error (_bfg ,"\u006a\u0062\u0069\u0067\u0032 \u0065\u006e\u0063\u006f\u0064\u0065\u0072\u0020\u0073\u0069\u007a\u0065\u0020h\u0061\u0075\u0073\u0020\u0076\u0061\u006c\u0075\u0065\u0020\u006e\u006f\u0074\u0020\u0069\u006e\u0020\u0072\u0061\u006e\u0067\u0065\u0020\u005b\u0031\u0020\u002d\u0020\u0031\u0030]");
};switch _adf .Components {case _ce .ComponentConn ,_ce .ComponentCharacters ,_ce .ComponentWords :default:return _e .Error (_bfg ,"\u0069n\u0076\u0061\u006c\u0069d\u0020\u0063\u006c\u0061\u0073s\u0065r\u0020c\u006f\u006d\u0070\u006f\u006e\u0065\u006et");
};return nil ;};func Init (settings Settings )(*Classer ,error ){const _f ="\u0063\u006c\u0061s\u0073\u0065\u0072\u002e\u0049\u006e\u0069\u0074";_bac :=&Classer {Settings :settings ,Widths :map[int ]int {},Heights :map[int ]int {},TemplatesSize :_c .IntsMap {},TemplateAreas :&_c .IntSlice {},ComponentPageNumbers :&_c .IntSlice {},ClassIDs :&_c .IntSlice {},ComponentsNumber :&_c .IntSlice {},CentroidPoints :&_ce .Points {},CentroidPointsTemplates :&_ce .Points {},UndilatedTemplates :&_ce .Bitmaps {},DilatedTemplates :&_ce .Bitmaps {},ClassInstances :&_ce .BitmapsArray {},FgTemplates :&_c .NumSlice {}};
if _be :=_bac .Settings .Validate ();_be !=nil {return nil ,_e .Wrap (_be ,_f ,"");};return _bac ,nil ;};func _ae (_dg *_ce .Bitmap ,_aab ,_bf ,_ac ,_edc int ,_agb *_ce .Bitmap )(_ffec _g .Point ,_eag error ){const _cd ="\u0066i\u006e\u0061\u006c\u0041l\u0069\u0067\u006e\u006d\u0065n\u0074P\u006fs\u0069\u0074\u0069\u006f\u006e\u0069\u006eg";
if _dg ==nil {return _ffec ,_e .Error (_cd ,"\u0073\u006f\u0075\u0072ce\u0020\u006e\u006f\u0074\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};if _agb ==nil {return _ffec ,_e .Error (_cd ,"t\u0065\u006d\u0070\u006cat\u0065 \u006e\u006f\u0074\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064");
};_ccf ,_aae :=_agb .Width ,_agb .Height ;_beff ,_aff :=_aab -_ac -JbAddedPixels ,_bf -_edc -JbAddedPixels ;_ba .Log .Trace ("\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0079\u003a\u0020\u0027\u0025\u0064'\u002c\u0020\u0077\u003a\u0020\u0027\u0025\u0064\u0027\u002c\u0020\u0068\u003a \u0027\u0025\u0064\u0027\u002c\u0020\u0062\u0078\u003a\u0020\u0027\u0025d'\u002c\u0020\u0062\u0079\u003a\u0020\u0027\u0025\u0064\u0027",_aab ,_bf ,_ccf ,_aae ,_beff ,_aff );
_aacb ,_eag :=_ce .Rect (_beff ,_aff ,_ccf ,_aae );if _eag !=nil {return _ffec ,_e .Wrap (_eag ,_cd ,"");};_edb ,_ ,_eag :=_dg .ClipRectangle (_aacb );if _eag !=nil {_ba .Log .Error ("\u0043a\u006e\u0027\u0074\u0020\u0063\u006c\u0069\u0070\u0020\u0072\u0065c\u0074\u0061\u006e\u0067\u006c\u0065\u003a\u0020\u0025\u0076",_aacb );
return _ffec ,_e .Wrap (_eag ,_cd ,"");};_ggb :=_ce .New (_edb .Width ,_edb .Height );_aaf :=_a .MaxInt32 ;var _cac ,_ggec ,_cca ,_beb ,_ga int ;for _cac =-1;_cac <=1;_cac ++{for _ggec =-1;_ggec <=1;_ggec ++{if _ ,_eag =_ce .Copy (_ggb ,_edb );_eag !=nil {return _ffec ,_e .Wrap (_eag ,_cd ,"");
};if _eag =_ggb .RasterOperation (_ggec ,_cac ,_ccf ,_aae ,_ce .PixSrcXorDst ,_agb ,0,0);_eag !=nil {return _ffec ,_e .Wrap (_eag ,_cd ,"");};_cca =_ggb .CountPixels ();if _cca < _aaf {_beb =_ggec ;_ga =_cac ;_aaf =_cca ;};};};_ffec .X =_beb ;_ffec .Y =_ga ;
return _ffec ,nil ;};func (_ffed *Classer )classifyRankHaus (_ggc *_ce .Boxes ,_cebc *_ce .Bitmaps ,_dca int )error {const _fff ="\u0063\u006ca\u0073\u0073\u0069f\u0079\u0052\u0061\u006e\u006b\u0048\u0061\u0075\u0073";if _ggc ==nil {return _e .Error (_fff ,"\u0062\u006fx\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");
};if _cebc ==nil {return _e .Error (_fff ,"\u0070\u0069x\u0061\u0020\u006eo\u0074\u0020\u0064\u0065\u0066\u0069\u006e\u0065\u0064");};_cfc :=len (_cebc .Values );if _cfc ==0{return _e .Error (_fff ,"e\u006dp\u0074\u0079\u0020\u006e\u0065\u0077\u0020\u0063o\u006d\u0070\u006f\u006een\u0074\u0073");
};_ccae :=_cebc .CountPixels ();_age :=_ffed .Settings .SizeHaus ;_cae :=_ce .SelCreateBrick (_age ,_age ,_age /2,_age /2,_ce .SelHit );_bad :=&_ce .Bitmaps {Values :make ([]*_ce .Bitmap ,_cfc )};_bebf :=&_ce .Bitmaps {Values :make ([]*_ce .Bitmap ,_cfc )};
var (_bdfa ,_ade ,_bdcb *_ce .Bitmap ;_bde error ;);for _acc :=0;_acc < _cfc ;_acc ++{_bdfa ,_bde =_cebc .GetBitmap (_acc );if _bde !=nil {return _e .Wrap (_bde ,_fff ,"");};_ade ,_bde =_bdfa .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _bde !=nil {return _e .Wrap (_bde ,_fff ,"");};_bdcb ,_bde =_ce .Dilate (nil ,_ade ,_cae );if _bde !=nil {return _e .Wrap (_bde ,_fff ,"");};_bad .Values [_cfc ]=_ade ;_bebf .Values [_cfc ]=_bdcb ;};_faf ,_bde :=_ce .Centroids (_bad .Values );if _bde !=nil {return _e .Wrap (_bde ,_fff ,"");
};if _bde =_faf .Add (_ffed .CentroidPoints );_bde !=nil {_ba .Log .Trace ("\u004e\u006f\u0020\u0063en\u0074\u0072\u006f\u0069\u0064\u0073\u0020\u0074\u006f\u0020\u0061\u0064\u0064");};if _ffed .Settings .RankHaus ==1.0{_bde =_ffed .classifyRankHouseOne (_ggc ,_cebc ,_bad ,_bebf ,_faf ,_dca );
}else {_bde =_ffed .classifyRankHouseNonOne (_ggc ,_cebc ,_bad ,_bebf ,_faf ,_ccae ,_dca );};if _bde !=nil {return _e .Wrap (_bde ,_fff ,"");};return nil ;};func (_ea *Classer )AddPage (inputPage *_ce .Bitmap ,pageNumber int ,method Method )(_gg error ){const _ff ="\u0043l\u0061s\u0073\u0065\u0072\u002e\u0041\u0064\u0064\u0050\u0061\u0067\u0065";
_ea .Widths [pageNumber ]=inputPage .Width ;_ea .Heights [pageNumber ]=inputPage .Height ;if _gg =_ea .verifyMethod (method );_gg !=nil {return _e .Wrap (_gg ,_ff ,"");};_cf ,_ag ,_gg :=inputPage .GetComponents (_ea .Settings .Components ,_ea .Settings .MaxCompWidth ,_ea .Settings .MaxCompHeight );
if _gg !=nil {return _e .Wrap (_gg ,_ff ,"");};_ba .Log .Debug ("\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074s\u003a\u0020\u0025\u0076",_cf );if _gg =_ea .addPageComponents (inputPage ,_ag ,_cf ,pageNumber ,method );_gg !=nil {return _e .Wrap (_gg ,_ff ,"");
};return nil ;};func (_dd *Classer )verifyMethod (_beeb Method )error {if _beeb !=RankHaus &&_beeb !=Correlation {return _e .Error ("\u0076\u0065\u0072i\u0066\u0079\u004d\u0065\u0074\u0068\u006f\u0064","\u0069\u006e\u0076\u0061li\u0064\u0020\u0063\u006c\u0061\u0073\u0073\u0065\u0072\u0020\u006d\u0065\u0074\u0068o\u0064");
};return nil ;};func (_cb *Classer )ComputeLLCorners ()(_cfd error ){const _fa ="\u0043l\u0061\u0073\u0073\u0065\u0072\u002e\u0043\u006f\u006d\u0070\u0075t\u0065\u004c\u004c\u0043\u006f\u0072\u006e\u0065\u0072\u0073";if _cb .PtaUL ==nil {return _e .Error (_fa ,"\u0055\u004c\u0020\u0043or\u006e\u0065\u0072\u0073\u0020\u006e\u006f\u0074\u0020\u0064\u0065\u0066\u0069\u006ee\u0064");
};_bee :=len (*_cb .PtaUL );_cb .PtaLL =&_ce .Points {};var (_fc ,_gd float32 ;_aa ,_fd int ;_cbf *_ce .Bitmap ;);for _af :=0;_af < _bee ;_af ++{_fc ,_gd ,_cfd =_cb .PtaUL .GetGeometry (_af );if _cfd !=nil {_ba .Log .Debug ("\u0047e\u0074\u0074\u0069\u006e\u0067\u0020\u0050\u0074\u0061\u0055\u004c \u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_cfd );
return _e .Wrap (_cfd ,_fa ,"\u0050\u0074\u0061\u0055\u004c\u0020\u0047\u0065\u006fm\u0065\u0074\u0072\u0079");};_aa ,_cfd =_cb .ClassIDs .Get (_af );if _cfd !=nil {_ba .Log .Debug ("\u0047\u0065\u0074\u0074\u0069\u006e\u0067\u0020\u0043\u006c\u0061s\u0073\u0049\u0044\u0020\u0066\u0061\u0069\u006c\u0065\u0064:\u0020\u0025\u0076",_cfd );
return _e .Wrap (_cfd ,_fa ,"\u0043l\u0061\u0073\u0073\u0049\u0044");};_cbf ,_cfd =_cb .UndilatedTemplates .GetBitmap (_aa );if _cfd !=nil {_ba .Log .Debug ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0055\u006ed\u0069\u006c\u0061\u0074\u0065\u0064\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_cfd );
return _e .Wrap (_cfd ,_fa ,"\u0055\u006e\u0064\u0069la\u0074\u0065\u0064\u0020\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073");};_fd =_cbf .Height ;_cb .PtaLL .AddPoint (_fc ,_gd +float32 (_fd ));};return nil ;};const JbAddedPixels =6;type Settings struct{MaxCompWidth int ;
MaxCompHeight int ;SizeHaus int ;RankHaus float64 ;Thresh float64 ;WeightFactor float64 ;KeepClassInstances bool ;Components _ce .Component ;Method Method ;};type Classer struct{BaseIndex int ;Settings Settings ;ComponentsNumber *_c .IntSlice ;TemplateAreas *_c .IntSlice ;
Widths map[int ]int ;Heights map[int ]int ;NumberOfClasses int ;ClassInstances *_ce .BitmapsArray ;UndilatedTemplates *_ce .Bitmaps ;DilatedTemplates *_ce .Bitmaps ;TemplatesSize _c .IntsMap ;FgTemplates *_c .NumSlice ;CentroidPoints *_ce .Points ;CentroidPointsTemplates *_ce .Points ;
ClassIDs *_c .IntSlice ;ComponentPageNumbers *_c .IntSlice ;PtaUL *_ce .Points ;PtaLL *_ce .Points ;};func (_ege *Classer )classifyRankHouseOne (_caff *_ce .Boxes ,_aga ,_ccaa ,_efd *_ce .Bitmaps ,_fag *_ce .Points ,_dcd int )(_egec error ){const _gfg ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_cdf ,_acd ,_bbe ,_add float32 ;_cdeb int ;_fgbf ,_dff ,_eacd ,_dda ,_gae *_ce .Bitmap ;_befd ,_bfce bool ;);for _fgae :=0;_fgae < len (_aga .Values );_fgae ++{_dff =_ccaa .Values [_fgae ];_eacd =_efd .Values [_fgae ];_cdf ,_acd ,_egec =_fag .GetGeometry (_fgae );
if _egec !=nil {return _e .Wrapf (_egec ,_gfg ,"\u0066\u0069\u0072\u0073\u0074\u0020\u0067\u0065\u006fm\u0065\u0074\u0072\u0079");};_aba :=len (_ege .UndilatedTemplates .Values );_befd =false ;_efe :=_aafg (_ege ,_dff );for _cdeb =_efe .Next ();_cdeb > -1;
{_dda ,_egec =_ege .UndilatedTemplates .GetBitmap (_cdeb );if _egec !=nil {return _e .Wrap (_egec ,_gfg ,"\u0062\u006d\u0033");};_gae ,_egec =_ege .DilatedTemplates .GetBitmap (_cdeb );if _egec !=nil {return _e .Wrap (_egec ,_gfg ,"\u0062\u006d\u0034");
};_bbe ,_add ,_egec =_ege .CentroidPointsTemplates .GetGeometry (_cdeb );if _egec !=nil {return _e .Wrap (_egec ,_gfg ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0054\u0065\u006d\u0070l\u0061\u0074\u0065\u0073");};_bfce ,_egec =_ce .HausTest (_dff ,_eacd ,_dda ,_gae ,_cdf -_bbe ,_acd -_add ,MaxDiffWidth ,MaxDiffHeight );
if _egec !=nil {return _e .Wrap (_egec ,_gfg ,"");};if _bfce {_befd =true ;if _egec =_ege .ClassIDs .Add (_cdeb );_egec !=nil {return _e .Wrap (_egec ,_gfg ,"");};if _egec =_ege .ComponentPageNumbers .Add (_dcd );_egec !=nil {return _e .Wrap (_egec ,_gfg ,"");
};if _ege .Settings .KeepClassInstances {_dbb ,_fdg :=_ege .ClassInstances .GetBitmaps (_cdeb );if _fdg !=nil {return _e .Wrap (_fdg ,_gfg ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_fgbf ,_fdg =_aga .GetBitmap (_fgae );if _fdg !=nil {return _e .Wrap (_fdg ,_gfg ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");
};_dbb .AddBitmap (_fgbf );_fgg ,_fdg :=_caff .Get (_fgae );if _fdg !=nil {return _e .Wrap (_fdg ,_gfg ,"\u004be\u0065\u0070\u0050\u0069\u0078\u0061a");};_dbb .AddBox (_fgg );};break ;};};if !_befd {if _egec =_ege .ClassIDs .Add (_aba );_egec !=nil {return _e .Wrap (_egec ,_gfg ,"");
};if _egec =_ege .ComponentPageNumbers .Add (_dcd );_egec !=nil {return _e .Wrap (_egec ,_gfg ,"");};_cfab :=&_ce .Bitmaps {};_fgbf ,_egec =_aga .GetBitmap (_fgae );if _egec !=nil {return _e .Wrap (_egec ,_gfg ,"\u0021\u0066\u006f\u0075\u006e\u0064");};
_cfab .Values =append (_cfab .Values ,_fgbf );_ee ,_egf :=_fgbf .Width ,_fgbf .Height ;_ege .TemplatesSize .Add (uint64 (_egf )*uint64 (_ee ),_aba );_fee ,_gcb :=_caff .Get (_fgae );if _gcb !=nil {return _e .Wrap (_gcb ,_gfg ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_cfab .AddBox (_fee );_ege .ClassInstances .AddBitmaps (_cfab );_ege .CentroidPointsTemplates .AddPoint (_cdf ,_acd );_ege .UndilatedTemplates .AddBitmap (_dff );_ege .DilatedTemplates .AddBitmap (_eacd );};};return nil ;};func DefaultSettings ()Settings {_adab :=&Settings {};
_adab .SetDefault ();return *_adab };func (_gea *Settings )SetDefault (){if _gea .MaxCompWidth ==0{switch _gea .Components {case _ce .ComponentConn :_gea .MaxCompWidth =MaxConnCompWidth ;case _ce .ComponentCharacters :_gea .MaxCompWidth =MaxCharCompWidth ;
case _ce .ComponentWords :_gea .MaxCompWidth =MaxWordCompWidth ;};};if _gea .MaxCompHeight ==0{_gea .MaxCompHeight =MaxCompHeight ;};if _gea .Thresh ==0.0{_gea .Thresh =0.9;};if _gea .WeightFactor ==0.0{_gea .WeightFactor =0.75;};if _gea .RankHaus ==0.0{_gea .RankHaus =0.97;
};if _gea .SizeHaus ==0{_gea .SizeHaus =2;};};func (_aaa *similarTemplatesFinder )Next ()int {var (_bacb ,_adc ,_fbe ,_bbfg int ;_fdd bool ;_afc *_ce .Bitmap ;_bea error ;);for {if _aaa .Index >=25{return -1;};_adc =_aaa .Width +TwoByTwoWalk [2*_aaa .Index ];
_bacb =_aaa .Height +TwoByTwoWalk [2*_aaa .Index +1];if _bacb < 1||_adc < 1{_aaa .Index ++;continue ;};if len (_aaa .CurrentNumbers )==0{_aaa .CurrentNumbers ,_fdd =_aaa .Classer .TemplatesSize .GetSlice (uint64 (_adc )*uint64 (_bacb ));if !_fdd {_aaa .Index ++;
continue ;};_aaa .N =0;};_fbe =len (_aaa .CurrentNumbers );for ;_aaa .N < _fbe ;_aaa .N ++{_bbfg =_aaa .CurrentNumbers [_aaa .N ];_afc ,_bea =_aaa .Classer .DilatedTemplates .GetBitmap (_bbfg );if _bea !=nil {_ba .Log .Debug ("\u0046\u0069\u006e\u0064\u004e\u0065\u0078\u0074\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u003a\u0020\u0074\u0065\u006d\u0070\u006c\u0061t\u0065\u0020\u006e\u006f\u0074 \u0066\u006fu\u006e\u0064\u003a\u0020");
return 0;};if _afc .Width -2*JbAddedPixels ==_adc &&_afc .Height -2*JbAddedPixels ==_bacb {return _bbfg ;};};_aaa .Index ++;_aaa .CurrentNumbers =nil ;};};func _aafg (_acfa *Classer ,_afff *_ce .Bitmap )*similarTemplatesFinder {return &similarTemplatesFinder {Width :_afff .Width ,Height :_afff .Height ,Classer :_acfa };
};func (_aag *Classer )classifyRankHouseNonOne (_dcfg *_ce .Boxes ,_dgg ,_fce ,_gaa *_ce .Bitmaps ,_dgb *_ce .Points ,_dgd *_c .NumSlice ,_bbf int )(_ged error ){const _ecd ="\u0043\u006c\u0061\u0073s\u0065\u0072\u002e\u0063\u006c\u0061\u0073\u0073\u0069\u0066y\u0052a\u006e\u006b\u0048\u006f\u0075\u0073\u0065O\u006e\u0065";
var (_egece ,_da ,_ffb ,_fba float32 ;_fbad ,_cfcc ,_dab int ;_feab ,_ccbe ,_bgf ,_ecdg ,_ccac *_ce .Bitmap ;_gb ,_ebf bool ;);_dfb :=_ce .MakePixelSumTab8 ();for _bab :=0;_bab < len (_dgg .Values );_bab ++{if _ccbe ,_ged =_fce .GetBitmap (_bab );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"b\u006d\u0073\u0031\u002e\u0047\u0065\u0074\u0028\u0069\u0029");
};if _fbad ,_ged =_dgd .GetInt (_bab );_ged !=nil {_ba .Log .Trace ("\u0047\u0065t\u0074\u0069\u006e\u0067 \u0046\u0047T\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u0073 \u0061\u0074\u003a\u0020\u0025\u0064\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_bab ,_ged );
};if _bgf ,_ged =_gaa .GetBitmap (_bab );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"b\u006d\u0073\u0032\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};if _egece ,_da ,_ged =_dgb .GetGeometry (_bab );_ged !=nil {return _e .Wrapf (_ged ,_ecd ,"\u0070t\u0061[\u0069\u005d\u002e\u0047\u0065\u006f\u006d\u0065\u0074\u0072\u0079");
};_dbf :=len (_aag .UndilatedTemplates .Values );_gb =false ;_fabc :=_aafg (_aag ,_ccbe );for _dab =_fabc .Next ();_dab > -1;{if _ecdg ,_ged =_aag .UndilatedTemplates .GetBitmap (_dab );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"\u0070\u0069\u0078\u0061\u0074\u002e\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");
};if _cfcc ,_ged =_aag .FgTemplates .GetInt (_dab );_ged !=nil {_ba .Log .Trace ("\u0047\u0065\u0074\u0074\u0069\u006eg\u0020\u0046\u0047\u0054\u0065\u006d\u0070\u006c\u0061\u0074\u0065\u005b\u0025d\u005d\u0020\u0066\u0061\u0069\u006c\u0065d\u003a\u0020\u0025\u0076",_dab ,_ged );
};if _ccac ,_ged =_aag .DilatedTemplates .GetBitmap (_dab );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"\u0070\u0069\u0078\u0061\u0074\u0064\u005b\u0069\u0043l\u0061\u0073\u0073\u005d");};if _ffb ,_fba ,_ged =_aag .CentroidPointsTemplates .GetGeometry (_dab );
_ged !=nil {return _e .Wrap (_ged ,_ecd ,"\u0043\u0065\u006et\u0072\u006f\u0069\u0064P\u006f\u0069\u006e\u0074\u0073\u0054\u0065m\u0070\u006c\u0061\u0074\u0065\u0073\u005b\u0069\u0043\u006c\u0061\u0073\u0073\u005d");};_ebf ,_ged =_ce .RankHausTest (_ccbe ,_bgf ,_ecdg ,_ccac ,_egece -_ffb ,_da -_fba ,MaxDiffWidth ,MaxDiffHeight ,_fbad ,_cfcc ,float32 (_aag .Settings .RankHaus ),_dfb );
if _ged !=nil {return _e .Wrap (_ged ,_ecd ,"");};if _ebf {_gb =true ;if _ged =_aag .ClassIDs .Add (_dab );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"");};if _ged =_aag .ComponentPageNumbers .Add (_bbf );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"");};if _aag .Settings .KeepClassInstances {_bdfg ,_fgeg :=_aag .ClassInstances .GetBitmaps (_dab );
if _fgeg !=nil {return _e .Wrap (_fgeg ,_ecd ,"\u0063\u002e\u0050\u0069\u0078\u0061\u0061\u002e\u0047\u0065\u0074B\u0069\u0074\u006d\u0061\u0070\u0073\u0028\u0069\u0043\u006ca\u0073\u0073\u0029");};if _feab ,_fgeg =_dgg .GetBitmap (_bab );_fgeg !=nil {return _e .Wrap (_fgeg ,_ecd ,"\u0070i\u0078\u0061\u005b\u0069\u005d");
};_bdfg .Values =append (_bdfg .Values ,_feab );_gbd ,_fgeg :=_dcfg .Get (_bab );if _fgeg !=nil {return _e .Wrap (_fgeg ,_ecd ,"b\u006f\u0078\u0061\u002e\u0047\u0065\u0074\u0028\u0069\u0029");};_bdfg .Boxes =append (_bdfg .Boxes ,_gbd );};break ;};};if !_gb {if _ged =_aag .ClassIDs .Add (_dbf );
_ged !=nil {return _e .Wrap (_ged ,_ecd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _ged =_aag .ComponentPageNumbers .Add (_bbf );_ged !=nil {return _e .Wrap (_ged ,_ecd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_feea :=&_ce .Bitmaps {};_feab =_dgg .Values [_bab ];
_feea .AddBitmap (_feab );_cgf ,_ffc :=_feab .Width ,_feab .Height ;_aag .TemplatesSize .Add (uint64 (_cgf )*uint64 (_ffc ),_dbf );_fcgc ,_gfe :=_dcfg .Get (_bab );if _gfe !=nil {return _e .Wrap (_gfe ,_ecd ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_feea .AddBox (_fcgc );
_aag .ClassInstances .AddBitmaps (_feea );_aag .CentroidPointsTemplates .AddPoint (_egece ,_da );_aag .UndilatedTemplates .AddBitmap (_ccbe );_aag .DilatedTemplates .AddBitmap (_bgf );_aag .FgTemplates .AddInt (_fbad );};};_aag .NumberOfClasses =len (_aag .UndilatedTemplates .Values );
return nil ;};const (MaxConnCompWidth =350;MaxCharCompWidth =350;MaxWordCompWidth =1000;MaxCompHeight =120;);func (_cab *Classer )classifyCorrelation (_bdc *_ce .Boxes ,_ceg *_ce .Bitmaps ,_ddd int )error {const _ced ="\u0063\u006c\u0061\u0073si\u0066\u0079\u0043\u006f\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e";
if _bdc ==nil {return _e .Error (_ced ,"\u006e\u0065\u0077\u0043\u006f\u006d\u0070\u006f\u006e\u0065\u006e\u0074\u0073\u0020\u0062\u006f\u0075\u006e\u0064\u0069\u006e\u0067\u0020\u0062o\u0078\u0065\u0073\u0020\u006eo\u0074\u0020f\u006f\u0075\u006e\u0064");
};if _ceg ==nil {return _e .Error (_ced ,"\u006e\u0065wC\u006f\u006d\u0070o\u006e\u0065\u006e\u0074s b\u0069tm\u0061\u0070\u0020\u0061\u0072\u0072\u0061y \u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");};_fgb :=len (_ceg .Values );if _fgb ==0{_ba .Log .Debug ("\u0063l\u0061\u0073s\u0069\u0066\u0079C\u006f\u0072\u0072\u0065\u006c\u0061\u0074i\u006f\u006e\u0020\u002d\u0020\u0070r\u006f\u0076\u0069\u0064\u0065\u0064\u0020\u0070\u0069\u0078\u0061s\u0020\u0069\u0073\u0020\u0065\u006d\u0070\u0074\u0079");
return nil ;};var (_acg ,_gfc *_ce .Bitmap ;_bcb error ;);_ccb :=&_ce .Bitmaps {Values :make ([]*_ce .Bitmap ,_fgb )};for _gc ,_bca :=range _ceg .Values {_gfc ,_bcb =_bca .AddBorderGeneral (JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,JbAddedPixels ,0);
if _bcb !=nil {return _e .Wrap (_bcb ,_ced ,"");};_ccb .Values [_gc ]=_gfc ;};_bgc :=_cab .FgTemplates ;_fbg :=_ce .MakePixelSumTab8 ();_bbc :=_ce .MakePixelCentroidTab8 ();_gdg :=make ([]int ,_fgb );_cdb :=make ([][]int ,_fgb );_ggf :=_ce .Points (make ([]_ce .Point ,_fgb ));
_bdf :=&_ggf ;var (_fea ,_cgc int ;_ebb ,_db ,_ef int ;_eac ,_efg int ;_abe byte ;);for _ge ,_edf :=range _ccb .Values {_cdb [_ge ]=make ([]int ,_edf .Height );_fea =0;_cgc =0;_db =(_edf .Height -1)*_edf .RowStride ;_ebb =0;for _efg =_edf .Height -1;_efg >=0;
_efg ,_db =_efg -1,_db -_edf .RowStride {_cdb [_ge ][_efg ]=_ebb ;_ef =0;for _eac =0;_eac < _edf .RowStride ;_eac ++{_abe =_edf .Data [_db +_eac ];_ef +=_fbg [_abe ];_fea +=_bbc [_abe ]+_eac *8*_fbg [_abe ];};_ebb +=_ef ;_cgc +=_ef *_efg ;};_gdg [_ge ]=_ebb ;
if _ebb > 0{(*_bdf )[_ge ]=_ce .Point {X :float32 (_fea )/float32 (_ebb ),Y :float32 (_cgc )/float32 (_ebb )};}else {(*_bdf )[_ge ]=_ce .Point {X :float32 (_edf .Width )/float32 (2),Y :float32 (_edf .Height )/float32 (2)};};};if _bcb =_cab .CentroidPoints .Add (_bdf );
_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0063\u0065\u006et\u0072\u006f\u0069\u0064\u0020\u0061\u0064\u0064");};var (_baa ,_ddb ,_fga int ;_ec float64 ;_fab ,_aed ,_dde ,_ceb float32 ;_bfc ,_bec _ce .Point ;_faa bool ;_df *similarTemplatesFinder ;_abef int ;
_egd *_ce .Bitmap ;_eaf *_g .Rectangle ;_cabg *_ce .Bitmaps ;);for _abef ,_gfc =range _ccb .Values {_ddb =_gdg [_abef ];if _fab ,_aed ,_bcb =_bdf .GetGeometry (_abef );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0070t\u0061\u0020\u002d\u0020\u0069");};
_faa =false ;_acf :=len (_cab .UndilatedTemplates .Values );_df =_aafg (_cab ,_gfc );for _fcg :=_df .Next ();_fcg > -1;{if _egd ,_bcb =_cab .UndilatedTemplates .GetBitmap (_fcg );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0075\u006e\u0069dl\u0061\u0074\u0065\u0064\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0062\u006d\u0032");
};if _fga ,_bcb =_bgc .GetInt (_fcg );_bcb !=nil {_ba .Log .Trace ("\u0046\u0047\u0020T\u0065\u006d\u0070\u006ca\u0074\u0065\u0020\u005b\u0069\u0063\u006ca\u0073\u0073\u005d\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_bcb );};if _dde ,_ceb ,_bcb =_cab .CentroidPointsTemplates .GetGeometry (_fcg );
_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0043\u0065\u006e\u0074\u0072\u006f\u0069\u0064\u0050\u006f\u0069\u006e\u0074T\u0065\u006d\u0070\u006c\u0061\u0074e\u0073\u005b\u0069\u0063\u006c\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u00782\u002c\u0079\u0032\u0020");
};if _cab .Settings .WeightFactor > 0.0{if _baa ,_bcb =_cab .TemplateAreas .Get (_fcg );_bcb !=nil {_ba .Log .Trace ("\u0054\u0065\u006dp\u006c\u0061\u0074\u0065A\u0072\u0065\u0061\u0073\u005b\u0069\u0063l\u0061\u0073\u0073\u005d\u0020\u003d\u0020\u0061\u0072\u0065\u0061\u0020\u0025\u0076",_bcb );
};_ec =_cab .Settings .Thresh +(1.0-_cab .Settings .Thresh )*_cab .Settings .WeightFactor *float64 (_fga )/float64 (_baa );}else {_ec =_cab .Settings .Thresh ;};_bcc ,_cde :=_ce .CorrelationScoreThresholded (_gfc ,_egd ,_ddb ,_fga ,_bfc .X -_bec .X ,_bfc .Y -_bec .Y ,MaxDiffWidth ,MaxDiffHeight ,_fbg ,_cdb [_abef ],float32 (_ec ));
if _cde !=nil {return _e .Wrap (_cde ,_ced ,"");};if _ada {var (_efgc ,_bebb float64 ;_dcf ,_agg int ;);_efgc ,_cde =_ce .CorrelationScore (_gfc ,_egd ,_ddb ,_fga ,_fab -_dde ,_aed -_ceb ,MaxDiffWidth ,MaxDiffHeight ,_fbg );if _cde !=nil {return _e .Wrap (_cde ,_ced ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_bebb ,_cde =_ce .CorrelationScoreSimple (_gfc ,_egd ,_ddb ,_fga ,_fab -_dde ,_aed -_ceb ,MaxDiffWidth ,MaxDiffHeight ,_fbg );if _cde !=nil {return _e .Wrap (_cde ,_ced ,"d\u0065\u0062\u0075\u0067Co\u0072r\u0065\u006c\u0061\u0074\u0069o\u006e\u0053\u0063\u006f\u0072\u0065");
};_dcf =int (_a .Sqrt (_efgc *float64 (_ddb )*float64 (_fga )));_agg =int (_a .Sqrt (_bebb *float64 (_ddb )*float64 (_fga )));if (_efgc >=_ec )!=(_bebb >=_ec ){return _e .Errorf (_ced ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043\u006f\u0072r\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020\u0073\u0063\u006f\u0072\u0065\u0020\u006d\u0069\u0073\u006d\u0061\u0074\u0063\u0068\u0020-\u0020\u0025d\u0028\u00250\u002e\u0034\u0066\u002c\u0020\u0025\u0076\u0029\u0020\u0076\u0073\u0020\u0025d(\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0076)\u0020\u0025\u0030\u002e\u0034\u0066",_dcf ,_efgc ,_efgc >=_ec ,_agg ,_bebb ,_bebb >=_ec ,_efgc -_bebb );
};if _efgc >=_ec !=_bcc {return _e .Errorf (_ced ,"\u0064\u0065\u0062\u0075\u0067\u0020\u0043o\u0072\u0072\u0065\u006c\u0061\u0074\u0069\u006f\u006e \u0073\u0063\u006f\u0072\u0065 \u004d\u0069\u0073\u006d\u0061t\u0063\u0068 \u0062\u0065\u0074w\u0065\u0065\u006e\u0020\u0063\u006frr\u0065\u006c\u0061\u0074\u0069\u006f\u006e\u0020/\u0020\u0074\u0068\u0072\u0065s\u0068\u006f\u006c\u0064\u002e\u0020\u0043\u006f\u006dpa\u0072\u0069\u0073\u006f\u006e:\u0020\u0025\u0030\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u002c\u0020\u0025\u0064\u0029\u0020\u003e\u003d\u0020\u00250\u002e\u0034\u0066\u0028\u0025\u0030\u002e\u0034\u0066\u0029\u0020\u0076\u0073\u0020\u0025\u0076",_efgc ,_efgc *float64 (_ddb )*float64 (_fga ),_dcf ,_ec ,float32 (_ec )*float32 (_ddb )*float32 (_fga ),_bcc );
};};if _bcc {_faa =true ;if _cde =_cab .ClassIDs .Add (_fcg );_cde !=nil {return _e .Wrap (_cde ,_ced ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");};if _cde =_cab .ComponentPageNumbers .Add (_ddd );_cde !=nil {return _e .Wrap (_cde ,_ced ,"\u006f\u0076\u0065\u0072\u0054\u0068\u0072\u0065\u0073\u0068\u006f\u006c\u0064");
};if _cab .Settings .KeepClassInstances {if _acg ,_cde =_ceg .GetBitmap (_abef );_cde !=nil {return _e .Wrap (_cde ,_ced ,"\u004b\u0065\u0065\u0070Cl\u0061\u0073\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065\u0073\u0020\u002d \u0069");};if _cabg ,_cde =_cab .ClassInstances .GetBitmaps (_fcg );
_cde !=nil {return _e .Wrap (_cde ,_ced ,"K\u0065\u0065\u0070\u0043\u006c\u0061s\u0073\u0049\u006e\u0073\u0074\u0061\u006e\u0063\u0065s\u0020\u002d\u0020i\u0043l\u0061\u0073\u0073");};_cabg .AddBitmap (_acg );if _eaf ,_cde =_bdc .Get (_abef );_cde !=nil {return _e .Wrap (_cde ,_ced ,"\u004be\u0065p\u0043\u006c\u0061\u0073\u0073I\u006e\u0073t\u0061\u006e\u0063\u0065\u0073");
};_cabg .AddBox (_eaf );};break ;};};if !_faa {if _bcb =_cab .ClassIDs .Add (_acf );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0021\u0066\u006f\u0075\u006e\u0064");};if _bcb =_cab .ComponentPageNumbers .Add (_ddd );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0021\u0066\u006f\u0075\u006e\u0064");
};_cabg =&_ce .Bitmaps {};if _acg ,_bcb =_ceg .GetBitmap (_abef );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_cabg .AddBitmap (_acg );_abc ,_fbb :=_acg .Width ,_acg .Height ;_cda :=uint64 (_fbb )*uint64 (_abc );_cab .TemplatesSize .Add (_cda ,_acf );
if _eaf ,_bcb =_bdc .Get (_abef );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0021\u0066\u006f\u0075\u006e\u0064");};_cabg .AddBox (_eaf );_cab .ClassInstances .AddBitmaps (_cabg );_cab .CentroidPointsTemplates .AddPoint (_fab ,_aed );_cab .FgTemplates .AddInt (_ddb );
_cab .UndilatedTemplates .AddBitmap (_acg );_baa =(_gfc .Width -2*JbAddedPixels )*(_gfc .Height -2*JbAddedPixels );if _bcb =_cab .TemplateAreas .Add (_baa );_bcb !=nil {return _e .Wrap (_bcb ,_ced ,"\u0021\u0066\u006f\u0075\u006e\u0064");};};};_cab .NumberOfClasses =len (_cab .UndilatedTemplates .Values );
return nil ;};