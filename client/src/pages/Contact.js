import commandlogo from "../images/command.svg"

function Contact() {
    return (
        <div className="contactMain">
            <div className="contList">
            <li className="solutions1">
                        <div className="solution-details1">
                        <img src={commandlogo} className="logoContact" alt="logo" />
                            <div className="solution-text">
                                <h3 className="solution-title">Contacto </h3>
                                <p>
                                    Actualmente en busca de empleo:
                                </p>
                            </div>
                        </div>
                        <div className="solution-footer1">
                            <div className="solutionFooter-cont">
                                <div className="solutionFooter-header">
                                    Contacto
                                </div>
                                <div className="solutionFooter-list">
                                    <ul>
                                        <li>Correo: bildev1958@gmail.com </li>
                                        <li>WhatsApp: 7294574940</li>
                                        <li>Direaccion: Chapultepec 104 Chapultepec Estado de mexico</li>
                                    </ul>
                                </div>
                            </div>
                        </div>
                    </li>
            </div>


        </div>
    )
}

export default Contact
