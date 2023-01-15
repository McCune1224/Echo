import React, { Component } from "react";
import Login from "./Login";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { useCookies } from 'react-cookie';

export default class Auth extends Component { 

    CheckAuthCode = () =>
    {  
        var path = window.location.pathname;
        if(!path.includes("code="))
            return;

        var args = path.split("code=")[1];
        console.log("Args : " + args);

        var code = args.split('&')[0];
        console.log("Code : " + code);

        var service = path.split("service=")[1];
        console.log("Service : " + service);

        //Now add to cookie
        const [cookies, setCookie] = useCookies([service]); 

        setCookie(service, code, { path: '/', expires: new Date(Date.now() + 1000)});
    }

    render() {  
        this.CheckAuthCode();


        const Container = () =>
        {
            return ( 
                <div className=" backdrop-blur-md bg-black/30 rounded-3xl
                w-1/3 min-w-[450px] max-w-[550px]
                h-3/4 min-h-[700px] max-h-[1200px]
                content-center
                flex">  
                    <Login/>
                </div>
            )
        }

        const Background = () =>
        {
            return (
                <div className="bg-cover bg-authBackground h-screen grid place-items-center">  
                    <Container/>
                </div>
            )    
        }

        return( 
            <div className=""> 
                <Background/>
            </div>
        )    
    }
}
