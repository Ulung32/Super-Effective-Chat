import { BsFillSendFill } from "react-icons/bs"

function MessageWriter() {
  return (
    <div className="w-full h-[20vh]">
        <div>
            <input type="text"/>
            <BsFillSendFill/>
        </div>
    </div>
  )
}

export default MessageWriter