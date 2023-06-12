import React from "react";
import Modal from "../components/modal";
import { useState } from "react";
import styles from "./hotel.module.css";



const Hotel = (props) => {
    const {img, title, getHotel, id}=props;
   // console.log(props);
   const [isOpen, setIsOpen] = useState(false);
    
return(
    
    <article className="Hotel">
      
       <img src={img} alt={title} ></img> 
       <h2>{title.toUpperCase()}</h2>
 
       <button className={styles.primaryBtn} onClick={() => setIsOpen(true)}>
       Informacion
      </button>
      {isOpen && <Modal setIsOpen={setIsOpen} />}
    </article>
   
)

}

export default Hotel;
// <button onClick={getHotel(id)}>Informacion</button>