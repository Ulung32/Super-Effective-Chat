import { useAuthStore } from "../../../store"
import DropDownAccount from "./DropDownAccount"

function Account() {
  const username = useAuthStore((state) => state.username)

  return (
    <div className="flex h-[10vh] bg-zinc-800 items-center justify-between w-full px-4">
      <div className="flex justify-start items-center">
        <div className="w-8 h-8 bg-black rounded-full"/>
        <span className="ml-3">{username}</span>
      </div>
      <DropDownAccount trigger={<button>
        ...
      </button>}/>
    </div>
  )
}

export default Account