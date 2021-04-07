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

// Package sighandler implements digital signature handlers for PDF signature validation and signing.
package sighandler ;import (_c "bytes";_dc "crypto";_bb "crypto/rand";_eg "crypto/rsa";_g "crypto/x509";_db "crypto/x509/pkix";_cf "encoding/asn1";_de "errors";_f "fmt";_fb "github.com/unidoc/pkcs7";_dd "github.com/unidoc/timestamp";_ddg "github.com/unidoc/unipdf/v3/core";
_ed "github.com/unidoc/unipdf/v3/model";_gd "github.com/unidoc/unipdf/v3/model/sigutil";_b "hash";_e "time";);

// Validate validates PdfSignature.
func (_ccb *docTimeStamp )Validate (sig *_ed .PdfSignature ,digest _ed .Hasher )(_ed .SignatureValidationResult ,error ){_bbe :=sig .Contents .Bytes ();_bdb ,_ffe :=_fb .Parse (_bbe );if _ffe !=nil {return _ed .SignatureValidationResult {},_ffe ;};if _ffe =_bdb .Verify ();
_ffe !=nil {return _ed .SignatureValidationResult {},_ffe ;};var _cae timestampInfo ;_ ,_ffe =_cf .Unmarshal (_bdb .Content ,&_cae );if _ffe !=nil {return _ed .SignatureValidationResult {},_ffe ;};_bfb ,_ffe :=_bed (_cae .MessageImprint .HashAlgorithm .Algorithm );
if _ffe !=nil {return _ed .SignatureValidationResult {},_ffe ;};_afd :=_bfb .New ();_fba :=digest .(*_c .Buffer );_afd .Write (_fba .Bytes ());_gdg :=_afd .Sum (nil );_beg :=_ed .SignatureValidationResult {IsSigned :true ,IsVerified :_c .Equal (_gdg ,_cae .MessageImprint .HashedMessage ),GeneralizedTime :_cae .GeneralizedTime };
return _beg ,nil ;};

// Sign sets the Contents fields for the PdfSignature.
func (_dga *adobeX509RSASHA1 )Sign (sig *_ed .PdfSignature ,digest _ed .Hasher )error {var _caf []byte ;var _gde error ;if _dga ._aeg !=nil {_caf ,_gde =_dga ._aeg (sig ,digest );if _gde !=nil {return _gde ;};}else {_daf ,_dde :=digest .(_b .Hash );if !_dde {return _de .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_cad ,_ :=_ee (_dga ._edb .SignatureAlgorithm );_caf ,_gde =_eg .SignPKCS1v15 (_bb .Reader ,_dga ._gag ,_cad ,_daf .Sum (nil ));if _gde !=nil {return _gde ;};};_caf ,_gde =_cf .Marshal (_caf );if _gde !=nil {return _gde ;};sig .Contents =_ddg .MakeHexString (string (_caf ));
return nil ;};

// InitSignature initialises the PdfSignature.
func (_bba *adobePKCS7Detached )InitSignature (sig *_ed .PdfSignature )error {if !_bba ._gf {if _bba ._fe ==nil {return _de .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _bba ._ec ==nil {return _de .New ("\u0070\u0072\u0069\u0076\u0061\u0074\u0065\u004b\u0065\u0079\u0020m\u0075\u0073\u0074\u0020\u006e\u006f\u0074\u0020\u0062\u0065 \u006e\u0069\u006c");};};_a :=*_bba ;sig .Handler =&_a ;sig .Filter =_ddg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");
sig .SubFilter =_ddg .MakeName ("\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064");sig .Reference =nil ;_ae ,_gdb :=_a .NewDigest (sig );if _gdb !=nil {return _gdb ;};_ae .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _a .Sign (sig ,_ae );};func _ee (_gdf _g .SignatureAlgorithm )(_dc .Hash ,bool ){var _bd _dc .Hash ;switch _gdf {case _g .SHA1WithRSA :_bd =_dc .SHA1 ;case _g .SHA256WithRSA :_bd =_dc .SHA256 ;case _g .SHA384WithRSA :_bd =_dc .SHA384 ;case _g .SHA512WithRSA :_bd =_dc .SHA512 ;
default:return _dc .SHA1 ,false ;};return _bd ,true ;};

// Validate validates PdfSignature.
func (_gg *adobePKCS7Detached )Validate (sig *_ed .PdfSignature ,digest _ed .Hasher )(_ed .SignatureValidationResult ,error ){_cd :=sig .Contents .Bytes ();_gbe ,_bf :=_fb .Parse (_cd );if _bf !=nil {return _ed .SignatureValidationResult {},_bf ;};_ab :=digest .(*_c .Buffer );
_gbe .Content =_ab .Bytes ();if _bf =_gbe .Verify ();_bf !=nil {return _ed .SignatureValidationResult {},_bf ;};return _ed .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;};type timestampInfo struct{Version int ;Policy _cf .RawValue ;
MessageImprint struct{HashAlgorithm _db .AlgorithmIdentifier ;HashedMessage []byte ;};SerialNumber _cf .RawValue ;GeneralizedTime _e .Time ;};

// NewAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached signature handler.
// Both parameters may be nil for the signature validation.
func NewAdobePKCS7Detached (privateKey *_eg .PrivateKey ,certificate *_g .Certificate )(_ed .SignatureHandler ,error ){return &adobePKCS7Detached {_fe :certificate ,_ec :privateKey },nil ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature
func (_gc *adobePKCS7Detached )IsApplicable (sig *_ed .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061\u0064\u0062\u0065.p\u006b\u0063\u0073\u0037\u002e\u0064\u0065\u0074\u0061\u0063\u0068\u0065\u0064";
};

// Sign sets the Contents fields.
func (_cc *adobePKCS7Detached )Sign (sig *_ed .PdfSignature ,digest _ed .Hasher )error {if _cc ._gf {_bc :=_cc ._fd ;if _bc <=0{_bc =8192;};sig .Contents =_ddg .MakeHexString (string (make ([]byte ,_bc )));return nil ;};_aaa :=digest .(*_c .Buffer );_ff ,_ga :=_fb .NewSignedData (_aaa .Bytes ());
if _ga !=nil {return _ga ;};if _af :=_ff .AddSigner (_cc ._fe ,_cc ._ec ,_fb .SignerInfoConfig {});_af !=nil {return _af ;};_ff .Detach ();_bbb ,_ga :=_ff .Finish ();if _ga !=nil {return _ga ;};_ebg :=make ([]byte ,8192);copy (_ebg ,_bbb );sig .Contents =_ddg .MakeHexString (string (_ebg ));
return nil ;};

// Validate validates PdfSignature.
func (_ca *adobeX509RSASHA1 )Validate (sig *_ed .PdfSignature ,digest _ed .Hasher )(_ed .SignatureValidationResult ,error ){_ce ,_ef :=_ca .getCertificate (sig );if _ef !=nil {return _ed .SignatureValidationResult {},_ef ;};_abb :=sig .Contents .Bytes ();
var _cfb []byte ;if _ ,_abc :=_cf .Unmarshal (_abb ,&_cfb );_abc !=nil {return _ed .SignatureValidationResult {},_abc ;};_ba ,_gfg :=digest .(_b .Hash );if !_gfg {return _ed .SignatureValidationResult {},_de .New ("\u0068a\u0073h\u0020\u0074\u0079\u0070\u0065\u0020\u0065\u0072\u0072\u006f\u0072");
};_ecd ,_ :=_ee (_ce .SignatureAlgorithm );if _ge :=_eg .VerifyPKCS1v15 (_ce .PublicKey .(*_eg .PublicKey ),_ecd ,_ba .Sum (nil ),_cfb );_ge !=nil {return _ed .SignatureValidationResult {},_ge ;};return _ed .SignatureValidationResult {IsSigned :true ,IsVerified :true },nil ;
};type adobeX509RSASHA1 struct{_gag *_eg .PrivateKey ;_edb *_g .Certificate ;_aeg SignFunc ;_ad bool ;};

// AdobeX509RSASHA1Opts defines options for configuring the adbe.x509.rsa_sha1
// signature handler.
type AdobeX509RSASHA1Opts struct{

// EstimateSize specifies whether the size of the signature contents
// should be estimated based on the modulus size of the public key
// extracted from the signing certificate. If set to false, a mock Sign
// call is made in order to estimate the size of the signature contents.
EstimateSize bool ;};

// DocTimeStampOpts defines options for configuring the timestamp handler.
type DocTimeStampOpts struct{

// SignatureSize is the estimated size of the signature contents in bytes.
// If not provided, a default signature size of 4192 is used.
// The signing process will report the model.ErrSignNotEnoughSpace error
// if the estimated signature size is smaller than the actual size of the
// signature.
SignatureSize int ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_df *docTimeStamp )IsApplicable (sig *_ed .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031";
};

// InitSignature initialises the PdfSignature.
func (_feb *docTimeStamp )InitSignature (sig *_ed .PdfSignature )error {_gae :=*_feb ;sig .Handler =&_gae ;sig .Filter =_ddg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_ddg .MakeName ("\u0045\u0054\u0053I\u002e\u0052\u0046\u0043\u0033\u0031\u0036\u0031");
sig .Reference =nil ;if _feb ._ddc > 0{sig .Contents =_ddg .MakeHexString (string (make ([]byte ,_feb ._ddc )));}else {_efe ,_aea :=_feb .NewDigest (sig );if _aea !=nil {return _aea ;};_efe .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
if _aea =_gae .Sign (sig ,_efe );_aea !=nil {return _aea ;};_feb ._ddc =_gae ._ddc ;};return nil ;};func (_gdc *adobeX509RSASHA1 )getCertificate (_ccd *_ed .PdfSignature )(*_g .Certificate ,error ){if _gdc ._edb !=nil {return _gdc ._edb ,nil ;};_aag ,_be :=_ccd .GetCerts ();
if _be !=nil {return nil ,_be ;};return _aag [0],nil ;};type docTimeStamp struct{_gec string ;_cg _dc .Hash ;_ddc int ;};

// NewAdobeX509RSASHA1Custom creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. Both the
// certificate and the sign function can be nil for the signature validation.
// NOTE: the handler will do a mock Sign when initializing the signature in
// order to estimate the signature size. Use NewAdobeX509RSASHA1CustomWithOpts
// for configuring the handler to estimate the signature size.
func NewAdobeX509RSASHA1Custom (certificate *_g .Certificate ,signFunc SignFunc )(_ed .SignatureHandler ,error ){return &adobeX509RSASHA1 {_edb :certificate ,_aeg :signFunc },nil ;};func _bed (_bgd _cf .ObjectIdentifier )(_dc .Hash ,error ){switch {case _bgd .Equal (_fb .OIDDigestAlgorithmSHA1 ),_bgd .Equal (_fb .OIDDigestAlgorithmECDSASHA1 ),_bgd .Equal (_fb .OIDDigestAlgorithmDSA ),_bgd .Equal (_fb .OIDDigestAlgorithmDSASHA1 ),_bgd .Equal (_fb .OIDEncryptionAlgorithmRSA ):return _dc .SHA1 ,nil ;
case _bgd .Equal (_fb .OIDDigestAlgorithmSHA256 ),_bgd .Equal (_fb .OIDDigestAlgorithmECDSASHA256 ):return _dc .SHA256 ,nil ;case _bgd .Equal (_fb .OIDDigestAlgorithmSHA384 ),_bgd .Equal (_fb .OIDDigestAlgorithmECDSASHA384 ):return _dc .SHA384 ,nil ;case _bgd .Equal (_fb .OIDDigestAlgorithmSHA512 ),_bgd .Equal (_fb .OIDDigestAlgorithmECDSASHA512 ):return _dc .SHA512 ,nil ;
};return _dc .Hash (0),_fb .ErrUnsupportedAlgorithm ;};

// NewAdobeX509RSASHA1CustomWithOpts creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler with a custom signing function. The
// handler is configured based on the provided options. If no options are
// provided, default options will be used. Both the certificate and the sign
// function can be nil for the signature validation.
func NewAdobeX509RSASHA1CustomWithOpts (certificate *_g .Certificate ,signFunc SignFunc ,opts *AdobeX509RSASHA1Opts )(_ed .SignatureHandler ,error ){if opts ==nil {opts =&AdobeX509RSASHA1Opts {};};return &adobeX509RSASHA1 {_edb :certificate ,_aeg :signFunc ,_ad :opts .EstimateSize },nil ;
};

// Sign sets the Contents fields for the PdfSignature.
func (_ebc *docTimeStamp )Sign (sig *_ed .PdfSignature ,digest _ed .Hasher )error {_ggg ,_dgaa :=_gd .NewTimestampRequest (digest .(*_c .Buffer ),&_dd .RequestOptions {Hash :_ebc ._cg ,Certificates :true });if _dgaa !=nil {return _dgaa ;};_ffc :=_gd .NewTimestampClient ();
_ebd ,_dgaa :=_ffc .GetEncodedToken (_ebc ._gec ,_ggg );if _dgaa !=nil {return _dgaa ;};_aae :=len (_ebd );if _ebc ._ddc > 0&&_aae > _ebc ._ddc {return _ed .ErrSignNotEnoughSpace ;};if _aae > 0{_ebc ._ddc =_aae +128;};sig .Contents =_ddg .MakeHexString (string (_ebd ));
return nil ;};

// NewAdobeX509RSASHA1 creates a new Adobe.PPKMS/Adobe.PPKLite
// adbe.x509.rsa_sha1 signature handler. Both the private key and the
// certificate can be nil for the signature validation.
func NewAdobeX509RSASHA1 (privateKey *_eg .PrivateKey ,certificate *_g .Certificate )(_ed .SignatureHandler ,error ){return &adobeX509RSASHA1 {_edb :certificate ,_gag :privateKey },nil ;};

// NewDigest creates a new digest.
func (_eb *adobePKCS7Detached )NewDigest (sig *_ed .PdfSignature )(_ed .Hasher ,error ){return _c .NewBuffer (nil ),nil ;};

// NewDigest creates a new digest.
func (_ecc *docTimeStamp )NewDigest (sig *_ed .PdfSignature )(_ed .Hasher ,error ){return _c .NewBuffer (nil ),nil ;};

// InitSignature initialises the PdfSignature.
func (_ccg *adobeX509RSASHA1 )InitSignature (sig *_ed .PdfSignature )error {if _ccg ._edb ==nil {return _de .New ("c\u0065\u0072\u0074\u0069\u0066\u0069c\u0061\u0074\u0065\u0020\u006d\u0075\u0073\u0074\u0020n\u006f\u0074\u0020b\u0065 \u006e\u0069\u006c");
};if _ccg ._gag ==nil &&_ccg ._aeg ==nil {return _de .New ("\u006d\u0075\u0073\u0074\u0020\u0070\u0072o\u0076\u0069\u0064e\u0020\u0065\u0069t\u0068\u0065r\u0020\u0061\u0020\u0070\u0072\u0069v\u0061te\u0020\u006b\u0065\u0079\u0020\u006f\u0072\u0020\u0061\u0020\u0073\u0069\u0067\u006e\u0069\u006e\u0067\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");
};_ag :=*_ccg ;sig .Handler =&_ag ;sig .Filter =_ddg .MakeName ("\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065");sig .SubFilter =_ddg .MakeName ("\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031");
sig .Cert =_ddg .MakeString (string (_ag ._edb .Raw ));sig .Reference =nil ;_dg ,_ebe :=_ag .NewDigest (sig );if _ebe !=nil {return _ebe ;};_dg .Write ([]byte ("\u0063\u0061\u006c\u0063\u0075\u006ca\u0074\u0065\u0020\u0074\u0068\u0065\u0020\u0043\u006f\u006e\u0074\u0065\u006et\u0073\u0020\u0066\u0069\u0065\u006c\u0064 \u0073\u0069\u007a\u0065"));
return _ag .sign (sig ,_dg ,_ccg ._ad );};

// NewDocTimeStamp creates a new DocTimeStamp signature handler.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
// NOTE: the handler will do a mock Sign when initializing the signature
// in order to estimate the signature size. Use NewDocTimeStampWithOpts
// for providing the signature size.
func NewDocTimeStamp (timestampServerURL string ,hashAlgorithm _dc .Hash )(_ed .SignatureHandler ,error ){return &docTimeStamp {_gec :timestampServerURL ,_cg :hashAlgorithm },nil ;};

// SignFunc represents a custom signing function. The function should return
// the computed signature.
type SignFunc func (_cb *_ed .PdfSignature ,_ece _ed .Hasher )([]byte ,error );func (_bg *adobeX509RSASHA1 )sign (_dee *_ed .PdfSignature ,_cbd _ed .Hasher ,_agd bool )error {if !_agd {return _bg .Sign (_dee ,_cbd );};_dgab ,_ddgc :=_bg ._edb .PublicKey .(*_eg .PublicKey );
if !_ddgc {return _f .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0075\u0062\u006c\u0069\u0063\u0020\u006b\u0065y\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",_dgab );};_cbg ,_fdc :=_cf .Marshal (make ([]byte ,_dgab .Size ()));if _fdc !=nil {return _fdc ;
};_dee .Contents =_ddg .MakeHexString (string (_cbg ));return nil ;};type adobePKCS7Detached struct{_ec *_eg .PrivateKey ;_fe *_g .Certificate ;_gf bool ;_fd int ;};

// IsApplicable returns true if the signature handler is applicable for the PdfSignature.
func (_ecdc *adobeX509RSASHA1 )IsApplicable (sig *_ed .PdfSignature )bool {if sig ==nil ||sig .Filter ==nil ||sig .SubFilter ==nil {return false ;};return (*sig .Filter =="A\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004d\u0053"||*sig .Filter =="\u0041\u0064\u006f\u0062\u0065\u002e\u0050\u0050\u004b\u004c\u0069\u0074\u0065")&&*sig .SubFilter =="\u0061d\u0062e\u002e\u0078\u0035\u0030\u0039.\u0072\u0073a\u005f\u0073\u0068\u0061\u0031";
};

// NewEmptyAdobePKCS7Detached creates a new Adobe.PPKMS/Adobe.PPKLite adbe.pkcs7.detached
// signature handler. The generated signature is empty and of size signatureLen.
// The signatureLen parameter can be 0 for the signature validation.
func NewEmptyAdobePKCS7Detached (signatureLen int )(_ed .SignatureHandler ,error ){return &adobePKCS7Detached {_gf :true ,_fd :signatureLen },nil ;};

// NewDocTimeStampWithOpts returns a new DocTimeStamp configured using the
// specified options. If no options are provided, default options will be used.
// Both the timestamp server URL and the hash algorithm can be empty for the
// signature validation.
// The following hash algorithms are supported:
// crypto.SHA1, crypto.SHA256, crypto.SHA384, crypto.SHA512.
func NewDocTimeStampWithOpts (timestampServerURL string ,hashAlgorithm _dc .Hash ,opts *DocTimeStampOpts )(_ed .SignatureHandler ,error ){if opts ==nil {opts =&DocTimeStampOpts {};};if opts .SignatureSize <=0{opts .SignatureSize =4192;};return &docTimeStamp {_gec :timestampServerURL ,_cg :hashAlgorithm ,_ddc :opts .SignatureSize },nil ;
};

// NewDigest creates a new digest.
func (_dcg *adobeX509RSASHA1 )NewDigest (sig *_ed .PdfSignature )(_ed .Hasher ,error ){_fdd ,_eef :=_dcg .getCertificate (sig );if _eef !=nil {return nil ,_eef ;};_adc ,_ :=_ee (_fdd .SignatureAlgorithm );return _adc .New (),nil ;};func (_gb *adobePKCS7Detached )getCertificate (_aa *_ed .PdfSignature )(*_g .Certificate ,error ){if _gb ._fe !=nil {return _gb ._fe ,nil ;
};_dbg ,_gbf :=_aa .GetCerts ();if _gbf !=nil {return nil ,_gbf ;};return _dbg [0],nil ;};func (_gba *docTimeStamp )getCertificate (_deg *_ed .PdfSignature )(*_g .Certificate ,error ){_dcc ,_bfa :=_deg .GetCerts ();if _bfa !=nil {return nil ,_bfa ;};return _dcc [0],nil ;
};