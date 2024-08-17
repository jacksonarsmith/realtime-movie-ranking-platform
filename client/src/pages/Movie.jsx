import { Box } from "@mui/material"
import Navbar from "../components/navbar/Navbar"
import MovieItem from "../components/movie-item/MovieItem"
import Footer from "../components/footer/Footer"

const Movie = () => {
  return (
    <Box>
        <Navbar />
        <MovieItem />
        <Footer />
    </Box>
  )
}

export default Movie