<script lang="ts">  
    import css from "./App.css";
    import { SetCookie } from "./lib/API/CookieManager" 
    import env from "./lib/API/Env"
    import PageManager from "./lib/Pages/PageManager.svelte";     

    //This is strictly for checking whether there is a return code, then close the thingy
    const CheckAuthCode = () =>
    {  
        //https://echo-frontend.up.railway.app//?token_spotify=BQAtWsxWOUjTD25J...
        var path = window.location.pathname;
        if(!path.includes("?token_"))
            return;

        var serviceEqualsToken = path.split("token_")[1];
        console.log("Service=Token : " + serviceEqualsToken);

        var [service, token] = serviceEqualsToken.split('=');   

        //Now add to cookie  
        SetCookie(service, token);
    }
</script>

<main>
    <!-- Emerald background for entire page in tailwindcss -->
    <PageManager/>
</main>
