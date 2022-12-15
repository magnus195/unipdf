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

package pdfutil ;import (_c "github.com/unidoc/unipdf/v3/common";_gf "github.com/unidoc/unipdf/v3/contentstream";_f "github.com/unidoc/unipdf/v3/contentstream/draw";_g "github.com/unidoc/unipdf/v3/core";_a "github.com/unidoc/unipdf/v3/model";);

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
func NormalizePage (page *_a .PdfPage )error {_ab ,_ad :=page .GetMediaBox ();if _ad !=nil {return _ad ;};_fb ,_ad :=page .GetRotate ();if _ad !=nil {_c .Log .Debug ("\u0045\u0052R\u004f\u0052\u003a\u0020\u0025\u0073\u0020\u002d\u0020\u0069\u0067\u006e\u006f\u0072\u0069\u006e\u0067\u0020\u0061\u006e\u0064\u0020\u0061\u0073\u0073\u0075\u006d\u0069\u006e\u0067\u0020\u006e\u006f\u0020\u0072\u006f\u0074\u0061\u0074\u0069\u006f\u006e\u000a",_ad .Error ());
};_ge :=_fb %360!=0&&_fb %90==0;_ab .Normalize ();_ga ,_df ,_ca ,_e :=_ab .Llx ,_ab .Lly ,_ab .Width (),_ab .Height ();_b :=_ga !=0||_df !=0;if !_ge &&!_b {return nil ;};_fg :=func (_fe ,_dc ,_aa float64 )_f .BoundingBox {return _f .Path {Points :[]_f .Point {_f .NewPoint (0,0).Rotate (_aa ),_f .NewPoint (_fe ,0).Rotate (_aa ),_f .NewPoint (0,_dc ).Rotate (_aa ),_f .NewPoint (_fe ,_dc ).Rotate (_aa )}}.GetBoundingBox ();
};_aac :=_gf .NewContentCreator ();var _bf float64 ;if _ge {_bf =-float64 (_fb );_eb :=_fg (_ca ,_e ,_bf );_aac .Translate ((_eb .Width -_ca )/2+_ca /2,(_eb .Height -_e )/2+_e /2);_aac .RotateDeg (_bf );_aac .Translate (-_ca /2,-_e /2);_ca ,_e =_eb .Width ,_eb .Height ;
};if _b {_aac .Translate (-_ga ,-_df );};_bg :=_aac .Operations ();_da ,_ad :=_g .MakeStream (_bg .Bytes (),_g .NewFlateEncoder ());if _ad !=nil {return _ad ;};_abc :=_g .MakeArray (_da );_abc .Append (page .GetContentStreamObjs ()...);*_ab =_a .PdfRectangle {Urx :_ca ,Ury :_e };
if _bd :=page .CropBox ;_bd !=nil {_bd .Normalize ();_cf ,_dab ,_bc ,_bb :=_bd .Llx -_ga ,_bd .Lly -_df ,_bd .Width (),_bd .Height ();if _ge {_dfb :=_fg (_bc ,_bb ,_bf );_bc ,_bb =_dfb .Width ,_dfb .Height ;};*_bd =_a .PdfRectangle {Llx :_cf ,Lly :_dab ,Urx :_cf +_bc ,Ury :_dab +_bb };
};_c .Log .Debug ("\u0052\u006f\u0074\u0061\u0074\u0065\u003d\u0025\u0066\u00b0\u0020\u004f\u0070\u0073\u003d%\u0071 \u004d\u0065\u0064\u0069\u0061\u0042\u006f\u0078\u003d\u0025\u002e\u0032\u0066",_bf ,_bg ,_ab );page .Contents =_abc ;page .Rotate =nil ;
return nil ;};