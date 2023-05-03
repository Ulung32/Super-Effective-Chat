import clsx from "clsx"
import { useChatAction, useChatStore } from "../../../store"
import { HiOutlineChat } from "react-icons/hi"
import { IoTrashBinOutline } from "react-icons/io5"

type HistoryItemProps = {
    name: string
    id: string
}

function HistoryItem({ name, id }: HistoryItemProps) {
    const selectedId = useChatStore((state) => state.idHistorySelected)
    const setSelectedId = useChatAction().setIdHistorySelected

    return (
        <div className={clsx("flex justify-start items-center px-2 py-2 rounded-2xl group hover:bg-yellow-200", { "bg-yellow-200 text-black": id==selectedId, "text-white" : id!=selectedId })} onClick={() => setSelectedId(id)}>
            <span className={clsx("group-hover:text-black", { "text-black" : id==selectedId})}>
                <HiOutlineChat size="24"/>
            </span>
            <p className={clsx("max-w-full p-2 mx-2 truncate hover:cursor-pointer group-hover:text-black", { "text-black" : id==selectedId})}>
                {name}
            </p>
            {
                selectedId == id &&
                <button className="text-black">
                    <IoTrashBinOutline size="24" />
                </button>
            }
        </div>
    )
}

export default HistoryItem