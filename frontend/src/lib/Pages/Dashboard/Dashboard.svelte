<script lang="ts">
    import axios from 'axios'
    import CreateAPIClient from '../../API/APIClient'
    import { GetCookie } from '../../API/CookieManager'
    import { onMount } from 'svelte'

    interface user {
        id: string
        username: string
        email: string
    }

    let user: user = {
        id: '',
        username: '',
        email: '',
    }
    let playlists: any = []

    const getUserInfo = async () => {
        const token = GetCookie('login')
        const APIClient = CreateAPIClient(token)
        const response = await APIClient.get('/user')
        user = response.data
    }

    //arrow function to fetch user playlists from /playlist endpoint using APIClient
    const getUserPlaylists = async () => {
        const token = GetCookie('login')
        const APIClient = CreateAPIClient(token)
        const response = await APIClient.get('/playlist')
        playlists = response.data
    }
    onMount(async () => {
        await getUserInfo()
    })
</script>

<!-- Make a simple html page that will display the logged in user's info, styling with tailwindcss-->
<div class="flex flex-col items-center justify-center h-screen">
    <h1 class="text-4xl font-bold">Welcome, {user.username}!</h1>
    <p class="text-2xl">Your email is {user.email}</p>
    <p>And your id is {user.id}</p>

    <!-- Button to fetch user playlists using getUserPlaylists -->
    <button on:click={getUserPlaylists}>Get Playlists</button>

    <!-- Render playlists via list after invoking getUserPlaylist button -->
    <ul>
        {#each playlists as playlist}
            <li>{playlist.name}</li>
        {/each}
    </ul>
</div>

