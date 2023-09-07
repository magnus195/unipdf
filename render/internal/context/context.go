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

package context ;import (_ag "errors";_afb "github.com/unidoc/freetype/truetype";_b "github.com/unidoc/unipdf/v3/core";_d "github.com/unidoc/unipdf/v3/internal/textencoding";_f "github.com/unidoc/unipdf/v3/internal/transform";_e "github.com/unidoc/unipdf/v3/model";
_ca "golang.org/x/image/font";_c "image";_af "image/color";);type FillRule int ;func (_aec *TextFont )CharcodeToRunes (charcode _d .CharCode )(_d .CharCode ,[]rune ){_ba :=[]_d .CharCode {charcode };if _aec ._fag ==nil ||_aec ._fag ==_aec .Font {if _aec .Font .IsSimple ()&&_aec ._cafb !=nil {if _eea :=_aec ._cafb .Index (rune (charcode ));
_eea > 0{return charcode ,[]rune {rune (charcode )};};};return charcode ,_aec .Font .CharcodesToUnicode (_ba );};_fgc :=_aec ._fag .CharcodesToUnicode (_ba );_ad ,_ :=_aec .Font .RunesToCharcodeBytes (_fgc );_ddd :=_aec .Font .BytesToCharcodes (_ad );_dea :=charcode ;
if len (_ddd )> 0&&_ddd [0]!=0{_dea =_ddd [0];};return _dea ,_fgc ;};type TextFont struct{Font *_e .PdfFont ;Size float64 ;_cafb *_afb .Font ;_fag *_e .PdfFont ;};func (_dag *TextState )ProcTj (data []byte ,ctx Context ){_fc :=_dag .Tf .Size ;_cge :=_dag .Th /100.0;
_bda :=_dag .GlobalScale ;_ecde :=_f .NewMatrix (_fc *_cge ,0,0,_fc ,0,_dag .Ts );_fcc :=ctx .Matrix ();_eba :=_fcc .Clone ().Mult (_dag .Tm .Clone ().Mult (_ecde )).ScalingFactorY ();_gef :=_dag .Tf .NewFace (_eba );_ea :=_dag .Tf .BytesToCharcodes (data );
for _ ,_bgd :=range _ea {_bce ,_ebce :=_dag .Tf .CharcodeToRunes (_bgd );_cde :=string (_ebce );if _cde =="\u0000"{continue ;};_gfb :=_fcc .Clone ().Mult (_dag .Tm .Clone ().Mult (_ecde ));_bbc :=_gfb .ScalingFactorY ();_gfb =_gfb .Scale (1/_bbc ,-1/_bbc );
if _dag .Tr !=TextRenderingModeInvisible {ctx .SetMatrix (_gfb );ctx .DrawString (_cde ,_gef ,0,0);ctx .SetMatrix (_fcc );};_aff :=0.0;if _cde =="\u0020"{_aff =_dag .Tw ;};_dbba ,_ ,_cgdb :=_dag .Tf .GetCharMetrics (_bce );if _cgdb {_dbba =_dbba *0.001*_fc ;
}else {_dbba ,_ =ctx .MeasureString (_cde ,_gef );_dbba =_dbba /_bda ;};_dba :=(_dbba +_dag .Tc +_aff )*_cge ;_dag .Tm =_dag .Tm .Mult (_f .TranslationMatrix (_dba ,0));};};func NewTextState ()TextState {return TextState {Th :100,Tm :_f .IdentityMatrix (),Tlm :_f .IdentityMatrix ()};
};type TextState struct{Tc float64 ;Tw float64 ;Th float64 ;Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _f .Matrix ;Tlm _f .Matrix ;Tr TextRenderingMode ;GlobalScale float64 ;};func (_gge *TextState )Reset (){_gge .Tm =_f .IdentityMatrix ();_gge .Tlm =_f .IdentityMatrix ()};
func (_dfe *TextFont )WithSize (size float64 ,originalFont *_e .PdfFont )*TextFont {return &TextFont {Font :_dfe .Font ,Size :size ,_cafb :_dfe ._cafb ,_fag :originalFont };};type Gradient interface{Pattern ;AddColorStop (_g float64 ,_bf _af .Color );};
const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);func (_edd *TextFont )GetCharMetrics (code _d .CharCode )(float64 ,float64 ,bool ){if _fae ,_eff :=_edd .Font .GetCharMetrics (code );_eff &&_fae .Wx !=0{return _fae .Wx ,_fae .Wy ,_eff ;};if _edd ._fag ==nil {return 0,0,false ;
};_afd ,_ebc :=_edd ._fag .GetCharMetrics (code );return _afd .Wx ,_afd .Wy ,_ebc &&_afd .Wx !=0;};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;);func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_dbb ,_cgge :=_e .NewPdfFontFromTTFFile (filePath );
if _cgge !=nil {return nil ,_cgge ;};return NewTextFont (_dbb ,size );};func NewTextFont (font *_e .PdfFont ,size float64 )(*TextFont ,error ){_gfg :=font .FontDescriptor ();if _gfg ==nil {return nil ,_ag .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");
};_aag ,_fgb :=_b .GetStream (_gfg .FontFile2 );if !_fgb {return nil ,_ag .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};_cfb ,_egc :=_b .DecodeStream (_aag );
if _egc !=nil {return nil ,_egc ;};_cbb ,_egc :=_afb .Parse (_cfb );if _egc !=nil {return nil ,_egc ;};return &TextFont {Font :font ,Size :size ,_cafb :_cbb },nil ;};type LineJoin int ;func (_egcc *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_egcc .Tm =_f .NewMatrix (a ,b ,c ,d ,e ,f );
_egcc .Tlm =_egcc .Tm .Clone ();};func (_fgba *TextState )ProcTD (tx ,ty float64 ){_fgba .Tl =-ty ;_fgba .ProcTd (tx ,ty )};func (_faf *TextFont )BytesToCharcodes (data []byte )[]_d .CharCode {if _faf ._fag !=nil {return _faf ._fag .BytesToCharcodes (data );
};return _faf .Font .BytesToCharcodes (data );};type Context interface{Push ();Pop ();Matrix ()_f .Matrix ;SetMatrix (_gb _f .Matrix );Translate (_ec ,_cd float64 );Scale (_bb ,_da float64 );Rotate (_dg float64 );MoveTo (_gf ,_fg float64 );LineTo (_ed ,_bbb float64 );
CubicTo (_caf ,_fa ,_cg ,_fe ,_cgg ,_ff float64 );QuadraticTo (_gfa ,_db ,_cff ,_gff float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();ResetClip ();LineWidth ()float64 ;SetLineWidth (_cgd float64 );SetLineCap (_ae LineCap );SetLineJoin (_aea LineJoin );
SetDash (_dd ...float64 );SetDashOffset (_aeag float64 );Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_gd ,_bba ,_gdd ,_cc float64 );SetFillRGBA (_ddg ,_aa ,_ece ,_ge float64 );SetFillStyle (_cgdg Pattern );SetFillRule (_geg FillRule );
SetStrokeRGBA (_ccb ,_feb ,_gdc ,_dda float64 );SetStrokeStyle (_ecd Pattern );FillPattern ()Pattern ;StrokePattern ()Pattern ;TextState ()*TextState ;DrawString (_fd string ,_gbb _ca .Face ,_ef ,_df float64 );MeasureString (_dgd string ,_ffd _ca .Face )(_cad ,_agf float64 );
DrawRectangle (_dc ,_age ,_eb ,_bd float64 );DrawImage (_eg _c .Image ,_bbd ,_bc int );DrawImageAnchored (_cb _c .Image ,_cdb ,_afc int ,_aef ,_fab float64 );Height ()int ;Width ()int ;};type TextRenderingMode int ;const (LineCapRound LineCap =iota ;LineCapButt ;
LineCapSquare ;);func (_bab *TextState )ProcQ (data []byte ,ctx Context ){_bab .ProcTStar ();_bab .ProcTj (data ,ctx )};type LineCap int ;func (_gg *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_gg .Tw =aw ;_gg .Tc =ac ;_gg .ProcQ (data ,ctx );
};func (_edb *TextState )Translate (tx ,ty float64 ){_edb .Tm =_edb .Tm .Mult (_f .TranslationMatrix (tx ,ty ));};type Pattern interface{ColorAt (_cf ,_ee int )_af .Color ;};const (TextRenderingModeFill TextRenderingMode =iota ;TextRenderingModeStroke ;
TextRenderingModeFillStroke ;TextRenderingModeInvisible ;TextRenderingModeFillClip ;TextRenderingModeStrokeClip ;TextRenderingModeFillStrokeClip ;TextRenderingModeClip ;);func (_aed *TextState )ProcTStar (){_aed .ProcTd (0,-_aed .Tl )};func (_bg *TextState )ProcTd (tx ,ty float64 ){_bg .Tlm .Concat (_f .TranslationMatrix (tx ,ty ));
_bg .Tm =_bg .Tlm .Clone ();};func (_de *TextFont )NewFace (size float64 )_ca .Face {return _afb .NewFace (_de ._cafb ,&_afb .Options {Size :size });};func (_ce *TextState )ProcTf (font *TextFont ){_ce .Tf =font };