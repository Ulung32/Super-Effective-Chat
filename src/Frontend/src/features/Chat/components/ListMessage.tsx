import { useQuery } from "@tanstack/react-query"
import MessageItem from "./MessageItem"
import { getChatByIdHistory } from "../api"
import { useChatAction, useChatStore } from "../../../store"
import { useEffect } from "react"
import { Chat } from "../../../models"

type ListMessageProps = {
  idHistory: string
}

function ListMessage({ idHistory }: ListMessageProps) {
  const chats = useChatStore((state) => state.chats)
  const setChats = useChatAction().setChats

  const { data, refetch } = useQuery({
    queryKey: ["get-chat-by-id-history", idHistory],
    queryFn: () => getChatByIdHistory(idHistory)
  })

  useEffect(() => {
    refetch()
  }, [chats])

  return (
    <div className="w- h-[100vh] overflow-y-scroll">
      {
        data?.data.map((msg: any) => (
          <MessageItem message={msg.chat} time="12.30" isBot={msg.isbot} />
        ))
      }
    </div>
  )
}

export default ListMessage