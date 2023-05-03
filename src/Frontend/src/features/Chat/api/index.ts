import axios from "axios";

export type ChatRequest = {
    idHistory: string 
    message: string
    method: string
}

export type HistoryRequest = {
    userId: string,
    name: string
}

export const getHistoryByIdUser = (id: string) => axios.get(`http://localhost:5000/stimaGPT/history?userID=${id}`)

export const getChatByIdHistory = (id: string) => axios.get(`http://localhost:5000/stimaGPT/histroy?id=${id}`)

export const getAnswer = (data: ChatRequest) => axios.post('http://localhost:5000/stimaGPT/chat', data)

export const createHistory = (data: HistoryRequest) => axios.post('http://localhost:5000/stimaGPT/history', data)