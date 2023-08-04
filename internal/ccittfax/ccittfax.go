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

package ccittfax ;import (_a "errors";_g "github.com/unidoc/unipdf/v3/internal/bitwise";_da "io";_d "math";);var (_dac =_a .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074\u0065\u0064\u0020\u0052T\u0043");_fdg =_a .New ("\u0063\u0063\u0069\u0074tf\u0061\u0078\u0020\u0045\u004f\u004c\u0020\u006e\u006f\u0074\u0020\u0066\u006f\u0075n\u0064");
);func init (){_e =&treeNode {_dgb :true ,_cgbg :_b };_cg =&treeNode {_cgbg :_ge ,_cae :_e };_cg ._fcf =_cg ;_f =&tree {_cdb :&treeNode {}};if _ed :=_f .fillWithNode (12,0,_cg );_ed !=nil {panic (_ed .Error ());};if _ae :=_f .fillWithNode (12,1,_e );_ae !=nil {panic (_ae .Error ());
};_ea =&tree {_cdb :&treeNode {}};for _gc :=0;_gc < len (_eg );_gc ++{for _ag :=0;_ag < len (_eg [_gc ]);_ag ++{if _fef :=_ea .fill (_gc +2,int (_eg [_gc ][_ag ]),int (_ba [_gc ][_ag ]));_fef !=nil {panic (_fef .Error ());};};};if _ga :=_ea .fillWithNode (12,0,_cg );
_ga !=nil {panic (_ga .Error ());};if _be :=_ea .fillWithNode (12,1,_e );_be !=nil {panic (_be .Error ());};_ca =&tree {_cdb :&treeNode {}};for _ee :=0;_ee < len (_begg );_ee ++{for _bf :=0;_bf < len (_begg [_ee ]);_bf ++{if _agb :=_ca .fill (_ee +4,int (_begg [_ee ][_bf ]),int (_cgb [_ee ][_bf ]));
_agb !=nil {panic (_agb .Error ());};};};if _dc :=_ca .fillWithNode (12,0,_cg );_dc !=nil {panic (_dc .Error ());};if _eae :=_ca .fillWithNode (12,1,_e );_eae !=nil {panic (_eae .Error ());};_eac =&tree {_cdb :&treeNode {}};if _aee :=_eac .fill (4,1,_fg );
_aee !=nil {panic (_aee .Error ());};if _cge :=_eac .fill (3,1,_ec );_cge !=nil {panic (_cge .Error ());};if _bc :=_eac .fill (1,1,0);_bc !=nil {panic (_bc .Error ());};if _ede :=_eac .fill (3,3,1);_ede !=nil {panic (_ede .Error ());};if _cc :=_eac .fill (6,3,2);
_cc !=nil {panic (_cc .Error ());};if _cab :=_eac .fill (7,3,3);_cab !=nil {panic (_cab .Error ());};if _beg :=_eac .fill (3,2,-1);_beg !=nil {panic (_beg .Error ());};if _bea :=_eac .fill (6,2,-2);_bea !=nil {panic (_bea .Error ());};if _fb :=_eac .fill (7,2,-3);
_fb !=nil {panic (_fb .Error ());};};var _begg =[...][]uint16 {{0x7,0x8,0xb,0xc,0xe,0xf},{0x12,0x13,0x14,0x1b,0x7,0x8},{0x17,0x18,0x2a,0x2b,0x3,0x34,0x35,0x7,0x8},{0x13,0x17,0x18,0x24,0x27,0x28,0x2b,0x3,0x37,0x4,0x8,0xc},{0x12,0x13,0x14,0x15,0x16,0x17,0x1a,0x1b,0x2,0x24,0x25,0x28,0x29,0x2a,0x2b,0x2c,0x2d,0x3,0x32,0x33,0x34,0x35,0x36,0x37,0x4,0x4a,0x4b,0x5,0x52,0x53,0x54,0x55,0x58,0x59,0x5a,0x5b,0x64,0x65,0x67,0x68,0xa,0xb},{0x98,0x99,0x9a,0x9b,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xd8,0xd9,0xda,0xdb},{},{0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f}};
func (_bbe *Decoder )tryFetchEOL1 ()(bool ,error ){_bab ,_fbg :=_bbe ._bcg .ReadBits (13);if _fbg !=nil {return false ,_fbg ;};return _bab ==0x3,nil ;};type tree struct{_cdb *treeNode };func (_af *Decoder )decodeRowType2 ()error {if _af ._ecg {_af ._bcg .Align ();
};if _bed :=_af .decode1D ();_bed !=nil {return _bed ;};return nil ;};var (_de map[int ]code ;_ef map[int ]code ;_dcg map[int ]code ;_dd map[int ]code ;_gg map[int ]code ;_eec map[int ]byte ;_fbf =code {Code :1<<4,BitsWritten :12};_fd =code {Code :3<<3,BitsWritten :13};
_bb =code {Code :2<<3,BitsWritten :13};_cf =code {Code :1<<12,BitsWritten :4};_gd =code {Code :1<<13,BitsWritten :3};_dg =code {Code :1<<15,BitsWritten :1};_bbf =code {Code :3<<13,BitsWritten :3};_fbfa =code {Code :3<<10,BitsWritten :6};_edd =code {Code :3<<9,BitsWritten :7};
_ecf =code {Code :2<<13,BitsWritten :3};_ege =code {Code :2<<10,BitsWritten :6};_fc =code {Code :2<<9,BitsWritten :7};);type tiffType int ;func init (){_de =make (map[int ]code );_de [0]=code {Code :13<<8|3<<6,BitsWritten :10};_de [1]=code {Code :2<<(5+8),BitsWritten :3};
_de [2]=code {Code :3<<(6+8),BitsWritten :2};_de [3]=code {Code :2<<(6+8),BitsWritten :2};_de [4]=code {Code :3<<(5+8),BitsWritten :3};_de [5]=code {Code :3<<(4+8),BitsWritten :4};_de [6]=code {Code :2<<(4+8),BitsWritten :4};_de [7]=code {Code :3<<(3+8),BitsWritten :5};
_de [8]=code {Code :5<<(2+8),BitsWritten :6};_de [9]=code {Code :4<<(2+8),BitsWritten :6};_de [10]=code {Code :4<<(1+8),BitsWritten :7};_de [11]=code {Code :5<<(1+8),BitsWritten :7};_de [12]=code {Code :7<<(1+8),BitsWritten :7};_de [13]=code {Code :4<<8,BitsWritten :8};
_de [14]=code {Code :7<<8,BitsWritten :8};_de [15]=code {Code :12<<8,BitsWritten :9};_de [16]=code {Code :5<<8|3<<6,BitsWritten :10};_de [17]=code {Code :6<<8,BitsWritten :10};_de [18]=code {Code :2<<8,BitsWritten :10};_de [19]=code {Code :12<<8|7<<5,BitsWritten :11};
_de [20]=code {Code :13<<8,BitsWritten :11};_de [21]=code {Code :13<<8|4<<5,BitsWritten :11};_de [22]=code {Code :6<<8|7<<5,BitsWritten :11};_de [23]=code {Code :5<<8,BitsWritten :11};_de [24]=code {Code :2<<8|7<<5,BitsWritten :11};_de [25]=code {Code :3<<8,BitsWritten :11};
_de [26]=code {Code :12<<8|10<<4,BitsWritten :12};_de [27]=code {Code :12<<8|11<<4,BitsWritten :12};_de [28]=code {Code :12<<8|12<<4,BitsWritten :12};_de [29]=code {Code :12<<8|13<<4,BitsWritten :12};_de [30]=code {Code :6<<8|8<<4,BitsWritten :12};_de [31]=code {Code :6<<8|9<<4,BitsWritten :12};
_de [32]=code {Code :6<<8|10<<4,BitsWritten :12};_de [33]=code {Code :6<<8|11<<4,BitsWritten :12};_de [34]=code {Code :13<<8|2<<4,BitsWritten :12};_de [35]=code {Code :13<<8|3<<4,BitsWritten :12};_de [36]=code {Code :13<<8|4<<4,BitsWritten :12};_de [37]=code {Code :13<<8|5<<4,BitsWritten :12};
_de [38]=code {Code :13<<8|6<<4,BitsWritten :12};_de [39]=code {Code :13<<8|7<<4,BitsWritten :12};_de [40]=code {Code :6<<8|12<<4,BitsWritten :12};_de [41]=code {Code :6<<8|13<<4,BitsWritten :12};_de [42]=code {Code :13<<8|10<<4,BitsWritten :12};_de [43]=code {Code :13<<8|11<<4,BitsWritten :12};
_de [44]=code {Code :5<<8|4<<4,BitsWritten :12};_de [45]=code {Code :5<<8|5<<4,BitsWritten :12};_de [46]=code {Code :5<<8|6<<4,BitsWritten :12};_de [47]=code {Code :5<<8|7<<4,BitsWritten :12};_de [48]=code {Code :6<<8|4<<4,BitsWritten :12};_de [49]=code {Code :6<<8|5<<4,BitsWritten :12};
_de [50]=code {Code :5<<8|2<<4,BitsWritten :12};_de [51]=code {Code :5<<8|3<<4,BitsWritten :12};_de [52]=code {Code :2<<8|4<<4,BitsWritten :12};_de [53]=code {Code :3<<8|7<<4,BitsWritten :12};_de [54]=code {Code :3<<8|8<<4,BitsWritten :12};_de [55]=code {Code :2<<8|7<<4,BitsWritten :12};
_de [56]=code {Code :2<<8|8<<4,BitsWritten :12};_de [57]=code {Code :5<<8|8<<4,BitsWritten :12};_de [58]=code {Code :5<<8|9<<4,BitsWritten :12};_de [59]=code {Code :2<<8|11<<4,BitsWritten :12};_de [60]=code {Code :2<<8|12<<4,BitsWritten :12};_de [61]=code {Code :5<<8|10<<4,BitsWritten :12};
_de [62]=code {Code :6<<8|6<<4,BitsWritten :12};_de [63]=code {Code :6<<8|7<<4,BitsWritten :12};_ef =make (map[int ]code );_ef [0]=code {Code :53<<8,BitsWritten :8};_ef [1]=code {Code :7<<(2+8),BitsWritten :6};_ef [2]=code {Code :7<<(4+8),BitsWritten :4};
_ef [3]=code {Code :8<<(4+8),BitsWritten :4};_ef [4]=code {Code :11<<(4+8),BitsWritten :4};_ef [5]=code {Code :12<<(4+8),BitsWritten :4};_ef [6]=code {Code :14<<(4+8),BitsWritten :4};_ef [7]=code {Code :15<<(4+8),BitsWritten :4};_ef [8]=code {Code :19<<(3+8),BitsWritten :5};
_ef [9]=code {Code :20<<(3+8),BitsWritten :5};_ef [10]=code {Code :7<<(3+8),BitsWritten :5};_ef [11]=code {Code :8<<(3+8),BitsWritten :5};_ef [12]=code {Code :8<<(2+8),BitsWritten :6};_ef [13]=code {Code :3<<(2+8),BitsWritten :6};_ef [14]=code {Code :52<<(2+8),BitsWritten :6};
_ef [15]=code {Code :53<<(2+8),BitsWritten :6};_ef [16]=code {Code :42<<(2+8),BitsWritten :6};_ef [17]=code {Code :43<<(2+8),BitsWritten :6};_ef [18]=code {Code :39<<(1+8),BitsWritten :7};_ef [19]=code {Code :12<<(1+8),BitsWritten :7};_ef [20]=code {Code :8<<(1+8),BitsWritten :7};
_ef [21]=code {Code :23<<(1+8),BitsWritten :7};_ef [22]=code {Code :3<<(1+8),BitsWritten :7};_ef [23]=code {Code :4<<(1+8),BitsWritten :7};_ef [24]=code {Code :40<<(1+8),BitsWritten :7};_ef [25]=code {Code :43<<(1+8),BitsWritten :7};_ef [26]=code {Code :19<<(1+8),BitsWritten :7};
_ef [27]=code {Code :36<<(1+8),BitsWritten :7};_ef [28]=code {Code :24<<(1+8),BitsWritten :7};_ef [29]=code {Code :2<<8,BitsWritten :8};_ef [30]=code {Code :3<<8,BitsWritten :8};_ef [31]=code {Code :26<<8,BitsWritten :8};_ef [32]=code {Code :27<<8,BitsWritten :8};
_ef [33]=code {Code :18<<8,BitsWritten :8};_ef [34]=code {Code :19<<8,BitsWritten :8};_ef [35]=code {Code :20<<8,BitsWritten :8};_ef [36]=code {Code :21<<8,BitsWritten :8};_ef [37]=code {Code :22<<8,BitsWritten :8};_ef [38]=code {Code :23<<8,BitsWritten :8};
_ef [39]=code {Code :40<<8,BitsWritten :8};_ef [40]=code {Code :41<<8,BitsWritten :8};_ef [41]=code {Code :42<<8,BitsWritten :8};_ef [42]=code {Code :43<<8,BitsWritten :8};_ef [43]=code {Code :44<<8,BitsWritten :8};_ef [44]=code {Code :45<<8,BitsWritten :8};
_ef [45]=code {Code :4<<8,BitsWritten :8};_ef [46]=code {Code :5<<8,BitsWritten :8};_ef [47]=code {Code :10<<8,BitsWritten :8};_ef [48]=code {Code :11<<8,BitsWritten :8};_ef [49]=code {Code :82<<8,BitsWritten :8};_ef [50]=code {Code :83<<8,BitsWritten :8};
_ef [51]=code {Code :84<<8,BitsWritten :8};_ef [52]=code {Code :85<<8,BitsWritten :8};_ef [53]=code {Code :36<<8,BitsWritten :8};_ef [54]=code {Code :37<<8,BitsWritten :8};_ef [55]=code {Code :88<<8,BitsWritten :8};_ef [56]=code {Code :89<<8,BitsWritten :8};
_ef [57]=code {Code :90<<8,BitsWritten :8};_ef [58]=code {Code :91<<8,BitsWritten :8};_ef [59]=code {Code :74<<8,BitsWritten :8};_ef [60]=code {Code :75<<8,BitsWritten :8};_ef [61]=code {Code :50<<8,BitsWritten :8};_ef [62]=code {Code :51<<8,BitsWritten :8};
_ef [63]=code {Code :52<<8,BitsWritten :8};_dcg =make (map[int ]code );_dcg [64]=code {Code :3<<8|3<<6,BitsWritten :10};_dcg [128]=code {Code :12<<8|8<<4,BitsWritten :12};_dcg [192]=code {Code :12<<8|9<<4,BitsWritten :12};_dcg [256]=code {Code :5<<8|11<<4,BitsWritten :12};
_dcg [320]=code {Code :3<<8|3<<4,BitsWritten :12};_dcg [384]=code {Code :3<<8|4<<4,BitsWritten :12};_dcg [448]=code {Code :3<<8|5<<4,BitsWritten :12};_dcg [512]=code {Code :3<<8|12<<3,BitsWritten :13};_dcg [576]=code {Code :3<<8|13<<3,BitsWritten :13};
_dcg [640]=code {Code :2<<8|10<<3,BitsWritten :13};_dcg [704]=code {Code :2<<8|11<<3,BitsWritten :13};_dcg [768]=code {Code :2<<8|12<<3,BitsWritten :13};_dcg [832]=code {Code :2<<8|13<<3,BitsWritten :13};_dcg [896]=code {Code :3<<8|18<<3,BitsWritten :13};
_dcg [960]=code {Code :3<<8|19<<3,BitsWritten :13};_dcg [1024]=code {Code :3<<8|20<<3,BitsWritten :13};_dcg [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_dcg [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_dcg [1216]=code {Code :119<<3,BitsWritten :13};
_dcg [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_dcg [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_dcg [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_dcg [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_dcg [1536]=code {Code :2<<8|26<<3,BitsWritten :13};
_dcg [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_dcg [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_dcg [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_dd =make (map[int ]code );_dd [64]=code {Code :27<<(3+8),BitsWritten :5};_dd [128]=code {Code :18<<(3+8),BitsWritten :5};
_dd [192]=code {Code :23<<(2+8),BitsWritten :6};_dd [256]=code {Code :55<<(1+8),BitsWritten :7};_dd [320]=code {Code :54<<8,BitsWritten :8};_dd [384]=code {Code :55<<8,BitsWritten :8};_dd [448]=code {Code :100<<8,BitsWritten :8};_dd [512]=code {Code :101<<8,BitsWritten :8};
_dd [576]=code {Code :104<<8,BitsWritten :8};_dd [640]=code {Code :103<<8,BitsWritten :8};_dd [704]=code {Code :102<<8,BitsWritten :9};_dd [768]=code {Code :102<<8|1<<7,BitsWritten :9};_dd [832]=code {Code :105<<8,BitsWritten :9};_dd [896]=code {Code :105<<8|1<<7,BitsWritten :9};
_dd [960]=code {Code :106<<8,BitsWritten :9};_dd [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_dd [1088]=code {Code :107<<8,BitsWritten :9};_dd [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_dd [1216]=code {Code :108<<8,BitsWritten :9};_dd [1280]=code {Code :108<<8|1<<7,BitsWritten :9};
_dd [1344]=code {Code :109<<8,BitsWritten :9};_dd [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_dd [1472]=code {Code :76<<8,BitsWritten :9};_dd [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_dd [1600]=code {Code :77<<8,BitsWritten :9};_dd [1664]=code {Code :24<<(2+8),BitsWritten :6};
_dd [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_gg =make (map[int ]code );_gg [1792]=code {Code :1<<8,BitsWritten :11};_gg [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_gg [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_gg [1984]=code {Code :1<<8|2<<4,BitsWritten :12};
_gg [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_gg [2112]=code {Code :1<<8|4<<4,BitsWritten :12};_gg [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_gg [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_gg [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_gg [2368]=code {Code :1<<8|12<<4,BitsWritten :12};
_gg [2432]=code {Code :1<<8|13<<4,BitsWritten :12};_gg [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_gg [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_eec =make (map[int ]byte );_eec [0]=0xFF;_eec [1]=0xFE;_eec [2]=0xFC;_eec [3]=0xF8;_eec [4]=0xF0;_eec [5]=0xE0;
_eec [6]=0xC0;_eec [7]=0x80;_eec [8]=0x00;};func (_edb *Decoder )decodeRow ()(_efd error ){if !_edb ._egd &&_edb ._ff > 0&&_edb ._ff ==_edb ._ddf {return _da .EOF ;};switch _edb ._cgd {case _cd :_efd =_edb .decodeRowType2 ();case _deb :_efd =_edb .decodeRowType4 ();
case _cabb :_efd =_edb .decodeRowType6 ();};if _efd !=nil {return _efd ;};_gee :=0;_gbd :=true ;_edb ._cff =0;for _fca :=0;_fca < _edb ._gab ;_fca ++{_feg :=_edb ._fdd ;if _fca !=_edb ._gab {_feg =_edb ._ab [_fca ];};if _feg > _edb ._fdd {_feg =_edb ._fdd ;
};_gf :=_gee /8;for _gee %8!=0&&_feg -_gee > 0{var _aa byte ;if !_gbd {_aa =1<<uint (7-(_gee %8));};_edb ._bfc [_gf ]|=_aa ;_gee ++;};if _gee %8==0{_gf =_gee /8;var _edec byte ;if !_gbd {_edec =0xff;};for _feg -_gee > 7{_edb ._bfc [_gf ]=_edec ;_gee +=8;
_gf ++;};};for _feg -_gee > 0{if _gee %8==0{_edb ._bfc [_gf ]=0;};var _abf byte ;if !_gbd {_abf =1<<uint (7-(_gee %8));};_edb ._bfc [_gf ]|=_abf ;_gee ++;};_gbd =!_gbd ;};if _gee !=_edb ._fdd {return _a .New ("\u0073\u0075\u006d\u0020\u006f\u0066 \u0072\u0075\u006e\u002d\u006c\u0065\u006e\u0067\u0074\u0068\u0073\u0020\u0064\u006f\u0065\u0073\u0020\u006e\u006f\u0074 \u0065\u0071\u0075\u0061\u006c\u0020\u0073\u0063\u0061\u006e\u0020\u006c\u0069\u006ee\u0020w\u0069\u0064\u0074\u0068");
};_edb ._fgc =(_gee +7)/8;_edb ._ddf ++;return nil ;};func (_dbdd *Decoder )decodeRun (_ddd *tree )(int ,error ){var _gfa int ;_dacb :=_ddd ._cdb ;for {_bd ,_caca :=_dbdd ._bcg .ReadBool ();if _caca !=nil {return 0,_caca ;};_dacb =_dacb .walk (_bd );if _dacb ==nil {return 0,_a .New ("\u0075\u006e\u006bno\u0077\u006e\u0020\u0063\u006f\u0064\u0065\u0020\u0069n\u0020H\u0075f\u0066m\u0061\u006e\u0020\u0052\u004c\u0045\u0020\u0073\u0074\u0072\u0065\u0061\u006d");
};if _dacb ._dgb {_gfa +=_dacb ._cgbg ;switch {case _dacb ._cgbg >=64:_dacb =_ddd ._cdb ;case _dacb ._cgbg >=0:return _gfa ,nil ;default:return _dbdd ._fdd ,nil ;};};};};func (_efadg *Decoder )decode1D ()error {var (_bcgg int ;_gca error ;);_cdg :=true ;
_efadg ._gab =0;for {var _egec int ;if _cdg {_egec ,_gca =_efadg .decodeRun (_ca );}else {_egec ,_gca =_efadg .decodeRun (_ea );};if _gca !=nil {return _gca ;};_bcgg +=_egec ;_efadg ._ab [_efadg ._gab ]=_bcgg ;_efadg ._gab ++;_cdg =!_cdg ;if _bcgg >=_efadg ._fdd {break ;
};};return nil ;};func _adbg (_bdd int ,_ded bool )(code ,int ,bool ){if _bdd < 64{if _ded {return _ef [_bdd ],0,true ;};return _de [_bdd ],0,true ;};_cea :=_bdd /64;if _cea > 40{return _gg [2560],_bdd -2560,false ;};if _cea > 27{return _gg [_cea *64],_bdd -_cea *64,false ;
};if _ded {return _dd [_cea *64],_bdd -_cea *64,false ;};return _dcg [_cea *64],_bdd -_cea *64,false ;};func (_gabf *Decoder )decodeRowType6 ()error {if _gabf ._ecg {_gabf ._bcg .Align ();};if _gabf ._egd {_gabf ._bcg .Mark ();_daf ,_ce :=_gabf .tryFetchEOL ();
if _ce !=nil {return _ce ;};if _daf {_daf ,_ce =_gabf .tryFetchEOL ();if _ce !=nil {return _ce ;};if _daf {return _da .EOF ;};};_gabf ._bcg .Reset ();};return _gabf .decode2D ();};var _cgb =[...][]uint16 {{2,3,4,5,6,7},{128,8,9,64,10,11},{192,1664,16,17,13,14,15,1,12},{26,21,28,27,18,24,25,22,256,23,20,19},{33,34,35,36,37,38,31,32,29,53,54,39,40,41,42,43,44,30,61,62,63,0,320,384,45,59,60,46,49,50,51,52,55,56,57,58,448,512,640,576,47,48},{1472,1536,1600,1728,704,768,832,896,960,1024,1088,1152,1216,1280,1344,1408},{},{1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560}};
const (_ tiffType =iota ;_cd ;_deb ;_cabb ;);func (_gfg *Decoder )looseFetchEOL ()(bool ,error ){_ged ,_fdf :=_gfg ._bcg .ReadBits (12);if _fdf !=nil {return false ,_fdf ;};switch _ged {case 0x1:return true ,nil ;case 0x0:for {_ggg ,_edbc :=_gfg ._bcg .ReadBool ();
if _edbc !=nil {return false ,_edbc ;};if _ggg {return true ,nil ;};};default:return false ,nil ;};};func _acc (_dea int )([]byte ,int ){var _ffbb []byte ;for _eggf :=0;_eggf < 2;_eggf ++{_ffbb ,_dea =_cfc (_ffbb ,_dea ,_fbf );};return _ffbb ,_dea %8;};
func _gbdg (_caa ,_fbdb int )code {var _cdce code ;switch _fbdb -_caa {case -1:_cdce =_bbf ;case -2:_cdce =_fbfa ;case -3:_cdce =_edd ;case 0:_cdce =_dg ;case 1:_cdce =_ecf ;case 2:_cdce =_ege ;case 3:_cdce =_fc ;};return _cdce ;};func (_eag *Decoder )tryFetchEOL ()(bool ,error ){_dgd ,_cffe :=_eag ._bcg .ReadBits (12);
if _cffe !=nil {return false ,_cffe ;};return _dgd ==0x1,nil ;};func (_ffae *Decoder )decodeRowType4 ()error {if !_ffae ._fda {return _ffae .decoderRowType41D ();};if _ffae ._ecg {_ffae ._bcg .Align ();};_ffae ._bcg .Mark ();_agfb ,_cfa :=_ffae .tryFetchEOL ();
if _cfa !=nil {return _cfa ;};if !_agfb &&_ffae ._dde {_ffae ._efa ++;if _ffae ._efa > _ffae ._cgea {return _fdg ;};_ffae ._bcg .Reset ();};if !_agfb {_ffae ._bcg .Reset ();};_gac ,_cfa :=_ffae ._bcg .ReadBool ();if _cfa !=nil {return _cfa ;};if _gac {if _agfb &&_ffae ._egd {if _cfa =_ffae .tryFetchRTC2D ();
_cfa !=nil {return _cfa ;};};_cfa =_ffae .decode1D ();}else {_cfa =_ffae .decode2D ();};if _cfa !=nil {return _cfa ;};return nil ;};func (_adf *Decoder )decode2D ()error {_adf ._ad =_adf ._gab ;_adf ._ab ,_adf ._bee =_adf ._bee ,_adf ._ab ;_cbb :=true ;
var (_cfg bool ;_fa int ;_fccb error ;);_adf ._gab =0;_cbe :for _fa < _adf ._fdd {_dcc :=_eac ._cdb ;for {_cfg ,_fccb =_adf ._bcg .ReadBool ();if _fccb !=nil {return _fccb ;};_dcc =_dcc .walk (_cfg );if _dcc ==nil {continue _cbe ;};if !_dcc ._dgb {continue ;
};switch _dcc ._cgbg {case _ec :var _afc int ;if _cbb {_afc ,_fccb =_adf .decodeRun (_ca );}else {_afc ,_fccb =_adf .decodeRun (_ea );};if _fccb !=nil {return _fccb ;};_fa +=_afc ;_adf ._ab [_adf ._gab ]=_fa ;_adf ._gab ++;if _cbb {_afc ,_fccb =_adf .decodeRun (_ea );
}else {_afc ,_fccb =_adf .decodeRun (_ca );};if _fccb !=nil {return _fccb ;};_fa +=_afc ;_adf ._ab [_adf ._gab ]=_fa ;_adf ._gab ++;case _fg :_afd :=_adf .getNextChangingElement (_fa ,_cbb )+1;if _afd >=_adf ._ad {_fa =_adf ._fdd ;}else {_fa =_adf ._bee [_afd ];
};default:_dfa :=_adf .getNextChangingElement (_fa ,_cbb );if _dfa >=_adf ._ad ||_dfa ==-1{_fa =_adf ._fdd +_dcc ._cgbg ;}else {_fa =_adf ._bee [_dfa ]+_dcc ._cgbg ;};_adf ._ab [_adf ._gab ]=_fa ;_adf ._gab ++;_cbb =!_cbb ;};continue _cbe ;};};return nil ;
};var (_gdd byte =1;_gea byte =0;);func _gagg (_bfg []byte ,_dfed int )([]byte ,int ){return _cfc (_bfg ,_dfed ,_cf )};func (_ffb *Decoder )tryFetchRTC2D ()(_cef error ){_ffb ._bcg .Mark ();var _gdc bool ;for _gfb :=0;_gfb < 5;_gfb ++{_gdc ,_cef =_ffb .tryFetchEOL1 ();
if _cef !=nil {if _a .Is (_cef ,_da .EOF ){if _gfb ==0{break ;};return _dac ;};};if _gdc {continue ;};if _gfb > 0{return _dac ;};break ;};if _gdc {return _da .EOF ;};_ffb ._bcg .Reset ();return _cef ;};func _adfc (_ddeb ,_ggfa []byte ,_baeb int ,_edcg bool )int {_ddfe :=_beeb (_ggfa ,_baeb );
if _ddfe < len (_ggfa )&&(_baeb ==-1&&_ggfa [_ddfe ]==_gdd ||_baeb >=0&&_baeb < len (_ddeb )&&_ddeb [_baeb ]==_ggfa [_ddfe ]||_baeb >=len (_ddeb )&&_edcg &&_ggfa [_ddfe ]==_gdd ||_baeb >=len (_ddeb )&&!_edcg &&_ggfa [_ddfe ]==_gea ){_ddfe =_beeb (_ggfa ,_ddfe );
};return _ddfe ;};type DecodeOptions struct{Columns int ;Rows int ;K int ;EncodedByteAligned bool ;BlackIsOne bool ;EndOfBlock bool ;EndOfLine bool ;DamagedRowsBeforeError int ;};func (_dfc *Encoder )encodeG4 (_afde [][]byte )[]byte {_fac :=make ([][]byte ,len (_afde ));
copy (_fac ,_afde );_fac =_gcd (_fac );var _fccf []byte ;var _baed int ;for _afe :=1;_afe < len (_fac );_afe ++{if _dfc .Rows > 0&&!_dfc .EndOfBlock &&_afe ==(_dfc .Rows +1){break ;};var _cecd []byte ;var _aac ,_ecd ,_fcg int ;_ac :=_baed ;_ffc :=-1;for _ffc < len (_fac [_afe ]){_aac =_beeb (_fac [_afe ],_ffc );
_ecd =_cbbb (_fac [_afe ],_fac [_afe -1],_ffc );_fcg =_beeb (_fac [_afe -1],_ecd );if _fcg < _aac {_cecd ,_ac =_cfc (_cecd ,_ac ,_cf );_ffc =_fcg ;}else {if _d .Abs (float64 (_ecd -_aac ))> 3{_cecd ,_ac ,_ffc =_gef (_fac [_afe ],_cecd ,_ac ,_ffc ,_aac );
}else {_cecd ,_ac =_fbd (_cecd ,_ac ,_aac ,_ecd );_ffc =_aac ;};};};_fccf =_dfc .appendEncodedRow (_fccf ,_cecd ,_baed );if _dfc .EncodedByteAlign {_ac =0;};_baed =_ac %8;};if _dfc .EndOfBlock {_ggda ,_ :=_acc (_baed );_fccf =_dfc .appendEncodedRow (_fccf ,_ggda ,_baed );
};return _fccf ;};func _gcd (_dfd [][]byte )[][]byte {_gfd :=make ([]byte ,len (_dfd [0]));for _gaca :=range _gfd {_gfd [_gaca ]=_gdd ;};_dfd =append (_dfd ,[]byte {});for _afg :=len (_dfd )-1;_afg > 0;_afg --{_dfd [_afg ]=_dfd [_afg -1];};_dfd [0]=_gfd ;
return _dfd ;};func _adfd (_egad []byte ,_gcfce int ,_adff int ,_fcb bool )([]byte ,int ){var (_fecf code ;_gegc bool ;);for !_gegc {_fecf ,_adff ,_gegc =_adbg (_adff ,_fcb );_egad ,_gcfce =_cfc (_egad ,_gcfce ,_fecf );};return _egad ,_gcfce ;};func (_fbfac *tree )fill (_fdgc ,_egea ,_baeg int )error {_bfcb :=_fbfac ._cdb ;
for _cbeed :=0;_cbeed < _fdgc ;_cbeed ++{_adbd :=_fdgc -1-_cbeed ;_bcgf :=((_egea >>uint (_adbd ))&1)!=0;_ccf :=_bfcb .walk (_bcgf );if _ccf !=nil {if _ccf ._dgb {return _a .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");
};_bfcb =_ccf ;continue ;};_ccf =&treeNode {};if _cbeed ==_fdgc -1{_ccf ._cgbg =_baeg ;_ccf ._dgb =true ;};if _egea ==0{_ccf ._effe =true ;};_bfcb .set (_bcgf ,_ccf );_bfcb =_ccf ;};return nil ;};func _gef (_dfag ,_adab []byte ,_gda ,_aafd ,_beggc int )([]byte ,int ,int ){_cbf :=_beeb (_dfag ,_beggc );
_adg :=_aafd >=0&&_dfag [_aafd ]==_gdd ||_aafd ==-1;_adab ,_gda =_cfc (_adab ,_gda ,_gd );var _bg int ;if _aafd > -1{_bg =_beggc -_aafd ;}else {_bg =_beggc -_aafd -1;};_adab ,_gda =_adfd (_adab ,_gda ,_bg ,_adg );_adg =!_adg ;_dgf :=_cbf -_beggc ;_adab ,_gda =_adfd (_adab ,_gda ,_dgf ,_adg );
_aafd =_cbf ;return _adab ,_gda ,_aafd ;};func (_cgead *Encoder )encodeG31D (_dgda [][]byte )[]byte {var _gff []byte ;_eacb :=0;for _eb :=range _dgda {if _cgead .Rows > 0&&!_cgead .EndOfBlock &&_eb ==_cgead .Rows {break ;};_agff ,_dda :=_gde (_dgda [_eb ],_eacb ,_fbf );
_gff =_cgead .appendEncodedRow (_gff ,_agff ,_eacb );if _cgead .EncodedByteAlign {_dda =0;};_eacb =_dda ;};if _cgead .EndOfBlock {_ddef ,_ :=_ecag (_eacb );_gff =_cgead .appendEncodedRow (_gff ,_ddef ,_eacb );};return _gff ;};func (_cfd *Encoder )encodeG32D (_eeg [][]byte )[]byte {var _dgg []byte ;
var _ggd int ;for _ccd :=0;_ccd < len (_eeg );_ccd +=_cfd .K {if _cfd .Rows > 0&&!_cfd .EndOfBlock &&_ccd ==_cfd .Rows {break ;};_egg ,_fde :=_gde (_eeg [_ccd ],_ggd ,_fd );_dgg =_cfd .appendEncodedRow (_dgg ,_egg ,_ggd );if _cfd .EncodedByteAlign {_fde =0;
};_ggd =_fde ;for _effd :=_ccd +1;_effd < (_ccd +_cfd .K )&&_effd < len (_eeg );_effd ++{if _cfd .Rows > 0&&!_cfd .EndOfBlock &&_effd ==_cfd .Rows {break ;};_aag ,_cbee :=_cfc (nil ,_ggd ,_bb );var _gag ,_adb ,_fag int ;_aeg :=-1;for _aeg < len (_eeg [_effd ]){_gag =_beeb (_eeg [_effd ],_aeg );
_adb =_cbbb (_eeg [_effd ],_eeg [_effd -1],_aeg );_fag =_beeb (_eeg [_effd -1],_adb );if _fag < _gag {_aag ,_cbee =_gagg (_aag ,_cbee );_aeg =_fag ;}else {if _d .Abs (float64 (_adb -_gag ))> 3{_aag ,_cbee ,_aeg =_gef (_eeg [_effd ],_aag ,_cbee ,_aeg ,_gag );
}else {_aag ,_cbee =_fbd (_aag ,_cbee ,_gag ,_adb );_aeg =_gag ;};};};_dgg =_cfd .appendEncodedRow (_dgg ,_aag ,_ggd );if _cfd .EncodedByteAlign {_cbee =0;};_ggd =_cbee %8;};};if _cfd .EndOfBlock {_fdee ,_ :=_abe (_ggd );_dgg =_cfd .appendEncodedRow (_dgg ,_fdee ,_ggd );
};return _dgg ;};func (_fbc *tree )fillWithNode (_bcge ,_adbde int ,_deba *treeNode )error {_ggef :=_fbc ._cdb ;for _ecfd :=0;_ecfd < _bcge ;_ecfd ++{_dga :=uint (_bcge -1-_ecfd );_gffe :=((_adbde >>_dga )&1)!=0;_bedd :=_ggef .walk (_gffe );if _bedd !=nil {if _bedd ._dgb {return _a .New ("\u006e\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u006c\u0065\u0061\u0066\u002c\u0020\u006eo\u0020o\u0074\u0068\u0065\u0072\u0020\u0066\u006f\u006c\u006c\u006f\u0077\u0069\u006e\u0067");
};_ggef =_bedd ;continue ;};if _ecfd ==_bcge -1{_bedd =_deba ;}else {_bedd =&treeNode {};};if _adbde ==0{_bedd ._effe =true ;};_ggef .set (_gffe ,_bedd );_ggef =_bedd ;};return nil ;};func (_geg *Decoder )Read (in []byte )(int ,error ){if _geg ._bac !=nil {return 0,_geg ._bac ;
};_gb :=len (in );var (_gge int ;_gec int ;);for _gb !=0{if _geg ._bae >=_geg ._fgc {if _gaf :=_geg .fetch ();_gaf !=nil {_geg ._bac =_gaf ;return 0,_gaf ;};};if _geg ._fgc ==-1{return _gge ,_da .EOF ;};switch {case _gb <=_geg ._fgc -_geg ._bae :_dcb :=_geg ._bfc [_geg ._bae :_geg ._bae +_gb ];
for _ ,_efad :=range _dcb {if !_geg ._ega {_efad =^_efad ;};in [_gec ]=_efad ;_gec ++;};_gge +=len (_dcb );_geg ._bae +=len (_dcb );return _gge ,nil ;default:_eacd :=_geg ._bfc [_geg ._bae :];for _ ,_feb :=range _eacd {if !_geg ._ega {_feb =^_feb ;};in [_gec ]=_feb ;
_gec ++;};_gge +=len (_eacd );_geg ._bae +=len (_eacd );_gb -=len (_eacd );};};return _gge ,nil ;};func _beeb (_gcg []byte ,_gfe int )int {if _gfe >=len (_gcg ){return _gfe ;};if _gfe < -1{_gfe =-1;};var _cgbe byte ;if _gfe > -1{_cgbe =_gcg [_gfe ];}else {_cgbe =_gdd ;
};_dad :=_gfe +1;for _dad < len (_gcg ){if _gcg [_dad ]!=_cgbe {break ;};_dad ++;};return _dad ;};func (_cac *Decoder )decodeG32D ()error {_cac ._ad =_cac ._gab ;_cac ._ab ,_cac ._bee =_cac ._bee ,_cac ._ab ;_bfb :=true ;var (_cdd bool ;_aff int ;_gcfc error ;
);_cac ._gab =0;_aae :for _aff < _cac ._fdd {_gbe :=_eac ._cdb ;for {_cdd ,_gcfc =_cac ._bcg .ReadBool ();if _gcfc !=nil {return _gcfc ;};_gbe =_gbe .walk (_cdd );if _gbe ==nil {continue _aae ;};if !_gbe ._dgb {continue ;};switch _gbe ._cgbg {case _ec :var _ggf int ;
if _bfb {_ggf ,_gcfc =_cac .decodeRun (_ca );}else {_ggf ,_gcfc =_cac .decodeRun (_ea );};if _gcfc !=nil {return _gcfc ;};_aff +=_ggf ;_cac ._ab [_cac ._gab ]=_aff ;_cac ._gab ++;if _bfb {_ggf ,_gcfc =_cac .decodeRun (_ea );}else {_ggf ,_gcfc =_cac .decodeRun (_ca );
};if _gcfc !=nil {return _gcfc ;};_aff +=_ggf ;_cac ._ab [_cac ._gab ]=_aff ;_cac ._gab ++;case _fg :_aec :=_cac .getNextChangingElement (_aff ,_bfb )+1;if _aec >=_cac ._ad {_aff =_cac ._fdd ;}else {_aff =_cac ._bee [_aec ];};default:_cb :=_cac .getNextChangingElement (_aff ,_bfb );
if _cb >=_cac ._ad ||_cb ==-1{_aff =_cac ._fdd +_gbe ._cgbg ;}else {_aff =_cac ._bee [_cb ]+_gbe ._cgbg ;};_cac ._ab [_cac ._gab ]=_aff ;_cac ._gab ++;_bfb =!_bfb ;};continue _aae ;};};return nil ;};func (_afa *Decoder )getNextChangingElement (_dbbf int ,_aba bool )int {_fbe :=0;
if !_aba {_fbe =1;};_affc :=int (uint32 (_afa ._cff )&0xFFFFFFFE)+_fbe ;if _affc > 2{_affc -=2;};if _dbbf ==0{return _affc ;};for _eaf :=_affc ;_eaf < _afa ._ad ;_eaf +=2{if _dbbf < _afa ._bee [_eaf ]{_afa ._cff =_eaf ;return _eaf ;};};return -1;};func _cfc (_bcgc []byte ,_febc int ,_abd code )([]byte ,int ){_ced :=0;
for _ced < _abd .BitsWritten {_ddec :=_febc /8;_gafc :=_febc %8;if _ddec >=len (_bcgc ){_bcgc =append (_bcgc ,0);};_dgeg :=8-_gafc ;_gdde :=_abd .BitsWritten -_ced ;if _dgeg > _gdde {_dgeg =_gdde ;};if _ced < 8{_bcgc [_ddec ]=_bcgc [_ddec ]|byte (_abd .Code >>uint (8+_gafc -_ced ))&_eec [8-_dgeg -_gafc ];
}else {_bcgc [_ddec ]=_bcgc [_ddec ]|(byte (_abd .Code <<uint (_ced -8))&_eec [8-_dgeg ])>>uint (_gafc );};_febc +=_dgeg ;_ced +=_dgeg ;};return _bcgc ,_febc ;};func (_df *Decoder )decoderRowType41D ()error {if _df ._ecg {_df ._bcg .Align ();};_df ._bcg .Mark ();
var (_bbfe bool ;_eeb error ;);if _df ._dde {_bbfe ,_eeb =_df .tryFetchEOL ();if _eeb !=nil {return _eeb ;};if !_bbfe {return _fdg ;};}else {_bbfe ,_eeb =_df .looseFetchEOL ();if _eeb !=nil {return _eeb ;};};if !_bbfe {_df ._bcg .Reset ();};if _bbfe &&_df ._egd {_df ._bcg .Mark ();
for _eff :=0;_eff < 5;_eff ++{_bbfe ,_eeb =_df .tryFetchEOL ();if _eeb !=nil {if _a .Is (_eeb ,_da .EOF ){if _eff ==0{break ;};return _dac ;};};if _bbfe {continue ;};if _eff > 0{return _dac ;};break ;};if _bbfe {return _da .EOF ;};_df ._bcg .Reset ();};
if _eeb =_df .decode1D ();_eeb !=nil {return _eeb ;};return nil ;};func _fdgd (_dfe []byte ,_gae bool ,_cabf int )(int ,int ){_abaf :=0;for _cabf < len (_dfe ){if _gae {if _dfe [_cabf ]!=_gdd {break ;};}else {if _dfe [_cabf ]!=_gea {break ;};};_abaf ++;
_cabf ++;};return _abaf ,_cabf ;};func (_agf tiffType )String ()string {switch _agf {case _cd :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u004d\u006f\u0064i\u0066\u0069\u0065\u0064\u0048\u0075\u0066\u0066\u006d\u0061n\u0052\u006c\u0065";case _deb :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0034";
case _cabb :return "\u0074\u0069\u0066\u0066\u0054\u0079\u0070\u0065\u0054\u0036";default:return "\u0075n\u0064\u0065\u0066\u0069\u006e\u0065d";};};type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;
BlackIs1 bool ;DamagedRowsBeforeError int ;};var _ba =[...][]uint16 {{3,2},{1,4},{6,5},{7},{9,8},{10,11,12},{13,14},{15},{16,17,0,18,64},{24,25,23,22,19,20,21,1792,1856,1920},{1984,2048,2112,2176,2240,2304,2368,2432,2496,2560,52,55,56,59,60,320,384,448,53,54,50,51,44,45,46,47,57,58,61,256,48,49,62,63,30,31,32,33,40,41,128,192,26,27,28,29,34,35,36,37,38,39,42,43},{640,704,768,832,1280,1344,1408,1472,1536,1600,1664,1728,512,576,896,960,1024,1088,1152,1216}};
func _cbbb (_dbe ,_effdf []byte ,_gbdf int )int {_ccbb :=_beeb (_effdf ,_gbdf );if _ccbb < len (_effdf )&&(_gbdf ==-1&&_effdf [_ccbb ]==_gdd ||_gbdf >=0&&_gbdf < len (_dbe )&&_dbe [_gbdf ]==_effdf [_ccbb ]||_gbdf >=len (_dbe )&&_dbe [_gbdf -1]!=_effdf [_ccbb ]){_ccbb =_beeb (_effdf ,_ccbb );
};return _ccbb ;};func NewDecoder (data []byte ,options DecodeOptions )(*Decoder ,error ){_dge :=&Decoder {_bcg :_g .NewReader (data ),_fdd :options .Columns ,_ff :options .Rows ,_cgea :options .DamagedRowsBeforeError ,_bfc :make ([]byte ,(options .Columns +7)/8),_bee :make ([]int ,options .Columns +2),_ab :make ([]int ,options .Columns +2),_ecg :options .EncodedByteAligned ,_ega :options .BlackIsOne ,_dde :options .EndOfLine ,_egd :options .EndOfBlock };
switch {case options .K ==0:_dge ._cgd =_deb ;if len (data )< 20{return nil ,_a .New ("\u0074o\u006f\u0020\u0073\u0068o\u0072\u0074\u0020\u0063\u0063i\u0074t\u0066a\u0078\u0020\u0073\u0074\u0072\u0065\u0061m");};_gcf :=data [:20];if _gcf [0]!=0||(_gcf [1]>>4!=1&&_gcf [1]!=1){_dge ._cgd =_cd ;
_db :=(uint16 (_gcf [0])<<8+uint16 (_gcf [1]&0xff))>>4;for _dbd :=12;_dbd < 160;_dbd ++{_db =(_db <<1)+uint16 ((_gcf [_dbd /8]>>uint16 (7-(_dbd %8)))&0x01);if _db &0xfff==1{_dge ._cgd =_deb ;break ;};};};case options .K < 0:_dge ._cgd =_cabb ;case options .K > 0:_dge ._cgd =_deb ;
_dge ._fda =true ;};switch _dge ._cgd {case _cd ,_deb ,_cabb :default:return nil ,_a .New ("\u0075\u006ek\u006e\u006f\u0077\u006e\u0020\u0063\u0063\u0069\u0074\u0074\u0066\u0061\u0078\u002e\u0044\u0065\u0063\u006f\u0064\u0065\u0072\u0020ty\u0070\u0065");
};return _dge ,nil ;};type treeNode struct{_fcf *treeNode ;_cae *treeNode ;_cgbg int ;_effe bool ;_dgb bool ;};func (_dag *treeNode )walk (_acd bool )*treeNode {if _acd {return _dag ._cae ;};return _dag ._fcf ;};type code struct{Code uint16 ;BitsWritten int ;
};func _ecag (_bda int )([]byte ,int ){var _aaf []byte ;for _bbc :=0;_bbc < 6;_bbc ++{_aaf ,_bda =_cfc (_aaf ,_bda ,_fbf );};return _aaf ,_bda %8;};func (_fec *Encoder )Encode (pixels [][]byte )[]byte {if _fec .BlackIs1 {_gdd =0;_gea =1;}else {_gdd =1;
_gea =0;};if _fec .K ==0{return _fec .encodeG31D (pixels );};if _fec .K > 0{return _fec .encodeG32D (pixels );};if _fec .K < 0{return _fec .encodeG4 (pixels );};return nil ;};func _fbd (_dee []byte ,_aaa ,_cgg ,_eggg int )([]byte ,int ){_eece :=_gbdg (_cgg ,_eggg );
_dee ,_aaa =_cfc (_dee ,_aaa ,_eece );return _dee ,_aaa ;};func (_ecgc *Encoder )appendEncodedRow (_abc ,_cdc []byte ,_bbgc int )[]byte {if len (_abc )> 0&&_bbgc !=0&&!_ecgc .EncodedByteAlign {_abc [len (_abc )-1]=_abc [len (_abc )-1]|_cdc [0];_abc =append (_abc ,_cdc [1:]...);
}else {_abc =append (_abc ,_cdc ...);};return _abc ;};var (_e *treeNode ;_cg *treeNode ;_ea *tree ;_ca *tree ;_f *tree ;_eac *tree ;_b =-2000;_ge =-1000;_fg =-3000;_ec =-4000;);func (_adc *Decoder )fetch ()error {if _adc ._fgc ==-1{return nil ;};if _adc ._bae < _adc ._fgc {return nil ;
};_adc ._fgc =0;_fcc :=_adc .decodeRow ();if _fcc !=nil {if !_a .Is (_fcc ,_da .EOF ){return _fcc ;};if _adc ._fgc !=0{return _fcc ;};_adc ._fgc =-1;};_adc ._bae =0;return nil ;};var _eg =[...][]uint16 {{0x2,0x3},{0x2,0x3},{0x2,0x3},{0x3},{0x4,0x5},{0x4,0x5,0x7},{0x4,0x7},{0x18},{0x17,0x18,0x37,0x8,0xf},{0x17,0x18,0x28,0x37,0x67,0x68,0x6c,0x8,0xc,0xd},{0x12,0x13,0x14,0x15,0x16,0x17,0x1c,0x1d,0x1e,0x1f,0x24,0x27,0x28,0x2b,0x2c,0x33,0x34,0x35,0x37,0x38,0x52,0x53,0x54,0x55,0x56,0x57,0x58,0x59,0x5a,0x5b,0x64,0x65,0x66,0x67,0x68,0x69,0x6a,0x6b,0x6c,0x6d,0xc8,0xc9,0xca,0xcb,0xcc,0xcd,0xd2,0xd3,0xd4,0xd5,0xd6,0xd7,0xda,0xdb},{0x4a,0x4b,0x4c,0x4d,0x52,0x53,0x54,0x55,0x5a,0x5b,0x64,0x65,0x6c,0x6d,0x72,0x73,0x74,0x75,0x76,0x77}};
func (_bca *treeNode )set (_adca bool ,_abg *treeNode ){if !_adca {_bca ._fcf =_abg ;}else {_bca ._cae =_abg ;};};func _gde (_cee []byte ,_dbf int ,_aca code )([]byte ,int ){_bbg :=true ;var _ebd []byte ;_ebd ,_dbf =_cfc (nil ,_dbf ,_aca );_ada :=0;var _effc int ;
for _ada < len (_cee ){_effc ,_ada =_fdgd (_cee ,_bbg ,_ada );_ebd ,_dbf =_adfd (_ebd ,_dbf ,_effc ,_bbg );_bbg =!_bbg ;};return _ebd ,_dbf %8;};func _abe (_gdce int )([]byte ,int ){var _edc []byte ;for _baa :=0;_baa < 6;_baa ++{_edc ,_gdce =_cfc (_edc ,_gdce ,_fd );
};return _edc ,_gdce %8;};type Decoder struct{_fdd int ;_ff int ;_ddf int ;_bfc []byte ;_cgea int ;_fda bool ;_ccb bool ;_eca bool ;_ega bool ;_dde bool ;_egd bool ;_ecg bool ;_fgc int ;_bae int ;_bee []int ;_ab []int ;_ad int ;_gab int ;_efa int ;_cff int ;
_bcg *_g .Reader ;_cgd tiffType ;_bac error ;};