import React, { Component, useState } from 'react' 
import { IDInputField, PasswordInputField } from '../../components/ModularComponent/LoginPageInputFields';
import axios from 'axios';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGoogle, faSpotify } from '@fortawesome/free-brands-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
library.add(faGoogle, faSpotify);

export default class Login { 

    OnClick = (id:string, pw:string) =>
    {
        axios
            .post(
                "https://echo-backend-production.up.railway.app/user/register",
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

    

    PlatformLogin = (text:string, color:string, icon:any, postUrl:string) =>
    { 
        const OnClick = () =>
        {
            open(postUrl);
        }
        return ( 
            <a className={"bg-" + color + " text-white h-[25px] w-[25px] font-[16px]"} onClick={OnClick}>
                <FontAwesomeIcon icon={icon} />
            </a>
        )
    }

    render() {  

        var [id, SetID] = useState(String); 
        var idInputField = new IDInputField(SetID);

        var [pw, SetPW] = useState(String);
        var pwInputField = new PasswordInputField(SetPW); 

        var center = "flex flex-col items-center justify-center"; 

        return ( 
            <div className={'w-full h-full flex-1 ' + center}>
                <div className={"w-full h-[100px] " + center}> 
                    <p className=' text-4xl text-white font-bold'>LOGIN</p>
                </div>
                <form className='flex w-4/5 h-2/3'>
                    <div className="flex flex-col items-center justify-center w-full h-full">
                        
                        <>{idInputField.Export()}</>
                        <>{pwInputField.Export()}</>

                        <div className="form-check w-full flex flex-row ...">
                            <input className="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-blue-600 checked:border-blue-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer" type="checkbox" value="" id="flexCheckDefault"/>
                            <label className="form-check-label inline-block text-white" htmlFor="flexCheckDefault">Remember My Credentials</label>
                        </div> 
                        <br/>
                        
                        <button className={
                            "top-6 w-full h-[60px] rounded-xl bg-blue-800 font-bold text-xl "} onClick={() => this.OnClick(id, pw)} type="submit">LOGIN</button>
                        <br/><br/> 

                        <span className='text-xl'> Or login with </span> 
                        <div className='px-9 flex flex-row'>
                            {<>{this.PlatformLogin("", "white", faGoogle, "#")}</>}
                            {<>{this.PlatformLogin("", "white", faSpotify, "https://echo-backend-production.up.railway.app/oauth/spotify")}</>}
                        </div>
                    </div>
                </form>
                <div className="text-center w-full p-t-115">
                    <span className=" text-xl">Not a member? </span>
                    <a className=" font-bold text-xl" href="#">Sign Up</a>
                </div>
            </div>
        )
    }
}
