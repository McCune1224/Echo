import React, { Component, useState } from 'react' 
import { IDInputField, PasswordInputField } from '../../components/ModularComponent/LoginPageInputFields';
import axios from 'axios';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGoogle, faSpotify } from '@fortawesome/free-brands-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
library.add(faGoogle, faSpotify);

const Login = () => 
{     
    const OnClick = (event:any, id:string, pw:string) =>
    {
        event.preventDefault();

        axios
            .post(
                "https://echo-backend-production.up.railway.app/user/login",
                {
                    username: id, 
                    password: pw
                }
            )
            .catch(function (error) {
                /*
                console.log(error);
                if (error.data.response.message.contain("User"))
                    emailInputField.inputField.SetElementState(3);*/
            });
    }

    

    const PlatformLogin = (text:string, color:string, icon:any, postUrl:string) =>
    { 
        const OnClick = () =>
        {
            window.open(postUrl, "_self");
        }
        return ( 
            <div className='px-2'>
                <a className={"h-[65px] rounded-2xl w-[65px] font-[16px]  flex flex-col items-center justify-center " + color} onClick={OnClick}>
                    <FontAwesomeIcon className={'w-3/4 h-3/4 text-white'} icon={icon} />
                </a>
            </div>
        )
    }

    var [id, SetID] = useState(String); 
    var idInputField = new IDInputField(SetID);
    
    var [pw, SetPW] = useState(String);
    var pwInputField = new PasswordInputField(SetPW); 

    var center = "flex flex-col items-center justify-center";  

    return ( 
        <div className={'text-white w-full h-full flex-1 ' + center}>
            <div className={"w-full h-[100px] " + center}> 
                <p className=' text-4xl text-white font-bold'>LOGIN</p>
            </div>
            <form className='flex w-4/5 h-2/3' onSubmit={(e) => OnClick(e, id, pw)}>
                <div className="flex flex-col items-center justify-center w-full h-full">
                    
                    <>{idInputField.Export()}</>
                    <>{pwInputField.Export()}</>

                    <div className="form-check w-full flex flex-row ...">
                        <input className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-blue-600 checked:border-blue-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer" type="checkbox" value="" id="flexCheckDefault"/>
                        <label className="form-check-label inline-block text-white" htmlFor="flexCheckDefault">Remember My Credentials</label>
                    </div> 
                    <br/>
                    
                    <button className={
                        "w-full h-[60px] rounded-xl bg-blue-800 font-bold text-xl text-white"} type="submit">LOGIN</button>
                    <br/>

                    <span className='text-xl'> Or login with </span> 
                    <div className='flex flex-row py-10'>
                        {<>{PlatformLogin("", "bg-blue-600", faGoogle, "#")}</>}
                        {<>{PlatformLogin("", "bg-green-600", faSpotify, "https://echo-backend-production.up.railway.app/oauth/spotify")}</>}
                    </div>
                    <br className=' py-10'/>
                </div>
            </form>
            <div className=" py-10 text-center w-full p-t-115 ">
                <span className=" text-xl">Not a member? </span>
                <a className=" font-bold text-xl" href="#">Sign Up</a>
            </div>
        </div>
    )
}

export default Login;