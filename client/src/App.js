import "./App.css"
import { Route, Routes } from "react-router-dom";
import Home from "./page/Home"
import Blog from "./page/Blog"
import Header from "./components/layouts/Header"
import Footer from "./components/layouts/Footer"
import Add from "./page/Add"
import Edit from "./page/Edit"
import Delete from "./page/Delete"

function App() {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/blog/:id" element={<Blog />} />
        <Route path="/add" element={<Add />} />
        <Route path="/edit/:id" element={<Edit />} />
        <Route path="/delete/:id" element={<Delete />} />
      </Routes>
      <Footer />
    </>
  );
}

export default App;
