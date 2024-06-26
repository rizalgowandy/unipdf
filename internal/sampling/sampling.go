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

package sampling ;import (_g "github.com/unidoc/unipdf/v3/internal/bitwise";_de "github.com/unidoc/unipdf/v3/internal/imageutil";_df "io";);func (_fa *Writer )WriteSample (sample uint32 )error {if _ ,_ada :=_fa ._eea .WriteBits (uint64 (sample ),_fa ._bd .BitsPerComponent );
_ada !=nil {return _ada ;};_fa ._eb --;if _fa ._eb ==0{_fa ._eb =_fa ._bd .ColorComponents ;_fa ._eee ++;};if _fa ._eee ==_fa ._bd .Width {if _fa ._gec {_fa ._eea .FinishByte ();};_fa ._eee =0;};return nil ;};func NewWriter (img _de .ImageBase )*Writer {return &Writer {_eea :_g .NewWriterMSB (img .Data ),_bd :img ,_eb :img .ColorComponents ,_gec :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };
};type SampleReader interface{ReadSample ()(uint32 ,error );ReadSamples (_a []uint32 )error ;};func (_ac *Reader )ReadSamples (samples []uint32 )(_ad error ){for _dfc :=0;_dfc < len (samples );_dfc ++{samples [_dfc ],_ad =_ac .ReadSample ();if _ad !=nil {return _ad ;
};};return nil ;};func (_afc *Writer )WriteSamples (samples []uint32 )error {for _bc :=0;_bc < len (samples );_bc ++{if _bgf :=_afc .WriteSample (samples [_bc ]);_bgf !=nil {return _bgf ;};};return nil ;};type SampleWriter interface{WriteSample (_ca uint32 )error ;
WriteSamples (_dea []uint32 )error ;};func ResampleBytes (data []byte ,bitsPerSample int )[]uint32 {var _bg []uint32 ;_fb :=bitsPerSample ;var _cc uint32 ;var _dc byte ;_dfe :=0;_ba :=0;_e :=0;for _e < len (data ){if _dfe > 0{_aa :=_dfe ;if _fb < _aa {_aa =_fb ;
};_cc =(_cc <<uint (_aa ))|uint32 (_dc >>uint (8-_aa ));_dfe -=_aa ;if _dfe > 0{_dc =_dc <<uint (_aa );}else {_dc =0;};_fb -=_aa ;if _fb ==0{_bg =append (_bg ,_cc );_fb =bitsPerSample ;_cc =0;_ba ++;};}else {_da :=data [_e ];_e ++;_gca :=8;if _fb < _gca {_gca =_fb ;
};_dfe =8-_gca ;_cc =(_cc <<uint (_gca ))|uint32 (_da >>uint (_dfe ));if _gca < 8{_dc =_da <<uint (_gca );};_fb -=_gca ;if _fb ==0{_bg =append (_bg ,_cc );_fb =bitsPerSample ;_cc =0;_ba ++;};};};for _dfe >=bitsPerSample {_cf :=_dfe ;if _fb < _cf {_cf =_fb ;
};_cc =(_cc <<uint (_cf ))|uint32 (_dc >>uint (8-_cf ));_dfe -=_cf ;if _dfe > 0{_dc =_dc <<uint (_cf );}else {_dc =0;};_fb -=_cf ;if _fb ==0{_bg =append (_bg ,_cc );_fb =bitsPerSample ;_cc =0;_ba ++;};};return _bg ;};func (_dfa *Reader )ReadSample ()(uint32 ,error ){if _dfa ._cg ==_dfa ._f .Height {return 0,_df .EOF ;
};_gc ,_gf :=_dfa ._b .ReadBits (byte (_dfa ._f .BitsPerComponent ));if _gf !=nil {return 0,_gf ;};_dfa ._cga --;if _dfa ._cga ==0{_dfa ._cga =_dfa ._f .ColorComponents ;_dfa ._c ++;};if _dfa ._c ==_dfa ._f .Width {if _dfa ._dff {_dfa ._b .ConsumeRemainingBits ();
};_dfa ._c =0;_dfa ._cg ++;};return uint32 (_gc ),nil ;};type Writer struct{_bd _de .ImageBase ;_eea *_g .Writer ;_eee ,_eb int ;_gec bool ;};func ResampleUint32 (data []uint32 ,bitsPerInputSample int ,bitsPerOutputSample int )[]uint32 {var _bga []uint32 ;
_ge :=bitsPerOutputSample ;var _ef uint32 ;var _gef uint32 ;_ab :=0;_eg :=0;_gb :=0;for _gb < len (data ){if _ab > 0{_bgc :=_ab ;if _ge < _bgc {_bgc =_ge ;};_ef =(_ef <<uint (_bgc ))|(_gef >>uint (bitsPerInputSample -_bgc ));_ab -=_bgc ;if _ab > 0{_gef =_gef <<uint (_bgc );
}else {_gef =0;};_ge -=_bgc ;if _ge ==0{_bga =append (_bga ,_ef );_ge =bitsPerOutputSample ;_ef =0;_eg ++;};}else {_ee :=data [_gb ];_gb ++;_cgf :=bitsPerInputSample ;if _ge < _cgf {_cgf =_ge ;};_ab =bitsPerInputSample -_cgf ;_ef =(_ef <<uint (_cgf ))|(_ee >>uint (_ab ));
if _cgf < bitsPerInputSample {_gef =_ee <<uint (_cgf );};_ge -=_cgf ;if _ge ==0{_bga =append (_bga ,_ef );_ge =bitsPerOutputSample ;_ef =0;_eg ++;};};};for _ab >=bitsPerOutputSample {_ace :=_ab ;if _ge < _ace {_ace =_ge ;};_ef =(_ef <<uint (_ace ))|(_gef >>uint (bitsPerInputSample -_ace ));
_ab -=_ace ;if _ab > 0{_gef =_gef <<uint (_ace );}else {_gef =0;};_ge -=_ace ;if _ge ==0{_bga =append (_bga ,_ef );_ge =bitsPerOutputSample ;_ef =0;_eg ++;};};if _ge > 0&&_ge < bitsPerOutputSample {_ef <<=uint (_ge );_bga =append (_bga ,_ef );};return _bga ;
};func NewReader (img _de .ImageBase )*Reader {return &Reader {_b :_g .NewReader (img .Data ),_f :img ,_cga :img .ColorComponents ,_dff :img .BytesPerLine *8!=img .ColorComponents *img .BitsPerComponent *img .Width };};type Reader struct{_f _de .ImageBase ;
_b *_g .Reader ;_c ,_cg ,_cga int ;_dff bool ;};