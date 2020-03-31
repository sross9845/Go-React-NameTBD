import React, {useState, useEffect} from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios'

function App() {
    const [data, setData] = useState([])
    const [link, setLink] = useState('')
    const [name, setName] = useState('')
    const [author, setAuthor] = useState('')
    const [description, setDescription] = useState('')
  
    useEffect(() => {
      axios.get("http://localhost:3001/resources")
      .then(resources => {
        console.log(resources.data)
        setData(resources.data)
      }).then(data => {console.log(data)})
    },[])
    
    var submit = (e) => {
      e.preventDefault()
      axios.post("http://localhost:3001/resources", {
        Link: link,
        Name: name,
        Author: author,
        Description: description
      })
      
    }
    var names = data.map(data => {

    return( 
            <>
            <p>Link: {data.Link}</p>
            <p>Author: {data.Author}</p>
            <p>Decription: {data.Description}</p>
            <hr />
            </>
      )
    })
    return (
        <>
        <h1>Test App!</h1>
        <p>{names}</p>
        <form onSubmit={submit}>
          Link: <input type="string" name="link" onChange={e => setLink(e.target.value)} value={link}/>
          <br />
          Name: <input type="string" name="name" onChange={e => setName(e.target.value)} value={name}/>
          <br />
          Author: <input type="string" name="author" onChange={e => setAuthor(e.target.value)} value={author}/>
          <br />
          Decription: <input type="string" name="description" onChange={e => setDescription(e.target.value)} value={description}/>
          <br />
          <input type="submit" value="submit"/>
        </form>
        </>
    );
  }
 


export default App;
