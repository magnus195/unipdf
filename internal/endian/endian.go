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

package endian ;import (_c "encoding/binary";_d "unsafe";);func IsLittle ()bool {return !_g };func init (){const _fd =int (_d .Sizeof (0));_fb :=1;_dg :=(*[_fd ]byte )(_d .Pointer (&_fb ));if _dg [0]==0{_g =true ;ByteOrder =_c .BigEndian ;}else {ByteOrder =_c .LittleEndian ;
};};var (ByteOrder _c .ByteOrder ;_g bool ;);func IsBig ()bool {return _g };