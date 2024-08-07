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

package security ;import (_ba "bytes";_e "crypto/aes";_ff "crypto/cipher";_c "crypto/md5";_ed "crypto/rand";_gg "crypto/rc4";_fg "crypto/sha256";_dc "crypto/sha512";_df "encoding/binary";_g "errors";_bac "fmt";_cf "github.com/unidoc/unipdf/v3/common";_b "hash";
_d "io";_a "math";);func (stdHandlerR4 )paddedPass (_fd []byte )[]byte {_ggc :=make ([]byte ,32);_gd :=copy (_ggc ,_fd );for ;_gd < 32;_gd ++{_ggc [_gd ]=_gfg [_gd -len (_fd )];};return _ggc ;};type ecb struct{_edd _ff .Block ;_gc int ;};var _ StdHandler =stdHandlerR4 {};
func (_ceb stdHandlerR4 )alg6 (_bed *StdEncryptDict ,_dcb []byte )([]byte ,error ){var (_gca []byte ;_ffe error ;);_fde :=_ceb .alg2 (_bed ,_dcb );if _bed .R ==2{_gca ,_ffe =_ceb .alg4 (_fde ,_dcb );}else if _bed .R >=3{_gca ,_ffe =_ceb .alg5 (_fde ,_dcb );
}else {return nil ,_g .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _ffe !=nil {return nil ,_ffe ;};_cf .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_gca ),string (_bed .U ));
_cd :=_gca ;_efdgd :=_bed .U ;if _bed .R >=3{if len (_cd )> 16{_cd =_cd [0:16];};if len (_efdgd )> 16{_efdgd =_efdgd [0:16];};};if !_ba .Equal (_cd ,_efdgd ){return nil ,nil ;};return _fde ,nil ;};

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_feff stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_ceg :=_feff .alg3 (d .R ,upass ,opass );if _ceg !=nil {_cf .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_ceg );
return nil ,_ceg ;};d .O =O ;_cf .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_ca :=_feff .alg2 (d ,upass );U ,_ceg :=_feff .alg5 (_ca ,upass );if _ceg !=nil {_cf .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_ceg );
return nil ,_ceg ;};d .U =U ;_cf .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _ca ,nil ;};

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_fec stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_fgee :=make ([]byte ,32);if _ ,_dcad :=_d .ReadFull (_ed .Reader ,_fgee );_dcad !=nil {return nil ,_dcad ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;
d .Perms =nil ;if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _bfe :=_fec .alg8 (d ,_fgee ,upass );_bfe !=nil {return nil ,_bfe ;};if _fgea :=_fec .alg9 (d ,_fgee ,opass );_fgea !=nil {return nil ,_fgea ;};if d .R ==5{return _fgee ,nil ;
};if _ffee :=_fec .alg10 (d ,_fgee );_ffee !=nil {return nil ,_ffee ;};return _fgee ,nil ;};func (_gdf stdHandlerR4 )alg3 (R int ,_cb ,_cg []byte )([]byte ,error ){var _cbb []byte ;if len (_cg )> 0{_cbb =_gdf .alg3Key (R ,_cg );}else {_cbb =_gdf .alg3Key (R ,_cb );
};_ae ,_efd :=_gg .NewCipher (_cbb );if _efd !=nil {return nil ,_g .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_edb :=_gdf .paddedPass (_cb );_gda :=make ([]byte ,len (_edb ));_ae .XORKeyStream (_gda ,_edb );
if R >=3{_daf :=make ([]byte ,len (_cbb ));for _fdf :=0;_fdf < 19;_fdf ++{for _fbf :=0;_fbf < len (_cbb );_fbf ++{_daf [_fbf ]=_cbb [_fbf ]^byte (_fdf +1);};_dafa ,_cbd :=_gg .NewCipher (_daf );if _cbd !=nil {return nil ,_g .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_dafa .XORKeyStream (_gda ,_gda );};};return _gda ,nil ;};

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};func _gf (_bg _ff .Block )*ecb {return &ecb {_edd :_bg ,_gc :_bg .BlockSize ()}};func (_de *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_de ._gc !=0{_cf .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_cf .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_de ._edd .Decrypt (dst ,src [:_de ._gc ]);src =src [_de ._gc :];dst =dst [_de ._gc :];};};

// Authenticate implements StdHandler interface.
func (_gdff stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _gdff .alg2a (d ,pass );};func (_bb errInvalidField )Error ()string {return _bac .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_bb .Func ,_bb .Field ,_bb .Exp ,_bb .Got );
};func (_fbfa stdHandlerR4 )alg7 (_efa *StdEncryptDict ,_bae []byte )([]byte ,error ){_abd :=_fbfa .alg3Key (_efa .R ,_bae );_gag :=make ([]byte ,len (_efa .O ));if _efa .R ==2{_gb ,_eff :=_gg .NewCipher (_abd );if _eff !=nil {return nil ,_g .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_gb .XORKeyStream (_gag ,_efa .O );}else if _efa .R >=3{_bab :=append ([]byte {},_efa .O ...);for _abg :=0;_abg < 20;_abg ++{_ecb :=append ([]byte {},_abd ...);for _eaf :=0;_eaf < len (_abd );_eaf ++{_ecb [_eaf ]^=byte (19-_abg );};_gbg ,_aba :=_gg .NewCipher (_ecb );
if _aba !=nil {return nil ,_g .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_gbg .XORKeyStream (_gag ,_bab );_bab =append ([]byte {},_gag ...);};}else {return nil ,_g .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");
};_acg ,_abga :=_fbfa .alg6 (_efa ,_gag );if _abga !=nil {return nil ,nil ;};return _acg ,nil ;};const _gfg ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";func _fce (_dag []byte )([]byte ,error ){_ced :=_fg .New ();
_ced .Write (_dag );return _ced .Sum (nil ),nil ;};

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;type ecbDecrypter ecb ;func _afb (_bfb []byte ,_gcab int ){_gef :=_gcab ;for _gef < len (_bfb ){copy (_bfb [_gef :],_bfb [:_gef ]);_gef *=2;};};func (_cbf stdHandlerR4 )alg5 (_aef []byte ,_ab []byte )([]byte ,error ){_fdfd :=_c .New ();
_fdfd .Write ([]byte (_gfg ));_fdfd .Write ([]byte (_cbf .ID0 ));_aa :=_fdfd .Sum (nil );_cf .Log .Trace ("\u0061\u006c\u0067\u0035");_cf .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_aef );_cf .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_cbf .ID0 );
if len (_aa )!=16{return nil ,_g .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");};_gae ,_aaa :=_gg .NewCipher (_aef );if _aaa !=nil {return nil ,_g .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_fge :=make ([]byte ,16);_gae .XORKeyStream (_fge ,_aa );_dfb :=make ([]byte ,len (_aef ));for _ag :=0;_ag < 19;_ag ++{for _dgg :=0;_dgg < len (_aef );_dgg ++{_dfb [_dgg ]=_aef [_dgg ]^byte (_ag +1);};_gae ,_aaa =_gg .NewCipher (_dfb );if _aaa !=nil {return nil ,_g .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_gae .XORKeyStream (_fge ,_fge );_cf .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_ag ,_dfb );_cf .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_ag ,_fge );
};_ge :=make ([]byte ,32);for _cfe :=0;_cfe < 16;_cfe ++{_ge [_cfe ]=_fge [_cfe ];};_ ,_aaa =_ed .Read (_ge [16:32]);if _aaa !=nil {return nil ,_g .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _ge ,nil ;};func _ef (_db _ff .Block )_ff .BlockMode {return (*ecbEncrypter )(_gf (_db ))};type errInvalidField struct{Func string ;Field string ;Exp int ;Got int ;};func (_fbb stdHandlerR6 )alg2a (_adg *StdEncryptDict ,_ace []byte )([]byte ,Permissions ,error ){if _fad :=_ac ("\u0061\u006c\u00672\u0061","\u004f",48,_adg .O );
_fad !=nil {return nil ,0,_fad ;};if _dbc :=_ac ("\u0061\u006c\u00672\u0061","\u0055",48,_adg .U );_dbc !=nil {return nil ,0,_dbc ;};if len (_ace )> 127{_ace =_ace [:127];};_eag ,_fdfa :=_fbb .alg12 (_adg ,_ace );if _fdfa !=nil {return nil ,0,_fdfa ;};
var (_ffec []byte ;_eg []byte ;_ddd []byte ;);var _fc Permissions ;if len (_eag )!=0{_fc =PermOwner ;_gdd :=make ([]byte ,len (_ace )+8+48);_baa :=copy (_gdd ,_ace );_baa +=copy (_gdd [_baa :],_adg .O [40:48]);copy (_gdd [_baa :],_adg .U [0:48]);_ffec =_gdd ;
_eg =_adg .OE ;_ddd =_adg .U [0:48];}else {_eag ,_fdfa =_fbb .alg11 (_adg ,_ace );if _fdfa ==nil &&len (_eag )==0{_eag ,_fdfa =_fbb .alg11 (_adg ,[]byte (""));};if _fdfa !=nil {return nil ,0,_fdfa ;}else if len (_eag )==0{return nil ,0,nil ;};_fc =_adg .P ;
_edf :=make ([]byte ,len (_ace )+8);_bbd :=copy (_edf ,_ace );copy (_edf [_bbd :],_adg .U [40:48]);_ffec =_edf ;_eg =_adg .UE ;_ddd =nil ;};if _deb :=_ac ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_eg );_deb !=nil {return nil ,0,_deb ;};_eg =_eg [:32];
_gaec ,_fdfa :=_fbb .alg2b (_adg .R ,_ffec ,_ace ,_ddd );if _fdfa !=nil {return nil ,0,_fdfa ;};_cfb ,_fdfa :=_e .NewCipher (_gaec [:32]);if _fdfa !=nil {return nil ,0,_fdfa ;};_bbg :=make ([]byte ,_e .BlockSize );_dbb :=_ff .NewCBCDecrypter (_cfb ,_bbg );
_ggd :=make ([]byte ,32);_dbb .CryptBlocks (_ggd ,_eg );if _adg .R ==5{return _ggd ,_fc ,nil ;};_fdfa =_fbb .alg13 (_adg ,_ggd );if _fdfa !=nil {return nil ,0,_fdfa ;};return _ggd ,_fc ,nil ;};func (_aca stdHandlerR4 )alg3Key (R int ,_bce []byte )[]byte {_ec :=_c .New ();
_dg :=_aca .paddedPass (_bce );_ec .Write (_dg );if R >=3{for _da :=0;_da < 50;_da ++{_bgdb :=_ec .Sum (nil );_ec =_c .New ();_ec .Write (_bgdb );};};_gcg :=_ec .Sum (nil );if R ==2{_gcg =_gcg [0:5];}else {_gcg =_gcg [0:_aca .Length /8];};return _gcg ;
};

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};type ecbEncrypter ecb ;type stdHandlerR6 struct{};

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_gfd *StdEncryptDict ,_bf ,_be []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_efe *StdEncryptDict ,_dd []byte )([]byte ,Permissions ,error );};var _ StdHandler =stdHandlerR6 {};func (_add stdHandlerR6 )alg13 (_fbbc *StdEncryptDict ,_gcc []byte )error {if _cga :=_ac ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_gcc );
_cga !=nil {return _cga ;};if _ege :=_ac ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_fbbc .Perms );_ege !=nil {return _ege ;};_dccc :=make ([]byte ,16);copy (_dccc ,_fbbc .Perms [:16]);_egb ,_gfdc :=_e .NewCipher (_gcc [:32]);if _gfdc !=nil {return _gfdc ;
};_ccg :=_ga (_egb );_ccg .CryptBlocks (_dccc ,_dccc );if !_ba .Equal (_dccc [9:12],[]byte ("\u0061\u0064\u0062")){return _g .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_ccb :=Permissions (_df .LittleEndian .Uint32 (_dccc [0:4]));if _ccb !=_fbbc .P {return _g .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _fdec bool ;if _dccc [8]=='T'{_fdec =true ;}else if _dccc [8]=='F'{_fdec =false ;}else {return _g .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _fdec !=_fbbc .EncryptMetadata {return _g .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};

// Authenticate implements StdHandler interface.
func (_fbec stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_cf .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_eac ,_aaag :=_fbec .alg7 (d ,pass );if _aaag !=nil {return nil ,0,_aaag ;};if _eac !=nil {_cf .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _eac ,PermOwner ,nil ;
};_cf .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_eac ,_aaag =_fbec .alg6 (d ,pass );if _aaag !=nil {return nil ,0,_aaag ;
};if _eac !=nil {_cf .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _eac ,d .P ,nil ;};return nil ,0,nil ;};func (_bc *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_bc ._gc !=0{_cf .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_cf .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_bc ._edd .Encrypt (dst ,src [:_bc ._gc ]);src =src [_bc ._gc :];dst =dst [_bc ._gc :];};};type stdHandlerR4 struct{Length int ;ID0 string ;};func (_bgc stdHandlerR6 )alg12 (_afg *StdEncryptDict ,_gee []byte )([]byte ,error ){if _cca :=_ac ("\u0061\u006c\u00671\u0032","\u0055",48,_afg .U );
_cca !=nil {return nil ,_cca ;};if _ccf :=_ac ("\u0061\u006c\u00671\u0032","\u004f",48,_afg .O );_ccf !=nil {return nil ,_ccf ;};_gge :=make ([]byte ,len (_gee )+8+48);_fee :=copy (_gge ,_gee );_fee +=copy (_gge [_fee :],_afg .O [32:40]);_fee +=copy (_gge [_fee :],_afg .U [0:48]);
_bec ,_gfge :=_bgc .alg2b (_afg .R ,_gge ,_gee ,_afg .U [0:48]);if _gfge !=nil {return nil ,_gfge ;};_bec =_bec [:32];if !_ba .Equal (_bec ,_afg .O [:32]){return nil ,nil ;};return _bec ,nil ;};func (_cab stdHandlerR6 )alg2b (R int ,_gfda ,_abb ,_caf []byte )([]byte ,error ){if R ==5{return _fce (_gfda );
};return _afe (_gfda ,_abb ,_caf );};func _ac (_cff ,_gaa string ,_dca int ,_ce []byte )error {if len (_ce )< _dca {return errInvalidField {Func :_cff ,Field :_gaa ,Exp :_dca ,Got :len (_ce )};};return nil ;};func (_dec stdHandlerR6 )alg11 (_cdd *StdEncryptDict ,_ebgf []byte )([]byte ,error ){if _fba :=_ac ("\u0061\u006c\u00671\u0031","\u0055",48,_cdd .U );
_fba !=nil {return nil ,_fba ;};_fda :=make ([]byte ,len (_ebgf )+8);_bfg :=copy (_fda ,_ebgf );_bfg +=copy (_fda [_bfg :],_cdd .U [32:40]);_cfda ,_fffd :=_dec .alg2b (_cdd .R ,_fda ,_ebgf ,nil );if _fffd !=nil {return nil ,_fffd ;};_cfda =_cfda [:32];
if !_ba .Equal (_cfda ,_cdd .U [:32]){return nil ,nil ;};return _cfda ,nil ;};func (_fe *ecbEncrypter )BlockSize ()int {return _fe ._gc };func (_bgd *ecbDecrypter )BlockSize ()int {return _bgd ._gc };func (_ebgc stdHandlerR6 )alg9 (_gcac *StdEncryptDict ,_cgea []byte ,_baf []byte )error {if _aaea :=_ac ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_cgea );
_aaea !=nil {return _aaea ;};if _cbdg :=_ac ("\u0061\u006c\u0067\u0039","\u0055",48,_gcac .U );_cbdg !=nil {return _cbdg ;};var _dcf [16]byte ;if _ ,_cfd :=_d .ReadFull (_ed .Reader ,_dcf [:]);_cfd !=nil {return _cfd ;};_bfc :=_dcf [0:8];_abgf :=_dcf [8:16];
_bfa :=_gcac .U [:48];_ada :=make ([]byte ,len (_baf )+len (_bfc )+len (_bfa ));_cgb :=copy (_ada ,_baf );_cgb +=copy (_ada [_cgb :],_bfc );_cgb +=copy (_ada [_cgb :],_bfa );_dcc ,_ecg :=_ebgc .alg2b (_gcac .R ,_ada ,_baf ,_bfa );if _ecg !=nil {return _ecg ;
};O :=make ([]byte ,len (_dcc )+len (_bfc )+len (_abgf ));_cgb =copy (O ,_dcc [:32]);_cgb +=copy (O [_cgb :],_bfc );_cgb +=copy (O [_cgb :],_abgf );_gcac .O =O ;_cgb =len (_baf );_cgb +=copy (_ada [_cgb :],_abgf );_dcc ,_ecg =_ebgc .alg2b (_gcac .R ,_ada ,_baf ,_bfa );
if _ecg !=nil {return _ecg ;};_beg ,_ecg :=_gbgc (_dcc [:32]);if _ecg !=nil {return _ecg ;};_afa :=make ([]byte ,_e .BlockSize );_begd :=_ff .NewCBCEncrypter (_beg ,_afa );OE :=make ([]byte ,32);_begd .CryptBlocks (OE ,_cgea [:32]);_gcac .OE =OE ;return nil ;
};

// Allowed checks if a set of permissions can be granted.
func (_acd Permissions )Allowed (p2 Permissions )bool {return _acd &p2 ==p2 };func (_ea stdHandlerR4 )alg2 (_af *StdEncryptDict ,_fgf []byte )[]byte {_cf .Log .Trace ("\u0061\u006c\u0067\u0032");_bd :=_ea .paddedPass (_fgf );_aff :=_c .New ();_aff .Write (_bd );
_aff .Write (_af .O );var _def [4]byte ;_df .LittleEndian .PutUint32 (_def [:],uint32 (_af .P ));_aff .Write (_def [:]);_cf .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_def );_aff .Write ([]byte (_ea .ID0 ));_cf .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_af .R ,_af .EncryptMetadata );
if (_af .R >=4)&&!_af .EncryptMetadata {_aff .Write ([]byte {0xff,0xff,0xff,0xff});};_fbe :=_aff .Sum (nil );if _af .R >=3{_aff =_c .New ();for _ggcf :=0;_ggcf < 50;_ggcf ++{_aff .Reset ();_aff .Write (_fbe [0:_ea .Length /8]);_fbe =_aff .Sum (nil );};
};if _af .R >=3{return _fbe [0:_ea .Length /8];};return _fbe [0:5];};func _afe (_gac ,_gfgd ,_ee []byte )([]byte ,error ){var (_egc ,_cba ,_dgd _b .Hash ;);_egc =_fg .New ();_egcg :=make ([]byte ,64);_gfa :=_egc ;_gfa .Write (_gac );K :=_gfa .Sum (_egcg [:0]);
_dbcg :=make ([]byte ,64*(127+64+48));_ecee :=func (_dce int )([]byte ,error ){_aaf :=len (_gfgd )+len (K )+len (_ee );_eef :=_dbcg [:_aaf ];_gaf :=copy (_eef ,_gfgd );_gaf +=copy (_eef [_gaf :],K [:]);_gaf +=copy (_eef [_gaf :],_ee );if _gaf !=_aaf {_cf .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_g .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_dbcg [:_aaf *64];_afb (K1 ,_aaf );_ede ,_eb :=_gbgc (K [0:16]);if _eb !=nil {return nil ,_eb ;};_dea :=_ff .NewCBCEncrypter (_ede ,K [16:32]);_dea .CryptBlocks (K1 ,K1 );
E :=K1 ;_cgg :=0;for _fff :=0;_fff < 16;_fff ++{_cgg +=int (E [_fff ]%3);};var _bcce _b .Hash ;switch _cgg %3{case 0:_bcce =_egc ;case 1:if _cba ==nil {_cba =_dc .New384 ();};_bcce =_cba ;case 2:if _dgd ==nil {_dgd =_dc .New ();};_bcce =_dgd ;};_bcce .Reset ();
_bcce .Write (E );K =_bcce .Sum (_egcg [:0]);return E ,nil ;};for _fac :=0;;{E ,_ggcd :=_ecee (_fac );if _ggcd !=nil {return nil ,_ggcd ;};_edec :=E [len (E )-1];_fac ++;if _fac >=64&&_edec <=uint8 (_fac -32){break ;};};return K [:32],nil ;};func (_aac stdHandlerR6 )alg10 (_aag *StdEncryptDict ,_gbb []byte )error {if _ead :=_ac ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_gbb );
_ead !=nil {return _ead ;};_efb :=uint64 (uint32 (_aag .P ))|(_a .MaxUint32 <<32);Perms :=make ([]byte ,16);_df .LittleEndian .PutUint64 (Perms [:8],_efb );if _aag .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");
if _ ,_dfbe :=_d .ReadFull (_ed .Reader ,Perms [12:16]);_dfbe !=nil {return _dfbe ;};_gcacd ,_adf :=_gbgc (_gbb [:32]);if _adf !=nil {return _adf ;};_eda :=_ef (_gcacd );_eda .CryptBlocks (Perms ,Perms );_aag .Perms =Perms [:16];return nil ;};func (_cc stdHandlerR6 )alg8 (_aefb *StdEncryptDict ,_gfdg []byte ,_aad []byte )error {if _abbc :=_ac ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_gfdg );
_abbc !=nil {return _abbc ;};var _abgc [16]byte ;if _ ,_bccc :=_d .ReadFull (_ed .Reader ,_abgc [:]);_bccc !=nil {return _bccc ;};_fcg :=_abgc [0:8];_cfa :=_abgc [8:16];_aae :=make ([]byte ,len (_aad )+len (_fcg ));_adb :=copy (_aae ,_aad );copy (_aae [_adb :],_fcg );
_ggb ,_bff :=_cc .alg2b (_aefb .R ,_aae ,_aad ,nil );if _bff !=nil {return _bff ;};U :=make ([]byte ,len (_ggb )+len (_fcg )+len (_cfa ));_adb =copy (U ,_ggb [:32]);_adb +=copy (U [_adb :],_fcg );copy (U [_adb :],_cfa );_aefb .U =U ;_adb =len (_aad );copy (_aae [_adb :],_cfa );
_ggb ,_bff =_cc .alg2b (_aefb .R ,_aae ,_aad ,nil );if _bff !=nil {return _bff ;};_ebg ,_bff :=_gbgc (_ggb [:32]);if _bff !=nil {return _bff ;};_dbg :=make ([]byte ,_e .BlockSize );_dfe :=_ff .NewCBCEncrypter (_ebg ,_dbg );UE :=make ([]byte ,32);_dfe .CryptBlocks (UE ,_gfdg [:32]);
_aefb .UE =UE ;return nil ;};const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e"););const (PermOwner =Permissions (_a .MaxUint32 );PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);
PermExtractGraphics =Permissions (1<<4);PermAnnotate =Permissions (1<<5);PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11););

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;func _gbgc (_ece []byte )(_ff .Block ,error ){_gde ,_cge :=_e .NewCipher (_ece );if _cge !=nil {_cf .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_cge );
return nil ,_cge ;};return _gde ,nil ;};func (_ecf stdHandlerR4 )alg4 (_fbg []byte ,_feg []byte )([]byte ,error ){_bcee ,_fa :=_gg .NewCipher (_fbg );if _fa !=nil {return nil ,_g .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_fef :=[]byte (_gfg );_bcc :=make ([]byte ,len (_fef ));_bcee .XORKeyStream (_bcc ,_fef );return _bcc ,nil ;};func _ga (_fb _ff .Block )_ff .BlockMode {return (*ecbDecrypter )(_gf (_fb ))};