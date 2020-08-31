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

package ccittfax ;import (_f "errors";_g "math";);func _ega (_edda *decodingTreeNode ,_bgdg code ,_ffc int ,_gag int ){_eef :=_fd (_bgdg .Code ,_ffc );_ffc ++;if _eef ==1{if _edda .Right ==nil {_edda .Right =&decodingTreeNode {Val :_eef };};if _ffc ==_bgdg .BitsWritten {_edda .Right .RunLen =&_gag ;_edda .Right .Code =&_bgdg ;}else {_ega (_edda .Right ,_bgdg ,_ffc ,_gag );};}else {if _edda .Left ==nil {_edda .Left =&decodingTreeNode {Val :_eef };};if _ffc ==_bgdg .BitsWritten {_edda .Left .RunLen =&_gag ;_edda .Left .Code =&_bgdg ;}else {_ega (_edda .Left ,_bgdg ,_ffc ,_gag );};};};func _caf (_fgef ,_cada []byte ,_edg int ,_cfd bool ,_bgca int )([]byte ,int ,int ,error ){_dc :=_edg ;var _bac error ;_cada ,_edg ,_bac =_bea (_fgef ,_cada ,_edg ,_cfd );if _bac !=nil {return _cada ,_dc ,_bgca ,_bac ;};_cfd =!_cfd ;_cada ,_edg ,_bac =_bea (_fgef ,_cada ,_edg ,_cfd );if _bac !=nil {return _cada ,_dc ,_bgca ,_bac ;};_bgca =len (_cada );return _cada ,_edg ,_bgca ,nil ;};func init (){_e =make (map[int ]code );_e [0]=code {Code :13<<8|3<<6,BitsWritten :10};_e [1]=code {Code :2<<(5+8),BitsWritten :3};_e [2]=code {Code :3<<(6+8),BitsWritten :2};_e [3]=code {Code :2<<(6+8),BitsWritten :2};_e [4]=code {Code :3<<(5+8),BitsWritten :3};_e [5]=code {Code :3<<(4+8),BitsWritten :4};_e [6]=code {Code :2<<(4+8),BitsWritten :4};_e [7]=code {Code :3<<(3+8),BitsWritten :5};_e [8]=code {Code :5<<(2+8),BitsWritten :6};_e [9]=code {Code :4<<(2+8),BitsWritten :6};_e [10]=code {Code :4<<(1+8),BitsWritten :7};_e [11]=code {Code :5<<(1+8),BitsWritten :7};_e [12]=code {Code :7<<(1+8),BitsWritten :7};_e [13]=code {Code :4<<8,BitsWritten :8};_e [14]=code {Code :7<<8,BitsWritten :8};_e [15]=code {Code :12<<8,BitsWritten :9};_e [16]=code {Code :5<<8|3<<6,BitsWritten :10};_e [17]=code {Code :6<<8,BitsWritten :10};_e [18]=code {Code :2<<8,BitsWritten :10};_e [19]=code {Code :12<<8|7<<5,BitsWritten :11};_e [20]=code {Code :13<<8,BitsWritten :11};_e [21]=code {Code :13<<8|4<<5,BitsWritten :11};_e [22]=code {Code :6<<8|7<<5,BitsWritten :11};_e [23]=code {Code :5<<8,BitsWritten :11};_e [24]=code {Code :2<<8|7<<5,BitsWritten :11};_e [25]=code {Code :3<<8,BitsWritten :11};_e [26]=code {Code :12<<8|10<<4,BitsWritten :12};_e [27]=code {Code :12<<8|11<<4,BitsWritten :12};_e [28]=code {Code :12<<8|12<<4,BitsWritten :12};_e [29]=code {Code :12<<8|13<<4,BitsWritten :12};_e [30]=code {Code :6<<8|8<<4,BitsWritten :12};_e [31]=code {Code :6<<8|9<<4,BitsWritten :12};_e [32]=code {Code :6<<8|10<<4,BitsWritten :12};_e [33]=code {Code :6<<8|11<<4,BitsWritten :12};_e [34]=code {Code :13<<8|2<<4,BitsWritten :12};_e [35]=code {Code :13<<8|3<<4,BitsWritten :12};_e [36]=code {Code :13<<8|4<<4,BitsWritten :12};_e [37]=code {Code :13<<8|5<<4,BitsWritten :12};_e [38]=code {Code :13<<8|6<<4,BitsWritten :12};_e [39]=code {Code :13<<8|7<<4,BitsWritten :12};_e [40]=code {Code :6<<8|12<<4,BitsWritten :12};_e [41]=code {Code :6<<8|13<<4,BitsWritten :12};_e [42]=code {Code :13<<8|10<<4,BitsWritten :12};_e [43]=code {Code :13<<8|11<<4,BitsWritten :12};_e [44]=code {Code :5<<8|4<<4,BitsWritten :12};_e [45]=code {Code :5<<8|5<<4,BitsWritten :12};_e [46]=code {Code :5<<8|6<<4,BitsWritten :12};_e [47]=code {Code :5<<8|7<<4,BitsWritten :12};_e [48]=code {Code :6<<8|4<<4,BitsWritten :12};_e [49]=code {Code :6<<8|5<<4,BitsWritten :12};_e [50]=code {Code :5<<8|2<<4,BitsWritten :12};_e [51]=code {Code :5<<8|3<<4,BitsWritten :12};_e [52]=code {Code :2<<8|4<<4,BitsWritten :12};_e [53]=code {Code :3<<8|7<<4,BitsWritten :12};_e [54]=code {Code :3<<8|8<<4,BitsWritten :12};_e [55]=code {Code :2<<8|7<<4,BitsWritten :12};_e [56]=code {Code :2<<8|8<<4,BitsWritten :12};_e [57]=code {Code :5<<8|8<<4,BitsWritten :12};_e [58]=code {Code :5<<8|9<<4,BitsWritten :12};_e [59]=code {Code :2<<8|11<<4,BitsWritten :12};_e [60]=code {Code :2<<8|12<<4,BitsWritten :12};_e [61]=code {Code :5<<8|10<<4,BitsWritten :12};_e [62]=code {Code :6<<8|6<<4,BitsWritten :12};_e [63]=code {Code :6<<8|7<<4,BitsWritten :12};_a =make (map[int ]code );_a [0]=code {Code :53<<8,BitsWritten :8};_a [1]=code {Code :7<<(2+8),BitsWritten :6};_a [2]=code {Code :7<<(4+8),BitsWritten :4};_a [3]=code {Code :8<<(4+8),BitsWritten :4};_a [4]=code {Code :11<<(4+8),BitsWritten :4};_a [5]=code {Code :12<<(4+8),BitsWritten :4};_a [6]=code {Code :14<<(4+8),BitsWritten :4};_a [7]=code {Code :15<<(4+8),BitsWritten :4};_a [8]=code {Code :19<<(3+8),BitsWritten :5};_a [9]=code {Code :20<<(3+8),BitsWritten :5};_a [10]=code {Code :7<<(3+8),BitsWritten :5};_a [11]=code {Code :8<<(3+8),BitsWritten :5};_a [12]=code {Code :8<<(2+8),BitsWritten :6};_a [13]=code {Code :3<<(2+8),BitsWritten :6};_a [14]=code {Code :52<<(2+8),BitsWritten :6};_a [15]=code {Code :53<<(2+8),BitsWritten :6};_a [16]=code {Code :42<<(2+8),BitsWritten :6};_a [17]=code {Code :43<<(2+8),BitsWritten :6};_a [18]=code {Code :39<<(1+8),BitsWritten :7};_a [19]=code {Code :12<<(1+8),BitsWritten :7};_a [20]=code {Code :8<<(1+8),BitsWritten :7};_a [21]=code {Code :23<<(1+8),BitsWritten :7};_a [22]=code {Code :3<<(1+8),BitsWritten :7};_a [23]=code {Code :4<<(1+8),BitsWritten :7};_a [24]=code {Code :40<<(1+8),BitsWritten :7};_a [25]=code {Code :43<<(1+8),BitsWritten :7};_a [26]=code {Code :19<<(1+8),BitsWritten :7};_a [27]=code {Code :36<<(1+8),BitsWritten :7};_a [28]=code {Code :24<<(1+8),BitsWritten :7};_a [29]=code {Code :2<<8,BitsWritten :8};_a [30]=code {Code :3<<8,BitsWritten :8};_a [31]=code {Code :26<<8,BitsWritten :8};_a [32]=code {Code :27<<8,BitsWritten :8};_a [33]=code {Code :18<<8,BitsWritten :8};_a [34]=code {Code :19<<8,BitsWritten :8};_a [35]=code {Code :20<<8,BitsWritten :8};_a [36]=code {Code :21<<8,BitsWritten :8};_a [37]=code {Code :22<<8,BitsWritten :8};_a [38]=code {Code :23<<8,BitsWritten :8};_a [39]=code {Code :40<<8,BitsWritten :8};_a [40]=code {Code :41<<8,BitsWritten :8};_a [41]=code {Code :42<<8,BitsWritten :8};_a [42]=code {Code :43<<8,BitsWritten :8};_a [43]=code {Code :44<<8,BitsWritten :8};_a [44]=code {Code :45<<8,BitsWritten :8};_a [45]=code {Code :4<<8,BitsWritten :8};_a [46]=code {Code :5<<8,BitsWritten :8};_a [47]=code {Code :10<<8,BitsWritten :8};_a [48]=code {Code :11<<8,BitsWritten :8};_a [49]=code {Code :82<<8,BitsWritten :8};_a [50]=code {Code :83<<8,BitsWritten :8};_a [51]=code {Code :84<<8,BitsWritten :8};_a [52]=code {Code :85<<8,BitsWritten :8};_a [53]=code {Code :36<<8,BitsWritten :8};_a [54]=code {Code :37<<8,BitsWritten :8};_a [55]=code {Code :88<<8,BitsWritten :8};_a [56]=code {Code :89<<8,BitsWritten :8};_a [57]=code {Code :90<<8,BitsWritten :8};_a [58]=code {Code :91<<8,BitsWritten :8};_a [59]=code {Code :74<<8,BitsWritten :8};_a [60]=code {Code :75<<8,BitsWritten :8};_a [61]=code {Code :50<<8,BitsWritten :8};_a [62]=code {Code :51<<8,BitsWritten :8};_a [63]=code {Code :52<<8,BitsWritten :8};_db =make (map[int ]code );_db [64]=code {Code :3<<8|3<<6,BitsWritten :10};_db [128]=code {Code :12<<8|8<<4,BitsWritten :12};_db [192]=code {Code :12<<8|9<<4,BitsWritten :12};_db [256]=code {Code :5<<8|11<<4,BitsWritten :12};_db [320]=code {Code :3<<8|3<<4,BitsWritten :12};_db [384]=code {Code :3<<8|4<<4,BitsWritten :12};_db [448]=code {Code :3<<8|5<<4,BitsWritten :12};_db [512]=code {Code :3<<8|12<<3,BitsWritten :13};_db [576]=code {Code :3<<8|13<<3,BitsWritten :13};_db [640]=code {Code :2<<8|10<<3,BitsWritten :13};_db [704]=code {Code :2<<8|11<<3,BitsWritten :13};_db [768]=code {Code :2<<8|12<<3,BitsWritten :13};_db [832]=code {Code :2<<8|13<<3,BitsWritten :13};_db [896]=code {Code :3<<8|18<<3,BitsWritten :13};_db [960]=code {Code :3<<8|19<<3,BitsWritten :13};_db [1024]=code {Code :3<<8|20<<3,BitsWritten :13};_db [1088]=code {Code :3<<8|21<<3,BitsWritten :13};_db [1152]=code {Code :3<<8|22<<3,BitsWritten :13};_db [1216]=code {Code :119<<3,BitsWritten :13};_db [1280]=code {Code :2<<8|18<<3,BitsWritten :13};_db [1344]=code {Code :2<<8|19<<3,BitsWritten :13};_db [1408]=code {Code :2<<8|20<<3,BitsWritten :13};_db [1472]=code {Code :2<<8|21<<3,BitsWritten :13};_db [1536]=code {Code :2<<8|26<<3,BitsWritten :13};_db [1600]=code {Code :2<<8|27<<3,BitsWritten :13};_db [1664]=code {Code :3<<8|4<<3,BitsWritten :13};_db [1728]=code {Code :3<<8|5<<3,BitsWritten :13};_b =make (map[int ]code );_b [64]=code {Code :27<<(3+8),BitsWritten :5};_b [128]=code {Code :18<<(3+8),BitsWritten :5};_b [192]=code {Code :23<<(2+8),BitsWritten :6};_b [256]=code {Code :55<<(1+8),BitsWritten :7};_b [320]=code {Code :54<<8,BitsWritten :8};_b [384]=code {Code :55<<8,BitsWritten :8};_b [448]=code {Code :100<<8,BitsWritten :8};_b [512]=code {Code :101<<8,BitsWritten :8};_b [576]=code {Code :104<<8,BitsWritten :8};_b [640]=code {Code :103<<8,BitsWritten :8};_b [704]=code {Code :102<<8,BitsWritten :9};_b [768]=code {Code :102<<8|1<<7,BitsWritten :9};_b [832]=code {Code :105<<8,BitsWritten :9};_b [896]=code {Code :105<<8|1<<7,BitsWritten :9};_b [960]=code {Code :106<<8,BitsWritten :9};_b [1024]=code {Code :106<<8|1<<7,BitsWritten :9};_b [1088]=code {Code :107<<8,BitsWritten :9};_b [1152]=code {Code :107<<8|1<<7,BitsWritten :9};_b [1216]=code {Code :108<<8,BitsWritten :9};_b [1280]=code {Code :108<<8|1<<7,BitsWritten :9};_b [1344]=code {Code :109<<8,BitsWritten :9};_b [1408]=code {Code :109<<8|1<<7,BitsWritten :9};_b [1472]=code {Code :76<<8,BitsWritten :9};_b [1536]=code {Code :76<<8|1<<7,BitsWritten :9};_b [1600]=code {Code :77<<8,BitsWritten :9};_b [1664]=code {Code :24<<(2+8),BitsWritten :6};_b [1728]=code {Code :77<<8|1<<7,BitsWritten :9};_dbe =make (map[int ]code );_dbe [1792]=code {Code :1<<8,BitsWritten :11};_dbe [1856]=code {Code :1<<8|4<<5,BitsWritten :11};_dbe [1920]=code {Code :1<<8|5<<5,BitsWritten :11};_dbe [1984]=code {Code :1<<8|2<<4,BitsWritten :12};_dbe [2048]=code {Code :1<<8|3<<4,BitsWritten :12};_dbe [2112]=code {Code :1<<8|4<<4,BitsWritten :12};_dbe [2176]=code {Code :1<<8|5<<4,BitsWritten :12};_dbe [2240]=code {Code :1<<8|6<<4,BitsWritten :12};_dbe [2304]=code {Code :1<<8|7<<4,BitsWritten :12};_dbe [2368]=code {Code :1<<8|12<<4,BitsWritten :12};_dbe [2432]=code {Code :1<<8|13<<4,BitsWritten :12};_dbe [2496]=code {Code :1<<8|14<<4,BitsWritten :12};_dbe [2560]=code {Code :1<<8|15<<4,BitsWritten :12};_dg =make (map[int ]byte );_dg [0]=0xFF;_dg [1]=0xFE;_dg [2]=0xFC;_dg [3]=0xF8;_dg [4]=0xF0;_dg [5]=0xE0;_dg [6]=0xC0;_dg [7]=0x80;_dg [8]=0x00;};func _ebg (_geg []byte ,_bfd int )(bool ,int ,error ){_face :=_bfd ;var _efe bool ;_efe ,_bfd =_bbf (_geg ,_bfd );if _efe {_efe ,_bfd =_bbf (_geg ,_bfd );if _efe {return true ,_bfd ,nil ;}else {return false ,_face ,_fb ;};};return false ,_face ,nil ;};func _abbd (_aaa int )([]byte ,int ){var _cfe []byte ;for _aaef :=0;_aaef < 6;_aaef ++{_cfe ,_aaa =_bae (_cfe ,_aaa ,_ad );};return _cfe ,_aaa %8;};func init (){for _da ,_cc :=range _a {_ega (_gcg ,_cc ,0,_da );};for _ef ,_cg :=range _b {_ega (_gcg ,_cg ,0,_ef );};for _ec ,_efb :=range _e {_ega (_bd ,_efb ,0,_ec );};for _dbf ,_af :=range _db {_ega (_bd ,_af ,0,_dbf );};for _dab ,_be :=range _dbe {_ega (_gcg ,_be ,0,_dab );_ega (_bd ,_be ,0,_dab );};_ega (_gcf ,_dgg ,0,0);_ega (_gcf ,_ab ,0,0);_ega (_gcf ,_ba ,0,0);_ega (_gcf ,_ff ,0,0);_ega (_gcf ,_ga ,0,0);_ega (_gcf ,_gb ,0,0);_ega (_gcf ,_fg ,0,0);_ega (_gcf ,_gc ,0,0);_ega (_gcf ,_ae ,0,0);};func (_aefe *Encoder )appendEncodedRow (_gcd ,_faa []byte ,_dcbd int )[]byte {if len (_gcd )> 0&&_dcbd !=0&&!_aefe .EncodedByteAlign {_gcd [len (_gcd )-1]=_gcd [len (_gcd )-1]|_faa [0];_gcd =append (_gcd ,_faa [1:]...);}else {_gcd =append (_gcd ,_faa ...);};return _gcd ;};func _bea (_aaf ,_bgcf []byte ,_gg int ,_cde bool )([]byte ,int ,error ){_eeb :=_gg ;var _fac int ;for _fac ,_gg =_de (_aaf ,_gg ,_cde );_fac !=-1;_fac ,_gg =_de (_aaf ,_gg ,_cde ){_bgcf =_dabg (_bgcf ,_cde ,_fac );if _fac < 64{break ;};};if _fac ==-1{return _bgcf ,_eeb ,_gd ;};return _bgcf ,_gg ,nil ;};func _cfb (_cda []byte ,_egb int )int {if _egb >=len (_cda ){return _egb ;};if _egb < -1{_egb =-1;};var _aecc byte ;if _egb > -1{_aecc =_cda [_egb ];}else {_aecc =_eeg ;};_fcc :=_egb +1;for _fcc < len (_cda ){if _cda [_fcc ]!=_aecc {break ;};_fcc ++;};return _fcc ;};func _dabg (_afgg []byte ,_cea bool ,_fgc int )[]byte {if _fgc < 0{return _afgg ;};_fgg :=make ([]byte ,_fgc );if _cea {for _ddg :=0;_ddg < len (_fgg );_ddg ++{_fgg [_ddg ]=_eeg ;};}else {for _gcfd :=0;_gcfd < len (_fgg );_gcfd ++{_fgg [_gcfd ]=_cebb ;};};_afgg =append (_afgg ,_fgg ...);return _afgg ;};func _fd (_cgc uint16 ,_ffg int )byte {if _ffg < 8{_cgc >>=8;};_ffg %=8;_dga :=byte (0x01<<(7-uint (_ffg )));return (byte (_cgc )&_dga )>>(7-uint (_ffg ));};func _gddf (_edb []byte ,_dbee int ,_cae code )(bool ,int ){_ddga :=_dbee ;var (_cee uint16 ;_cbc int ;);_cee ,_cbc ,_dbee =_fgb (_edb ,_dbee );if _cbc > 3{return false ,_ddga ;};_cee >>=uint (3-_cbc );_cee <<=3;if _cee !=_cae .Code {return false ,_ddga ;}else {return true ,_dbee -3+_cbc ;};};func _ace (_eda [][]byte ,_gec []byte ,_bgd bool ,_ee ,_eaf int )([]byte ,int ){_abb :=_ecb (_gec ,_eda [len (_eda )-1],_ee ,_bgd );_bgb :=_abb +_eaf ;if _ee ==-1{_gec =_dabg (_gec ,_bgd ,_bgb -_ee -1);}else {_gec =_dabg (_gec ,_bgd ,_bgb -_ee );};_ee =_bgb ;return _gec ,_ee ;};func _abba (_gge []byte ,_bbd int )(bool ,int ,error ){_fgdd :=_bbd ;var _fad =false ;for _cfa :=0;_cfa < 6;_cfa ++{_fad ,_bbd =_cbf (_gge ,_bbd );if !_fad {if _cfa > 1{return false ,_fgdd ,_ge ;}else {_bbd =_fgdd ;break ;};};};return _fad ,_bbd ,nil ;};func _fdf (_eegd int )([]byte ,int ){var _dba []byte ;for _eegc :=0;_eegc < 6;_eegc ++{_dba ,_eegd =_bae (_dba ,_eegd ,_ea );};return _dba ,_eegd %8;};func _cgd (_fbf *decodingTreeNode ,_dcdc uint16 ,_cebg int )(*int ,*code ){if _fbf ==nil {return nil ,nil ;};if _cebg ==16{return _fbf .RunLen ,_fbf .Code ;};_begc :=_fd (_dcdc ,_cebg );_cebg ++;var _cab *int ;var _bge *code ;if _begc ==1{_cab ,_bge =_cgd (_fbf .Right ,_dcdc ,_cebg );}else {_cab ,_bge =_cgd (_fbf .Left ,_dcdc ,_cebg );};if _cab ==nil {_cab =_fbf .RunLen ;_bge =_fbf .Code ;};return _cab ,_bge ;};func _ecac (_baea ,_eacf int )code {var _faf code ;switch _eacf -_baea {case -1:_faf =_ff ;case -2:_faf =_ga ;case -3:_faf =_gb ;case 0:_faf =_ba ;case 1:_faf =_fg ;case 2:_faf =_gc ;case 3:_faf =_ae ;};return _faf ;};func _fgb (_afg []byte ,_ccb int )(uint16 ,int ,int ){_babc :=_ccb ;_efd :=_ccb /8;_ccb %=8;if _efd >=len (_afg ){return 0,16,_babc ;};_fgf :=byte (0xFF>>uint (_ccb ));_cbb :=uint16 ((_afg [_efd ]&_fgf )<<uint (_ccb ))<<8;_cfgb :=8-_ccb ;_efd ++;_ccb =0;if _efd >=len (_afg ){return _cbb >>(16-uint (_cfgb )),16-_cfgb ,_babc +_cfgb ;};_cbb |=uint16 (_afg [_efd ])<<(8-uint (_cfgb ));_cfgb +=8;_efd ++;_ccb =0;if _efd >=len (_afg ){return _cbb >>(16-uint (_cfgb )),16-_cfgb ,_babc +_cfgb ;};if _cfgb ==16{return _cbb ,0,_babc +_cfgb ;};_cec :=16-_cfgb ;_cbb |=uint16 (_afg [_efd ]>>(8-uint (_cec )));return _cbb ,0,_babc +16;};func (_eb *Encoder )decodeG31D (_bab []byte )([][]byte ,error ){var _cf [][]byte ;var _cad int ;for (_cad /8)< len (_bab ){var _ffd bool ;_ffd ,_cad =_bbf (_bab ,_cad );if !_ffd {if _eb .EndOfLine {return nil ,_abc ;};}else {for _fbg :=0;_fbg < 5;_fbg ++{_ffd ,_cad =_bbf (_bab ,_cad );if !_ffd {if _fbg ==0{break ;};return nil ,_aa ;};};if _ffd {break ;};};var _cd []byte ;_cd ,_cad =_eb .decodeRow1D (_bab ,_cad );if _eb .EncodedByteAlign &&_cad %8!=0{_cad +=8-_cad %8;};_cf =append (_cf ,_cd );if _eb .Rows > 0&&!_eb .EndOfBlock &&len (_cf )>=_eb .Rows {break ;};};return _cf ,nil ;};func _bbaf (_fda ,_ece []byte ,_cabf ,_gee ,_cba int )([]byte ,int ,int ){_acgg :=_cfb (_fda ,_cba );_fcd :=_gee >=0&&_fda [_gee ]==_eeg ||_gee ==-1;_ece ,_cabf =_bae (_ece ,_cabf ,_ab );var _acgf int ;if _gee > -1{_acgf =_cba -_gee ;}else {_acgf =_cba -_gee -1;};_ece ,_cabf =_eed (_ece ,_cabf ,_acgf ,_fcd );_fcd =!_fcd ;_efa :=_acgg -_cba ;_ece ,_cabf =_eed (_ece ,_cabf ,_efa ,_fcd );_gee =_acgg ;return _ece ,_cabf ,_gee ;};func (_edbf *Encoder )encodeG32D (_ccf [][]byte )[]byte {var _bed []byte ;var _cdca int ;for _ggf :=0;_ggf < len (_ccf );_ggf +=_edbf .K {if _edbf .Rows > 0&&!_edbf .EndOfBlock &&_ggf ==_edbf .Rows {break ;};_gcc ,_egag :=_ffbf (_ccf [_ggf ],_cdca ,_ea );_bed =_edbf .appendEncodedRow (_bed ,_gcc ,_cdca );if _edbf .EncodedByteAlign {_egag =0;};_cdca =_egag ;for _bgbc :=_ggf +1;_bgbc < (_ggf +_edbf .K )&&_bgbc < len (_ccf );_bgbc ++{if _edbf .Rows > 0&&!_edbf .EndOfBlock &&_bgbc ==_edbf .Rows {break ;};_edga ,_efg :=_bae (nil ,_cdca ,_ac );var _gfg ,_aafc ,_dafa int ;_fag :=-1;for _fag < len (_ccf [_bgbc ]){_gfg =_cfb (_ccf [_bgbc ],_fag );_aafc =_fdg (_ccf [_bgbc ],_ccf [_bgbc -1],_fag );_dafa =_cfb (_ccf [_bgbc -1],_aafc );if _dafa < _gfg {_edga ,_efg =_dbag (_edga ,_efg );_fag =_dafa ;}else {if _g .Abs (float64 (_aafc -_gfg ))> 3{_edga ,_efg ,_fag =_bbaf (_ccf [_bgbc ],_edga ,_efg ,_fag ,_gfg );}else {_edga ,_efg =_dbeb (_edga ,_efg ,_gfg ,_aafc );_fag =_gfg ;};};};_bed =_edbf .appendEncodedRow (_bed ,_edga ,_cdca );if _edbf .EncodedByteAlign {_efg =0;};_cdca =_efg %8;};};if _edbf .EndOfBlock {_cfag ,_ :=_fdf (_cdca );_bed =_edbf .appendEncodedRow (_bed ,_cfag ,_cdca );};return _bed ;};func _cge (_abcc []byte ,_agf int )(bool ,int ){return _gddf (_abcc ,_agf ,_ac )};type decodingTreeNode struct{Val byte ;RunLen *int ;Code *code ;Left *decodingTreeNode ;Right *decodingTreeNode ;};func _ffbf (_gbfg []byte ,_ead int ,_babb code )([]byte ,int ){_gaga :=true ;var _bdc []byte ;_bdc ,_ead =_bae (nil ,_ead ,_babb );_ceab :=0;var _eca int ;for _ceab < len (_gbfg ){_eca ,_ceab =_abd (_gbfg ,_gaga ,_ceab );_bdc ,_ead =_eed (_bdc ,_ead ,_eca ,_gaga );_gaga =!_gaga ;};return _bdc ,_ead %8;};func _bbf (_aeba []byte ,_dbc int )(bool ,int ){_bbe :=_dbc ;var (_daf uint16 ;_gfd int ;);_daf ,_gfd ,_dbc =_fgb (_aeba ,_dbc );if _gfd > 4{return false ,_bbe ;};_daf >>=uint (4-_gfd );_daf <<=4;if _daf !=_ad .Code {return false ,_bbe ;}else {return true ,_dbc -4+_gfd ;};};var (_fb =_f .New ("\u0045\u004f\u0046\u0042 c\u006f\u0064\u0065\u0020\u0069\u0073\u0020\u0063\u006f\u0072\u0072\u0075\u0070\u0074e\u0064");_ge =_f .New ("R\u0054\u0043\u0020\u0063od\u0065 \u0069\u0073\u0020\u0063\u006fr\u0072\u0075\u0070\u0074\u0065\u0064");_gd =_f .New ("\u0077\u0072o\u006e\u0067\u0020\u0063\u006f\u0064\u0065\u0020\u0069\u006e\u0020\u0068\u006f\u0072\u0069\u007a\u006f\u006e\u0074\u0061\u006c\u0020mo\u0064\u0065");_abc =_f .New ("\u006e\u006f\u0020\u0045\u004f\u004c\u0020\u0066\u006f\u0075\u006e\u0064\u0020\u0077\u0068\u0069\u006c\u0065 \u0074\u0068\u0065\u0020\u0045\u006e\u0064O\u0066\u004c\u0069\u006e\u0065\u0020\u0070\u0061\u0072\u0061\u006de\u0074\u0065\u0072\u0020\u0069\u0073\u0020\u0074\u0072\u0075\u0065");_aa =_f .New ("i\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0045\u004f\u004c");_fe =_f .New ("\u0069n\u0076a\u006c\u0069\u0064\u0020\u0032\u0044\u0020\u0063\u006f\u0064\u0065"););func _dbeb (_gega []byte ,_aeg ,_dec ,_cffd int )([]byte ,int ){_abfb :=_ecac (_dec ,_cffd );_gega ,_aeg =_bae (_gega ,_aeg ,_abfb );return _gega ,_aeg ;};func (_beg *Encoder )decodeG4 (_beb []byte )([][]byte ,error ){_gdb :=make ([]byte ,_beg .Columns );for _bg :=range _gdb {_gdb [_bg ]=_eeg ;};_acg :=make ([][]byte ,1);_acg [0]=_gdb ;var (_fgd bool ;_dbb error ;_ebc int ;);for (_ebc /8)< len (_beb ){_fgd ,_ebc ,_dbb =_ebg (_beb ,_ebc );if _dbb !=nil {return nil ,_dbb ;};if _fgd {break ;};var (_agc code ;_gdd bool ;);_cac :=true ;var _bgc []byte ;_cff :=-1;for _cff < _beg .Columns {_agc ,_ebc ,_gdd =_gdbg (_beb ,_ebc );if !_gdd {return nil ,_fe ;};switch _agc {case _dgg :_bgc ,_cff =_dae (_acg ,_bgc ,_cac ,_cff );case _ab :_bgc ,_ebc ,_cff ,_dbb =_caf (_beb ,_bgc ,_ebc ,_cac ,_cff );if _dbb !=nil {return nil ,_dbb ;};case _ba :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,0);_cac =!_cac ;case _ff :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,1);_cac =!_cac ;case _ga :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,2);_cac =!_cac ;case _gb :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,3);_cac =!_cac ;case _fg :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,-1);_cac =!_cac ;case _gc :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,-2);_cac =!_cac ;case _ae :_bgc ,_cff =_ace (_acg ,_bgc ,_cac ,_cff ,-3);_cac =!_cac ;};if len (_bgc )>=_beg .Columns {break ;};};if _beg .EncodedByteAlign &&_ebc %8!=0{_ebc +=8-_ebc %8;};_acg =append (_acg ,_bgc );if _beg .Rows > 0&&!_beg .EndOfBlock &&len (_acg )>=(_beg .Rows +1){break ;};};_acg =_acg [1:];return _acg ,nil ;};func _dae (_gbg [][]byte ,_cb []byte ,_bec bool ,_dde int )([]byte ,int ){_cdc :=_ecb (_cb ,_gbg [len (_gbg )-1],_dde ,_bec );_geb :=_cfb (_gbg [len (_gbg )-1],_cdc );if _dde ==-1{_cb =_dabg (_cb ,_bec ,_geb -_dde -1);}else {_cb =_dabg (_cb ,_bec ,_geb -_dde );};_dde =_geb ;return _cb ,_dde ;};func (_ebd *Encoder )Encode (pixels [][]byte )[]byte {if _ebd .BlackIs1 {_eeg =0;_cebb =1;}else {_eeg =1;_cebb =0;};if _ebd .K ==0{return _ebd .encodeG31D (pixels );};if _ebd .K > 0{return _ebd .encodeG32D (pixels );};if _ebd .K < 0{return _ebd .encodeG4 (pixels );};return nil ;};func _cbf (_aea []byte ,_gebc int )(bool ,int ){return _gddf (_aea ,_gebc ,_ea )};func _de (_bdb []byte ,_bgde int ,_aec bool )(int ,int ){var (_feg uint16 ;_abca int ;_cgg int ;);_cgg =_bgde ;_feg ,_abca ,_bgde =_fgb (_bdb ,_bgde );_bc ,_gf :=_aag (_feg ,_abca ,_aec );if _bc ==-1{return -1,_cgg ;};return _bc ,_cgg +_gf .BitsWritten ;};var (_gcg =&decodingTreeNode {Val :255};_bd =&decodingTreeNode {Val :255};_gcf =&decodingTreeNode {Val :255};);var (_eeg byte =1;_cebb byte =0;);func _gdbg (_cga []byte ,_gga int )(code ,int ,bool ){var (_dcd uint16 ;_feb int ;_fab int ;);_fab =_gga ;_dcd ,_feb ,_gga =_fgb (_cga ,_gga );_deb ,_eg :=_ecc (_dcd ,_feb );if !_eg {return code {},_fab ,false ;};return _deb ,_fab +_deb .BitsWritten ,true ;};func _aag (_ffb uint16 ,_dda int ,_gebb bool )(int ,code ){var _afe *int ;var _edd *code ;if _gebb {_afe ,_edd =_cgd (_gcg ,_ffb ,_dda );}else {_afe ,_edd =_cgd (_bd ,_ffb ,_dda );};if _afe ==nil {return -1,code {};};return *_afe ,*_edd ;};type Encoder struct{K int ;EndOfLine bool ;EncodedByteAlign bool ;Columns int ;Rows int ;EndOfBlock bool ;BlackIs1 bool ;DamagedRowsBeforeError int ;};func _eed (_afd []byte ,_gfe int ,_debgc int ,_facb bool )([]byte ,int ){var (_gcfb code ;_bad bool ;);for !_bad {_gcfb ,_debgc ,_bad =_cdb (_debgc ,_facb );_afd ,_gfe =_bae (_afd ,_gfe ,_gcfb );};return _afd ,_gfe ;};func _abd (_bba []byte ,_bdd bool ,_cfda int )(int ,int ){_bde :=0;for _cfda < len (_bba ){if _bdd {if _bba [_cfda ]!=_eeg {break ;};}else {if _bba [_cfda ]!=_cebb {break ;};};_bde ++;_cfda ++;};return _bde ,_cfda ;};func (_fc *Encoder )decodeRow1D (_acb []byte ,_eac int )([]byte ,int ){var _debg []byte ;_aeb :=true ;var _def int ;_def ,_eac =_de (_acb ,_eac ,_aeb );for _def !=-1{_debg =_dabg (_debg ,_aeb ,_def );if _def < 64{if len (_debg )>=_fc .Columns {break ;};_aeb =!_aeb ;};_def ,_eac =_de (_acb ,_eac ,_aeb );};return _debg ,_eac ;};func (_ca *Encoder )Decode (encoded []byte )([][]byte ,error ){if _ca .BlackIs1 {_eeg =0;_cebb =1;}else {_eeg =1;_cebb =0;};if _ca .K ==0{return _ca .decodeG31D (encoded );};if _ca .K > 0{return _ca .decodeG32D (encoded );};if _ca .K < 4{return _ca .decodeG4 (encoded );};return nil ,nil ;};func _bae (_efef []byte ,_bebc int ,_dgb code )([]byte ,int ){_ecab :=0;for _ecab < _dgb .BitsWritten {_cgcc :=_bebc /8;_cgee :=_bebc %8;if _cgcc >=len (_efef ){_efef =append (_efef ,0);};_fgefc :=8-_cgee ;_fbbc :=_dgb .BitsWritten -_ecab ;if _fgefc > _fbbc {_fgefc =_fbbc ;};if _ecab < 8{_efef [_cgcc ]=_efef [_cgcc ]|byte (_dgb .Code >>uint (8+_cgee -_ecab ))&_dg [8-_fgefc -_cgee ];}else {_efef [_cgcc ]=_efef [_cgcc ]|(byte (_dgb .Code <<uint (_ecab -8))&_dg [8-_fgefc ])>>uint (_cgee );};_bebc +=_fgefc ;_ecab +=_fgefc ;};return _efef ,_bebc ;};func (_bfb *Encoder )encodeG4 (_bebb [][]byte )[]byte {_cbg :=make ([][]byte ,len (_bebb ));copy (_cbg ,_bebb );_cbg =_eba (_cbg );var _ddab []byte ;var _eaa int ;for _cag :=1;_cag < len (_cbg );_cag ++{if _bfb .Rows > 0&&!_bfb .EndOfBlock &&_cag ==(_bfb .Rows +1){break ;};var _eag []byte ;var _cdcc ,_feeg ,_fdb int ;_cdg :=_eaa ;_ccbf :=-1;for _ccbf < len (_cbg [_cag ]){_cdcc =_cfb (_cbg [_cag ],_ccbf );_feeg =_fdg (_cbg [_cag ],_cbg [_cag -1],_ccbf );_fdb =_cfb (_cbg [_cag -1],_feeg );if _fdb < _cdcc {_eag ,_cdg =_bae (_eag ,_cdg ,_dgg );_ccbf =_fdb ;}else {if _g .Abs (float64 (_feeg -_cdcc ))> 3{_eag ,_cdg ,_ccbf =_bbaf (_cbg [_cag ],_eag ,_cdg ,_ccbf ,_cdcc );}else {_eag ,_cdg =_dbeb (_eag ,_cdg ,_cdcc ,_feeg );_ccbf =_cdcc ;};};};_ddab =_bfb .appendEncodedRow (_ddab ,_eag ,_eaa );if _bfb .EncodedByteAlign {_cdg =0;};_eaa =_cdg %8;};if _bfb .EndOfBlock {_cdef ,_ :=_eefd (_eaa );_ddab =_bfb .appendEncodedRow (_ddab ,_cdef ,_eaa );};return _ddab ;};func _ecc (_daa uint16 ,_ada int )(code ,bool ){_ ,_cfg :=_cgd (_gcf ,_daa ,_ada );if _cfg ==nil {return code {},false ;};return *_cfg ,true ;};func (_dd *Encoder )decodeG32D (_aac []byte )([][]byte ,error ){var (_ag [][]byte ;_bf int ;_eae error ;);_add :for (_bf /8)< len (_aac ){var _fge bool ;_fge ,_bf ,_eae =_abba (_aac ,_bf );if _eae !=nil {return nil ,_eae ;};if _fge {break ;};_fge ,_bf =_cbf (_aac ,_bf );if !_fge {if _dd .EndOfLine {return nil ,_abc ;};};var _bfg []byte ;_bfg ,_bf =_dd .decodeRow1D (_aac ,_bf );if _dd .EncodedByteAlign &&_bf %8!=0{_bf +=8-_bf %8;};if _bfg !=nil {_ag =append (_ag ,_bfg );};if _dd .Rows > 0&&!_dd .EndOfBlock &&len (_ag )>=_dd .Rows {break ;};for _bb :=1;_bb < _dd .K &&(_bf /8)< len (_aac );_bb ++{_fge ,_bf =_cge (_aac ,_bf );if !_fge {_fge ,_bf ,_eae =_abba (_aac ,_bf );if _eae !=nil {return nil ,_eae ;};if _fge {break _add ;}else {if _dd .EndOfLine {return nil ,_abc ;};};};var (_ce code ;_ed bool ;);_fbb :=true ;var _fec []byte ;_ceb :=-1;for _ce ,_bf ,_ed =_gdbg (_aac ,_bf );_ed ;_ce ,_bf ,_ed =_gdbg (_aac ,_bf ){switch _ce {case _dgg :_fec ,_ceb =_dae (_ag ,_fec ,_fbb ,_ceb );case _ab :_fec ,_bf ,_ceb ,_eae =_caf (_aac ,_fec ,_bf ,_fbb ,_ceb );if _eae !=nil {return nil ,_eae ;};case _ba :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,0);_fbb =!_fbb ;case _ff :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,1);_fbb =!_fbb ;case _ga :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,2);_fbb =!_fbb ;case _gb :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,3);_fbb =!_fbb ;case _fg :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,-1);_fbb =!_fbb ;case _gc :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,-2);_fbb =!_fbb ;case _ae :_fec ,_ceb =_ace (_ag ,_fec ,_fbb ,_ceb ,-3);_fbb =!_fbb ;};if len (_fec )>=_dd .Columns {break ;};};if _dd .EncodedByteAlign &&_bf %8!=0{_bf +=8-_bf %8;};if _fec !=nil {_ag =append (_ag ,_fec );};if _dd .Rows > 0&&!_dd .EndOfBlock &&len (_ag )>=_dd .Rows {break _add ;};};};return _ag ,nil ;};var (_e map[int ]code ;_a map[int ]code ;_db map[int ]code ;_b map[int ]code ;_dbe map[int ]code ;_dg map[int ]byte ;_ad =code {Code :1<<4,BitsWritten :12};_ea =code {Code :3<<3,BitsWritten :13};_ac =code {Code :2<<3,BitsWritten :13};_dgg =code {Code :1<<12,BitsWritten :4};_ab =code {Code :1<<13,BitsWritten :3};_ba =code {Code :1<<15,BitsWritten :1};_ff =code {Code :3<<13,BitsWritten :3};_ga =code {Code :3<<10,BitsWritten :6};_gb =code {Code :3<<9,BitsWritten :7};_fg =code {Code :2<<13,BitsWritten :3};_gc =code {Code :2<<10,BitsWritten :6};_ae =code {Code :2<<9,BitsWritten :7};);type code struct{Code uint16 ;BitsWritten int ;};func _ecb (_dce ,_adaa []byte ,_gcdf int ,_dca bool )int {_eebe :=_cfb (_adaa ,_gcdf );if _eebe < len (_adaa )&&(_gcdf ==-1&&_adaa [_eebe ]==_eeg ||_gcdf >=0&&_gcdf < len (_dce )&&_dce [_gcdf ]==_adaa [_eebe ]||_gcdf >=len (_dce )&&_dca &&_adaa [_eebe ]==_eeg ||_gcdf >=len (_dce )&&!_dca &&_adaa [_eebe ]==_cebb ){_eebe =_cfb (_adaa ,_eebe );};return _eebe ;};func _eba (_gff [][]byte )[][]byte {_bbg :=make ([]byte ,len (_gff [0]));for _cged :=range _bbg {_bbg [_cged ]=_eeg ;};_gff =append (_gff ,[]byte {});for _ded :=len (_gff )-1;_ded > 0;_ded --{_gff [_ded ]=_gff [_ded -1];};_gff [0]=_bbg ;return _gff ;};func (_dafg *Encoder )encodeG31D (_dcb [][]byte )[]byte {var _gbf []byte ;_gab :=0;for _edf :=range _dcb {if _dafg .Rows > 0&&!_dafg .EndOfBlock &&_edf ==_dafg .Rows {break ;};_fee ,_dafb :=_ffbf (_dcb [_edf ],_gab ,_ad );_gbf =_dafg .appendEncodedRow (_gbf ,_fee ,_gab );if _dafg .EncodedByteAlign {_dafb =0;};_gab =_dafb ;};if _dafg .EndOfBlock {_bdg ,_ :=_abbd (_gab );_gbf =_dafg .appendEncodedRow (_gbf ,_bdg ,_gab );};return _gbf ;};func _dbag (_defa []byte ,_fdd int )([]byte ,int ){return _bae (_defa ,_fdd ,_dgg )};func _cdb (_aef int ,_efgg bool )(code ,int ,bool ){if _aef < 64{if _efgg {return _a [_aef ],0,true ;}else {return _e [_aef ],0,true ;};}else {_adb :=_aef /64;if _adb > 40{return _dbe [2560],_aef -2560,false ;};if _adb > 27{return _dbe [_adb *64],_aef -_adb *64,false ;};if _efgg {return _b [_adb *64],_aef -_adb *64,false ;}else {return _db [_adb *64],_aef -_adb *64,false ;};};};func _eefd (_edgd int )([]byte ,int ){var _adc []byte ;for _dgaf :=0;_dgaf < 2;_dgaf ++{_adc ,_edgd =_bae (_adc ,_edgd ,_ad );};return _adc ,_edgd %8;};func _fdg (_dad ,_adce []byte ,_abf int )int {_agcf :=_cfb (_adce ,_abf );if _agcf < len (_adce )&&(_abf ==-1&&_adce [_agcf ]==_eeg ||_abf >=0&&_abf < len (_dad )&&_dad [_abf ]==_adce [_agcf ]||_abf >=len (_dad )&&_dad [_abf -1]!=_adce [_agcf ]){_agcf =_cfb (_adce ,_agcf );};return _agcf ;};