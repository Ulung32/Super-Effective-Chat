import {
  createBrowserRouter,
  Outlet,
  RouterProvider,
} from "react-router-dom";
import { Chat, Login, Questions, Register } from "./features";

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
    element: <Outlet/>,
    children: [
      {
        path: "chat",
        element: <Chat/>
      },
      {
        path: "questions",
        element: <Questions/>
      }
    ]
  },
])
function App() {
  return (
    <RouterProvider router={router}/>
  )
}

export default App
