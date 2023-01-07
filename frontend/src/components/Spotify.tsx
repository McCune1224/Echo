import { useState } from "react";

const SpotifyAuthentication = () => {
 
    const [text, setText] = useState("Spotify Login");

    const handleClick = () => {  
            setText(text + " " + text);  
      } 

    return <div><button onClick={handleClick}>{text}</button></div>
}

export default SpotifyAuthentication; 