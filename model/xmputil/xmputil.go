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
package xmputil ;import (_d "errors";_bc "fmt";_be "github.com/trimmer-io/go-xmp/models/pdf";_a "github.com/trimmer-io/go-xmp/models/xmp_mm";_e "github.com/trimmer-io/go-xmp/xmp";_fg "github.com/unidoc/unipdf/v3/core";_df "github.com/unidoc/unipdf/v3/internal/timeutils";
_da "github.com/unidoc/unipdf/v3/internal/uuid";_ag "github.com/unidoc/unipdf/v3/model/xmputil/pdfaextension";_c "github.com/unidoc/unipdf/v3/model/xmputil/pdfaid";_b "strconv";_f "time";);

// NewDocument creates a new document without any previous xmp information.
func NewDocument ()*Document {_daf :=_e .NewDocument ();return &Document {_ee :_daf }};

// MarshalIndent the document into xml byte stream with predefined prefix and indent.
func (_eed *Document )MarshalIndent (prefix ,indent string )([]byte ,error ){if _eed ._ee .IsDirty (){if _eb :=_eed ._ee .SyncModels ();_eb !=nil {return nil ,_eb ;};};return _e .MarshalIndent (_eed ._ee ,prefix ,indent );};

// GetMediaManagement gets the media management metadata from provided xmp document.
func (_ec *Document )GetMediaManagement ()(*MediaManagement ,bool ){_deb :=_a .FindModel (_ec ._ee );if _deb ==nil {return nil ,false ;};_bec :=make ([]MediaManagementVersion ,len (_deb .Versions ));for _ca ,_egg :=range _deb .Versions {_bec [_ca ]=MediaManagementVersion {VersionID :_egg .Version ,ModifyDate :_egg .ModifyDate .Value (),Comments :_egg .Comments ,Modifier :_egg .Modifier };
};_fbd :=&MediaManagement {OriginalDocumentID :GUID (_deb .OriginalDocumentID .Value ()),DocumentID :GUID (_deb .DocumentID .Value ()),InstanceID :GUID (_deb .InstanceID .Value ()),VersionID :_deb .VersionID ,Versions :_bec };if _deb .DerivedFrom !=nil {_fbd .DerivedFrom =&MediaManagementDerivedFrom {OriginalDocumentID :GUID (_deb .DerivedFrom .OriginalDocumentID ),DocumentID :GUID (_deb .DerivedFrom .DocumentID ),InstanceID :GUID (_deb .DerivedFrom .InstanceID ),VersionID :_deb .DerivedFrom .VersionID };
};return _fbd ,true ;};

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
ModifyDate _f .Time ;

// Modifier is a person who did the modification.
Modifier string ;};

// MediaManagementDerivedFrom is a structure that contains references of identifiers and versions
// from which given document was derived.
type MediaManagementDerivedFrom struct{OriginalDocumentID GUID ;DocumentID GUID ;InstanceID GUID ;VersionID string ;};

// SetMediaManagement sets up XMP media management metadata: namespace xmpMM.
func (_ebb *Document )SetMediaManagement (options *MediaManagementOptions )error {_ffd ,_ab :=_a .MakeModel (_ebb ._ee );if _ab !=nil {return _ab ;};if options ==nil {options =new (MediaManagementOptions );};_gbg :=_a .ResourceRef {};if _ffd .OriginalDocumentID .IsZero (){if options .OriginalDocumentID !=""{_ffd .OriginalDocumentID =_e .GUID (options .OriginalDocumentID );
}else {_ggf ,_bff :=_da .NewUUID ();if _bff !=nil {return _bff ;};_ffd .OriginalDocumentID =_e .GUID (_ggf .String ());};}else {_gbg .OriginalDocumentID =_ffd .OriginalDocumentID ;};switch {case options .DocumentID !="":_ffd .DocumentID =_e .GUID (options .DocumentID );
case options .NewDocumentID ||_ffd .DocumentID .IsZero ():if !_ffd .DocumentID .IsZero (){_gbg .DocumentID =_ffd .DocumentID ;};_de ,_cf :=_da .NewUUID ();if _cf !=nil {return _cf ;};_ffd .DocumentID =_e .GUID (_de .String ());};if !_ffd .InstanceID .IsZero (){_gbg .InstanceID =_ffd .InstanceID ;
};_ffd .InstanceID =_e .GUID (options .InstanceID );if _ffd .InstanceID ==""{_eeb ,_cbb :=_da .NewUUID ();if _cbb !=nil {return _cbb ;};_ffd .InstanceID =_e .GUID (_eeb .String ());};if !_gbg .IsZero (){_ffd .DerivedFrom =&_gbg ;};_dde :=options .VersionID ;
if _ffd .VersionID !=""{_dbb ,_cff :=_b .Atoi (_ffd .VersionID );if _cff !=nil {_dde =_b .Itoa (len (_ffd .Versions )+1);}else {_dde =_b .Itoa (_dbb +1);};};if _dde ==""{_dde ="\u0031";};_ffd .VersionID =_dde ;if _ab =_ffd .SyncToXMP (_ebb ._ee );_ab !=nil {return _ab ;
};return nil ;};

// PdfInfoOptions are the options used for setting pdf info.
type PdfInfoOptions struct{InfoDict _fg .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;

// Overwrite if set to true, overwrites all values found in the current pdf info xmp model to the ones provided.
Overwrite bool ;};

// GetPdfInfo gets the document pdf info.
func (_ff *Document )GetPdfInfo ()(*PdfInfo ,bool ){_af ,_bfc :=_ff ._ee .FindModel (_be .NsPDF ).(*_be .PDFInfo );if !_bfc {return nil ,false ;};_db :=PdfInfo {};var _dd *_fg .PdfObjectDictionary ;_db .Copyright =_af .Copyright ;_db .PdfVersion =_af .PDFVersion ;
_db .Marked =bool (_af .Marked );_aag :=func (_bgb string ,_ce _fg .PdfObject ){if _dd ==nil {_dd =_fg .MakeDict ();};_dd .Set (_fg .PdfObjectName (_bgb ),_ce );};if len (_af .Title )> 0{_aag ("\u0054\u0069\u0074l\u0065",_fg .MakeString (_af .Title .Default ()));
};if len (_af .Author )> 0{_aag ("\u0041\u0075\u0074\u0068\u006f\u0072",_fg .MakeString (_af .Author [0]));};if _af .Keywords !=""{_aag ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073",_fg .MakeString (_af .Keywords ));};if len (_af .Subject )> 0{_aag ("\u0053u\u0062\u006a\u0065\u0063\u0074",_fg .MakeString (_af .Subject .Default ()));
};if _af .Creator !=""{_aag ("\u0043r\u0065\u0061\u0074\u006f\u0072",_fg .MakeString (string (_af .Creator )));};if _af .Producer !=""{_aag ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072",_fg .MakeString (string (_af .Producer )));};if _af .Trapped {_aag ("\u0054r\u0061\u0070\u0070\u0065\u0064",_fg .MakeName ("\u0054\u0072\u0075\u0065"));
};if !_af .CreationDate .IsZero (){_aag ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065",_fg .MakeString (_df .FormatPdfTime (_af .CreationDate .Value ())));};if !_af .ModifyDate .IsZero (){_aag ("\u004do\u0064\u0044\u0061\u0074\u0065",_fg .MakeString (_df .FormatPdfTime (_af .ModifyDate .Value ())));
};_db .InfoDict =_dd ;return &_db ,true ;};

// GetPdfaExtensionSchemas gets a pdfa extension schemas.
func (_daa *Document )GetPdfaExtensionSchemas ()([]_ag .Schema ,error ){_cb :=_daa ._ee .FindModel (_ag .Namespace );if _cb ==nil {return nil ,nil ;};_aa ,_gfg :=_cb .(*_ag .Model );if !_gfg {return nil ,_bc .Errorf ("\u0069\u006eva\u006c\u0069\u0064 \u006d\u006f\u0064\u0065l f\u006fr \u0070\u0064\u0066\u0061\u0045\u0078\u0074en\u0073\u0069\u006f\u006e\u0073\u003a\u0020%\u0054",_cb );
};return _aa .Schemas ,nil ;};

// Marshal the document into xml byte stream.
func (_cg *Document )Marshal ()([]byte ,error ){if _cg ._ee .IsDirty (){if _cc :=_cg ._ee .SyncModels ();_cc !=nil {return nil ,_cc ;};};return _e .Marshal (_cg ._ee );};

// GetGoXmpDocument gets direct access to the go-xmp.Document.
// All changes done to specified document would result in change of this document 'd'.
func (_bg *Document )GetGoXmpDocument ()*_e .Document {return _bg ._ee };

// GetPdfAID gets the pdfaid xmp metadata model.
func (_bac *Document )GetPdfAID ()(*PdfAID ,bool ){_beg ,_ac :=_bac ._ee .FindModel (_c .Namespace ).(*_c .Model );if !_ac {return nil ,false ;};return &PdfAID {Part :_beg .Part ,Conformance :_beg .Conformance },true ;};

// Document is an implementation of the xmp document.
// It is a wrapper over go-xmp/xmp.Document that provides some Pdf predefined functionality.
type Document struct{_ee *_e .Document };

// PdfInfo is the xmp document pdf info.
type PdfInfo struct{InfoDict _fg .PdfObject ;PdfVersion string ;Copyright string ;Marked bool ;};

// LoadDocument loads up the xmp document from provided input stream.
func LoadDocument (stream []byte )(*Document ,error ){_ae :=_e .NewDocument ();if _ge :=_e .Unmarshal (stream ,_ae );_ge !=nil {return nil ,_ge ;};return &Document {_ee :_ae },nil ;};

// PdfAID is the result of the XMP pdfaid metadata.
type PdfAID struct{Part int ;Conformance string ;};

// SetPdfInfo sets the pdf info into selected document.
func (_eee *Document )SetPdfInfo (options *PdfInfoOptions )error {if options ==nil {return _d .New ("\u006ei\u006c\u0020\u0070\u0064\u0066\u0020\u006f\u0070\u0074\u0069\u006fn\u0073\u0020\u0070\u0072\u006f\u0076\u0069\u0064\u0065\u0064");};_ed ,_agg :=_be .MakeModel (_eee ._ee );
if _agg !=nil {return _agg ;};if options .Overwrite {*_ed =_be .PDFInfo {};};if options .InfoDict !=nil {_ga ,_ad :=_fg .GetDict (options .InfoDict );if !_ad {return _bc .Errorf ("i\u006e\u0076\u0061\u006c\u0069\u0064 \u0070\u0064\u0066\u0020\u006f\u0062\u006a\u0065\u0063t\u0020\u0074\u0079p\u0065:\u0020\u0025\u0054",options .InfoDict );
};var _dgd *_fg .PdfObjectString ;for _ ,_eg :=range _ga .Keys (){switch _eg {case "\u0054\u0069\u0074l\u0065":_dgd ,_ad =_fg .GetString (_ga .Get ("\u0054\u0069\u0074l\u0065"));if _ad {_ed .Title =_e .NewAltString (_dgd );};case "\u0041\u0075\u0074\u0068\u006f\u0072":_dgd ,_ad =_fg .GetString (_ga .Get ("\u0041\u0075\u0074\u0068\u006f\u0072"));
if _ad {_ed .Author =_e .NewStringList (_dgd .String ());};case "\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073":_dgd ,_ad =_fg .GetString (_ga .Get ("\u004b\u0065\u0079\u0077\u006f\u0072\u0064\u0073"));if _ad {_ed .Keywords =_dgd .String ();};case "\u0043r\u0065\u0061\u0074\u006f\u0072":_dgd ,_ad =_fg .GetString (_ga .Get ("\u0043r\u0065\u0061\u0074\u006f\u0072"));
if _ad {_ed .Creator =_e .AgentName (_dgd .String ());};case "\u0053u\u0062\u006a\u0065\u0063\u0074":_dgd ,_ad =_fg .GetString (_ga .Get ("\u0053u\u0062\u006a\u0065\u0063\u0074"));if _ad {_ed .Subject =_e .NewAltString (_dgd .String ());};case "\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072":_dgd ,_ad =_fg .GetString (_ga .Get ("\u0050\u0072\u006f\u0064\u0075\u0063\u0065\u0072"));
if _ad {_ed .Producer =_e .AgentName (_dgd .String ());};case "\u0054r\u0061\u0070\u0070\u0065\u0064":_gb ,_aea :=_fg .GetName (_ga .Get ("\u0054r\u0061\u0070\u0070\u0065\u0064"));if _aea {switch _gb .String (){case "\u0054\u0072\u0075\u0065":_ed .Trapped =true ;
case "\u0046\u0061\u006cs\u0065":_ed .Trapped =false ;default:_ed .Trapped =true ;};};case "\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065":if _ba ,_gg :=_fg .GetString (_ga .Get ("\u0043\u0072\u0065a\u0074\u0069\u006f\u006e\u0044\u0061\u0074\u0065"));
_gg &&_ba .String ()!=""{_fc ,_dc :=_df .ParsePdfTime (_ba .String ());if _dc !=nil {return _bc .Errorf ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0043\u0072e\u0061\u0074\u0069\u006f\u006e\u0044\u0061t\u0065\u0020\u0066\u0069\u0065\u006c\u0064\u003a\u0020\u0025\u0077",_dc );
};_ed .CreationDate =_e .NewDate (_fc );};case "\u004do\u0064\u0044\u0061\u0074\u0065":if _fb ,_cbd :=_fg .GetString (_ga .Get ("\u004do\u0064\u0044\u0061\u0074\u0065"));_cbd &&_fb .String ()!=""{_ada ,_fgg :=_df .ParsePdfTime (_fb .String ());if _fgg !=nil {return _bc .Errorf ("\u0069n\u0076\u0061\u006c\u0069d\u0020\u004d\u006f\u0064\u0044a\u0074e\u0020f\u0069\u0065\u006c\u0064\u003a\u0020\u0025w",_fgg );
};_ed .ModifyDate =_e .NewDate (_ada );};};};};if options .PdfVersion !=""{_ed .PDFVersion =options .PdfVersion ;};if options .Marked {_ed .Marked =_e .Bool (options .Marked );};if options .Copyright !=""{_ed .Copyright =options .Copyright ;};if _agg =_ed .SyncToXMP (_eee ._ee );
_agg !=nil {return _agg ;};return nil ;};

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

// SetPdfAID sets up pdfaid xmp metadata.
// In example: Part: '1' Conformance: 'B' states for PDF/A 1B.
func (_fge *Document )SetPdfAID (part int ,conformance string )error {_dfe ,_fbg :=_c .MakeModel (_fge ._ee );if _fbg !=nil {return _fbg ;};_dfe .Part =part ;_dfe .Conformance =conformance ;if _gbe :=_dfe .SyncToXMP (_fge ._ee );_gbe !=nil {return _gbe ;
};return nil ;};

// GUID is a string representing a globally unique identifier.
type GUID string ;

// SetPdfAExtension sets the pdfaExtension XMP metadata.
func (_gf *Document )SetPdfAExtension ()error {_gd ,_dg :=_ag .MakeModel (_gf ._ee );if _dg !=nil {return _dg ;};if _dg =_ag .FillModel (_gf ._ee ,_gd );_dg !=nil {return _dg ;};if _dg =_gd .SyncToXMP (_gf ._ee );_dg !=nil {return _dg ;};return nil ;};


// MediaManagementVersion is the version of the media management xmp metadata.
type MediaManagementVersion struct{VersionID string ;ModifyDate _f .Time ;Comments string ;Modifier string ;};