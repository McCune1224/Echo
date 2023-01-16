<script lang='ts'> 
    import CreateAPIClient from "../API/APIClient";
    import Login from "./Auth/Login.svelte"
    import Dashboard from "./Dashboard/Dashboard.svelte"
    
    var platforms = "login";
    var loginToken = "";
    var tokenValid = false;

    function ValidateCurrentSession()
    {
        tokenValid = true;

        //Display Attributes
        shrinkLoginField = true;
    }

    const GetCookie = (key: string) =>
    {
        var cookie = document.cookie; 
        if(!cookie.includes(key + '=')) return null;

        var value = '; ' + cookie;
        const parts = value.split('; ' + key + '=');
        if(parts.length === 2) return parts.pop().split(';').shift();
    } 
    
    function RefreshValidity()
    { 
        loginToken = GetCookie(platforms);  

        if(loginToken != null)  
        { 
            const client = CreateAPIClient(loginToken);

            client.get('/user')
            .then(response => 
            {
                console.log(response)
                if(response.status == 200)
                    ValidateCurrentSession();
            })
            .catch(error => { console.log("poopoo, you don't exist. You're fake."); })  
        }  
    }


    var expand = false; 
    function SetExpand(state: boolean)
    {
        expand = state;
    }
        
    var displayLoginField = true;
    var shrinkLoginField = false;

    var shrinkDashboard = true;

    function EnableDashboard()
    {
        shrinkDashboard = false;
        console.log("shrink dashboard : ", shrinkDashboard);
    }

    function DisableLoginField()
    {
        displayLoginField = false;
        console.log(displayLoginField);
        //Chain
        expand = true;
        EnableDashboard();
    }

    function ContainerTransitionEnd()
    {
        
    }

    //Login -> shrink the page.
 
    RefreshValidity();

</script>

<div class="bg-cover bg-authBackground h-screen grid place-items-center">  
    <div class={(expand ? "main-container-expanded" : "main-container") + " transition-slow"} on:transitionend={ContainerTransitionEnd}>
        {#if displayLoginField}

            <div class={(shrinkLoginField ? "shrinkable-close" : "shrinkable-open") + " transition-normal w-full h-full"} 
            on:transitionend={DisableLoginField}>
            <Login OnSubmit={RefreshValidity}/></div>
        {:else}
            <div class={(shrinkDashboard ? "shrinkable-close" : "shrinkable-open") + " transition-slow w-full h-full"}>
            <Dashboard/>
            </div>
        {/if}
 
    </div> 
</div>