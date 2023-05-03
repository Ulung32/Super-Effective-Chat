export type History = {
    id: string,
    idUser: string,
    name: string
}

export type Chat = {
    id: string
    idHistory: string
    message: string
    isBot: boolean
}

export type User = {
    id: string,
    username: string,
    password: string
}

export type QnA = {
    id: string
    question: string
    answer: string
}
