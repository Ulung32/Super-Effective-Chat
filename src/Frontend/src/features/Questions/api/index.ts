import axios from "axios"

export type QnA = {
    id: string
    question: string
    answer: string
}

export type QnARequest = {
    question: string
    answer: string
}

export const createQnA = (data: QnARequest) => axios.post(`https://secchatbot-production.up.railway.app/stimaGPT`, data)

export const deleteQnA = (id: string) => axios.delete(`https://secchatbot-production.up.railway.app/stimaGPT/delete?id=${id}`)

export const getAllQnA = () => axios.get("https://secchatbot-production.up.railway.app/stimaGPT")

