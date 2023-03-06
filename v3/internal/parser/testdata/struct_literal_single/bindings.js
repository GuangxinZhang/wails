// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

import {main} from './models';

function GreetService(method) {
    return {
        packageName: "main",
        serviceName: "GreetService",
        methodName: method,
        args: Array.prototype.slice.call(arguments, 1),
    };
}

/**
 * GreetService.ArrayInt
 * 
 * @param _in {number[]}
 * @returns {Promise<void>}
 **/
function ArrayInt(_in) {
    return wails.Call(GreetService("ArrayInt", _in));
}

/**
 * GreetService.BoolInBoolOut
 * 
 * @param _in {boolean}
 * @returns {Promise<boolean>}
 **/
function BoolInBoolOut(_in) {
    return wails.Call(GreetService("BoolInBoolOut", _in));
}

/**
 * GreetService.Float32InFloat32Out
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function Float32InFloat32Out(_in) {
    return wails.Call(GreetService("Float32InFloat32Out", _in));
}

/**
 * GreetService.Float64InFloat64Out
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function Float64InFloat64Out(_in) {
    return wails.Call(GreetService("Float64InFloat64Out", _in));
}

/**
 * GreetService.Greet
 * Greet someone
 * @param name {string}
 * @returns {Promise<string>}
 **/
function Greet(name) {
    return wails.Call(GreetService("Greet", name));
}

/**
 * GreetService.Int16InIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function Int16InIntOut(_in) {
    return wails.Call(GreetService("Int16InIntOut", _in));
}

/**
 * GreetService.Int16PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function Int16PointerInAndOutput(_in) {
    return wails.Call(GreetService("Int16PointerInAndOutput", _in));
}

/**
 * GreetService.Int32InIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function Int32InIntOut(_in) {
    return wails.Call(GreetService("Int32InIntOut", _in));
}

/**
 * GreetService.Int32PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function Int32PointerInAndOutput(_in) {
    return wails.Call(GreetService("Int32PointerInAndOutput", _in));
}

/**
 * GreetService.Int64InIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function Int64InIntOut(_in) {
    return wails.Call(GreetService("Int64InIntOut", _in));
}

/**
 * GreetService.Int64PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function Int64PointerInAndOutput(_in) {
    return wails.Call(GreetService("Int64PointerInAndOutput", _in));
}

/**
 * GreetService.Int8InIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function Int8InIntOut(_in) {
    return wails.Call(GreetService("Int8InIntOut", _in));
}

/**
 * GreetService.Int8PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function Int8PointerInAndOutput(_in) {
    return wails.Call(GreetService("Int8PointerInAndOutput", _in));
}

/**
 * GreetService.IntInIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function IntInIntOut(_in) {
    return wails.Call(GreetService("IntInIntOut", _in));
}

/**
 * GreetService.IntPointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function IntPointerInAndOutput(_in) {
    return wails.Call(GreetService("IntPointerInAndOutput", _in));
}

/**
 * GreetService.IntPointerInputNamedOutputs
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null, void>}
 **/
function IntPointerInputNamedOutputs(_in) {
    return wails.Call(GreetService("IntPointerInputNamedOutputs", _in));
}

/**
 * GreetService.MapIntInt
 * 
 * @param _in {map}
 * @returns {Promise<void>}
 **/
function MapIntInt(_in) {
    return wails.Call(GreetService("MapIntInt", _in));
}

/**
 * GreetService.MapIntPointerInt
 * 
 * @param _in {map}
 * @returns {Promise<void>}
 **/
function MapIntPointerInt(_in) {
    return wails.Call(GreetService("MapIntPointerInt", _in));
}

/**
 * GreetService.MapIntSliceInt
 * 
 * @param _in {map}
 * @returns {Promise<void>}
 **/
function MapIntSliceInt(_in) {
    return wails.Call(GreetService("MapIntSliceInt", _in));
}

/**
 * GreetService.MapIntSliceIntInMapIntSliceIntOut
 * 
 * @param _in {map}
 * @returns {Promise<map>}
 **/
function MapIntSliceIntInMapIntSliceIntOut(_in) {
    return wails.Call(GreetService("MapIntSliceIntInMapIntSliceIntOut", _in));
}

/**
 * GreetService.NoInputsStringOut
 * 
 *
 * @returns {Promise<string>}
 **/
function NoInputsStringOut() {
    return wails.Call(GreetService("NoInputsStringOut"));
}

/**
 * GreetService.PointerBoolInBoolOut
 * 
 * @param _in {boolean | null}
 * @returns {Promise<boolean | null>}
 **/
function PointerBoolInBoolOut(_in) {
    return wails.Call(GreetService("PointerBoolInBoolOut", _in));
}

/**
 * GreetService.PointerFloat32InFloat32Out
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function PointerFloat32InFloat32Out(_in) {
    return wails.Call(GreetService("PointerFloat32InFloat32Out", _in));
}

/**
 * GreetService.PointerFloat64InFloat64Out
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function PointerFloat64InFloat64Out(_in) {
    return wails.Call(GreetService("PointerFloat64InFloat64Out", _in));
}

/**
 * GreetService.PointerMapIntInt
 * 
 * @param _in {map | null}
 * @returns {Promise<void>}
 **/
function PointerMapIntInt(_in) {
    return wails.Call(GreetService("PointerMapIntInt", _in));
}

/**
 * GreetService.PointerStringInStringOut
 * 
 * @param _in {string | null}
 * @returns {Promise<string | null>}
 **/
function PointerStringInStringOut(_in) {
    return wails.Call(GreetService("PointerStringInStringOut", _in));
}

/**
 * GreetService.StringArrayInputNamedOutput
 * 
 * @param _in {string[]}
 * @returns {Promise<string[]>}
 **/
function StringArrayInputNamedOutput(_in) {
    return wails.Call(GreetService("StringArrayInputNamedOutput", _in));
}

/**
 * GreetService.StringArrayInputNamedOutputs
 * 
 * @param _in {string[]}
 * @returns {Promise<string[], void>}
 **/
function StringArrayInputNamedOutputs(_in) {
    return wails.Call(GreetService("StringArrayInputNamedOutputs", _in));
}

/**
 * GreetService.StringArrayInputStringArrayOut
 * 
 * @param _in {string[]}
 * @returns {Promise<string[]>}
 **/
function StringArrayInputStringArrayOut(_in) {
    return wails.Call(GreetService("StringArrayInputStringArrayOut", _in));
}

/**
 * GreetService.StringArrayInputStringOut
 * 
 * @param _in {string[]}
 * @returns {Promise<string>}
 **/
function StringArrayInputStringOut(_in) {
    return wails.Call(GreetService("StringArrayInputStringOut", _in));
}

/**
 * GreetService.StructInputStructOutput
 * 
 * @param _in {main.Person}
 * @returns {Promise<main.Person>}
 **/
function StructInputStructOutput(_in) {
    return wails.Call(GreetService("StructInputStructOutput", _in));
}

/**
 * GreetService.StructPointerInputErrorOutput
 * 
 * @param _in {main.Person | null}
 * @returns {Promise<void>}
 **/
function StructPointerInputErrorOutput(_in) {
    return wails.Call(GreetService("StructPointerInputErrorOutput", _in));
}

/**
 * GreetService.StructPointerInputStructPointerOutput
 * 
 * @param _in {main.Person | null}
 * @returns {Promise<main.Person | null>}
 **/
function StructPointerInputStructPointerOutput(_in) {
    return wails.Call(GreetService("StructPointerInputStructPointerOutput", _in));
}

/**
 * GreetService.UInt16InUIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function UInt16InUIntOut(_in) {
    return wails.Call(GreetService("UInt16InUIntOut", _in));
}

/**
 * GreetService.UInt16PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function UInt16PointerInAndOutput(_in) {
    return wails.Call(GreetService("UInt16PointerInAndOutput", _in));
}

/**
 * GreetService.UInt32InUIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function UInt32InUIntOut(_in) {
    return wails.Call(GreetService("UInt32InUIntOut", _in));
}

/**
 * GreetService.UInt32PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function UInt32PointerInAndOutput(_in) {
    return wails.Call(GreetService("UInt32PointerInAndOutput", _in));
}

/**
 * GreetService.UInt64InUIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function UInt64InUIntOut(_in) {
    return wails.Call(GreetService("UInt64InUIntOut", _in));
}

/**
 * GreetService.UInt64PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function UInt64PointerInAndOutput(_in) {
    return wails.Call(GreetService("UInt64PointerInAndOutput", _in));
}

/**
 * GreetService.UInt8InUIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function UInt8InUIntOut(_in) {
    return wails.Call(GreetService("UInt8InUIntOut", _in));
}

/**
 * GreetService.UInt8PointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function UInt8PointerInAndOutput(_in) {
    return wails.Call(GreetService("UInt8PointerInAndOutput", _in));
}

/**
 * GreetService.UIntInUIntOut
 * 
 * @param _in {number}
 * @returns {Promise<number>}
 **/
function UIntInUIntOut(_in) {
    return wails.Call(GreetService("UIntInUIntOut", _in));
}

/**
 * GreetService.UIntPointerInAndOutput
 * 
 * @param _in {number | null}
 * @returns {Promise<number | null>}
 **/
function UIntPointerInAndOutput(_in) {
    return wails.Call(GreetService("UIntPointerInAndOutput", _in));
}

window.go = window.go || {};
window.go.main = {
    GreetService: {
        ArrayInt,
        BoolInBoolOut,
        Float32InFloat32Out,
        Float64InFloat64Out,
        Greet,
        Int16InIntOut,
        Int16PointerInAndOutput,
        Int32InIntOut,
        Int32PointerInAndOutput,
        Int64InIntOut,
        Int64PointerInAndOutput,
        Int8InIntOut,
        Int8PointerInAndOutput,
        IntInIntOut,
        IntPointerInAndOutput,
        IntPointerInputNamedOutputs,
        MapIntInt,
        MapIntPointerInt,
        MapIntSliceInt,
        MapIntSliceIntInMapIntSliceIntOut,
        NoInputsStringOut,
        PointerBoolInBoolOut,
        PointerFloat32InFloat32Out,
        PointerFloat64InFloat64Out,
        PointerMapIntInt,
        PointerStringInStringOut,
        StringArrayInputNamedOutput,
        StringArrayInputNamedOutputs,
        StringArrayInputStringArrayOut,
        StringArrayInputStringOut,
        StructInputStructOutput,
        StructPointerInputErrorOutput,
        StructPointerInputStructPointerOutput,
        UInt16InUIntOut,
        UInt16PointerInAndOutput,
        UInt32InUIntOut,
        UInt32PointerInAndOutput,
        UInt64InUIntOut,
        UInt64PointerInAndOutput,
        UInt8InUIntOut,
        UInt8PointerInAndOutput,
        UIntInUIntOut,
        UIntPointerInAndOutput,
    },
};

