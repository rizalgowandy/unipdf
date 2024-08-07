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

package basic ;import _e "github.com/unidoc/unipdf/v3/internal/jbig2/errors";func (_bd *Stack )Push (v interface{}){_bd .Data =append (_bd .Data ,v )};func (_dfa NumSlice )GetIntSlice ()[]int {_bc :=make ([]int ,len (_dfa ));for _ebb ,_da :=range _dfa {_bc [_ebb ]=int (_da );
};return _bc ;};type IntsMap map[uint64 ][]int ;func Ceil (numerator ,denominator int )int {if numerator %denominator ==0{return numerator /denominator ;};return (numerator /denominator )+1;};func (_ge *NumSlice )AddInt (v int ){*_ge =append (*_ge ,float32 (v ))};
func (_ag IntSlice )Get (index int )(int ,error ){if index > len (_ag )-1{return 0,_e .Errorf ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069\u006e\u0064\u0065x:\u0020\u0025\u0064\u0020\u006f\u0075\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006eg\u0065",index );
};return _ag [index ],nil ;};func NewIntSlice (i int )*IntSlice {_beb :=IntSlice (make ([]int ,i ));return &_beb };func (_c IntsMap )Add (key uint64 ,value int ){_c [key ]=append (_c [key ],value )};func (_aa IntsMap )Delete (key uint64 ){delete (_aa ,key )};
func (_dc NumSlice )Get (i int )(float32 ,error ){if i < 0||i > len (_dc )-1{return 0,_e .Errorf ("\u004e\u0075\u006dS\u006c\u0069\u0063\u0065\u002e\u0047\u0065\u0074","\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};return _dc [i ],nil ;};func (_b IntsMap )GetSlice (key uint64 )([]int ,bool ){_cg ,_be :=_b [key ];if !_be {return nil ,false ;};return _cg ,true ;};func (_ca *Stack )top ()int {return len (_ca .Data )-1};func (_db *IntSlice )Copy ()*IntSlice {_ba :=IntSlice (make ([]int ,len (*_db )));
copy (_ba ,*_db );return &_ba ;};func (_f IntsMap )Get (key uint64 )(int ,bool ){_d ,_a :=_f [key ];if !_a {return 0,false ;};if len (_d )==0{return 0,false ;};return _d [0],true ;};type IntSlice []int ;func (_fgg *Stack )Peek ()(_bg interface{},_gf bool ){return _fgg .peek ()};
type Stack struct{Data []interface{};Aux *Stack ;};func (_dfb *Stack )peek ()(interface{},bool ){_ac :=_dfb .top ();if _ac ==-1{return nil ,false ;};return _dfb .Data [_ac ],true ;};func Sign (v float32 )float32 {if v >=0.0{return 1.0;};return -1.0;};func Max (x ,y int )int {if x > y {return x ;
};return y ;};func (_edg *Stack )Len ()int {return len (_edg .Data )};func (_cb *NumSlice )Add (v float32 ){*_cb =append (*_cb ,v )};func Abs (v int )int {if v > 0{return v ;};return -v ;};func (_fe *IntSlice )Add (v int )error {if _fe ==nil {return _e .Error ("\u0049\u006e\u0074S\u006c\u0069\u0063\u0065\u002e\u0041\u0064\u0064","\u0073\u006c\u0069\u0063\u0065\u0020\u006e\u006f\u0074\u0020\u0064\u0065f\u0069\u006e\u0065\u0064");
};*_fe =append (*_fe ,v );return nil ;};func (_fg IntSlice )Size ()int {return len (_fg )};func (_aaf *Stack )Pop ()(_ad interface{},_ec bool ){_ad ,_ec =_aaf .peek ();if !_ec {return nil ,_ec ;};_aaf .Data =_aaf .Data [:_aaf .top ()];return _ad ,true ;
};func (_df NumSlice )GetInt (i int )(int ,error ){const _eb ="\u0047\u0065\u0074\u0049\u006e\u0074";if i < 0||i > len (_df )-1{return 0,_e .Errorf (_eb ,"\u0069n\u0064\u0065\u0078\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u006fu\u0074\u0020\u006f\u0066\u0020\u0072\u0061\u006e\u0067\u0065",i );
};_ebf :=_df [i ];return int (_ebf +Sign (_ebf )*0.5),nil ;};func Min (x ,y int )int {if x < y {return x ;};return y ;};type NumSlice []float32 ;func NewNumSlice (i int )*NumSlice {_ed :=NumSlice (make ([]float32 ,i ));return &_ed };