import axios from 'axios'



export const createBackendAPIClient = (auth_token: string) => {
    return axios.create({
        baseURL: import.meta.env.VITE_API_URL,
        headers: {
            Authorization: `Bearer ${auth_token}`
        }
    })
}

