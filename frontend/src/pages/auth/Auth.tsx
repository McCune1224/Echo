import React, { Component } from "react";
import Login from "./Login";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

export default class Auth extends Component {


    render() {
        
        var login = new Login();

        const Container = () =>
        {
            return ( 
                <div className=" backdrop-blur-md bg-black/30 rounded-3xl
                w-1/3 min-w-[450px] max-w-[550px]
                h-3/4 min-h-[700px] max-h-[1200px]
                content-center
                flex">  
                    <>{login.render()}</>
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
