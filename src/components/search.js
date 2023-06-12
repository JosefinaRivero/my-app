import React, { useRef, useState } from 'react';
import "../index.css"


const DatePicker = () => {

  const [date, setDate] = useState('');
  const dateInputRef = useRef(null);

  const handleChange = (e) => {
    setDate(e.target.value);
  };


  return (
    <div id="datepick" class="datahotel">
        <div  class="datacontent"><h3>Reserva tu hotel</h3>
       
      <input
        type="date"
        onChange={handleChange}
        ref={dateInputRef}
      />  
      <input
      type="date"
      onChange={handleChange}
      ref={dateInputRef}
    />
   
   <button>ENVIA</button>
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
export default DatePicker;