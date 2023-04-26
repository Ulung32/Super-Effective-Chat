import MessageItem from "./MessageItem"

function ListMessage() {
  return (
    <div className="w- h-[100vh] overflow-hidden">
        <MessageItem message="Halo" isBot={false} time="15.30"/>
        <MessageItem message="Halo juga" isBot={true} time="15.30"/>
        <MessageItem message="Apakah ada yang bisa saya bantu?" isBot={true} time="15.30"/>
    </div>
  )
}

export default ListMessage