import { FiCheck } from "react-icons/fi"

function ListAlgo() {
  return (
    <div className='w-full h-[27vh] bg-zinc-800 mt-6 px-8 py-4'>
        <div className="flex items-center justify-center my-2">
          <input type="radio" id="kmb" name="hosting" value="kmb" className="hidden peer/kmb" required/>
          <label htmlFor="kmb" className="w-full px-2 py-3 text-gray-500 bg-black rounded-2xl cursor-pointer hidden peer-checked/kmb:block">
            <div className="flex items-center">
              <div className="w-6 h-6 rounded-lg bg-yellow-200 flex items-center justify-center">
                <FiCheck color="black"/>
              </div>
              <div className="ml-4 text-lg font-semibold">KMP</div>
            </div>
          </label>
          <label htmlFor="kmb" className="w-full px-2 py-3 text-gray-500 bg-zinc-900 rounded-2xl cursor-pointer block peer-checked/kmb:hidden hover:bg-slate-800">
            <div className="flex items-center">
              <div className="w-6 h-6 rounded-lg border border-yellow-200 flex items-center justify-center">
              </div>
              <div className="ml-4 text-lg font-semibold">KMP</div>
            </div>
          </label>
        </div>
        <div className="flex items-center justify-center my-2">
          <input type="radio" id="bm" name="hosting" value="bm" className="hidden peer/bm" required/>
          <label htmlFor="bm" className="w-full px-2 py-3 text-gray-500 bg-black rounded-2xl cursor-pointer hidden peer-checked/bm:block">
            <div className="flex items-center">
              <div className="w-6 h-6 rounded-lg bg-yellow-200 flex items-center justify-center">
                <FiCheck color="black"/>
              </div>
              <div className="ml-4 text-lg font-semibold">BM</div>
            </div>
          </label>
          <label htmlFor="bm" className="w-full px-2 py-3 text-gray-500 bg-zinc-900 rounded-2xl cursor-pointer block peer-checked/bm:hidden hover:bg-slate-800">
            <div className="flex items-center">
              <div className="w-6 h-6 rounded-lg border border-yellow-200 flex items-center justify-center">
              </div>
              <div className="ml-4 text-lg font-semibold">BM</div>
            </div>
          </label>
        </div>
    </div>
  )
}

export default ListAlgo