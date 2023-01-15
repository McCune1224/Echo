import React, { useState } from "react"; 
import * as InputFields from "../ModularComponent/ValidationInputFields";
import axios from "axios";

const AccountCreationPage = () => {
    //This should be hard defined
    var maximumDisableInquires = 10;

    var [disableInquiries, SetDisableInquiries] = React.useState(
        Array<boolean>(maximumDisableInquires)
    );
    var [disableSubmit, SetDisableSubmit] = React.useState(false);

    //Called Every Time the inquire has a change
    const UpdateInquiries = () => {
        var disable = false;
        disableInquiries.forEach((state) => {
            if (state == true) {
                disable = true;
                return;
            }
        });
        SetDisableSubmit(disable);
    };

    const AssignSubmitAffecter = (inputField: InputFields.InputField) =>
        inputField.AffectSubmitButton(Callback, disableInquiries.length - 1);

    //Important Callback Function
    const Callback = (state: boolean, id: number) => {
        var tempInquiries = disableInquiries;
        tempInquiries[id] = state;

        SetDisableInquiries(tempInquiries);
        UpdateInquiries();
    };

    //Variable Field
    var [emailInput, SetEmailInput] = useState(String); //brianmjk@gmail.com
    var emailInputField = new InputFields.EmailInputField(SetEmailInput);

    var [passwordInput, SetPasswordInput] = useState(String);
    var passwordInputField = new InputFields.PasswordInputField(
        SetPasswordInput
    );
    AssignSubmitAffecter(passwordInputField);

    var [usernameInput, SetUsernameInput] = useState(String);
    var usernameInputField = new InputFields.UsernameInputField(
        SetUsernameInput
    );
    AssignSubmitAffecter(usernameInputField);

    //Post Variables
    const [data, setData] = useState("");
    const [query, setQuery] = useState("");

    const OnSubmit = (event: React.FormEvent<HTMLFormElement>) => 
    {
        console.log("SUBMITTED!");

        event.preventDefault(); 
        axios
            .post(
                "https://echo-backend-production.up.railway.app/user/register",
                {
                    email: emailInput,
                    username: usernameInput,
                    password: passwordInput,
                }
            )
            .catch(function (error) {
                console.log(error);
                if (error.data.response.message.contain("User"))
                    emailInputField.inputField.SetElementState(3);
            });
    };

    const submitButton = () => {
        if (disableSubmit)
            return (
                <button
                    disabled={true}
                    className="bg-gray-500 hover:bg-gray-700 text-white font-bold py-2 px-5 rounded"
                >
                    Submit
                </button>
            );
        return (
            <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-5 rounded">
                Submit
            </button>
        );
    };

    return (
        <div>
            <form onSubmit={(event) => OnSubmit(event)}>
                <>{emailInputField.Export()}</>
                <>{usernameInputField.Export()}</>
                <>{passwordInputField.Export()}</>
                <div className="button-container">
                    <>{submitButton()}</>
                </div>
            </form>
        </div>
    );
};

export default AccountCreationPage;
