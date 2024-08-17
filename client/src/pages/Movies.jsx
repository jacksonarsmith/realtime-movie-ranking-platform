import { Box } from "@mui/material"
import Navbar from "../components/navbar/Navbar"
import MovieList from "../components/movie-list/MovieList"
import Footer from "../components/footer/Footer"

const Movies = () => {
  return (
    <Box>
        <Navbar />
        <MovieList />
        <Footer />
    </Box>
  )
}

export default Movies