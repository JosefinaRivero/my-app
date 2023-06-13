import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css"
import { LoginForm } from "./components/Authorization/LoginForm"
import HotelList from "./pages/hotelList"
import SearchBar from "./components/search"
import NavBar from "./components/navbar"
import Footer from "./components/footer";

import {
   BrowserRouter as Router,
   Switch,
   Route,
} from "react-router-dom";



const HotellWebSite = () => {

   const [isAuthenticated, setIsAuthenticated] = React.useState(false);
   const [user, setUser] = React.useState(undefined);
   const [startDate, setStartDate] = React.useState(new Date("2014/02/08"));
   const [endDate, setEndDate] = React.useState(new Date("2014/02/10"));
   const [hotels, setHotels] = React.useState([]);
   return (
      <Router>
         <section>
            <NavBar />
            <Switch>
               <Route path="/about">
                  {/* <HotelList hotelsApi={hotels} /> */}
               </Route>
               <Route path="/hotel/:id">
                  {/* <SearchBar  startDate={startDate} endDate={endDate}
                  setStartDate={setStartDate} setEndDate={setEndDate} setHotels={setHotels} /> */}
               </Route>
               <Route path="/">
                  <SearchBar 
                  startDate={startDate} endDate={endDate}
                  setStartDate={setStartDate} setEndDate={setEndDate}
                   setHotels={setHotels} />
                  <HotelList hotelsApi={hotels} />
               </Route>
               <Route path="/hotels">
                  <SearchBar 
                  startDate={startDate} endDate={endDate}
                  setStartDate={setStartDate} setEndDate={setEndDate}
                   setHotels={setHotels} />
                  <HotelList hotelsApi={hotels} />
               </Route>
            </Switch>
            <Footer />
         </section>
      </Router>)

};

const root = ReactDOM.createRoot(document.getElementById('root'))

root.render(<HotellWebSite />);