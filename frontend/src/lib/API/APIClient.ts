import axios from 'axios'; 

const CreateAPIClient = (token) =>
{ 
    return axios.create({
        baseURL: 'https://echo-backend-production.up.railway.app/',
        headers: {'Authorization': 'Bearer ' + token}
    });
} 

export default CreateAPIClient;