import React, { useState, ChangeEvent, FocusEvent, ReactEventHandler } from 'react' 
import Bindable from '../Core/Bindable/Bindable';

class ValidationStyle
{
    // Base Layout of how this should be made
    // 0 - default
    // 1 - success
    // 2 - error
    // 3 - from here is all custom stylings.
    styles: Array<string>;

    //Static Values
    public static readonly BaseLabel = "block mb-2 text-4lx font-medium text-gray-700 dark:text-white-500";
    public static readonly SuccessLabel = "block mb-2 text-4lx font-medium text-green-700 dark:text-green-500";
    public static readonly FailLabel = "block mb-2 text-4lx font-medium text-red-700 dark:text-red-500";

    public static DefaultLabel = () => new ValidationStyle(ValidationStyle.BaseLabel, ValidationStyle.SuccessLabel, ValidationStyle.FailLabel);
    
    public static readonly BaseInput = 
    "bg-gray-50 p-4 border border-gray-500 text-gray-900 text-3lx rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500";
    public static readonly SuccessInput = 
    "bg-green-50 p-4 border border-green-500 text-green-900 dark:text-green-400 placeholder-green-700 dark:placeholder-green-500 text-3lx rounded-lg focus:ring-green-500 focus:border-green-500 block w-full dark:bg-gray-700 dark:border-green-500";
    public static readonly FailInput = 
    "bg-red-50 p-4 border border-red-500 text-red-900 placeholder-red-700 text-3lx rounded-lg focus:ring-red-500 dark:bg-gray-700 focus:border-red-500 block w-full dark:text-red-500 dark:placeholder-red-500 dark:border-red-500";

    public static DefaultInput = () => new ValidationStyle(ValidationStyle.BaseInput, ValidationStyle.SuccessInput, ValidationStyle.FailInput);

    public static readonly BaseMessage = "mt-2 text-sm text-600 dark:text-500";
    public static readonly SuccessMessage = "mt-2 text-sm text-green-600 dark:text-green-500";
    public static readonly FailMessage = "mt-2 text-sm text-red-600 dark:text-red-500";

    public static DefaultMessage = () => new ValidationStyle(ValidationStyle.BaseMessage, ValidationStyle.SuccessMessage, ValidationStyle.FailMessage);
 
    constructor(...args: string[])  
    {  
        this.styles = [];
        for(var i=0; i<args.length; i++)
        {
            this.styles.push(args[i]);
        } 
    }
 

    Add = (style: string) => this.styles.push(style); 

    AddArray = (styles: Array<string>) =>
    {
        styles.forEach(style => {
            this.Add(style);
        }); 
    }

    //Default Parameters
    Base = () => this.styles[0];
    Success = () => this.styles[1];
    Fail = () => this.styles[2]; 

    //Getters and Getting Methods 
    Output = (success:number): string => 
    {
        return this.styles[success];
    }
}   

class ValidationInputField
{ 
    //Validations
    labelValidation: ValidationStyle;
    inputValidation: ValidationStyle;
    messageValidation: ValidationStyle;
    customMessages: ValidationStyle;
 
    //Use States
    labelStyle: Bindable<string>; 
    //labelStyle: string; setLabelStyle: Function;
    inputStyle: Bindable<string>;
    //inputStyle: string; setInputStyle: Function;
    messageStyle: Bindable<string>;
    //messageStyle: string; setMessageStyle: Function; 
    customMessage: Bindable<string>;
    //customMessage: string; setCustomMessage: Function;

    //Static One Set Values
    inputName: string;  
    htmlType: string;
    placeholder: string;
    successMessage: string;
    failMessage:string; 
    condition: Function;  
    required: boolean;

    //* Input Field Parent Object Functions 
    //Parent Field
    inputField: InputField;
    _value: Bindable<string>;

    AddCustomValidations = (
        additionalLabelValidation: Array<string> = [],
        additionalInputValidation: Array<string> = [],
        additionalMessageValidation: Array<string> = [],
        additionalCustomMessages: Array<string> = []
        ) =>
    {  
        for(let i = 0; i < additionalLabelValidation.length; i++) 
        {
            this.labelValidation.Add(additionalLabelValidation[i]);
            this.inputValidation.Add(additionalInputValidation[i]);
            this.messageValidation.Add(additionalMessageValidation[i]); 
            this.customMessages.Add(additionalCustomMessages[i]); 
        }        
    }

    GetValue = () => this._value.value;

    SetValue = (value:string) => 
    {
        this._value.ChangeValue(value);
        this.inputField.SetRealValue(value);
    }

    constructor(inputField: InputField, inputName: string, placeholder: string, successMessage: string, failMessage:string, condition: Function, 
        type: string = "text",
        customMessageDefault: string = "",
        required: boolean = false
        )
    { 
        this._value = new Bindable<string>("");

        //Parent Function
        this.inputField = inputField;

        this.inputName = inputName;  
        this.placeholder = placeholder;
        this.successMessage = successMessage;
        this.failMessage = failMessage;
        this.condition = condition;
        this.htmlType = type; 
        this.required = required;

        //They all must have the same length 
        this.labelValidation = ValidationStyle.DefaultLabel();
        this.inputValidation = ValidationStyle.DefaultInput();
        this.messageValidation = ValidationStyle.DefaultMessage();

        this.customMessages = new ValidationStyle(customMessageDefault, successMessage, failMessage);

        this.labelStyle = new Bindable<string>(this.labelValidation.Base());
        this.inputStyle = new Bindable<string>(this.inputValidation.Base()); 
        this.messageStyle = new Bindable<string>(this.messageValidation.Base()); 
        this.customMessage = new Bindable<string>(this.customMessages.Base());  
    } 
    
    SetElementState = (value: number) =>
    { 
        this.labelStyle.ChangeValue(this.labelValidation.Output(value));
        this.inputStyle.ChangeValue(this.inputValidation.Output(value));
        this.messageStyle.ChangeValue(this.messageValidation.Output(value));
        this.customMessage.ChangeValue(this.customMessages.Output(value));

        if(this.inputField.disableSubmit) 
            this.inputField.SetButtonActive(value != 1); 
    }

    UpdateValues = (inputString: string) =>
    { 
        var state:number = this.condition(inputString); 
 
        //This is temporary 
        this.SetValue(inputString); 

        //Apply
        this.SetElementState(state);
    }

    OnChange = (event: ChangeEvent<HTMLInputElement>) =>
    { 
        this.SetValue(event.target.value); 
    }

    OnSubmit = (event: React.FormEvent<HTMLInputElement>) =>
    {
        this.UpdateValues(event.currentTarget.value);
    }

    Export = () => {};
}

class ValidationInputFieldOnChange extends ValidationInputField 
{  
    OnChange = (event: ChangeEvent<HTMLInputElement>) =>
    { 
        this.SetValue(event.target.value); 
        this.UpdateValues(event.target.value); 
    }

    Export = () =>
    {
        return ( 
            <div className="mb-6 w-full">
                <label /*htmlFor={htmlFor}*/ className={this.labelStyle.value}>{this.inputName}</label>
                {this.required ? 
                (
                    <input onChange={(event) => this.OnChange(event)} type={this.htmlType} /*id={id}*/ className={this.inputStyle.value} placeholder={this.placeholder} required/>
                ) : 
                (
                    <input onChange={(event) => this.OnChange(event)} type={this.htmlType} /*id={id}*/ className={this.inputStyle.value} placeholder={this.placeholder}/>
                ) 
                }
                <p className={this.messageStyle.value}><span className="font-medium">{this.customMessage.value}</span></p>
            </div>  
        )
    }
}

class ValidationInputFieldOnSubmit extends ValidationInputField 
{ 
    Export = () =>
    {
        return ( 
            <div className="mb-6 w-full">
                <label /*htmlFor={htmlFor}*/ className={this.labelStyle.value}>{this.inputName}</label>
                <input onSubmit={(event) => this.OnSubmit(event)} onChange={(event) => this.OnChange(event)} type={this.htmlType} /*id={id}*/ className={this.inputStyle.value} placeholder={this.placeholder}/>
                <p className={this.messageStyle.value}><span className="font-medium">{this.customMessage.value}</span></p>
            </div>  
        )
    }
} 

/// return value should be
// 0 - default
// 1 - success
// 2 - failed
// 3 - from here is custom. Make sure to have the appropriate additional properties. If not it breaks.

type ValidationFunction = (inputFieldText: string) => number;

class InputField
{     
    disableSubmit: boolean;
    submitButtonActive: boolean; SetSubmitButtonActive: Function; //This should only be affect when ```disableSubmit``` is true.
    buttonStatusCallback!: Function; callbackID!: number;
    inputField!: ValidationInputField; 

    //This is modular, make sure to give them it.
    SetRealValue: Function;

    AffectSubmitButton = (callback:Function, id:number) => 
    {
        this.disableSubmit = true;
        this.buttonStatusCallback = callback;
        this.callbackID = id;
    }

    SetButtonActive = (state: boolean) =>
    {
        this.SetSubmitButtonActive(state);
        this.buttonStatusCallback(state, this.callbackID);
    }

    GetValue = () => this.inputField.GetValue();

    constructor(SetRealValue:Function)
    { 
        this.SetRealValue = SetRealValue;
        this.disableSubmit = false; 
        [this.submitButtonActive, this.SetSubmitButtonActive] = useState(true); 
    }
    
    Export = () => this.inputField.Export();
}

class UsernameInputField extends InputField
{
    //Condition:
    // Only Alphabets and Numbers are allowed. 
    // Max Length = 16
    // 2 - Fail, can only contain letters and numbers.
    // 3 - It must be under 16 characters. Too long!

    Condition: ValidationFunction = (inputFieldText):number =>
    {
        if(inputFieldText.length == 0) return 0;
        var alphaNumeralCheck = !Boolean(inputFieldText.match(/^[A-Za-z0-9]*$/));
        var lengthCheck = inputFieldText.length > 16;
        if(alphaNumeralCheck)
            return 2;
        else if(lengthCheck)
            return 3; 
        return 1;
    }
    
    constructor(SetRealValue: Function)
    { 
        super(SetRealValue);   

        var customLabelStyles = [ValidationStyle.FailLabel];
        var customInputStyles = [ValidationStyle.FailInput];
        var customMessageStyles = [ValidationStyle.FailMessage];
        var customMessages = ["The given username is too long. It must be less than equal to 16 characters."] 

        this.inputField = new ValidationInputFieldOnChange(this, "Username", "Enter a username", "✓", "Special characters are not allowed.", this.Condition, 
            "username", "Enter a Username. [a~z], [A~Z], [0~9]", true); 

        this.inputField.AddCustomValidations(customLabelStyles, customInputStyles, customMessageStyles, customMessages); 
    } 
}

//Validation Objects  
class EmailInputField extends InputField
{  

    //Condition:
    // 0 - Empty
    // 1 - Checks for @
    // 2 - Success
    // 3 - Already Exists - Custom
    Condition: ValidationFunction = (inputFieldText):number =>
    {
        if(inputFieldText.length == 0) return 0;
        if(inputFieldText.includes('@')) return 1;
        return 2;
    }

    constructor(SetRealValue: Function)
    { 
        super(SetRealValue);  

        var customLabelStyles = [ValidationStyle.FailLabel];
        var customInputStyles = [ValidationStyle.FailInput];
        var customMessageStyles = [ValidationStyle.FailMessage];
        var customMessages = ["The given email is already registered."]  

        this.inputField = new ValidationInputFieldOnSubmit(this, "Email Address", "example@example.com", "Good", "Bad", this.Condition, 
            "email", "Enter Email", true); 
 
        this.inputField.AddCustomValidations(customLabelStyles, customInputStyles, customMessageStyles, customMessages); 
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
        if(inputFieldText.length == 0) return 0;
        var length: number = inputFieldText.length;
        if(length < 8) return 3;
        else if(length > 16) return 4;

        return 1;
    }

    constructor(SetRealValue: Function)
    { 
        super(SetRealValue);   

        var customLabelStyles = [ValidationStyle.FailLabel, ValidationStyle.FailLabel, ValidationStyle.FailLabel];
        var customInputStyles = [ValidationStyle.FailInput, ValidationStyle.FailInput, ValidationStyle.FailInput];
        var customMessageStyles = [ValidationStyle.FailMessage, ValidationStyle.FailMessage, ValidationStyle.FailMessage];
        var customMessages = ["The given password is too short.", "The given password is too long.", "The given password is too simple."] 

        this.inputField = new ValidationInputFieldOnChange(this, "Password", "Enter Password", "✓", "Unacceptable", this.Condition, 
            "password", "Enter a password. [a~z], [A~Z], [0~9]... Special Characters", true); 

        this.inputField.AddCustomValidations(customLabelStyles, customInputStyles, customMessageStyles, customMessages); 
    } 
}
export {
    InputField,
    EmailInputField,
    PasswordInputField,
    UsernameInputField,
    ValidationStyle,
    ValidationInputField,
    ValidationInputFieldOnChange,
    ValidationInputFieldOnSubmit
};
export type { ValidationFunction };
