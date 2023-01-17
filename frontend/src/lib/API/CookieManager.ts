function GetCookie(key: string)
{
    var cookie = document.cookie; 
    if(!cookie.includes(key + '=')) return null;

    var value = '; ' + cookie;
    const parts = value.split('; ' + key + '=');
    if(parts.length === 2) return parts.pop().split(';').shift();
}   

function SetCookie(key: string, value: string, expires: number = null)
{
    document.cookie = key + '=' + value + '; path=/;';
}

export
{
    GetCookie,
    SetCookie
}; 