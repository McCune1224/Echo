import { useState } from "react";    
import Auth from "./pages/auth/Auth"; 
import { useCookies } from 'react-cookie';
import Dashboard from "./pages/Dashboard/Dashboard";


function App() 
{   
  var authProviders = ["spotify", "google"];
  for (var i = 0; i < authProviders.length; i++)
  {
    var key = authProviders[i];
    const [value] = useCookies([key]); 
    var val = value[key];
    if(val != undefined)  
      return <Dashboard/>;  
  }

  return <Auth/>; 
}

export default App;
