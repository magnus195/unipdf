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

package textshaping ;import (_b "github.com/unidoc/garabic";_g "golang.org/x/text/unicode/bidi";_cg "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_e :=_g .Paragraph {};_e .SetString (text );_d ,_f :=_e .Order ();if _f !=nil {return "",_f ;};for _cf :=0;_cf < _d .NumRuns ();_cf ++{_a :=_d .Run (_cf );_ea :=_a .String ();if _a .Direction ()==_g .RightToLeft {var (_cgd =_b .Shape (_ea );
_fg =[]rune (_cgd );_cfc =make ([]rune ,len (_fg )););_gg :=0;for _da :=len (_fg )-1;_da >=0;_da --{_cfc [_gg ]=_fg [_da ];_gg ++;};_ea =string (_cfc );text =_cg .Replace (text ,_cg .TrimSpace (_a .String ()),_ea ,1);};};return text ,nil ;};