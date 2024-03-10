import {
    createBrowserRouter,
    RouterProvider,
  } from "react-router-dom";
import { Home } from "../pages/Home";
import { Destination } from "../pages/Destination";

const routes = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/destination/:id",
      element: <Destination />,
    },
  ]);
  export default function RoutesProvider() {
    return <RouterProvider router={routes} />
  }