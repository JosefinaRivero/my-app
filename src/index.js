
import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css"


const hotels=[
 {
    title:'Hotel Miami', 
    img:"./images/hotel1.jpg",
    id: 1,
},
 {
    title:'hotel la',
    img:"./images/hotel3.jpg",
    id: 2,
},
 {
    title:'hotel la',
    img:"./images/hotel4.jpg",
    id: 3,
},
 {
    title:'hotel la',
    img:"./images/hotel5.jpg",
    id: 4,
},
{
    title:'hotel la',
    img:"./images/hotel6.jpg",
    id: 5,
},
 {
    title:'hotel la',
    img:"./images/hotel2.jpg",
    id: 6,

},
];



const HotelList=() => {
    return (
        <section className="HotelList">
        {hotels.map((hotel)=>{
        
            return (
                <Hotel {...hotel} key={hotel.id}/>
            )
        })}
        </section>
    );
};


const Hotel = (props) => {
    const {img, title}=props;
    console.log(props);
return(
    
    <article className="Hotel">
      
       <img src={img} alt={title} ></img> 
       <h2>{title.toUpperCase()}</h2>
      
    </article>
   
)

}



const root=ReactDOM.createRoot(document.getElementById('root'))

root.render(<HotelList/>);

