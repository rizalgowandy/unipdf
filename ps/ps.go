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

// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
//
// Package ps implements various functionalities needed for handling Postscript for PDF uses, in particular
// for PDF function type 4.
package ps ;import (_ea "bufio";_d "bytes";_e "errors";_dg "fmt";_ac "github.com/unidoc/unipdf/v3/common";_g "github.com/unidoc/unipdf/v3/core";_a "io";_b "math";);func (_ebee *PSOperand )sin (_dfbd *PSStack )error {_dcd ,_efgab :=_dfbd .PopNumberAsFloat64 ();
if _efgab !=nil {return _efgab ;};_cdg :=_b .Sin (_dcd *_b .Pi /180.0);_efgab =_dfbd .Push (MakeReal (_cdg ));return _efgab ;};func (_bg *PSReal )String ()string {return _dg .Sprintf ("\u0025\u002e\u0035\u0066",_bg .Val )};func (_acfc *PSBoolean )String ()string {return _dg .Sprintf ("\u0025\u0076",_acfc .Val )};
func (_gdbb *PSParser )skipSpaces ()(int ,error ){_ccfb :=0;for {_ebdb ,_fddga :=_gdbb ._aeb .Peek (1);if _fddga !=nil {return 0,_fddga ;};if _g .IsWhiteSpace (_ebdb [0]){_gdbb ._aeb .ReadByte ();_ccfb ++;}else {break ;};};return _ccfb ,nil ;};var ErrUnsupportedOperand =_e .New ("\u0075\u006e\u0073\u0075pp\u006f\u0072\u0074\u0065\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064");


// String returns a string representation of the stack.
func (_cfb *PSStack )String ()string {_ccg :="\u005b\u0020";for _ ,_ebaa :=range *_cfb {_ccg +=_ebaa .String ();_ccg +="\u0020";};_ccg +="\u005d";return _ccg ;};func (_cbc *PSOperand )xor (_dab *PSStack )error {_edcf ,_dge :=_dab .Pop ();if _dge !=nil {return _dge ;
};_gagda ,_dge :=_dab .Pop ();if _dge !=nil {return _dge ;};if _feae ,_bfec :=_edcf .(*PSBoolean );_bfec {_gad ,_fbdc :=_gagda .(*PSBoolean );if !_fbdc {return ErrTypeCheck ;};_dge =_dab .Push (MakeBool (_feae .Val !=_gad .Val ));return _dge ;};if _dcgd ,_fga :=_edcf .(*PSInteger );
_fga {_cfa ,_gdeg :=_gagda .(*PSInteger );if !_gdeg {return ErrTypeCheck ;};_dge =_dab .Push (MakeInteger (_dcgd .Val ^_cfa .Val ));return _dge ;};return ErrTypeCheck ;};

// PSObject represents a postscript object.
type PSObject interface{

// Duplicate makes a fresh copy of the PSObject.
Duplicate ()PSObject ;

// DebugString returns a descriptive representation of the PSObject with more information than String()
// for debugging purposes.
DebugString ()string ;

// String returns a string representation of the PSObject.
String ()string ;};var ErrStackUnderflow =_e .New ("\u0073t\u0061c\u006b\u0020\u0075\u006e\u0064\u0065\u0072\u0066\u006c\u006f\u0077");func (_cd *PSOperand )String ()string {return string (*_cd )};func (_dggbd *PSOperand )eq (_fdc *PSStack )error {_efge ,_dff :=_fdc .Pop ();
if _dff !=nil {return _dff ;};_eff ,_dff :=_fdc .Pop ();if _dff !=nil {return _dff ;};_abfa ,_ggg :=_efge .(*PSBoolean );_ccc ,_dbd :=_eff .(*PSBoolean );if _ggg ||_dbd {var _begd error ;if _ggg &&_dbd {_begd =_fdc .Push (MakeBool (_abfa .Val ==_ccc .Val ));
}else {_begd =_fdc .Push (MakeBool (false ));};return _begd ;};var _adgb float64 ;var _bfe float64 ;if _eae ,_cgbfc :=_efge .(*PSInteger );_cgbfc {_adgb =float64 (_eae .Val );}else if _bfea ,_efc :=_efge .(*PSReal );_efc {_adgb =_bfea .Val ;}else {return ErrTypeCheck ;
};if _aea ,_afa :=_eff .(*PSInteger );_afa {_bfe =float64 (_aea .Val );}else if _gcea ,_gcc :=_eff .(*PSReal );_gcc {_bfe =_gcea .Val ;}else {return ErrTypeCheck ;};if _b .Abs (_bfe -_adgb )< _gc {_dff =_fdc .Push (MakeBool (true ));}else {_dff =_fdc .Push (MakeBool (false ));
};return _dff ;};func (_gfd *PSOperand )truncate (_fde *PSStack )error {_dad ,_dfba :=_fde .Pop ();if _dfba !=nil {return _dfba ;};if _fbdf ,_gccc :=_dad .(*PSReal );_gccc {_egddg :=int (_fbdf .Val );_dfba =_fde .Push (MakeReal (float64 (_egddg )));}else if _bcgd ,_bfef :=_dad .(*PSInteger );
_bfef {_dfba =_fde .Push (MakeInteger (_bcgd .Val ));}else {return ErrTypeCheck ;};return _dfba ;};func (_aeea *PSOperand )idiv (_cacg *PSStack )error {_aaf ,_ddc :=_cacg .Pop ();if _ddc !=nil {return _ddc ;};_aafd ,_ddc :=_cacg .Pop ();if _ddc !=nil {return _ddc ;
};_bbf ,_gaa :=_aaf .(*PSInteger );if !_gaa {return ErrTypeCheck ;};if _bbf .Val ==0{return ErrUndefinedResult ;};_bce ,_gaa :=_aafd .(*PSInteger );if !_gaa {return ErrTypeCheck ;};_dcg :=_bce .Val /_bbf .Val ;_ddc =_cacg .Push (MakeInteger (_dcg ));return _ddc ;
};

// MakeOperand returns a new PSOperand object based on string `val`.
func MakeOperand (val string )*PSOperand {_bbgg :=PSOperand (val );return &_bbgg };func (_gde *PSOperand )sub (_aceg *PSStack )error {_ecea ,_debac :=_aceg .Pop ();if _debac !=nil {return _debac ;};_aeae ,_debac :=_aceg .Pop ();if _debac !=nil {return _debac ;
};_ecca ,_cdgc :=_ecea .(*PSReal );_ecg ,_caa :=_ecea .(*PSInteger );if !_cdgc &&!_caa {return ErrTypeCheck ;};_eggd ,_eegg :=_aeae .(*PSReal );_fgc ,_fdg :=_aeae .(*PSInteger );if !_eegg &&!_fdg {return ErrTypeCheck ;};if _caa &&_fdg {_dgge :=_fgc .Val -_ecg .Val ;
_cegd :=_aceg .Push (MakeInteger (_dgge ));return _cegd ;};var _abd float64 =0;if _eegg {_abd =_eggd .Val ;}else {_abd =float64 (_fgc .Val );};if _cdgc {_abd -=_ecca .Val ;}else {_abd -=float64 (_ecg .Val );};_debac =_aceg .Push (MakeReal (_abd ));return _debac ;
};func (_cgf *PSOperand )bitshift (_cdb *PSStack )error {_acfca ,_bdb :=_cdb .PopInteger ();if _bdb !=nil {return _bdb ;};_baga ,_bdb :=_cdb .PopInteger ();if _bdb !=nil {return _bdb ;};var _geef int ;if _acfca >=0{_geef =_baga <<uint (_acfca );}else {_geef =_baga >>uint (-_acfca );
};_bdb =_cdb .Push (MakeInteger (_geef ));return _bdb ;};func (_egc *PSOperand )atan (_eab *PSStack )error {_edbf ,_eeg :=_eab .PopNumberAsFloat64 ();if _eeg !=nil {return _eeg ;};_def ,_eeg :=_eab .PopNumberAsFloat64 ();if _eeg !=nil {return _eeg ;};if _edbf ==0{var _gcb error ;
if _def < 0{_gcb =_eab .Push (MakeReal (270));}else {_gcb =_eab .Push (MakeReal (90));};return _gcb ;};_dga :=_def /_edbf ;_aad :=_b .Atan (_dga )*180/_b .Pi ;_eeg =_eab .Push (MakeReal (_aad ));return _eeg ;};func (_dcag *PSParser )parseFunction ()(*PSProgram ,error ){_defb ,_ :=_dcag ._aeb .ReadByte ();
if _defb !='{'{return nil ,_e .New ("\u0069\u006ev\u0061\u006c\u0069d\u0020\u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");};_bgb :=NewPSProgram ();for {_dcag .skipSpaces ();_dcag .skipComments ();_ccb ,_bef :=_dcag ._aeb .Peek (2);if _bef !=nil {if _bef ==_a .EOF {break ;
};return nil ,_bef ;};_ac .Log .Trace ("\u0050e\u0065k\u0020\u0073\u0074\u0072\u0069\u006e\u0067\u003a\u0020\u0025\u0073",string (_ccb ));if _ccb [0]=='}'{_ac .Log .Trace ("\u0045\u004f\u0046 \u0066\u0075\u006e\u0063\u0074\u0069\u006f\u006e");_dcag ._aeb .ReadByte ();
break ;}else if _ccb [0]=='{'{_ac .Log .Trace ("\u0046u\u006e\u0063\u0074\u0069\u006f\u006e!");_dagg ,_aeeb :=_dcag .parseFunction ();if _aeeb !=nil {return nil ,_aeeb ;};_bgb .Append (_dagg );}else if _g .IsDecimalDigit (_ccb [0])||(_ccb [0]=='-'&&_g .IsDecimalDigit (_ccb [1])){_ac .Log .Trace ("\u002d>\u004e\u0075\u006d\u0062\u0065\u0072!");
_facd ,_fbea :=_dcag .parseNumber ();if _fbea !=nil {return nil ,_fbea ;};_bgb .Append (_facd );}else {_ac .Log .Trace ("\u002d>\u004fp\u0065\u0072\u0061\u006e\u0064 \u006f\u0072 \u0062\u006f\u006f\u006c\u003f");_ccb ,_ =_dcag ._aeb .Peek (5);_gfb :=string (_ccb );
_ac .Log .Trace ("\u0050\u0065\u0065k\u0020\u0073\u0074\u0072\u003a\u0020\u0025\u0073",_gfb );if (len (_gfb )> 4)&&(_gfb [:5]=="\u0066\u0061\u006cs\u0065"){_eabc ,_egb :=_dcag .parseBool ();if _egb !=nil {return nil ,_egb ;};_bgb .Append (_eabc );}else if (len (_gfb )> 3)&&(_gfb [:4]=="\u0074\u0072\u0075\u0065"){_eda ,_ecce :=_dcag .parseBool ();
if _ecce !=nil {return nil ,_ecce ;};_bgb .Append (_eda );}else {_cdgd ,_ade :=_dcag .parseOperand ();if _ade !=nil {return nil ,_ade ;};_bgb .Append (_cdgd );};};};return _bgb ,nil ;};func (_bbae *PSOperand )neg (_acd *PSStack )error {_cecf ,_deg :=_acd .Pop ();
if _deg !=nil {return _deg ;};if _acdd ,_bbg :=_cecf .(*PSReal );_bbg {_deg =_acd .Push (MakeReal (-_acdd .Val ));return _deg ;}else if _adf ,_bcfe :=_cecf .(*PSInteger );_bcfe {_deg =_acd .Push (MakeInteger (-_adf .Val ));return _deg ;}else {return ErrTypeCheck ;
};};func (_aga *PSOperand )div (_aac *PSStack )error {_dgaf ,_ded :=_aac .Pop ();if _ded !=nil {return _ded ;};_egdd ,_ded :=_aac .Pop ();if _ded !=nil {return _ded ;};_gdce ,_cdbf :=_dgaf .(*PSReal );_baf ,_gdf :=_dgaf .(*PSInteger );if !_cdbf &&!_gdf {return ErrTypeCheck ;
};if _cdbf &&_gdce .Val ==0{return ErrUndefinedResult ;};if _gdf &&_baf .Val ==0{return ErrUndefinedResult ;};_acfd ,_bfa :=_egdd .(*PSReal );_ffa ,_fade :=_egdd .(*PSInteger );if !_bfa &&!_fade {return ErrTypeCheck ;};var _fca float64 ;if _bfa {_fca =_acfd .Val ;
}else {_fca =float64 (_ffa .Val );};if _cdbf {_fca /=_gdce .Val ;}else {_fca /=float64 (_baf .Val );};_ded =_aac .Push (MakeReal (_fca ));return _ded ;};

// NewPSParser returns a new instance of the PDF Postscript parser from input data.
func NewPSParser (content []byte )*PSParser {_dgbb :=PSParser {};_egf :=_d .NewBuffer (content );_dgbb ._aeb =_ea .NewReader (_egf );return &_dgbb ;};

// Exec executes the program, typically leaving output values on the stack.
func (_afc *PSProgram )Exec (stack *PSStack )error {for _ ,_bbe :=range *_afc {var _aff error ;switch _ggc :=_bbe .(type ){case *PSInteger :_dgg :=_ggc ;_aff =stack .Push (_dgg );case *PSReal :_gb :=_ggc ;_aff =stack .Push (_gb );case *PSBoolean :_ad :=_ggc ;
_aff =stack .Push (_ad );case *PSProgram :_dc :=_ggc ;_aff =stack .Push (_dc );case *PSOperand :_fae :=_ggc ;_aff =_fae .Exec (stack );default:return ErrTypeCheck ;};if _aff !=nil {return _aff ;};};return nil ;};var ErrStackOverflow =_e .New ("\u0073\u0074\u0061\u0063\u006b\u0020\u006f\u0076\u0065r\u0066\u006c\u006f\u0077");


// MakeReal returns a new PSReal object initialized with `val`.
func MakeReal (val float64 )*PSReal {_ebeed :=PSReal {};_ebeed .Val =val ;return &_ebeed };func (_fac *PSOperand )cvi (_ead *PSStack )error {_fcf ,_beb :=_ead .Pop ();if _beb !=nil {return _beb ;};if _bf ,_ebge :=_fcf .(*PSReal );_ebge {_cad :=int (_bf .Val );
_beb =_ead .Push (MakeInteger (_cad ));}else if _ccf ,_bbbf :=_fcf .(*PSInteger );_bbbf {_bbc :=_ccf .Val ;_beb =_ead .Push (MakeInteger (_bbc ));}else {return ErrTypeCheck ;};return _beb ;};

// PSProgram defines a Postscript program which is a series of PS objects (arguments, commands, programs etc).
type PSProgram []PSObject ;

// PSReal represents a real number.
type PSReal struct{Val float64 ;};func (_gba *PSOperand )cos (_eea *PSStack )error {_afbf ,_cff :=_eea .PopNumberAsFloat64 ();if _cff !=nil {return _cff ;};_dde :=_b .Cos (_afbf *_b .Pi /180.0);_cff =_eea .Push (MakeReal (_dde ));return _cff ;};

// PSExecutor has its own execution stack and is used to executre a PS routine (program).
type PSExecutor struct{Stack *PSStack ;_ge *PSProgram ;};

// MakeInteger returns a new PSInteger object initialized with `val`.
func MakeInteger (val int )*PSInteger {_fbdb :=PSInteger {};_fbdb .Val =val ;return &_fbdb };func (_ggcb *PSOperand )add (_gec *PSStack )error {_dgb ,_eec :=_gec .Pop ();if _eec !=nil {return _eec ;};_ggde ,_eec :=_gec .Pop ();if _eec !=nil {return _eec ;
};_ag ,_cf :=_dgb .(*PSReal );_cgb ,_bgf :=_dgb .(*PSInteger );if !_cf &&!_bgf {return ErrTypeCheck ;};_cga ,_ff :=_ggde .(*PSReal );_eecg ,_ab :=_ggde .(*PSInteger );if !_ff &&!_ab {return ErrTypeCheck ;};if _bgf &&_ab {_fadd :=_cgb .Val +_eecg .Val ;
_dea :=_gec .Push (MakeInteger (_fadd ));return _dea ;};var _abf float64 ;if _cf {_abf =_ag .Val ;}else {_abf =float64 (_cgb .Val );};if _ff {_abf +=_cga .Val ;}else {_abf +=float64 (_eecg .Val );};_eec =_gec .Push (MakeReal (_abf ));return _eec ;};func (_efe *PSOperand )ifCondition (_egdb *PSStack )error {_faa ,_begf :=_egdb .Pop ();
if _begf !=nil {return _begf ;};_da ,_begf :=_egdb .Pop ();if _begf !=nil {return _begf ;};_fbbf ,_gfe :=_faa .(*PSProgram );if !_gfe {return ErrTypeCheck ;};_dggd ,_gfe :=_da .(*PSBoolean );if !_gfe {return ErrTypeCheck ;};if _dggd .Val {_dda :=_fbbf .Exec (_egdb );
return _dda ;};return nil ;};func (_ca *PSOperand )ceiling (_geg *PSStack )error {_aba ,_efg :=_geg .Pop ();if _efg !=nil {return _efg ;};if _afb ,_be :=_aba .(*PSReal );_be {_efg =_geg .Push (MakeReal (_b .Ceil (_afb .Val )));}else if _agd ,_bc :=_aba .(*PSInteger );
_bc {_efg =_geg .Push (MakeInteger (_agd .Val ));}else {_efg =ErrTypeCheck ;};return _efg ;};func (_ace *PSOperand )gt (_bad *PSStack )error {_agf ,_ebd :=_bad .PopNumberAsFloat64 ();if _ebd !=nil {return _ebd ;};_gdb ,_ebd :=_bad .PopNumberAsFloat64 ();
if _ebd !=nil {return _ebd ;};if _b .Abs (_gdb -_agf )< _gc {_fcde :=_bad .Push (MakeBool (false ));return _fcde ;}else if _gdb > _agf {_fbd :=_bad .Push (MakeBool (true ));return _fbd ;}else {_efca :=_bad .Push (MakeBool (false ));return _efca ;};};func (_fcd *PSProgram )String ()string {_edb :="\u007b\u0020";
for _ ,_acb :=range *_fcd {_edb +=_acb .String ();_edb +="\u0020";};_edb +="\u007d";return _edb ;};const _gc =0.000001;func (_dbf *PSOperand )exp (_fbf *PSStack )error {_bgge ,_dcf :=_fbf .PopNumberAsFloat64 ();if _dcf !=nil {return _dcf ;};_fea ,_dcf :=_fbf .PopNumberAsFloat64 ();
if _dcf !=nil {return _dcf ;};if _b .Abs (_bgge )< 1&&_fea < 0{return ErrUndefinedResult ;};_gcec :=_b .Pow (_fea ,_bgge );_dcf =_fbf .Push (MakeReal (_gcec ));return _dcf ;};func (_cba *PSOperand )index (_gbac *PSStack )error {_ffc ,_bff :=_gbac .Pop ();
if _bff !=nil {return _bff ;};_cabb ,_deba :=_ffc .(*PSInteger );if !_deba {return ErrTypeCheck ;};if _cabb .Val < 0{return ErrRangeCheck ;};if _cabb .Val > len (*_gbac )-1{return ErrStackUnderflow ;};_fgd :=(*_gbac )[len (*_gbac )-1-_cabb .Val ];_bff =_gbac .Push (_fgd .Duplicate ());
return _bff ;};func (_fgf *PSOperand )not (_fbfc *PSStack )error {_abb ,_bdf :=_fbfc .Pop ();if _bdf !=nil {return _bdf ;};if _affa ,_ggca :=_abb .(*PSBoolean );_ggca {_bdf =_fbfc .Push (MakeBool (!_affa .Val ));return _bdf ;}else if _cag ,_badbb :=_abb .(*PSInteger );
_badbb {_bdf =_fbfc .Push (MakeInteger (^_cag .Val ));return _bdf ;}else {return ErrTypeCheck ;};};func (_dgbd *PSOperand )ln (_feb *PSStack )error {_bggb ,_aegg :=_feb .PopNumberAsFloat64 ();if _aegg !=nil {return _aegg ;};_badb :=_b .Log (_bggb );_aegg =_feb .Push (MakeReal (_badb ));
return _aegg ;};func (_ggd *PSOperand )DebugString ()string {return _dg .Sprintf ("\u006fp\u003a\u0027\u0025\u0073\u0027",*_ggd );};func (_cggg *PSOperand )exch (_dbc *PSStack )error {_fg ,_abac :=_dbc .Pop ();if _abac !=nil {return _abac ;};_cac ,_abac :=_dbc .Pop ();
if _abac !=nil {return _abac ;};_abac =_dbc .Push (_fg );if _abac !=nil {return _abac ;};_abac =_dbc .Push (_cac );return _abac ;};func (_aag *PSOperand )pop (_fda *PSStack )error {_ ,_aacc :=_fda .Pop ();if _aacc !=nil {return _aacc ;};return nil ;};func (_ebf *PSOperand )or (_cgfg *PSStack )error {_dgag ,_aaca :=_cgfg .Pop ();
if _aaca !=nil {return _aaca ;};_adge ,_aaca :=_cgfg .Pop ();if _aaca !=nil {return _aaca ;};if _dag ,_gefa :=_dgag .(*PSBoolean );_gefa {_cgfa ,_gfeg :=_adge .(*PSBoolean );if !_gfeg {return ErrTypeCheck ;};_aaca =_cgfg .Push (MakeBool (_dag .Val ||_cgfa .Val ));
return _aaca ;};if _fab ,_eeae :=_dgag .(*PSInteger );_eeae {_gdcg ,_caba :=_adge .(*PSInteger );if !_caba {return ErrTypeCheck ;};_aaca =_cgfg .Push (MakeInteger (_fab .Val |_gdcg .Val ));return _aaca ;};return ErrTypeCheck ;};func (_eef *PSOperand )mod (_cadb *PSStack )error {_fgdf ,_bbd :=_cadb .Pop ();
if _bbd !=nil {return _bbd ;};_ffae ,_bbd :=_cadb .Pop ();if _bbd !=nil {return _bbd ;};_bde ,_faac :=_fgdf .(*PSInteger );if !_faac {return ErrTypeCheck ;};if _bde .Val ==0{return ErrUndefinedResult ;};_faca ,_faac :=_ffae .(*PSInteger );if !_faac {return ErrTypeCheck ;
};_ebc :=_faca .Val %_bde .Val ;_bbd =_cadb .Push (MakeInteger (_ebc ));return _bbd ;};

// Push pushes an object on top of the stack.
func (_aega *PSStack )Push (obj PSObject )error {if len (*_aega )> 100{return ErrStackOverflow ;};*_aega =append (*_aega ,obj );return nil ;};

// NewPSStack returns an initialized PSStack.
func NewPSStack ()*PSStack {return &PSStack {}};func (_ebeec *PSParser )parseBool ()(*PSBoolean ,error ){_bbfd ,_bgafe :=_ebeec ._aeb .Peek (4);if _bgafe !=nil {return MakeBool (false ),_bgafe ;};if (len (_bbfd )>=4)&&(string (_bbfd [:4])=="\u0074\u0072\u0075\u0065"){_ebeec ._aeb .Discard (4);
return MakeBool (true ),nil ;};_bbfd ,_bgafe =_ebeec ._aeb .Peek (5);if _bgafe !=nil {return MakeBool (false ),_bgafe ;};if (len (_bbfd )>=5)&&(string (_bbfd [:5])=="\u0066\u0061\u006cs\u0065"){_ebeec ._aeb .Discard (5);return MakeBool (false ),nil ;};
return MakeBool (false ),_e .New ("\u0075n\u0065\u0078\u0070\u0065c\u0074\u0065\u0064\u0020\u0062o\u006fl\u0065a\u006e\u0020\u0073\u0074\u0072\u0069\u006eg");};func _bbcf (_cbcf int )int {if _cbcf < 0{return -_cbcf ;};return _cbcf ;};func (_gccf *PSOperand )le (_dcfb *PSStack )error {_aegf ,_bddb :=_dcfb .PopNumberAsFloat64 ();
if _bddb !=nil {return _bddb ;};_ffca ,_bddb :=_dcfb .PopNumberAsFloat64 ();if _bddb !=nil {return _bddb ;};if _b .Abs (_ffca -_aegf )< _gc {_acg :=_dcfb .Push (MakeBool (true ));return _acg ;}else if _ffca < _aegf {_cgfd :=_dcfb .Push (MakeBool (true ));
return _cgfd ;}else {_dggc :=_dcfb .Push (MakeBool (false ));return _dggc ;};};func (_deb *PSOperand )dup (_dbb *PSStack )error {_aadg ,_gag :=_dbb .Pop ();if _gag !=nil {return _gag ;};_gag =_dbb .Push (_aadg );if _gag !=nil {return _gag ;};_gag =_dbb .Push (_aadg .Duplicate ());
return _gag ;};func (_bbb *PSBoolean )Duplicate ()PSObject {_ee :=PSBoolean {};_ee .Val =_bbb .Val ;return &_ee };func (_ed *PSInteger )Duplicate ()PSObject {_ef :=PSInteger {};_ef .Val =_ed .Val ;return &_ef };

// PSParser is a basic Postscript parser.
type PSParser struct{_aeb *_ea .Reader };

// PopNumberAsFloat64 pops and return the numeric value of the top of the stack as a float64.
// Real or integer only.
func (_ccfbc *PSStack )PopNumberAsFloat64 ()(float64 ,error ){_ffd ,_dfc :=_ccfbc .Pop ();if _dfc !=nil {return 0,_dfc ;};if _aeeaf ,_eaa :=_ffd .(*PSReal );_eaa {return _aeeaf .Val ,nil ;}else if _cdbg ,_eed :=_ffd .(*PSInteger );_eed {return float64 (_cdbg .Val ),nil ;
}else {return 0,ErrTypeCheck ;};};

// Exec executes the operand `op` in the state specified by `stack`.
func (_bba *PSOperand )Exec (stack *PSStack )error {_fad :=ErrUnsupportedOperand ;switch *_bba {case "\u0061\u0062\u0073":_fad =_bba .abs (stack );case "\u0061\u0064\u0064":_fad =_bba .add (stack );case "\u0061\u006e\u0064":_fad =_bba .and (stack );case "\u0061\u0074\u0061\u006e":_fad =_bba .atan (stack );
case "\u0062\u0069\u0074\u0073\u0068\u0069\u0066\u0074":_fad =_bba .bitshift (stack );case "\u0063e\u0069\u006c\u0069\u006e\u0067":_fad =_bba .ceiling (stack );case "\u0063\u006f\u0070\u0079":_fad =_bba .copy (stack );case "\u0063\u006f\u0073":_fad =_bba .cos (stack );
case "\u0063\u0076\u0069":_fad =_bba .cvi (stack );case "\u0063\u0076\u0072":_fad =_bba .cvr (stack );case "\u0064\u0069\u0076":_fad =_bba .div (stack );case "\u0064\u0075\u0070":_fad =_bba .dup (stack );case "\u0065\u0071":_fad =_bba .eq (stack );case "\u0065\u0078\u0063\u0068":_fad =_bba .exch (stack );
case "\u0065\u0078\u0070":_fad =_bba .exp (stack );case "\u0066\u006c\u006fo\u0072":_fad =_bba .floor (stack );case "\u0067\u0065":_fad =_bba .ge (stack );case "\u0067\u0074":_fad =_bba .gt (stack );case "\u0069\u0064\u0069\u0076":_fad =_bba .idiv (stack );
case "\u0069\u0066":_fad =_bba .ifCondition (stack );case "\u0069\u0066\u0065\u006c\u0073\u0065":_fad =_bba .ifelse (stack );case "\u0069\u006e\u0064e\u0078":_fad =_bba .index (stack );case "\u006c\u0065":_fad =_bba .le (stack );case "\u006c\u006f\u0067":_fad =_bba .log (stack );
case "\u006c\u006e":_fad =_bba .ln (stack );case "\u006c\u0074":_fad =_bba .lt (stack );case "\u006d\u006f\u0064":_fad =_bba .mod (stack );case "\u006d\u0075\u006c":_fad =_bba .mul (stack );case "\u006e\u0065":_fad =_bba .ne (stack );case "\u006e\u0065\u0067":_fad =_bba .neg (stack );
case "\u006e\u006f\u0074":_fad =_bba .not (stack );case "\u006f\u0072":_fad =_bba .or (stack );case "\u0070\u006f\u0070":_fad =_bba .pop (stack );case "\u0072\u006f\u0075n\u0064":_fad =_bba .round (stack );case "\u0072\u006f\u006c\u006c":_fad =_bba .roll (stack );
case "\u0073\u0069\u006e":_fad =_bba .sin (stack );case "\u0073\u0071\u0072\u0074":_fad =_bba .sqrt (stack );case "\u0073\u0075\u0062":_fad =_bba .sub (stack );case "\u0074\u0072\u0075\u006e\u0063\u0061\u0074\u0065":_fad =_bba .truncate (stack );case "\u0078\u006f\u0072":_fad =_bba .xor (stack );
};return _fad ;};

// MakeBool returns a new PSBoolean object initialized with `val`.
func MakeBool (val bool )*PSBoolean {_bdcb :=PSBoolean {};_bdcb .Val =val ;return &_bdcb };

// PopInteger specificially pops an integer from the top of the stack, returning the value as an int.
func (_faaf *PSStack )PopInteger ()(int ,error ){_gdfc ,_dbe :=_faaf .Pop ();if _dbe !=nil {return 0,_dbe ;};if _ggea ,_dfd :=_gdfc .(*PSInteger );_dfd {return _ggea .Val ,nil ;};return 0,ErrTypeCheck ;};func (_fcc *PSOperand )log (_fedf *PSStack )error {_fee ,_ebb :=_fedf .PopNumberAsFloat64 ();
if _ebb !=nil {return _ebb ;};_gdbc :=_b .Log10 (_fee );_ebb =_fedf .Push (MakeReal (_gdbc ));return _ebb ;};func (_gef *PSInteger )DebugString ()string {return _dg .Sprintf ("\u0069\u006e\u0074\u003a\u0025\u0064",_gef .Val );};func (_cfg *PSOperand )ifelse (_bdd *PSStack )error {_edc ,_baa :=_bdd .Pop ();
if _baa !=nil {return _baa ;};_cabc ,_baa :=_bdd .Pop ();if _baa !=nil {return _baa ;};_ffe ,_baa :=_bdd .Pop ();if _baa !=nil {return _baa ;};_eega ,_cfd :=_edc .(*PSProgram );if !_cfd {return ErrTypeCheck ;};_cdc ,_cfd :=_cabc .(*PSProgram );if !_cfd {return ErrTypeCheck ;
};_cec ,_cfd :=_ffe .(*PSBoolean );if !_cfd {return ErrTypeCheck ;};if _cec .Val {_acea :=_cdc .Exec (_bdd );return _acea ;};_baa =_eega .Exec (_bdd );return _baa ;};

// Pop pops an object from the top of the stack.
func (_gfef *PSStack )Pop ()(PSObject ,error ){if len (*_gfef )< 1{return nil ,ErrStackUnderflow ;};_cfga :=(*_gfef )[len (*_gfef )-1];*_gfef =(*_gfef )[0:len (*_gfef )-1];return _cfga ,nil ;};

// PSBoolean represents a boolean value.
type PSBoolean struct{Val bool ;};func (_gae *PSOperand )and (_acc *PSStack )error {_faf ,_bd :=_acc .Pop ();if _bd !=nil {return _bd ;};_cgbf ,_bd :=_acc .Pop ();if _bd !=nil {return _bd ;};if _adg ,_cbbd :=_faf .(*PSBoolean );_cbbd {_bbaf ,_geb :=_cgbf .(*PSBoolean );
if !_geb {return ErrTypeCheck ;};_bd =_acc .Push (MakeBool (_adg .Val &&_bbaf .Val ));return _bd ;};if _df ,_aa :=_faf .(*PSInteger );_aa {_dd ,_afd :=_cgbf .(*PSInteger );if !_afd {return ErrTypeCheck ;};_bd =_acc .Push (MakeInteger (_df .Val &_dd .Val ));
return _bd ;};return ErrTypeCheck ;};

// Execute executes the program for an input parameters `objects` and returns a slice of output objects.
func (_gf *PSExecutor )Execute (objects []PSObject )([]PSObject ,error ){for _ ,_cbb :=range objects {_acf :=_gf .Stack .Push (_cbb );if _acf !=nil {return nil ,_acf ;};};_gge :=_gf ._ge .Exec (_gf .Stack );if _gge !=nil {_ac .Log .Debug ("\u0045x\u0065c\u0020\u0066\u0061\u0069\u006c\u0065\u0064\u003a\u0020\u0025\u0076",_gge );
return nil ,_gge ;};_cbba :=[]PSObject (*_gf .Stack );_gf .Stack .Empty ();return _cbba ,nil ;};func (_caf *PSParser )skipComments ()error {if _ ,_edcfg :=_caf .skipSpaces ();_edcfg !=nil {return _edcfg ;};_abe :=true ;for {_cef ,_gaf :=_caf ._aeb .Peek (1);
if _gaf !=nil {_ac .Log .Debug ("\u0045\u0072\u0072\u006f\u0072\u0020\u0025\u0073",_gaf .Error ());return _gaf ;};if _abe &&_cef [0]!='%'{return nil ;};_abe =false ;if (_cef [0]!='\r')&&(_cef [0]!='\n'){_caf ._aeb .ReadByte ();}else {break ;};};return _caf .skipComments ();
};

// PSObjectArrayToFloat64Array converts []PSObject into a []float64 array. Each PSObject must represent a number,
// otherwise a ErrTypeCheck error occurs.
func PSObjectArrayToFloat64Array (objects []PSObject )([]float64 ,error ){var _f []float64 ;for _ ,_fb :=range objects {if _fe ,_fcb :=_fb .(*PSInteger );_fcb {_f =append (_f ,float64 (_fe .Val ));}else if _cc ,_af :=_fb .(*PSReal );_af {_f =append (_f ,_cc .Val );
}else {return nil ,ErrTypeCheck ;};};return _f ,nil ;};

// NewPSProgram returns an empty, initialized PSProgram.
func NewPSProgram ()*PSProgram {return &PSProgram {}};

// NewPSExecutor returns an initialized PSExecutor for an input `program`.
func NewPSExecutor (program *PSProgram )*PSExecutor {_cb :=&PSExecutor {};_cb .Stack =NewPSStack ();_cb ._ge =program ;return _cb ;};func (_ecc *PSOperand )round (_efga *PSStack )error {_gca ,_ece :=_efga .Pop ();if _ece !=nil {return _ece ;};if _cacgc ,_dgbf :=_gca .(*PSReal );
_dgbf {_ece =_efga .Push (MakeReal (_b .Floor (_cacgc .Val +0.5)));}else if _dfb ,_ccd :=_gca .(*PSInteger );_ccd {_ece =_efga .Push (MakeInteger (_dfb .Val ));}else {return ErrTypeCheck ;};return _ece ;};

// DebugString returns a descriptive string representation of the stack - intended for debugging.
func (_gaab *PSStack )DebugString ()string {_fdbc :="\u005b\u0020";for _ ,_ccdd :=range *_gaab {_fdbc +=_ccdd .DebugString ();_fdbc +="\u0020";};_fdbc +="\u005d";return _fdbc ;};var ErrUndefinedResult =_e .New ("\u0075\u006e\u0064\u0065fi\u006e\u0065\u0064\u0020\u0072\u0065\u0073\u0075\u006c\u0074\u0020\u0065\u0072\u0072o\u0072");
func (_cg *PSOperand )abs (_de *PSStack )error {_eba ,_dca :=_de .Pop ();if _dca !=nil {return _dca ;};if _gga ,_cbd :=_eba .(*PSReal );_cbd {_eg :=_gga .Val ;if _eg < 0{_dca =_de .Push (MakeReal (-_eg ));}else {_dca =_de .Push (MakeReal (_eg ));};}else if _ebg ,_bag :=_eba .(*PSInteger );
_bag {_ebe :=_ebg .Val ;if _ebe < 0{_dca =_de .Push (MakeInteger (-_ebe ));}else {_dca =_de .Push (MakeInteger (_ebe ));};}else {return ErrTypeCheck ;};return _dca ;};

// Empty empties the stack.
func (_caca *PSStack )Empty (){*_caca =[]PSObject {}};func (_ba *PSBoolean )DebugString ()string {return _dg .Sprintf ("\u0062o\u006f\u006c\u003a\u0025\u0076",_ba .Val );};func (_bge *PSParser )parseNumber ()(PSObject ,error ){_fgad ,_abc :=_g .ParseNumber (_bge ._aeb );
if _abc !=nil {return nil ,_abc ;};switch _bgga :=_fgad .(type ){case *_g .PdfObjectFloat :return MakeReal (float64 (*_bgga )),nil ;case *_g .PdfObjectInteger :return MakeInteger (int (*_bgga )),nil ;};return nil ,_dg .Errorf ("\u0075n\u0068\u0061\u006e\u0064\u006c\u0065\u0064\u0020\u006e\u0075\u006db\u0065\u0072\u0020\u0074\u0079\u0070\u0065\u0020\u0025\u0054",_fgad );
};

// Append appends an object to the PSProgram.
func (_gee *PSProgram )Append (obj PSObject ){*_gee =append (*_gee ,obj )};var ErrRangeCheck =_e .New ("\u0072\u0061\u006e\u0067\u0065\u0020\u0063\u0068\u0065\u0063\u006b\u0020e\u0072\u0072\u006f\u0072");func (_ga *PSProgram )Duplicate ()PSObject {_eag :=&PSProgram {};
for _ ,_aeg :=range *_ga {_eag .Append (_aeg .Duplicate ());};return _eag ;};func (_ceg *PSOperand )mul (_cgbfe *PSStack )error {_aca ,_dac :=_cgbfe .Pop ();if _dac !=nil {return _dac ;};_efa ,_dac :=_cgbfe .Pop ();if _dac !=nil {return _dac ;};_cfc ,_fbe :=_aca .(*PSReal );
_deac ,_dbcd :=_aca .(*PSInteger );if !_fbe &&!_dbcd {return ErrTypeCheck ;};_bea ,_gfc :=_efa .(*PSReal );_fdd ,_gfa :=_efa .(*PSInteger );if !_gfc &&!_gfa {return ErrTypeCheck ;};if _dbcd &&_gfa {_bagad :=_deac .Val *_fdd .Val ;_fbfb :=_cgbfe .Push (MakeInteger (_bagad ));
return _fbfb ;};var _bgaf float64 ;if _fbe {_bgaf =_cfc .Val ;}else {_bgaf =float64 (_deac .Val );};if _gfc {_bgaf *=_bea .Val ;}else {_bgaf *=float64 (_fdd .Val );};_dac =_cgbfe .Push (MakeReal (_bgaf ));return _dac ;};func (_bgd *PSOperand )ge (_ce *PSStack )error {_dfe ,_bdc :=_ce .PopNumberAsFloat64 ();
if _bdc !=nil {return _bdc ;};_fbc ,_bdc :=_ce .PopNumberAsFloat64 ();if _bdc !=nil {return _bdc ;};if _b .Abs (_fbc -_dfe )< _gc {_bcg :=_ce .Push (MakeBool (true ));return _bcg ;}else if _fbc > _dfe {_dgd :=_ce .Push (MakeBool (true ));return _dgd ;}else {_gagd :=_ce .Push (MakeBool (false ));
return _gagd ;};};

// PSOperand represents a Postscript operand (text string).
type PSOperand string ;

// PSInteger represents an integer.
type PSInteger struct{Val int ;};func (_gccd *PSOperand )ne (_feg *PSStack )error {_dded :=_gccd .eq (_feg );if _dded !=nil {return _dded ;};_dded =_gccd .not (_feg );return _dded ;};func (_fec *PSOperand )Duplicate ()PSObject {_aee :=*_fec ;return &_aee };
func (_egg *PSOperand )lt (_gcbg *PSStack )error {_gfg ,_bgdb :=_gcbg .PopNumberAsFloat64 ();if _bgdb !=nil {return _bgdb ;};_dee ,_bgdb :=_gcbg .PopNumberAsFloat64 ();if _bgdb !=nil {return _bgdb ;};if _b .Abs (_dee -_gfg )< _gc {_dggbc :=_gcbg .Push (MakeBool (false ));
return _dggbc ;}else if _dee < _gfg {_adc :=_gcbg .Push (MakeBool (true ));return _adc ;}else {_dfg :=_gcbg .Push (MakeBool (false ));return _dfg ;};};func (_fd *PSReal )DebugString ()string {return _dg .Sprintf ("\u0072e\u0061\u006c\u003a\u0025\u002e\u0035f",_fd .Val );
};

// Parse parses the postscript and store as a program that can be executed.
func (_fcca *PSParser )Parse ()(*PSProgram ,error ){_fcca .skipSpaces ();_faab ,_bed :=_fcca ._aeb .Peek (2);if _bed !=nil {return nil ,_bed ;};if _faab [0]!='{'{return nil ,_e .New ("\u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0050\u0053\u0020\u0050\u0072\u006f\u0067\u0072\u0061\u006d\u0020\u006e\u006f\u0074\u0020\u0073t\u0061\u0072\u0074\u0069\u006eg\u0020\u0077i\u0074\u0068\u0020\u007b");
};_bae ,_bed :=_fcca .parseFunction ();if _bed !=nil &&_bed !=_a .EOF {return nil ,_bed ;};return _bae ,_bed ;};func (_beg *PSOperand )cvr (_edf *PSStack )error {_fed ,_dggb :=_edf .Pop ();if _dggb !=nil {return _dggb ;};if _fdf ,_fbb :=_fed .(*PSReal );
_fbb {_dggb =_edf .Push (MakeReal (_fdf .Val ));}else if _cgg ,_gdc :=_fed .(*PSInteger );_gdc {_dggb =_edf .Push (MakeReal (float64 (_cgg .Val )));}else {return ErrTypeCheck ;};return _dggb ;};func (_bb *PSInteger )String ()string {return _dg .Sprintf ("\u0025\u0064",_bb .Val )};
func (_cda *PSOperand )copy (_egd *PSStack )error {_bcf ,_ec :=_egd .PopInteger ();if _ec !=nil {return _ec ;};if _bcf < 0{return ErrRangeCheck ;};if _bcf > len (*_egd ){return ErrRangeCheck ;};*_egd =append (*_egd ,(*_egd )[len (*_egd )-_bcf :]...);return nil ;
};func (_gd *PSReal )Duplicate ()PSObject {_gce :=PSReal {};_gce .Val =_gd .Val ;return &_gce };func (_fdb *PSParser )parseOperand ()(*PSOperand ,error ){var _cgfdc []byte ;for {_ffb ,_aae :=_fdb ._aeb .Peek (1);if _aae !=nil {if _aae ==_a .EOF {break ;
};return nil ,_aae ;};if _g .IsDelimiter (_ffb [0]){break ;};if _g .IsWhiteSpace (_ffb [0]){break ;};_fgcf ,_ :=_fdb ._aeb .ReadByte ();_cgfdc =append (_cgfdc ,_fgcf );};if len (_cgfdc )==0{return nil ,_e .New ("\u0069\u006e\u0076al\u0069\u0064\u0020\u006f\u0070\u0065\u0072\u0061\u006e\u0064\u0020\u0028\u0065\u006d\u0070\u0074\u0079\u0029");
};return MakeOperand (string (_cgfdc )),nil ;};

// PSStack defines a stack of PSObjects. PSObjects can be pushed on or pull from the stack.
type PSStack []PSObject ;func (_ae *PSProgram )DebugString ()string {_fa :="\u007b\u0020";for _ ,_bgg :=range *_ae {_fa +=_bgg .DebugString ();_fa +="\u0020";};_fa +="\u007d";return _fa ;};var ErrTypeCheck =_e .New ("\u0074\u0079p\u0065\u0020\u0063h\u0065\u0063\u006b\u0020\u0065\u0072\u0072\u006f\u0072");
func (_ggdb *PSOperand )floor (_dedc *PSStack )error {_ffg ,_cab :=_dedc .Pop ();if _cab !=nil {return _cab ;};if _bbed ,_cbdd :=_ffg .(*PSReal );_cbdd {_cab =_dedc .Push (MakeReal (_b .Floor (_bbed .Val )));}else if _bga ,_agc :=_ffg .(*PSInteger );_agc {_cab =_dedc .Push (MakeInteger (_bga .Val ));
}else {return ErrTypeCheck ;};return _cab ;};func (_agg *PSOperand )sqrt (_gbg *PSStack )error {_dbba ,_age :=_gbg .PopNumberAsFloat64 ();if _age !=nil {return _age ;};if _dbba < 0{return ErrRangeCheck ;};_feef :=_b .Sqrt (_dbba );_age =_gbg .Push (MakeReal (_feef ));
return _age ;};func (_aaff *PSOperand )roll (_ggda *PSStack )error {_fddg ,_dbg :=_ggda .Pop ();if _dbg !=nil {return _dbg ;};_dgf ,_dbg :=_ggda .Pop ();if _dbg !=nil {return _dbg ;};_gcg ,_gbab :=_fddg .(*PSInteger );if !_gbab {return ErrTypeCheck ;};
_dbcdf ,_gbab :=_dgf .(*PSInteger );if !_gbab {return ErrTypeCheck ;};if _dbcdf .Val < 0{return ErrRangeCheck ;};if _dbcdf .Val ==0||_dbcdf .Val ==1{return nil ;};if _dbcdf .Val > len (*_ggda ){return ErrStackUnderflow ;};for _fadg :=0;_fadg < _bbcf (_gcg .Val );
_fadg ++{var _edgg []PSObject ;_edgg =(*_ggda )[len (*_ggda )-(_dbcdf .Val ):len (*_ggda )];if _gcg .Val > 0{_bgce :=_edgg [len (_edgg )-1];_edgg =append ([]PSObject {_bgce },_edgg [0:len (_edgg )-1]...);}else {_gbad :=_edgg [len (_edgg )-_dbcdf .Val ];
_edgg =append (_edgg [1:],_gbad );};_cfcg :=append ((*_ggda )[0:len (*_ggda )-_dbcdf .Val ],_edgg ...);_ggda =&_cfcg ;};return nil ;};