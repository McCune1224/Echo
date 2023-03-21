import React, { useEffect } from 'react'
import axios, { AxiosError } from 'axios'

interface SignUpFields {
    username: string
    email: string
    password: string
}

const SignUp = () => {


    const [submitError, setSubmitError] = React.useState<string>('')
    const [username, setUsername] = React.useState('')
    const [email, setEmail] = React.useState('')
    const [password, setPassword] = React.useState('')

    const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setUsername(event.target.value)

    }
    const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setEmail(event.target.value)
    }

    const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(event.target.value)

    }

    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()
        const signUpFields: SignUpFields = {
            username: username,
            email: email,
            password: password
        }
        try {


            const response = await axios.post('http://localhost:42069/user/register', signUpFields)
            if (response.status === 400) {
                console.log(submitError)
            }
        }
        catch (error) {
            console.log(error)

        }
    }


    return (
        <div>
            <form onSubmit={(e) => handleFormSubmit(e)}
                className=" bg-slate-200 container text-black">
                <h2 className="text-7xl py-9">
                    Sign Up to get started
                </h2>
                <label className="text-5xl">Username
                    <input type="text" name="username" onChange={(e) => handleUsernameChange(e)}></input>
                </label>
                <label className="text-5xl">Email
                    <input type="text" name="email" onChange={(e) => handleEmailChange(e)}></input>
                </label>
                <label className="text-5xl">Password
                    <input type="password" name="password" onChange={(e) => handlePasswordChange(e)}></input>
                </label>
                <input type="submit"
                    className={`bg-pink-400 text-7xl font-bold py-2 px-4 
                          border-2 border-white hover:bg-pink-500 `}
                >
                </input>
                {submitError && <p className="text-red-500">{submitError}</p>}
            </form>

        </div>
    )
}

export default SignUp
