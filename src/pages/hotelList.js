import  React from "react"
import { hotels } from "../data/hotels"
import  Hotel  from "./hotel"
import {getHotels} from "../api/hotels"


const HotelList = (props) => {

    const getHotel = (id) => {
        const hotel = hotels.find((hotel) => hotel.id = id)
        if(props.isAuthenticated){
            //show page reservation
        }else{
            //show dialog registration
        }
    }
    
    return (
        <section className="HotelList">
            
            {props.hotelsApi.map((hotel) => {

                return (
                    
                    <Hotel {...hotel} key={hotel.id} getHotel={getHotel} />
                )
            })}


        </section>
    );
};
export default HotelList;