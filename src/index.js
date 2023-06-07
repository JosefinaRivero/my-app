
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
<<<<<<< HEAD
       title:'hotel cancun',
=======
       title:'hotel best',
>>>>>>> bb9446b8541216d50f25913c5d38064cf6ab5572
       img:"./images/hotel4.jpg",
       id: 3,
   },
    {
<<<<<<< HEAD
       title:'hotel las vegas',
=======
       title:'hotel ocean',
>>>>>>> bb9446b8541216d50f25913c5d38064cf6ab5572
       img:"./images/hotel5.jpg",
       id: 4,
   },
   {
<<<<<<< HEAD
       title:'hotel hilton',
=======
       title:'hotel teh beach',
>>>>>>> bb9446b8541216d50f25913c5d38064cf6ab5572
       img:"./images/hotel6.jpg",
       id: 5,
   },
    {
<<<<<<< HEAD
       title:'hotel texas',
=======
       title:'hotel sun',
>>>>>>> bb9446b8541216d50f25913c5d38064cf6ab5572
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

