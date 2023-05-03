import { useQuery } from "@tanstack/react-query"
import HistoryItem from "./HistoryItem"
import { getHistoryByIdUser } from "../api"
import CreateHistoryDialog from "./CreateHistoryDialog"
import { useEffect } from "react"
import { useChatStore } from "../../../store"

type ListHistoryProps = {
  idUSer: string
}

function ListHistory({idUSer}: ListHistoryProps) {
  const histories = useChatStore((state) => state.histories)

  const { data, refetch } = useQuery({
    queryKey: ["get-history-by-id-user", idUSer],
    queryFn: () => getHistoryByIdUser(idUSer)
  })

  useEffect(() => {
    refetch()
  }, [histories])

  return (
    <div className='w-full text-center h-[60vh] px-8 pt-4'>
        <CreateHistoryDialog trigger={
          <button className="w-full flex items-center px-2 py-4 rounded-2xl bg-indigo-600 hover:bg-violet-500 group">
            <span className="block w-8 h-8 flex justify-center items-center bg-indigo-500 text-white rounded-lg text-lg font-bold group-hover:bg-indigo-400">
              +
            </span>
            <span className="ml-4 text-lg">
              New Chat
            </span>
          </button>
        }/>
        <div className="max-h-[50vh] mt-1 overflow-y-scroll scroll">
          {
            data?.data?.map((el: any) => 
              <HistoryItem name={el.name} id={el._id}/>)
          }
        </div>
    </div>
  )
}

export default ListHistory