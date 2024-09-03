import { useEffect, useState } from 'react';
import { Box, Typography, Button, Divider } from '@mui/material'
import Marquee from '../../utils/marquee/Marquee';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Hero = () => {

  const [featuredMovies, setFeaturedMovies] = useState([]);


    useEffect(() => {
      const fetchFeaturedMovieList = async () => {
        try {
          const response = await axios.get(`${import.meta.env.VITE_DEV_API}/movies/featured`);
          setFeaturedMovies(response.data);
        } catch (error) {
          console.error(error);
        }
      }

      fetchFeaturedMovieList();
    }, [])

    return (
      <Box sx={{ display: 'flex',  flexDirection: 'column', justifyContent: 'center', alignItems: 'center', m: 2, p: 2, height: '100vh' }}>
          <Typography variant="h1">
              Our Newest Movies
          </Typography>
          <Divider sx={{ mb: 2, width: '100%' }}/>
          <Typography variant="body1">
              Explore the hottest new releases and see what&apos;s popular among movie lovers.
          </Typography>
          <Button variant="contained" sx={{ m: 2 }}>
              <Link to="/movies" style={{ textDecoration: 'none' }}>
                Browse Movies
              </Link>
          </Button>
          <Box 
              sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                width: '75vw',
                boxShadow: 3,
                borderRadius: 3,
                pt: 2,
                pb: 2,
                m: 2
              }}
          >
            <Marquee movies={featuredMovies}/>
          </Box>
      </Box>
    )
}

export default Hero