import { BiPaperPlane } from "react-icons/bi"

function MessageWriter() {
  return (
    <div className="w-[80vw] h-24 absolute bottom-0 flex items-center justify-center">
        <div className="w-[65%] p-4 bg-zinc-950 flex justify-between items-center rounded-3xl">
            <input type="text" className="w-[90%] bg-zinc-950 text-white outline-none"/>
            <button className="w-8 h-8 rounded-xl bg-yellow-200 flex items-center justify-center">
                <BiPaperPlane size="24" color="black"/>
            </button>
        </div>
    </div>
  )
}

export default MessageWriter