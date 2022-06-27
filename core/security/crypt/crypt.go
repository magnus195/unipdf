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

package crypt ;import (_d "crypto/aes";_ef "crypto/cipher";_be "crypto/md5";_bg "crypto/rand";_ab "crypto/rc4";_b "fmt";_c "github.com/unidoc/unipdf/v3/common";_eg "github.com/unidoc/unipdf/v3/core/security";_a "io";);func init (){_ec ("\u0041\u0045\u0053V\u0032",_f )};


// PDFVersion implements Filter interface.
func (_ce filterV2 )PDFVersion ()[2]int {return [2]int {}};type filterAESV3 struct{filterAES };

// NewFilterV2 creates a RC4-based filter with a specified key length (in bytes).
func NewFilterV2 (length int )Filter {_dff ,_fg :=_gbe (FilterDict {Length :length });if _fg !=nil {_c .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020R\u0043\u0034\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_fg );
return filterV2 {_fgb :length };};return _dff ;};func _gbe (_abg FilterDict )(Filter ,error ){if _abg .Length %8!=0{return nil ,_b .Errorf ("\u0063\u0072\u0079p\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069\u0070\u006c\u0065\u0020o\u0066\u0020\u0038\u0020\u0028\u0025\u0064\u0029",_abg .Length );
};if _abg .Length < 5||_abg .Length > 16{if _abg .Length ==40||_abg .Length ==64||_abg .Length ==128{_c .Log .Debug ("\u0053\u0054\u0041\u004e\u0044AR\u0044\u0020V\u0049\u004f\u004c\u0041\u0054\u0049\u004f\u004e\u003a\u0020\u0043\u0072\u0079\u0070\u0074\u0020\u004c\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072s\u0020\u0074\u006f \u0062\u0065\u0020\u0069\u006e\u0020\u0062\u0069\u0074\u0073\u0020\u0072\u0061t\u0068\u0065\u0072\u0020\u0074h\u0061\u006e\u0020\u0062\u0079\u0074\u0065\u0073\u0020-\u0020\u0061s\u0073u\u006d\u0069\u006e\u0067\u0020\u0062\u0069t\u0073\u0020\u0028\u0025\u0064\u0029",_abg .Length );
_abg .Length /=8;}else {return nil ,_b .Errorf ("\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074\u0065\u0072\u0020\u006c\u0065\u006e\u0067\u0074h\u0020\u006e\u006f\u0074\u0020\u0069\u006e \u0072\u0061\u006e\u0067\u0065\u0020\u0034\u0030\u0020\u002d\u00201\u0032\u0038\u0020\u0062\u0069\u0074\u0020\u0028\u0025\u0064\u0029",_abg .Length );
};};return filterV2 {_fgb :_abg .Length },nil ;};

// Name implements Filter interface.
func (filterAESV3 )Name ()string {return "\u0041\u0045\u0053V\u0033"};

// MakeKey implements Filter interface.
func (filterAESV3 )MakeKey (_ ,_ uint32 ,ekey []byte )([]byte ,error ){return ekey ,nil };func (filterIdentity )EncryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };func _fb (_de FilterDict )(Filter ,error ){if _de .Length ==256{_c .Log .Debug ("\u0041\u0045S\u0056\u0033\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_de .Length );
_de .Length /=8;};if _de .Length !=0&&_de .Length !=32{return nil ,_b .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0033\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_de .Length );
};return filterAESV3 {},nil ;};func (filterIdentity )MakeKey (objNum ,genNum uint32 ,fkey []byte )([]byte ,error ){return fkey ,nil };func init (){_ec ("\u0041\u0045\u0053V\u0033",_fb )};func _f (_ad FilterDict )(Filter ,error ){if _ad .Length ==128{_c .Log .Debug ("\u0041\u0045S\u0056\u0032\u0020c\u0072\u0079\u0070\u0074\u0020f\u0069\u006c\u0074\u0065\u0072 l\u0065\u006e\u0067\u0074\u0068\u0020\u0061\u0070\u0070\u0065\u0061\u0072\u0073\u0020\u0074\u006f\u0020\u0062e\u0020i\u006e\u0020\u0062\u0069\u0074\u0073 ra\u0074\u0068\u0065\u0072\u0020\u0074\u0068\u0061\u006e\u0020\u0062\u0079te\u0073 \u002d\u0020\u0061\u0073s\u0075m\u0069n\u0067\u0020b\u0069\u0074s \u0028\u0025\u0064\u0029",_ad .Length );
_ad .Length /=8;};if _ad .Length !=0&&_ad .Length !=16{return nil ,_b .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0041\u0045\u0053\u0056\u0032\u0020\u0063\u0072\u0079\u0070\u0074\u0020\u0066\u0069\u006c\u0074e\u0072\u0020\u006c\u0065\u006eg\u0074\u0068 \u0028\u0025\u0064\u0029",_ad .Length );
};return filterAESV2 {},nil ;};type filterIdentity struct{};func (filterAES )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_ag ,_ge :=_d .NewCipher (okey );if _ge !=nil {return nil ,_ge ;};_c .Log .Trace ("A\u0045\u0053\u0020\u0045nc\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
const _ae =_d .BlockSize ;_cfa :=_ae -len (buf )%_ae ;for _da :=0;_da < _cfa ;_da ++{buf =append (buf ,byte (_cfa ));};_c .Log .Trace ("\u0050a\u0064d\u0065\u0064\u0020\u0074\u006f \u0025\u0064 \u0062\u0079\u0074\u0065\u0073",len (buf ));_gb :=make ([]byte ,_ae +len (buf ));
_fc :=_gb [:_ae ];if _ ,_dad :=_a .ReadFull (_bg .Reader ,_fc );_dad !=nil {return nil ,_dad ;};_fd :=_ef .NewCBCEncrypter (_ag ,_fc );_fd .CryptBlocks (_gb [_ae :],buf );buf =_gb ;_c .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );
return buf ,nil ;};

// DecryptBytes implements Filter interface.
func (filterV2 )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_cdd ,_gg :=_ab .NewCipher (okey );if _gg !=nil {return nil ,_gg ;};_c .Log .Trace ("\u0052\u00434\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_cdd .XORKeyStream (buf ,buf );
_c .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};var _ Filter =filterAESV3 {};

// NewIdentity creates an identity filter that bypasses all data without changes.
func NewIdentity ()Filter {return filterIdentity {}};

// HandlerVersion implements Filter interface.
func (filterAESV2 )HandlerVersion ()(V ,R int ){V ,R =4,4;return ;};

// NewFilterAESV3 creates an AES-based filter with a 256 bit key (AESV3).
func NewFilterAESV3 ()Filter {_aa ,_g :=_fb (FilterDict {});if _g !=nil {_c .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0033\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_g );
return filterAESV3 {};};return _aa ;};

// KeyLength implements Filter interface.
func (filterAESV3 )KeyLength ()int {return 256/8};

// MakeKey implements Filter interface.
func (_gd filterV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ba (objNum ,genNum ,ekey ,false );};type filterAESV2 struct{filterAES };

// HandlerVersion implements Filter interface.
func (filterAESV3 )HandlerVersion ()(V ,R int ){V ,R =5,6;return ;};func init (){_ec ("\u0056\u0032",_gbe )};var _ Filter =filterAESV2 {};

// HandlerVersion implements Filter interface.
func (_gc filterV2 )HandlerVersion ()(V ,R int ){V ,R =2,3;return ;};func (filterIdentity )PDFVersion ()[2]int {return [2]int {}};

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
MakeKey (_fa ,_cbd uint32 ,_gdf []byte )([]byte ,error );

// EncryptBytes encrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and encrypt data in-place.
EncryptBytes (_cfag []byte ,_cgc []byte )([]byte ,error );

// DecryptBytes decrypts a buffer using object encryption key, as returned by MakeKey.
// Implementation may reuse a buffer and decrypt data in-place.
DecryptBytes (_cdb []byte ,_cgb []byte )([]byte ,error );};

// PDFVersion implements Filter interface.
func (filterAESV3 )PDFVersion ()[2]int {return [2]int {2,0}};func _ec (_aba string ,_bgb filterFunc ){if _ ,_faa :=_fbc [_aba ];_faa {panic ("\u0061l\u0072e\u0061\u0064\u0079\u0020\u0072e\u0067\u0069s\u0074\u0065\u0072\u0065\u0064");};_fbc [_aba ]=_bgb ;
};

// KeyLength implements Filter interface.
func (_bf filterV2 )KeyLength ()int {return _bf ._fgb };

// NewFilter creates CryptFilter from a corresponding dictionary.
func NewFilter (d FilterDict )(Filter ,error ){_bae ,_dea :=_fba (d .CFM );if _dea !=nil {return nil ,_dea ;};_ca ,_dea :=_bae (d );if _dea !=nil {return nil ,_dea ;};return _ca ,nil ;};

// NewFilterAESV2 creates an AES-based filter with a 128 bit key (AESV2).
func NewFilterAESV2 ()Filter {_egf ,_ee :=_f (FilterDict {});if _ee !=nil {_c .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075l\u0064\u0020\u006e\u006f\u0074\u0020\u0063re\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0056\u0032\u0020\u0063\u0072\u0079pt\u0020\u0066i\u006c\u0074\u0065\u0072\u003a\u0020\u0025\u0076",_ee );
return filterAESV2 {};};return _egf ;};type filterAES struct{};

// EncryptBytes implements Filter interface.
func (filterV2 )EncryptBytes (buf []byte ,okey []byte )([]byte ,error ){_bad ,_ac :=_ab .NewCipher (okey );if _ac !=nil {return nil ,_ac ;};_c .Log .Trace ("\u0052\u00434\u0020\u0045\u006ec\u0072\u0079\u0070\u0074\u003a\u0020\u0025\u0020\u0078",buf );_bad .XORKeyStream (buf ,buf );
_c .Log .Trace ("\u0074o\u003a\u0020\u0025\u0020\u0078",buf );return buf ,nil ;};func (filterAES )DecryptBytes (buf []byte ,okey []byte )([]byte ,error ){_aea ,_efbe :=_d .NewCipher (okey );if _efbe !=nil {return nil ,_efbe ;};if len (buf )< 16{_c .Log .Debug ("\u0045R\u0052\u004f\u0052\u0020\u0041\u0045\u0053\u0020\u0069\u006e\u0076a\u006c\u0069\u0064\u0020\u0062\u0075\u0066\u0020\u0025\u0073",buf );
return buf ,_b .Errorf ("\u0041\u0045\u0053\u003a B\u0075\u0066\u0020\u006c\u0065\u006e\u0020\u003c\u0020\u0031\u0036\u0020\u0028\u0025d\u0029",len (buf ));};_df :=buf [:16];buf =buf [16:];if len (buf )%16!=0{_c .Log .Debug ("\u0020\u0069\u0076\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (_df ),_df );
_c .Log .Debug ("\u0062\u0075\u0066\u0020\u0028\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );return buf ,_b .Errorf ("\u0041\u0045\u0053\u0020\u0062\u0075\u0066\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006e\u006f\u0074\u0020\u006d\u0075\u006c\u0074\u0069p\u006c\u0065\u0020\u006f\u0066 \u0031\u0036 \u0028\u0025\u0064\u0029",len (buf ));
};_cb :=_ef .NewCBCDecrypter (_aea ,_df );_c .Log .Trace ("A\u0045\u0053\u0020\u0044ec\u0072y\u0070\u0074\u0020\u0028\u0025d\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );_c .Log .Trace ("\u0063\u0068\u006f\u0070\u0020\u0041\u0045\u0053\u0020\u0044\u0065c\u0072\u0079\u0070\u0074\u0020\u0028\u0025\u0064\u0029\u003a \u0025\u0020\u0078",len (buf ),buf );
_cb .CryptBlocks (buf ,buf );_c .Log .Trace ("\u0074\u006f\u0020(\u0025\u0064\u0029\u003a\u0020\u0025\u0020\u0078",len (buf ),buf );if len (buf )==0{_c .Log .Trace ("\u0045\u006d\u0070\u0074\u0079\u0020b\u0075\u0066\u002c\u0020\u0072\u0065\u0074\u0075\u0072\u006e\u0069\u006e\u0067 \u0065\u006d\u0070\u0074\u0079\u0020\u0073t\u0072\u0069\u006e\u0067");
return buf ,nil ;};_bbg :=int (buf [len (buf )-1]);if _bbg > len (buf ){_c .Log .Debug ("\u0049\u006c\u006c\u0065g\u0061\u006c\u0020\u0070\u0061\u0064\u0020\u006c\u0065\u006eg\u0074h\u0020\u0028\u0025\u0064\u0020\u003e\u0020%\u0064\u0029",_bbg ,len (buf ));
return buf ,_b .Errorf ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0070a\u0064\u0020l\u0065\u006e\u0067\u0074\u0068");};buf =buf [:len (buf )-_bbg ];return buf ,nil ;};

// Name implements Filter interface.
func (filterV2 )Name ()string {return "\u0056\u0032"};

// MakeKey implements Filter interface.
func (filterAESV2 )MakeKey (objNum ,genNum uint32 ,ekey []byte )([]byte ,error ){return _ba (objNum ,genNum ,ekey ,true );};var _ Filter =filterV2 {};func (filterIdentity )Name ()string {return "\u0049\u0064\u0065\u006e\u0074\u0069\u0074\u0079"};type filterV2 struct{_fgb int };
func (filterIdentity )DecryptBytes (p []byte ,okey []byte )([]byte ,error ){return p ,nil };func (filterIdentity )KeyLength ()int {return 0};type filterFunc func (_ga FilterDict )(Filter ,error );func _ba (_db ,_dc uint32 ,_bee []byte ,_cd bool )([]byte ,error ){_ff :=make ([]byte ,len (_bee )+5);
for _dcg :=0;_dcg < len (_bee );_dcg ++{_ff [_dcg ]=_bee [_dcg ];};for _dg :=0;_dg < 3;_dg ++{_efbf :=byte ((_db >>uint32 (8*_dg ))&0xff);_ff [_dg +len (_bee )]=_efbf ;};for _ded :=0;_ded < 2;_ded ++{_dfb :=byte ((_dc >>uint32 (8*_ded ))&0xff);_ff [_ded +len (_bee )+3]=_dfb ;
};if _cd {_ff =append (_ff ,0x73);_ff =append (_ff ,0x41);_ff =append (_ff ,0x6C);_ff =append (_ff ,0x54);};_gec :=_be .New ();_gec .Write (_ff );_bc :=_gec .Sum (nil );if len (_bee )+5< 16{return _bc [0:len (_bee )+5],nil ;};return _bc ,nil ;};func _fba (_gf string )(filterFunc ,error ){_cfe :=_fbc [_gf ];
if _cfe ==nil {return nil ,_b .Errorf ("\u0075\u006e\u0073\u0075p\u0070\u006f\u0072\u0074\u0065\u0064\u0020\u0063\u0072\u0079p\u0074 \u0066\u0069\u006c\u0074\u0065\u0072\u003a \u0025\u0071",_gf );};return _cfe ,nil ;};var (_fbc =make (map[string ]filterFunc );
);

// FilterDict represents information from a CryptFilter dictionary.
type FilterDict struct{CFM string ;AuthEvent _eg .AuthEvent ;Length int ;};

// KeyLength implements Filter interface.
func (filterAESV2 )KeyLength ()int {return 128/8};func (filterIdentity )HandlerVersion ()(V ,R int ){return ;};

// PDFVersion implements Filter interface.
func (filterAESV2 )PDFVersion ()[2]int {return [2]int {1,5}};

// Name implements Filter interface.
func (filterAESV2 )Name ()string {return "\u0041\u0045\u0053V\u0032"};