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

export const createQnA = (data: QnARequest) => axios.post(`http://localhost:1323/stimaGPT/QnA`, data)

export const deleteQnA = (id: string) => axios.delete(`http://localhost:1323/stimaGPT/QnA?id=${id}`)

