import { Outlet } from "react-router-dom"
import Footer from "./components/Footer";
import NavbarMobile from "./components/NavbarMobile";

function App() {
    return (
        <div>
            <NavbarMobile />
            <Outlet />
            <Footer />
        </div>
    )
}

export default App;
