import { Box, Typography, Card, CardContent, CardMedia, Skeleton } from '@mui/material';
import { styled } from '@mui/material/styles';
import PropTypes from 'prop-types';

const MarqueeContainer = styled(Box)({
  width: '100%',
  overflow: 'hidden',
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
});

const MarqueeContent = styled(Box)({
  display: 'inline-block',
  padding: '0 2rem',
  animation: 'marquee 20s linear infinite',
  whiteSpace: 'nowrap',
  boxSizing: 'border-box',
  '@keyframes marquee': {
    '0%': {
      transform: 'translate(0, 0)',
    },
    '100%': {
      transform: 'translate(-50%, 0)',
    },
  },
});

const Marquee = ({ movies }) => {
  return (
    <MarqueeContainer>
      <MarqueeContent>
        {[...movies, ...movies].map((movie, index) => (
          <Card key={index} sx={{ display: 'inline-block', margin: '0 1rem', minWidth: 200 }}>
            {movie.poster ? (
              <CardMedia
                component="img"
                height="140"
                image={movie.poster}
                alt={movie.title}
              />
            ) : (
              <Skeleton variant="rectangular" width="100%" height={140} />
            )}
            <CardContent>
              <Typography variant="h6" component="div">
                {movie.title}
              </Typography>
              <Typography variant="body2" color="text.secondary">
                {movie.description}
              </Typography>
            </CardContent>
          </Card>
        ))}
      </MarqueeContent>
    </MarqueeContainer>
  );
};

Marquee.propTypes = {
  movies: PropTypes.array.isRequired,
};

export default Marquee;