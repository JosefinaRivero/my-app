import React from "react";
import Modal from "../components/modal";
import { useState } from "react";
import {
  useParams
  
  } from "react-router-dom";


function HotelReserva() {
    // We can use the `useParams` hook here to access
    // the dynamic pieces of the URL.
    let { id} = useParams();
  
    return (
      <div>
        <h3>ID: {id}</h3>
      </div>
    );
  }
export default HotelReserva;
