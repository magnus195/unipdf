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

package context ;import (_f "errors";_ce "github.com/unidoc/freetype/truetype";_cf "github.com/unidoc/unipdf/v3/core";_ag "github.com/unidoc/unipdf/v3/internal/cmap";_cg "github.com/unidoc/unipdf/v3/internal/textencoding";_eb "github.com/unidoc/unipdf/v3/internal/transform";
_ff "github.com/unidoc/unipdf/v3/model";_b "golang.org/x/image/font";_c "image";_e "image/color";_d "strings";);func (_efd *TextFont )GetCharMetrics (code _cg .CharCode )(float64 ,float64 ,bool ){if _gce ,_dbc :=_efd .Font .GetCharMetrics (code );_dbc &&_gce .Wx !=0{return _gce .Wx ,_gce .Wy ,_dbc ;
};if _efd ._cgcc ==nil {return 0,0,false ;};_eff ,_fcb :=_efd ._cgcc .GetCharMetrics (code );return _eff .Wx ,_eff .Wy ,_fcb &&_eff .Wx !=0;};type TextState struct{Tc float64 ;Tw float64 ;Th float64 ;Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _eb .Matrix ;
Tlm _eb .Matrix ;Tr TextRenderingMode ;GlobalScale float64 ;};func (_abg *TextState )ProcTj (data []byte ,ctx Context ){_dfb :=_abg .Tf .Size ;_gfa :=_abg .Th /100.0;_cee :=_abg .GlobalScale ;_ege :=_eb .NewMatrix (_dfb *_gfa ,0,0,_dfb ,0,_abg .Ts );_cfb :=ctx .Matrix ();
_cgg :=_cfb .Clone ().Mult (_abg .Tm .Clone ().Mult (_ege )).ScalingFactorY ();_bg :=_abg .Tf .NewFace (_cgg );_aba :=_abg .Tf .BytesToCharcodes (data );for _ ,_aaa :=range _aba {_bfb ,_fee :=_abg .Tf .CharcodeToRunes (_aaa );_bge :=string (_fee );if _bge =="\u0000"{continue ;
};_fd :=_cfb .Clone ().Mult (_abg .Tm .Clone ().Mult (_ege ));_dga :=_fd .ScalingFactorY ();_fd =_fd .Scale (1/_dga ,-1/_dga );if _abg .Tr !=TextRenderingModeInvisible {ctx .SetMatrix (_fd );ctx .DrawString (_bge ,_bg ,0,0);ctx .SetMatrix (_cfb );};_cd :=0.0;
if _bge =="\u0020"{_cd =_abg .Tw ;};_gbb ,_ ,_bef :=_abg .Tf .GetCharMetrics (_bfb );if _bef {_gbb =_gbb *0.001*_dfb ;}else {_gbb ,_ =ctx .MeasureString (_bge ,_bg );_gbb =_gbb /_cee ;};_eea :=(_gbb +_abg .Tc +_cd )*_gfa ;_abg .Tm =_abg .Tm .Mult (_eb .TranslationMatrix (_eea ,0));
};};func (_ddd *TextFont )NewFace (size float64 )_b .Face {return _ce .NewFace (_ddd ._bf ,&_ce .Options {Size :size });};func (_cgef *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_cgef .Tw =aw ;_cgef .Tc =ac ;_cgef .ProcQ (data ,ctx );
};type Pattern interface{ColorAt (_ed ,_bd int )_e .Color ;};func (_fge *TextFont )CharcodeToRunes (charcode _cg .CharCode )(_cg .CharCode ,[]rune ){_fbc :=[]_cg .CharCode {charcode };if _fge ._cgcc ==nil ||_fge ._cgcc ==_fge .Font {return _fge .charcodeToRunesSimple (charcode );
};_cad :=_fge ._cgcc .CharcodesToUnicode (_fbc );_dcf ,_ :=_fge .Font .RunesToCharcodeBytes (_cad );_ge :=_fge .Font .BytesToCharcodes (_dcf );_fcg :=charcode ;if len (_ge )> 0&&_ge [0]!=0{_fcg =_ge [0];};if string (_cad )==string (_ag .MissingCodeRune )&&_fge ._cgcc .BaseFont ()==_fge .Font .BaseFont (){return _fge .charcodeToRunesSimple (charcode );
};return _fcg ,_cad ;};func (_gdb *TextFont )charcodeToRunesSimple (_fcc _cg .CharCode )(_cg .CharCode ,[]rune ){_cgeg :=[]_cg .CharCode {_fcc };if _gdb .Font .IsSimple ()&&_gdb ._bf !=nil {if _gea :=_gdb ._bf .Index (rune (_fcc ));_gea > 0{return _fcc ,[]rune {rune (_fcc )};
};};if _gdb ._bf !=nil &&!_gdb ._bf .HasCmap ()&&_d .Contains (_gdb .Font .Encoder ().String (),"\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-"){if _dgb :=_gdb ._bf .Index (rune (_fcc ));_dgb > 0{return _fcc ,[]rune {rune (_fcc )};};};return _fcc ,_gdb .Font .CharcodesToUnicode (_cgeg );
};type FillRule int ;type TextFont struct{Font *_ff .PdfFont ;Size float64 ;_bf *_ce .Font ;_cgcc *_ff .PdfFont ;};const (LineCapRound LineCap =iota ;LineCapButt ;LineCapSquare ;);func (_dca *TextState )Reset (){_dca .Tm =_eb .IdentityMatrix ();_dca .Tlm =_eb .IdentityMatrix ()};
type TextRenderingMode int ;func (_acb *TextFont )WithSize (size float64 ,originalFont *_ff .PdfFont )*TextFont {return &TextFont {Font :_acb .Font ,Size :size ,_bf :_acb ._bf ,_cgcc :originalFont };};type Gradient interface{Pattern ;AddColorStop (_cff float64 ,_g _e .Color );
};func (_bff *TextState )ProcQ (data []byte ,ctx Context ){_bff .ProcTStar ();_bff .ProcTj (data ,ctx )};func (_edf *TextState )ProcTD (tx ,ty float64 ){_edf .Tl =-ty ;_edf .ProcTd (tx ,ty )};func NewTextState ()TextState {return TextState {Th :100,Tm :_eb .IdentityMatrix (),Tlm :_eb .IdentityMatrix ()};
};type LineJoin int ;const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);func (_df *TextState )ProcTd (tx ,ty float64 ){_df .Tlm .Concat (_eb .TranslationMatrix (tx ,ty ));_df .Tm =_df .Tlm .Clone ();};type LineCap int ;func (_ecd *TextState )ProcTf (font *TextFont ){_ecd .Tf =font };
type Context interface{Push ();Pop ();Matrix ()_eb .Matrix ;SetMatrix (_cge _eb .Matrix );Translate (_ea ,_af float64 );Scale (_edb ,_cb float64 );Rotate (_dg float64 );MoveTo (_da ,_cgc float64 );LineTo (_gg ,_ca float64 );CubicTo (_afa ,_cbc ,_de ,_ffa ,_ggc ,_cfd float64 );
QuadraticTo (_afd ,_gd ,_ba ,_ac float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();ResetClip ();LineWidth ()float64 ;SetLineWidth (_be float64 );SetLineCap (_bc LineCap );SetLineJoin (_cfdd LineJoin );SetDash (_agd ...float64 );
SetDashOffset (_dab float64 );Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_fe ,_fc ,_eg ,_dabf float64 );SetFillRGBA (_bee ,_gda ,_dd ,_bb float64 );SetFillStyle (_gc Pattern );SetFillRule (_ede FillRule );SetStrokeRGBA (_cbe ,_ga ,_db ,_gf float64 );
SetStrokeStyle (_dc Pattern );FillPattern ()Pattern ;StrokePattern ()Pattern ;TextState ()*TextState ;DrawString (_ee string ,_fg _b .Face ,_dbf ,_afac float64 );MeasureString (_agdc string ,_ffe _b .Face )(_ggf ,_ae float64 );DrawRectangle (_ggcc ,_bbf ,_cc ,_fb float64 );
DrawImage (_ef _c .Image ,_aca ,_ec int );DrawImageAnchored (_caf _c .Image ,_caa ,_dbd int ,_dce ,_fce float64 );Height ()int ;Width ()int ;};func (_dgbd *TextState )ProcTStar (){_dgbd .ProcTd (0,-_dgbd .Tl )};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;
);func (_ad *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_ad .Tm =_eb .NewMatrix (a ,b ,c ,d ,e ,f );_ad .Tlm =_ad .Tm .Clone ();};func NewTextFont (font *_ff .PdfFont ,size float64 )(*TextFont ,error ){_baa :=font .FontDescriptor ();if _baa ==nil {return nil ,_f .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");
};_gdf ,_ab :=_cf .GetStream (_baa .FontFile2 );if !_ab {return nil ,_f .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};_bca ,_aa :=_cf .DecodeStream (_gdf );
if _aa !=nil {return nil ,_aa ;};_fbf ,_aa :=_ce .Parse (_bca );if _aa !=nil {return nil ,_aa ;};_gb :=font .FontDescriptor ().FontName .String ();_bda :=len (_gb )> 7&&_gb [6]=='+';if !_fbf .HasCmap ()&&(!_d .Contains (font .Encoder ().String (),"\u0049d\u0065\u006e\u0074\u0069\u0074\u0079-")||!_bda ){return nil ,_f .New ("\u006e\u006f c\u006d\u0061\u0070 \u0061\u006e\u0064\u0020enc\u006fdi\u006e\u0067\u0020\u0069\u0073\u0020\u006eot\u0020\u0069\u0064\u0065\u006e\u0074\u0069t\u0079");
};return &TextFont {Font :font ,Size :size ,_bf :_fbf },nil ;};func (_gae *TextState )Translate (tx ,ty float64 ){_gae .Tm =_gae .Tm .Mult (_eb .TranslationMatrix (tx ,ty ));};func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_deg ,_dba :=_ff .NewPdfFontFromTTFFile (filePath );
if _dba !=nil {return nil ,_dba ;};return NewTextFont (_deg ,size );};const (TextRenderingModeFill TextRenderingMode =iota ;TextRenderingModeStroke ;TextRenderingModeFillStroke ;TextRenderingModeInvisible ;TextRenderingModeFillClip ;TextRenderingModeStrokeClip ;
TextRenderingModeFillStrokeClip ;TextRenderingModeClip ;);func (_afb *TextFont )BytesToCharcodes (data []byte )[]_cg .CharCode {if _afb ._cgcc !=nil {return _afb ._cgcc .BytesToCharcodes (data );};return _afb .Font .BytesToCharcodes (data );};