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

package textshaping ;import (_f "github.com/unidoc/garabic";_d "golang.org/x/text/unicode/bidi";_c "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_a :=_d .Paragraph {};_a .SetString (text );_ff ,_gg :=_a .Order ();if _gg !=nil {return "",_gg ;};for _ac :=0;_ac < _ff .NumRuns ();_ac ++{_dg :=_ff .Run (_ac );_fc :=_dg .String ();if _dg .Direction ()==_d .RightToLeft {var (_ad =_f .Shape (_fc );
_ca =[]rune (_ad );_e =make ([]rune ,len (_ca )););_fg :=0;for _gb :=len (_ca )-1;_gb >=0;_gb --{_e [_fg ]=_ca [_gb ];_fg ++;};_fc =string (_e );text =_c .Replace (text ,_c .TrimSpace (_dg .String ()),_fc ,1);};};return text ,nil ;};