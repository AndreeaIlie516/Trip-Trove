import { Home } from "./pages/Home";
import { DestinationPage } from "./pages/Destination";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./App.css";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/destination/:id",
      element: <DestinationPage />,
    },
  ]);
  return <RouterProvider router={router} />;
}

export default App;
