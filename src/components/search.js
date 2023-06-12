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

export default DatePicker;