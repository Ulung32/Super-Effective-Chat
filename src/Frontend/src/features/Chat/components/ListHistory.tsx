import HistoryItem from "./HistoryItem"

function ListHistory() {
  return (
    <div className='w-full text-center h-[60vh] px-8 pt-4'>
        <button className="w-full flex items-center px-2 py-4 rounded-2xl bg-indigo-600 hover:bg-violet-500 group">
          <span className="block w-8 h-8 flex justify-center items-center bg-indigo-500 text-white rounded-lg text-lg font-bold group-hover:bg-indigo-400">
            +
          </span>
          <span className="ml-4 text-lg">
            New Chat
          </span>
        </button>
        <div className="max-h-[50vh] mt-1 overflow-y-scroll scroll">
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
            <HistoryItem name="cerj feriobf eriofn erf eriufkrjenf reifher i  h i" id="cerjh"/>
        </div>
    </div>
  )
}

export default ListHistory