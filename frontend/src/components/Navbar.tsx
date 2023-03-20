import React from 'react'
import { Link } from 'react-router-dom'

const Navbar = () => {
    //Navbar with content on the left and right side of the page
    return (
        <div className="flex justify-between items-center bg-zinc-900 border-b-4 border-pink-400 text-white p-4">
            <ul className="text-3xl flex items-center flex-auto">
                <img src="https://i.imgur.com/4Q1QY2y.png" alt="logo" className="w-10 h-10" />
                <h1 className="text-7xl font-bold ml-2 mx-9">Echo</h1>
                <li className="hover:bg-zinc-700 rounded-lg text-3xl px-8">
                    Home
                </li>
                <li className="hover:bg-zinc-700 rounded-lg text-3xl px-8">
                    Sync
                </li>
                <li className="hover:bg-zinc-700 rounded-lg text-3xl px-8" >
                    Connected Services
                </li>
            </ul>
            <div className="flex items-center">
                <Link to="/account">
                    <i className="fa-solid fa-user"></i>
                </Link>
                <button onClick={() => console.log("foobarba")}>
                    <h1 className="text-xl font-bold mr-2">Logout</h1>
                    </button>
            </div>
        </div>
    )
}

export default Navbar
