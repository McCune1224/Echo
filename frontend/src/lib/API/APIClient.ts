import axios from 'axios'
import env from './Env'

const CreateAPIClient = (token) => {
    return axios.create({
        baseURL: env.VITE_BACKEND_LOGIN_URL,
        headers: { Authorization: 'Bearer ' + token },
    })
}

export default CreateAPIClient
