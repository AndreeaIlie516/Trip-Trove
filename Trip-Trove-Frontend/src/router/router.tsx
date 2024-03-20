import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Home } from "../pages/Home";
import { AllDestinations } from "../pages/AllDestinations";
import { Destination } from "../pages/Destination";
import { AddDestination } from "../components/AddDestination";
import { UpdateDestination } from "../components/UpdateDestination";

const routes = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/destinations",
    Component: AllDestinations,
  },
  {
    path: "/destination/:id",
    element: <Destination />,
  },
  {
    path: "/destinations/update/:id",
    Component: UpdateDestination,
  },
  {
    path: "/destinations/add",
    Component: AddDestination,
  },
]);
export default function RoutesProvider() {
  return <RouterProvider router={routes} />;
}
