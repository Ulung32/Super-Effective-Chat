import axios from "axios";

export type ChatRequest = {
    historyId: string 
    query: string
    method: string
}

export type HistoryRequest = {
    userId: string,
    name: string
}

export const getHistoryByIdUser = (id: string) => axios.get(`https://secchatbot-production.up.railway.app/stimaGPT/history?userID=${id}`)

export const getChatByIdHistory = (id: string) => axios.get(`https://secchatbot-production.up.railway.app/stimaGPT/chat?HisID=${id}`)

export const getAnswer = (data: ChatRequest) => axios.post('https://secchatbot-production.up.railway.app/stimaGPT/chat', data)

export const createHistory = (data: HistoryRequest) => axios.post('https://secchatbot-production.up.railway.app/stimaGPT/history', data)

export const deleteHistory = (id: string) => axios.delete(`https://secchatbot-production.up.railway.app/stimaGPT/history?historyID=${id}`)