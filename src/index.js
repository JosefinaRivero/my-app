
import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css"
import Hotel from "./pages/hotel";




const hotels=[
    {
       title:'Hotel Miami', 
       img:"./images/hotel1.jpg",
       id: 1,
   },
    {
       title:'hotel miami 2',
       img:"./images/hotel3.jpg",
       id: 2,
   },
    {
       title:'hotel best',
       img:"./images/hotel4.jpg",
       id: 3,
   },
    {
       title:'hotel ocean',
       img:"./images/hotel5.jpg",
       id: 4,
   },
   {
       title:'hotel teh beach',
       img:"./images/hotel6.jpg",
       id: 5,
   },
    {
       title:'hotel sun',
       img:"./images/hotel2.jpg",
       id: 6,
   
   },
   ];
   




const HotelList=() => {
 
 const getHotel = (id) => {
    const hotel =hotels.find((hotel)=> hotel.id=id)
     console.log (hotel);
}
    return (
        <section className="HotelList">
        
        {hotels.map((hotel)=>{
        
        return (
            <Hotel {...hotel} key={hotel.id} getHotel={getHotel}/>
        )
    })}

        
        </section>
    );
};





const root=ReactDOM.createRoot(document.getElementById('root'))

root.render(<HotelList/>);

