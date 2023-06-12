import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css"
import {LoginForm} from "./components/Authorization/LoginForm"
import HotelList from "./pages/hotelList"
import DatePicker from "./components/search"
import NavBar from "./components/navbar"
import Footer from "./components/footer";




   
 const HotellWebSite=() => {
 
    const [isAuthenticated, setIsAuthenticated] = React.useState(false);
    const [user, setUser] = React.useState(undefined);
    return (<section>
      <NavBar/>
      <DatePicker/>
       <HotelList />
       <Footer/>
       </section>)
    //    return (<section>
    //     {!isAuthenticated ? 
    //     <LoginForm setIsAuthenticated={setIsAuthenticated} setUser={setUser}></LoginForm> :
    //     user.isAdmin ? <>Load admin pages</> :
    //     <HotelList></HotelList>
    // }
    //    </section>
           
    //    );
   };

const root=ReactDOM.createRoot(document.getElementById('root'))

root.render(<HotellWebSite/>);