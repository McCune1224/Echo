
<script>  
    import env from '../../../API/Env' 
    import { faGoogle } from '@fortawesome/free-brands-svg-icons';
    import PlatformLoginButton from '../../../Elements/PlatformLoginButton.svelte'
</script>    

<PlatformLoginButton text="" color="text-blue-500" icon={faGoogle} postUrl={env.VITE_SPOTIFY_OAUTH_API_URL}/>