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

package mmr ;import (_cd "errors";_g "fmt";_e "github.com/unidoc/unipdf/v3/common";_aa "github.com/unidoc/unipdf/v3/internal/bitwise";_b "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_c "io";);var (_fd =[][3]int {{4,0x1,int (_gc )},{3,0x1,int (_db )},{1,0x1,int (_f )},{3,0x3,int (_ebd )},{6,0x3,int (_cgc )},{7,0x3,int (_ea )},{3,0x2,int (_ae )},{6,0x2,int (_aec )},{7,0x2,int (_ca )},{10,0xf,int (_bgf )},{12,0xf,int (_dgb )},{12,0x1,int (EOL )}};
_bfc =[][3]int {{4,0x07,2},{4,0x08,3},{4,0x0B,4},{4,0x0C,5},{4,0x0E,6},{4,0x0F,7},{5,0x12,128},{5,0x13,8},{5,0x14,9},{5,0x1B,64},{5,0x07,10},{5,0x08,11},{6,0x17,192},{6,0x18,1664},{6,0x2A,16},{6,0x2B,17},{6,0x03,13},{6,0x34,14},{6,0x35,15},{6,0x07,1},{6,0x08,12},{7,0x13,26},{7,0x17,21},{7,0x18,28},{7,0x24,27},{7,0x27,18},{7,0x28,24},{7,0x2B,25},{7,0x03,22},{7,0x37,256},{7,0x04,23},{7,0x08,20},{7,0xC,19},{8,0x12,33},{8,0x13,34},{8,0x14,35},{8,0x15,36},{8,0x16,37},{8,0x17,38},{8,0x1A,31},{8,0x1B,32},{8,0x02,29},{8,0x24,53},{8,0x25,54},{8,0x28,39},{8,0x29,40},{8,0x2A,41},{8,0x2B,42},{8,0x2C,43},{8,0x2D,44},{8,0x03,30},{8,0x32,61},{8,0x33,62},{8,0x34,63},{8,0x35,0},{8,0x36,320},{8,0x37,384},{8,0x04,45},{8,0x4A,59},{8,0x4B,60},{8,0x5,46},{8,0x52,49},{8,0x53,50},{8,0x54,51},{8,0x55,52},{8,0x58,55},{8,0x59,56},{8,0x5A,57},{8,0x5B,58},{8,0x64,448},{8,0x65,512},{8,0x67,640},{8,0x68,576},{8,0x0A,47},{8,0x0B,48},{9,0x01,_cac },{9,0x98,1472},{9,0x99,1536},{9,0x9A,1600},{9,0x9B,1728},{9,0xCC,704},{9,0xCD,768},{9,0xD2,832},{9,0xD3,896},{9,0xD4,960},{9,0xD5,1024},{9,0xD6,1088},{9,0xD7,1152},{9,0xD8,1216},{9,0xD9,1280},{9,0xDA,1344},{9,0xDB,1408},{10,0x01,_cac },{11,0x01,_cac },{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560}};
_dgg =[][3]int {{2,0x02,3},{2,0x03,2},{3,0x02,1},{3,0x03,4},{4,0x02,6},{4,0x03,5},{5,0x03,7},{6,0x04,9},{6,0x05,8},{7,0x04,10},{7,0x05,11},{7,0x07,12},{8,0x04,13},{8,0x07,14},{9,0x01,_cac },{9,0x18,15},{10,0x01,_cac },{10,0x17,16},{10,0x18,17},{10,0x37,0},{10,0x08,18},{10,0x0F,64},{11,0x01,_cac },{11,0x17,24},{11,0x18,25},{11,0x28,23},{11,0x37,22},{11,0x67,19},{11,0x68,20},{11,0x6C,21},{11,0x08,1792},{11,0x0C,1856},{11,0x0D,1920},{12,0x00,EOF },{12,0x01,EOL },{12,0x12,1984},{12,0x13,2048},{12,0x14,2112},{12,0x15,2176},{12,0x16,2240},{12,0x17,2304},{12,0x1C,2368},{12,0x1D,2432},{12,0x1E,2496},{12,0x1F,2560},{12,0x24,52},{12,0x27,55},{12,0x28,56},{12,0x2B,59},{12,0x2C,60},{12,0x33,320},{12,0x34,384},{12,0x35,448},{12,0x37,53},{12,0x38,54},{12,0x52,50},{12,0x53,51},{12,0x54,44},{12,0x55,45},{12,0x56,46},{12,0x57,47},{12,0x58,57},{12,0x59,58},{12,0x5A,61},{12,0x5B,256},{12,0x64,48},{12,0x65,49},{12,0x66,62},{12,0x67,63},{12,0x68,30},{12,0x69,31},{12,0x6A,32},{12,0x6B,33},{12,0x6C,40},{12,0x6D,41},{12,0xC8,128},{12,0xC9,192},{12,0xCA,26},{12,0xCB,27},{12,0xCC,28},{12,0xCD,29},{12,0xD2,34},{12,0xD3,35},{12,0xD4,36},{12,0xD5,37},{12,0xD6,38},{12,0xD7,39},{12,0xDA,42},{12,0xDB,43},{13,0x4A,640},{13,0x4B,704},{13,0x4C,768},{13,0x4D,832},{13,0x52,1280},{13,0x53,1344},{13,0x54,1408},{13,0x55,1472},{13,0x5A,1536},{13,0x5B,1600},{13,0x64,1664},{13,0x65,1728},{13,0x6C,512},{13,0x6D,576},{13,0x72,896},{13,0x73,960},{13,0x74,1024},{13,0x75,1088},{13,0x76,1152},{13,0x77,1216}};
);func _ga (_dg ,_ad int )int {if _dg > _ad {return _ad ;};return _dg ;};func (_ec *Decoder )detectAndSkipEOL ()error {for {_gag ,_dec :=_ec ._ag .uncompressGetCode (_ec ._cda );if _dec !=nil {return _dec ;};if _gag !=nil &&_gag ._cc ==EOL {_ec ._ag ._cbed +=_gag ._gd ;
}else {return nil ;};};};func (_age *Decoder )uncompress1d (_dee *runData ,_acc []int ,_cffd int )(int ,error ){var (_def =true ;_fge int ;_bd *code ;_bb int ;_dgf error ;);_ged :for _fge < _cffd {_fac :for {if _def {_bd ,_dgf =_dee .uncompressGetCode (_age ._fc );
if _dgf !=nil {return 0,_dgf ;};}else {_bd ,_dgf =_dee .uncompressGetCode (_age ._cf );if _dgf !=nil {return 0,_dgf ;};};_dee ._cbed +=_bd ._gd ;if _bd ._cc < 0{break _ged ;};_fge +=_bd ._cc ;if _bd ._cc < 64{_def =!_def ;_acc [_bb ]=_fge ;_bb ++;break _fac ;
};};};if _acc [_bb ]!=_cffd {_acc [_bb ]=_cffd ;};_da :=EOL ;if _bd !=nil &&_bd ._cc !=EOL {_da =_bb ;};return _da ,nil ;};func _faf (_ecg *_aa .SubstreamReader )(*runData ,error ){_fafb :=&runData {_ccc :_ecg ,_cbed :0,_gbe :1};_fca :=_ga (_de (_acad ,int (_ecg .Length ())),_gg );
_fafb ._ffca =make ([]byte ,_fca );if _bbed :=_fafb .fillBuffer (0);_bbed !=nil {if _bbed ==_c .EOF {_fafb ._ffca =make ([]byte ,10);_e .Log .Debug ("F\u0069\u006c\u006c\u0042uf\u0066e\u0072\u0020\u0066\u0061\u0069l\u0065\u0064\u003a\u0020\u0025\u0076",_bbed );
}else {return nil ,_bbed ;};};return _fafb ,nil ;};const (_gg int =1024<<7;_acad int =3;_gcg uint =24;);func (_cgb *runData )uncompressGetCode (_fgd []*code )(*code ,error ){return _cgb .uncompressGetCodeLittleEndian (_fgd );};const (EOF =-3;_cac =-2;EOL =-1;
_ff =8;_ead =(1<<_ff )-1;_ee =5;_dbg =(1<<_ee )-1;);type code struct{_gd int ;_ce int ;_cc int ;_bg []*code ;_eb bool ;};func New (r _aa .StreamReader ,width ,height int ,dataOffset ,dataLength int64 )(*Decoder ,error ){_ge :=&Decoder {_cae :width ,_cef :height };
_eac ,_gdd :=_aa .NewSubstreamReader (r ,uint64 (dataOffset ),uint64 (dataLength ));if _gdd !=nil {return nil ,_gdd ;};_ba ,_gdd :=_faf (_eac );if _gdd !=nil {return nil ,_gdd ;};_ge ._ag =_ba ;if _cb :=_ge .initTables ();_cb !=nil {return nil ,_cb ;};
return _ge ,nil ;};func (_bfa *runData )uncompressGetCodeLittleEndian (_deeb []*code )(*code ,error ){_fab ,_ddf :=_bfa .uncompressGetNextCodeLittleEndian ();if _ddf !=nil {_e .Log .Debug ("\u0055n\u0063\u006fm\u0070\u0072\u0065\u0073s\u0047\u0065\u0074N\u0065\u0078\u0074\u0043\u006f\u0064\u0065\u004c\u0069tt\u006c\u0065\u0045n\u0064\u0069a\u006e\u0020\u0066\u0061\u0069\u006ce\u0064\u003a \u0025\u0076",_ddf );
return nil ,_ddf ;};_fab &=0xffffff;_fbc :=_fab >>(_gcg -_ff );_gaf :=_deeb [_fbc ];if _gaf !=nil &&_gaf ._eb {_fbc =(_fab >>(_gcg -_ff -_ee ))&_dbg ;_gaf =_gaf ._bg [_fbc ];};return _gaf ,nil ;};type mmrCode int ;const (_gc mmrCode =iota ;_db ;_f ;_ebd ;
_cgc ;_ea ;_ae ;_aec ;_ca ;_bgf ;_dgb ;);func (_bgb *Decoder )uncompress2d (_af *runData ,_ced []int ,_bfe int ,_gb []int ,_ceb int )(int ,error ){var (_bbf int ;_cgg int ;_ebgd int ;_begf =true ;_bbe error ;_fdga *code ;);_ced [_bfe ]=_ceb ;_ced [_bfe +1]=_ceb ;
_ced [_bfe +2]=_ceb +1;_ced [_bfe +3]=_ceb +1;_egf :for _ebgd < _ceb {_fdga ,_bbe =_af .uncompressGetCode (_bgb ._cda );if _bbe !=nil {return EOL ,nil ;};if _fdga ==nil {_af ._cbed ++;break _egf ;};_af ._cbed +=_fdga ._gd ;switch mmrCode (_fdga ._cc ){case _f :_ebgd =_ced [_bbf ];
case _ebd :_ebgd =_ced [_bbf ]+1;case _ae :_ebgd =_ced [_bbf ]-1;case _db :for {var _dff []*code ;if _begf {_dff =_bgb ._fc ;}else {_dff =_bgb ._cf ;};_fdga ,_bbe =_af .uncompressGetCode (_dff );if _bbe !=nil {return 0,_bbe ;};if _fdga ==nil {break _egf ;
};_af ._cbed +=_fdga ._gd ;if _fdga ._cc < 64{if _fdga ._cc < 0{_gb [_cgg ]=_ebgd ;_cgg ++;_fdga =nil ;break _egf ;};_ebgd +=_fdga ._cc ;_gb [_cgg ]=_ebgd ;_cgg ++;break ;};_ebgd +=_fdga ._cc ;};_aff :=_ebgd ;_cdf :for {var _bda []*code ;if !_begf {_bda =_bgb ._fc ;
}else {_bda =_bgb ._cf ;};_fdga ,_bbe =_af .uncompressGetCode (_bda );if _bbe !=nil {return 0,_bbe ;};if _fdga ==nil {break _egf ;};_af ._cbed +=_fdga ._gd ;if _fdga ._cc < 64{if _fdga ._cc < 0{_gb [_cgg ]=_ebgd ;_cgg ++;break _egf ;};_ebgd +=_fdga ._cc ;
if _ebgd < _ceb ||_ebgd !=_aff {_gb [_cgg ]=_ebgd ;_cgg ++;};break _cdf ;};_ebgd +=_fdga ._cc ;};for _ebgd < _ceb &&_ced [_bbf ]<=_ebgd {_bbf +=2;};continue _egf ;case _gc :_bbf ++;_ebgd =_ced [_bbf ];_bbf ++;continue _egf ;case _cgc :_ebgd =_ced [_bbf ]+2;
case _aec :_ebgd =_ced [_bbf ]-2;case _ea :_ebgd =_ced [_bbf ]+3;case _ca :_ebgd =_ced [_bbf ]-3;default:if _af ._cbed ==12&&_fdga ._cc ==EOL {_af ._cbed =0;if _ ,_bbe =_bgb .uncompress1d (_af ,_ced ,_ceb );_bbe !=nil {return 0,_bbe ;};_af ._cbed ++;if _ ,_bbe =_bgb .uncompress1d (_af ,_gb ,_ceb );
_bbe !=nil {return 0,_bbe ;};_egc ,_gdc :=_bgb .uncompress1d (_af ,_ced ,_ceb );if _gdc !=nil {return EOF ,_gdc ;};_af ._cbed ++;return _egc ,nil ;};_ebgd =_ceb ;continue _egf ;};if _ebgd <=_ceb {_begf =!_begf ;_gb [_cgg ]=_ebgd ;_cgg ++;if _bbf > 0{_bbf --;
}else {_bbf ++;};for _ebgd < _ceb &&_ced [_bbf ]<=_ebgd {_bbf +=2;};};};if _gb [_cgg ]!=_ceb {_gb [_cgg ]=_ceb ;};if _fdga ==nil {return EOL ,nil ;};return _cgg ,nil ;};type runData struct{_ccc *_aa .SubstreamReader ;_cbed int ;_gbe int ;_cfd int ;_ffca []byte ;
_afg int ;_ebga int ;};type Decoder struct{_cae ,_cef int ;_ag *runData ;_fc []*code ;_cf []*code ;_cda []*code ;};func _de (_bf ,_dd int )int {if _bf < _dd {return _dd ;};return _bf ;};func (_ebe *Decoder )UncompressMMR ()(_gde *_b .Bitmap ,_bad error ){_gde =_b .New (_ebe ._cae ,_ebe ._cef );
_cff :=make ([]int ,_gde .Width +5);_ac :=make ([]int ,_gde .Width +5);_ac [0]=_gde .Width ;_baf :=1;var _ccd int ;for _aca :=0;_aca < _gde .Height ;_aca ++{_ccd ,_bad =_ebe .uncompress2d (_ebe ._ag ,_ac ,_baf ,_cff ,_gde .Width );if _bad !=nil {return nil ,_bad ;
};if _ccd ==EOF {break ;};if _ccd > 0{_bad =_ebe .fillBitmap (_gde ,_aca ,_cff ,_ccd );if _bad !=nil {return nil ,_bad ;};};_ac ,_cff =_cff ,_ac ;_baf =_ccd ;};if _bad =_ebe .detectAndSkipEOL ();_bad !=nil {return nil ,_bad ;};_ebe ._ag .align ();return _gde ,nil ;
};func (_gbb *runData )fillBuffer (_cgff int )error {_gbb ._afg =_cgff ;_ ,_aea :=_gbb ._ccc .Seek (int64 (_cgff ),_c .SeekStart );if _aea !=nil {if _aea ==_c .EOF {_e .Log .Debug ("\u0053\u0065\u0061\u006b\u0020\u0045\u004f\u0046");_gbb ._ebga =-1;}else {return _aea ;
};};if _aea ==nil {_gbb ._ebga ,_aea =_gbb ._ccc .Read (_gbb ._ffca );if _aea !=nil {if _aea ==_c .EOF {_e .Log .Trace ("\u0052\u0065\u0061\u0064\u0020\u0045\u004f\u0046");_gbb ._ebga =-1;}else {return _aea ;};};};if _gbb ._ebga > -1&&_gbb ._ebga < 3{for _gbb ._ebga < 3{_fff ,_ebdb :=_gbb ._ccc .ReadByte ();
if _ebdb !=nil {if _ebdb ==_c .EOF {_gbb ._ffca [_gbb ._ebga ]=0;}else {return _ebdb ;};}else {_gbb ._ffca [_gbb ._ebga ]=_fff &0xFF;};_gbb ._ebga ++;};};_gbb ._ebga -=3;if _gbb ._ebga < 0{_gbb ._ffca =make ([]byte ,len (_gbb ._ffca ));_gbb ._ebga =len (_gbb ._ffca )-3;
};return nil ;};func (_ffc *Decoder )fillBitmap (_addg *_b .Bitmap ,_cbe int ,_cbg []int ,_df int )error {var _ffe byte ;_caf :=0;_gaa :=_addg .GetByteIndex (_caf ,_cbe );for _fdg :=0;_fdg < _df ;_fdg ++{_gdde :=byte (1);_dc :=_cbg [_fdg ];if (_fdg &1)==0{_gdde =0;
};for _caf < _dc {_ffe =(_ffe <<1)|_gdde ;_caf ++;if (_caf &7)==0{if _fa :=_addg .SetByte (_gaa ,_ffe );_fa !=nil {return _fa ;};_gaa ++;_ffe =0;};};};if (_caf &7)!=0{_ffe <<=uint (8-(_caf &7));if _ebg :=_addg .SetByte (_gaa ,_ffe );_ebg !=nil {return _ebg ;
};};return nil ;};func (_eeg *runData )align (){_eeg ._cbed =((_eeg ._cbed +7)>>3)<<3};func (_cg *code )String ()string {return _g .Sprintf ("\u0025\u0064\u002f\u0025\u0064\u002f\u0025\u0064",_cg ._gd ,_cg ._ce ,_cg ._cc );};func (_dae *runData )uncompressGetNextCodeLittleEndian ()(int ,error ){_cdd :=_dae ._cbed -_dae ._gbe ;
if _cdd < 0||_cdd > 24{_ded :=(_dae ._cbed >>3)-_dae ._afg ;if _ded >=_dae ._ebga {_ded +=_dae ._afg ;if _bec :=_dae .fillBuffer (_ded );_bec !=nil {return 0,_bec ;};_ded -=_dae ._afg ;};_egg :=(uint32 (_dae ._ffca [_ded ]&0xFF)<<16)|(uint32 (_dae ._ffca [_ded +1]&0xFF)<<8)|(uint32 (_dae ._ffca [_ded +2]&0xFF));
_aab :=uint32 (_dae ._cbed &7);_egg <<=_aab ;_dae ._cfd =int (_egg );}else {_gdg :=_dae ._gbe &7;_dfd :=7-_gdg ;if _cdd <=_dfd {_dae ._cfd <<=uint (_cdd );}else {_fcf :=(_dae ._gbe >>3)+3-_dae ._afg ;if _fcf >=_dae ._ebga {_fcf +=_dae ._afg ;if _fe :=_dae .fillBuffer (_fcf );
_fe !=nil {return 0,_fe ;};_fcf -=_dae ._afg ;};_gdg =8-_gdg ;for {_dae ._cfd <<=uint (_gdg );_dae ._cfd |=int (uint (_dae ._ffca [_fcf ])&0xFF);_cdd -=_gdg ;_fcf ++;_gdg =8;if !(_cdd >=8){break ;};};_dae ._cfd <<=uint (_cdd );};};_dae ._gbe =_dae ._cbed ;
return _dae ._cfd ,nil ;};func (_edd *Decoder )initTables ()(_gab error ){if _edd ._fc ==nil {_edd ._fc ,_gab =_edd .createLittleEndianTable (_bfc );if _gab !=nil {return ;};_edd ._cf ,_gab =_edd .createLittleEndianTable (_dgg );if _gab !=nil {return ;
};_edd ._cda ,_gab =_edd .createLittleEndianTable (_fd );if _gab !=nil {return ;};};return nil ;};func (_ed *Decoder )createLittleEndianTable (_fb [][3]int )([]*code ,error ){_cgf :=make ([]*code ,_ead +1);for _fg :=0;_fg < len (_fb );_fg ++{_eg :=_be (_fb [_fg ]);
if _eg ._gd <=_ff {_ddc :=_ff -_eg ._gd ;_ab :=_eg ._ce <<uint (_ddc );for _eea :=(1<<uint (_ddc ))-1;_eea >=0;_eea --{_beg :=_ab |_eea ;_cgf [_beg ]=_eg ;};}else {_egd :=_eg ._ce >>uint (_eg ._gd -_ff );if _cgf [_egd ]==nil {var _bc =_be ([3]int {});_bc ._bg =make ([]*code ,_dbg +1);
_cgf [_egd ]=_bc ;};if _eg ._gd <=_ff +_ee {_add :=_ff +_ee -_eg ._gd ;_caed :=(_eg ._ce <<uint (_add ))&_dbg ;_cgf [_egd ]._eb =true ;for _cba :=(1<<uint (_add ))-1;_cba >=0;_cba --{_cgf [_egd ]._bg [_caed |_cba ]=_eg ;};}else {return nil ,_cd .New ("\u0043\u006f\u0064\u0065\u0020\u0074a\u0062\u006c\u0065\u0020\u006f\u0076\u0065\u0072\u0066\u006c\u006f\u0077\u0020i\u006e\u0020\u004d\u004d\u0052\u0044\u0065c\u006f\u0064\u0065\u0072");
};};};return _cgf ,nil ;};func _be (_d [3]int )*code {return &code {_gd :_d [0],_ce :_d [1],_cc :_d [2]}};