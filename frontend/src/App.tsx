import { useState } from "react";
import SpotifyAuthentication from "./components/Spotify"; 
import Nav from "./components/Nav/Nav"
import LoginPage from "./components/Authentication/LoginPage"; 
import AccountCreationPage from "./components/Authentication/AccountCreationPage";
import InputField from "./components/Core/InputField";


function App() 
{  
  /*
  const PrintValue = () =>
  {
    console.log(inputField.valueHolder.value); 
  }

  var inputField = new InputField("inputField1");
  inputField.valueHolder.InvokeOnChange = PrintValue;
*/

  return ( 
    <div>
      <AccountCreationPage/>
    </div>  
  );
}

export default App;
