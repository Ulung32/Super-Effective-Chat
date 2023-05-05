import axios from "axios"

const API_URL = import.meta.env.API_URL

export type UserRequest = {
    username: string
    password: string
}

const createUser = (data: UserRequest) => axios.post(`https://secchatbot-production.up.railway.app/stimaGPT/User`, data)


export {createUser}