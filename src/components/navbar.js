import React from 'react';
import "../index.css"
import { Link
} from "react-router-dom";


const Navbar = () => {
  return (
  
    <div id="header" class="headerhotel">
        <div class="headerlogo"><img class="logo" src="./images/logo-sunset-hotels.png" alt="sunset hotels"/></div>
        <div class="navbar"><ul>
        <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/about">About</Link>
            </li>
            <li>
              <Link to="/users">Users</Link>
            </li>
          {/* <li><a href="#">HOME</a></li>
          <li><a href="#">CONTACT</a></li>
          <li><a href="#">AREA CLIENT</a></li> */}
          </ul></div>
    </div>
    
 
  );
};

export default Navbar;
