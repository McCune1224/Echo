import axios from 'axios'
import React, { useEffect } from 'react'
import { useCookies } from 'react-cookie'
import { useNavigate } from 'react-router-dom'
import { createBackendAPIClient } from '../api/BackendAPI'
import { createSpotifyAPI } from '../api/SpotifyAPI'
import Navbar from './Navbar'

interface User {
    username: string
    id: number
    email: string
}

const Dashboard = () => {

    const [authToken, _] = useCookies(['auth_token'])


    const [user, setUser] = React.useState<User>()
    const backendClient = createBackendAPIClient(authToken.auth_token)
    useEffect(() => {
        const fetchUser = async () => {
            const response = await backendClient.get('/users')
            console.log(response)
            setUser(response.data)
        }

        fetchUser()
    }, [])

    const spotifyAuth = async () => {
        const response = await backendClient.get('oauth/spotify')
    }



    return (
        <div>
            <Navbar />
            <div className="flex justify-center items-center h-screen">
                <div className="container flex flex-col justify-center items-center">
                    <h2 className="text-7xl font-bold">Welcome {user?.username}</h2>
                    <button onClick={() => spotifyAuth} className="bg-pink-400 text-white text-3xl font-bold px-4 py-2 rounded-lg">
                        <i className="fa-brands fa-spotify text-7xl"></i>
                        Connect Spotify</button>
                </div>
            </div>
        </div>
    )
}

export default Dashboard
