import { useState } from "react"; 
import SpotifyAuthentication from "./components/Spotify"; 
import Nav from "./components/Nav/Nav"
import LoginPage from "./components/Authentication/LoginPage"; 
import AccountCreationPage from "./components/Authentication/AccountCreationPage";
import InputField from "./components/Core/InputField"; 
import Auth from "./pages/auth/Auth";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";


function App() 
{  
  /*
  const PrintValue = () =>
  {
    console.log(inputField.valueHolder.value); 
  }

<div>
<Auth/>
</div> 

  var inputField = new InputField("inputField1");
  inputField.valueHolder.InvokeOnChange = PrintValue;
*/
return <Auth/>;
/*
  return (  
    <Router>
      <Routes>
        <Route path="/auth" element={<Auth/>}/> 
      </Routes>
    </Router>
  );*/
}

export default App;
