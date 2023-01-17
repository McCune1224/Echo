
<script>  
    import env from '../../../API/Env' 
    import { faSpotify } from '@fortawesome/free-brands-svg-icons';
    import PlatformLoginButton from '../../../Elements/PlatformLoginButton.svelte'
</script>    

<PlatformLoginButton text="" color="text-green-500" icon={faSpotify} postUrl={env.VITE_SPOTIFY_OAUTH_API_URL}/>