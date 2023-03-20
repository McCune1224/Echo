import axios from 'axios';
import SpotifyWebAPI from 'spotify-web-api-js';


export const createSpotifyAPI = (accessToken: string) => {
    const spotifyAPI = new SpotifyWebAPI();
    spotifyAPI.setAccessToken(accessToken);
    return spotifyAPI;
}
