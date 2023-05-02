import { useQuery } from "@tanstack/react-query"
import MessageItem from "./MessageItem"
import { getChatByIdHistory } from "../api"
import { useChatAction, useChatStore } from "../../../store"
import { useEffect } from "react"

type ListMessageProps = {
  idHistory: string
}

function ListMessage({ idHistory }: ListMessageProps) {
  const chats = useChatStore((state) => state.chats)
  const setChats = useChatAction().setChats

  const { data } = useQuery({
    queryKey: ["get-chat-by-id-history", idHistory],
    queryFn: () => getChatByIdHistory(idHistory)
  })

  useEffect(() => {
    if(data){
      setChats(data.data)
    }
  }, [])

  return (
    <div className="w- h-[100vh] overflow-hidden">
        <MessageItem message="Halo" isBot={false} time="15.30"/>
        <MessageItem message="Halo juga" isBot={true} time="15.30"/>
        <MessageItem message="Apakah ada yang bisa saya bantu?" isBot={true} time="15.30"/>
    </div>
  )
}

export default ListMessage