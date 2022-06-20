import logo from './logo.svg';
import './App.css';
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import NavbarComponent from "./components/NavbarComponent"
import Buku from './pages/Buku';
import  Footer from './components/Footer';
import { Route, Routes } from "react-router-dom"
import Pinjam_buku from './pages/Pinjam_buku';

function App() {
  return (
    <>
    <NavbarComponent />
    <br/>
      <div classname="App">
        <Routes>
          <Route path="/" element={<Buku imgSrc={require('./greybg.jpg')} title="Card Title" text="Some quick example text to build on the card title and make up the
              bulk of the card's content."/>} />
          <Route path="/pinjamBuku" element={<Pinjam_buku />}/>
        </Routes>
      </div>
    <Footer />
    </>
  );
}

export default App;
