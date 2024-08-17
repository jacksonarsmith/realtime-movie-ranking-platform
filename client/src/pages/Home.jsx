import { Box } from '@mui/material'
import Navbar from '../components/navbar/Navbar'
import Hero from '../components/hero/Hero'
import Footer from '../components/footer/Footer'

const Home = () => {
  return (
    <Box>
        <Navbar />
        <Hero />
        <Footer />
    </Box>
  )
}

export default Home