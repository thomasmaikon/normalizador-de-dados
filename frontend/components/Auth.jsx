import Cookies from "js-cookie";
import { createContext, useState } from "react";

export const MyContextAuth = createContext(null)

export default function Auth({children}){

    function saveToken(newToken){
        const token = newToken
        Cookies.set('hubla-token', token)
    }

    function getToken(){
        return Cookies.get('hubla-token')
    }
    return (
        <MyContextAuth.Provider value={{getToken, saveToken}}>
            {children}
        </MyContextAuth.Provider>
   )
}