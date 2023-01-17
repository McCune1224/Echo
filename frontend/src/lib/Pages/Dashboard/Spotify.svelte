<script>
    import axios from "axios"
    import { GetCookie } from "../../API/CookieManager"
    import LoginButtonSpotify from "../Auth/Platform/LoginButtonSpotify.svelte"


    var currentToken = GetCookie("spotify");
    var tokenValid = currentToken != null;
    
    var headerToken = () => "Bearer " + currentToken;  

    const config =
    { 
        headers:
        {
        "Authorization": headerToken(),
        "Content-Type": "application/json",
        "Accept": "application/json",
        }
    };

    console.log(config);

    function ValidateSession()
    {    
        axios.get("https://api.spotify.com/v1/me", config)
        .then(response => {
            console.log("Spotify Response: ", response);
        })
        .catch(error => { 
            if(error.response.status == 401) 
                alert(error.response.data.error['message']); 
            console.log("Spotify Error: ", error);
            tokenValid = false;
        })
    }

    ValidateSession();
</script>

<div>
{#if !tokenValid}
    <LoginButtonSpotify/>
{:else} 
    <h1 class=" font-bold text-white text-4xl" > Spotify is Connected! </h1>
{/if} 
</div>
