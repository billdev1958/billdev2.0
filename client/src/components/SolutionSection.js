import weblogo from "../images/sphere-dark.svg";
import commandlogo from "../images/command.svg";
import gearlogo from "../images/gear.svg";

function SolutionSection() {
	return (
		<div className="section-solutions">
			<div className="section-grid">
				<li className="solutions">
					<div className="solution-details">
						<img
							src={commandlogo}
							className="weblogo"
							alt="weblogo"
						/>
						<div className="solution-text">
							<h3 className="solution-title">
								Desarrollo de
								software{" "}
							</h3>
							<p>
								Me gusta lo
								clasico tambien
								que es crear
								aplicaciones de
								escritorio
							</p>
						</div>
					</div>
					<div className="solution-footer">
						<div className="solutionFooter-cont">
							<div className="solutionFooter-header"></div>
							<div className="solutionFooter-list">
								<ul>
									<li>
										{" "}
										CLI{" "}
									</li>
									<li>
										Sistemas
										para
										windows
									</li>
									<li>
										Sistemas
										para
										linux
									</li>
								</ul>
							</div>
						</div>
					</div>
				</li>

				<li className="solutions">
					<div className="solution-details">
						<img
							src={weblogo}
							className="weblogo"
							alt="weblogo"
						/>
						<div className="solution-text">
							<h3 className="solution-title">
								Desarrollo web
							</h3>

							<p>
								Planeo y
								desarrollo
								soluciones web
								como:
							</p>
						</div>
					</div>
					<div className="solution-footer">
						<div className="solutionFooter-cont">
							<div className="solutionFooter-header"></div>
							<div className="solutionFooter-list">
								<ul>
									<li>
										{" "}
										Paginas
										web
										estaticas{" "}
									</li>
									<li>
										Aplicaciones
										web
									</li>
									<li>
										API
										REST
									</li>
									<li>
										Microservicios
									</li>
								</ul>
							</div>
						</div>
					</div>
				</li>

				<li className="solutions">
					<div className="solution-details">
						<img
							src={gearlogo}
							className="weblogo"
							alt="weblogo"
						/>
						<div className="solution-text">
							<h3 className="solution-title">
								Stack de
								tecnologias
							</h3>

							<p>
								Lo que uso
								actualmente
							</p>
						</div>
					</div>
					<div className="solution-footer">
						<div className="solutionFooter-cont">
							<div className="solutionFooter-header"></div>
							<div className="solutionFooter-list">
								<ul>
									<li>
										{" "}
										Golang{" "}
									</li>
									<li>
										Docker
									</li>
									<li>
										Postgres
									</li>
									<li>
										React
									</li>
									<li>
										Sass
									</li>
									<li>
										Tailwind
										css
									</li>
								</ul>
							</div>
						</div>
					</div>
				</li>
			</div>
		</div>
	);
}

export default SolutionSection;
