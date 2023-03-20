import React, { useEffect } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'
import { useCookies } from 'react-cookie'
import SignUp from './SignUp'

const Home = () => {





    return (
        <div className="flex justify-center items-center h-screen">
            <div className="container flex flex-col justify-center items-center text-center">
                <p className="text-7xl py-9">
                    <b className="text-9xl text-pink-400">Echo</b> is a
                    simple and fast website to transfer and sync playlists
                    from one streaming service to another.
                </p>
                <button className="bg-pink-400 text-7xl font-bold py-2 px-4 border-2 border-white hover:bg-pink-500 ">
                    <Link to="/signup">Sign-Up</Link>
                </button>
            </div>
        </div>
    )
}

export default Home
