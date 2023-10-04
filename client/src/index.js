
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import Homepage from "./pages/Homepage";
import Blog from "./pages/Blog";
import Proyects from "./pages/Proyects";
import Contact from "./pages/Contact";

const router = createBrowserRouter([
	{
		path: "/",
		element: <App />,

		children: [
			{
				path: "",
				element: <Homepage />,
			},
			{
				path: "blog",
				element: <Blog />,
			},
			{
				path: "proyects",
				element: <Proyects />,
			},
			{
				path: "contact",
				element: <Contact />,
			},
		],
	},
]);

const root = ReactDOM.createRoot(document.getElementById("root"));
root.render(
	<React.StrictMode>
		<RouterProvider router={router} />
	</React.StrictMode>
);
