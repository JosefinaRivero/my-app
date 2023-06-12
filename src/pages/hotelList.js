import  React from "react"
import { hotels } from "../data/hotels"
import  Hotel  from "./hotel"
import {getHotels} from "../api/hotels"

const HotelList = () => {

    //example how load on init
    // React.useEffect(async () => {
    //     getHotels().then(result => {
    //         // load hotels in state 
    //     });
    // }, []) //empty array means that il is callaed only on fisrt render

    const getHotel = (id) => {
        const hotel = hotels.find((hotel) => hotel.id = id)
        console.log(hotel);
    }

   const searchHotels = (search, from, to) => {
        console.log(search)
        console.log(from)
        console.log(to)
    } 
    
    return (
        <section className="HotelList">
            
            {hotels.map((hotel) => {

                return (
                    
                    <Hotel {...hotel} key={hotel.id} getHotel={getHotel} />
                )
            })}


        </section>
    );
};
export default HotelList;