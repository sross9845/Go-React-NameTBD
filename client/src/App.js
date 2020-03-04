import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios'

function App() {
    const [links, setLink] = useState([])
  
    useEffect(() => {
      axios.get("http://localhost:3001/resources")
      .then(resources => {
        console.log(resources.data)
        setLink(resources.data)
      }).then(data => {console.log(links)})
    },[])
    
    var names = links.map(link => {

    return( 
            <>
            <p>Link: {link.Link}</p>
            <p>Author: {link.Author}</p>
            <hr />
            </>
      )
    })
    return (
        <p>{names}</p>
    );
  }
 


export default App;
