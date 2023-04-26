import HistoryItem from "./HistoryItem"

function ListHistory() {
  return (
    <div className='w-full text-center h-[70vh]'>
        <button className="w-full flex items-center px-2 py-4 rounded-2xl bg-indigo-600">
          <span className="block w-10 h-10 flex justify-center items-center bg-indigo-500 text-white rounded-2xl font-black">
            +
          </span>
          <span className="ml-4">
            New Chat
          </span>
        </button>
        <div className="max-h-[50vh] my-1 overflow-y-scroll scroll">
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