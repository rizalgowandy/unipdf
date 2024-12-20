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

// Package fjson provides support for loading PDF form field data from JSON data/files.
package fjson ;import (_f "encoding/json";_cc "github.com/unidoc/unipdf/v3/common";_e "github.com/unidoc/unipdf/v3/core";_ac "github.com/unidoc/unipdf/v3/model";_a "io";_c "os";);

// LoadFromPDFFile loads form field data from a PDF file.
func LoadFromPDFFile (filePath string )(*FieldData ,error ){_ad ,_cfe :=_c .Open (filePath );if _cfe !=nil {return nil ,_cfe ;};defer _ad .Close ();return LoadFromPDF (_ad );};

// LoadFromJSON loads JSON form data from `r`.
func LoadFromJSON (r _a .Reader )(*FieldData ,error ){var _ag FieldData ;_ed :=_f .NewDecoder (r ).Decode (&_ag ._de );if _ed !=nil {return nil ,_ed ;};return &_ag ,nil ;};

// LoadFromPDF loads form field data from a PDF.
func LoadFromPDF (rs _a .ReadSeeker )(*FieldData ,error ){_eda ,_ace :=_ac .NewPdfReader (rs );if _ace !=nil {return nil ,_ace ;};if _eda .AcroForm ==nil {return nil ,nil ;};var _cg []fieldValue ;_be :=_eda .AcroForm .AllFields ();for _ ,_bg :=range _be {var _fg []string ;
_cf :=make (map[string ]struct{});_cfb ,_eb :=_bg .FullName ();if _eb !=nil {return nil ,_eb ;};if _ee ,_fe :=_bg .V .(*_e .PdfObjectString );_fe {_cg =append (_cg ,fieldValue {Name :_cfb ,Value :_ee .Decoded ()});continue ;};var _da string ;for _ ,_df :=range _bg .Annotations {_g ,_bd :=_e .GetName (_df .AS );
if _bd {_da =_g .String ();};_fgf ,_gd :=_e .GetDict (_df .AP );if !_gd {continue ;};_gg ,_ :=_e .GetDict (_fgf .Get ("\u004e"));for _ ,_fd :=range _gg .Keys (){_fec :=_fd .String ();if _ ,_bec :=_cf [_fec ];!_bec {_fg =append (_fg ,_fec );_cf [_fec ]=struct{}{};
};};_fda ,_ :=_e .GetDict (_fgf .Get ("\u0044"));for _ ,_ea :=range _fda .Keys (){_ebd :=_ea .String ();if _ ,_ae :=_cf [_ebd ];!_ae {_fg =append (_fg ,_ebd );_cf [_ebd ]=struct{}{};};};};_ge :=fieldValue {Name :_cfb ,Value :_da ,Options :_fg };_cg =append (_cg ,_ge );
};_bf :=FieldData {_de :_cg };return &_bf ,nil ;};

// FieldImageValues implements model.FieldImageProvider interface.
func (_ga *FieldData )FieldImageValues ()(map[string ]*_ac .Image ,error ){_ffdc :=make (map[string ]*_ac .Image );for _ ,_gdf :=range _ga ._de {if _gdf .ImageValue !=nil {_ffdc [_gdf .Name ]=_gdf .ImageValue ;};};return _ffdc ,nil ;};

// FieldValues implements model.FieldValueProvider interface.
func (_gga *FieldData )FieldValues ()(map[string ]_e .PdfObject ,error ){_ede :=make (map[string ]_e .PdfObject );for _ ,_af :=range _gga ._de {if len (_af .Value )> 0{_ede [_af .Name ]=_e .MakeString (_af .Value );};};return _ede ,nil ;};

// SetImageFromFile assign image file to a specific field identified by fieldName.
func (_eac *FieldData )SetImageFromFile (fieldName string ,imagePath string ,opt []string )error {_cd ,_gdg :=_c .Open (imagePath );if _gdg !=nil {return _gdg ;};defer _cd .Close ();_cfc ,_gdg :=_ac .ImageHandling .Read (_cd );if _gdg !=nil {_cc .Log .Error ("\u0045\u0072\u0072or\u0020\u006c\u006f\u0061\u0064\u0069\u006e\u0067\u0020\u0069\u006d\u0061\u0067\u0065\u003a\u0020\u0025\u0073",_gdg );
return _gdg ;};return _eac .SetImage (fieldName ,_cfc ,opt );};

// JSON returns the field data as a string in JSON format.
func (_fgg FieldData )JSON ()(string ,error ){_fgd ,_ffd :=_f .MarshalIndent (_fgg ._de ,"","\u0020\u0020\u0020\u0020");return string (_fgd ),_ffd ;};

// LoadFromJSONFile loads form field data from a JSON file.
func LoadFromJSONFile (filePath string )(*FieldData ,error ){_b ,_bc :=_c .Open (filePath );if _bc !=nil {return nil ,_bc ;};defer _b .Close ();return LoadFromJSON (_b );};

// FieldData represents form field data loaded from JSON file.
type FieldData struct{_de []fieldValue };

// SetImage assign model.Image to a specific field identified by fieldName.
func (_ce *FieldData )SetImage (fieldName string ,img *_ac .Image ,opt []string )error {_edf :=fieldValue {Name :fieldName ,ImageValue :img ,Options :opt };_ce ._de =append (_ce ._de ,_edf );return nil ;};type fieldValue struct{Name string `json:"name"`;
Value string `json:"value"`;ImageValue *_ac .Image `json:"-"`;

// Options lists allowed values if present.
Options []string `json:"options,omitempty"`;};