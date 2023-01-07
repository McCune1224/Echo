import { useState } from "react";
import axios from "axios";

const SpotifyReponse = () => {

}

const SpotifyAuthentication = () => {  

    const [text, setText] = useState("Spotify Login");

    const handleClick = () => {    
        var endpoint : string = "https://echo-backend-production.up.railway.app/oauth/spotify";
        axios.get(endpoint)
        .then(function (response) {
          // handle success
          return <div><button onClick={handleClick}>{response.status}</button></div>
        })
        .catch(function (error) {
          // handle error
          console.log(error);
        })
        .then(function () {
          // always executed
        });
      } 

    return <div><button onClick={handleClick}>{text}</button></div>
}

export default SpotifyAuthentication; 