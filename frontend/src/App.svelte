<script lang="ts">  
    import css from "./App.css";
    import { SetCookie } from "./lib/API/CookieManager" 
    import env from "./lib/API/Env"
    import PageManager from "./lib/Pages/PageManager.svelte";     

    //This is strictly for checking whether there is a return code, then close the thingy
    const CheckAuthCode = () =>
    {  
        var path = window.location.pathname;
        if(!path.includes("code="))
            return;

        var args = path.split("code=")[1];
        console.log("Args : " + args);

        var code = args.split('&')[0];
        console.log("Code : " + code);
 
        var serviceQuarry = path.split("service=");
        var service = "";

        //This is because alex has not fixed the fucking code yet, please add &service=spotify or &service=google or fukcing something.
        if(serviceQuarry.length == 0)
        {
            service = "spotify";
            alert("Alex, please add &service={service name} so i can actually use it. go to App.svelte line 23 and it's a god damn mess. Currently," +
            " the only service that works is spotify.");
        }
        else
        {
            service = serviceQuarry[1];
            console.log("Service : " + service);
        }
        //Now add to cookie  
        SetCookie(service, code);
    }
</script>

<main>
    <!-- Emerald background for entire page in tailwindcss -->
    <PageManager/>
</main>
