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

package endian ;import (_d "encoding/binary";_e "unsafe";);func IsBig ()bool {return _g };func init (){const _ea =int (_e .Sizeof (0));_a :=1;_f :=(*[_ea ]byte )(_e .Pointer (&_a ));if _f [0]==0{_g =true ;ByteOrder =_d .BigEndian ;}else {ByteOrder =_d .LittleEndian ;
};};func IsLittle ()bool {return !_g };var (ByteOrder _d .ByteOrder ;_g bool ;);