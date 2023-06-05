
import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css"


const HotelList=() => {
    return (
        <section className="HotelList">
            <Hotel title={firstHotel.title} img={firstHotel.img}/>
            <Hotel title={secondHotel.title} img={secondHotel.img}/>
            <Hotel title={tercerHotel.title} img={tercerHotel.img}/>
            <Hotel title={cuartoHotel.title} img={cuartoHotel.img}/>
            <Hotel title={quintoHotel.title} img={quintoHotel.img}/>
            <Hotel title={sextoHotel.title} img={sextoHotel.img}/>
           
        </section>
    )
}

const firstHotel = {
    title:'Hotel Miami', 
    img:"./images/hotel1.jpg",
}
const tercerHotel = {
    title:'hotel la',
    img:"./images/hotel3.jpg",
}
const cuartoHotel = {
    title:'hotel la',
    img:"./images/hotel4.jpg",
}
const quintoHotel = {
    title:'hotel la',
    img:"./images/hotel5.jpg",
}
const sextoHotel = {
    title:'hotel la',
    img:"./images/hotel6.jpg",
}
const secondHotel = {
    title:'hotel la',
    img:"./images/hotel2.jpg",
}




const Hotel = (props) => {
    console.log(props);
    const {img, title}=props;
return(
    <article className="Hotel">
       <img src={img} alt={title} ></img> 
       <h3>{title.toUpperCase()}</h3>
    </article>
   
)

}



const root=ReactDOM.createRoot(document.getElementById('root'))

root.render(<HotelList/>);

