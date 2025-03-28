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

package textshaping ;import (_g "github.com/unidoc/garabic";_a "golang.org/x/text/unicode/bidi";_c "strings";);

// ArabicShape returns shaped arabic glyphs string.
func ArabicShape (text string )(string ,error ){_b :=_a .Paragraph {};_b .SetString (text );_gc ,_af :=_b .Order ();if _af !=nil {return "",_af ;};for _ce :=0;_ce < _gc .NumRuns ();_ce ++{_f :=_gc .Run (_ce );_aa :=_f .String ();if _f .Direction ()==_a .RightToLeft {var (_fe =_g .Shape (_aa );
_fef =[]rune (_fe );_ca =make ([]rune ,len (_fef )););_dd :=0;for _ge :=len (_fef )-1;_ge >=0;_ge --{_ca [_dd ]=_fef [_ge ];_dd ++;};_aa =string (_ca );text =_c .Replace (text ,_c .TrimSpace (_f .String ()),_aa ,1);};};return text ,nil ;};