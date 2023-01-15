import React, { useState, ChangeEvent, FocusEvent, ReactEventHandler } from 'react';
import { InputField, ValidationFunction, ValidationInputField, ValidationInputFieldOnChange, ValidationStyle } from './ValidationInputFields';

class IDInputField extends InputField
{
    inputField: ValidationInputField;
    
    Condition: ValidationFunction = (inputFieldText):number =>
    {
        return 0;
    }

    constructor(SetRealValue: Function)
    { 
        super(SetRealValue);    

        this.inputField = new ValidationInputFieldOnChange(this, "ID", "Enter ID", "", "", this.Condition, 
            "", "", true); 
 
    } 
}

class PasswordInputField extends InputField
{ 
    inputField: ValidationInputField;  

    //Custom State:
    // 3 - Too Short
    // 4 - Too Long
    // 5 - Too Many Same Characters/Simple
    // 6 - Etc. 
    Condition: ValidationFunction = (inputFieldText):number =>
    {
        return 0;
    }

    constructor(SetRealValue: Function)
    { 
        super(SetRealValue);    

        this.inputField = new ValidationInputFieldOnChange(this, "Password", "Enter Password", "", "", this.Condition, 
            "password", "", true); 
 
    } 
}

export
{
    PasswordInputField,
    IDInputField
}

