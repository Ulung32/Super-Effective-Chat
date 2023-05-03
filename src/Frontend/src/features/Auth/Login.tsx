import { Link, useNavigate } from "react-router-dom"
import Textbox from "./components/Textbox"
import Button from "./components/Button"
import { useForm } from "react-hook-form"
import { UserRequest } from "./api"
import axios from "axios"
import { useAuthAction, useAuthStore } from "../../store"

function Login() {
  const {
    formState: { isDirty, isSubmitting },
    handleSubmit,
    register,
  } = useForm<UserRequest>()

  const navigate = useNavigate()
  const setId = useAuthAction().setId
  const setUsername = useAuthAction().setUsername

  const onSubmit = async (req: UserRequest) => {
    try {
      const res = await axios.get(`http://localhost:5000/stimaGPT/User?username=${req.username}&password=${req.password}`)
      setId(res.data._id)
      setUsername(res.data.UserName)
      navigate("/")
    }catch(err){
      console.log(err)
    }
  };

  return (
    <div className="w-full h-[100vh] overflow-hidden bg-zinc-900 flex flex-col items-center justify-center font-sono">
      <span className="block my-4 text-3xl font-bold text-white">Login</span>
      <form className="w-[400px] max-w-[80%]" onSubmit={handleSubmit(onSubmit)}>
        <Textbox label="username" className="my-4 p-4 focus:outline-white" name="username" register={register} required={true} />
        <Textbox label="password" type="password" className="my-4 focus:outline-white" name="password" register={register} required={true} />
        <div className="flex justify-center">
          <Button label="submit" type="submit" className="bg-yellow-200 mx-4 cursor-pointer" disabled={!isDirty} loading={isSubmitting} />
          <Link to="/auth/register">
            <Button label="register" className="bg-indigo-600 text-white" />
          </Link>
        </div>
      </form>
    </div>
  )
}

export default Login