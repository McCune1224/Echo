import { useState } from "react";

const PooPoo = () => {
    const [show, setShow] = useState(false); 
    const [text, setText] = useState("Boobs");
  
    const handleClick = () => {
      setShow(!show);
      
      setText("Boobs"); 
      if(show) 
        setText("Dick"); 
    } 
   
  
    return <div><button onClick={handleClick}>Button of the day: {text}</button></div>
  };

const SpotifyAuthentication = () => {
 
    const [text, setText] = useState("Spotify Login");

    const handleClick = () => {  
            setText(text + " " + text);  
      } 

    return <div><button onClick={handleClick}>{text}</button></div>
}

export default SpotifyAuthentication; 