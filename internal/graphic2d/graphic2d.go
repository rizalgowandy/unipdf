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

package graphic2d ;import (_fb "image/color";_e "math";);func _ba (_df ,_baa ,_ce ,_ffa ,_ge float64 ,_gd ,_db bool ,_fa ,_ab float64 )(float64 ,float64 ,float64 ,float64 ){if _fcf (_df ,_fa )&&_fcf (_baa ,_ab ){return _df ,_baa ,0.0,0.0;};_bbe ,_dc :=_e .Sincos (_ge );
_eba :=_dc *(_df -_fa )/2.0+_bbe *(_baa -_ab )/2.0;_afg :=-_bbe *(_df -_fa )/2.0+_dc *(_baa -_ab )/2.0;_dd :=_eba *_eba /_ce /_ce +_afg *_afg /_ffa /_ffa ;if _dd > 1.0{_ce *=_e .Sqrt (_dd );_ffa *=_e .Sqrt (_dd );};_ec :=(_ce *_ce *_ffa *_ffa -_ce *_ce *_afg *_afg -_ffa *_ffa *_eba *_eba )/(_ce *_ce *_afg *_afg +_ffa *_ffa *_eba *_eba );
if _ec < 0.0{_ec =0.0;};_ecf :=_e .Sqrt (_ec );if _gd ==_db {_ecf =-_ecf ;};_cbc :=_ecf *_ce *_afg /_ffa ;_fc :=_ecf *-_ffa *_eba /_ce ;_ebad :=_dc *_cbc -_bbe *_fc +(_df +_fa )/2.0;_gg :=_bbe *_cbc +_dc *_fc +(_baa +_ab )/2.0;_fbce :=(_eba -_cbc )/_ce ;
_gdf :=(_afg -_fc )/_ffa ;_gf :=-(_eba +_cbc )/_ce ;_cd :=-(_afg +_fc )/_ffa ;_gba :=_e .Acos (_fbce /_e .Sqrt (_fbce *_fbce +_gdf *_gdf ));if _gdf < 0.0{_gba =-_gba ;};_gba =_cbb (_gba );_bf :=(_fbce *_gf +_gdf *_cd )/_e .Sqrt ((_fbce *_fbce +_gdf *_gdf )*(_gf *_gf +_cd *_cd ));
_bf =_e .Min (1.0,_e .Max (-1.0,_bf ));_gcg :=_e .Acos (_bf );if _fbce *_cd -_gdf *_gf < 0.0{_gcg =-_gcg ;};if !_db &&_gcg > 0.0{_gcg -=2.0*_e .Pi ;}else if _db &&_gcg < 0.0{_gcg +=2.0*_e .Pi ;};return _ebad ,_gg ,_gba ,_gba +_gcg ;};func (_bgb Point )Mul (f float64 )Point {return Point {f *_bgb .X ,f *_bgb .Y }};
func QuadraticToCubicBezier (startX ,startY ,x1 ,y1 ,x ,y float64 )(Point ,Point ){_fba :=Point {X :startX ,Y :startY };_cdd :=Point {X :x1 ,Y :y1 };_ee :=Point {X :x ,Y :y };_ega :=_fba .Interpolate (_cdd ,2.0/3.0);_ggg :=_ee .Interpolate (_cdd ,2.0/3.0);
return _ega ,_ggg ;};func _bcd (_ggc ,_cbee ,_ac float64 ,_ca bool ,_bdb float64 )(float64 ,float64 ){_be ,_ddc :=_e .Sincos (_bdb );_ga ,_cf :=_e .Sincos (_ac );_aeb :=-_ggc *_be *_cf -_cbee *_ddc *_ga ;_dg :=-_ggc *_be *_ga +_cbee *_ddc *_cf ;if !_ca {return -_aeb ,-_dg ;
};return _aeb ,_dg ;};func (_cde Point )Interpolate (q Point ,t float64 )Point {return Point {(1-t )*_cde .X +t *q .X ,(1-t )*_cde .Y +t *q .Y };};type Point struct{X ,Y float64 ;};const _gaf =1e-10;var (Aliceblue =_fb .RGBA {0xf0,0xf8,0xff,0xff};Antiquewhite =_fb .RGBA {0xfa,0xeb,0xd7,0xff};
Aqua =_fb .RGBA {0x00,0xff,0xff,0xff};Aquamarine =_fb .RGBA {0x7f,0xff,0xd4,0xff};Azure =_fb .RGBA {0xf0,0xff,0xff,0xff};Beige =_fb .RGBA {0xf5,0xf5,0xdc,0xff};Bisque =_fb .RGBA {0xff,0xe4,0xc4,0xff};Black =_fb .RGBA {0x00,0x00,0x00,0xff};Blanchedalmond =_fb .RGBA {0xff,0xeb,0xcd,0xff};
Blue =_fb .RGBA {0x00,0x00,0xff,0xff};Blueviolet =_fb .RGBA {0x8a,0x2b,0xe2,0xff};Brown =_fb .RGBA {0xa5,0x2a,0x2a,0xff};Burlywood =_fb .RGBA {0xde,0xb8,0x87,0xff};Cadetblue =_fb .RGBA {0x5f,0x9e,0xa0,0xff};Chartreuse =_fb .RGBA {0x7f,0xff,0x00,0xff};Chocolate =_fb .RGBA {0xd2,0x69,0x1e,0xff};
Coral =_fb .RGBA {0xff,0x7f,0x50,0xff};Cornflowerblue =_fb .RGBA {0x64,0x95,0xed,0xff};Cornsilk =_fb .RGBA {0xff,0xf8,0xdc,0xff};Crimson =_fb .RGBA {0xdc,0x14,0x3c,0xff};Cyan =_fb .RGBA {0x00,0xff,0xff,0xff};Darkblue =_fb .RGBA {0x00,0x00,0x8b,0xff};Darkcyan =_fb .RGBA {0x00,0x8b,0x8b,0xff};
Darkgoldenrod =_fb .RGBA {0xb8,0x86,0x0b,0xff};Darkgray =_fb .RGBA {0xa9,0xa9,0xa9,0xff};Darkgreen =_fb .RGBA {0x00,0x64,0x00,0xff};Darkgrey =_fb .RGBA {0xa9,0xa9,0xa9,0xff};Darkkhaki =_fb .RGBA {0xbd,0xb7,0x6b,0xff};Darkmagenta =_fb .RGBA {0x8b,0x00,0x8b,0xff};
Darkolivegreen =_fb .RGBA {0x55,0x6b,0x2f,0xff};Darkorange =_fb .RGBA {0xff,0x8c,0x00,0xff};Darkorchid =_fb .RGBA {0x99,0x32,0xcc,0xff};Darkred =_fb .RGBA {0x8b,0x00,0x00,0xff};Darksalmon =_fb .RGBA {0xe9,0x96,0x7a,0xff};Darkseagreen =_fb .RGBA {0x8f,0xbc,0x8f,0xff};
Darkslateblue =_fb .RGBA {0x48,0x3d,0x8b,0xff};Darkslategray =_fb .RGBA {0x2f,0x4f,0x4f,0xff};Darkslategrey =_fb .RGBA {0x2f,0x4f,0x4f,0xff};Darkturquoise =_fb .RGBA {0x00,0xce,0xd1,0xff};Darkviolet =_fb .RGBA {0x94,0x00,0xd3,0xff};Deeppink =_fb .RGBA {0xff,0x14,0x93,0xff};
Deepskyblue =_fb .RGBA {0x00,0xbf,0xff,0xff};Dimgray =_fb .RGBA {0x69,0x69,0x69,0xff};Dimgrey =_fb .RGBA {0x69,0x69,0x69,0xff};Dodgerblue =_fb .RGBA {0x1e,0x90,0xff,0xff};Firebrick =_fb .RGBA {0xb2,0x22,0x22,0xff};Floralwhite =_fb .RGBA {0xff,0xfa,0xf0,0xff};
Forestgreen =_fb .RGBA {0x22,0x8b,0x22,0xff};Fuchsia =_fb .RGBA {0xff,0x00,0xff,0xff};Gainsboro =_fb .RGBA {0xdc,0xdc,0xdc,0xff};Ghostwhite =_fb .RGBA {0xf8,0xf8,0xff,0xff};Gold =_fb .RGBA {0xff,0xd7,0x00,0xff};Goldenrod =_fb .RGBA {0xda,0xa5,0x20,0xff};
Gray =_fb .RGBA {0x80,0x80,0x80,0xff};Green =_fb .RGBA {0x00,0x80,0x00,0xff};Greenyellow =_fb .RGBA {0xad,0xff,0x2f,0xff};Grey =_fb .RGBA {0x80,0x80,0x80,0xff};Honeydew =_fb .RGBA {0xf0,0xff,0xf0,0xff};Hotpink =_fb .RGBA {0xff,0x69,0xb4,0xff};Indianred =_fb .RGBA {0xcd,0x5c,0x5c,0xff};
Indigo =_fb .RGBA {0x4b,0x00,0x82,0xff};Ivory =_fb .RGBA {0xff,0xff,0xf0,0xff};Khaki =_fb .RGBA {0xf0,0xe6,0x8c,0xff};Lavender =_fb .RGBA {0xe6,0xe6,0xfa,0xff};Lavenderblush =_fb .RGBA {0xff,0xf0,0xf5,0xff};Lawngreen =_fb .RGBA {0x7c,0xfc,0x00,0xff};Lemonchiffon =_fb .RGBA {0xff,0xfa,0xcd,0xff};
Lightblue =_fb .RGBA {0xad,0xd8,0xe6,0xff};Lightcoral =_fb .RGBA {0xf0,0x80,0x80,0xff};Lightcyan =_fb .RGBA {0xe0,0xff,0xff,0xff};Lightgoldenrodyellow =_fb .RGBA {0xfa,0xfa,0xd2,0xff};Lightgray =_fb .RGBA {0xd3,0xd3,0xd3,0xff};Lightgreen =_fb .RGBA {0x90,0xee,0x90,0xff};
Lightgrey =_fb .RGBA {0xd3,0xd3,0xd3,0xff};Lightpink =_fb .RGBA {0xff,0xb6,0xc1,0xff};Lightsalmon =_fb .RGBA {0xff,0xa0,0x7a,0xff};Lightseagreen =_fb .RGBA {0x20,0xb2,0xaa,0xff};Lightskyblue =_fb .RGBA {0x87,0xce,0xfa,0xff};Lightslategray =_fb .RGBA {0x77,0x88,0x99,0xff};
Lightslategrey =_fb .RGBA {0x77,0x88,0x99,0xff};Lightsteelblue =_fb .RGBA {0xb0,0xc4,0xde,0xff};Lightyellow =_fb .RGBA {0xff,0xff,0xe0,0xff};Lime =_fb .RGBA {0x00,0xff,0x00,0xff};Limegreen =_fb .RGBA {0x32,0xcd,0x32,0xff};Linen =_fb .RGBA {0xfa,0xf0,0xe6,0xff};
Magenta =_fb .RGBA {0xff,0x00,0xff,0xff};Maroon =_fb .RGBA {0x80,0x00,0x00,0xff};Mediumaquamarine =_fb .RGBA {0x66,0xcd,0xaa,0xff};Mediumblue =_fb .RGBA {0x00,0x00,0xcd,0xff};Mediumorchid =_fb .RGBA {0xba,0x55,0xd3,0xff};Mediumpurple =_fb .RGBA {0x93,0x70,0xdb,0xff};
Mediumseagreen =_fb .RGBA {0x3c,0xb3,0x71,0xff};Mediumslateblue =_fb .RGBA {0x7b,0x68,0xee,0xff};Mediumspringgreen =_fb .RGBA {0x00,0xfa,0x9a,0xff};Mediumturquoise =_fb .RGBA {0x48,0xd1,0xcc,0xff};Mediumvioletred =_fb .RGBA {0xc7,0x15,0x85,0xff};Midnightblue =_fb .RGBA {0x19,0x19,0x70,0xff};
Mintcream =_fb .RGBA {0xf5,0xff,0xfa,0xff};Mistyrose =_fb .RGBA {0xff,0xe4,0xe1,0xff};Moccasin =_fb .RGBA {0xff,0xe4,0xb5,0xff};Navajowhite =_fb .RGBA {0xff,0xde,0xad,0xff};Navy =_fb .RGBA {0x00,0x00,0x80,0xff};Oldlace =_fb .RGBA {0xfd,0xf5,0xe6,0xff};
Olive =_fb .RGBA {0x80,0x80,0x00,0xff};Olivedrab =_fb .RGBA {0x6b,0x8e,0x23,0xff};Orange =_fb .RGBA {0xff,0xa5,0x00,0xff};Orangered =_fb .RGBA {0xff,0x45,0x00,0xff};Orchid =_fb .RGBA {0xda,0x70,0xd6,0xff};Palegoldenrod =_fb .RGBA {0xee,0xe8,0xaa,0xff};
Palegreen =_fb .RGBA {0x98,0xfb,0x98,0xff};Paleturquoise =_fb .RGBA {0xaf,0xee,0xee,0xff};Palevioletred =_fb .RGBA {0xdb,0x70,0x93,0xff};Papayawhip =_fb .RGBA {0xff,0xef,0xd5,0xff};Peachpuff =_fb .RGBA {0xff,0xda,0xb9,0xff};Peru =_fb .RGBA {0xcd,0x85,0x3f,0xff};
Pink =_fb .RGBA {0xff,0xc0,0xcb,0xff};Plum =_fb .RGBA {0xdd,0xa0,0xdd,0xff};Powderblue =_fb .RGBA {0xb0,0xe0,0xe6,0xff};Purple =_fb .RGBA {0x80,0x00,0x80,0xff};Red =_fb .RGBA {0xff,0x00,0x00,0xff};Rosybrown =_fb .RGBA {0xbc,0x8f,0x8f,0xff};Royalblue =_fb .RGBA {0x41,0x69,0xe1,0xff};
Saddlebrown =_fb .RGBA {0x8b,0x45,0x13,0xff};Salmon =_fb .RGBA {0xfa,0x80,0x72,0xff};Sandybrown =_fb .RGBA {0xf4,0xa4,0x60,0xff};Seagreen =_fb .RGBA {0x2e,0x8b,0x57,0xff};Seashell =_fb .RGBA {0xff,0xf5,0xee,0xff};Sienna =_fb .RGBA {0xa0,0x52,0x2d,0xff};
Silver =_fb .RGBA {0xc0,0xc0,0xc0,0xff};Skyblue =_fb .RGBA {0x87,0xce,0xeb,0xff};Slateblue =_fb .RGBA {0x6a,0x5a,0xcd,0xff};Slategray =_fb .RGBA {0x70,0x80,0x90,0xff};Slategrey =_fb .RGBA {0x70,0x80,0x90,0xff};Snow =_fb .RGBA {0xff,0xfa,0xfa,0xff};Springgreen =_fb .RGBA {0x00,0xff,0x7f,0xff};
Steelblue =_fb .RGBA {0x46,0x82,0xb4,0xff};Tan =_fb .RGBA {0xd2,0xb4,0x8c,0xff};Teal =_fb .RGBA {0x00,0x80,0x80,0xff};Thistle =_fb .RGBA {0xd8,0xbf,0xd8,0xff};Tomato =_fb .RGBA {0xff,0x63,0x47,0xff};Turquoise =_fb .RGBA {0x40,0xe0,0xd0,0xff};Violet =_fb .RGBA {0xee,0x82,0xee,0xff};
Wheat =_fb .RGBA {0xf5,0xde,0xb3,0xff};White =_fb .RGBA {0xff,0xff,0xff,0xff};Whitesmoke =_fb .RGBA {0xf5,0xf5,0xf5,0xff};Yellow =_fb .RGBA {0xff,0xff,0x00,0xff};Yellowgreen =_fb .RGBA {0x9a,0xcd,0x32,0xff};);func (_da Point )Add (q Point )Point {return Point {_da .X +q .X ,_da .Y +q .Y }};
func _cg (_gb ,_agd ,_fe ,_ebc ,_bcc ,_ffb float64 )(float64 ,float64 ){_cbe ,_agf :=_e .Sincos (_ffb );_bb ,_fd :=_e .Sincos (_fe );_fbc :=_ebc +_gb *_agf *_fd -_agd *_cbe *_bb ;_aed :=_bcc +_gb *_agf *_bb +_agd *_cbe *_fd ;return _fbc ,_aed ;};func _fcf (_bdd ,_ddd float64 )bool {return _e .Abs (_bdd -_ddd )<=_gaf };
func EllipseToCubicBeziers (startX ,startY ,rx ,ry ,rot float64 ,large ,sweep bool ,endX ,endY float64 )[][4]Point {rx =_e .Abs (rx );ry =_e .Abs (ry );if rx < ry {rx ,ry =ry ,rx ;rot +=90.0;};_eb :=_cbb (rot *_e .Pi /180.0);if _e .Pi <=_eb {_eb -=_e .Pi ;
};_d ,_a ,_ff ,_ae :=_ba (startX ,startY ,rx ,ry ,_eb ,large ,sweep ,endX ,endY );_c :=_e .Pi /2.0;_ag :=int (_e .Ceil (_e .Abs (_ae -_ff )/_c ));_c =_e .Abs (_ae -_ff )/float64 (_ag );_af :=_e .Sin (_c )*(_e .Sqrt (4.0+3.0*_e .Pow (_e .Tan (_c /2.0),2.0))-1.0)/3.0;
if !sweep {_c =-_c ;};_b :=Point {X :startX ,Y :startY };_cb ,_ef :=_bcd (rx ,ry ,_eb ,sweep ,_ff );_aa :=Point {X :_cb ,Y :_ef };_eg :=[][4]Point {};for _bg :=1;_bg < _ag +1;_bg ++{_aaf :=_ff +float64 (_bg )*_c ;_de ,_ebg :=_cg (rx ,ry ,_eb ,_d ,_a ,_aaf );
_egb :=Point {X :_de ,Y :_ebg };_g ,_gc :=_bcd (rx ,ry ,_eb ,sweep ,_aaf );_bc :=Point {X :_g ,Y :_gc };_bd :=_b .Add (_aa .Mul (_af ));_bcg :=_egb .Sub (_bc .Mul (_af ));_eg =append (_eg ,[4]Point {_b ,_bd ,_bcg ,_egb });_aa =_bc ;_b =_egb ;};return _eg ;
};var Names =[]string {"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e","\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065","\u0061\u0071\u0075\u0061","\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065","\u0061\u007a\u0075r\u0065","\u0062\u0065\u0069g\u0065","\u0062\u0069\u0073\u0071\u0075\u0065","\u0062\u006c\u0061c\u006b","\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064","\u0062\u006c\u0075\u0065","\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074","\u0062\u0072\u006fw\u006e","\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd","\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e","\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065","\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e","\u0063\u006f\u0072a\u006c","\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065","\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b","\u0063r\u0069\u006d\u0073\u006f\u006e","\u0063\u0079\u0061\u006e","\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065","\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e","\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064","\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079","\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n","\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079","\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi","d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061","\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e","\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065","\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064","\u0064a\u0072\u006b\u0072\u0065\u0064","\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e","\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e","\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065","\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079","\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079","\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065","\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074","\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b","d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065","\u0064i\u006d\u0067\u0072\u0061\u0079","\u0064i\u006d\u0067\u0072\u0065\u0079","\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065","\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k","f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065","f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e","\u0066u\u0063\u0068\u0073\u0069\u0061","\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o","\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065","\u0067\u006f\u006c\u0064","\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd","\u0067\u0072\u0061\u0079","\u0067\u0072\u0065e\u006e","g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077","\u0067\u0072\u0065\u0079","\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077","\u0068o\u0074\u0070\u0069\u006e\u006b","\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d","\u0069\u006e\u0064\u0069\u0067\u006f","\u0069\u0076\u006fr\u0079","\u006b\u0068\u0061k\u0069","\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072","\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068","\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n","\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e","\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e","\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c","\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n","l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077","\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y","\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e","\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y","\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek","l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e","\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e","\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065","\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079","\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079","\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065","l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077","\u006c\u0069\u006d\u0065","\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n","\u006c\u0069\u006ee\u006e","\u006da\u0067\u0065\u006e\u0074\u0061","\u006d\u0061\u0072\u006f\u006f\u006e","\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065","\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065","\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064","\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065","\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e","\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065","\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e","\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065","\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064","\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065","\u006di\u006e\u0074\u0063\u0072\u0065\u0061m","\u006di\u0073\u0074\u0079\u0072\u006f\u0073e","\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e","n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065","\u006e\u0061\u0076\u0079","\u006fl\u0064\u006c\u0061\u0063\u0065","\u006f\u006c\u0069v\u0065","\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b","\u006f\u0072\u0061\u006e\u0067\u0065","\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d","\u006f\u0072\u0063\u0068\u0069\u0064","\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064","\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n","\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065","\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064","\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070","\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f","\u0070\u0065\u0072\u0075","\u0070\u0069\u006e\u006b","\u0070\u006c\u0075\u006d","\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065","\u0070\u0075\u0072\u0070\u006c\u0065","\u0072\u0065\u0064","\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n","\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e","s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e","\u0073\u0061\u006c\u006d\u006f\u006e","\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e","\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e","\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c","\u0073\u0069\u0065\u006e\u006e\u0061","\u0073\u0069\u006c\u0076\u0065\u0072","\u0073k\u0079\u0062\u006c\u0075\u0065","\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e","\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y","\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y","\u0073\u006e\u006f\u0077","s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e","\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e","\u0074\u0061\u006e","\u0074\u0065\u0061\u006c","\u0074h\u0069\u0073\u0074\u006c\u0065","\u0074\u006f\u006d\u0061\u0074\u006f","\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e","\u0076\u0069\u006f\u006c\u0065\u0074","\u0077\u0068\u0065a\u0074","\u0077\u0068\u0069t\u0065","\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065","\u0079\u0065\u006c\u006c\u006f\u0077","y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e"};
func _cbb (_afb float64 )float64 {_afb =_e .Mod (_afb ,2.0*_e .Pi );if _afb < 0.0{_afb +=2.0*_e .Pi ;};return _afb ;};var ColorMap =map[string ]_fb .RGBA {"\u0061l\u0069\u0063\u0065\u0062\u006c\u0075e":_fb .RGBA {0xf0,0xf8,0xff,0xff},"\u0061\u006e\u0074i\u0071\u0075\u0065\u0077\u0068\u0069\u0074\u0065":_fb .RGBA {0xfa,0xeb,0xd7,0xff},"\u0061\u0071\u0075\u0061":_fb .RGBA {0x00,0xff,0xff,0xff},"\u0061\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065":_fb .RGBA {0x7f,0xff,0xd4,0xff},"\u0061\u007a\u0075r\u0065":_fb .RGBA {0xf0,0xff,0xff,0xff},"\u0062\u0065\u0069g\u0065":_fb .RGBA {0xf5,0xf5,0xdc,0xff},"\u0062\u0069\u0073\u0071\u0075\u0065":_fb .RGBA {0xff,0xe4,0xc4,0xff},"\u0062\u006c\u0061c\u006b":_fb .RGBA {0x00,0x00,0x00,0xff},"\u0062\u006c\u0061\u006e\u0063\u0068\u0065\u0064\u0061l\u006d\u006f\u006e\u0064":_fb .RGBA {0xff,0xeb,0xcd,0xff},"\u0062\u006c\u0075\u0065":_fb .RGBA {0x00,0x00,0xff,0xff},"\u0062\u006c\u0075\u0065\u0076\u0069\u006f\u006c\u0065\u0074":_fb .RGBA {0x8a,0x2b,0xe2,0xff},"\u0062\u0072\u006fw\u006e":_fb .RGBA {0xa5,0x2a,0x2a,0xff},"\u0062u\u0072\u006c\u0079\u0077\u006f\u006fd":_fb .RGBA {0xde,0xb8,0x87,0xff},"\u0063a\u0064\u0065\u0074\u0062\u006c\u0075e":_fb .RGBA {0x5f,0x9e,0xa0,0xff},"\u0063\u0068\u0061\u0072\u0074\u0072\u0065\u0075\u0073\u0065":_fb .RGBA {0x7f,0xff,0x00,0xff},"\u0063h\u006f\u0063\u006f\u006c\u0061\u0074e":_fb .RGBA {0xd2,0x69,0x1e,0xff},"\u0063\u006f\u0072a\u006c":_fb .RGBA {0xff,0x7f,0x50,0xff},"\u0063\u006f\u0072\u006e\u0066\u006c\u006f\u0077\u0065r\u0062\u006c\u0075\u0065":_fb .RGBA {0x64,0x95,0xed,0xff},"\u0063\u006f\u0072\u006e\u0073\u0069\u006c\u006b":_fb .RGBA {0xff,0xf8,0xdc,0xff},"\u0063r\u0069\u006d\u0073\u006f\u006e":_fb .RGBA {0xdc,0x14,0x3c,0xff},"\u0063\u0079\u0061\u006e":_fb .RGBA {0x00,0xff,0xff,0xff},"\u0064\u0061\u0072\u006b\u0062\u006c\u0075\u0065":_fb .RGBA {0x00,0x00,0x8b,0xff},"\u0064\u0061\u0072\u006b\u0063\u0079\u0061\u006e":_fb .RGBA {0x00,0x8b,0x8b,0xff},"\u0064\u0061\u0072\u006b\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064":_fb .RGBA {0xb8,0x86,0x0b,0xff},"\u0064\u0061\u0072\u006b\u0067\u0072\u0061\u0079":_fb .RGBA {0xa9,0xa9,0xa9,0xff},"\u0064a\u0072\u006b\u0067\u0072\u0065\u0065n":_fb .RGBA {0x00,0x64,0x00,0xff},"\u0064\u0061\u0072\u006b\u0067\u0072\u0065\u0079":_fb .RGBA {0xa9,0xa9,0xa9,0xff},"\u0064a\u0072\u006b\u006b\u0068\u0061\u006bi":_fb .RGBA {0xbd,0xb7,0x6b,0xff},"d\u0061\u0072\u006b\u006d\u0061\u0067\u0065\u006e\u0074\u0061":_fb .RGBA {0x8b,0x00,0x8b,0xff},"\u0064\u0061\u0072\u006b\u006f\u006c\u0069\u0076\u0065g\u0072\u0065\u0065\u006e":_fb .RGBA {0x55,0x6b,0x2f,0xff},"\u0064\u0061\u0072\u006b\u006f\u0072\u0061\u006e\u0067\u0065":_fb .RGBA {0xff,0x8c,0x00,0xff},"\u0064\u0061\u0072\u006b\u006f\u0072\u0063\u0068\u0069\u0064":_fb .RGBA {0x99,0x32,0xcc,0xff},"\u0064a\u0072\u006b\u0072\u0065\u0064":_fb .RGBA {0x8b,0x00,0x00,0xff},"\u0064\u0061\u0072\u006b\u0073\u0061\u006c\u006d\u006f\u006e":_fb .RGBA {0xe9,0x96,0x7a,0xff},"\u0064\u0061\u0072k\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x8f,0xbc,0x8f,0xff},"\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065":_fb .RGBA {0x48,0x3d,0x8b,0xff},"\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0061\u0079":_fb .RGBA {0x2f,0x4f,0x4f,0xff},"\u0064\u0061\u0072\u006b\u0073\u006c\u0061\u0074\u0065\u0067\u0072\u0065\u0079":_fb .RGBA {0x2f,0x4f,0x4f,0xff},"\u0064\u0061\u0072\u006b\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065":_fb .RGBA {0x00,0xce,0xd1,0xff},"\u0064\u0061\u0072\u006b\u0076\u0069\u006f\u006c\u0065\u0074":_fb .RGBA {0x94,0x00,0xd3,0xff},"\u0064\u0065\u0065\u0070\u0070\u0069\u006e\u006b":_fb .RGBA {0xff,0x14,0x93,0xff},"d\u0065\u0065\u0070\u0073\u006b\u0079\u0062\u006c\u0075\u0065":_fb .RGBA {0x00,0xbf,0xff,0xff},"\u0064i\u006d\u0067\u0072\u0061\u0079":_fb .RGBA {0x69,0x69,0x69,0xff},"\u0064i\u006d\u0067\u0072\u0065\u0079":_fb .RGBA {0x69,0x69,0x69,0xff},"\u0064\u006f\u0064\u0067\u0065\u0072\u0062\u006c\u0075\u0065":_fb .RGBA {0x1e,0x90,0xff,0xff},"\u0066i\u0072\u0065\u0062\u0072\u0069\u0063k":_fb .RGBA {0xb2,0x22,0x22,0xff},"f\u006c\u006f\u0072\u0061\u006c\u0077\u0068\u0069\u0074\u0065":_fb .RGBA {0xff,0xfa,0xf0,0xff},"f\u006f\u0072\u0065\u0073\u0074\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x22,0x8b,0x22,0xff},"\u0066u\u0063\u0068\u0073\u0069\u0061":_fb .RGBA {0xff,0x00,0xff,0xff},"\u0067a\u0069\u006e\u0073\u0062\u006f\u0072o":_fb .RGBA {0xdc,0xdc,0xdc,0xff},"\u0067\u0068\u006f\u0073\u0074\u0077\u0068\u0069\u0074\u0065":_fb .RGBA {0xf8,0xf8,0xff,0xff},"\u0067\u006f\u006c\u0064":_fb .RGBA {0xff,0xd7,0x00,0xff},"\u0067o\u006c\u0064\u0065\u006e\u0072\u006fd":_fb .RGBA {0xda,0xa5,0x20,0xff},"\u0067\u0072\u0061\u0079":_fb .RGBA {0x80,0x80,0x80,0xff},"\u0067\u0072\u0065e\u006e":_fb .RGBA {0x00,0x80,0x00,0xff},"g\u0072\u0065\u0065\u006e\u0079\u0065\u006c\u006c\u006f\u0077":_fb .RGBA {0xad,0xff,0x2f,0xff},"\u0067\u0072\u0065\u0079":_fb .RGBA {0x80,0x80,0x80,0xff},"\u0068\u006f\u006e\u0065\u0079\u0064\u0065\u0077":_fb .RGBA {0xf0,0xff,0xf0,0xff},"\u0068o\u0074\u0070\u0069\u006e\u006b":_fb .RGBA {0xff,0x69,0xb4,0xff},"\u0069n\u0064\u0069\u0061\u006e\u0072\u0065d":_fb .RGBA {0xcd,0x5c,0x5c,0xff},"\u0069\u006e\u0064\u0069\u0067\u006f":_fb .RGBA {0x4b,0x00,0x82,0xff},"\u0069\u0076\u006fr\u0079":_fb .RGBA {0xff,0xff,0xf0,0xff},"\u006b\u0068\u0061k\u0069":_fb .RGBA {0xf0,0xe6,0x8c,0xff},"\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072":_fb .RGBA {0xe6,0xe6,0xfa,0xff},"\u006c\u0061\u0076\u0065\u006e\u0064\u0065\u0072\u0062\u006c\u0075\u0073\u0068":_fb .RGBA {0xff,0xf0,0xf5,0xff},"\u006ca\u0077\u006e\u0067\u0072\u0065\u0065n":_fb .RGBA {0x7c,0xfc,0x00,0xff},"\u006c\u0065\u006do\u006e\u0063\u0068\u0069\u0066\u0066\u006f\u006e":_fb .RGBA {0xff,0xfa,0xcd,0xff},"\u006ci\u0067\u0068\u0074\u0062\u006c\u0075e":_fb .RGBA {0xad,0xd8,0xe6,0xff},"\u006c\u0069\u0067\u0068\u0074\u0063\u006f\u0072\u0061\u006c":_fb .RGBA {0xf0,0x80,0x80,0xff},"\u006ci\u0067\u0068\u0074\u0063\u0079\u0061n":_fb .RGBA {0xe0,0xff,0xff,0xff},"l\u0069g\u0068\u0074\u0067\u006f\u006c\u0064\u0065\u006er\u006f\u0064\u0079\u0065ll\u006f\u0077":_fb .RGBA {0xfa,0xfa,0xd2,0xff},"\u006ci\u0067\u0068\u0074\u0067\u0072\u0061y":_fb .RGBA {0xd3,0xd3,0xd3,0xff},"\u006c\u0069\u0067\u0068\u0074\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x90,0xee,0x90,0xff},"\u006ci\u0067\u0068\u0074\u0067\u0072\u0065y":_fb .RGBA {0xd3,0xd3,0xd3,0xff},"\u006ci\u0067\u0068\u0074\u0070\u0069\u006ek":_fb .RGBA {0xff,0xb6,0xc1,0xff},"l\u0069\u0067\u0068\u0074\u0073\u0061\u006c\u006d\u006f\u006e":_fb .RGBA {0xff,0xa0,0x7a,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x20,0xb2,0xaa,0xff},"\u006c\u0069\u0067h\u0074\u0073\u006b\u0079\u0062\u006c\u0075\u0065":_fb .RGBA {0x87,0xce,0xfa,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0061\u0079":_fb .RGBA {0x77,0x88,0x99,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u006c\u0061\u0074e\u0067\u0072\u0065\u0079":_fb .RGBA {0x77,0x88,0x99,0xff},"\u006c\u0069\u0067\u0068\u0074\u0073\u0074\u0065\u0065l\u0062\u006c\u0075\u0065":_fb .RGBA {0xb0,0xc4,0xde,0xff},"l\u0069\u0067\u0068\u0074\u0079\u0065\u006c\u006c\u006f\u0077":_fb .RGBA {0xff,0xff,0xe0,0xff},"\u006c\u0069\u006d\u0065":_fb .RGBA {0x00,0xff,0x00,0xff},"\u006ci\u006d\u0065\u0067\u0072\u0065\u0065n":_fb .RGBA {0x32,0xcd,0x32,0xff},"\u006c\u0069\u006ee\u006e":_fb .RGBA {0xfa,0xf0,0xe6,0xff},"\u006da\u0067\u0065\u006e\u0074\u0061":_fb .RGBA {0xff,0x00,0xff,0xff},"\u006d\u0061\u0072\u006f\u006f\u006e":_fb .RGBA {0x80,0x00,0x00,0xff},"\u006d\u0065d\u0069\u0075\u006da\u0071\u0075\u0061\u006d\u0061\u0072\u0069\u006e\u0065":_fb .RGBA {0x66,0xcd,0xaa,0xff},"\u006d\u0065\u0064\u0069\u0075\u006d\u0062\u006c\u0075\u0065":_fb .RGBA {0x00,0x00,0xcd,0xff},"\u006d\u0065\u0064i\u0075\u006d\u006f\u0072\u0063\u0068\u0069\u0064":_fb .RGBA {0xba,0x55,0xd3,0xff},"\u006d\u0065\u0064i\u0075\u006d\u0070\u0075\u0072\u0070\u006c\u0065":_fb .RGBA {0x93,0x70,0xdb,0xff},"\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0065\u0061g\u0072\u0065\u0065\u006e":_fb .RGBA {0x3c,0xb3,0x71,0xff},"\u006de\u0064i\u0075\u006d\u0073\u006c\u0061\u0074\u0065\u0062\u006c\u0075\u0065":_fb .RGBA {0x7b,0x68,0xee,0xff},"\u006d\u0065\u0064\u0069\u0075\u006d\u0073\u0070\u0072\u0069\u006e\u0067g\u0072\u0065\u0065\u006e":_fb .RGBA {0x00,0xfa,0x9a,0xff},"\u006de\u0064i\u0075\u006d\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065":_fb .RGBA {0x48,0xd1,0xcc,0xff},"\u006de\u0064i\u0075\u006d\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064":_fb .RGBA {0xc7,0x15,0x85,0xff},"\u006d\u0069\u0064n\u0069\u0067\u0068\u0074\u0062\u006c\u0075\u0065":_fb .RGBA {0x19,0x19,0x70,0xff},"\u006di\u006e\u0074\u0063\u0072\u0065\u0061m":_fb .RGBA {0xf5,0xff,0xfa,0xff},"\u006di\u0073\u0074\u0079\u0072\u006f\u0073e":_fb .RGBA {0xff,0xe4,0xe1,0xff},"\u006d\u006f\u0063\u0063\u0061\u0073\u0069\u006e":_fb .RGBA {0xff,0xe4,0xb5,0xff},"n\u0061\u0076\u0061\u006a\u006f\u0077\u0068\u0069\u0074\u0065":_fb .RGBA {0xff,0xde,0xad,0xff},"\u006e\u0061\u0076\u0079":_fb .RGBA {0x00,0x00,0x80,0xff},"\u006fl\u0064\u006c\u0061\u0063\u0065":_fb .RGBA {0xfd,0xf5,0xe6,0xff},"\u006f\u006c\u0069v\u0065":_fb .RGBA {0x80,0x80,0x00,0xff},"\u006fl\u0069\u0076\u0065\u0064\u0072\u0061b":_fb .RGBA {0x6b,0x8e,0x23,0xff},"\u006f\u0072\u0061\u006e\u0067\u0065":_fb .RGBA {0xff,0xa5,0x00,0xff},"\u006fr\u0061\u006e\u0067\u0065\u0072\u0065d":_fb .RGBA {0xff,0x45,0x00,0xff},"\u006f\u0072\u0063\u0068\u0069\u0064":_fb .RGBA {0xda,0x70,0xd6,0xff},"\u0070\u0061\u006c\u0065\u0067\u006f\u006c\u0064\u0065\u006e\u0072\u006f\u0064":_fb .RGBA {0xee,0xe8,0xaa,0xff},"\u0070a\u006c\u0065\u0067\u0072\u0065\u0065n":_fb .RGBA {0x98,0xfb,0x98,0xff},"\u0070\u0061\u006c\u0065\u0074\u0075\u0072\u0071\u0075\u006f\u0069\u0073\u0065":_fb .RGBA {0xaf,0xee,0xee,0xff},"\u0070\u0061\u006c\u0065\u0076\u0069\u006f\u006c\u0065\u0074\u0072\u0065\u0064":_fb .RGBA {0xdb,0x70,0x93,0xff},"\u0070\u0061\u0070\u0061\u0079\u0061\u0077\u0068\u0069\u0070":_fb .RGBA {0xff,0xef,0xd5,0xff},"\u0070e\u0061\u0063\u0068\u0070\u0075\u0066f":_fb .RGBA {0xff,0xda,0xb9,0xff},"\u0070\u0065\u0072\u0075":_fb .RGBA {0xcd,0x85,0x3f,0xff},"\u0070\u0069\u006e\u006b":_fb .RGBA {0xff,0xc0,0xcb,0xff},"\u0070\u006c\u0075\u006d":_fb .RGBA {0xdd,0xa0,0xdd,0xff},"\u0070\u006f\u0077\u0064\u0065\u0072\u0062\u006c\u0075\u0065":_fb .RGBA {0xb0,0xe0,0xe6,0xff},"\u0070\u0075\u0072\u0070\u006c\u0065":_fb .RGBA {0x80,0x00,0x80,0xff},"\u0072\u0065\u0064":_fb .RGBA {0xff,0x00,0x00,0xff},"\u0072o\u0073\u0079\u0062\u0072\u006f\u0077n":_fb .RGBA {0xbc,0x8f,0x8f,0xff},"\u0072o\u0079\u0061\u006c\u0062\u006c\u0075e":_fb .RGBA {0x41,0x69,0xe1,0xff},"s\u0061\u0064\u0064\u006c\u0065\u0062\u0072\u006f\u0077\u006e":_fb .RGBA {0x8b,0x45,0x13,0xff},"\u0073\u0061\u006c\u006d\u006f\u006e":_fb .RGBA {0xfa,0x80,0x72,0xff},"\u0073\u0061\u006e\u0064\u0079\u0062\u0072\u006f\u0077\u006e":_fb .RGBA {0xf4,0xa4,0x60,0xff},"\u0073\u0065\u0061\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x2e,0x8b,0x57,0xff},"\u0073\u0065\u0061\u0073\u0068\u0065\u006c\u006c":_fb .RGBA {0xff,0xf5,0xee,0xff},"\u0073\u0069\u0065\u006e\u006e\u0061":_fb .RGBA {0xa0,0x52,0x2d,0xff},"\u0073\u0069\u006c\u0076\u0065\u0072":_fb .RGBA {0xc0,0xc0,0xc0,0xff},"\u0073k\u0079\u0062\u006c\u0075\u0065":_fb .RGBA {0x87,0xce,0xeb,0xff},"\u0073l\u0061\u0074\u0065\u0062\u006c\u0075e":_fb .RGBA {0x6a,0x5a,0xcd,0xff},"\u0073l\u0061\u0074\u0065\u0067\u0072\u0061y":_fb .RGBA {0x70,0x80,0x90,0xff},"\u0073l\u0061\u0074\u0065\u0067\u0072\u0065y":_fb .RGBA {0x70,0x80,0x90,0xff},"\u0073\u006e\u006f\u0077":_fb .RGBA {0xff,0xfa,0xfa,0xff},"s\u0070\u0072\u0069\u006e\u0067\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x00,0xff,0x7f,0xff},"\u0073t\u0065\u0065\u006c\u0062\u006c\u0075e":_fb .RGBA {0x46,0x82,0xb4,0xff},"\u0074\u0061\u006e":_fb .RGBA {0xd2,0xb4,0x8c,0xff},"\u0074\u0065\u0061\u006c":_fb .RGBA {0x00,0x80,0x80,0xff},"\u0074h\u0069\u0073\u0074\u006c\u0065":_fb .RGBA {0xd8,0xbf,0xd8,0xff},"\u0074\u006f\u006d\u0061\u0074\u006f":_fb .RGBA {0xff,0x63,0x47,0xff},"\u0074u\u0072\u0071\u0075\u006f\u0069\u0073e":_fb .RGBA {0x40,0xe0,0xd0,0xff},"\u0076\u0069\u006f\u006c\u0065\u0074":_fb .RGBA {0xee,0x82,0xee,0xff},"\u0077\u0068\u0065a\u0074":_fb .RGBA {0xf5,0xde,0xb3,0xff},"\u0077\u0068\u0069t\u0065":_fb .RGBA {0xff,0xff,0xff,0xff},"\u0077\u0068\u0069\u0074\u0065\u0073\u006d\u006f\u006b\u0065":_fb .RGBA {0xf5,0xf5,0xf5,0xff},"\u0079\u0065\u006c\u006c\u006f\u0077":_fb .RGBA {0xff,0xff,0x00,0xff},"y\u0065\u006c\u006c\u006f\u0077\u0067\u0072\u0065\u0065\u006e":_fb .RGBA {0x9a,0xcd,0x32,0xff}};
func (_cc Point )Sub (q Point )Point {return Point {_cc .X -q .X ,_cc .Y -q .Y }};