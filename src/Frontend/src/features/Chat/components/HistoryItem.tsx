import { HiOutlineChat } from "react-icons/hi"
import { IoTrashBinOutline } from "react-icons/io5"

type HistoryItemProps = {
    name: string
    id: string
}

function HistoryItem({ name, id }: HistoryItemProps) {
    return (
        <div className="flex justify-between items-center px-2 py-2 rounded-2xl group hover:bg-yellow-200">
            <span className="text-white group-hover:text-black">
                <HiOutlineChat size="24"/>
            </span>
            <p className="max-w-full p-2 mx-2 truncate hover:cursor-pointer text-white group-hover:text-black">
                {name}
            </p>
            <button className="group-hover:text-black">
                <IoTrashBinOutline size="24" />
            </button>
        </div>
    )
}

export default HistoryItem