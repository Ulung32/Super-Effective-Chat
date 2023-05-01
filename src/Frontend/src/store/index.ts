import { create } from "zustand"

type AuthState = {
    id: string
    username: string

    actions: {
        setId: (id: string) => void
        setUsername: (username: string) => void
    }
}

export const useAuthStore = create<AuthState>((set) => ({
    id: "efe",
    username: "",

    actions: {
        setId(id){
            set({id})
        },
        setUsername(username){
            set({username})
        }
    }
}))


export function useAuthAction(){
    return useAuthStore((state) => state.actions)
}