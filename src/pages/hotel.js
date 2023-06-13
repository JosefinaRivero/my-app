import React from "react";
import Modal from "../components/modal";
import { useState } from "react";
import styles from "./hotel.module.css";
import {addReservation} from "../api/login"


const Hotel = (props) => {
    const {img, title, getHotel, id}=props;
   // console.log(props);
   const [isOpen, setIsOpen] = useState(false);
   const [result, setResult] = useState(undefined);
    
return(
    
    <article className="Hotel">
      
       <img src={img} alt={title} ></img> 
       <h2>{title.toUpperCase()}</h2>
       <button className={styles.primaryBtn} onClick={() => setIsOpen(true)}>
       Informacion
      </button >
      {isOpen && <Modal setIsOpen={setIsOpen} />}
      
      <button onClick={() => {
        addReservation({}).then(result => {
        setResult(true)
      }).catch(result => {
        setResult(false)
      })}}>Confirmar reservacion</button>

      {result !== undefined && <div>
        <p>Resultado : {result ? <div>OK</div> :<div>Error</div>}</p>
        </div>}
    </article>
   
)

}

export default Hotel;
// <button onClick={getHotel(id)}>Informacion</button>