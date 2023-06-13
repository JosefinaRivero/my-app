import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import "../index.css"
import "react-datepicker/dist/react-datepicker.css";
import { getHotelsWithDate } from '../api/hotels';
import { useHistory,useParams } from "react-router-dom";
const SearchBar = (props) => {




  let history = useHistory();
  let { id } = useParams();
  return (
    <div id="datepick" class="datahotel">
      <div>Id : {id}</div>
      <div class="datacontent"><h3>Reserva tu hotel</h3>

        <DatePicker
          selected={props.startDate}
          onChange={(date) => props.setStartDate(date)}
          selectsStart
          startDate={props.startDate}
          endDate={props.endDate}
        />
        <DatePicker
          selected={props.endDate}
          onChange={(date) => props.setEndDate(date)}
          selectsEnd
          startDate={props.startDate}
          endDate={props.endDate}
          minDate={props.startDate}
        />

        <button className='button' onClick={() => {
          getHotelsWithDate(props.startDate, props.endDate).then(result => {
            props.setHotels(result)
          })
          history.push("/users");
        }}>ENVIA</button>
      </div>

    </div>
  );
};

/*
const DatePicker = () =>{
  const [startDate, setStartDate] = useState(new Date());
  const [endDate, setEndDate] = useState(null);
  return (
    <>
      <DatePicker
        selected={startDate}
        onChange={(date) => setStartDate(date)}
        selectsStart
        startDate={startDate}
        endDate={endDate}
      />
      <DatePicker
        selected={endDate}
        onChange={(date) => setEndDate(date)}
        selectsEnd
        startDate={startDate}
        endDate={endDate}
        minDate={startDate}
      />
    </>
  );
};*/
export default SearchBar;