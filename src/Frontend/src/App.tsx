import {
  createBrowserRouter,
  Navigate,
  Outlet,
  RouterProvider,
  useLocation,
} from "react-router-dom";
import { Chat, Login, Questions, Register } from "./features";
import { useAuthStore } from "./store";
import Error from "./features/ErrorElement";

const router = createBrowserRouter([
  {
    path: "/auth",
    element: <Outlet/>,
    children: [
      {
        path: "login",
        element: <Login/>
      },
      {
        path: "register",
        element: <Register/>
      }
    ]
  },
  {
    path: "/",
    element: <Root/>,
    children: [
      {
        path: "chat",
        element: <Chat/>
      },
      {
        path: "questions",
        element: <Questions/>
      }
    ],
    errorElement: <Error/>
  }
])

function Root() {
  const location = useLocation()
  const authStore = useAuthStore()

  if(authStore.id == "" && authStore.username == ""){
    return <Navigate to="/auth/login"/>
  }

  if (location.pathname === "/") {
    return <Navigate to="/chat" />
  }

  return (
    <>
      <Outlet />
    </>
  )
}

function App() {
  return (
    <RouterProvider router={router}/>
  )
}

export default App
