import { useQuery } from "@tanstack/react-query"
import MessageItem from "./MessageItem"
import { getChatByIdHistory } from "../api"
import { useChatStore } from "../../../store"
import { useEffect } from "react"

type ListMessageProps = {
  idHistory: string
}

function ListMessage({ idHistory }: ListMessageProps) {
  const chats = useChatStore((state) => state.chats)

  const { data, refetch, isLoading } = useQuery({
    queryKey: ["get-chat-by-id-history", idHistory],
    queryFn: () => getChatByIdHistory(idHistory)
  })

  useEffect(() => {
    refetch()
  }, [chats])

  return (
    <div className="h-[100vh] overflow-y-scroll pb-20">
      { isLoading ? 
        <p>ciuehiufhe</p>
        :
        data?.data.map((msg: any) => (
          <MessageItem message={msg.chat} isBot={msg.isbot} />
        ))
      }
    </div>
  )
}

export default ListMessage