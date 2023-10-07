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

package pdfutil ;import (_e "github.com/unidoc/unipdf/v3/common";_ea "github.com/unidoc/unipdf/v3/contentstream";_f "github.com/unidoc/unipdf/v3/contentstream/draw";_b "github.com/unidoc/unipdf/v3/core";_d "github.com/unidoc/unipdf/v3/model";);

// NormalizePage performs the following operations on the passed in page:
// - Normalize the page rotation.
//   Rotates the contents of the page according to the Rotate entry, thus
//   flattening the rotation. The Rotate entry of the page is set to nil.
// - Normalize the media box.
//   If the media box of the page is offsetted (Llx != 0 or Lly != 0),
//   the contents of the page are translated to (-Llx, -Lly). After
//   normalization, the media box is updated (Llx and Lly are set to 0 and
//   Urx and Ury are updated accordingly).
// - Normalize the crop box.
//   The crop box of the page is updated based on the previous operations.
// After normalization, the page should look the same if openend using a
// PDF viewer.
// NOTE: This function does not normalize annotations, outlines other parts
// that are not part of the basic geometry and page content streams.
func NormalizePage (page *_d .PdfPage )error {_bf ,_a :=page .GetMediaBox ();if _a !=nil {return _a ;};_cc ,_a :=page .GetRotate ();if _a !=nil {_e .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a",_a .Error ());
};_df :=_cc %360!=0&&_cc %90==0;_bf .Normalize ();_g ,_ee ,_ead ,_cb :=_bf .Llx ,_bf .Lly ,_bf .Width (),_bf .Height ();_eaf :=_g !=0||_ee !=0;if !_df &&!_eaf {return nil ;};_bc :=func (_cf ,_ff ,_bfc float64 )_f .BoundingBox {return _f .Path {Points :[]_f .Point {_f .NewPoint (0,0).Rotate (_bfc ),_f .NewPoint (_cf ,0).Rotate (_bfc ),_f .NewPoint (0,_ff ).Rotate (_bfc ),_f .NewPoint (_cf ,_ff ).Rotate (_bfc )}}.GetBoundingBox ();
};_gg :=_ea .NewContentCreator ();var _fb float64 ;if _df {_fb =-float64 (_cc );_ed :=_bc (_ead ,_cb ,_fb );_gg .Translate ((_ed .Width -_ead )/2+_ead /2,(_ed .Height -_cb )/2+_cb /2);_gg .RotateDeg (_fb );_gg .Translate (-_ead /2,-_cb /2);_ead ,_cb =_ed .Width ,_ed .Height ;
};if _eaf {_gg .Translate (-_g ,-_ee );};_gf :=_gg .Operations ();_cfd ,_a :=_b .MakeStream (_gf .Bytes (),_b .NewFlateEncoder ());if _a !=nil {return _a ;};_ac :=_b .MakeArray (_cfd );_ac .Append (page .GetContentStreamObjs ()...);*_bf =_d .PdfRectangle {Urx :_ead ,Ury :_cb };
if _db :=page .CropBox ;_db !=nil {_db .Normalize ();_eg ,_ba ,_ad ,_gd :=_db .Llx -_g ,_db .Lly -_ee ,_db .Width (),_db .Height ();if _df {_bfa :=_bc (_ad ,_gd ,_fb );_ad ,_gd =_bfa .Width ,_bfa .Height ;};*_db =_d .PdfRectangle {Llx :_eg ,Lly :_ba ,Urx :_eg +_ad ,Ury :_ba +_gd };
};_e .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_fb ,_gf ,_bf );page .Contents =_ac ;page .Rotate =nil ;
return nil ;};