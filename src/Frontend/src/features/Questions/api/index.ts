export type QnA = {
    id: string
    question: string
    answer: string
}

export type QnARequest = {
    question: string
    answer: string
}

export const createQnA = (data: UserRequest) => axios.post(`http://localhost:1323/stimaGPT/User`, data)

