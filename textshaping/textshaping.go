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

package textshaping ;import (_b "github.com/unidoc/garabic";_be "golang.org/x/text/unicode/bidi";_d "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_c :=_be .Paragraph {};_c .SetString (text );_ac ,_ba :=_c .Order ();if _ba !=nil {return "",_ba ;};for _dc :=0;_dc < _ac .NumRuns ();_dc ++{_f :=_ac .Run (_dc );_cg :=_f .String ();if _f .Direction ()==_be .RightToLeft {var (_e =_b .Shape (_cg );
_g =[]rune (_e );_dg =make ([]rune ,len (_g )););_bec :=0;for _ab :=len (_g )-1;_ab >=0;_ab --{_dg [_bec ]=_g [_ab ];_bec ++;};_cg =string (_dg );text =_d .Replace (text ,_d .TrimSpace (_f .String ()),_cg ,1);};};return text ,nil ;};