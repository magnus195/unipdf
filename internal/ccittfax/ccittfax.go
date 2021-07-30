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

package ccittfax ;import (_g "errors";_e "github.com/unidoc/unipdf/v3/common";_b "math";);func _ddde (_agb ,_cfc []byte ,_gfec int ,_eccf bool )int {_aaag :=_dfc (_cfc ,_gfec );if _aaag < len (_cfc )&&(_gfec ==-1&&_cfc [_aaag ]==_aga ||_gfec >=0&&_gfec < len (_agb )&&_agb [_gfec ]==_cfc [_aaag ]||_gfec >=len (_agb )&&_eccf &&_cfc [_aaag ]==_aga ||_gfec >=len (_agb )&&!_eccf &&_cfc [_aaag ]==_cdd ){_aaag =_dfc (_cfc ,_aaag );
};return _aaag ;};func _bfab (_fcd *decodingTreeNode ,_cbg uint16 ,_gdg int )(*int ,*code ){if _fcd ==nil {return nil ,nil ;};if _gdg ==16{return _fcd .RunLen ,_fcd .Code ;};_aaeb :=_abce (_cbg ,_gdg );_gdg ++;var _affd *int ;var _aag *code ;if _aaeb ==1{_affd ,_aag =_bfab (_fcd .Right ,_cbg ,_gdg );
}else {_affd ,_aag =_bfab (_fcd .Left ,_cbg ,_gdg );};if _affd ==nil {_affd =_fcd .RunLen ;_aag =_fcd .Code ;};return _affd ,_aag ;};func init (){_ga =make (map[int ]code );_ga [0]=code {Code :13<<8|3<<6,BitsWritten :10};_ga [1]=code {Code :2<<(5+8),BitsWritten :3};
_ga [2]=code {Code :3<<(6+8),BitsWritten :2};_ga [3]=code {Code :2<<(6+8),BitsWritten :2};_ga [4]=code {Code :3<<(5+8),BitsWritten :3};_ga [5]=code {Code :3<<(4+8),BitsWritten :4};_ga [6]=code {Code :2<<(4+8),BitsWritten :4};_ga [7]=code {Code :3<<(3+8),BitsWritten :5};
_ga [8]=code {Code :5<<(2+8),BitsWritten :6};_ga [9]=code {Code :4<<(2+8),BitsWritten :6};_ga [10]=code {Code :4<<(1+8),BitsWritten :7};_ga [11]=code {Code :5<<(1+8),BitsWritten :7};_ga [12]=code {Code :7<<(1+8),BitsWritten :7};_ga [13]=code {Code :4<<8,BitsWritten :8};
_ga [14]=code {Code :7<<8,BitsWritten :8};_ga [15]=code {Code :12<<8,BitsWritten :9};_ga [16]=code {Code :5<<8|3<<6,BitsWritten :10};_ga [17]=code {Code :6<<8,BitsWritten :10};_ga [18]=code {Code :2<<8,BitsWritten :10};_ga [19]=code {Code :12<<8|7<<5,BitsWritten :11};
_ga [20]=code {Code :13<<8,BitsWritten :11};_ga [21]=code {Code :13<<8|4<<5,BitsWritten :11};_ga [22]=code {Code :6<<8|7<<5,BitsWritten :11};_ga [23]=code {Code :5<<8,BitsWritten :11};_ga [24]=code {Code :2<<8|7<<5,BitsWritten :11};_ga [25]=code {Code :3<<8,BitsWritten :11};
_ga [26]=code {Code :12<<8|10<<4,BitsWritten :12};_ga [27]=code {Code :12<<8|11<<4,BitsWritten :12};_ga [28]=code {Code :12<<8|12<<4,BitsWritten :12};_ga [29]=code {Code :12<<8|13<<4,BitsWritten :12};_ga [30]=code {Code :6<<8|8<<4,BitsWritten :12};_ga [31]=code {Code :6<<8|9<<4,BitsWritten :12};
_ga [32]=code {Code :6<<8|10<<4,BitsWritten :12};_ga [33]=code {Code :6<<8|11<<4,BitsWritten :12};_ga [34]=code {Code :13<<8|2<<4,BitsWritten :12};_ga [35]=code {Code :13<<8|3<<4,BitsWritten :12};_ga [36]=code {Code :13<<8|4<<4,BitsWritten :12};_ga [37]=code {Code :13<<8|5<<4,BitsWritten :12};
_ga [38]=code {Code :13<<8|6<<4,BitsWritten :12};_ga [39]=code {Code :13<<8|7<<4,BitsWritten :12};_ga [40]=code {Code :6<<8|12<<4,BitsWritten :12};_ga [41]=code {Code :6<<8|13<<4,BitsWritten :12};_ga [42]=code {Code :13<<8|10<<4,BitsWritten :12};_ga [43]=code {Code :13<<8|11<<4,BitsWritten :12};
_ga [44]=code {Code :5<<8|4<<4,BitsWritten :12};_ga [45]=code {Code :5<<8|5<<4,BitsWritten :12};_ga [46]=code {Code :5<<8|6<<4,BitsWritten :12};_ga [47]=code {Code :5<<8|7<<4,BitsWritten :12};_ga [48]=code {Code :6<<8|4<<4,BitsWritten :12};_ga [49]=code {Code :6<<8|5<<4,BitsWritten :12};
_ga [50]=code {Code :5<<8|2<<4,BitsWritten :12};_ga [51]=code {Code :5<<8|3<<4,BitsWritten :12};_ga [52]=code {Code :2<<8|4<<4,BitsWritten :12};_ga [53]=code {Code :3<<8|7<<4,BitsWritten :12};_ga [54]=code {Code :3<<8|8<<4,BitsWritten :12};_ga [55]=code {Code :2<<8|7<<4,BitsWritten :12};
_ga [56]=code {Code :2<<8|8<<4,BitsWritten :12};_ga [57]=code {Code :5<<8|8<<4,BitsWritten :12};_ga [58]=code {Code :5<<8|9<<4,BitsWritten :12};_ga [59]=code {Code :2<<8|11<<4,BitsWritten :12};_ga [60]=code {Code :2<<8|12<<4,BitsWritten :12};_ga [61]=code {Code :5<<8|10<<4,BitsWritten :12};
_ga [62]=code {Code :6<<8|6<<4,BitsWritten :12};_ga [63]=code {Code :6<<8|7<<4,BitsWritten :12};_f =make (map[int ]code );_f [0]=code {Code :53<<8,BitsWritten :8};_f [1]=code {Code :7<<(2+8),BitsWritten :6};_f [2]=code {Code :7<<(4+8),BitsWritten :4};_f [3]=code {Code :8<<(4+8),BitsWritten :4};
_f [4]=code {Code :11<<(4+8),BitsWritten :4};_f [5]=code {Code :12<<(4+8),BitsWritten :4};_f [6]=code {Code :14<<(4+8),BitsWritten :4};_f [7]=code {Code :15<<(4+8),BitsWritten :4};_f [8]=code {Code :19<<(3+8),BitsWritten :5};_f [9]=code {Code :20<<(3+8),BitsWritten :5};
_f [10]=code {Code :7<<(3+8),BitsWritten :5};_f [11]=code {Code :8<<(3+8),BitsWritten :5};_f [12]=code {Code :8<<(2+8),BitsWritten :6};_f [13]=code {Code :3<<(2+8),BitsWritten :6};_f [14]=code {Code :52<<(2+8),BitsWritten :6};_f [15]=code {Code :53<<(2+8),BitsWritten :6};
_f [16]=code {Code :42<<(2+8),BitsWritten :6};_f [17]=code {Code :43<<(2+8),BitsWritten :6};_f [18]=code {Code :39<<(1+8),BitsWritten :7};_f [19]=code {Code :12<<(1+8),BitsWritten :7};_f [20]=code {Code :8<<(1+8),BitsWritten :7};_f [21]=code {Code :23<<(1+8),BitsWritten :7};
_f [22]=code {Code :3<<(1+8),BitsWritten :7};_f [23]=code {Code :4<<(1+8),BitsWritten :7};_f [24]=code {Code :40<<(1+8),BitsWritten :7};_f [25]=code {Code :43<<(1+8),BitsWritten :7};_f [26]=code {Code :19<<(1+8),BitsWritten :7};_f [27]=code {Code :36<<(1+8),BitsWritten :7};
_f [28]=code {Code :24<<(1+8),BitsWritten :7};_f [29]=code {Code :2<<8,BitsWritten :8};_f [30]=code {Code :3<<8,BitsWritten :8};_f [31]=code {Code :26<<8,BitsWritten :8};_f [32]=code {Code :27<<8,BitsWritten :8};_f [33]=code {Code :18<<8,BitsWritten :8};
_f [34]=code {Code :19<<8,BitsWritten :8};_f [35]=code {Code :20<<8,BitsWritten :8};_f [36]=code {Code :21<<8,BitsWritten :8};_f [37]=code {Code :22<<8,BitsWritten :8};_f [38]=code {Code :23<<8,BitsWritten :8};_f [39]=code {Code :40<<8,BitsWritten :8};
_f [40]=code {Code :41<<8,BitsWritten :8};_f [41]=code {Code :42<<8,BitsWritten :8};_f [42]=code {Code :43<<8,BitsWritten :8};_f [43]=code {Code :44<<8,BitsWritten :8};_f [44]=code {Code :45<<8,BitsWritten :8};_f [45]=code {Code :4<<8,BitsWritten :8};_f [46]=code {Code :5<<8,BitsWritten :8};
_f [47]=code {Code :10<<8,BitsWritten :8};_f [48]=code {Code :11<<8,BitsWritten :8};_f [49]=code {Code :82<<8,BitsWritten :8};_f [50]=code {Code :83<<8,BitsWritten :8};_f [51]=code {Code :84<<8,BitsWritten :8};_f [52]=code {Code :85<<8,BitsWritten :8};
_f [53]=code {Code :36<<8,BitsWritten :8};_f [54]=code {Code :37<<8,BitsWritten :8};_f [55]=code {Code :88<<8,BitsWritten :8};_f [56]=code {Code :89<<8,BitsWritten :8};_f [57]=code {Code :90<<8,BitsWritten :8};_f [58]=code {Code :91<<8,BitsWritten :8};
_f [59]=code {Code :74<<8,BitsWritten :8};_f [60]=code {Code :75<<8,BitsWritten :8};_f [61]=code {Code :50<<8,BitsWritten :8};_f [62]=code {Code :51<<8,BitsWritten :8};_f [63]=code {Code :52<<8,BitsWritten :8};_aa =make (map[int ]code );_aa [64]=code {Code :3<<8|3<<6,BitsWritten :10};
_aa [128]=code {Code :12<<8|8<<4,BitsWritten :12};_aa [192]=code {Code :12<<8|9<<4,BitsWritten :12};_aa [256]=code {Code :5<<8|11<<4,BitsWritten :12};_aa [320]=code {Code :3<<8|3<<4,BitsWritten :12};_aa [384]=code {Code :3<<8|4<<4,BitsWritten :12};_aa [448]=code {Code :3<<8|5<<4,BitsWritten :12};
_aa [512]=code {Code :3<<8|12<<3,BitsWritten :13};_aa [576]=code {Code :3<<8|13<<3,BitsWritten :13};_aa [640]=code {Code :2<<8|10<<3,BitsWritten :13};_aa [704]=code {Code :2<<8|11<<3,BitsWritten :13};_aa [768]=code {Code :2<<8|12<<3,BitsWritten :13};_aa [832]=code {Code :2<<8|13<<3,BitsWritten :13};
_aa [896]=code {Code :3<<8|18<<3,BitsWritten :13};_aa [960]=code {Code :3<<8|19<<3,BitsWritten :13};_aa [1024]=code {Code :3<<8|20<<3,BitsWritten :13};_aa [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_aa [1152]=code {Code :3<<8|22<<3,BitsWritten :13};
_aa [1216]=code {Code :119<<3,BitsWritten :13};_aa [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_aa [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_aa [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_aa [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_aa [1536]=code {Code :2<<8|26<<3,BitsWritten :13};
_aa [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_aa [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_aa [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_c =make (map[int ]code );_c [64]=code {Code :27<<(3+8),BitsWritten :5};_c [128]=code {Code :18<<(3+8),BitsWritten :5};
_c [192]=code {Code :23<<(2+8),BitsWritten :6};_c [256]=code {Code :55<<(1+8),BitsWritten :7};_c [320]=code {Code :54<<8,BitsWritten :8};_c [384]=code {Code :55<<8,BitsWritten :8};_c [448]=code {Code :100<<8,BitsWritten :8};_c [512]=code {Code :101<<8,BitsWritten :8};
_c [576]=code {Code :104<<8,BitsWritten :8};_c [640]=code {Code :103<<8,BitsWritten :8};_c [704]=code {Code :102<<8,BitsWritten :9};_c [768]=code {Code :102<<8|1<<7,BitsWritten :9};_c [832]=code {Code :105<<8,BitsWritten :9};_c [896]=code {Code :105<<8|1<<7,BitsWritten :9};
_c [960]=code {Code :106<<8,BitsWritten :9};_c [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_c [1088]=code {Code :107<<8,BitsWritten :9};_c [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_c [1216]=code {Code :108<<8,BitsWritten :9};_c [1280]=code {Code :108<<8|1<<7,BitsWritten :9};
_c [1344]=code {Code :109<<8,BitsWritten :9};_c [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_c [1472]=code {Code :76<<8,BitsWritten :9};_c [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_c [1600]=code {Code :77<<8,BitsWritten :9};_c [1664]=code {Code :24<<(2+8),BitsWritten :6};
_c [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_d =make (map[int ]code );_d [1792]=code {Code :1<<8,BitsWritten :11};_d [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_d [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_d [1984]=code {Code :1<<8|2<<4,BitsWritten :12};
_d [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_d [2112]=code {Code :1<<8|4<<4,BitsWritten :12};_d [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_d [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_d [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_d [2368]=code {Code :1<<8|12<<4,BitsWritten :12};
_d [2432]=code {Code :1<<8|13<<4,BitsWritten :12};_d [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_d [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_ce =make (map[int ]byte );_ce [0]=0xFF;_ce [1]=0xFE;_ce [2]=0xFC;_ce [3]=0xF8;_ce [4]=0xF0;_ce [5]=0xE0;
_ce [6]=0xC0;_ce [7]=0x80;_ce [8]=0x00;};func _beb (_bdd [][]byte )[][]byte {_bdca :=make ([]byte ,len (_bdd [0]));for _bfcd :=range _bdca {_bdca [_bfcd ]=_aga ;};_bdd =append (_bdd ,[]byte {});for _feb :=len (_bdd )-1;_feb > 0;_feb --{_bdd [_feb ]=_bdd [_feb -1];
};_bdd [0]=_bdca ;return _bdd ;};func _fgb (_bdae []byte ,_dbb int )([]byte ,int ){return _aac (_bdae ,_dbb ,_dg )};func init (){for _dc ,_aff :=range _f {_bee (_gdb ,_aff ,0,_dc );};for _eb ,_gfe :=range _c {_bee (_gdb ,_gfe ,0,_eb );};for _ab ,_eg :=range _ga {_bee (_fg ,_eg ,0,_ab );
};for _ae ,_db :=range _aa {_bee (_fg ,_db ,0,_ae );};for _fgc ,_ee :=range _d {_bee (_gdb ,_ee ,0,_fgc );_bee (_fg ,_ee ,0,_fgc );};_bee (_de ,_dg ,0,0);_bee (_de ,_gd ,0,0);_bee (_de ,_gf ,0,0);_bee (_de ,_bd ,0,0);_bee (_de ,_da ,0,0);_bee (_de ,_ea ,0,0);
_bee (_de ,_bg ,0,0);_bee (_de ,_ge ,0,0);_bee (_de ,_bb ,0,0);};func _efd (_ba uint16 ,_ag int )(code ,bool ){_ ,_cdg :=_bfab (_de ,_ba ,_ag );if _cdg ==nil {return code {},false ;};return *_cdg ,true ;};func _abce (_dca uint16 ,_cdc int )byte {if _cdc < 8{_dca >>=8;
};_cdc %=8;_dff :=byte (0x01<<(7-uint (_cdc )));return (byte (_dca )&_dff )>>(7-uint (_cdc ));};func (_gdbd *Encoder )decodeG31D (_dac []byte )([][]byte ,error ){var _dge [][]byte ;var _bf int ;for (_bf /8)< len (_dac ){var _deb bool ;_deb ,_bf =_fbf (_dac ,_bf );
if !_deb {if _gdbd .EndOfLine {return nil ,_gc ;};}else {for _dad :=0;_dad < 5;_dad ++{_deb ,_bf =_fbf (_dac ,_bf );if !_deb {if _dad ==0{break ;};return nil ,_fc ;};};if _deb {break ;};};var _cc []byte ;_cc ,_bf =_gdbd .decodeRow1D (_dac ,_bf );if _gdbd .EncodedByteAlign &&_bf %8!=0{_bf +=8-_bf %8;
};_dge =append (_dge ,_cc );if _gdbd .Rows > 0&&!_gdbd .EndOfBlock &&len (_dge )>=_gdbd .Rows {break ;};};return _dge ,nil ;};func _edf (_ecc int )([]byte ,int ){var _bbba []byte ;for _abca :=0;_abca < 6;_abca ++{_bbba ,_ecc =_aac (_bbba ,_ecc ,_cf );};
return _bbba ,_ecc %8;};func _ccb (_dgf uint16 ,_fae int ,_dfb bool )(int ,code ){var _ff *int ;var _cba *code ;if _dfb {_ff ,_cba =_bfab (_gdb ,_dgf ,_fae );}else {_ff ,_cba =_bfab (_fg ,_dgf ,_fae );};if _ff ==nil {return -1,code {};};return *_ff ,*_cba ;
};func (_ed *Encoder )Decode (encoded []byte )([][]byte ,error ){if _ed .BlackIs1 {_aga =0;_cdd =1;}else {_aga =1;_cdd =0;};if _ed .K ==0{return _ed .decodeG31D (encoded );};if _ed .K > 0{return _ed .decodeG32D (encoded );};if _ed .K < 4{return _ed .decodeG4 (encoded );
};return nil ,nil ;};func _ebg (_bec []byte ,_efge int )(bool ,int ){return _cbb (_bec ,_efge ,_cf )};func _bfc (_gef ,_eec []byte ,_dcb int ,_fgf bool ,_edg int )([]byte ,int ,int ,error ){_fdb :=_dcb ;var _afb error ;_eec ,_dcb ,_afb =_aad (_gef ,_eec ,_dcb ,_fgf );
if _afb !=nil {return _eec ,_fdb ,_edg ,_afb ;};_fgf =!_fgf ;_eec ,_dcb ,_afb =_aad (_gef ,_eec ,_dcb ,_fgf );if _afb !=nil {return _eec ,_fdb ,_edg ,_afb ;};_edg =len (_eec );return _eec ,_dcb ,_edg ,nil ;};func _eed (_bbg []byte ,_dbaa int )(code ,int ,bool ){var (_ca uint16 ;
_cgg int ;_gb int ;);_gb =_dbaa ;_ca ,_cgg ,_ =_dde (_bbg ,_dbaa );_cfe ,_gab :=_efd (_ca ,_cgg );if !_gab {return code {},_gb ,false ;};return _cfe ,_gb +_cfe .BitsWritten ,true ;};func _bee (_ec *decodingTreeNode ,_dfbc code ,_edd int ,_bfb int ){_gca :=_abce (_dfbc .Code ,_edd );
_edd ++;if _gca ==1{if _ec .Right ==nil {_ec .Right =&decodingTreeNode {Val :_gca };};if _edd ==_dfbc .BitsWritten {_ec .Right .RunLen =&_bfb ;_ec .Right .Code =&_dfbc ;}else {_bee (_ec .Right ,_dfbc ,_edd ,_bfb );};}else {if _ec .Left ==nil {_ec .Left =&decodingTreeNode {Val :_gca };
};if _edd ==_dfbc .BitsWritten {_ec .Left .RunLen =&_bfb ;_ec .Left .Code =&_dfbc ;}else {_bee (_ec .Left ,_dfbc ,_edd ,_bfb );};};};func _add (_ebd []byte ,_faf int )(bool ,int ,error ){_ggc :=_faf ;var _fabc bool ;_fabc ,_faf =_fbf (_ebd ,_faf );if _fabc {_fabc ,_faf =_fbf (_ebd ,_faf );
if _fabc {return true ,_faf ,nil ;};return false ,_ggc ,_af ;};return false ,_ggc ,nil ;};func _dfc (_feg []byte ,_aea int )int {if _aea >=len (_feg ){return _aea ;};if _aea < -1{_aea =-1;};var _gfee byte ;if _aea > -1{_gfee =_feg [_aea ];}else {_gfee =_aga ;
};_eedf :=_aea +1;for _eedf < len (_feg ){if _feg [_eedf ]!=_gfee {break ;};_eedf ++;};return _eedf ;};func _dgg (_gaa []byte ,_adde int ,_fgcg bool )(int ,int ){var (_aaa uint16 ;_efgf int ;_ddb int ;);_ddb =_adde ;_aaa ,_efgf ,_ =_dde (_gaa ,_adde );
_cb ,_fad :=_ccb (_aaa ,_efgf ,_fgcg );if _cb ==-1{return -1,_ddb ;};return _cb ,_ddb +_fad .BitsWritten ;};var (_gdb =&decodingTreeNode {Val :255};_fg =&decodingTreeNode {Val :255};_de =&decodingTreeNode {Val :255};);func _ef (_gea [][]byte ,_ddg []byte ,_aae bool ,_dfa ,_geb int )([]byte ,int ){_dgc :=_ddde (_ddg ,_gea [len (_gea )-1],_dfa ,_aae );
_dab :=_dgc +_geb ;if _dfa ==-1{_ddg =_fbc (_ddg ,_aae ,_dab -_dfa -1);}else {_ddg =_fbc (_ddg ,_aae ,_dab -_dfa );};_dfa =_dab ;return _ddg ,_dfa ;};var (_aga byte =1;_cdd byte =0;);func _cbb (_bcg []byte ,_bge int ,_cfgc code )(bool ,int ){_feea :=_bge ;
var (_bgb uint16 ;_dabee int ;);_bgb ,_dabee ,_bge =_dde (_bcg ,_bge );if _dabee > 3{return false ,_feea ;};_bgb >>=uint (3-_dabee );_bgb <<=3;if _bgb !=_cfgc .Code {return false ,_feea ;};return true ,_bge -3+_dabee ;};func _fbf (_gcg []byte ,_fag int )(bool ,int ){_bdc :=_fag ;
var (_ggcc uint16 ;_ddd int ;);_ggcc ,_ddd ,_fag =_dde (_gcg ,_fag );if _ddd > 4{return false ,_bdc ;};_ggcc >>=uint (4-_ddd );_ggcc <<=4;if _ggcc !=_bc .Code {return false ,_bdc ;};return true ,_fag -4+_ddd ;};func _gga (_affdc []byte ,_ceec int ,_bgf int ,_fec bool )([]byte ,int ){var (_aeeb code ;
_ddbe bool ;);for !_ddbe {_aeeb ,_bgf ,_ddbe =_ace (_bgf ,_fec );_affdc ,_ceec =_aac (_affdc ,_ceec ,_aeeb );};return _affdc ,_ceec ;};func _fgad (_geaf []byte ,_facd bool ,_fdf int )(int ,int ){_bdcd :=0;for _fdf < len (_geaf ){if _facd {if _geaf [_fdf ]!=_aga {break ;
};}else {if _geaf [_fdf ]!=_cdd {break ;};};_bdcd ++;_fdf ++;};return _bdcd ,_fdf ;};func (_afc *Encoder )appendEncodedRow (_daab ,_gcf []byte ,_fdg int )[]byte {if len (_daab )> 0&&_fdg !=0&&!_afc .EncodedByteAlign {_daab [len (_daab )-1]=_daab [len (_daab )-1]|_gcf [0];
_daab =append (_daab ,_gcf [1:]...);}else {_daab =append (_daab ,_gcf ...);};return _daab ;};func (_cfed *Encoder )decodeRow1D (_gebdg []byte ,_bgg int )([]byte ,int ){var _cff []byte ;_dbf :=true ;var _dae int ;_dae ,_bgg =_dgg (_gebdg ,_bgg ,_dbf );for _dae !=-1{_cff =_fbc (_cff ,_dbf ,_dae );
if _dae < 64{if len (_cff )>=_cfed .Columns {break ;};_dbf =!_dbf ;};_dae ,_bgg =_dgg (_gebdg ,_bgg ,_dbf );};return _cff ,_bgg ;};func (_bfe *Encoder )encodeG31D (_bef [][]byte )[]byte {var _adb []byte ;_dfed :=0;for _ggg :=range _bef {if _bfe .Rows > 0&&!_bfe .EndOfBlock &&_ggg ==_bfe .Rows {break ;
};_abf ,_cbf :=_aef (_bef [_ggg ],_dfed ,_bc );_adb =_bfe .appendEncodedRow (_adb ,_abf ,_dfed );if _bfe .EncodedByteAlign {_cbf =0;};_dfed =_cbf ;};if _bfe .EndOfBlock {_bgbe ,_ :=_ebb (_dfed );_adb =_bfe .appendEncodedRow (_adb ,_bgbe ,_dfed );};return _adb ;
};type decodingTreeNode struct{Val byte ;RunLen *int ;Code *code ;Left *decodingTreeNode ;Right *decodingTreeNode ;};func _facg (_fdd []byte ,_cfb ,_fecb ,_gcb int )([]byte ,int ){_aebd :=_fge (_fecb ,_gcb );_fdd ,_cfb =_aac (_fdd ,_cfb ,_aebd );return _fdd ,_cfb ;
};var (_af =_g .New ("\u0045\u004f\u0046\u0042 c\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074e\u0064");_cee =_g .New ("R\u0054\u0043\u0020\u0063od\u0065 \u0069\u0073\u0020\u0063\u006fr\u0072\u0075\u0070\u0074\u0065\u0064");
_eac =_g .New ("\u0077\u0072o\u006e\u0067\u0020\u0063\u006f\u0064\u0065\u0020\u0069\u006e\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020mo\u0064\u0065");_gc =_g .New ("\u006e\u006f\u0020\u0045\u004f\u004c\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0074\u0068\u0065\u0020\u0045\u006e\u0064O\u0066\u004c\u0069\u006e\u0065\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0065");
_fc =_g .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u004f\u004c");_gee =_g .New ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0032\u0044\u0020\u0063\u006f\u0064\u0065"););func (_acf *Encoder )Encode (pixels [][]byte )[]byte {if _acf .BlackIs1 {_aga =0;
_cdd =1;}else {_aga =1;_cdd =0;};if _acf .K ==0{return _acf .encodeG31D (pixels );};if _acf .K > 0{return _acf .encodeG32D (pixels );};if _acf .K < 0{return _acf .encodeG4 (pixels );};return nil ;};func _cfab (_bbbc int )([]byte ,int ){var _eecb []byte ;
for _gefg :=0;_gefg < 2;_gefg ++{_eecb ,_bbbc =_aac (_eecb ,_bbbc ,_bc );};return _eecb ,_bbbc %8;};func _dde (_dgb []byte ,_gebd int )(uint16 ,int ,int ){_dfe :=_gebd ;_bfag :=_gebd /8;_gebd %=8;if _bfag >=len (_dgb ){return 0,16,_dfe ;};_baa :=byte (0xFF>>uint (_gebd ));
_gfd :=uint16 ((_dgb [_bfag ]&_baa )<<uint (_gebd ))<<8;_bab :=8-_gebd ;_bfag ++;_gebd =0;if _bfag >=len (_dgb ){return _gfd >>(16-uint (_bab )),16-_bab ,_dfe +_bab ;};_gfd |=uint16 (_dgb [_bfag ])<<(8-uint (_bab ));_bab +=8;_bfag ++;_gebd =0;if _bfag >=len (_dgb ){return _gfd >>(16-uint (_bab )),16-_bab ,_dfe +_bab ;
};if _bab ==16{return _gfd ,0,_dfe +_bab ;};_fee :=16-_bab ;_gfd |=uint16 (_dgb [_bfag ]>>(8-uint (_fee )));return _gfd ,0,_dfe +16;};func _ebb (_baf int )([]byte ,int ){var _egc []byte ;for _effc :=0;_effc < 6;_effc ++{_egc ,_baf =_aac (_egc ,_baf ,_bc );
};return _egc ,_baf %8;};var (_ga map[int ]code ;_f map[int ]code ;_aa map[int ]code ;_c map[int ]code ;_d map[int ]code ;_ce map[int ]byte ;_bc =code {Code :1<<4,BitsWritten :12};_cf =code {Code :3<<3,BitsWritten :13};_fd =code {Code :2<<3,BitsWritten :13};
_dg =code {Code :1<<12,BitsWritten :4};_gd =code {Code :1<<13,BitsWritten :3};_gf =code {Code :1<<15,BitsWritten :1};_bd =code {Code :3<<13,BitsWritten :3};_da =code {Code :3<<10,BitsWritten :6};_ea =code {Code :3<<9,BitsWritten :7};_bg =code {Code :2<<13,BitsWritten :3};
_ge =code {Code :2<<10,BitsWritten :6};_bb =code {Code :2<<9,BitsWritten :7};);type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func _aad (_gdc ,_bbf []byte ,_ccd int ,_ad bool )([]byte ,int ,error ){_bfa :=_ccd ;
var _bbe int ;for _bbe ,_ccd =_dgg (_gdc ,_ccd ,_ad );_bbe !=-1;_bbe ,_ccd =_dgg (_gdc ,_ccd ,_ad ){_bbf =_fbc (_bbf ,_ad ,_bbe );if _bbe < 64{break ;};};if _bbe ==-1{return _bbf ,_bfa ,_eac ;};return _bbf ,_ccd ,nil ;};func _fbc (_cbd []byte ,_cfg bool ,_cffg int )[]byte {if _cffg < 0{return _cbd ;
};_agd :=make ([]byte ,_cffg );if _cfg {for _bda :=0;_bda < len (_agd );_bda ++{_agd [_bda ]=_aga ;};}else {for _geag :=0;_geag < len (_agd );_geag ++{_agd [_geag ]=_cdd ;};};_cbd =append (_cbd ,_agd ...);return _cbd ;};func _ace (_gda int ,_fgag bool )(code ,int ,bool ){if _gda < 64{if _fgag {return _f [_gda ],0,true ;
};return _ga [_gda ],0,true ;};_gdgb :=_gda /64;if _gdgb > 40{return _d [2560],_gda -2560,false ;};if _gdgb > 27{return _d [_gdgb *64],_gda -_gdgb *64,false ;};if _fgag {return _c [_gdgb *64],_gda -_gdgb *64,false ;};return _aa [_gdgb *64],_gda -_gdgb *64,false ;
};func _fgd (_faa []byte ,_eacb int )(bool ,int ){return _cbb (_faa ,_eacb ,_fd )};func (_ceb *Encoder )decodeG4 (_dbc []byte )([][]byte ,error ){_ac :=make ([]byte ,_ceb .Columns );for _cd :=range _ac {_ac [_cd ]=_aga ;};_fca :=make ([][]byte ,1);_fca [0]=_ac ;
var (_aeb bool ;_df error ;_egd int ;);for (_egd /8)< len (_dbc ){_aeb ,_egd ,_df =_add (_dbc ,_egd );if _df !=nil {return nil ,_df ;};if _aeb {break ;};var (_dfg code ;_cgc bool ;);_fdc :=true ;var _ceeb []byte ;_dd :=-1;_fe :=true ;for _dd < _ceb .Columns {_dfg ,_egd ,_cgc =_eed (_dbc ,_egd );
if !_cgc {_e .Log .Warning ("E\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064\u0069\u006e\u0067\u0020\u0063\u0063i\u0074\u0074\u0066a\u0078:\u0020\u0025\u0076",_gee );_fe =false ;break ;};switch _dfg {case _dg :_ceeb ,_dd =_cge (_fca ,_ceeb ,_fdc ,_dd );
case _gd :_ceeb ,_egd ,_dd ,_df =_bfc (_dbc ,_ceeb ,_egd ,_fdc ,_dd );if _df !=nil {_e .Log .Error ("\u0045\u0072\u0072\u006f\u0072\u0020\u0064e\u0063\u006f\u0064i\u006e\u0067\u0020\u0063c\u0069\u0074\u0074\u0066\u0061\u0078\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020\u006d\u006f\u0064\u0065\u003a\u0020\u0025\u0076",_df );
_fe =false ;break ;};case _gf :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,0);_fdc =!_fdc ;case _bd :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,1);_fdc =!_fdc ;case _da :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,2);_fdc =!_fdc ;case _ea :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,3);
_fdc =!_fdc ;case _bg :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,-1);_fdc =!_fdc ;case _ge :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,-2);_fdc =!_fdc ;case _bb :_ceeb ,_dd =_ef (_fca ,_ceeb ,_fdc ,_dd ,-3);_fdc =!_fdc ;};if len (_ceeb )>=_ceb .Columns {break ;
};};if !_fe {_egd +=8-_egd %8;continue ;};if _ceb .EncodedByteAlign &&_egd %8!=0{_egd +=8-_egd %8;};_fca =append (_fca ,_ceeb );if _ceb .Rows > 0&&!_ceb .EndOfBlock &&len (_fca )>=(_ceb .Rows +1){break ;};};_fca =_fca [1:];return _fca ,nil ;};func (_bbga *Encoder )encodeG4 (_gce [][]byte )[]byte {_cdge :=make ([][]byte ,len (_gce ));
copy (_cdge ,_gce );_cdge =_beb (_cdge );var _gdbg []byte ;var _daf int ;for _bfae :=1;_bfae < len (_cdge );_bfae ++{if _bbga .Rows > 0&&!_bbga .EndOfBlock &&_bfae ==(_bbga .Rows +1){break ;};var _gcgg []byte ;var _bdg ,_agdf ,_aadf int ;_cfa :=_daf ;_cef :=-1;
for _cef < len (_cdge [_bfae ]){_bdg =_dfc (_cdge [_bfae ],_cef );_agdf =_dce (_cdge [_bfae ],_cdge [_bfae -1],_cef );_aadf =_dfc (_cdge [_bfae -1],_agdf );if _aadf < _bdg {_gcgg ,_cfa =_aac (_gcgg ,_cfa ,_dg );_cef =_aadf ;}else {if _b .Abs (float64 (_agdf -_bdg ))> 3{_gcgg ,_cfa ,_cef =_geg (_cdge [_bfae ],_gcgg ,_cfa ,_cef ,_bdg );
}else {_gcgg ,_cfa =_facg (_gcgg ,_cfa ,_bdg ,_agdf );_cef =_bdg ;};};};_gdbg =_bbga .appendEncodedRow (_gdbg ,_gcgg ,_daf );if _bbga .EncodedByteAlign {_cfa =0;};_daf =_cfa %8;};if _bbga .EndOfBlock {_ged ,_ :=_cfab (_daf );_gdbg =_bbga .appendEncodedRow (_gdbg ,_ged ,_daf );
};return _gdbg ;};func _geg (_eeb ,_dgbc []byte ,_caf ,_ecf ,_eee int )([]byte ,int ,int ){_bcb :=_dfc (_eeb ,_eee );_ddbg :=_ecf >=0&&_eeb [_ecf ]==_aga ||_ecf ==-1;_dgbc ,_caf =_aac (_dgbc ,_caf ,_gd );var _fdad int ;if _ecf > -1{_fdad =_eee -_ecf ;}else {_fdad =_eee -_ecf -1;
};_dgbc ,_caf =_gga (_dgbc ,_caf ,_fdad ,_ddbg );_ddbg =!_ddbg ;_fgbe :=_bcb -_eee ;_dgbc ,_caf =_gga (_dgbc ,_caf ,_fgbe ,_ddbg );_ecf =_bcb ;return _dgbc ,_caf ,_ecf ;};type code struct{Code uint16 ;BitsWritten int ;};func _cge (_abc [][]byte ,_efg []byte ,_fb bool ,_dbd int )([]byte ,int ){_cgf :=_ddde (_efg ,_abc [len (_abc )-1],_dbd ,_fb );
_dda :=_dfc (_abc [len (_abc )-1],_cgf );if _dbd ==-1{_efg =_fbc (_efg ,_fb ,_dda -_dbd -1);}else {_efg =_fbc (_efg ,_fb ,_dda -_dbd );};_dbd =_dda ;return _efg ,_dbd ;};func _fge (_ddf ,_aba int )code {var _bcc code ;switch _aba -_ddf {case -1:_bcc =_bd ;
case -2:_bcc =_da ;case -3:_bcc =_ea ;case 0:_bcc =_gf ;case 1:_bcc =_bg ;case 2:_bcc =_ge ;case 3:_bcc =_bb ;};return _bcc ;};func (_gdf *Encoder )decodeG32D (_fa []byte )([][]byte ,error ){var (_bbc [][]byte ;_ceg int ;_daa error ;);_daaa :for (_ceg /8)< len (_fa ){var _dee bool ;
_dee ,_ceg ,_daa =_cea (_fa ,_ceg );if _daa !=nil {return nil ,_daa ;};if _dee {break ;};_dee ,_ceg =_ebg (_fa ,_ceg );if !_dee {if _gdf .EndOfLine {return nil ,_gc ;};};var _fce []byte ;_fce ,_ceg =_gdf .decodeRow1D (_fa ,_ceg );if _gdf .EncodedByteAlign &&_ceg %8!=0{_ceg +=8-_ceg %8;
};if _fce !=nil {_bbc =append (_bbc ,_fce );};if _gdf .Rows > 0&&!_gdf .EndOfBlock &&len (_bbc )>=_gdf .Rows {break ;};for _afd :=1;_afd < _gdf .K &&(_ceg /8)< len (_fa );_afd ++{_dee ,_ceg =_fgd (_fa ,_ceg );if !_dee {_dee ,_ceg ,_daa =_cea (_fa ,_ceg );
if _daa !=nil {return nil ,_daa ;};if _dee {break _daaa ;}else {if _gdf .EndOfLine {return nil ,_gc ;};};};var (_cg code ;_gg bool ;);_gag :=true ;var _dba []byte ;_eace :=-1;for _cg ,_ceg ,_gg =_eed (_fa ,_ceg );_gg ;_cg ,_ceg ,_gg =_eed (_fa ,_ceg ){switch _cg {case _dg :_dba ,_eace =_cge (_bbc ,_dba ,_gag ,_eace );
case _gd :_dba ,_ceg ,_eace ,_daa =_bfc (_fa ,_dba ,_ceg ,_gag ,_eace );if _daa !=nil {return nil ,_daa ;};case _gf :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,0);_gag =!_gag ;case _bd :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,1);_gag =!_gag ;case _da :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,2);
_gag =!_gag ;case _ea :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,3);_gag =!_gag ;case _bg :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,-1);_gag =!_gag ;case _ge :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,-2);_gag =!_gag ;case _bb :_dba ,_eace =_ef (_bbc ,_dba ,_gag ,_eace ,-3);
_gag =!_gag ;};if len (_dba )>=_gdf .Columns {break ;};};if _gdf .EncodedByteAlign &&_ceg %8!=0{_ceg +=8-_ceg %8;};if _dba !=nil {_bbc =append (_bbc ,_dba );};if _gdf .Rows > 0&&!_gdf .EndOfBlock &&len (_bbc )>=_gdf .Rows {break _daaa ;};};};return _bbc ,nil ;
};func _aac (_fcae []byte ,_dbcc int ,_gbc code )([]byte ,int ){_fade :=0;for _fade < _gbc .BitsWritten {_cbeb :=_dbcc /8;_ede :=_dbcc %8;if _cbeb >=len (_fcae ){_fcae =append (_fcae ,0);};_geec :=8-_ede ;_bcd :=_gbc .BitsWritten -_fade ;if _geec > _bcd {_geec =_bcd ;
};if _fade < 8{_fcae [_cbeb ]=_fcae [_cbeb ]|byte (_gbc .Code >>uint (8+_ede -_fade ))&_ce [8-_geec -_ede ];}else {_fcae [_cbeb ]=_fcae [_cbeb ]|(byte (_gbc .Code <<uint (_fade -8))&_ce [8-_geec ])>>uint (_ede );};_dbcc +=_geec ;_fade +=_geec ;};return _fcae ,_dbcc ;
};func _aef (_afe []byte ,_cbe int ,_afda code )([]byte ,int ){_bbgf :=true ;var _aee []byte ;_aee ,_cbe =_aac (nil ,_cbe ,_afda );_dgbg :=0;var _cfd int ;for _dgbg < len (_afe ){_cfd ,_dgbg =_fgad (_afe ,_bbgf ,_dgbg );_aee ,_cbe =_gga (_aee ,_cbe ,_cfd ,_bbgf );
_bbgf =!_bbgf ;};return _aee ,_cbe %8;};func _dce (_cgge ,_fdbb []byte ,_dgeg int )int {_acfg :=_dfc (_fdbb ,_dgeg );if _acfg < len (_fdbb )&&(_dgeg ==-1&&_fdbb [_acfg ]==_aga ||_dgeg >=0&&_dgeg < len (_cgge )&&_cgge [_dgeg ]==_fdbb [_acfg ]||_dgeg >=len (_cgge )&&_cgge [_dgeg -1]!=_fdbb [_acfg ]){_acfg =_dfc (_fdbb ,_acfg );
};return _acfg ;};func _cea (_bbb []byte ,_gad int )(bool ,int ,error ){_cdb :=_gad ;var _fab =false ;for _fbe :=0;_fbe < 6;_fbe ++{_fab ,_gad =_ebg (_bbb ,_gad );if !_fab {if _fbe > 1{return false ,_cdb ,_cee ;};_gad =_cdb ;break ;};};return _fab ,_gad ,nil ;
};func (_eaa *Encoder )encodeG32D (_fac [][]byte )[]byte {var _fda []byte ;var _bca int ;for _adda :=0;_adda < len (_fac );_adda +=_eaa .K {if _eaa .Rows > 0&&!_eaa .EndOfBlock &&_adda ==_eaa .Rows {break ;};_bce ,_dcg :=_aef (_fac [_adda ],_bca ,_cf );
_fda =_eaa .appendEncodedRow (_fda ,_bce ,_bca );if _eaa .EncodedByteAlign {_dcg =0;};_bca =_dcg ;for _dfad :=_adda +1;_dfad < (_adda +_eaa .K )&&_dfad < len (_fac );_dfad ++{if _eaa .Rows > 0&&!_eaa .EndOfBlock &&_dfad ==_eaa .Rows {break ;};_gbd ,_cde :=_aac (nil ,_bca ,_fd );
var _cdce ,_fcc ,_ddeg int ;_ccdf :=-1;for _ccdf < len (_fac [_dfad ]){_cdce =_dfc (_fac [_dfad ],_ccdf );_fcc =_dce (_fac [_dfad ],_fac [_dfad -1],_ccdf );_ddeg =_dfc (_fac [_dfad -1],_fcc );if _ddeg < _cdce {_gbd ,_cde =_fgb (_gbd ,_cde );_ccdf =_ddeg ;
}else {if _b .Abs (float64 (_fcc -_cdce ))> 3{_gbd ,_cde ,_ccdf =_geg (_fac [_dfad ],_gbd ,_cde ,_ccdf ,_cdce );}else {_gbd ,_cde =_facg (_gbd ,_cde ,_cdce ,_fcc );_ccdf =_cdce ;};};};_fda =_eaa .appendEncodedRow (_fda ,_gbd ,_bca );if _eaa .EncodedByteAlign {_cde =0;
};_bca =_cde %8;};};if _eaa .EndOfBlock {_bcf ,_ :=_edf (_bca );_fda =_eaa .appendEncodedRow (_fda ,_bcf ,_bca );};return _fda ;};