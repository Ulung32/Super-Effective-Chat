function ListAlgo() {
  return (
    <div className='w-full h-[20vh]'>
        <div className="flex items-center justify-center">
          <input type="radio" id="hosting-small" name="hosting" value="hosting-small" className="hidden peer" required/>
          <label htmlFor="hosting-small" className="inline-flex items-center justify-between w-full p-4 text-gray-500 bg-white rounded-2xl cursor-pointer peer-checked:bg-slate-300 hover:text-gray-600 hover:bg-gray-100">                           
            <div className="w-full text-lg font-semibold">KMP</div>
          </label>
        </div>
        <div className="flex items-center justify-center">
          <input type="radio" id="hosting-big" name="hosting" value="hosting-small" className="hidden peer" required/>
          <label htmlFor="hosting-big" className="inline-flex items-center justify-between w-full p-4 text-gray-500 bg-white rounded-2xl cursor-pointer peer-checked:bg-slate-300 hover:text-gray-600 hover:bg-gray-100">                           
            <div className="w-full text-lg font-semibold">KMP</div>
          </label>
        </div>
    </div>
  )
}

export default ListAlgo