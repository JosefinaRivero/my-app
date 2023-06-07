import React from "react";




const Hotel = (props) => {
    const {img, title, getHotel, id}=props;
   // console.log(props);
    
return(
    
    <article className="Hotel">
      
       <img src={img} alt={title} ></img> 
       <h2>{title.toUpperCase()}</h2>
       <button onClick={getHotel(id)}>click me</button>
      
    </article>
   
)

}

export default Hotel