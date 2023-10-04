import vimIcon from "../images/vimIcon.png"
import bill from "../images/bill.jpg"
import weblogo from "../images/sphere-dark.svg"
import commandlogo from "../images/command.svg"
import gearlogo from "../images/gear.svg"
function CardBlog() {

return (

 <div className="section-solutions">
                <div className="section-grid">
                    <li className="cardContent">
                        <div className="solution-details">
                          	 <div class="blogImage">
					<img src={vimIcon} className="imageBlog" alt="weblogo" />
				</div> 
                            <div className="articleText">
                                <h3 className="articleTitle">Desarrollo de software </h3>
                                <p>
                                    Me gusta lo clasico tambien que es crear aplicaciones de escritorio 
                                </p>
                            </div>
                        </div>
                        <div className="cardFooter">
                            <div className="solutionFooter-cont">
                                <div className="solutionFooter-header">

                                </div>
                                <div className="solutionFooter-list">
					<ul>
						<li>
							<img src={bill} className="avatar" alt="weblogo" />
	
						</li>
						<li>User</li>
						<li>Publicacion</li>
					</ul>
                                </div>
                            </div>
                        </div>
                    </li>



                </div>
            </div>
)
}

export default CardBlog;
