import axios, { AxiosError } from 'axios'
import React from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import { useCookies } from 'react-cookie'

interface LoginFields {
    username?: string
    email?: string
    password: string
}

const Login = () => {

    const navigate = useNavigate()
    const location = useLocation()

    const [auth_token, setAuth_token] = useCookies(['auth_token'])
    const [username, setUsername] = React.useState('')
    const [email, setEmail] = React.useState('')
    const [password, setPassword] = React.useState('')
    const [submitError, setSubmitError] = React.useState<string>('')

    const handleFormSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault()

        const loginFields: LoginFields = {
            username: username,
            email: email,
            password: password
        }

        try {
            const baseURL = import.meta.env.VITE_API_URL
            const response = await axios.post(baseURL + "/users/login", loginFields)
            if (response.status === 200 && response.data?.token === "") {
                setSubmitError(response.data?.message)
                console.log(submitError)
                return
            }
            setAuth_token('auth_token', response.data?.token,
                {
                    path: '/',
                    expires: new Date(response.data.exp * 1000)
                }
            )
            navigate(location.state?.from || '/dashboard')


        }

        catch (error) {
            console.log(error)
        }
    }
    return (
        <div>
            <form className="bg-slate-200 container text-black" onSubmit={(e) => handleFormSubmit(e)}>
                <label>
                    Username:
                    <input type="text" name="username" onChange={(e) => setUsername(e.target.value)} />
                </label>
                <label>
                    Password
                    <input type="password" name="password" onChange={(e) => setPassword(e.target.value)} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </div>
    )
}

export default Login
