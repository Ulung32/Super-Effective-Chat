import axios from "axios"

const API_URL = import.meta.env.API_URL

export type User = {
    id: string,
    username: string,
    password: string
}

export type UserRequest = {
    username: string
    password: string
}

const createUser = (data: UserRequest) => axios.post(`http://localhost:1323/stimaGPT/User`, data)


export {createUser}