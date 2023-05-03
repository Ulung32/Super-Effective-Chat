import axios from "axios"

const API_URL = import.meta.env.API_URL

export type UserRequest = {
    username: string
    password: string
}

const createUser = (data: UserRequest) => axios.post(`http://localhost:5000/stimaGPT/User`, data)


export {createUser}