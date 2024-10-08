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

package crypt ;import (_g "crypto/aes";_c "crypto/cipher";_f "crypto/md5";_a "crypto/rand";_eg "crypto/rc4";_b "fmt";_fc "github.com/unidoc/unipdf/v3/common";_d "github.com/unidoc/unipdf/v3/core/security";_bb "io";);func init (){_df ("\u0041\u0045\u0053V\u0032",_fa )};


// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};type filterAESV2 struct{filterAES };

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_dd ,_ee :=_da (FilterDict {});if _ee !=nil {_fc .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ee );
return filterAESV3 {};};return _dd ;};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_aa ,_bc :=_fa (FilterDict {});if _bc !=nil {_fc .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_bc );
return filterAESV2 {};};return _aa ;};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _dbf (objNum ,genNum ,ekey ,true );};

// MakeKey implements Filter interface.
func (_bga filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _dbf (objNum ,genNum ,ekey ,false );};var _ Filter =filterV2 {};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };type filterAES struct{};func init (){_df ("\u0041\u0045\u0053V\u0033",_da )};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_egb ,_ecc :=_eg .NewCipher (okey );if _ecc !=nil {return nil ,_ecc ;};_fc .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );
_egb .XORKeyStream (buf ,buf );_fc .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};type filterAESV3 struct{filterAES };var _ Filter =filterAESV2 {};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_cg ,_fb :=_g .NewCipher (okey );
if _fb !=nil {return nil ,_fb ;};if len (buf )< 16{_fc .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );return buf ,_b .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));
};_gdb :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_fc .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_gdb ),_gdb );_fc .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
return buf ,_b .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));
};_fe :=_c .NewCBCDecrypter (_cg ,_gdb );_fc .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_fc .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );
_fe .CryptBlocks (buf ,buf );_fc .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_fc .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");
return buf ,nil ;};_ga :=int (buf [len (buf )-1]);if _ga > len (buf ){_fc .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_ga ,len (buf ));
return buf ,_b .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_ga ];return buf ,nil ;};func _dbf (_gbc ,_ag uint32 ,_dc []byte ,_dbd bool )([]byte ,error ){_gc :=make ([]byte ,len (_dc )+5);
copy (_gc ,_dc );for _ff :=0;_ff < 3;_ff ++{_bf :=byte ((_gbc >>uint32 (8*_ff ))&0xff);_gc [_ff +len (_dc )]=_bf ;};for _bda :=0;_bda < 2;_bda ++{_eae :=byte ((_ag >>uint32 (8*_bda ))&0xff);_gc [_bda +len (_dc )+3]=_eae ;};if _dbd {_gc =append (_gc ,0x73);
_gc =append (_gc ,0x41);_gc =append (_gc ,0x6C);_gc =append (_gc ,0x54);};_bfa :=_f .New ();_bfa .Write (_gc );_bdb :=_bfa .Sum (nil );if len (_dc )+5< 16{return _bdb [0:len (_dc )+5],nil ;};return _bdb ,nil ;};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};

// Filter is a common interface for crypt filter methods.
type Filter interface{

// Name returns a name of the filter that should be used in CFM field of Encrypt dictionary.
Name ()string ;

// KeyLength returns a length of the encryption key in bytes.
KeyLength ()int ;

// PDFVersion reports the minimal version of PDF document that introduced this filter.
PDFVersion ()[2]int ;

// HandlerVersion reports V and R parameters that should be used for this filter.
HandlerVersion ()(V ,R int );

// MakeKey generates a object encryption key based on file encryption key and object numbers.
// Used only for legacy filters - AESV3 doesn't change the key for each object.
MakeKey (_eag ,_agb uint32 ,_agd []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_egg []byte ,_cc []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_gba []byte ,_ef []byte )([]byte ,error );};

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bd ,_ad :=_g .NewCipher (okey );if _ad !=nil {return nil ,_ad ;
};_fc .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );const _cd =_g .BlockSize ;_db :=_cd -len (buf )%_cd ;for _fg :=0;_fg < _db ;_fg ++{buf =append (buf ,byte (_db ));
};_fc .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_de :=make ([]byte ,_cd +len (buf ));_ba :=_de [:_cd ];if _ ,_gd :=_bb .ReadFull (_a .Reader ,_ba );_gd !=nil {return nil ,_gd ;};
_ed :=_c .NewCBCEncrypter (_bd ,_ba );_ed .CryptBlocks (_de [_cd :],buf );buf =_de ;_fc .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,nil ;};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};type filterV2 struct{_cga int };func init (){_df ("\u0056\u0032",_aeg )};

// PDFVersion implements Filter interface.
func (_daf filterV2 )PDFVersion ()[2]int {return [2]int {}};func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };func _da (_bg FilterDict )(Filter ,error ){if _bg .Length ==256{_fc .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_bg .Length );
_bg .Length /=8;};if _bg .Length !=0&&_bg .Length !=32{return nil ,_b .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_bg .Length );
};return filterAESV3 {},nil ;};type filterFunc func (_dbc FilterDict )(Filter ,error );

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};func _fa (_ca FilterDict )(Filter ,error ){if _ca .Length ==128{_fc .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_ca .Length );
_ca .Length /=8;};if _ca .Length !=0&&_ca .Length !=16{return nil ,_b .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_ca .Length );
};return filterAESV2 {},nil ;};func (filterIdentity )KeyLength ()int {return 0};

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_dg ,_bbfd :=_bbg (d .CFM );if _bbfd !=nil {return nil ,_bbfd ;};_ge ,_bbfd :=_dg (d );if _bbfd !=nil {return nil ,_bbfd ;};return _ge ,nil ;};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};

// KeyLength implements Filter interface.
func (_bbf filterV2 )KeyLength ()int {return _bbf ._cga };func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};type filterIdentity struct{};

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _d .AuthEvent ;Length int ;};func _aeg (_bgb FilterDict )(Filter ,error ){if _bgb .Length %8!=0{return nil ,_b .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_bgb .Length );
};if _bgb .Length < 5||_bgb .Length > 16{if _bgb .Length ==40||_bgb .Length ==64||_bgb .Length ==128{_fc .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_bgb .Length );
_bgb .Length /=8;}else {return nil ,_b .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_bgb .Length );
};};return filterV2 {_cga :_bgb .Length },nil ;};

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};func _df (_ecf string ,_eee filterFunc ){if _ ,_egf :=_cb [_ecf ];_egf {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_cb [_ecf ]=_eee ;};

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_cf ,_gdf :=_eg .NewCipher (okey );if _gdf !=nil {return nil ,_gdf ;};_fc .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );
_cf .XORKeyStream (buf ,buf );_fc .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_ec ,_fcd :=_aeg (FilterDict {Length :length });if _fcd !=nil {_fc .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_fcd );
return filterV2 {_cga :length };};return _ec ;};

// HandlerVersion implements Filter interface.
func (_bdg filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };var (_cb =make (map[string ]filterFunc ););func _bbg (_eaf string )(filterFunc ,error ){_cbc :=_cb [_eaf ];
if _cbc ==nil {return nil ,_b .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_eaf );};return _cbc ,nil ;};var _ Filter =filterAESV3 {};func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };


// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};