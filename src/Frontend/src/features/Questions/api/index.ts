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

export const createQnA = (data: QnARequest) => axios.post(`http://localhost:5000/stimaGPT/QnA`, data)

export const deleteQnA = (id: string) => axios.delete(`http://localhost:5000/stimaGPT/QnA?id=${id}`)

export const getAllQnA = () => axios.get("http://localhost:5000/stimaGPT")

