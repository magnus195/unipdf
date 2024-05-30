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

// Package xmputil provides abstraction used by the pdf document XMP Metadata.
package xmputil ;import (_e "errors";_ad "fmt";_b "github.com/trimmer-io/go-xmp/models/pdf";_g "github.com/trimmer-io/go-xmp/models/xmp_base";_ab "github.com/trimmer-io/go-xmp/models/xmp_mm";_d "github.com/trimmer-io/go-xmp/xmp";
	_ca "github.com/magnus195/unipdf/v3/core";
	_f "github.com/magnus195/unipdf/v3/internal/timeutils";
	_de "github.com/magnus195/unipdf/v3/internal/uuid";
	_eb "github.com/magnus195/unipdf/v3/model/xmputil/pdfaextension";
	_ef "github.com/magnus195/unipdf/v3/model/xmputil/pdfaid";_da "strconv";_c "time";);

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_af *Document )SetPdfAExtension ()error {_ag ,_eg :=_eb .MakeModel (_af ._bc );if _eg !=nil {return _eg ;};if _eg =_eb .FillModel (_af ._bc ,_ag );_eg !=nil {return _eg ;};if _eg =_ag .SyncToXMP (_af ._bc );_eg !=nil {return _eg ;};return nil ;};


// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument (stream []byte )(*Document ,error ){_ee :=_d .NewDocument ();if _adf :=_d .Unmarshal (stream ,_ee );_adf !=nil {return nil ,_adf ;};return &Document {_bc :_ee },nil ;};

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_ega *Document )GetMediaManagement ()(*MediaManagement ,bool ){_fd :=_ab .FindModel (_ega ._bc );if _fd ==nil {return nil ,false ;};_eae :=make ([]MediaManagementVersion ,len (_fd .Versions ));for _bcg ,_bcf :=range _fd .Versions {_eae [_bcg ]=MediaManagementVersion {VersionID :_bcf .Version ,ModifyDate :_bcf .ModifyDate .Value (),Comments :_bcf .Comments ,Modifier :_bcf .Modifier };
};_ce :=&MediaManagement {OriginalDocumentID :GUID (_fd .OriginalDocumentID .Value ()),DocumentID :GUID (_fd .DocumentID .Value ()),InstanceID :GUID (_fd .InstanceID .Value ()),VersionID :_fd .VersionID ,Versions :_eae };if _fd .DerivedFrom !=nil {_ce .DerivedFrom =&MediaManagementDerivedFrom {OriginalDocumentID :GUID (_fd .DerivedFrom .OriginalDocumentID ),DocumentID :GUID (_fd .DerivedFrom .DocumentID ),InstanceID :GUID (_fd .DerivedFrom .InstanceID ),VersionID :_fd .DerivedFrom .VersionID };
};return _ce ,true ;};

// SetPdfInfo sets the pdf info into selected document.
func (_deg *Document )SetPdfInfo (options *PdfInfoOptions )error {if options ==nil {return _e .New ("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};_egf ,_gea :=_b .MakeModel (_deg ._bc );
if _gea !=nil {return _gea ;};if options .Overwrite {*_egf =_b .PDFInfo {};};if options .InfoDict !=nil {_cd ,_fa :=_ca .GetDict (options .InfoDict );if !_fa {return _ad .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",options .InfoDict );
};var _dec *_ca .PdfObjectString ;for _ ,_gb :=range _cd .Keys (){switch _gb {case "\u0054\u0069\u0074l\u0065":_dec ,_fa =_ca .GetString (_cd .Get ("\u0054\u0069\u0074l\u0065"));if _fa {_egf .Title =_d .NewAltString (_dec );};case "\u0041\u0075\u0074\u0068\u006f\u0072":_dec ,_fa =_ca .GetString (_cd .Get ("\u0041\u0075\u0074\u0068\u006f\u0072"));
if _fa {_egf .Author =_d .NewStringList (_dec .String ());};case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":_dec ,_fa =_ca .GetString (_cd .Get ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"));if _fa {_egf .Keywords =_dec .String ();};case "\u0043r\u0065\u0061\u0074\u006f\u0072":_dec ,_fa =_ca .GetString (_cd .Get ("\u0043r\u0065\u0061\u0074\u006f\u0072"));
if _fa {_egf .Creator =_d .AgentName (_dec .String ());};case "\u0053u\u0062\u006a\u0065\u0063\u0074":_dec ,_fa =_ca .GetString (_cd .Get ("\u0053u\u0062\u006a\u0065\u0063\u0074"));if _fa {_egf .Subject =_d .NewAltString (_dec .String ());};case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":_dec ,_fa =_ca .GetString (_cd .Get ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"));
if _fa {_egf .Producer =_d .AgentName (_dec .String ());};case "\u0054r\u0061\u0070\u0070\u0065\u0064":_ac ,_ege :=_ca .GetName (_cd .Get ("\u0054r\u0061\u0070\u0070\u0065\u0064"));if _ege {switch _ac .String (){case "\u0054\u0072\u0075\u0065":_egf .Trapped =true ;
case "\u0046\u0061\u006cs\u0065":_egf .Trapped =false ;default:_egf .Trapped =true ;};};case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":if _ea ,_cc :=_ca .GetString (_cd .Get ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"));
_cc &&_ea .String ()!=""{_gd ,_fad :=_f .ParsePdfTime (_ea .String ());if _fad !=nil {return _ad .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077",_fad );
};_egf .CreationDate =_d .NewDate (_gd );};case "\u004do\u0064\u0044\u0061\u0074\u0065":if _daa ,_ec :=_ca .GetString (_cd .Get ("\u004do\u0064\u0044\u0061\u0074\u0065"));_ec &&_daa .String ()!=""{_ccg ,_ba :=_f .ParsePdfTime (_daa .String ());if _ba !=nil {return _ad .Errorf ("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w",_ba );
};_egf .ModifyDate =_d .NewDate (_ccg );};};};};if options .PdfVersion !=""{_egf .PDFVersion =options .PdfVersion ;};if options .Marked {_egf .Marked =_d .Bool (options .Marked );};if options .Copyright !=""{_egf .Copyright =options .Copyright ;};if _gea =_egf .SyncToXMP (_deg ._bc );
_gea !=nil {return _gea ;};return nil ;};

// NewDocument creates a new document without any previous xmp information.
func NewDocument ()*Document {_aa :=_d .NewDocument ();return &Document {_bc :_aa }};

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_dbg *Document )SetMediaManagement (options *MediaManagementOptions )error {_agc ,_gdb :=_ab .MakeModel (_dbg ._bc );if _gdb !=nil {return _gdb ;};if options ==nil {options =new (MediaManagementOptions );};_ff :=_ab .ResourceRef {};if _agc .OriginalDocumentID .IsZero (){if options .OriginalDocumentID !=""{_agc .OriginalDocumentID =_d .GUID (options .OriginalDocumentID );
}else {_efa ,_bdc :=_de .NewUUID ();if _bdc !=nil {return _bdc ;};_agc .OriginalDocumentID =_d .GUID (_efa .String ());};}else {_ff .OriginalDocumentID =_agc .OriginalDocumentID ;};switch {case options .DocumentID !="":_agc .DocumentID =_d .GUID (options .DocumentID );
case options .NewDocumentID ||_agc .DocumentID .IsZero ():if !_agc .DocumentID .IsZero (){_ff .DocumentID =_agc .DocumentID ;};_gc ,_ebe :=_de .NewUUID ();if _ebe !=nil {return _ebe ;};_agc .DocumentID =_d .GUID (_gc .String ());};if !_agc .InstanceID .IsZero (){_ff .InstanceID =_agc .InstanceID ;
};_agc .InstanceID =_d .GUID (options .InstanceID );if _agc .InstanceID ==""{_egec ,_afab :=_de .NewUUID ();if _afab !=nil {return _afab ;};_agc .InstanceID =_d .GUID (_egec .String ());};if !_ff .IsZero (){_agc .DerivedFrom =&_ff ;};_ggc :=options .VersionID ;
if _agc .VersionID !=""{_ace ,_cce :=_da .Atoi (_agc .VersionID );if _cce !=nil {_ggc =_da .Itoa (len (_agc .Versions )+1);}else {_ggc =_da .Itoa (_ace +1);};};if _ggc ==""{_ggc ="\u0031";};_agc .VersionID =_ggc ;if _gdb =_agc .SyncToXMP (_dbg ._bc );_gdb !=nil {return _gdb ;
};return nil ;};

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_bdf *Document )GetPdfaExtensionSchemas ()([]_eb .Schema ,error ){_cb :=_bdf ._bc .FindModel (_eb .Namespace );if _cb ==nil {return nil ,nil ;};_gg ,_fe :=_cb .(*_eb .Model );if !_fe {return nil ,_ad .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054",_cb );
};return _gg .Schemas ,nil ;};

// MediaManagementOptions are the options for the Media management xmp metadata.
type MediaManagementOptions struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
// By default, this value is generated.
OriginalDocumentID string ;

// NewDocumentID is a flag which generates a new Document identifier while setting media management.
// This value should be set to true only if the document is stored and saved as new document.
// Otherwise, if the document is modified and overwrites previous file, it should be set to false.
NewDocumentID bool ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
// By default, this value is generated if NewDocumentID is true or previous doesn't exist.
DocumentID string ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
// By default, this value is generated.
InstanceID string ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
// By default, the derived from structure is filled from previous XMP metadata (if exists).
DerivedFrom string ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
// By default, this values is incremented or set to the next version number.
VersionID string ;

// ModifyComment is a comment to given modification
ModifyComment string ;

// ModifyDate is a custom modification date for the versions.
// By default, this would be set to time.Now().
ModifyDate _c .Time ;

// Modifier is a person who did the modification.
Modifier string ;};

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct{InfoDict _ca .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;

// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
Overwrite bool ;};

// GUID is a string representing a globally unique identifier.
type GUID string ;

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_bd *Document )GetGoXmpDocument ()*_d .Document {return _bd ._bc };

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{_bc *_d .Document };

// PdfInfo is the xmp document pdf info.
type PdfInfo struct{InfoDict _ca .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;};

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_cdc *Document )SetPdfAID (part int ,conformance string )error {_be ,_gbe :=_ef .MakeModel (_cdc ._bc );if _gbe !=nil {return _gbe ;};_be .Part =part ;_be .Conformance =conformance ;if _ae :=_be .SyncToXMP (_cdc ._bc );_ae !=nil {return _ae ;};return nil ;
};

// MediaManagement are the values from the document media management metadata.
type MediaManagement struct{

// OriginalDocumentID  as media is imported and projects is started, an original-document ID
// must be created to identify a new document. This identifies a document as a conceptual entity.
OriginalDocumentID GUID ;

// DocumentID when a document is copied to a new file path or converted to a new format with
// Save As, another new document ID should usually be assigned. This identifies a general version or
// branch of a document. You can use it to track different versions or extracted portions of a document
// with the same original-document ID.
DocumentID GUID ;

// InstanceID to track a document’s editing history, you must assign a new instance ID
// whenever a document is saved after any changes. This uniquely identifies an exact version of a
// document. It is used in resource references (to identify both the document or part itself and the
// referenced or referencing documents), and in document-history resource events (to identify the
// document instance that resulted from the change).
InstanceID GUID ;

// DerivedFrom references the source document from which this one is derived,
// typically through a Save As operation that changes the file name or format. It is a minimal reference;
// missing components can be assumed to be unchanged. For example, a new version might only need
// to specify the instance ID and version number of the previous version, or a rendition might only need
// to specify the instance ID and rendition class of the original.
DerivedFrom *MediaManagementDerivedFrom ;

// VersionID are meant to associate the document with a product version that is part of a release process. They can be useful in tracking the
// document history, but should not be used to identify a document uniquely in any context.
// Usually it simply works by incrementing integers 1,2,3...
VersionID string ;

// Versions is the history of the document versions along with the comments, timestamps and issuers.
Versions []MediaManagementVersion ;};

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_ed *Document )MarshalIndent (prefix ,indent string )([]byte ,error ){if _ed ._bc .IsDirty (){if _fc :=_ed ._bc .SyncModels ();_fc !=nil {return nil ,_fc ;};};return _d .MarshalIndent (_ed ._bc ,prefix ,indent );};

// GetPdfAID gets the pdfaid xmp metadata model.
func (_ga *Document )GetPdfAID ()(*PdfAID ,bool ){_ffe ,_egea :=_ga ._bc .FindModel (_ef .Namespace ).(*_ef .Model );if !_egea {return nil ,false ;};return &PdfAID {Part :_ffe .Part ,Conformance :_ffe .Conformance },true ;};

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct{Part int ;Conformance string ;};

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct{OriginalDocumentID GUID ;DocumentID GUID ;InstanceID GUID ;VersionID string ;};

// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct{VersionID string ;ModifyDate _c .Time ;Comments string ;Modifier string ;};

// GetPdfInfo gets the document pdf info.
func (_dg *Document )GetPdfInfo ()(*PdfInfo ,bool ){_bda :=PdfInfo {};var _fg *_ca .PdfObjectDictionary ;_ccge :=func (_gde string ,_db _ca .PdfObject ){if _fg ==nil {_fg =_ca .MakeDict ();};_fg .Set (_ca .PdfObjectName (_gde ),_db );};_dge ,_fgb :=_dg ._bc .FindModel (_b .NsPDF ).(*_b .PDFInfo );
if !_fgb {_def ,_afa :=_dg ._bc .FindModel (_g .NsXmp ).(*_g .XmpBase );if !_afa {return nil ,false ;};if _def .CreatorTool !=""{_ccge ("\u0043r\u0065\u0061\u0074\u006f\u0072",_ca .MakeString (string (_def .CreatorTool )));};if !_def .CreateDate .IsZero (){_ccge ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_ca .MakeString (_f .FormatPdfTime (_def .CreateDate .Value ())));
};if !_def .ModifyDate .IsZero (){_ccge ("\u004do\u0064\u0044\u0061\u0074\u0065",_ca .MakeString (_f .FormatPdfTime (_def .ModifyDate .Value ())));};_bda .InfoDict =_fg ;return &_bda ,true ;};_bda .Copyright =_dge .Copyright ;_bda .PdfVersion =_dge .PDFVersion ;
_bda .Marked =bool (_dge .Marked );if len (_dge .Title )> 0{_ccge ("\u0054\u0069\u0074l\u0065",_ca .MakeString (_dge .Title .Default ()));};if len (_dge .Author )> 0{_ccge ("\u0041\u0075\u0074\u0068\u006f\u0072",_ca .MakeString (_dge .Author [0]));};if _dge .Keywords !=""{_ccge ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",_ca .MakeString (_dge .Keywords ));
};if len (_dge .Subject )> 0{_ccge ("\u0053u\u0062\u006a\u0065\u0063\u0074",_ca .MakeString (_dge .Subject .Default ()));};if _dge .Creator !=""{_ccge ("\u0043r\u0065\u0061\u0074\u006f\u0072",_ca .MakeString (string (_dge .Creator )));};if _dge .Producer !=""{_ccge ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",_ca .MakeString (string (_dge .Producer )));
};if _dge .Trapped {_ccge ("\u0054r\u0061\u0070\u0070\u0065\u0064",_ca .MakeName ("\u0054\u0072\u0075\u0065"));};if !_dge .CreationDate .IsZero (){_ccge ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_ca .MakeString (_f .FormatPdfTime (_dge .CreationDate .Value ())));
};if !_dge .ModifyDate .IsZero (){_ccge ("\u004do\u0064\u0044\u0061\u0074\u0065",_ca .MakeString (_f .FormatPdfTime (_dge .ModifyDate .Value ())));};_bda .InfoDict =_fg ;return &_bda ,true ;};

// Marshal the document into xml byte stream.
func (_ade *Document )Marshal ()([]byte ,error ){if _ade ._bc .IsDirty (){if _ge :=_ade ._bc .SyncModels ();_ge !=nil {return nil ,_ge ;};};return _d .Marshal (_ade ._bc );};