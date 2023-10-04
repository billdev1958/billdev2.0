import githubLogo from "../images/github.svg"
function Footer() {
    return (
        <div className="footer-m">
            <div className="footer-c">
                <div className="logoF">
                </div>
                <div>
                    <h3>Billdev</h3>
                </div>
            </div>
            <div className="footer-list">
                    <div className="footer-link">
                        <a href="https://github.com/billdev1958"> <img src={githubLogo} className="github-logo" alt="logo" /></a>
                    </div>
                </div>
        </div>
    )
}

export default Footer
