import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css"
import {LoginForm} from "./components/Authorization/LoginForm"
import HotelList from "./pages/hotelList"




   
 const HotellWebSite=() => {
 
    const [isAuthenticated, setIsAuthenticated] = React.useState(false);
    const [user, setUser] = React.useState(undefined);
    return (<section> <HotelList /></section>)
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