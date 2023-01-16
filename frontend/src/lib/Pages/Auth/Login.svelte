<script lang="ts">

    import InputField from "../../Elements/InputField.svelte";
    import PlatformLoginButton from "../../Elements/PlatformLoginButton.svelte"
    import axios from "axios";
    
    var center = "flex flex-col items-center justify-center";
    
    var id = "a";
    var pw = "a";   

    export var OnSubmit;

    const ChangeID = (value: string) =>
    {
        id = value; 
    }
    const ChangePW = (value: string) =>
    {
        pw = value; 
    }

    const OnClick = (event:Event & {currentTarget: EventTarget & HTMLFormElement;}) =>
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
            .then((response) =>
            { 
                console.log(response); 
                
                var token = response.data["token"]; 

                var date = new Date();  
                var expireDate = new Date(date.setDate(date.getDate() + 3));

                console.log(expireDate);  
                document.cookie = "login=" + token + "; expires=" + expireDate.toUTCString() + " path=/;";  

                OnSubmit();
            })
            .catch((error) =>{
                console.log(error)
            })
    }


</script> 

<div class={'text-white w-full h-full flex-1 ' + center}>
    <div class={"w-full h-[100px] " + center}> 
        <p class=' text-4xl text-white font-bold'>LOGIN</p>
    </div>
    <form class='flex w-4/5 h-2/3' on:submit={(e) => OnClick(e)}>
        <div class="flex flex-col items-center justify-center w-full h-full">
            
            <InputField label="ID" placeholder="Enter ID" requiredField={true} type="username" callback={ChangeID} />
            <InputField label="Password" placeholder="Enter Password" requiredField={true} type="password" callback={ChangePW} />

            <div class="form-check w-full flex flex-row ...">
                <input class="form-check-input appearance-none h-4 w-4 border border-gray-300 rounded-sm bg-white checked:bg-blue-600 checked:border-blue-600 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-2 cursor-pointer" type="checkbox" value="" id="flexCheckDefault"/>
                <label class="form-check-label inline-block text-white" for="flexCheckDefault">Remember My Credentials</label>
            </div> 
            <br/>
            
            <button class={
                "w-full h-[60px] rounded-xl bg-blue-800 font-bold text-xl text-white"} type="submit">LOGIN</button>
            <br/>

            <span class='text-xl'> Or login with </span> 
            <div class='flex flex-row py-10'>
                <PlatformLoginButton/>
                <PlatformLoginButton/>
                
            </div>
            <br class=' py-10'/>
        </div>
    </form>
    <div class=" py-10 text-center w-full p-t-115 ">
        <span class=" text-xl">Not a member? </span>
        <a class=" font-bold text-xl" href="#">Sign Up</a>
    </div>
</div>