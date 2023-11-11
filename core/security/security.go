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

package security ;import (_cg "bytes";_c "crypto/aes";_ff "crypto/cipher";_cc "crypto/md5";_e "crypto/rand";_g "crypto/rc4";_b "crypto/sha256";_dd "crypto/sha512";_fa "encoding/binary";_fb "errors";_cb "fmt";_ad "github.com/unidoc/unipdf/v3/common";_da "hash";
_d "io";_a "math";);func (_gef stdHandlerR4 )alg3 (R int ,_cce ,_bcd []byte )([]byte ,error ){var _gf []byte ;if len (_bcd )> 0{_gf =_gef .alg3Key (R ,_bcd );}else {_gf =_gef .alg3Key (R ,_cce );};_fab ,_dgg :=_g .NewCipher (_gf );if _dgg !=nil {return nil ,_fb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_fbb :=_gef .paddedPass (_cce );_bga :=make ([]byte ,len (_fbb ));_fab .XORKeyStream (_bga ,_fbb );if R >=3{_cag :=make ([]byte ,len (_gf ));for _aeg :=0;_aeg < 19;_aeg ++{for _cee :=0;_cee < len (_gf );_cee ++{_cag [_cee ]=_gf [_cee ]^byte (_aeg +1);
};_ced ,_acd :=_g .NewCipher (_cag );if _acd !=nil {return nil ,_fb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_ced .XORKeyStream (_bga ,_bga );};};return _bga ,nil ;};func _bf (_ffb ,_ffe string ,_ce int ,_bea []byte )error {if len (_bea )< _ce {return errInvalidField {Func :_ffb ,Field :_ffe ,Exp :_ce ,Got :len (_bea )};
};return nil ;};func (_cdf stdHandlerR4 )alg4 (_abcd []byte ,_ega []byte )([]byte ,error ){_gdb ,_dc :=_g .NewCipher (_abcd );if _dc !=nil {return nil ,_fb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_abg :=[]byte (_adf );
_bcf :=make ([]byte ,len (_abg ));_gdb .XORKeyStream (_bcf ,_abg );return _bcf ,nil ;};func (_ed *ecbEncrypter )BlockSize ()int {return _ed ._ec };type ecbEncrypter ecb ;func (_dac stdHandlerR4 )alg2 (_dga *StdEncryptDict ,_ca []byte )[]byte {_ad .Log .Trace ("\u0061\u006c\u0067\u0032");
_gaa :=_dac .paddedPass (_ca );_ecb :=_cc .New ();_ecb .Write (_gaa );_ecb .Write (_dga .O );var _cfd [4]byte ;_fa .LittleEndian .PutUint32 (_cfd [:],uint32 (_dga .P ));_ecb .Write (_cfd [:]);_ad .Log .Trace ("\u0067o\u0020\u0050\u003a\u0020\u0025\u0020x",_cfd );
_ecb .Write ([]byte (_dac .ID0 ));_ad .Log .Trace ("\u0074\u0068\u0069\u0073\u002e\u0052\u0020\u003d\u0020\u0025d\u0020\u0065\u006e\u0063\u0072\u0079\u0070t\u004d\u0065\u0074\u0061\u0064\u0061\u0074\u0061\u0020\u0025\u0076",_dga .R ,_dga .EncryptMetadata );
if (_dga .R >=4)&&!_dga .EncryptMetadata {_ecb .Write ([]byte {0xff,0xff,0xff,0xff});};_ge :=_ecb .Sum (nil );if _dga .R >=3{_ecb =_cc .New ();for _cea :=0;_cea < 50;_cea ++{_ecb .Reset ();_ecb .Write (_ge [0:_dac .Length /8]);_ge =_ecb .Sum (nil );};};
if _dga .R >=3{return _ge [0:_dac .Length /8];};return _ge [0:5];};

// NewHandlerR6 creates a new standard security handler for R=5 and R=6.
func NewHandlerR6 ()StdHandler {return stdHandlerR6 {}};func (_gbe stdHandlerR6 )alg2a (_gba *StdEncryptDict ,_dgaf []byte )([]byte ,Permissions ,error ){if _dde :=_bf ("\u0061\u006c\u00672\u0061","\u004f",48,_gba .O );_dde !=nil {return nil ,0,_dde ;};
if _bdff :=_bf ("\u0061\u006c\u00672\u0061","\u0055",48,_gba .U );_bdff !=nil {return nil ,0,_bdff ;};if len (_dgaf )> 127{_dgaf =_dgaf [:127];};_bbd ,_eda :=_gbe .alg12 (_gba ,_dgaf );if _eda !=nil {return nil ,0,_eda ;};var (_gdd []byte ;_dae []byte ;
_dbe []byte ;);var _aga Permissions ;if len (_bbd )!=0{_aga =PermOwner ;_fce :=make ([]byte ,len (_dgaf )+8+48);_fbf :=copy (_fce ,_dgaf );_fbf +=copy (_fce [_fbf :],_gba .O [40:48]);copy (_fce [_fbf :],_gba .U [0:48]);_gdd =_fce ;_dae =_gba .OE ;_dbe =_gba .U [0:48];
}else {_bbd ,_eda =_gbe .alg11 (_gba ,_dgaf );if _eda ==nil &&len (_bbd )==0{_bbd ,_eda =_gbe .alg11 (_gba ,[]byte (""));};if _eda !=nil {return nil ,0,_eda ;}else if len (_bbd )==0{return nil ,0,nil ;};_aga =_gba .P ;_eeg :=make ([]byte ,len (_dgaf )+8);
_acf :=copy (_eeg ,_dgaf );copy (_eeg [_acf :],_gba .U [40:48]);_gdd =_eeg ;_dae =_gba .UE ;_dbe =nil ;};if _dcf :=_bf ("\u0061\u006c\u00672\u0061","\u004b\u0065\u0079",32,_dae );_dcf !=nil {return nil ,0,_dcf ;};_dae =_dae [:32];_ef ,_eda :=_gbe .alg2b (_gba .R ,_gdd ,_dgaf ,_dbe );
if _eda !=nil {return nil ,0,_eda ;};_eff ,_eda :=_c .NewCipher (_ef [:32]);if _eda !=nil {return nil ,0,_eda ;};_cfb :=make ([]byte ,_c .BlockSize );_efc :=_ff .NewCBCDecrypter (_eff ,_cfb );_aef :=make ([]byte ,32);_efc .CryptBlocks (_aef ,_dae );if _gba .R ==5{return _aef ,_aga ,nil ;
};_eda =_gbe .alg13 (_gba ,_aef );if _eda !=nil {return nil ,0,_eda ;};return _aef ,_aga ,nil ;};func (_faea stdHandlerR6 )alg12 (_dda *StdEncryptDict ,_bdb []byte )([]byte ,error ){if _ddd :=_bf ("\u0061\u006c\u00671\u0032","\u0055",48,_dda .U );_ddd !=nil {return nil ,_ddd ;
};if _aafa :=_bf ("\u0061\u006c\u00671\u0032","\u004f",48,_dda .O );_aafa !=nil {return nil ,_aafa ;};_bbfg :=make ([]byte ,len (_bdb )+8+48);_bac :=copy (_bbfg ,_bdb );_bac +=copy (_bbfg [_bac :],_dda .O [32:40]);_bac +=copy (_bbfg [_bac :],_dda .U [0:48]);
_gfa ,_abd :=_faea .alg2b (_dda .R ,_bbfg ,_bdb ,_dda .U [0:48]);if _abd !=nil {return nil ,_abd ;};_gfa =_gfa [:32];if !_cg .Equal (_gfa ,_dda .O [:32]){return nil ,nil ;};return _gfa ,nil ;};func (_adb stdHandlerR6 )alg2b (R int ,_eb ,_eba ,_dfc []byte )([]byte ,error ){if R ==5{return _cgg (_eb );
};return _edf (_eb ,_eba ,_dfc );};

// Permissions is a bitmask of access permissions for a PDF file.
type Permissions uint32 ;

// GenerateParams generates and sets O and U parameters for the encryption dictionary.
// It expects R, P and EncryptMetadata fields to be set.
func (_cff stdHandlerR4 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){O ,_ffba :=_cff .alg3 (d .R ,upass ,opass );if _ffba !=nil {_ad .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_ffba );
return nil ,_ffba ;};d .O =O ;_ad .Log .Trace ("\u0067\u0065\u006e\u0020\u004f\u003a\u0020\u0025\u0020\u0078",O );_aa :=_cff .alg2 (d ,upass );U ,_ffba :=_cff .alg5 (_aa ,upass );if _ffba !=nil {_ad .Log .Debug ("\u0045R\u0052\u004fR\u003a\u0020\u0045r\u0072\u006f\u0072\u0020\u0067\u0065\u006ee\u0072\u0061\u0074\u0069\u006e\u0067 \u004f\u0020\u0066\u006f\u0072\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u0069\u006f\u006e\u0020\u0028\u0025\u0073\u0029",_ffba );
return nil ,_ffba ;};d .U =U ;_ad .Log .Trace ("\u0067\u0065\u006e\u0020\u0055\u003a\u0020\u0025\u0020\u0078",U );return _aa ,nil ;};func (_ea stdHandlerR4 )alg3Key (R int ,_ba []byte )[]byte {_ecd :=_cc .New ();_bc :=_ea .paddedPass (_ba );_ecd .Write (_bc );
if R >=3{for _eg :=0;_eg < 50;_eg ++{_ddf :=_ecd .Sum (nil );_ecd =_cc .New ();_ecd .Write (_ddf );};};_fe :=_ecd .Sum (nil );if R ==2{_fe =_fe [0:5];}else {_fe =_fe [0:_ea .Length /8];};return _fe ;};func (_ac errInvalidField )Error ()string {return _cb .Sprintf ("\u0025s\u003a\u0020e\u0078\u0070\u0065\u0063t\u0065\u0064\u0020%\u0073\u0020\u0066\u0069\u0065\u006c\u0064\u0020\u0074o \u0062\u0065\u0020%\u0064\u0020b\u0079\u0074\u0065\u0073\u002c\u0020g\u006f\u0074 \u0025\u0064",_ac .Func ,_ac .Field ,_ac .Exp ,_ac .Got );
};type stdHandlerR4 struct{Length int ;ID0 string ;};func _dab (_ee []byte )(_ff .Block ,error ){_ccbd ,_gb :=_c .NewCipher (_ee );if _gb !=nil {_ad .Log .Error ("\u0045\u0052\u0052\u004f\u0052\u003a\u0020\u0063\u006f\u0075\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0063\u0072\u0065\u0061\u0074\u0065\u0020A\u0045\u0053\u0020\u0063\u0069p\u0068\u0065r\u003a\u0020\u0025\u0076",_gb );
return nil ,_gb ;};return _ccbd ,nil ;};func _edf (_gfb ,_dcad ,_aec []byte )([]byte ,error ){var (_gfbd ,_cbg ,_bae _da .Hash ;);_gfbd =_b .New ();_dace :=make ([]byte ,64);_gea :=_gfbd ;_gea .Write (_gfb );K :=_gea .Sum (_dace [:0]);_eef :=make ([]byte ,64*(127+64+48));
_aca :=func (_ada int )([]byte ,error ){_edb :=len (_dcad )+len (K )+len (_aec );_fbbg :=_eef [:_edb ];_ffg :=copy (_fbbg ,_dcad );_ffg +=copy (_fbbg [_ffg :],K [:]);_ffg +=copy (_fbbg [_ffg :],_aec );if _ffg !=_edb {_ad .Log .Error ("E\u0052\u0052\u004f\u0052\u003a\u0020u\u006e\u0065\u0078\u0070\u0065\u0063t\u0065\u0064\u0020\u0072\u006f\u0075\u006ed\u0020\u0069\u006e\u0070\u0075\u0074\u0020\u0073\u0069\u007ae\u002e");
return nil ,_fb .New ("\u0077\u0072\u006f\u006e\u0067\u0020\u0073\u0069\u007a\u0065");};K1 :=_eef [:_edb *64];_effc (K1 ,_edb );_ecbg ,_gaf :=_dab (K [0:16]);if _gaf !=nil {return nil ,_gaf ;};_gabg :=_ff .NewCBCEncrypter (_ecbg ,K [16:32]);_gabg .CryptBlocks (K1 ,K1 );
E :=K1 ;_beec :=0;for _bbe :=0;_bbe < 16;_bbe ++{_beec +=int (E [_bbe ]%3);};var _df _da .Hash ;switch _beec %3{case 0:_df =_gfbd ;case 1:if _cbg ==nil {_cbg =_dd .New384 ();};_df =_cbg ;case 2:if _bae ==nil {_bae =_dd .New ();};_df =_bae ;};_df .Reset ();
_df .Write (E );K =_df .Sum (_dace [:0]);return E ,nil ;};for _gcae :=0;;{E ,_eca :=_aca (_gcae );if _eca !=nil {return nil ,_eca ;};_bca :=E [len (E )-1];_gcae ++;if _gcae >=64&&_bca <=uint8 (_gcae -32){break ;};};return K [:32],nil ;};func (_edg stdHandlerR6 )alg9 (_ebag *StdEncryptDict ,_fcg []byte ,_eaf []byte )error {if _gac :=_bf ("\u0061\u006c\u0067\u0039","\u004b\u0065\u0079",32,_fcg );
_gac !=nil {return _gac ;};if _dggg :=_bf ("\u0061\u006c\u0067\u0039","\u0055",48,_ebag .U );_dggg !=nil {return _dggg ;};var _bdg [16]byte ;if _ ,_cfbg :=_d .ReadFull (_e .Reader ,_bdg [:]);_cfbg !=nil {return _cfbg ;};_geg :=_bdg [0:8];_cgb :=_bdg [8:16];
_aefb :=_ebag .U [:48];_dacg :=make ([]byte ,len (_eaf )+len (_geg )+len (_aefb ));_cagcc :=copy (_dacg ,_eaf );_cagcc +=copy (_dacg [_cagcc :],_geg );_cagcc +=copy (_dacg [_cagcc :],_aefb );_egc ,_cggf :=_edg .alg2b (_ebag .R ,_dacg ,_eaf ,_aefb );if _cggf !=nil {return _cggf ;
};O :=make ([]byte ,len (_egc )+len (_geg )+len (_cgb ));_cagcc =copy (O ,_egc [:32]);_cagcc +=copy (O [_cagcc :],_geg );_cagcc +=copy (O [_cagcc :],_cgb );_ebag .O =O ;_cagcc =len (_eaf );_cagcc +=copy (_dacg [_cagcc :],_cgb );_egc ,_cggf =_edg .alg2b (_ebag .R ,_dacg ,_eaf ,_aefb );
if _cggf !=nil {return _cggf ;};_fde ,_cggf :=_dab (_egc [:32]);if _cggf !=nil {return _cggf ;};_gaag :=make ([]byte ,_c .BlockSize );_ede :=_ff .NewCBCEncrypter (_fde ,_gaag );OE :=make ([]byte ,32);_ede .CryptBlocks (OE ,_fcg [:32]);_ebag .OE =OE ;return nil ;
};func _cf (_faf _ff .Block )_ff .BlockMode {return (*ecbEncrypter )(_bb (_faf ))};

// Authenticate implements StdHandler interface.
func (_fddd stdHandlerR6 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){return _fddd .alg2a (d ,pass );};const _adf ="\x28\277\116\136\x4e\x75\x8a\x41\x64\000\x4e\x56\377"+"\xfa\001\010\056\x2e\x00\xb6\xd0\x68\076\x80\x2f\014"+"\251\xfe\x64\x53\x69\172";
type ecb struct{_gc _ff .Block ;_ec int ;};func (_fea stdHandlerR6 )alg8 (_eegb *StdEncryptDict ,_cgc []byte ,_gfe []byte )error {if _ffdc :=_bf ("\u0061\u006c\u0067\u0038","\u004b\u0065\u0079",32,_cgc );_ffdc !=nil {return _ffdc ;};var _ffc [16]byte ;
if _ ,_faa :=_d .ReadFull (_e .Reader ,_ffc [:]);_faa !=nil {return _faa ;};_eagg :=_ffc [0:8];_adg :=_ffc [8:16];_ebaf :=make ([]byte ,len (_gfe )+len (_eagg ));_cac :=copy (_ebaf ,_gfe );copy (_ebaf [_cac :],_eagg );_abga ,_caa :=_fea .alg2b (_eegb .R ,_ebaf ,_gfe ,nil );
if _caa !=nil {return _caa ;};U :=make ([]byte ,len (_abga )+len (_eagg )+len (_adg ));_cac =copy (U ,_abga [:32]);_cac +=copy (U [_cac :],_eagg );copy (U [_cac :],_adg );_eegb .U =U ;_cac =len (_gfe );copy (_ebaf [_cac :],_adg );_abga ,_caa =_fea .alg2b (_eegb .R ,_ebaf ,_gfe ,nil );
if _caa !=nil {return _caa ;};_ggc ,_caa :=_dab (_abga [:32]);if _caa !=nil {return _caa ;};_cef :=make ([]byte ,_c .BlockSize );_gee :=_ff .NewCBCEncrypter (_ggc ,_cef );UE :=make ([]byte ,32);_gee .CryptBlocks (UE ,_cgc [:32]);_eegb .UE =UE ;return nil ;
};

// Allowed checks if a set of permissions can be granted.
func (_abc Permissions )Allowed (p2 Permissions )bool {return _abc &p2 ==p2 };func _bb (_cgf _ff .Block )*ecb {return &ecb {_gc :_cgf ,_ec :_cgf .BlockSize ()}};type stdHandlerR6 struct{};func (_fag stdHandlerR4 )alg6 (_db *StdEncryptDict ,_cfe []byte )([]byte ,error ){var (_fdd []byte ;
_bgb error ;);_fbc :=_fag .alg2 (_db ,_cfe );if _db .R ==2{_fdd ,_bgb =_fag .alg4 (_fbc ,_cfe );}else if _db .R >=3{_fdd ,_bgb =_fag .alg5 (_fbc ,_cfe );}else {return nil ,_fb .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};if _bgb !=nil {return nil ,_bgb ;
};_ad .Log .Trace ("\u0063\u0068\u0065\u0063k:\u0020\u0025\u0020\u0078\u0020\u003d\u003d\u0020\u0025\u0020\u0078\u0020\u003f",string (_fdd ),string (_db .U ));_ccb :=_fdd ;_ffd :=_db .U ;if _db .R >=3{if len (_ccb )> 16{_ccb =_ccb [0:16];};if len (_ffd )> 16{_ffd =_ffd [0:16];
};};if !_cg .Equal (_ccb ,_ffd ){return nil ,nil ;};return _fbc ,nil ;};const (EventDocOpen =AuthEvent ("\u0044o\u0063\u004f\u0070\u0065\u006e");EventEFOpen =AuthEvent ("\u0045\u0046\u004f\u0070\u0065\u006e"););func _fg (_ab _ff .Block )_ff .BlockMode {return (*ecbDecrypter )(_bb (_ab ))};
func _effc (_gagd []byte ,_fda int ){_bfc :=_fda ;for _bfc < len (_gagd ){copy (_gagd [_bfc :],_gagd [:_bfc ]);_bfc *=2;};};func (_eaac stdHandlerR6 )alg10 (_egcg *StdEncryptDict ,_ecdc []byte )error {if _adc :=_bf ("\u0061\u006c\u00671\u0030","\u004b\u0065\u0079",32,_ecdc );
_adc !=nil {return _adc ;};_beea :=uint64 (uint32 (_egcg .P ))|(_a .MaxUint32 <<32);Perms :=make ([]byte ,16);_fa .LittleEndian .PutUint64 (Perms [:8],_beea );if _egcg .EncryptMetadata {Perms [8]='T';}else {Perms [8]='F';};copy (Perms [9:12],"\u0061\u0064\u0062");
if _ ,_bfce :=_d .ReadFull (_e .Reader ,Perms [12:16]);_bfce !=nil {return _bfce ;};_bce ,_aaf :=_dab (_ecdc [:32]);if _aaf !=nil {return _aaf ;};_ccbf :=_cf (_bce );_ccbf .CryptBlocks (Perms ,Perms );_egcg .Perms =Perms [:16];return nil ;};type errInvalidField struct{Func string ;
Field string ;Exp int ;Got int ;};func (stdHandlerR4 )paddedPass (_fd []byte )[]byte {_cd :=make ([]byte ,32);_fbe :=copy (_cd ,_fd );for ;_fbe < 32;_fbe ++{_cd [_fbe ]=_adf [_fbe -len (_fd )];};return _cd ;};func (_ded stdHandlerR6 )alg11 (_dfa *StdEncryptDict ,_caff []byte )([]byte ,error ){if _af :=_bf ("\u0061\u006c\u00671\u0031","\u0055",48,_dfa .U );
_af !=nil {return nil ,_af ;};_fae :=make ([]byte ,len (_caff )+8);_cbb :=copy (_fae ,_caff );_cbb +=copy (_fae [_cbb :],_dfa .U [32:40]);_gda ,_fagc :=_ded .alg2b (_dfa .R ,_fae ,_caff ,nil );if _fagc !=nil {return nil ,_fagc ;};_gda =_gda [:32];if !_cg .Equal (_gda ,_dfa .U [:32]){return nil ,nil ;
};return _gda ,nil ;};

// StdHandler is an interface for standard security handlers.
type StdHandler interface{

// GenerateParams uses owner and user passwords to set encryption parameters and generate an encryption key.
// It assumes that R, P and EncryptMetadata are already set.
GenerateParams (_ga *StdEncryptDict ,_be ,_gd []byte )([]byte ,error );

// Authenticate uses encryption dictionary parameters and the password to calculate
// the document encryption key. It also returns permissions that should be granted to a user.
// In case of failed authentication, it returns empty key and zero permissions with no error.
Authenticate (_ae *StdEncryptDict ,_bg []byte )([]byte ,Permissions ,error );};

// AuthEvent is an event type that triggers authentication.
type AuthEvent string ;func (_cdd stdHandlerR6 )alg13 (_fdde *StdEncryptDict ,_bed []byte )error {if _dgf :=_bf ("\u0061\u006c\u00671\u0033","\u004b\u0065\u0079",32,_bed );_dgf !=nil {return _dgf ;};if _cgee :=_bf ("\u0061\u006c\u00671\u0033","\u0050\u0065\u0072m\u0073",16,_fdde .Perms );
_cgee !=nil {return _cgee ;};_ceb :=make ([]byte ,16);copy (_ceb ,_fdde .Perms [:16]);_gad ,_cceb :=_c .NewCipher (_bed [:32]);if _cceb !=nil {return _cceb ;};_fba :=_fg (_gad );_fba .CryptBlocks (_ceb ,_ceb );if !_cg .Equal (_ceb [9:12],[]byte ("\u0061\u0064\u0062")){return _fb .New ("\u0064\u0065\u0063o\u0064\u0065\u0064\u0020p\u0065\u0072\u006d\u0069\u0073\u0073\u0069o\u006e\u0073\u0020\u0061\u0072\u0065\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};_cgd :=Permissions (_fa .LittleEndian .Uint32 (_ceb [0:4]));if _cgd !=_fdde .P {return _fb .New ("\u0070\u0065r\u006d\u0069\u0073\u0073\u0069\u006f\u006e\u0073\u0020\u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066\u0061il\u0065\u0064");
};var _add bool ;if _ceb [8]=='T'{_add =true ;}else if _ceb [8]=='F'{_add =false ;}else {return _fb .New ("\u0064\u0065\u0063\u006f\u0064\u0065\u0064 \u006d\u0065\u0074a\u0064\u0061\u0074\u0061 \u0065\u006e\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e\u0020\u0066\u006c\u0061\u0067\u0020\u0069\u0073\u0020\u0069\u006e\u0076\u0061\u006c\u0069\u0064");
};if _add !=_fdde .EncryptMetadata {return _fb .New ("\u006d\u0065t\u0061\u0064\u0061\u0074a\u0020\u0065n\u0063\u0072\u0079\u0070\u0074\u0069\u006f\u006e \u0076\u0061\u006c\u0069\u0064\u0061\u0074\u0069\u006f\u006e\u0020\u0066a\u0069\u006c\u0065\u0064");
};return nil ;};

// Authenticate implements StdHandler interface.
func (_gag stdHandlerR4 )Authenticate (d *StdEncryptDict ,pass []byte )([]byte ,Permissions ,error ){_ad .Log .Trace ("\u0044\u0065b\u0075\u0067\u0067\u0069n\u0067\u0020a\u0075\u0074\u0068\u0065\u006e\u0074\u0069\u0063a\u0074\u0069\u006f\u006e\u0020\u002d\u0020\u006f\u0077\u006e\u0065\u0072 \u0070\u0061\u0073\u0073");
_gff ,_bag :=_gag .alg7 (d ,pass );if _bag !=nil {return nil ,0,_bag ;};if _gff !=nil {_ad .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _gff ,PermOwner ,nil ;
};_ad .Log .Trace ("\u0044\u0065bu\u0067\u0067\u0069n\u0067\u0020\u0061\u0075the\u006eti\u0063\u0061\u0074\u0069\u006f\u006e\u0020- \u0075\u0073\u0065\u0072\u0020\u0070\u0061s\u0073");_gff ,_bag =_gag .alg6 (d ,pass );if _bag !=nil {return nil ,0,_bag ;
};if _gff !=nil {_ad .Log .Trace ("\u0074h\u0069\u0073\u002e\u0061u\u0074\u0068\u0065\u006e\u0074i\u0063a\u0074e\u0064\u0020\u003d\u0020\u0054\u0072\u0075e");return _gff ,d .P ,nil ;};return nil ,0,nil ;};type ecbDecrypter ecb ;const (PermOwner =Permissions (_a .MaxUint32 );
PermPrinting =Permissions (1<<2);PermModify =Permissions (1<<3);PermExtractGraphics =Permissions (1<<4);PermAnnotate =Permissions (1<<5);PermFillForms =Permissions (1<<8);PermDisabilityExtract =Permissions (1<<9);PermRotateInsert =Permissions (1<<10);PermFullPrintQuality =Permissions (1<<11);
);var _ StdHandler =stdHandlerR4 {};

// StdEncryptDict is a set of additional fields used in standard encryption dictionary.
type StdEncryptDict struct{R int ;P Permissions ;EncryptMetadata bool ;O ,U []byte ;OE ,UE []byte ;Perms []byte ;};

// GenerateParams is the algorithm opposite to alg2a (R>=5).
// It generates U,O,UE,OE,Perms fields using AESv3 encryption.
// There is no algorithm number assigned to this function in the spec.
// It expects R, P and EncryptMetadata fields to be set.
func (_ddeg stdHandlerR6 )GenerateParams (d *StdEncryptDict ,opass ,upass []byte )([]byte ,error ){_ggg :=make ([]byte ,32);if _ ,_ddc :=_d .ReadFull (_e .Reader ,_ggg );_ddc !=nil {return nil ,_ddc ;};d .U =nil ;d .O =nil ;d .UE =nil ;d .OE =nil ;d .Perms =nil ;
if len (upass )> 127{upass =upass [:127];};if len (opass )> 127{opass =opass [:127];};if _feb :=_ddeg .alg8 (d ,_ggg ,upass );_feb !=nil {return nil ,_feb ;};if _cbbb :=_ddeg .alg9 (d ,_ggg ,opass );_cbbb !=nil {return nil ,_cbbb ;};if d .R ==5{return _ggg ,nil ;
};if _fff :=_ddeg .alg10 (d ,_ggg );_fff !=nil {return nil ,_fff ;};return _ggg ,nil ;};func (_dg *ecbDecrypter )CryptBlocks (dst ,src []byte ){if len (src )%_dg ._ec !=0{_ad .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0064\u0065\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_ad .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0064\u0065\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_dg ._gc .Decrypt (dst ,src [:_dg ._ec ]);src =src [_dg ._ec :];dst =dst [_dg ._ec :];};};func _cgg (_daca []byte )([]byte ,error ){_aed :=_b .New ();_aed .Write (_daca );return _aed .Sum (nil ),nil ;};func (_gg *ecbEncrypter )CryptBlocks (dst ,src []byte ){if len (src )%_gg ._ec !=0{_ad .Log .Error ("\u0045\u0052\u0052\u004f\u0052:\u0020\u0045\u0043\u0042\u0020\u0065\u006e\u0063\u0072\u0079\u0070\u0074\u003a \u0069\u006e\u0070\u0075\u0074\u0020\u006e\u006f\u0074\u0020\u0066\u0075\u006c\u006c\u0020\u0062\u006c\u006f\u0063\u006b\u0073");
return ;};if len (dst )< len (src ){_ad .Log .Error ("\u0045R\u0052\u004fR\u003a\u0020\u0045C\u0042\u0020\u0065\u006e\u0063\u0072\u0079p\u0074\u003a\u0020\u006f\u0075\u0074p\u0075\u0074\u0020\u0073\u006d\u0061\u006c\u006c\u0065\u0072\u0020t\u0068\u0061\u006e\u0020\u0069\u006e\u0070\u0075\u0074");
return ;};for len (src )> 0{_gg ._gc .Encrypt (dst ,src [:_gg ._ec ]);src =src [_gg ._ec :];dst =dst [_gg ._ec :];};};func (_eaa stdHandlerR4 )alg5 (_gca []byte ,_fdg []byte )([]byte ,error ){_fc :=_cc .New ();_fc .Write ([]byte (_adf ));_fc .Write ([]byte (_eaa .ID0 ));
_cge :=_fc .Sum (nil );_ad .Log .Trace ("\u0061\u006c\u0067\u0035");_ad .Log .Trace ("\u0065k\u0065\u0079\u003a\u0020\u0025\u0020x",_gca );_ad .Log .Trace ("\u0049D\u003a\u0020\u0025\u0020\u0078",_eaa .ID0 );if len (_cge )!=16{return nil ,_fb .New ("\u0068a\u0073\u0068\u0020\u006c\u0065\u006e\u0067\u0074\u0068\u0020\u006eo\u0074\u0020\u0031\u0036\u0020\u0062\u0079\u0074\u0065\u0073");
};_cagc ,_cdc :=_g .NewCipher (_gca );if _cdc !=nil {return nil ,_fb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");};_dca :=make ([]byte ,16);_cagc .XORKeyStream (_dca ,_cge );_acb :=make ([]byte ,len (_gca ));
for _daa :=0;_daa < 19;_daa ++{for _eag :=0;_eag < len (_gca );_eag ++{_acb [_eag ]=_gca [_eag ]^byte (_daa +1);};_cagc ,_cdc =_g .NewCipher (_acb );if _cdc !=nil {return nil ,_fb .New ("\u0066a\u0069l\u0065\u0064\u0020\u0072\u0063\u0034\u0020\u0063\u0069\u0070\u0068");
};_cagc .XORKeyStream (_dca ,_dca );_ad .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u002c\u0020\u0065\u006b\u0065\u0079:\u0020\u0025\u0020\u0078",_daa ,_acb );_ad .Log .Trace ("\u0069\u0020\u003d\u0020\u0025\u0064\u0020\u002d\u003e\u0020\u0025\u0020\u0078",_daa ,_dca );
};_egf :=make ([]byte ,32);for _fga :=0;_fga < 16;_fga ++{_egf [_fga ]=_dca [_fga ];};_ ,_cdc =_e .Read (_egf [16:32]);if _cdc !=nil {return nil ,_fb .New ("\u0066a\u0069\u006c\u0065\u0064 \u0074\u006f\u0020\u0067\u0065n\u0020r\u0061n\u0064\u0020\u006e\u0075\u006d\u0062\u0065r");
};return _egf ,nil ;};var _ StdHandler =stdHandlerR6 {};

// NewHandlerR4 creates a new standard security handler for R<=4.
func NewHandlerR4 (id0 string ,length int )StdHandler {return stdHandlerR4 {ID0 :id0 ,Length :length }};func (_caf stdHandlerR4 )alg7 (_bbf *StdEncryptDict ,_fbg []byte )([]byte ,error ){_dcg :=_caf .alg3Key (_bbf .R ,_fbg );_cad :=make ([]byte ,len (_bbf .O ));
if _bbf .R ==2{_ag ,_daf :=_g .NewCipher (_dcg );if _daf !=nil {return nil ,_fb .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");};_ag .XORKeyStream (_cad ,_bbf .O );}else if _bbf .R >=3{_agd :=append ([]byte {},_bbf .O ...);
for _bdf :=0;_bdf < 20;_bdf ++{_aea :=append ([]byte {},_dcg ...);for _bee :=0;_bee < len (_dcg );_bee ++{_aea [_bee ]^=byte (19-_bdf );};_gab ,_ade :=_g .NewCipher (_aea );if _ade !=nil {return nil ,_fb .New ("\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0063\u0069\u0070\u0068\u0065\u0072");
};_gab .XORKeyStream (_cad ,_agd );_agd =append ([]byte {},_cad ...);};}else {return nil ,_fb .New ("\u0069n\u0076\u0061\u006c\u0069\u0064\u0020R");};_abca ,_abf :=_caf .alg6 (_bbf ,_cad );if _abf !=nil {return nil ,nil ;};return _abca ,nil ;};func (_ggf *ecbDecrypter )BlockSize ()int {return _ggf ._ec };
