import fondo from "../images/fondo.jpg";
import SolutionSection from "../components/SolutionSection";

const scrollDown = () => {
  window.scrollTo({
    top: 700,
    behavior: "smooth",
  });
};

function Homepage() {
  return (
    <div className="home-cont">
      <div className="header-logo" style={{ backgroundImage: `url(${fondo})` }}>
        <h1>Mi nombre es billy </h1>
        <h2>Soy desarrollador backend</h2>
        <div className="home-button">
          <button className="buttonM1" onClick={scrollDown}>
            Explorar
          </button>
        </div>
      </div>

      <div className="header-introduction">
        <div className="intro-text">
          <div className="textI">
            <p>
              Actualmente mi lenguaje principal es golang aunque tambien
              desarrollo frontend, me gusta planear y desarrollar soluciones
              simples para problemas complejos siempre con un enfoque de
              eficiencia de recursos y buenas practicas.{" "}
            </p>
          </div>
        </div>
      </div>

      <SolutionSection />
    </div>
  );
}

export default Homepage;
