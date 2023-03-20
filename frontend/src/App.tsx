import { Routes, Route, Navigate } from "react-router-dom"
import { useState } from 'react'
import Home from "./components/Home"
import SignUp from "./components/SignUp"
import Login from "./components/Login"
import Dashboard from "./components/Dashboard"
import { useCookies } from "react-cookie"

function App() {

    //redirect to login if not logged in
    function RequireAuth({ children, redirectTo }: any) {
        const [authToken, _] = useCookies(['auth_token'])
        return authToken['auth_token'] ? children : <Navigate to={redirectTo} />
    }

    //Redirect to dashbaord if logged in
    function AuthRedirect({ children, redirectTo }: any) {
        const [authToken, _] = useCookies(['auth_token'])
        return authToken['auth_token'] ? <Navigate to={redirectTo} /> : children
    }


    return (
        <>
            <Routes>
                <Route path="*" element={<h1>Ruh Roh Raggy 404</h1>} />
                <Route path="/" element={
                    <AuthRedirect redirectTo="/dashboard"> <Home /> </AuthRedirect>
                } />
                <Route path="/signup" element={
                    <AuthRedirect redirectTo="/dashboard"> <SignUp /> </AuthRedirect>
                } />
                <Route path="/login" element={
                    <AuthRedirect redirectTo="/dashboard"> <Login /> </AuthRedirect>
                } />
                <Route path="/dashboard" element={
                    <RequireAuth redirectTo="/login"> <Dashboard /> </RequireAuth>} />
            </Routes>
        </>)

}

export default App
