import React from 'react';

class UserInformation
{
    emailAddress: String;
    username: String;
    password: String;

    constructor(email: String, username: String, password: String) {
        this.emailAddress = email;
        this.username = username;
        this.password = password;
    } 
}

export default UserInformation;