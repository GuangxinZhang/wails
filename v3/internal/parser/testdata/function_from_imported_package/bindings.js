// @ts-check
// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

import {main, services} from './models';

function OtherService(method) {
    return {
        packageName: "services",
        serviceName: "OtherService",
        methodName: method,
        args: Array.prototype.slice.call(arguments, 1),
    };
}

/**
 * OtherService.Yay
 * 
 *
 * @returns {Promise<services.Address | null>}
 **/
function Yay() {
    return wails.Call(OtherService("Yay"));
}
function GreetService(method) {
    return {
        packageName: "main",
        serviceName: "GreetService",
        methodName: method,
        args: Array.prototype.slice.call(arguments, 1),
    };
}

/**
 * GreetService.Greet
 * Greet does XYZ
 * @param name {string}
 * @returns {Promise<string>}
 **/
function Greet(name) {
    return wails.Call(GreetService("Greet", name));
}

/**
 * GreetService.NewPerson
 * NewPerson creates a new person
 * @param name {string}
 * @returns {Promise<main.Person | null>}
 **/
function NewPerson(name) {
    return wails.Call(GreetService("NewPerson", name));
}

window.go = window.go || {};
window.go.main = {
    GreetService: {
        Greet,
        NewPerson,
    },
};
window.go.services = {
    OtherService: {
        Yay,
    },
};

